package velero

import (
	"context"
	"time"
	commonoadp "github.com/openshift/oadp-operator/pkg/common"
	v1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	scheme "k8s.io/client-go/kubernetes/scheme"
	appsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	rest "k8s.io/client-go/rest"
)

// DeploymentsGetter has a method to return a DeploymentInterface.
// A group's client should implement this interface.
type DeploymentsGetter interface {
	Deployments(namespace string) DeploymentInterface
}

// DeploymentInterface has methods to work with Deployment resources.
type DeploymentInterface interface {
	Get() (*v1.Deployment, error)
	List(opts metav1.ListOptions) (*v1.DeploymentList, error)
}

// deployments implements DeploymentInterface
type deployments struct {
	client rest.Interface
	ns     string
}

// newDeployments returns a Deployments
func newDeployments(c *appsv1.AppsV1Client, namespace string) *deployments {
	return &deployments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the deployment, and returns the corresponding deployment object, and an error if there is any.
func (c *deployments) Get() (result *v1.Deployment, err error) {
	result = &v1.Deployment{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("deployments").
		Name("velero").
		VersionedParams(&metav1.GetOptions{}, scheme.ParameterCodec).
		Do(context.Background()).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Deployments that match those selectors.
func (c *deployments) List(opts metav1.ListOptions) (result *v1.DeploymentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.DeploymentList{}
	opts.LabelSelector = "app.kubernetes.io/name=velero"
	err = c.client.Get().
		Namespace(c.ns).
		Resource("deployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(context.Background()).
		Into(result)
	return
}

func (vd *deployments) Tolerations() (t []corev1.Toleration) {
	deploy, _ := vd.Get()
	return deploy.Spec.Template.Spec.Tolerations
}

func (vd *deployments) Requests() (t corev1.ResourceList) {
	deploy, _ := vd.Get()
	for _, c := range deploy.Spec.Template.Spec.Containers {
		if c.Name == commonoadp.Velero {
			return c.Resources.Requests
		}
	}
	return nil
}

func (vd *deployments) ResourcesLimits() (t corev1.ResourceList) {
	deploy, _ := vd.Get()
	for _, c := range deploy.Spec.Template.Spec.Containers {
		if c.Name == commonoadp.Velero {
			return c.Resources.Limits
		}
	}
	return nil
}
