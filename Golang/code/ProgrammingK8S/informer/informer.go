package informer

import (
	podv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	v1 "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"time"
)

func GetInformer(kubeconfig string) (podInformer v1.PodInformer, err error) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return
	}
	informerFactory := informers.NewSharedInformerFactory(clientset, time.Second*30)
	podInformer = informerFactory.Core().V1().Pods()
	podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{}) // Add handlers
	informerFactory.Start(wait.NeverStop)
	informerFactory.WaitForCacheSync(wait.NeverStop)
	return
}

func GetPodByInformer(kubeconfig, ns, podName string) (pod *podv1.Pod, err error) {
	podInformer, err := GetInformer(kubeconfig)
	if err != nil {
		return
	}
	pod, err = podInformer.Lister().Pods(ns).Get(podName)
	return
}
