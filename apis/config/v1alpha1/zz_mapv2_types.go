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

type MapV2Observation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	ResourceVersion *string `json:"resourceVersion,omitempty" tf:"resource_version,omitempty"`
}

type MapV2Parameters struct {

	// Annotations of the resource
	// +kubebuilder:validation:Optional
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// K8s cluster ID
	// +kubebuilder:validation:Required
	ClusterID *string `json:"clusterId" tf:"cluster_id,omitempty"`

	// ConfigMap V2 data map
	// +kubebuilder:validation:Required
	Data map[string]*string `json:"data" tf:"data,omitempty"`

	// If set to true, ensures that data stored in the ConfigMap cannot be updated
	// +kubebuilder:validation:Optional
	Immutable *bool `json:"immutable,omitempty" tf:"immutable,omitempty"`

	// Labels of the resource
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// ConfigMap V2 namespace
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`
}

// MapV2Spec defines the desired state of MapV2
type MapV2Spec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MapV2Parameters `json:"forProvider"`
}

// MapV2Status defines the observed state of MapV2.
type MapV2Status struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MapV2Observation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MapV2 is the Schema for the MapV2s API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,rancherjet}
type MapV2 struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MapV2Spec   `json:"spec"`
	Status            MapV2Status `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MapV2List contains a list of MapV2s
type MapV2List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MapV2 `json:"items"`
}

// Repository type metadata.
var (
	MapV2_Kind             = "MapV2"
	MapV2_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MapV2_Kind}.String()
	MapV2_KindAPIVersion   = MapV2_Kind + "." + CRDGroupVersion.String()
	MapV2_GroupVersionKind = CRDGroupVersion.WithKind(MapV2_Kind)
)

func init() {
	SchemeBuilder.Register(&MapV2{}, &MapV2List{})
}
