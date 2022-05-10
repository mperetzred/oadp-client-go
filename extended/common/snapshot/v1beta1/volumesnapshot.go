package v1

import (
	"context"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1beta1 "github.com/kubernetes-csi/external-snapshotter/client/v4/apis/volumesnapshot/v1beta1"
	snapshotv1beta1 "github.com/kubernetes-csi/external-snapshotter/client/v4/clientset/versioned/typed/volumesnapshot/v1beta1"
)

type VolumeSnapshotExpantionInterface interface {
	IsVolumeSnapshotReadyToUse(voulmeSnapshot *v1beta1.VolumeSnapshot) (bool, error)
}

type VolumeSnapshotExpantion struct {
	snapshotv1beta1.VolumeSnapshotInterface
}

func (vs *VolumeSnapshotExpantion) IsVolumeSnapshotReadyToUse(voulmeSnapshot *v1beta1.VolumeSnapshot) (bool, error) {
	log.Printf("Checking if volumesnapshot is ready to use...")

	volumeSnapshot, err := vs.Get(context.Background(), voulmeSnapshot.Name, metav1.GetOptions{})
	return *volumeSnapshot.Status.ReadyToUse, err

}
