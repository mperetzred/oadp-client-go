package v1

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1 "github.com/kubernetes-csi/external-snapshotter/client/v6/apis/volumesnapshot/v1"
	snapshotv1 "github.com/kubernetes-csi/external-snapshotter/client/v6/clientset/versioned/typed/volumesnapshot/v1"
	snapshotv1common "github.com/mperetzred/oadp-client-go/extended/common/snapshot/v1"
)

type VolumeSnapshotGetter interface {
	VolumeSnapshot(namespace string) VeleroVolumeSnapshotInterface
}

type VeleroVolumeSnapshotInterface interface {
	ListByLabelSelector(backupName string) (*v1.VolumeSnapshotList, error)
}

type velerovolumesnapshot struct {
	name           string
	volumesnapshot snapshotv1.VolumeSnapshotInterface
	labelkey       string
}

func newDefaultVolumeSnapshot(volumesnapshot snapshotv1.VolumeSnapshotInterface) *velerovolumesnapshot {
	return &velerovolumesnapshot{
		"",
		&snapshotv1common.VolumeSnapshotExpantion{volumesnapshot},
		"velero.io/backup-name",
	}
}

func (vs *velerovolumesnapshot) ListByLabelSelector(backupName string) (*v1.VolumeSnapshotList, error) {
	return vs.volumesnapshot.List(context.TODO(), metav1.ListOptions{LabelSelector: vs.labelkey + "=" + backupName})
}
