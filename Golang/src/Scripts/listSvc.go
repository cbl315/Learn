package main

import (
	"context"
	"flag"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

const (
	targetFinalizer = "foregroundDeletion"
)

func main() {
	var kubeconfig *string
	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	namespaceList, err := clientset.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _, namespace := range namespaceList.Items {
		svcList, err := clientset.CoreV1().Services(namespace.Name).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		for _, svc := range svcList.Items {
			if len(svc.Finalizers) > 0 {
				for _, finalizer := range svc.Finalizers {
					if finalizer == targetFinalizer {
						log.Printf("Delete service %s for contain finalizer %s, namespace: %s", svc.Name, targetFinalizer, svc.Name)
						deletePolicy := metav1.DeletePropagationBackground
						err = clientset.CoreV1().Services(namespace.Name).Delete(context.Background(), svc.Name, metav1.DeleteOptions{PropagationPolicy: &deletePolicy})
						if err != nil && !errors.IsNotFound(err) {
							log.Printf("Delete service %s for contain finalizer %s, namespace: %s fail, err: %s", svc.Name, targetFinalizer, svc.Name, err.Error())
						}
						break
					}
				}
			}
		}
		log.Println("Finish clean namespace", namespace.Name)
	}
}
