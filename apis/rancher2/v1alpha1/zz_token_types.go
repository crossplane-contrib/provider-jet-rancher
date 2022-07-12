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

type TokenObservation struct {
	AccessKey *string `json:"accessKey,omitempty" tf:"access_key,omitempty"`

	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	Expired *bool `json:"expired,omitempty" tf:"expired,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

type TokenParameters struct {

	// Annotations of the resource
	// +kubebuilder:validation:Optional
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// Cluster ID for scoped token
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Token description
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Labels of the resource
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Renew expired or disabled token
	// +kubebuilder:validation:Optional
	Renew *bool `json:"renew,omitempty" tf:"renew,omitempty"`

	// Token time to live in seconds
	// +kubebuilder:validation:Optional
	TTL *float64 `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

// TokenSpec defines the desired state of Token
type TokenSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TokenParameters `json:"forProvider"`
}

// TokenStatus defines the observed state of Token.
type TokenStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TokenObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Token is the Schema for the Tokens API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,rancherjet}
type Token struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TokenSpec   `json:"spec"`
	Status            TokenStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TokenList contains a list of Tokens
type TokenList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Token `json:"items"`
}

// Repository type metadata.
var (
	Token_Kind             = "Token"
	Token_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Token_Kind}.String()
	Token_KindAPIVersion   = Token_Kind + "." + CRDGroupVersion.String()
	Token_GroupVersionKind = CRDGroupVersion.WithKind(Token_Kind)
)

func init() {
	SchemeBuilder.Register(&Token{}, &TokenList{})
}
