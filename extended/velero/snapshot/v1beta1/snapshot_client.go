package v1beta1

import (
	snapshot "github.com/kubernetes-csi/external-snapshotter/client/v4/clientset/versioned/typed/volumesnapshot/v1beta1"
)

type SnapshotV1beta1Interface interface {
	VolumeSnapshotClassGetter
	VolumeSnapshotGetter
}

type SnapshotV1beta1Client struct {
	snapshot.SnapshotV1beta1Interface
}

func (v *SnapshotV1beta1Client) VolumeSnapshotClass() VeleroVolumeSnapshotClassInterface {
	// select Velero pod with this label
	return newDefaultVolumeSnapshotClass(v.VolumeSnapshotClasses())
}

func (v *SnapshotV1beta1Client) VolumeSnapshot(namespace string) VeleroVolumeSnapshotInterface {
	// select Velero pod with this label
	return newDefaultVolumeSnapshot(v.VolumeSnapshots(namespace))
}
