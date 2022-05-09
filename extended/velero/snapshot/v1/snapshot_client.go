package v1

import (
	snapshot "github.com/kubernetes-csi/external-snapshotter/client/v6/clientset/versioned/typed/volumesnapshot/v1"
)

type SnapshotV1Interface interface {
	VolumeSnapshotClassGetter
	VolumeSnapshotGetter
}

type SnapshotV1Client struct {
	snapshot.SnapshotV1Interface
}

func (v *SnapshotV1Client) VolumeSnapshotClass() VeleroVolumeSnapshotClassInterface {
	// select Velero pod with this label
	return newDefaultVolumeSnapshotClass(v.VolumeSnapshotClasses())
}

func (v *SnapshotV1Client) VolumeSnapshot(namespace string) VeleroVolumeSnapshotInterface {
	// select Velero pod with this label
	return newDefaultVolumeSnapshot(v.VolumeSnapshots(namespace))
}
