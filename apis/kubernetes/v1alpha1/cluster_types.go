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

type Cluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec,omitempty"`
	Status            ClusterStatus `json:"status,omitempty"`
}

type ClusterSpecMaintenancePolicy struct {
	// +optional
	Day *string `json:"day,omitempty" tf:"day"`
	// +optional
	Duration *string `json:"duration,omitempty" tf:"duration"`
	// +optional
	StartTime *string `json:"startTime,omitempty" tf:"start_time"`
}

type ClusterSpecNodePoolNodes struct {
	// +optional
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at"`
	// +optional
	DropletID *string `json:"dropletID,omitempty" tf:"droplet_id"`
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at"`
}

type ClusterSpecNodePoolTaint struct {
	Effect *string `json:"effect" tf:"effect"`
	Key    *string `json:"key" tf:"key"`
	Value  *string `json:"value" tf:"value"`
}

type ClusterSpecNodePool struct {
	// +optional
	ActualNodeCount *int64 `json:"actualNodeCount,omitempty" tf:"actual_node_count"`
	// +optional
	AutoScale *bool `json:"autoScale,omitempty" tf:"auto_scale"`
	// +optional
	ID *string `json:"ID,omitempty" tf:"id"`
	// +optional
	Labels *map[string]string `json:"labels,omitempty" tf:"labels"`
	// +optional
	MaxNodes *int64 `json:"maxNodes,omitempty" tf:"max_nodes"`
	// +optional
	MinNodes *int64  `json:"minNodes,omitempty" tf:"min_nodes"`
	Name     *string `json:"name" tf:"name"`
	// +optional
	NodeCount *int64 `json:"nodeCount,omitempty" tf:"node_count"`
	// +optional
	Nodes []ClusterSpecNodePoolNodes `json:"nodes,omitempty" tf:"nodes"`
	Size  *string                    `json:"size" tf:"size"`
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags"`
	// +optional
	Taint []ClusterSpecNodePoolTaint `json:"taint,omitempty" tf:"taint"`
}

type ClusterSpec struct {
	State *ClusterSpecResource `json:"state,omitempty" tf:"-"`

	Resource ClusterSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ClusterSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AutoUpgrade *bool `json:"autoUpgrade,omitempty" tf:"auto_upgrade"`
	// +optional
	ClusterSubnet *string `json:"clusterSubnet,omitempty" tf:"cluster_subnet"`
	// +optional
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at"`
	// +optional
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint"`
	// +optional
	Ha *bool `json:"ha,omitempty" tf:"ha"`
	// +optional
	Ipv4Address *string `json:"ipv4Address,omitempty" tf:"ipv4_address"`
	// +optional
	// +optional
	MaintenancePolicy *ClusterSpecMaintenancePolicy `json:"maintenancePolicy,omitempty" tf:"maintenance_policy"`
	Name              *string                       `json:"name" tf:"name"`
	NodePool          *ClusterSpecNodePool          `json:"nodePool" tf:"node_pool"`
	Region            *string                       `json:"region" tf:"region"`
	// +optional
	ServiceSubnet *string `json:"serviceSubnet,omitempty" tf:"service_subnet"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	SurgeUpgrade *bool `json:"surgeUpgrade,omitempty" tf:"surge_upgrade"`
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags"`
	// +optional
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at"`
	// +optional
	Urn     *string `json:"urn,omitempty" tf:"urn"`
	Version *string `json:"version" tf:"version"`
	// +optional
	VpcUUID *string `json:"vpcUUID,omitempty" tf:"vpc_uuid"`
}

type ClusterStatus struct {
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

// ClusterList is a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Cluster CRD objects
	Items []Cluster `json:"items,omitempty"`
}
