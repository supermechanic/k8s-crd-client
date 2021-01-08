package client

import (
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
)

var client dynamic.Interface
var err error

func init() {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	client, err = dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}
}
