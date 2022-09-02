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

type DailyScheduleObservation struct {
}

type DailyScheduleParameters struct {

	// The number of days between snapshots.
	// +kubebuilder:validation:Required
	DaysInCycle *float64 `json:"daysInCycle" tf:"days_in_cycle,omitempty"`

	// +kubebuilder:validation:Required
	StartTime *string `json:"startTime" tf:"start_time,omitempty"`
}

type DayOfWeeksObservation struct {
}

type DayOfWeeksParameters struct {

	// The day of the week to create the snapshot. e.g. MONDAY
	// Possible values are MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY, and SUNDAY.
	// +kubebuilder:validation:Required
	Day *string `json:"day" tf:"day,omitempty"`

	// +kubebuilder:validation:Required
	StartTime *string `json:"startTime" tf:"start_time,omitempty"`
}

type GroupPlacementPolicyObservation struct {
}

type GroupPlacementPolicyParameters struct {

	// The number of availability domains instances will be spread across. If two instances are in different
	// availability domain, they will not be put in the same low latency network
	// +kubebuilder:validation:Optional
	AvailabilityDomainCount *float64 `json:"availabilityDomainCount,omitempty" tf:"availability_domain_count,omitempty"`

	// Collocation specifies whether to place VMs inside the same availability domain on the same low-latency network.
	// Specify COLLOCATED to enable collocation. Can only be specified with vm_count. If compute instances are created
	// with a COLLOCATED policy, then exactly vm_count instances must be created at the same time with the resource policy
	// attached.
	// Possible values are COLLOCATED.
	// +kubebuilder:validation:Optional
	Collocation *string `json:"collocation,omitempty" tf:"collocation,omitempty"`

	// Number of vms in this placement group.
	// +kubebuilder:validation:Optional
	VMCount *float64 `json:"vmCount,omitempty" tf:"vm_count,omitempty"`
}

type HourlyScheduleObservation struct {
}

type HourlyScheduleParameters struct {

	// The number of hours between snapshots.
	// +kubebuilder:validation:Required
	HoursInCycle *float64 `json:"hoursInCycle" tf:"hours_in_cycle,omitempty"`

	// +kubebuilder:validation:Required
	StartTime *string `json:"startTime" tf:"start_time,omitempty"`
}

type InstanceSchedulePolicyObservation struct {
}

type InstanceSchedulePolicyParameters struct {

	// The expiration time of the schedule. The timestamp is an RFC3339 string.
	// +kubebuilder:validation:Optional
	ExpirationTime *string `json:"expirationTime,omitempty" tf:"expiration_time,omitempty"`

	// +kubebuilder:validation:Optional
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// Specifies the time zone to be used in interpreting the schedule. The value of this field must be a time zone name
	// from the tz database: http://en.wikipedia.org/wiki/Tz_database.
	// +kubebuilder:validation:Required
	TimeZone *string `json:"timeZone" tf:"time_zone,omitempty"`

	// Specifies the schedule for starting instances.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	VMStartSchedule []VMStartScheduleParameters `json:"vmStartSchedule,omitempty" tf:"vm_start_schedule,omitempty"`

	// Specifies the schedule for stopping instances.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	VMStopSchedule []VMStopScheduleParameters `json:"vmStopSchedule,omitempty" tf:"vm_stop_schedule,omitempty"`
}

type ResourcePolicyObservation struct {

	// an identifier for the resource with format projects/{{project}}/regions/{{region}}/resourcePolicies/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
}

type ResourcePolicyParameters struct {

	// An optional description of this resource. Provide this property when you create the resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Resource policy for instances used for placement configuration.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	GroupPlacementPolicy []GroupPlacementPolicyParameters `json:"groupPlacementPolicy,omitempty" tf:"group_placement_policy,omitempty"`

	// Resource policy for scheduling instance operations.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	InstanceSchedulePolicy []InstanceSchedulePolicyParameters `json:"instanceSchedulePolicy,omitempty" tf:"instance_schedule_policy,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Region where resource policy resides.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// Policy for creating snapshots of persistent disks.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	SnapshotSchedulePolicy []SnapshotSchedulePolicyParameters `json:"snapshotSchedulePolicy,omitempty" tf:"snapshot_schedule_policy,omitempty"`
}

type RetentionPolicyObservation struct {
}

