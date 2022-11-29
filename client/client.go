package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/akamensky/argparse"

	"google.golang.org/grpc"

	labeler "github.com/dimgatz98/labeler/pkg/server"
)

func main() {
	parser := argparse.NewParser("Labeler", "Label nodes/pods using grpc calls")

	nodeIP := parser.String("i", "node-ip", &argparse.Options{Default: "0.0.0.0", Help: "Node IP"})
	nodePort := parser.String("p", "port", &argparse.Options{Default: "9000", Help: "Node port"})
	label := parser.String("l", "label", &argparse.Options{Required: true, Help: "String used for label in the form `(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])?`:`(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])?`"})
	config := parser.String("c", "kube-config", &argparse.Options{Default: "", Help: "Kubernetes configuration path, only needed when running locally"})

	pod := parser.NewCommand("pod", "Label a pod")
	namespace := pod.String("n", "namespace", &argparse.Options{Default: "default", Help: "Namespace of the pod to be labeled"})
	podName := pod.String("o", "pod-name", &argparse.Options{Required: true, Help: "Name of the pod to be labeled"})

	node := parser.NewCommand("node", "Label a node")
	nodeName := node.String("n", "node-name", &argparse.Options{Required: true, Help: "Name of the node to be labeled"})

	// Parse input
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Printf(parser.Usage(err))
		os.Exit(1)
	}

	var conn *grpc.ClientConn
	conn, err = grpc.Dial(*nodeIP+":"+*nodePort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	c := labeler.NewLabelerServiceClient(conn)

	if *nodeName == "" {
		data := labeler.PodLabel{
			Namespace:  *namespace,
			Pod:        *podName,
			Label:      *label,
			KubeConfig: *config,
		}
		response, err := c.LabelPod(context.Background(), &data)
		if err != nil {
			log.Fatalf("Error when calling LabelNode: %v", err)
		}

		log.Printf("Response from server: %v", response)
	} else {
		data := labeler.NodeLabel{
			Node:       *nodeName,
			Label:      *label,
			KubeConfig: *config,
		}
		response, err := c.LabelNode(context.Background(), &data)
		if err != nil {
			log.Fatalf("Error when calling LabelNode: %v", err)
		}

		log.Printf("Response from server: %v", response)
	}
}
