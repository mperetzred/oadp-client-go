package velero

import (
	"github.com/mperetzred/oadp-client-go/extended/common"
	commonoadp "github.com/openshift/oadp-operator/pkg/common"
	corev1 "k8s.io/api/core/v1"
	appsv1client "k8s.io/client-go/kubernetes/typed/apps/v1"
)

type DeploymentsGetter interface {
	Deployments(namespace string) VeleroDeploymentInterface
}

type VeleroDeploymentInterface interface {
	common.DeploymentExpansionInterface
}

type velerodeployment struct {
	*common.DeploymentExpansion
}

func newDefaultVeleroDeployment(deployment appsv1client.DeploymentInterface) *velerodeployment {
	return &velerodeployment{
		&common.DeploymentExpansion{
			"velero",
			"app.kubernetes.io/name=velero",
			deployment,
		},
	}
}

func (vd *velerodeployment) Tolerations() (t []corev1.Toleration) {
	deploy, _ := vd.Get()
	return deploy.Spec.Template.Spec.Tolerations
}

func (vd *velerodeployment) Requests() (t corev1.ResourceList) {
	deploy, _ := vd.Get()
	for _, c := range deploy.Spec.Template.Spec.Containers {
		if c.Name == commonoadp.Velero {
			return c.Resources.Requests
		}
	}
	return nil
}

func (vd *velerodeployment) ResourcesLimits() (t corev1.ResourceList) {
	deploy, _ := vd.Get()
	for _, c := range deploy.Spec.Template.Spec.Containers {
		if c.Name == commonoadp.Velero {
			return c.Resources.Limits
		}
	}
	return nil
}
