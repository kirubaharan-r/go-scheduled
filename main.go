package main

import (
	"context"
	"fmt"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	// var kubeconfig *string
	// if home := homedir.HomeDir(); home != "" {
	// 	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	// } else {
	// 	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	// }
	// flag.Parse()

	// config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	config, err := rest.InClusterConfig()

	if err != nil {
		fmt.Println(os.Getenv("GOOS"))
		panic(err.Error())
	}

	ctx := context.Background()
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	poes, err := clientset.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(len(poes.Items))
	fmt.Println("Pod_Name  Status   life_status")

	for _, tes := range poes.Items {
		var a string = string(tes.Status.Phase)
		var b int = int(len(poes.Items))
		c := 10

		if b <= c {
			fmt.Println(tes.Name, "-", a, "-", "schdeuled pod")
		} else {
			fmt.Println(tes.Name, "-", a, "-", "unschdeuled pod")
		}

		//b := "running"

		// if a == b {
		// 	fmt.Print("pod unschedulable")
		// }

	}

}
