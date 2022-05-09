package restic

import (
	"github.com/mperetzred/oadp-client-go/extended/common"
	corev1 "k8s.io/api/core/v1"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
)

type PodsGetter interface {
	Pods(namespace string) ResticPodsInterface
}

type ResticPodsInterface interface {
	common.PodsExpansionInterface
}

type resticpods struct {
	*common.PodsExpansion
}

func newDefaultResticPods(pods corev1client.PodInterface) *resticpods {
	podLogOpts := corev1.PodLogOptions{
		Container: "restic",
	}

	return &resticpods{
		&common.PodsExpansion{
			&podLogOpts,
			"name=restic",
			pods,
		},
	}

}
