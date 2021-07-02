/*
Copyright AppsCode Inc. and Contributors

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

	v1alpha1 "kubeform.dev/provider-digitalocean-api/apis/droplet/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSnapshots implements SnapshotInterface
type FakeSnapshots struct {
	Fake *FakeDropletV1alpha1
	ns   string
}

var snapshotsResource = schema.GroupVersionResource{Group: "droplet.digitalocean.kubeform.com", Version: "v1alpha1", Resource: "snapshots"}

var snapshotsKind = schema.GroupVersionKind{Group: "droplet.digitalocean.kubeform.com", Version: "v1alpha1", Kind: "Snapshot"}

// Get takes name of the snapshot, and returns the corresponding snapshot object, and an error if there is any.
func (c *FakeSnapshots) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Snapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(snapshotsResource, c.ns, name), &v1alpha1.Snapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Snapshot), err
}

// List takes label and field selectors, and returns the list of Snapshots that match those selectors.
func (c *FakeSnapshots) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SnapshotList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(snapshotsResource, snapshotsKind, c.ns, opts), &v1alpha1.SnapshotList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SnapshotList{ListMeta: obj.(*v1alpha1.SnapshotList).ListMeta}
	for _, item := range obj.(*v1alpha1.SnapshotList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested snapshots.
func (c *FakeSnapshots) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(snapshotsResource, c.ns, opts))

}

// Create takes the representation of a snapshot and creates it.  Returns the server's representation of the snapshot, and an error, if there is any.
func (c *FakeSnapshots) Create(ctx context.Context, snapshot *v1alpha1.Snapshot, opts v1.CreateOptions) (result *v1alpha1.Snapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(snapshotsResource, c.ns, snapshot), &v1alpha1.Snapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Snapshot), err
}

// Update takes the representation of a snapshot and updates it. Returns the server's representation of the snapshot, and an error, if there is any.
func (c *FakeSnapshots) Update(ctx context.Context, snapshot *v1alpha1.Snapshot, opts v1.UpdateOptions) (result *v1alpha1.Snapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(snapshotsResource, c.ns, snapshot), &v1alpha1.Snapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Snapshot), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSnapshots) UpdateStatus(ctx context.Context, snapshot *v1alpha1.Snapshot, opts v1.UpdateOptions) (*v1alpha1.Snapshot, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(snapshotsResource, "status", c.ns, snapshot), &v1alpha1.Snapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Snapshot), err
}

// Delete takes name of the snapshot and deletes it. Returns an error if one occurs.
func (c *FakeSnapshots) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(snapshotsResource, c.ns, name), &v1alpha1.Snapshot{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSnapshots) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(snapshotsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SnapshotList{})
	return err
}

// Patch applies the patch and returns the patched snapshot.
func (c *FakeSnapshots) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Snapshot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(snapshotsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Snapshot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Snapshot), err
}
