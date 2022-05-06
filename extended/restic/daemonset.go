package restic

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	appsv1client "k8s.io/client-go/kubernetes/typed/apps/v1"
)

type ResticDaemonSetInterface interface {
	Get() (*appsv1.DaemonSet, error)
}

type resticdaemonset struct {
	Name          string
	LabelSelector string
	DaemonSet     appsv1client.DaemonSetInterface
}

func newDefaultResticDaemonSet(daemonset appsv1client.DaemonSetInterface) *resticdaemonset {
	return &resticdaemonset{
		"restic",
		"metadata.name=restic",
		daemonset,
	}
}

func (d *resticdaemonset) Get() (*appsv1.DaemonSet, error) {
	return d.DaemonSet.Get(context.TODO(), d.Name, metav1.GetOptions{})
}

func (d *resticdaemonset) List() (*appsv1.DaemonSetList, error) {
	return d.DaemonSet.List(context.TODO(), metav1.ListOptions{
		LabelSelector: d.LabelSelector,
	})
}

func (d *resticdaemonset) IsAvailable() (bool, error) {

	resticDaemeonSet, err := d.List()
	if err != nil {
		return false, nil
	}
	var numScheduled int32
	var numDesired int32

	for _, daemonSetInfo := range (*resticDaemeonSet).Items {
		numScheduled = daemonSetInfo.Status.CurrentNumberScheduled
		numDesired = daemonSetInfo.Status.DesiredNumberScheduled
	}
	// check correct num of Restic pods are initialized
	if numScheduled != 0 && numDesired != 0 {
		if numScheduled == numDesired {
			return true, nil
		}
	}
	if numDesired == 0 {
		return true, nil
	}
	return false, err

}
