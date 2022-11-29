package labeler

import (
	"context"
	"encoding/json"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
)

type PatchStringValue struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value"`
}

type PatchNodeParam struct {
	Node         string                 `json:"node"`
	OperatorType string                 `json:"operator_type"`
	OperatorPath string                 `json:"operator_path"`
	OperatorData map[string]interface{} `json:"operator_data"`
}

type PatchPodParam struct {
	Namespace    string                 `json:"namespace"`
	Pod          string                 `json:"pod"`
	OperatorType string                 `json:"operator_type"`
	OperatorPath string                 `json:"operator_path"`
	OperatorData map[string]interface{} `json:"operator_data"`
}

func ListNodes(clientset *kubernetes.Clientset) (list *v1.NodeList, err error) {
	nodes, err := clientset.CoreV1().Nodes().List(
		context.TODO(),
		metav1.ListOptions{},
	)
	if err != nil {
		return nil, err
	}
	return nodes, nil
}

func PatchNode(clientset *kubernetes.Clientset, param PatchNodeParam) (result *v1.Node, err error) {
	coreV1 := clientset.CoreV1()
	node := param.Node

	operatorData := param.OperatorData
	operatorType := param.OperatorType
	operatorPath := param.OperatorPath

	var payloads []interface{}

	for key, value := range operatorData {
		payload := PatchStringValue{
			Op:    operatorType,
			Path:  operatorPath + key,
			Value: value,
		}

		payloads = append(payloads, payload)
	}

	payloadBytes, _ := json.Marshal(payloads)

	result, err = coreV1.Nodes().Patch(context.TODO(), node, types.JSONPatchType, payloadBytes, metav1.PatchOptions{})
	if err != nil {
		return nil, err
	}

	return result, err
}

func PatchPod(clientset *kubernetes.Clientset, param PatchPodParam) (result *v1.Pod, err error) {
	coreV1 := clientset.CoreV1()

	operatorData := param.OperatorData
	operatorType := param.OperatorType
	operatorPath := param.OperatorPath

	var payloads []interface{}

	for key, value := range operatorData {
		payload := PatchStringValue{
			Op:    operatorType,
			Path:  operatorPath + key,
			Value: value,
		}

		payloads = append(payloads, payload)

	}

	payloadBytes, _ := json.Marshal(payloads)

	result, err = coreV1.Pods(param.Namespace).Patch(context.TODO(), param.Pod, types.JSONPatchType, payloadBytes, metav1.PatchOptions{})
	if err != nil {
		return nil, err
	}

	return result, err
}
