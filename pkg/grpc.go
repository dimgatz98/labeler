package labeler

import (
	fmt "fmt"
	"log"

	labeler "github.com/dimgatz98/labeler/pkg/server"

	"golang.org/x/net/context"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"strings"
)

type Server struct {
}

func (s *Server) LabelNode(ctx context.Context, label *labeler.NodeLabel) (*labeler.Info, error) {
	log.Printf("Received label from client: %v", label)

	slice := strings.Split(label.Label, ":")
	if len(slice) > 2 || len(slice) < 2 {
		return nil, fmt.Errorf("Invalid label")
	}

	var config *rest.Config
	var err error
	config, err = rest.InClusterConfig()
	if err != nil {
		config, err = clientcmd.BuildConfigFromFlags("", label.KubeConfig)
		if err != nil {
			return nil, err
		}
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	params := PatchNodeParam{
		Node:         label.Node,
		OperatorType: "replace",
		OperatorPath: "/metadata/labels/",
		OperatorData: map[string]interface{}{
			slice[0]: slice[1],
		},
	}
	res, err := PatchNode(clientset, params)
	if err != nil {
		return nil, err
	}

	log.Println(params)
	return &labeler.Info{Info: fmt.Sprintf("NodeInfo:\n%v", res)}, nil
}

func (s *Server) LabelPod(ctx context.Context, label *labeler.PodLabel) (*labeler.Info, error) {
	slice := strings.Split(label.Label, ":")
	if len(slice) > 2 || len(slice) < 2 {
		return nil, fmt.Errorf("Invalid label")
	}

	var config *rest.Config
	var err error
	config, err = rest.InClusterConfig()
	if err != nil {
		config, err = clientcmd.BuildConfigFromFlags("", label.KubeConfig)
		if err != nil {
			return nil, err
		}
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	params := PatchPodParam{
		Namespace:    label.Namespace,
		Pod:          label.Pod,
		OperatorType: "replace",
		OperatorPath: "/metadata/labels/",
		OperatorData: map[string]interface{}{
			slice[0]: slice[1],
		},
	}
	res, err := PatchPod(clientset, params)
	if err != nil {
		return nil, err
	}

	log.Println(params)
	return &labeler.Info{Info: fmt.Sprintf("NodeInfo:\n%v", res)}, nil
}
