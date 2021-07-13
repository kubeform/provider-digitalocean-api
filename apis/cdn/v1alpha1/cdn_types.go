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

type Cdn struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CdnSpec   `json:"spec,omitempty"`
	Status            CdnStatus `json:"status,omitempty"`
}

type CdnSpec struct {
	State *CdnSpecResource `json:"state,omitempty" tf:"-"`

	Resource CdnSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type CdnSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// ID of a DigitalOcean managed TLS certificate for use with custom domains
	// +optional
	// Deprecated
	CertificateID *string `json:"certificateID,omitempty" tf:"certificate_id"`
	// +optional
	CertificateName *string `json:"certificateName,omitempty" tf:"certificate_name"`
	// The date and time (ISO8601) of when the CDN endpoint was created.
	// +optional
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at"`
	// fully qualified domain name (FQDN) for custom subdomain, (requires certificate_id)
	// +optional
	CustomDomain *string `json:"customDomain,omitempty" tf:"custom_domain"`
	// fully qualified domain name (FQDN) to serve the CDN content
	// +optional
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint"`
	// fully qualified domain name (FQDN) for the origin server
	Origin *string `json:"origin" tf:"origin"`
	// The amount of time the content is cached in the CDN
	// +optional
	Ttl *int64 `json:"ttl,omitempty" tf:"ttl"`
}

type CdnStatus struct {
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

// CdnList is a list of Cdns
type CdnList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Cdn CRD objects
	Items []Cdn `json:"items,omitempty"`
}
