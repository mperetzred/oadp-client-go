package restic

import (
	"k8s.io/client-go/kubernetes"
	appsv1client "k8s.io/client-go/kubernetes/typed/apps/v1"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
)

type ResticInterface interface {
	PodsGetter
	DaemonSetsGetter
}

type ResticClient struct {
	corev1Client *corev1client.CoreV1Client
	appsv1Client *appsv1client.AppsV1Client
}

func (v *ResticClient) Pods(namespace string) ResticPodsInterface {
	// select Velero pod with this label
	return newDefaultResticPods(v.corev1Client.Pods(namespace))
}

func (v *ResticClient) DeamonSet(namespace string) ResticDaemonSetInterface {
	// select Velero pod with this label
	return newDefaultResticDaemonSet(v.appsv1Client.DaemonSets(namespace))
}

func NewForConfig(c *rest.Config) (*ResticClient, error) {
	configShallowCopy := *c

	if configShallowCopy.UserAgent == "" {
		configShallowCopy.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	// share the transport between all clients
	var cs ResticClient
	kubernetesClient, err := kubernetes.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.appsv1Client = kubernetesClient.AppsV1().(*appsv1client.AppsV1Client)

	cs.corev1Client = kubernetesClient.CoreV1().(*corev1client.CoreV1Client)

	return &cs, nil
}
