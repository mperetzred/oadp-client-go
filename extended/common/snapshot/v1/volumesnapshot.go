package v1

import (
	"context"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1 "github.com/kubernetes-csi/external-snapshotter/client/v4/apis/volumesnapshot/v1"
	snapshotv1 "github.com/kubernetes-csi/external-snapshotter/client/v4/clientset/versioned/typed/volumesnapshot/v1"
)

type VolumeSnapshotExpantionInterface interface {
	IsReadyToUse(voulmeSnapshot *v1.VolumeSnapshot) (bool, error)
}

type VolumeSnapshotExpantion struct {
	snapshotv1.VolumeSnapshotInterface
}

func (vs *VolumeSnapshotExpantion) IsReadyToUse(voulmeSnapshot *v1.VolumeSnapshot) (bool, error) {
	log.Printf("Checking if volumesnapshot is ready to use...")

	volumeSnapshot, err := vs.Get(context.Background(), voulmeSnapshot.Name, metav1.GetOptions{})
	return *volumeSnapshot.Status.ReadyToUse, err

}
