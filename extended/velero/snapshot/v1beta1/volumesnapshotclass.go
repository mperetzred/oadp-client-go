package v1beta1

import (
	"context"

	v1beta1 "github.com/kubernetes-csi/external-snapshotter/client/v4/apis/volumesnapshot/v1beta1"
	snapshotv1beta1 "github.com/kubernetes-csi/external-snapshotter/client/v4/clientset/versioned/typed/volumesnapshot/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

type VolumeSnapshotClassGetter interface {
	VolumeSnapshotClass() VeleroVolumeSnapshotClassInterface
}

type VeleroVolumeSnapshotClassInterface interface {
	Create(name string, driver string, parameters map[string]string) (*v1beta1.VolumeSnapshotClass, error)
}

type velerovolumesnapshotclass struct {
	name                string
	volumesnapshotclass snapshotv1beta1.VolumeSnapshotClassInterface
	annotations         map[string]string
	labels              map[string]string
}

func newDefaultVolumeSnapshotClass(volumesnapshotclass snapshotv1beta1.VolumeSnapshotClassInterface) *velerovolumesnapshotclass {
	return &velerovolumesnapshotclass{
		"",
		volumesnapshotclass,
		map[string]string{
			"snapshot.storage.kubernetes.io/is-default-class": "true",
		},
		map[string]string{
			"velero.io/csi-volumesnapshot-class": "true",
		},
	}
}

func (vsc *velerovolumesnapshotclass) Create(name string, driver string, parameters map[string]string) (*v1beta1.VolumeSnapshotClass, error) {
	vs := v1beta1.VolumeSnapshotClass{
		ObjectMeta: metav1.ObjectMeta{
			Name:        name,
			Annotations: vsc.annotations,
			Labels:      vsc.labels,
		},
		Driver:         driver,
		DeletionPolicy: v1beta1.VolumeSnapshotContentRetain,
		Parameters:     parameters,
	}

	if name != "" {
		vs.ObjectMeta.Name = name
	} else {
		name = "velero-snapclass-"
		vs.ObjectMeta.GenerateName = name
	}
	vsc.name = name
	return vsc.volumesnapshotclass.Create(context.TODO(), &vs, metav1.CreateOptions{})
}

func (vsc *velerovolumesnapshotclass) CreateWithDefaultName(driver string, parameters map[string]string) (*v1beta1.VolumeSnapshotClass, error) {
	return vsc.Create(vsc.name, driver, parameters)
}

func (vsc *velerovolumesnapshotclass) Delete() error {
	return vsc.volumesnapshotclass.Delete(context.TODO(), vsc.name, metav1.DeleteOptions{})
}

func (vsc *velerovolumesnapshotclass) DeleteCollection() error {
	return vsc.volumesnapshotclass.DeleteCollection(context.TODO(), metav1.DeleteOptions{}, metav1.ListOptions{
		LabelSelector: labels.SelectorFromSet(vsc.labels).String(),
	})
}

func (vsc *velerovolumesnapshotclass) List() (*v1beta1.VolumeSnapshotClassList, error) {
	return vsc.volumesnapshotclass.List(context.TODO(), metav1.ListOptions{
		LabelSelector: labels.SelectorFromSet(vsc.labels).String(),
	})
}
