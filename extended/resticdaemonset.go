package extended

import (
	"context"
	"time"

	"github.com/mperetzred/oadp-client-go/generated/oadp/clientset/versioned/scheme"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	appsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	"k8s.io/client-go/rest"
)

// DaemonSetsGetter has a method to return a DaemonSetInterface.
// A group's client should implement this interface.
type ResticDaemonSetsGetter interface {
	DaemonSets(namespace string) ResticDaemonSetInterface
}

// DaemonSetInterface has methods to work with DaemonSet resources.
type ResticDaemonSetInterface interface {
	Get() (*v1.DaemonSet, error)
	List(opts metav1.ListOptions) (*v1.DaemonSetList, error)
}

// daemonSets implements DaemonSetInterface
type daemonSets struct {
	client rest.Interface
	ns     string
}

// newDaemonSets returns a DaemonSets
func newDaemonSets(c *appsv1.AppsV1Client, namespace string) *daemonSets {
	return &daemonSets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the daemonSet, and returns the corresponding daemonSet object, and an error if there is any.
func (c *daemonSets) Get() (result *v1.DaemonSet, err error) {
	result = &v1.DaemonSet{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("daemonsets").
		Name("restic").
		VersionedParams(&metav1.GetOptions{}, scheme.ParameterCodec).
		Do(context.Background()).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DaemonSets that match those selectors.
func (c *daemonSets) List(opts metav1.ListOptions) (result *v1.DaemonSetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.DaemonSetList{}
	opts.LabelSelector = "metadata.name=restic"
	err = c.client.Get().
		Namespace(c.ns).
		Resource("daemonsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(context.Background()).
		Into(result)
	return
}

func (c *daemonSets) IsAvailable() (bool, error) {
	ds, err := c.Get()
	if err != nil {
		return false, err
	}
	numScheduled := ds.Status.CurrentNumberScheduled
	numDesired := ds.Status.DesiredNumberScheduled

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
