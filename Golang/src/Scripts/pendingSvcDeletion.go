package main

import (
	"context"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	discoveryv1beta1 "k8s.io/api/discovery/v1beta1"
	k8sError "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"log"
	"sync"
	"time"
)

func createPendingDeletionSvc(ns string, clientset *kubernetes.Clientset) {
	ctx := context.Background()
	var svcLst []string
	wg := &sync.WaitGroup{}
	lock := &sync.Mutex{}
	for i := 0; i < 600; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			name := fmt.Sprintf("charles-test-%d", i)
			// Create svc
			svc := &corev1.Service{
				ObjectMeta: metav1.ObjectMeta{Name: name, Labels: map[string]string{"charles": "test"}},
				Spec:       corev1.ServiceSpec{Ports: []corev1.ServicePort{{Port: int32(20000 + i)}}},
			}
			svc, err := clientset.CoreV1().Services(ns).Create(ctx, svc, metav1.CreateOptions{})
			if err != nil && !k8sError.IsAlreadyExists(err) {
				log.Println("create service fail", i, err)
			} else {
				lock.Lock()
				svcLst = append(svcLst, name)
				lock.Unlock()
			}
			// Create endpointslice
			trueB := true
			es := &discoveryv1beta1.EndpointSlice{
				AddressType: discoveryv1beta1.AddressTypeIPv4,
				ObjectMeta: metav1.ObjectMeta{
					GenerateName: name,
					OwnerReferences: []metav1.OwnerReference{
						{
							APIVersion:         "v1",
							Kind:               "Service",
							Name:               name,
							UID:                svc.UID,
							Controller:         &trueB,
							BlockOwnerDeletion: &trueB,
						},
					}},
			}
			es, err = clientset.DiscoveryV1beta1().EndpointSlices(ns).Create(ctx, es, metav1.CreateOptions{})
			if err != nil && !k8sError.IsAlreadyExists(err) {
				log.Println("create endpointslice fail", i, err)
			}

		}(i)
	}
	wg.Wait()
	// Delete svc
	time.Sleep(10)

	for _, svcName := range svcLst {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			log.Println("Start delete", name)
			propagation := metav1.DeletePropagationForeground
			err := clientset.CoreV1().Services(ns).Delete(ctx, name, metav1.DeleteOptions{PropagationPolicy: &propagation})
			if err != nil {
				log.Println("delete service fail", name, err)
			}
		}(svcName)
	}
	wg.Wait()
	log.Println("Finish")
}
