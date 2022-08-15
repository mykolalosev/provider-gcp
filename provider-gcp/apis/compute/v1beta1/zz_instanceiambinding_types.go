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

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type InstanceIAMBindingConditionObservation struct {
}

type InstanceIAMBindingConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Required
	Title *string `json:"title" tf:"title,omitempty"`
}

type InstanceIAMBindingObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type InstanceIAMBindingParameters struct {

	// +kubebuilder:validation:Optional
	Condition []InstanceIAMBindingConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Required
	InstanceName *string `json:"instanceName" tf:"instance_name,omitempty"`

	// +kubebuilder:validation:Required
	Members []*string `json:"members" tf:"members,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`

	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// InstanceIAMBindingSpec defines the desired state of InstanceIAMBinding
type InstanceIAMBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     InstanceIAMBindingParameters `json:"forProvider"`
}

// InstanceIAMBindingStatus defines the observed state of InstanceIAMBinding.
type InstanceIAMBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        InstanceIAMBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceIAMBinding is the Schema for the InstanceIAMBindings API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type InstanceIAMBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceIAMBindingSpec   `json:"spec"`
	Status            InstanceIAMBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// InstanceIAMBindingList contains a list of InstanceIAMBindings
type InstanceIAMBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstanceIAMBinding `json:"items"`
}

// Repository type metadata.
var (
	InstanceIAMBinding_Kind             = "InstanceIAMBinding"
	InstanceIAMBinding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: InstanceIAMBinding_Kind}.String()
	InstanceIAMBinding_KindAPIVersion   = InstanceIAMBinding_Kind + "." + CRDGroupVersion.String()
	InstanceIAMBinding_GroupVersionKind = CRDGroupVersion.WithKind(InstanceIAMBinding_Kind)
)

func init() {
	SchemeBuilder.Register(&InstanceIAMBinding{}, &InstanceIAMBindingList{})
}
