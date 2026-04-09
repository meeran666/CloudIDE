package main

import (
	"context"
	"fmt"
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

// Read & parse YAML
func readAndParseKubeYaml(filePath string, replId string) ([]map[string]interface{}, error) {
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
			break
		}

		docBytes, _ := yaml.Marshal(doc)
		docString := strings.ReplaceAll(string(docBytes), "service_name", replId)

		var finalDoc map[string]interface{}
		yaml.Unmarshal([]byte(docString), &finalDoc)

		fmt.Println(docString)

		manifests = append(manifests, finalDoc)
	}

	return manifests, nil
}

func ContainerHandler(golet_id string) {
	fmt.Println("value29")
	// Load kubeconfig
	kubeconfig := os.Getenv("KUBECONFIG")
	if kubeconfig == "" {
		kubeconfig = os.ExpandEnv("$HOME/.kube/config")
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	namespace := "default"

	manifests, err := readAndParseKubeYaml("./service.yaml", golet_id)
	if err != nil {
		// http.Error(w, err.Error(), 500)
		fmt.Println(err)
		return
	}

	for _, manifest := range manifests {
		kind, ok := manifest["kind"].(string)
		if !ok {
			continue
		}

		switch kind {

		case "Deployment":
			var deployment appsv1.Deployment
			bytes, _ := yaml.Marshal(manifest)
			yaml.Unmarshal(bytes, &deployment)

			_, err := clientset.AppsV1().
				Deployments(namespace).
				Create(context.TODO(), &deployment, metav1.CreateOptions{})

			if err != nil {
				log.Println("Deployment error:", err)
			}

		case "Service":
			var service corev1.Service
			bytes, _ := yaml.Marshal(manifest)
			yaml.Unmarshal(bytes, &service)

			_, err := clientset.CoreV1().
				Services(namespace).
				Create(context.TODO(), &service, metav1.CreateOptions{})

			if err != nil {
				log.Println("Service error:", err)
			}

		case "Ingress":
			var ingress networkingv1.Ingress
			bytes, _ := yaml.Marshal(manifest)
			yaml.Unmarshal(bytes, &ingress)

			_, err := clientset.NetworkingV1().
				Ingresses(namespace).
				Create(context.TODO(), &ingress, metav1.CreateOptions{})

			if err != nil {
				log.Println("Ingress error:", err)
			}

		default:
			log.Println("Unsupported kind:", kind)
		}
	}

}
