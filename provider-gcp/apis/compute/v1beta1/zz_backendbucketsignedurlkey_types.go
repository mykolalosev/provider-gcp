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

type BackendBucketSignedURLKeyObservation struct {

	// an identifier for the resource with format projects/{{project}}/global/backendBuckets/{{backend_bucket}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type BackendBucketSignedURLKeyParameters struct {

	// The backend bucket this signed URL key belongs.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-gcp/apis/compute/v1beta1.BackendBucket
	// +kubebuilder:validation:Optional
	BackendBucket *string `json:"backendBucket,omitempty" tf:"backend_bucket,omitempty"`

	// Reference to a BackendBucket in compute to populate backendBucket.
	// +kubebuilder:validation:Optional
	BackendBucketRef *v1.Reference `json:"backendBucketRef,omitempty" tf:"-"`

	// Selector for a BackendBucket in compute to populate backendBucket.
	// +kubebuilder:validation:Optional
	BackendBucketSelector *v1.Selector `json:"backendBucketSelector,omitempty" tf:"-"`

	// 128-bit key value used for signing the URL. The key value must be a
	// valid RFC 4648 Section 5 base64url encoded string.
	// Note: This property is sensitive and will not be displayed in the plan.
	// +kubebuilder:validation:Required
	KeyValueSecretRef v1.SecretKeySelector `json:"keyValueSecretRef" tf:"-"`

	// Name of the signed URL key.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

// BackendBucketSignedURLKeySpec defines the desired state of BackendBucketSignedURLKey
type BackendBucketSignedURLKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BackendBucketSignedURLKeyParameters `json:"forProvider"`
}

// BackendBucketSignedURLKeyStatus defines the observed state of BackendBucketSignedURLKey.
type BackendBucketSignedURLKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BackendBucketSignedURLKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BackendBucketSignedURLKey is the Schema for the BackendBucketSignedURLKeys API. A key for signing Cloud CDN signed URLs for BackendBuckets.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type BackendBucketSignedURLKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BackendBucketSignedURLKeySpec   `json:"spec"`
	Status            BackendBucketSignedURLKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackendBucketSignedURLKeyList contains a list of BackendBucketSignedURLKeys
type BackendBucketSignedURLKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackendBucketSignedURLKey `json:"items"`
}

// Repository type metadata.
var (
	BackendBucketSignedURLKey_Kind             = "BackendBucketSignedURLKey"
	BackendBucketSignedURLKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BackendBucketSignedURLKey_Kind}.String()
	BackendBucketSignedURLKey_KindAPIVersion   = BackendBucketSignedURLKey_Kind + "." + CRDGroupVersion.String()
	BackendBucketSignedURLKey_GroupVersionKind = CRDGroupVersion.WithKind(BackendBucketSignedURLKey_Kind)
)

func init() {
	SchemeBuilder.Register(&BackendBucketSignedURLKey{}, &BackendBucketSignedURLKeyList{})
}
