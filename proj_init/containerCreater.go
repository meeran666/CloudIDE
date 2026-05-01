package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8syaml "k8s.io/apimachinery/pkg/util/yaml"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func getter_image_name(stack string) string {
	ide_image := map[string]string{
		"Node": "meeran666/ide_image_node",
		"Go":   "meeran666/ide_image_go",
	}

	return ide_image[stack]
}

// Read & parse YAML
func readAndParseKubeYaml(filePath, golet_id, ide_image, workspace_name string) ([]map[string]interface{}, error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	decoder := k8syaml.NewYAMLOrJSONDecoder(strings.NewReader(string(fileContent)), 4096)

	var manifests []map[string]interface{}

	for {
		var doc map[string]interface{}
		err := decoder.Decode(&doc)
		if err != nil {
			if err == io.EOF {
				return manifests, nil
			}
			fmt.Println(err)
			return manifests, err
		}

		docBytes, _ := yaml.Marshal(doc)
		docString := strings.ReplaceAll(string(docBytes), "service_name", golet_id)
		docString = strings.ReplaceAll(string(docString), "ide_image", ide_image)
		docString = strings.ReplaceAll(string(docString), "workspace_name", workspace_name)
		var finalDoc map[string]interface{}
		yaml.Unmarshal([]byte(docString), &finalDoc)

		manifests = append(manifests, finalDoc)
	}

}

func ContainerCreater(golet_id, stack, workspace_name string) error {
	// Load kubeconfig
	kubeconfig := os.Getenv("KUBECONFIG")
	if kubeconfig == "" {
		kubeconfig = os.ExpandEnv("$HOME/.kube/config")
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	namespace := "default"

	//call map function which tells you stack is mapped to which image name
	ide_image := getter_image_name(stack)
	manifests, err := readAndParseKubeYaml("./service.yaml", golet_id, ide_image, workspace_name)

	if err != nil {
		return err
	}
	for _, manifest := range manifests {
		kind, ok := manifest["kind"].(string)
		if !ok {
			continue
		}

		switch kind {

		case "Deployment":
			var deployment appsv1.Deployment
			// bytes, _ := yaml.Marshal(manifest)
			// yaml.Unmarshal(bytes, &deployment)
			jsonBytes, _ := json.Marshal(manifest)
			json.Unmarshal(jsonBytes, &deployment)

			_, err := clientset.AppsV1().
				Deployments(namespace).
				Create(context.TODO(), &deployment, metav1.CreateOptions{})

			if err != nil {
				return err

			}

		case "Service":
			var service corev1.Service
			jsonBytes, _ := json.Marshal(manifest)
			json.Unmarshal(jsonBytes, &service)

			_, err := clientset.CoreV1().
				Services(namespace).
				Create(context.TODO(), &service, metav1.CreateOptions{})

			if err != nil {
				return err

			}

		case "Ingress":
			var ingress networkingv1.Ingress
			jsonBytes, _ := json.Marshal(manifest)
			json.Unmarshal(jsonBytes, &ingress)

			_, err := clientset.NetworkingV1().
				Ingresses(namespace).
				Create(context.TODO(), &ingress, metav1.CreateOptions{})

			if err != nil {
				return err
			}

		default:
			log.Println("Unsupported kind:", kind)
		}
	}
	return nil
}
