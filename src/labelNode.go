package labeler

import (
	fmt "fmt"
	"log"

	labeler "github.com/dimgatz98/labeler/src/server"

	"golang.org/x/net/context"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"strings"
)

type Server struct {
}

func (s *Server) LabelNode(ctx context.Context, label *labeler.Label) (*labeler.NodeInfo, error) {
	log.Printf("Received label from client: %v", label)

	slice := strings.Split(label.Label, ":")
	if len(slice) > 1 || len(slice) < 1 {
		return nil, fmt.Errorf("Invalid label")
	}

	config, err := clientcmd.BuildConfigFromFlags("", "/home/dimitris/.kube/config")
	if err != nil {
		return nil, err
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

	fmt.Println(params)
	return &labeler.NodeInfo{Info: fmt.Sprintf("NodeInfo:\n%v", res)}, nil
}
