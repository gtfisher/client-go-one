package main

import (
	"fmt"
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	fmt.Println("Home is" + os.Getenv("HOME"))
	pl, _ := listpods()
	fmt.Println(pl)
}

func listpods() ([]string, error) {
	var po []string
	config, err := clientcmd.BuildConfigFromFlags("", os.Getenv("HOME")+"/.kube/config")
	if err != nil {
		return po, err
	}
	cs, err := kubernetes.NewForConfig(config)
	if err != nil {
		return po, err
	}
	namespace := "kube-system"
	pods, err := cs.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	if err != nil {
		return po, err
	}
	for _, p := range pods.Items {
		fmt.Printf("The type of po %T\n", p)
		po = append(po, p.GetName()+"\n")
	}
	return po, nil
}
