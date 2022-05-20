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
	RestoresGetter
	velerov1.BackupStorageLocationsGetter
	velerov1.DeleteBackupRequestsGetter
	velerov1.DownloadRequestsGetter
	velerov1.PodVolumeBackupsGetter
	velerov1.PodVolumeRestoresGetter
	velerov1.ResticRepositoriesGetter
	velerov1.SchedulesGetter
	velerov1.ServerStatusRequestsGetter
	velerov1.VolumeSnapshotLocationsGetter
	Snapshot() snapshot.Interface
}

type VeleroV1Client struct {
	corev1Client   *corev1client.CoreV1Client
	appsv1Client   *appsv1client.AppsV1Client
	snapshotClient *snapshot.Clientset
	velerov1Client *velerov1.VeleroV1Client
}

func (v *VeleroV1Client) Pods(namespace string) PodInterface {
	// select Velero pod with this label
	return newPods(v.corev1Client, namespace)
}

func (v *VeleroV1Client) Deployments(namespace string) DeploymentInterface {
	// select Velero pod with this label
	return newDeployments(v.appsv1Client, namespace)
}

func (v *VeleroV1Client) Snapshot() snapshot.Interface {
	// select Velero pod with this label
	return v.snapshotClient
}

func (v *VeleroV1Client) Backups(namespace string) BackupExpansionInterface {
	// select Velero pod with this label
	return newBackupExpansion(v.velerov1Client.Backups(namespace))
}

func (v *VeleroV1Client) Restores(namespace string) RestoreExpansionInterface {
	// select Velero pod with this label
	return newRestoreExpansion(v.velerov1Client.Restores(namespace))
}

func (c *VeleroV1Client) BackupStorageLocations(namespace string) velerov1.BackupStorageLocationInterface {
	return c.velerov1Client.BackupStorageLocations(namespace)
}

func (c *VeleroV1Client) DeleteBackupRequests(namespace string) velerov1.DeleteBackupRequestInterface {
	return c.velerov1Client.DeleteBackupRequests(namespace)
}

func (c *VeleroV1Client) DownloadRequests(namespace string) velerov1.DownloadRequestInterface {
	return c.velerov1Client.DownloadRequests(namespace)
}

func (c *VeleroV1Client) PodVolumeBackups(namespace string) velerov1.PodVolumeBackupInterface {
	return c.velerov1Client.PodVolumeBackups(namespace)
}

func (c *VeleroV1Client) PodVolumeRestores(namespace string) velerov1.PodVolumeRestoreInterface {
	return c.velerov1Client.PodVolumeRestores(namespace)
}

func (c *VeleroV1Client) ResticRepositories(namespace string) velerov1.ResticRepositoryInterface {
	return c.velerov1Client.ResticRepositories(namespace)
}

func (c *VeleroV1Client) Schedules(namespace string) velerov1.ScheduleInterface {
	return c.velerov1Client.Schedules(namespace)
}

func (c *VeleroV1Client) ServerStatusRequests(namespace string) velerov1.ServerStatusRequestInterface {
	return c.velerov1Client.ServerStatusRequests(namespace)
}

func (c *VeleroV1Client) VolumeSnapshotLocations(namespace string) velerov1.VolumeSnapshotLocationInterface {
	return c.velerov1Client.VolumeSnapshotLocations(namespace)
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

	cs.velerov1Client, err = velerov1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	return &cs, nil
}
