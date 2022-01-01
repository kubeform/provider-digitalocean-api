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

	v1alpha1 "kubeform.dev/provider-digitalocean-api/apis/monitor/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAlerts implements AlertInterface
type FakeAlerts struct {
	Fake *FakeMonitorV1alpha1
	ns   string
}

var alertsResource = schema.GroupVersionResource{Group: "monitor.digitalocean.kubeform.com", Version: "v1alpha1", Resource: "alerts"}

var alertsKind = schema.GroupVersionKind{Group: "monitor.digitalocean.kubeform.com", Version: "v1alpha1", Kind: "Alert"}

// Get takes name of the alert, and returns the corresponding alert object, and an error if there is any.
func (c *FakeAlerts) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Alert, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(alertsResource, c.ns, name), &v1alpha1.Alert{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Alert), err
}

// List takes label and field selectors, and returns the list of Alerts that match those selectors.
func (c *FakeAlerts) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AlertList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(alertsResource, alertsKind, c.ns, opts), &v1alpha1.AlertList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AlertList{ListMeta: obj.(*v1alpha1.AlertList).ListMeta}
	for _, item := range obj.(*v1alpha1.AlertList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested alerts.
func (c *FakeAlerts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(alertsResource, c.ns, opts))

}

// Create takes the representation of a alert and creates it.  Returns the server's representation of the alert, and an error, if there is any.
func (c *FakeAlerts) Create(ctx context.Context, alert *v1alpha1.Alert, opts v1.CreateOptions) (result *v1alpha1.Alert, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(alertsResource, c.ns, alert), &v1alpha1.Alert{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Alert), err
}

// Update takes the representation of a alert and updates it. Returns the server's representation of the alert, and an error, if there is any.
func (c *FakeAlerts) Update(ctx context.Context, alert *v1alpha1.Alert, opts v1.UpdateOptions) (result *v1alpha1.Alert, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(alertsResource, c.ns, alert), &v1alpha1.Alert{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Alert), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAlerts) UpdateStatus(ctx context.Context, alert *v1alpha1.Alert, opts v1.UpdateOptions) (*v1alpha1.Alert, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(alertsResource, "status", c.ns, alert), &v1alpha1.Alert{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Alert), err
}

// Delete takes name of the alert and deletes it. Returns an error if one occurs.
func (c *FakeAlerts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(alertsResource, c.ns, name), &v1alpha1.Alert{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAlerts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(alertsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AlertList{})
	return err
}

// Patch applies the patch and returns the patched alert.
func (c *FakeAlerts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Alert, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(alertsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Alert{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Alert), err
}