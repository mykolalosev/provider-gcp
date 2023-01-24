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

type FilesObservation struct {
}

type FilesParameters struct {

	// Textual Content.
	// +kubebuilder:validation:Required
	Content *string `json:"content" tf:"content,omitempty"`

	// Fingerprint (e.g. github sha) associated with the File.
	// +kubebuilder:validation:Optional
	Fingerprint *string `json:"fingerprint,omitempty" tf:"fingerprint,omitempty"`

	// Output only. Name of the Ruleset. The ruleset_id is auto generated by the service. Format: projects/{project_id}/rulesets/{ruleset_id}
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type MetadataObservation struct {
	Services []*string `json:"services,omitempty" tf:"services,omitempty"`
}

type MetadataParameters struct {
}

type RulesetObservation struct {

	// Output only. Time the Ruleset was created.
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// an identifier for the resource with format projects/{{project}}/rulesets/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Output only. The metadata for this ruleset.
	Metadata []MetadataObservation `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// Output only. Name of the Ruleset. The ruleset_id is auto generated by the service. Format: projects/{project_id}/rulesets/{ruleset_id}
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type RulesetParameters struct {

	// The project for the resource
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Source for the Ruleset.
	// +kubebuilder:validation:Required
	Source []SourceParameters `json:"source" tf:"source,omitempty"`
}

type SourceObservation struct {
}

type SourceParameters struct {

	// File set constituting the Source bundle.
	// +kubebuilder:validation:Required
	Files []FilesParameters `json:"files" tf:"files,omitempty"`

	// Language of the Source bundle. If unspecified, the language will default to FIREBASE_RULES. Possible values: LANGUAGE_UNSPECIFIED, FIREBASE_RULES, EVENT_FLOW_TRIGGERS
	// +kubebuilder:validation:Optional
	Language *string `json:"language,omitempty" tf:"language,omitempty"`
}

// RulesetSpec defines the desired state of Ruleset
type RulesetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RulesetParameters `json:"forProvider"`
}

// RulesetStatus defines the observed state of Ruleset.
type RulesetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RulesetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Ruleset is the Schema for the Rulesets API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Ruleset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RulesetSpec   `json:"spec"`
	Status            RulesetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RulesetList contains a list of Rulesets
type RulesetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ruleset `json:"items"`
}

// Repository type metadata.
var (
	Ruleset_Kind             = "Ruleset"
	Ruleset_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Ruleset_Kind}.String()
	Ruleset_KindAPIVersion   = Ruleset_Kind + "." + CRDGroupVersion.String()
	Ruleset_GroupVersionKind = CRDGroupVersion.WithKind(Ruleset_Kind)
)

func init() {
	SchemeBuilder.Register(&Ruleset{}, &RulesetList{})
}
