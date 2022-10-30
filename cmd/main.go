package main

import (
	"fmt"
	"log"

	labeler "github.com/dimgatz98/labeler"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "/home/dimitris/.kube/config")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	_, err = labeler.ListNodes(clientset)
	if err != nil {
		panic(err)
	}

	params := labeler.PatchNodeParam{
		Node:         "k3d-k3s-default-server-0",
		OperatorType: "replace",
		OperatorPath: "/metadata/labels/",
		OperatorData: map[string]interface{}{
			"all-2g.12gb": "true",
		},
	}
	res, err := labeler.PatchNode(clientset, params)
	if err != nil {
		panic(err)

	}
	fmt.Println(res)
}
