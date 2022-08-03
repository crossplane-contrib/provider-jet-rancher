/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AzureADObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type AzureADParameters struct {

	// +kubebuilder:validation:Optional
	AccessMode *string `json:"accessMode,omitempty" tf:"access_mode,omitempty"`

	// +kubebuilder:validation:Optional
	AllowedPrincipalIds []*string `json:"allowedPrincipalIds,omitempty" tf:"allowed_principal_ids,omitempty"`

	// Annotations of the resource
	// +kubebuilder:validation:Optional
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// +kubebuilder:validation:Required
	ApplicationIDSecretRef v1.SecretKeySelector `json:"applicationIdSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	ApplicationSecretSecretRef v1.SecretKeySelector `json:"applicationSecretSecretRef" tf:"-"`

	// +kubebuilder:validation:Required
	AuthEndpoint *string `json:"authEndpoint" tf:"auth_endpoint,omitempty"`

	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// +kubebuilder:validation:Required
	GraphEndpoint *string `json:"graphEndpoint" tf:"graph_endpoint,omitempty"`

	// Labels of the resource
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Required
	RancherURL *string `json:"rancherUrl" tf:"rancher_url,omitempty"`

	// +kubebuilder:validation:Required
	TenantID *string `json:"tenantId" tf:"tenant_id,omitempty"`

	// +kubebuilder:validation:Required
	TokenEndpoint *string `json:"tokenEndpoint" tf:"token_endpoint,omitempty"`
}

// AzureADSpec defines the desired state of AzureAD
type AzureADSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AzureADParameters `json:"forProvider"`
}

// AzureADStatus defines the observed state of AzureAD.
type AzureADStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AzureADObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AzureAD is the Schema for the AzureADs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,rancherjet}
type AzureAD struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzureADSpec   `json:"spec"`
	Status            AzureADStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AzureADList contains a list of AzureADs
type AzureADList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AzureAD `json:"items"`
}

// Repository type metadata.
var (
	AzureAD_Kind             = "AzureAD"
	AzureAD_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AzureAD_Kind}.String()
	AzureAD_KindAPIVersion   = AzureAD_Kind + "." + CRDGroupVersion.String()
	AzureAD_GroupVersionKind = CRDGroupVersion.WithKind(AzureAD_Kind)
)

func init() {
	SchemeBuilder.Register(&AzureAD{}, &AzureADList{})
}