package main

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	corev1 "k8s.io/api/core/v1"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", ".kube/config")
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	podClient := clientset.CoreV1().Pods(v1.NamespaceDefault)

	pod := buildWorkerPod()
	newPod, err := podClient.Create(context.Background(), &pod, v1.CreateOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Printf(newPod.Name)

	list, err := podClient.List(context.Background(), v1.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, d := range list.Items {
		fmt.Printf("Namespace: %v \t Name: %v \t Status: %+v\n", d.Namespace, d.Name, d.Status.Phase)
	}

}

func buildWorkerPod() corev1.Pod {
	pod := corev1.Pod{
		TypeMeta: v1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Pod",
		},
		ObjectMeta: v1.ObjectMeta{
			GenerateName:               "test-",
			Labels: map[string]string{
				"app": "test",
			},
			Annotations: map[string]string{
				"app": "test",
			},
		},
		Spec:       corev1.PodSpec{
			Containers:  []corev1.Container{
				{
					Name:                     "test",
					Image:                    "nicolaka/netshoot",
					Command:                  []string{"/bin/ash", "-c", "trap : TERM INT; sleep infinity & wait"},
				},
			},
		},
	}
	return pod
}