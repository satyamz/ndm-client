package main

import (
	"fmt"

	ndmClient "github.com/openebs/node-disk-manager/pkg/client/clientset/versioned"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

func main() {
	fmt.Println("Starting NDM Client")
	config, err := rest.InClusterConfig()
	fmt.Println(config, "  ", err)
	clientset, err := ndmClient.NewForConfig(config)
	if err != nil {
		fmt.Println("Unable to create client")
		return
	}
	diskList, err := clientset.OpenebsV1alpha1().Disks().List(v1.ListOptions{})
	if err != nil {
		fmt.Println("Unable to fetch disk details")
		return
	}
	for _, v := range diskList.Items {
		fmt.Println(v.Name)
	}
	return
}
