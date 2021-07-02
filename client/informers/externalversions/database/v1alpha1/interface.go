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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "kubeform.dev/provider-digitalocean-api/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Clusters returns a ClusterInformer.
	Clusters() ClusterInformer
	// ConnectionPools returns a ConnectionPoolInformer.
	ConnectionPools() ConnectionPoolInformer
	// Dbs returns a DbInformer.
	Dbs() DbInformer
	// Firewalls returns a FirewallInformer.
	Firewalls() FirewallInformer
	// Replicas returns a ReplicaInformer.
	Replicas() ReplicaInformer
	// Users returns a UserInformer.
	Users() UserInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Clusters returns a ClusterInformer.
func (v *version) Clusters() ClusterInformer {
	return &clusterInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ConnectionPools returns a ConnectionPoolInformer.
func (v *version) ConnectionPools() ConnectionPoolInformer {
	return &connectionPoolInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Dbs returns a DbInformer.
func (v *version) Dbs() DbInformer {
	return &dbInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Firewalls returns a FirewallInformer.
func (v *version) Firewalls() FirewallInformer {
	return &firewallInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Replicas returns a ReplicaInformer.
func (v *version) Replicas() ReplicaInformer {
	return &replicaInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Users returns a UserInformer.
func (v *version) Users() UserInformer {
	return &userInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
