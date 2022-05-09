package velero

import (
	snapshot "github.com/mperetzred/oadp-client-go/extended/velero/snapshot"
	velerov1 "github.com/vmware-tanzu/velero/pkg/generated/clientset/versioned/typed/velero/v1"
	"k8s.io/client-go/kubernetes"
	appsv1client "k8s.io/client-go/kubernetes/typed/apps/v1"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
)

type VeleroV1Interface interface {
	PodsGetter
	DeploymentsGetter
	BackupsGetter
	Snapshot() snapshot.Interface
	velerov1.VeleroV1Interface // extends velerov1.VeleroV1Interface
}

type VeleroV1Client struct {
	corev1Client   *corev1client.CoreV1Client
	appsv1Client   *appsv1client.AppsV1Client
	snapshotClient *snapshot.Clientset
	*velerov1.VeleroV1Client
}

func (v *VeleroV1Client) Pods(namespace string) VeleroPodsInterface {
	// select Velero pod with this label
	return newDefaultVeleroPods(v.corev1Client.Pods(namespace))
}

func (v *VeleroV1Client) Deployments(namespace string) VeleroDeploymentInterface {
	// select Velero pod with this label
	return newDefaultVeleroDeployment(v.appsv1Client.Deployments(namespace))
}

func (v *VeleroV1Client) Snapshot() snapshot.Interface {
	// select Velero pod with this label
	return v.snapshotClient
}

func (v *VeleroV1Client) BackupExpansion(namespace string) BackupExpansionInterface {
	// select Velero pod with this label
	return newBackupExpansion(v.Backups(namespace))
}

func (v *VeleroV1Client) RestoreExpansion(namespace string) RestoreExpansionInterface {
	// select Velero pod with this label
	return newRestoreExpansion(v.Restores(namespace))
}

func NewForConfig(c *rest.Config) (*VeleroV1Client, error) {
	configShallowCopy := *c

	if configShallowCopy.UserAgent == "" {
		configShallowCopy.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	// share the transport between all clients
	var cs VeleroV1Client
	kubernetesClient, err := kubernetes.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.appsv1Client = kubernetesClient.AppsV1().(*appsv1client.AppsV1Client)

	cs.corev1Client = kubernetesClient.CoreV1().(*corev1client.CoreV1Client)

	cs.snapshotClient, err = snapshot.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.VeleroV1Client, err = velerov1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	return &cs, nil
}
