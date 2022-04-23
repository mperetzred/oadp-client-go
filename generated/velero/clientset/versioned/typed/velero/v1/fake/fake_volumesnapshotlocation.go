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

// FakeVolumeSnapshotLocations implements VolumeSnapshotLocationInterface
type FakeVolumeSnapshotLocations struct {
	Fake *FakeVeleroV1
	ns   string
}

var volumesnapshotlocationsResource = schema.GroupVersionResource{Group: "velero.io", Version: "v1", Resource: "volumesnapshotlocations"}

var volumesnapshotlocationsKind = schema.GroupVersionKind{Group: "velero.io", Version: "v1", Kind: "VolumeSnapshotLocation"}

// Get takes name of the volumeSnapshotLocation, and returns the corresponding volumeSnapshotLocation object, and an error if there is any.
func (c *FakeVolumeSnapshotLocations) Get(ctx context.Context, name string, options v1.GetOptions) (result *velerov1.VolumeSnapshotLocation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(volumesnapshotlocationsResource, c.ns, name), &velerov1.VolumeSnapshotLocation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*velerov1.VolumeSnapshotLocation), err
}

// List takes label and field selectors, and returns the list of VolumeSnapshotLocations that match those selectors.
func (c *FakeVolumeSnapshotLocations) List(ctx context.Context, opts v1.ListOptions) (result *velerov1.VolumeSnapshotLocationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(volumesnapshotlocationsResource, volumesnapshotlocationsKind, c.ns, opts), &velerov1.VolumeSnapshotLocationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &velerov1.VolumeSnapshotLocationList{ListMeta: obj.(*velerov1.VolumeSnapshotLocationList).ListMeta}
	for _, item := range obj.(*velerov1.VolumeSnapshotLocationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested volumeSnapshotLocations.
func (c *FakeVolumeSnapshotLocations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(volumesnapshotlocationsResource, c.ns, opts))

}

// Create takes the representation of a volumeSnapshotLocation and creates it.  Returns the server's representation of the volumeSnapshotLocation, and an error, if there is any.
func (c *FakeVolumeSnapshotLocations) Create(ctx context.Context, volumeSnapshotLocation *velerov1.VolumeSnapshotLocation, opts v1.CreateOptions) (result *velerov1.VolumeSnapshotLocation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(volumesnapshotlocationsResource, c.ns, volumeSnapshotLocation), &velerov1.VolumeSnapshotLocation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*velerov1.VolumeSnapshotLocation), err
}

// Update takes the representation of a volumeSnapshotLocation and updates it. Returns the server's representation of the volumeSnapshotLocation, and an error, if there is any.
func (c *FakeVolumeSnapshotLocations) Update(ctx context.Context, volumeSnapshotLocation *velerov1.VolumeSnapshotLocation, opts v1.UpdateOptions) (result *velerov1.VolumeSnapshotLocation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(volumesnapshotlocationsResource, c.ns, volumeSnapshotLocation), &velerov1.VolumeSnapshotLocation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*velerov1.VolumeSnapshotLocation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVolumeSnapshotLocations) UpdateStatus(ctx context.Context, volumeSnapshotLocation *velerov1.VolumeSnapshotLocation, opts v1.UpdateOptions) (*velerov1.VolumeSnapshotLocation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(volumesnapshotlocationsResource, "status", c.ns, volumeSnapshotLocation), &velerov1.VolumeSnapshotLocation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*velerov1.VolumeSnapshotLocation), err
}

// Delete takes name of the volumeSnapshotLocation and deletes it. Returns an error if one occurs.
func (c *FakeVolumeSnapshotLocations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(volumesnapshotlocationsResource, c.ns, name, opts), &velerov1.VolumeSnapshotLocation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVolumeSnapshotLocations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(volumesnapshotlocationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &velerov1.VolumeSnapshotLocationList{})
	return err
}

// Patch applies the patch and returns the patched volumeSnapshotLocation.
func (c *FakeVolumeSnapshotLocations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *velerov1.VolumeSnapshotLocation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(volumesnapshotlocationsResource, c.ns, name, pt, data, subresources...), &velerov1.VolumeSnapshotLocation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*velerov1.VolumeSnapshotLocation), err
}
