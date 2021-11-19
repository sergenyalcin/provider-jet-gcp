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

type InstanceGroupNamedPortObservation struct {
}

type InstanceGroupNamedPortParameters struct {

	// The name of the instance group.
	// +kubebuilder:validation:Required
	Group *string `json:"group" tf:"group,omitempty"`

	// The name for this named port. The name must be 1-63 characters
	// long, and comply with RFC1035.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The port number, which can be a value between 1 and 65535.
	// +kubebuilder:validation:Required
	Port *int64 `json:"port" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The zone of the instance group.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// InstanceGroupNamedPortSpec defines the desired state of InstanceGroupNamedPort
type InstanceGroupNamedPortSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceGroupNamedPortParameters `json:"forProvider"`
}

// InstanceGroupNamedPortStatus defines the observed state of InstanceGroupNamedPort.
type InstanceGroupNamedPortStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceGroupNamedPortObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceGroupNamedPort is the Schema for the InstanceGroupNamedPorts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type InstanceGroupNamedPort struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceGroupNamedPortSpec   `json:"spec"`
	Status            InstanceGroupNamedPortStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceGroupNamedPortList contains a list of InstanceGroupNamedPorts
type InstanceGroupNamedPortList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstanceGroupNamedPort `json:"items"`
}

// Repository type metadata.
var (
	InstanceGroupNamedPort_Kind             = "InstanceGroupNamedPort"
	InstanceGroupNamedPort_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InstanceGroupNamedPort_Kind}.String()
	InstanceGroupNamedPort_KindAPIVersion   = InstanceGroupNamedPort_Kind + "." + CRDGroupVersion.String()
	InstanceGroupNamedPort_GroupVersionKind = CRDGroupVersion.WithKind(InstanceGroupNamedPort_Kind)
)

func init() {
	SchemeBuilder.Register(&InstanceGroupNamedPort{}, &InstanceGroupNamedPortList{})
}