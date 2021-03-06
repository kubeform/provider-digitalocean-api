//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerRegistry) DeepCopyInto(out *ContainerRegistry) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerRegistry.
func (in *ContainerRegistry) DeepCopy() *ContainerRegistry {
	if in == nil {
		return nil
	}
	out := new(ContainerRegistry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ContainerRegistry) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerRegistryList) DeepCopyInto(out *ContainerRegistryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ContainerRegistry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerRegistryList.
func (in *ContainerRegistryList) DeepCopy() *ContainerRegistryList {
	if in == nil {
		return nil
	}
	out := new(ContainerRegistryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ContainerRegistryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerRegistrySpec) DeepCopyInto(out *ContainerRegistrySpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ContainerRegistrySpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerRegistrySpec.
func (in *ContainerRegistrySpec) DeepCopy() *ContainerRegistrySpec {
	if in == nil {
		return nil
	}
	out := new(ContainerRegistrySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerRegistrySpecResource) DeepCopyInto(out *ContainerRegistrySpecResource) {
	*out = *in
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ServerURL != nil {
		in, out := &in.ServerURL, &out.ServerURL
		*out = new(string)
		**out = **in
	}
	if in.SubscriptionTierSlug != nil {
		in, out := &in.SubscriptionTierSlug, &out.SubscriptionTierSlug
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerRegistrySpecResource.
func (in *ContainerRegistrySpecResource) DeepCopy() *ContainerRegistrySpecResource {
	if in == nil {
		return nil
	}
	out := new(ContainerRegistrySpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerRegistryStatus) DeepCopyInto(out *ContainerRegistryStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerRegistryStatus.
func (in *ContainerRegistryStatus) DeepCopy() *ContainerRegistryStatus {
	if in == nil {
		return nil
	}
	out := new(ContainerRegistryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerCredentials) DeepCopyInto(out *DockerCredentials) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerCredentials.
func (in *DockerCredentials) DeepCopy() *DockerCredentials {
	if in == nil {
		return nil
	}
	out := new(DockerCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DockerCredentials) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerCredentialsList) DeepCopyInto(out *DockerCredentialsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DockerCredentials, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerCredentialsList.
func (in *DockerCredentialsList) DeepCopy() *DockerCredentialsList {
	if in == nil {
		return nil
	}
	out := new(DockerCredentialsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DockerCredentialsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerCredentialsSpec) DeepCopyInto(out *DockerCredentialsSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(DockerCredentialsSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerCredentialsSpec.
func (in *DockerCredentialsSpec) DeepCopy() *DockerCredentialsSpec {
	if in == nil {
		return nil
	}
	out := new(DockerCredentialsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerCredentialsSpecResource) DeepCopyInto(out *DockerCredentialsSpecResource) {
	*out = *in
	if in.CredentialExpirationTime != nil {
		in, out := &in.CredentialExpirationTime, &out.CredentialExpirationTime
		*out = new(string)
		**out = **in
	}
	if in.DockerCredentials != nil {
		in, out := &in.DockerCredentials, &out.DockerCredentials
		*out = new(string)
		**out = **in
	}
	if in.ExpirySeconds != nil {
		in, out := &in.ExpirySeconds, &out.ExpirySeconds
		*out = new(int64)
		**out = **in
	}
	if in.RegistryName != nil {
		in, out := &in.RegistryName, &out.RegistryName
		*out = new(string)
		**out = **in
	}
	if in.Write != nil {
		in, out := &in.Write, &out.Write
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerCredentialsSpecResource.
func (in *DockerCredentialsSpecResource) DeepCopy() *DockerCredentialsSpecResource {
	if in == nil {
		return nil
	}
	out := new(DockerCredentialsSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerCredentialsStatus) DeepCopyInto(out *DockerCredentialsStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerCredentialsStatus.
func (in *DockerCredentialsStatus) DeepCopy() *DockerCredentialsStatus {
	if in == nil {
		return nil
	}
	out := new(DockerCredentialsStatus)
	in.DeepCopyInto(out)
	return out
}
