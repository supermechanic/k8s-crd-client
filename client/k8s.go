package client

import (
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
)

//DynamicClient k8s dynanic client
var DynamicClient dynamic.Interface
var err error

func init() {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	DynamicClient, err = dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}
}
