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

type ProjectMetadataObservation struct {
}

type ProjectMetadataParameters struct {

	// A series of key value pairs.
	// +kubebuilder:validation:Required
	Metadata map[string]*string `json:"metadata" tf:"metadata,omitempty"`

	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

// ProjectMetadataSpec defines the desired state of ProjectMetadata
type ProjectMetadataSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectMetadataParameters `json:"forProvider"`
}

// ProjectMetadataStatus defines the observed state of ProjectMetadata.
type ProjectMetadataStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectMetadataObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectMetadata is the Schema for the ProjectMetadatas API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type ProjectMetadata struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProjectMetadataSpec   `json:"spec"`
	Status            ProjectMetadataStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectMetadataList contains a list of ProjectMetadatas
type ProjectMetadataList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectMetadata `json:"items"`
}

// Repository type metadata.
var (
	ProjectMetadata_Kind             = "ProjectMetadata"
	ProjectMetadata_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProjectMetadata_Kind}.String()
	ProjectMetadata_KindAPIVersion   = ProjectMetadata_Kind + "." + CRDGroupVersion.String()
	ProjectMetadata_GroupVersionKind = CRDGroupVersion.WithKind(ProjectMetadata_Kind)
)

func init() {
	SchemeBuilder.Register(&ProjectMetadata{}, &ProjectMetadataList{})
}