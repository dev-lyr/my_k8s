package main

import (
	"flag"
	"fmt"
        "time"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
        watch "k8s.io/apimachinery/pkg/watch"
)

func listPods(podClient *corev1.PodInterface) {
	podList, err := (*podClient).List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, pod := range podList.Items {
		fmt.Println(pod.Name)
	}
}

func createPod(podClient *corev1.PodInterface, podName string) {
	pod := &apiv1.Pod{
		ObjectMeta: metav1.ObjectMeta{Name: podName},
		Spec: apiv1.PodSpec{
			Containers: []apiv1.Container{
				{
					Name:  "web",
					Image: "nginx:1.12",
					Ports: []apiv1.ContainerPort{
						{
							Name:          "http",
							Protocol:      apiv1.ProtocolTCP,
							ContainerPort: 80,
						},
					},
				},
			},
		},
	}
        newPod, err := (*podClient).Create(pod)
        if err != nil{
            panic(err)
        }
        fmt.Printf("create pod:%s success\n", newPod.Name)
}

func deletePod(podClient *corev1.PodInterface, podName string) {
    err := (*podClient).Delete(podName, nil)
    if err != nil{
        panic(err)
    }
    fmt.Printf("delete pod:%s success\n", podName)
}

func watchPod(podClient *corev1.PodInterface){
    podWatch, err := (*podClient).Watch(metav1.ListOptions{})
    if err != nil{
        panic(err)
    }
    defer podWatch.Stop()
    for{
        podEventChan := podWatch.ResultChan()
        var podEvent watch.Event
        select {
            case podEvent = <-podEventChan:
                podObj := podEvent.Object.(*apiv1.Pod)
                fmt.Printf("receive event:%s, obj=%s\n", podEvent.Type, podObj.Name)
                break
        }
    }
}

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "kubeconfig file path(optional)")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	podClient := clientset.CoreV1().Pods(apiv1.NamespaceDefault)

        go watchPod(&podClient)

	//list pod
	//listPods(&podClient)

	//create pod
	createPod(&podClient, "test")
        time.Sleep(time.Duration(20)*time.Second)

	//list pod after create
	//listPods(&podClient)

	//delete pod
	deletePod(&podClient, "test")

	//list pod after delete
	//listPods(&podClient)

        time.Sleep(time.Duration(20)*time.Second)
}
