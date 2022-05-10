package v1beta1

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1beta1 "github.com/kubernetes-csi/external-snapshotter/client/v4/apis/volumesnapshot/v1beta1"
	snapshotv1beta1 "github.com/kubernetes-csi/external-snapshotter/client/v4/clientset/versioned/typed/volumesnapshot/v1beta1"
	snapshotv1beta1common "github.com/mperetzred/oadp-client-go/extended/common/snapshot/v1beta1"
)

type VolumeSnapshotGetter interface {
	VolumeSnapshot(namespace string) VeleroVolumeSnapshotInterface
}

type VeleroVolumeSnapshotInterface interface {
	ListByLabelSelector(backupName string) (*v1beta1.VolumeSnapshotList, error)
}

type velerovolumesnapshot struct {
	name           string
	volumesnapshot *snapshotv1beta1common.VolumeSnapshotExpantion
	labelkey       string
}

func newDefaultVolumeSnapshot(volumesnapshot snapshotv1beta1.VolumeSnapshotInterface) *velerovolumesnapshot {
	return &velerovolumesnapshot{
		"",
		&snapshotv1beta1common.VolumeSnapshotExpantion{volumesnapshot},
		"velero.io/backup-name",
	}
}

func (vs *velerovolumesnapshot) ListByLabelSelector(backupName string) (*v1beta1.VolumeSnapshotList, error) {
	return vs.volumesnapshot.List(context.TODO(), metav1.ListOptions{LabelSelector: vs.labelkey + "=" + backupName})
}
