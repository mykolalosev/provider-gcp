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

type ServiceAccountKeyObservation struct {

	// an identifier for the resource with format projects/{{project}}/serviceAccounts/{{account}}/keys/{{key}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name used for this key pair
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The public key, base64 encoded
	PublicKey *string `json:"publicKey,omitempty" tf:"public_key,omitempty"`

	// The key can be used after this timestamp. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	ValidAfter *string `json:"validAfter,omitempty" tf:"valid_after,omitempty"`

	// The key can be used before this timestamp.
	// A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	ValidBefore *string `json:"validBefore,omitempty" tf:"valid_before,omitempty"`
}

type ServiceAccountKeyParameters struct {

	// Arbitrary map of values that, when changed, will trigger a new key to be generated.
	// +kubebuilder:validation:Optional
	Keepers map[string]string `json:"keepers,omitempty" tf:"keepers,omitempty"`

	// The algorithm used to generate the key. KEY_ALG_RSA_2048 is the default algorithm.
	// Valid values are listed at
	// ServiceAccountPrivateKeyType
	// +kubebuilder:validation:Optional
	KeyAlgorithm *string `json:"keyAlgorithm,omitempty" tf:"key_algorithm,omitempty"`

	// The output format of the private key. TYPE_GOOGLE_CREDENTIALS_FILE is the default output format.
	// +kubebuilder:validation:Optional
	PrivateKeyType *string `json:"privateKeyType,omitempty" tf:"private_key_type,omitempty"`

	// Public key data to create a service account key for given service account. The expected format for this field is a base64 encoded X509_PEM and it conflicts with public_key_type and private_key_type.
	// +kubebuilder:validation:Optional
	PublicKeyData *string `json:"publicKeyData,omitempty" tf:"public_key_data,omitempty"`

	// The output format of the public key requested. TYPE_X509_PEM_FILE is the default output format.
	// +kubebuilder:validation:Optional
	PublicKeyType *string `json:"publicKeyType,omitempty" tf:"public_key_type,omitempty"`

	// The Service account id of the Key. This can be a string in the format
	// {ACCOUNT} or projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}. If the {ACCOUNT}-only syntax is used, either
	// the full email address of the service account or its name can be specified as a value, in which case the project will
	// automatically be inferred from the account. Otherwise, if the projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}
	// syntax is used, the {ACCOUNT} specified can be the full email address of the service account or the service account's
	// unique id. Substituting - as a wildcard for the {PROJECT_ID} will infer the project from the account.
	// +crossplane:generate:reference:type=ServiceAccount
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-gcp/config/common.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ServiceAccountID *string `json:"serviceAccountId,omitempty" tf:"service_account_id,omitempty"`

	// Reference to a ServiceAccount to populate serviceAccountId.
	// +kubebuilder:validation:Optional
	ServiceAccountIDRef *v1.Reference `json:"serviceAccountIdRef,omitempty" tf:"-"`

	// Selector for a ServiceAccount to populate serviceAccountId.
	// +kubebuilder:validation:Optional
	ServiceAccountIDSelector *v1.Selector `json:"serviceAccountIdSelector,omitempty" tf:"-"`
}

// ServiceAccountKeySpec defines the desired state of ServiceAccountKey
type ServiceAccountKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceAccountKeyParameters `json:"forProvider"`
}

// ServiceAccountKeyStatus defines the observed state of ServiceAccountKey.
type ServiceAccountKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceAccountKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceAccountKey is the Schema for the ServiceAccountKeys API. Allows management of a Google Cloud Platform service account Key
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type ServiceAccountKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceAccountKeySpec   `json:"spec"`
	Status            ServiceAccountKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceAccountKeyList contains a list of ServiceAccountKeys
type ServiceAccountKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceAccountKey `json:"items"`
}

// Repository type metadata.
var (
	ServiceAccountKey_Kind             = "ServiceAccountKey"
	ServiceAccountKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServiceAccountKey_Kind}.String()
	ServiceAccountKey_KindAPIVersion   = ServiceAccountKey_Kind + "." + CRDGroupVersion.String()
	ServiceAccountKey_GroupVersionKind = CRDGroupVersion.WithKind(ServiceAccountKey_Kind)
)

func init() {
	SchemeBuilder.Register(&ServiceAccountKey{}, &ServiceAccountKeyList{})
}
