package kubecli

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type oadp struct {
	kubeconfig string
	restClient *rest.RESTClient
	config     *rest.Config
	*kubernetes.Clientset
}
