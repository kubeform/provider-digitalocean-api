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

	v1alpha1 "kubeform.dev/provider-digitalocean-api/apis/containerregistry/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDockerCredentialses implements DockerCredentialsInterface
type FakeDockerCredentialses struct {
	Fake *FakeContainerregistryV1alpha1
	ns   string
}

var dockercredentialsesResource = schema.GroupVersionResource{Group: "containerregistry.digitalocean.kubeform.com", Version: "v1alpha1", Resource: "dockercredentialses"}

var dockercredentialsesKind = schema.GroupVersionKind{Group: "containerregistry.digitalocean.kubeform.com", Version: "v1alpha1", Kind: "DockerCredentials"}

// Get takes name of the dockerCredentials, and returns the corresponding dockerCredentials object, and an error if there is any.
func (c *FakeDockerCredentialses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DockerCredentials, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dockercredentialsesResource, c.ns, name), &v1alpha1.DockerCredentials{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DockerCredentials), err
}

// List takes label and field selectors, and returns the list of DockerCredentialses that match those selectors.
func (c *FakeDockerCredentialses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DockerCredentialsList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dockercredentialsesResource, dockercredentialsesKind, c.ns, opts), &v1alpha1.DockerCredentialsList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DockerCredentialsList{ListMeta: obj.(*v1alpha1.DockerCredentialsList).ListMeta}
	for _, item := range obj.(*v1alpha1.DockerCredentialsList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dockerCredentialses.
func (c *FakeDockerCredentialses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dockercredentialsesResource, c.ns, opts))

}

// Create takes the representation of a dockerCredentials and creates it.  Returns the server's representation of the dockerCredentials, and an error, if there is any.
func (c *FakeDockerCredentialses) Create(ctx context.Context, dockerCredentials *v1alpha1.DockerCredentials, opts v1.CreateOptions) (result *v1alpha1.DockerCredentials, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dockercredentialsesResource, c.ns, dockerCredentials), &v1alpha1.DockerCredentials{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DockerCredentials), err
}

// Update takes the representation of a dockerCredentials and updates it. Returns the server's representation of the dockerCredentials, and an error, if there is any.
func (c *FakeDockerCredentialses) Update(ctx context.Context, dockerCredentials *v1alpha1.DockerCredentials, opts v1.UpdateOptions) (result *v1alpha1.DockerCredentials, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dockercredentialsesResource, c.ns, dockerCredentials), &v1alpha1.DockerCredentials{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DockerCredentials), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDockerCredentialses) UpdateStatus(ctx context.Context, dockerCredentials *v1alpha1.DockerCredentials, opts v1.UpdateOptions) (*v1alpha1.DockerCredentials, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dockercredentialsesResource, "status", c.ns, dockerCredentials), &v1alpha1.DockerCredentials{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DockerCredentials), err
}

// Delete takes name of the dockerCredentials and deletes it. Returns an error if one occurs.
func (c *FakeDockerCredentialses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dockercredentialsesResource, c.ns, name), &v1alpha1.DockerCredentials{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDockerCredentialses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dockercredentialsesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DockerCredentialsList{})
	return err
}

// Patch applies the patch and returns the patched dockerCredentials.
func (c *FakeDockerCredentialses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DockerCredentials, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dockercredentialsesResource, c.ns, name, pt, data, subresources...), &v1alpha1.DockerCredentials{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DockerCredentials), err
}
