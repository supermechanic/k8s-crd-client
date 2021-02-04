package seldon

import (
	"context"
	"fmt"
	"model-service/client"
	"model-service/config"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

//ServerType supported seldon server
//go:generate stringer -type=ServerType
type ServerType int32

const (
	MLFLOW_SERVER ServerType = iota
	SKLEARN_SERVER
	TENSORFLOW_SERVER
	PYTORCH_SERVER
	TRITON_SERVER
	XGBOOST_SERVER
)

//SetSpecName ///
func SetSpecName(obj *unstructured.Unstructured, name string) {
	if len(name) == 0 {
		unstructured.RemoveNestedField(obj.Object, "spec", "name")
	}
	unstructured.SetNestedField(obj.Object, name)
}

//SetSpecAnnonation //
func SetSpecAnnonation(obj *unstructured.Unstructured, annotations map[string]string) {
	if annotations == nil {
		unstructured.RemoveNestedField(obj.Object, "spec", "annotations")
		return
	}
	unstructured.SetNestedStringMap(obj.Object, annotations, "spec", "annotations")
}

//SetSpecPredictors ///
func SetSpecPredictors(obj *unstructured.Unstructured, predictors []map[string]interface{}) {
	if predictors == nil {
		unstructured.RemoveNestedField(obj.Object, "spec", "predictors")
		return
	}
	unstructured.SetNestedField(obj.Object, predictors, "spec", "predictors")
}

//CreateBasicSdep ///
func CreateBasicSdep(deployName string) (obj *unstructured.Unstructured) {
	obj.SetAPIVersion("machinelearning.seldon.io/v1")
	obj.SetKind("SeldonDeployment")
	obj.SetName(deployName)
	obj.SetNamespace(config.Config.Namespace)
	return
}

//CreateSeldonDeploy ///
func CreateSeldonDeployV1(deployName, modelPath string, params map[int32]int32, serverType ServerType) error {
	deploymentRes := schema.GroupVersionResource{Group: "machinelearning.seldon.io", Version: "v1", Resource: "seldondeployments"}
	deployment := CreateBasicSdep(deployName)
	annotations := map[string]string{
		"seldon.io/executor": "true",
	}
	SetSpecAnnonation(deployment, annotations)
	SetSpecName(deployment, deployName+"_pred")
	predictors := []map[string]interface{}{
		{
			"componentSpecs": nil,
			"name":           "default",
			"replicas":       1,
		},
	}
	graph := map[string]interface{}{
		"children":         nil,
		"implementation":   serverType.String(),
		"modelUri":         "s3://" + modelPath,
		"envSecretRefName": config.Config.SecretName,
		"name":             "predictor",
	}

	parameterInfos, err := GetServerParams(serverType, params)
	if err != nil {
		return err
	}
	parameters := []map[string]interface{}{}
	for _, v := range parameterInfos {
		parameter := map[string]interface{}{
			"name":  v.Name,
			"type":  v.Type,
			"value": v.Value,
		}
		parameters = append(parameters, parameter)
	}
	graph["parameters"] = parameters
	predictors[0]["graph"] = graph
	SetSpecPredictors(deployment, predictors)
	result, err := client.DynamicClient.Resource(deploymentRes).Namespace(config.Config.Namespace).Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	fmt.Println((result))
	return nil
}
