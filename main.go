package main

import (
	"flag"
	"fmt"
	"time"

	clientset "github.com/tcontro/pkg/generated/clientset/versioned"
	informers "github.com/tcontro/pkg/generated/informers/externalversions"
	"github.com/tcontro/pkg/signals"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
)

var (
	masterURL  string
	kubeconfig string
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()
	stopCh := signals.SetupSignalHandler()

	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		klog.Fatalf("Error building kubeconfig: %s", err.Error())
	}
	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}
	helloClient, err := clientset.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building example clientset: %s", err.Error())
	}

	kubeInformerFactory := kubeinformers.NewSharedInformerFactory(kubeClient, time.Second*30)
	helloInformerFactory := informers.NewSharedInformerFactory(helloClient, time.Second*30)

	fmt.Println(kubeInformerFactory, helloInformerFactory)

	controller := NewController(kubeClient, helloClient,
		kubeInformerFactory.Apps().V1().Deployments(),
		helloInformerFactory.Bar().V1alpha1().HelloTypes())

	kubeInformerFactory.Start(stopCh)
	helloInformerFactory.Start(stopCh)

	if err = controller.Run(2, stopCh); err != nil {
		klog.Fatalf("Error running controller: %s", err.Error())
	}

}