type RetentionPolicyParameters struct {

	// Maximum age of the snapshot that is allowed to be kept.
	// +kubebuilder:validation:Required
	MaxRetentionDays *float64 `json:"maxRetentionDays" tf:"max_retention_days,omitempty"`

	// Specifies the behavior to apply to scheduled snapshots when
	// the source disk is deleted.
	// Default value is KEEP_AUTO_SNAPSHOTS.
	// Possible values are KEEP_AUTO_SNAPSHOTS and APPLY_RETENTION_POLICY.
	// +kubebuilder:validation:Optional
	OnSourceDiskDelete *string `json:"onSourceDiskDelete,omitempty" tf:"on_source_disk_delete,omitempty"`
}

type ScheduleObservation struct {
}

type ScheduleParameters struct {

	// The policy will execute every nth day at the specified time.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	DailySchedule []DailyScheduleParameters `json:"dailySchedule,omitempty" tf:"daily_schedule,omitempty"`

	// The policy will execute every nth hour starting at the specified time.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	HourlySchedule []HourlyScheduleParameters `json:"hourlySchedule,omitempty" tf:"hourly_schedule,omitempty"`

	// Allows specifying a snapshot time for each day of the week.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	WeeklySchedule []WeeklyScheduleParameters `json:"weeklySchedule,omitempty" tf:"weekly_schedule,omitempty"`
}

type SnapshotPropertiesObservation struct {
}

type SnapshotPropertiesParameters struct {

	// Whether to perform a 'guest aware' snapshot.
	// +kubebuilder:validation:Optional
	GuestFlush *bool `json:"guestFlush,omitempty" tf:"guest_flush,omitempty"`

	// A set of key-value pairs.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Cloud Storage bucket location to store the auto snapshot
	// +kubebuilder:validation:Optional
	StorageLocations []*string `json:"storageLocations,omitempty" tf:"storage_locations,omitempty"`
}

type SnapshotSchedulePolicyObservation struct {
}

type SnapshotSchedulePolicyParameters struct {

	// Retention policy applied to snapshots created by this resource policy.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	RetentionPolicy []RetentionPolicyParameters `json:"retentionPolicy,omitempty" tf:"retention_policy,omitempty"`

	// +kubebuilder:validation:Required
	Schedule []ScheduleParameters `json:"schedule" tf:"schedule,omitempty"`

	// Properties with which the snapshots are created, such as labels.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	SnapshotProperties []SnapshotPropertiesParameters `json:"snapshotProperties,omitempty" tf:"snapshot_properties,omitempty"`
}

type VMStartScheduleObservation struct {
}

type VMStartScheduleParameters struct {

	// +kubebuilder:validation:Required
	Schedule *string `json:"schedule" tf:"schedule,omitempty"`
}

type VMStopScheduleObservation struct {
}

type VMStopScheduleParameters struct {

	// +kubebuilder:validation:Required
	Schedule *string `json:"schedule" tf:"schedule,omitempty"`
}

type WeeklyScheduleObservation struct {
}

type WeeklyScheduleParameters struct {

	// May contain up to seven  snapshot times.
	// Structure is documented below.
	// +kubebuilder:validation:Required
	DayOfWeeks []DayOfWeeksParameters `json:"dayOfWeeks" tf:"day_of_weeks,omitempty"`
}

// ResourcePolicySpec defines the desired state of ResourcePolicy
type ResourcePolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResourcePolicyParameters `json:"forProvider"`
}

// ResourcePolicyStatus defines the observed state of ResourcePolicy.
type ResourcePolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResourcePolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ResourcePolicy is the Schema for the ResourcePolicys API. A policy that can be attached to a resource to specify or schedule actions on that resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type ResourcePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResourcePolicySpec   `json:"spec"`
	Status            ResourcePolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourcePolicyList contains a list of ResourcePolicys
type ResourcePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourcePolicy `json:"items"`
}

// Repository type metadata.
var (
	ResourcePolicy_Kind             = "ResourcePolicy"
	ResourcePolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ResourcePolicy_Kind}.String()
	ResourcePolicy_KindAPIVersion   = ResourcePolicy_Kind + "." + CRDGroupVersion.String()
	ResourcePolicy_GroupVersionKind = CRDGroupVersion.WithKind(ResourcePolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&ResourcePolicy{}, &ResourcePolicyList{})
}
