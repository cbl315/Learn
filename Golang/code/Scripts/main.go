package main

import (
	"flag"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func parseArgs() (ns string, clientset *kubernetes.Clientset) {
	var kubeconfig *string
	kubeconfig = flag.String("kubeconfig", "~/.kube/config", "absolute path to the kubeconfig file")
	nsP := flag.String("ns", "", "namespace")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return *nsP, clientset
}

func main() {
	_, cs := parseArgs()
	deletePendingDeletionSvc(cs)
	//createPendingDeletionSvc(ns, cs)
}
