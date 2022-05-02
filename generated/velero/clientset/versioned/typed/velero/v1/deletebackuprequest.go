/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	scheme "github.com/mperetzred/oadp-client-go/generated/velero/clientset/versioned/scheme"
	v1 "github.com/vmware-tanzu/velero/pkg/apis/velero/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DeleteBackupRequestsGetter has a method to return a DeleteBackupRequestInterface.
// A group's client should implement this interface.
type DeleteBackupRequestsGetter interface {
	DeleteBackupRequests(namespace string) DeleteBackupRequestInterface
}

// DeleteBackupRequestInterface has methods to work with DeleteBackupRequest resources.
type DeleteBackupRequestInterface interface {
	Create(ctx context.Context, deleteBackupRequest *v1.DeleteBackupRequest, opts metav1.CreateOptions) (*v1.DeleteBackupRequest, error)
	Update(ctx context.Context, deleteBackupRequest *v1.DeleteBackupRequest, opts metav1.UpdateOptions) (*v1.DeleteBackupRequest, error)
	UpdateStatus(ctx context.Context, deleteBackupRequest *v1.DeleteBackupRequest, opts metav1.UpdateOptions) (*v1.DeleteBackupRequest, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.DeleteBackupRequest, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.DeleteBackupRequestList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.DeleteBackupRequest, err error)
	DeleteBackupRequestExpansion
}

// deleteBackupRequests implements DeleteBackupRequestInterface
type deleteBackupRequests struct {
	client rest.Interface
	ns     string
}

// newDeleteBackupRequests returns a DeleteBackupRequests
func newDeleteBackupRequests(c *VeleroV1Client, namespace string) *deleteBackupRequests {
	return &deleteBackupRequests{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the deleteBackupRequest, and returns the corresponding deleteBackupRequest object, and an error if there is any.
func (c *deleteBackupRequests) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.DeleteBackupRequest, err error) {
	result = &v1.DeleteBackupRequest{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("deletebackuprequests").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DeleteBackupRequests that match those selectors.
func (c *deleteBackupRequests) List(ctx context.Context, opts metav1.ListOptions) (result *v1.DeleteBackupRequestList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.DeleteBackupRequestList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("deletebackuprequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested deleteBackupRequests.
func (c *deleteBackupRequests) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("deletebackuprequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a deleteBackupRequest and creates it.  Returns the server's representation of the deleteBackupRequest, and an error, if there is any.
func (c *deleteBackupRequests) Create(ctx context.Context, deleteBackupRequest *v1.DeleteBackupRequest, opts metav1.CreateOptions) (result *v1.DeleteBackupRequest, err error) {
	result = &v1.DeleteBackupRequest{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("deletebackuprequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(deleteBackupRequest).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a deleteBackupRequest and updates it. Returns the server's representation of the deleteBackupRequest, and an error, if there is any.
func (c *deleteBackupRequests) Update(ctx context.Context, deleteBackupRequest *v1.DeleteBackupRequest, opts metav1.UpdateOptions) (result *v1.DeleteBackupRequest, err error) {
	result = &v1.DeleteBackupRequest{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("deletebackuprequests").
		Name(deleteBackupRequest.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(deleteBackupRequest).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *deleteBackupRequests) UpdateStatus(ctx context.Context, deleteBackupRequest *v1.DeleteBackupRequest, opts metav1.UpdateOptions) (result *v1.DeleteBackupRequest, err error) {
	result = &v1.DeleteBackupRequest{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("deletebackuprequests").
		Name(deleteBackupRequest.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(deleteBackupRequest).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the deleteBackupRequest and deletes it. Returns an error if one occurs.
func (c *deleteBackupRequests) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("deletebackuprequests").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *deleteBackupRequests) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("deletebackuprequests").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched deleteBackupRequest.
func (c *deleteBackupRequests) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.DeleteBackupRequest, err error) {
	result = &v1.DeleteBackupRequest{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("deletebackuprequests").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}