package v1

import (
	snapshot "github.com/kubernetes-csi/external-snapshotter/client/v6/clientset/versioned/typed/volumesnapshot/v1"
)

type SnapshotV1Interface interface {
	VolumeSnapshotClass() VeleroVolumeSnapshotClassInterface
}

type SnapshotV1Client struct {
	snapshot.SnapshotV1Interface
}

func (v *SnapshotV1Client) VolumeSnapshotClass() VeleroVolumeSnapshotClassInterface {
	// select Velero pod with this label
	return newDefaultVolumeSnapshotClass(v.VolumeSnapshotClasses())
}
