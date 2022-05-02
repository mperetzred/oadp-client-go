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

package fake

import (
	"context"

	velerov1 "github.com/vmware-tanzu/velero/pkg/apis/velero/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePodVolumeBackups implements PodVolumeBackupInterface
type FakePodVolumeBackups struct {
	Fake *FakeVeleroV1
	ns   string
}

var podvolumebackupsResource = schema.GroupVersionResource{Group: "velero.io", Version: "v1", Resource: "podvolumebackups"}

var podvolumebackupsKind = schema.GroupVersionKind{Group: "velero.io", Version: "v1", Kind: "PodVolumeBackup"}

// Get takes name of the podVolumeBackup, and returns the corresponding podVolumeBackup object, and an error if there is any.
func (c *FakePodVolumeBackups) Get(ctx context.Context, name string, options v1.GetOptions) (result *velerov1.PodVolumeBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(podvolumebackupsResource, c.ns, name), &velerov1.PodVolumeBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*velerov1.PodVolumeBackup), err
}

// List takes label and field selectors, and returns the list of PodVolumeBackups that match those selectors.
func (c *FakePodVolumeBackups) List(ctx context.Context, opts v1.ListOptions) (result *velerov1.PodVolumeBackupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(podvolumebackupsResource, podvolumebackupsKind, c.ns, opts), &velerov1.PodVolumeBackupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &velerov1.PodVolumeBackupList{ListMeta: obj.(*velerov1.PodVolumeBackupList).ListMeta}
	for _, item := range obj.(*velerov1.PodVolumeBackupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested podVolumeBackups.
func (c *FakePodVolumeBackups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(podvolumebackupsResource, c.ns, opts))

}

// Create takes the representation of a podVolumeBackup and creates it.  Returns the server's representation of the podVolumeBackup, and an error, if there is any.
func (c *FakePodVolumeBackups) Create(ctx context.Context, podVolumeBackup *velerov1.PodVolumeBackup, opts v1.CreateOptions) (result *velerov1.PodVolumeBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(podvolumebackupsResource, c.ns, podVolumeBackup), &velerov1.PodVolumeBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*velerov1.PodVolumeBackup), err
}

// Update takes the representation of a podVolumeBackup and updates it. Returns the server's representation of the podVolumeBackup, and an error, if there is any.
func (c *FakePodVolumeBackups) Update(ctx context.Context, podVolumeBackup *velerov1.PodVolumeBackup, opts v1.UpdateOptions) (result *velerov1.PodVolumeBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(podvolumebackupsResource, c.ns, podVolumeBackup), &velerov1.PodVolumeBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*velerov1.PodVolumeBackup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePodVolumeBackups) UpdateStatus(ctx context.Context, podVolumeBackup *velerov1.PodVolumeBackup, opts v1.UpdateOptions) (*velerov1.PodVolumeBackup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(podvolumebackupsResource, "status", c.ns, podVolumeBackup), &velerov1.PodVolumeBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*velerov1.PodVolumeBackup), err
}

// Delete takes name of the podVolumeBackup and deletes it. Returns an error if one occurs.
func (c *FakePodVolumeBackups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(podvolumebackupsResource, c.ns, name, opts), &velerov1.PodVolumeBackup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePodVolumeBackups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(podvolumebackupsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &velerov1.PodVolumeBackupList{})
	return err
}

// Patch applies the patch and returns the patched podVolumeBackup.
func (c *FakePodVolumeBackups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *velerov1.PodVolumeBackup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(podvolumebackupsResource, c.ns, name, pt, data, subresources...), &velerov1.PodVolumeBackup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*velerov1.PodVolumeBackup), err
}
