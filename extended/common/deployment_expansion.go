package common

import (
	"context"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	appsv1client "k8s.io/client-go/kubernetes/typed/apps/v1"
)

type DeploymentExpansionInterface interface {
	Get() (*appsv1.Deployment, error)
}

type DeploymentExpansion struct {
	Name          string
	LabelSelector string
	Deployment    appsv1client.DeploymentInterface
}

func (d *DeploymentExpansion) Get() (*appsv1.Deployment, error) {
	return d.Deployment.Get(context.TODO(), d.Name, metav1.GetOptions{})
}

func (d *DeploymentExpansion) List() (*appsv1.DeploymentList, error) {
	return d.Deployment.List(context.TODO(), metav1.ListOptions{
		LabelSelector: d.LabelSelector,
	})
}

func (d *DeploymentExpansion) IsAvailable() (bool, error) {

	deploymentList, err := d.List()
	if err != nil {
		return false, nil
	}
	if len(deploymentList.Items) == 0 {
		return false, nil
	}
	// loop until deployment status is 'Running' or timeout
	for _, deploymentInfo := range deploymentList.Items {
		for _, conditions := range deploymentInfo.Status.Conditions {
			if conditions.Type == appsv1.DeploymentAvailable && conditions.Status != corev1.ConditionTrue {
				return false, fmt.Errorf("%s deployment is not yet available.\nconditions: %v", deploymentInfo.Name, deploymentInfo.Status.Conditions)
			}
		}
	}
	return true, nil

}
