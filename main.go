package main

import (
	"flag"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"k8s.io/client-go/util/homedir"
)

var kubeconfig *string

func init() {
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
}

//CreateSeldonDeploy //
func CreateSeldonDeploy() {

}
func publishModel(c *gin.Context) {

}
func main() {
	flag.Parse()
	r := gin.Default()
	r.POST("/api/v1/service/model/publish", publishModel)
}
