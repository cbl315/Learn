package client_go

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"runtime"
)

func GetPod(kubeconfig string) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}
	config.UserAgent = fmt.Sprintf("book-example/v1.0 (%s/%s) kubernetes/v1.0", runtime.GOOS, runtime.GOARCH)
	// switch protocol to protobuf from http
	config.AcceptContentTypes = "application/vnd.kubernetes.protobuf,application/json"
	config.ContentType = "application/vnd.kubernetes.protobuf"
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	pod, err := clientset.CoreV1().Pods("book").Get(context.Background(), "example", metav1.GetOptions{})
	log.Printf("%+v", pod)
}
