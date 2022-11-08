package main

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "./kube-config")
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	podClient := clientset.CoreV1().Pods(v1.NamespaceDefault)

	list, err := podClient.List(context.Background(), v1.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, d := range list.Items {
		fmt.Printf("Namespace: %v \t Name: %v \t Status: %+v\n", d.Namespace, d.Name, d.Status.Phase)
	}

	discoveryClient, err := discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		panic(err)
	}

	_, apiResourceList, err := discoveryClient.ServerGroupsAndResources()
	if err !=  nil {
		panic(err)
	}
	for _, list := range apiResourceList {
		gv, err := schema.ParseGroupVersion(list.GroupVersion)
		if err != nil {
			panic(err)
		}
		for _, resource := range list.APIResources {
			fmt.Printf("%v \t %v\n", gv, resource)
		}
	}

}
