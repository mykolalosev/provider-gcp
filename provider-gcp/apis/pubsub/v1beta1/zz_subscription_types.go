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

type DeadLetterPolicyObservation struct {
}

type DeadLetterPolicyParameters struct {

	// The name of the topic to which dead letter messages should be published.
	// Format is projects/{project}/topics/{topic}.
	// The Cloud Pub/Sub service account associated with the enclosing subscription's
	// parent project  must have
	// permission to Publish() to this topic.
	// The operation will fail if the topic does not exist.
	// Users should ensure that there is a subscription attached to this topic
	// since messages published to a topic with no subscriptions are lost.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-gcp/apis/pubsub/v1beta1.Topic
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DeadLetterTopic *string `json:"deadLetterTopic,omitempty" tf:"dead_letter_topic,omitempty"`

	// +kubebuilder:validation:Optional
	DeadLetterTopicRef *v1.Reference `json:"deadLetterTopicRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	DeadLetterTopicSelector *v1.Selector `json:"deadLetterTopicSelector,omitempty" tf:"-"`

	// The maximum number of delivery attempts for any message. The value must be
	// between 5 and 100.
	// The number of delivery attempts is defined as 1 + .
	// A NACK is any call to ModifyAckDeadline with a 0 deadline. Note that
	// client libraries may automatically extend ack_deadlines.
	// This field will be honored on a best effort basis.
	// If this parameter is 0, a default value of 5 is used.
	// +kubebuilder:validation:Optional
	MaxDeliveryAttempts *float64 `json:"maxDeliveryAttempts,omitempty" tf:"max_delivery_attempts,omitempty"`
}

type ExpirationPolicyObservation struct {
}

type ExpirationPolicyParameters struct {

	// Specifies the "time-to-live" duration for an associated resource. The
	// resource expires if it is not active for a period of ttl.
	// If ttl is not set, the associated resource never expires.
	// A duration in seconds with up to nine fractional digits, terminated by 's'.
	// Example - "3.5s".
	// +kubebuilder:validation:Required
	TTL *string `json:"ttl" tf:"ttl,omitempty"`
}

type OidcTokenObservation struct {
}

type OidcTokenParameters struct {

	// Audience to be used when generating OIDC token. The audience claim
	// identifies the recipients that the JWT is intended for. The audience
	// value is a single case-sensitive string. Having multiple values
	// for the audience field is not supported. More info about the OIDC JWT
	// token audience here: https://tools.ietf.org/html/rfc7519#section-4.1.3
	// Note: if not specified, the Push endpoint URL will be used.
	// +kubebuilder:validation:Optional
	Audience *string `json:"audience,omitempty" tf:"audience,omitempty"`

	// Service account email to be used for generating the OIDC token.
	// The caller  must have the
	// iam.serviceAccounts.actAs permission for the service account.
	// +kubebuilder:validation:Required
	ServiceAccountEmail *string `json:"serviceAccountEmail" tf:"service_account_email,omitempty"`
}

type PushConfigObservation struct {
}

type PushConfigParameters struct {

	// Endpoint configuration attributes.
	// Every endpoint has a set of API supported attributes that can
	// be used to control different aspects of the message delivery.
	// The currently supported attribute is x-goog-version, which you
	// can use to change the format of the pushed message. This
	// attribute indicates the version of the data expected by
	// the endpoint. This controls the shape of the pushed message
	// . The endpoint version is
	// based on the version of the Pub/Sub API.
	// If not present during the subscriptions.create call,
	// it will default to the version of the API used to make
	// such call. If not present during a subscriptions.modifyPushConfig
	// call, its value will not be changed. subscriptions.get
	// calls will always return a valid version, even if the
	// subscription was created without this attribute.
	// The possible values for this attribute are:
	// +kubebuilder:validation:Optional
	Attributes map[string]*string `json:"attributes,omitempty" tf:"attributes,omitempty"`

	// If specified, Pub/Sub will generate and attach an OIDC JWT token as
	// an Authorization header in the HTTP request for every pushed message.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	OidcToken []OidcTokenParameters `json:"oidcToken,omitempty" tf:"oidc_token,omitempty"`

	// A URL locating the endpoint to which messages should be pushed.
	// For example, a Webhook endpoint might use
	// "https://example.com/push".
	// +kubebuilder:validation:Required
	PushEndpoint *string `json:"pushEndpoint" tf:"push_endpoint,omitempty"`
}

type RetryPolicyObservation struct {
}

type RetryPolicyParameters struct {

	// The maximum delay between consecutive deliveries of a given message. Value should be between 0 and 600 seconds. Defaults to 600 seconds.
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	// +kubebuilder:validation:Optional
	MaximumBackoff *string `json:"maximumBackoff,omitempty" tf:"maximum_backoff,omitempty"`

	// The minimum delay between consecutive deliveries of a given message. Value should be between 0 and 600 seconds. Defaults to 10 seconds.
	// A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	// +kubebuilder:validation:Optional
	MinimumBackoff *string `json:"minimumBackoff,omitempty" tf:"minimum_backoff,omitempty"`
}

type SubscriptionObservation struct {

	// an identifier for the resource with format projects/{{project}}/subscriptions/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SubscriptionParameters struct {

	// This value is the maximum time after a subscriber receives a message
	// before the subscriber should acknowledge the message. After message
	// delivery but before the ack deadline expires and before the message is
	// acknowledged, it is an outstanding message and will not be delivered
	// again during that time .
	// For pull subscriptions, this value is used as the initial value for
	// the ack deadline. To override this value for a given message, call
	// subscriptions.modifyAckDeadline with the corresponding ackId if using
	// pull. The minimum custom deadline you can specify is 10 seconds. The
	// maximum custom deadline you can specify is 600 seconds .
	// If this parameter is 0, a default value of 10 seconds is used.
	// For push delivery, this value is also used to set the request timeout
	// for the call to the push endpoint.
	// If the subscriber never acknowledges the message, the Pub/Sub system
	// will eventually redeliver the message.
	// +kubebuilder:validation:Optional
	AckDeadlineSeconds *float64 `json:"ackDeadlineSeconds,omitempty" tf:"ack_deadline_seconds,omitempty"`

	// A policy that specifies the conditions for dead lettering messages in
	// this subscription. If dead_letter_policy is not set, dead lettering
	// is disabled.
	// The Cloud Pub/Sub service account associated with this subscription's
	// parent project  must have
	// permission to Acknowledge() messages on this subscription.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	DeadLetterPolicy []DeadLetterPolicyParameters `json:"deadLetterPolicy,omitempty" tf:"dead_letter_policy,omitempty"`

	// If 'true', Pub/Sub provides the following guarantees for the delivery
	// of a message with a given value of messageId on this Subscriptions':
	//
	// - The message sent to a subscriber is guaranteed not to be resent before the message's acknowledgement deadline expires.
	//
	// - An acknowledged message will not be resent to a subscriber.
	//
	// Note that subscribers may still receive multiple copies of a message when 'enable_exactly_once_delivery'
	// is true if the message was published multiple times by a publisher client. These copies are considered distinct by Pub/Sub and have distinct messageId values
	// +kubebuilder:validation:Optional
	EnableExactlyOnceDelivery *bool `json:"enableExactlyOnceDelivery,omitempty" tf:"enable_exactly_once_delivery,omitempty"`

	// If true, messages published with the same orderingKey in PubsubMessage will be delivered to
	// the subscribers in the order in which they are received by the Pub/Sub system. Otherwise, they
	// may be delivered in any order.
	// +kubebuilder:validation:Optional
	EnableMessageOrdering *bool `json:"enableMessageOrdering,omitempty" tf:"enable_message_ordering,omitempty"`

	// A policy that specifies the conditions for this subscription's expiration.
	// A subscription is considered active as long as any connected subscriber
	// is successfully consuming messages from the subscription or is issuing
	// operations on the subscription. If expirationPolicy is not set, a default
	// policy with ttl of 31 days will be used.  If it is set but ttl is "", the
	// resource never expires.  The minimum allowed value for expirationPolicy.ttl
	// is 1 day.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	ExpirationPolicy []ExpirationPolicyParameters `json:"expirationPolicy,omitempty" tf:"expiration_policy,omitempty"`

	// The subscription only delivers the messages that match the filter.
	// Pub/Sub automatically acknowledges the messages that don't match the filter. You can filter messages
	// by their attributes. The maximum length of a filter is 256 bytes. After creating the subscription,
	// you can't modify the filter.
	// +kubebuilder:validation:Optional
	Filter *string `json:"filter,omitempty" tf:"filter,omitempty"`

	// A set of key/value label pairs to assign to this Subscription.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// How long to retain unacknowledged messages in the subscription's
	// backlog, from the moment a message is published. If
	// retain_acked_messages is true, then this also configures the retention
	// of acknowledged messages, and thus configures how far back in time a
	// subscriptions.seek can be done. Defaults to 7 days. Cannot be more
	// than 7 days  or less than 10 minutes .
	// A duration in seconds with up to nine fractional digits, terminated
	// by 's'. Example: "600.5s".
	// +kubebuilder:validation:Optional
	MessageRetentionDuration *string `json:"messageRetentionDuration,omitempty" tf:"message_retention_duration,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// If push delivery is used with this subscription, this field is used to
	// configure it. An empty pushConfig signifies that the subscriber will
	// pull and ack messages using API methods.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	PushConfig []PushConfigParameters `json:"pushConfig,omitempty" tf:"push_config,omitempty"`

	// Indicates whether to retain acknowledged messages. If true, then
	// messages are not expunged from the subscription's backlog, even if
	// they are acknowledged, until they fall out of the
	// messageRetentionDuration window.
	// +kubebuilder:validation:Optional
	RetainAckedMessages *bool `json:"retainAckedMessages,omitempty" tf:"retain_acked_messages,omitempty"`

	// A policy that specifies how Pub/Sub retries message delivery for this subscription.
	// If not set, the default retry policy is applied. This generally implies that messages will be retried as soon as possible for healthy subscribers.
	// RetryPolicy will be triggered on NACKs or acknowledgement deadline exceeded events for a given message
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	RetryPolicy []RetryPolicyParameters `json:"retryPolicy,omitempty" tf:"retry_policy,omitempty"`

	// A reference to a Topic resource.
	// +crossplane:generate:reference:type=Topic
	// +kubebuilder:validation:Optional
	Topic *string `json:"topic,omitempty" tf:"topic,omitempty"`

	// +kubebuilder:validation:Optional
	TopicRef *v1.Reference `json:"topicRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	TopicSelector *v1.Selector `json:"topicSelector,omitempty" tf:"-"`
}

// SubscriptionSpec defines the desired state of Subscription
type SubscriptionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubscriptionParameters `json:"forProvider"`
}

// SubscriptionStatus defines the observed state of Subscription.
type SubscriptionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubscriptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Subscription is the Schema for the Subscriptions API. A named resource representing the stream of messages from a single, specific topic, to be delivered to the subscribing application.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type Subscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SubscriptionSpec   `json:"spec"`
	Status            SubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubscriptionList contains a list of Subscriptions
type SubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Subscription `json:"items"`
}

// Repository type metadata.
var (
	Subscription_Kind             = "Subscription"
	Subscription_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Subscription_Kind}.String()
	Subscription_KindAPIVersion   = Subscription_Kind + "." + CRDGroupVersion.String()
	Subscription_GroupVersionKind = CRDGroupVersion.WithKind(Subscription_Kind)
)

func init() {
	SchemeBuilder.Register(&Subscription{}, &SubscriptionList{})
}
