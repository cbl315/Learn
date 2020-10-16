package main

import (
	"flag"
	clientGo "github.com/hfutcbl/learn/golang/ProgrammingK8S/client-go"
	"github.com/hfutcbl/learn/golang/ProgrammingK8S/informer"
	"log"
)

func getClient() (kubecofnig string) {
	var kubeconfigArg = flag.String("kubeconfig", "~/.kube/config", "kubeconfig file")
	flag.Parse()
	return *kubeconfigArg
}

func main() {
	kubeconfig := getClient()
	clientGo.GetPod(kubeconfig)
	ns, podName := "namespace", "podName"
	pod, err := informer.GetPodByInformer(kubeconfig, ns, podName)
	if err != nil {
		log.Printf("get pod fail, err: %v", err)
		return
	}
	log.Printf("Get pod success, pod: %+v", pod)
}
