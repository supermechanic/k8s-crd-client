package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var kubeconfig *string
var client dynamic.Interface

func init() {
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
}

//CreateSeldonDeploy //
func CreateSeldonDeploy(name, modelname, server string) {
	deploymentRes := schema.GroupVersionResource{Group: "machinelearning.seldon.io", Version: "v1", Resource: "deployments"}
	deployment := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "machinelearning.seldon.io/v1",
			"kind":       "SeldonDeployment",
			"metadata": map[string]interface{}{
				"name": name,
			},
			"spec": map[string]interface{}{
				"annotations": map[string]interface{}{
					"seldon.io/executor": "true",
				},
				"name": name + "-clf",
				"predictors": map[string]interface{}{
					"componentSpecs": nil,
					"graph": map[string]interface{}{
						"children":         nil,
						"implementation":   server,
						"modelUri":         "s3://" + modelname,
						"envSecretRefName": "envSecretRefName",
						"name":             "classifier",
					},
					"name":     "default",
					"replicas": 1,
				},
			},
		},
	}
	fmt.Println("Creating deployment...")
	result, err := client.Resource(deploymentRes).Namespace("mayxe").Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created deployment %q.\n", result.GetName())
}
func publishModel(c *gin.Context) {
	CreateSeldonDeploy("iris-sk-v1", "iris", "SKLEARN_SERVER")
}
func main() {
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	client, err = dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r.POST("/api/v1/service/model/publish", publishModel)
}
