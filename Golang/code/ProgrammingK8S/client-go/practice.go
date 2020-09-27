package client_go

import (
	"flag"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func main() {
	var kubeconfig = flag.String("kubeconfig", "~/.kube/config", "kubeconfig file")
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Fatal("")
	}
	clientset, err := kubernetes.New
}
