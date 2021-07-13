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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type Firewall struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FirewallSpec   `json:"spec,omitempty"`
	Status            FirewallStatus `json:"status,omitempty"`
}

type FirewallSpecInboundRule struct {
	// +optional
	PortRange *string `json:"portRange,omitempty" tf:"port_range"`
	Protocol  *string `json:"protocol" tf:"protocol"`
	// +optional
	SourceAddresses []string `json:"sourceAddresses,omitempty" tf:"source_addresses"`
	// +optional
	SourceDropletIDS []int64 `json:"sourceDropletIDS,omitempty" tf:"source_droplet_ids"`
	// +optional
	SourceLoadBalancerUids []string `json:"sourceLoadBalancerUids,omitempty" tf:"source_load_balancer_uids"`
	// +optional
	SourceTags []string `json:"sourceTags,omitempty" tf:"source_tags"`
}

type FirewallSpecOutboundRule struct {
	// +optional
	DestinationAddresses []string `json:"destinationAddresses,omitempty" tf:"destination_addresses"`
	// +optional
	DestinationDropletIDS []int64 `json:"destinationDropletIDS,omitempty" tf:"destination_droplet_ids"`
	// +optional
	DestinationLoadBalancerUids []string `json:"destinationLoadBalancerUids,omitempty" tf:"destination_load_balancer_uids"`
	// +optional
	DestinationTags []string `json:"destinationTags,omitempty" tf:"destination_tags"`
	// +optional
	PortRange *string `json:"portRange,omitempty" tf:"port_range"`
	Protocol  *string `json:"protocol" tf:"protocol"`
}

type FirewallSpecPendingChanges struct {
	// +optional
	DropletID *int64 `json:"dropletID,omitempty" tf:"droplet_id"`
	// +optional
	Removing *bool `json:"removing,omitempty" tf:"removing"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
}

type FirewallSpec struct {
	State *FirewallSpecResource `json:"state,omitempty" tf:"-"`

	Resource FirewallSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type FirewallSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at"`
	// +optional
	DropletIDS []int64 `json:"dropletIDS,omitempty" tf:"droplet_ids"`
	// +optional
	InboundRule []FirewallSpecInboundRule `json:"inboundRule,omitempty" tf:"inbound_rule"`
	Name        *string                   `json:"name" tf:"name"`
	// +optional
	OutboundRule []FirewallSpecOutboundRule `json:"outboundRule,omitempty" tf:"outbound_rule"`
	// +optional
	PendingChanges []FirewallSpecPendingChanges `json:"pendingChanges,omitempty" tf:"pending_changes"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags"`
}

type FirewallStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// FirewallList is a list of Firewalls
type FirewallList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Firewall CRD objects
	Items []Firewall `json:"items,omitempty"`
}
