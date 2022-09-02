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

type ActiveDirectoryConfigObservation struct {
}

type ActiveDirectoryConfigParameters struct {

	// The domain name for the active directory .
	// Can only be used with SQL Server.
	// +kubebuilder:validation:Required
	Domain *string `json:"domain" tf:"domain,omitempty"`
}

type AuthorizedNetworksObservation struct {
}

type AuthorizedNetworksParameters struct {

	// The RFC 3339
	// formatted date time string indicating when this whitelist expires.
	// +kubebuilder:validation:Optional
	ExpirationTime *string `json:"expirationTime,omitempty" tf:"expiration_time,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type BackupConfigurationObservation struct {
}

type BackupConfigurationParameters struct {

	// Backup retention settings. The configuration is detailed below.
	// +kubebuilder:validation:Optional
	BackupRetentionSettings []BackupRetentionSettingsParameters `json:"backupRetentionSettings,omitempty" tf:"backup_retention_settings,omitempty"`

	// True if binary logging is enabled.
	// Can only be used with MySQL.
	// +kubebuilder:validation:Optional
	BinaryLogEnabled *bool `json:"binaryLogEnabled,omitempty" tf:"binary_log_enabled,omitempty"`

	// True if backup configuration is enabled.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The region where the backup will be stored
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// True if Point-in-time recovery is enabled. Will restart database if enabled after instance creation. Valid only for PostgreSQL instances.
	// +kubebuilder:validation:Optional
	PointInTimeRecoveryEnabled *bool `json:"pointInTimeRecoveryEnabled,omitempty" tf:"point_in_time_recovery_enabled,omitempty"`

	// HH:MM format time indicating when backup
	// configuration starts.
	// +kubebuilder:validation:Optional
	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`

	// The number of days of transaction logs we retain for point in time restore, from 1-7.
	// +kubebuilder:validation:Optional
	TransactionLogRetentionDays *float64 `json:"transactionLogRetentionDays,omitempty" tf:"transaction_log_retention_days,omitempty"`
}

type BackupRetentionSettingsObservation struct {
}

type BackupRetentionSettingsParameters struct {

	// Depending on the value of retention_unit, this is used to determine if a backup needs to be deleted. If retention_unit
	// is 'COUNT', we will retain this many backups.
	// +kubebuilder:validation:Required
	RetainedBackups *float64 `json:"retainedBackups" tf:"retained_backups,omitempty"`

	// The unit that 'retained_backups' represents. Defaults to COUNT.
	// +kubebuilder:validation:Optional
	RetentionUnit *string `json:"retentionUnit,omitempty" tf:"retention_unit,omitempty"`
}

type CloneObservation struct {
}

type CloneParameters struct {

	// +kubebuilder:validation:Optional
	AllocatedIPRange *string `json:"allocatedIpRange,omitempty" tf:"allocated_ip_range,omitempty"`

	// The timestamp of the point in time that should be restored.
	// +kubebuilder:validation:Optional
	PointInTime *string `json:"pointInTime,omitempty" tf:"point_in_time,omitempty"`

	// Name of the source instance which will be cloned.
	// +kubebuilder:validation:Required
	SourceInstanceName *string `json:"sourceInstanceName" tf:"source_instance_name,omitempty"`
}

type DatabaseFlagsObservation struct {
}

type DatabaseFlagsParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type DatabaseInstanceObservation struct {

	// The connection name of the instance to be used in
	// connection strings. For example, when connecting with Cloud SQL Proxy.
	ConnectionName *string `json:"connectionName,omitempty" tf:"connection_name,omitempty"`

	// The first IPv4 address of any type assigned. This is to
	// support accessing the first address in the list in a terraform output
	// when the resource is configured with a count.
	FirstIPAddress *string `json:"firstIpAddress,omitempty" tf:"first_ip_address,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	IPAddress []IPAddressObservation `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	// The first private  IPv4 address assigned. This is
	// a workaround for an issue fixed in Terraform 0.12
	// but also provides a convenient way to access an IP of a specific type without
	// performing filtering in a Terraform config.
	PrivateIPAddress *string `json:"privateIpAddress,omitempty" tf:"private_ip_address,omitempty"`

	// The first public  IPv4 address assigned. This is
	// a workaround for an issue fixed in Terraform 0.12
	// but also provides a convenient way to access an IP of a specific type without
	// performing filtering in a Terraform config.
	PublicIPAddress *string `json:"publicIpAddress,omitempty" tf:"public_ip_address,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	ServerCACert []ServerCACertObservation `json:"serverCaCert,omitempty" tf:"server_ca_cert,omitempty"`

	// The service account email address assigned to the
	// instance.
	ServiceAccountEmailAddress *string `json:"serviceAccountEmailAddress,omitempty" tf:"service_account_email_address,omitempty"`

	// The settings to use for the database. The
	// configuration is detailed below. Required if clone is not set.
	// +kubebuilder:validation:Optional
	Settings []SettingsObservation `json:"settings,omitempty" tf:"settings,omitempty"`
}

type DatabaseInstanceParameters struct {

	// The context needed to create this instance as a clone of another instance. When this field is set during
	// resource creation, Terraform will attempt to clone another instance as indicated in the context. The
	// configuration is detailed below.
	// +kubebuilder:validation:Optional
	Clone []CloneParameters `json:"clone,omitempty" tf:"clone,omitempty"`

	// The MySQL, PostgreSQL or
	// SQL Server version to use. Supported values include MYSQL_5_6,
	// MYSQL_5_7, MYSQL_8_0, POSTGRES_9_6,POSTGRES_10, POSTGRES_11,
	// POSTGRES_12, POSTGRES_13, SQLSERVER_2017_STANDARD,
	// SQLSERVER_2017_ENTERPRISE, SQLSERVER_2017_EXPRESS, SQLSERVER_2017_WEB.
	// SQLSERVER_2019_STANDARD, SQLSERVER_2019_ENTERPRISE, SQLSERVER_2019_EXPRESS,
	// SQLSERVER_2019_WEB.
	// Database Version Policies
	// includes an up-to-date reference of supported versions.
	// +kubebuilder:validation:Required
	DatabaseVersion *string `json:"databaseVersion" tf:"database_version,omitempty"`

	// Whether or not to allow Terraform to destroy the instance. Unless this field is set to false
	// in Terraform state, a terraform destroy or terraform apply command that deletes the instance will fail.
	// +kubebuilder:validation:Optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// The name of the existing instance that will
	// act as the master in the replication setup. Note, this requires the master to
	// have binary_log_enabled set, as well as existing backups.
	// +kubebuilder:validation:Optional
	MasterInstanceName *string `json:"masterInstanceName,omitempty" tf:"master_instance_name,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region the instance will sit in. If a region is not provided in the resource definition,
	// the provider region will be used instead.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// The configuration for replication. The
	// configuration is detailed below. Valid only for MySQL instances.
	// +kubebuilder:validation:Optional
	ReplicaConfiguration []ReplicaConfigurationParameters `json:"replicaConfiguration,omitempty" tf:"replica_configuration,omitempty"`

	// The context needed to restore the database to a backup run. This field will
	// cause Terraform to trigger the database to restore from the backup run indicated. The configuration is detailed below.
	// NOTE: Restoring from a backup is an imperative action and not recommended via Terraform. Adding or modifying this
	// block during resource creation/update will trigger the restore action after the resource is created/updated.
	// +kubebuilder:validation:Optional
	RestoreBackupContext []RestoreBackupContextParameters `json:"restoreBackupContext,omitempty" tf:"restore_backup_context,omitempty"`

	// Initial root password. Required for MS SQL Server, ignored by MySQL and PostgreSQL.
	// +kubebuilder:validation:Optional
	RootPasswordSecretRef *v1.SecretKeySelector `json:"rootPasswordSecretRef,omitempty" tf:"-"`

	// The settings to use for the database. The
	// configuration is detailed below. Required if clone is not set.
	// +kubebuilder:validation:Optional
	Settings []SettingsParameters `json:"settings,omitempty" tf:"settings,omitempty"`
}

type IPAddressObservation struct {
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address,omitempty"`

	TimeToRetire *string `json:"timeToRetire,omitempty" tf:"time_to_retire,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type IPAddressParameters struct {
}

type IPConfigurationObservation struct {
}

type IPConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	AllocatedIPRange *string `json:"allocatedIpRange,omitempty" tf:"allocated_ip_range,omitempty"`

	// +kubebuilder:validation:Optional
	AuthorizedNetworks []AuthorizedNetworksParameters `json:"authorizedNetworks,omitempty" tf:"authorized_networks,omitempty"`

	// Whether this Cloud SQL instance should be assigned
	// a public IPV4 address. At least ipv4_enabled must be enabled or a
	// private_network must be configured.
	// +kubebuilder:validation:Optional
	IPv4Enabled *bool `json:"ipv4Enabled,omitempty" tf:"ipv4_enabled,omitempty"`

	// The VPC network from which the Cloud SQL
	// instance is accessible for private IP. For example, projects/myProject/global/networks/default.
	// Specifying a network enables private IP.
	// At least ipv4_enabled must be enabled or a private_network must be configured.
	// This setting can be updated, but it cannot be removed after it is set.
	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-gcp/apis/compute/v1beta1.Network
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PrivateNetwork *string `json:"privateNetwork,omitempty" tf:"private_network,omitempty"`

	// Reference to a Network in compute to populate privateNetwork.
	// +kubebuilder:validation:Optional
	PrivateNetworkRef *v1.Reference `json:"privateNetworkRef,omitempty" tf:"-"`

	// Selector for a Network in compute to populate privateNetwork.
	// +kubebuilder:validation:Optional
	PrivateNetworkSelector *v1.Selector `json:"privateNetworkSelector,omitempty" tf:"-"`

	// Whether SSL connections over IP are enforced or not.
	// +kubebuilder:validation:Optional
	RequireSSL *bool `json:"requireSsl,omitempty" tf:"require_ssl,omitempty"`
}

type InsightsConfigObservation struct {
}

type InsightsConfigParameters struct {

	// True if Query Insights feature is enabled.
	// +kubebuilder:validation:Optional
	QueryInsightsEnabled *bool `json:"queryInsightsEnabled,omitempty" tf:"query_insights_enabled,omitempty"`

	// Maximum query length stored in bytes. Between 256 and 4500. Default to 1024.
	// +kubebuilder:validation:Optional
	QueryStringLength *float64 `json:"queryStringLength,omitempty" tf:"query_string_length,omitempty"`

	// True if Query Insights will record application tags from query when enabled.
	// +kubebuilder:validation:Optional
	RecordApplicationTags *bool `json:"recordApplicationTags,omitempty" tf:"record_application_tags,omitempty"`

	// True if Query Insights will record client address when enabled.
	// +kubebuilder:validation:Optional
	RecordClientAddress *bool `json:"recordClientAddress,omitempty" tf:"record_client_address,omitempty"`
}

type LocationPreferenceObservation struct {
}

type LocationPreferenceParameters struct {

	// A GAE application whose zone to remain
	// in. Must be in the same region as this instance.
	// +kubebuilder:validation:Optional
	FollowGaeApplication *string `json:"followGaeApplication,omitempty" tf:"follow_gae_application,omitempty"`

	// The preferred compute engine
	// zone.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type MaintenanceWindowObservation struct {
}

type MaintenanceWindowParameters struct {

	// Day of week , starting on Monday
	// +kubebuilder:validation:Optional
	Day *float64 `json:"day,omitempty" tf:"day,omitempty"`

	// Hour of day , ignored if day not set
	// +kubebuilder:validation:Optional
	Hour *float64 `json:"hour,omitempty" tf:"hour,omitempty"`

	// Receive updates earlier  or later
	// +kubebuilder:validation:Optional
	UpdateTrack *string `json:"updateTrack,omitempty" tf:"update_track,omitempty"`
}

type ReplicaConfigurationObservation struct {
}

type ReplicaConfigurationParameters struct {

	// PEM representation of the trusted CA's x509
	// certificate.
	// +kubebuilder:validation:Optional
	CACertificate *string `json:"caCertificate,omitempty" tf:"ca_certificate,omitempty"`

	// PEM representation of the replica's x509
	// certificate.
	// +kubebuilder:validation:Optional
	ClientCertificate *string `json:"clientCertificate,omitempty" tf:"client_certificate,omitempty"`

	// PEM representation of the replica's private key. The
	// corresponding public key in encoded in the client_certificate.
	// +kubebuilder:validation:Optional
	ClientKey *string `json:"clientKey,omitempty" tf:"client_key,omitempty"`

	// The number of seconds
	// between connect retries.
	// +kubebuilder:validation:Optional
	ConnectRetryInterval *float64 `json:"connectRetryInterval,omitempty" tf:"connect_retry_interval,omitempty"`

	// Path to a SQL file in GCS from which replica
	// instances are created. Format is gs://bucket/filename.
	// +kubebuilder:validation:Optional
	DumpFilePath *string `json:"dumpFilePath,omitempty" tf:"dump_file_path,omitempty"`

	// Specifies if the replica is the failover target.
	// If the field is set to true the replica will be designated as a failover replica.
	// If the master instance fails, the replica instance will be promoted as
	// the new master instance.
	// +kubebuilder:validation:Optional
	FailoverTarget *bool `json:"failoverTarget,omitempty" tf:"failover_target,omitempty"`

	// Time in ms between replication
	// heartbeats.
	// +kubebuilder:validation:Optional
	MasterHeartbeatPeriod *float64 `json:"masterHeartbeatPeriod,omitempty" tf:"master_heartbeat_period,omitempty"`

	// Password for the replication connection.
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SSLCipher *string `json:"sslCipher,omitempty" tf:"ssl_cipher,omitempty"`

	// Username for replication connection.
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`

	// True if the master's common name
	// value is checked during the SSL handshake.
	// +kubebuilder:validation:Optional
	VerifyServerCertificate *bool `json:"verifyServerCertificate,omitempty" tf:"verify_server_certificate,omitempty"`
}

type RestoreBackupContextObservation struct {
}

type RestoreBackupContextParameters struct {

	// The ID of the backup run to restore from.
	// +kubebuilder:validation:Required
	BackupRunID *float64 `json:"backupRunId" tf:"backup_run_id,omitempty"`

	// The ID of the instance that the backup was taken from. If left empty,
	// this instance's ID will be used.
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type ServerCACertObservation struct {
	Cert *string `json:"cert,omitempty" tf:"cert,omitempty"`

	CommonName *string `json:"commonName,omitempty" tf:"common_name,omitempty"`

	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	// The RFC 3339
	// formatted date time string indicating when this whitelist expires.
	ExpirationTime *string `json:"expirationTime,omitempty" tf:"expiration_time,omitempty"`

	Sha1Fingerprint *string `json:"sha1Fingerprint,omitempty" tf:"sha1_fingerprint,omitempty"`
}

type ServerCACertParameters struct {
}

type SettingsObservation struct {
	Version *float64 `json:"version,omitempty" tf:"version,omitempty"`
}

type SettingsParameters struct {

	// This specifies when the instance should be
	// active. Can be either ALWAYS, NEVER or ON_DEMAND.
	// +kubebuilder:validation:Optional
	ActivationPolicy *string `json:"activationPolicy,omitempty" tf:"activation_policy,omitempty"`

	// +kubebuilder:validation:Optional
	ActiveDirectoryConfig []ActiveDirectoryConfigParameters `json:"activeDirectoryConfig,omitempty" tf:"active_directory_config,omitempty"`

	// The availability type of the Cloud SQL
	// instance, high availability  or single zone .' For all instances, ensure that
	// settings.backup_configuration.enabled is set to true.
	// For MySQL instances, ensure that settings.backup_configuration.binary_log_enabled is set to true.
	// For Postgres instances, ensure that settings.backup_configuration.point_in_time_recovery_enabled
	// is set to true.
	// +kubebuilder:validation:Optional
	AvailabilityType *string `json:"availabilityType,omitempty" tf:"availability_type,omitempty"`

	// +kubebuilder:validation:Optional
	BackupConfiguration []BackupConfigurationParameters `json:"backupConfiguration,omitempty" tf:"backup_configuration,omitempty"`

	// The name of server instance collation.
	// +kubebuilder:validation:Optional
	Collation *string `json:"collation,omitempty" tf:"collation,omitempty"`

	// +kubebuilder:validation:Optional
	DatabaseFlags []DatabaseFlagsParameters `json:"databaseFlags,omitempty" tf:"database_flags,omitempty"`

	// Enables auto-resizing of the storage size. Set to false if you want to set disk_size.
	// +kubebuilder:validation:Optional
	DiskAutoresize *bool `json:"diskAutoresize,omitempty" tf:"disk_autoresize,omitempty"`

	// +kubebuilder:validation:Optional
	DiskAutoresizeLimit *float64 `json:"diskAutoresizeLimit,omitempty" tf:"disk_autoresize_limit,omitempty"`

	// The size of data disk, in GB. Size of a running instance cannot be reduced but can be increased. If you want to set this field, set disk_autoresize to false.
	// +kubebuilder:validation:Optional
	DiskSize *float64 `json:"diskSize,omitempty" tf:"disk_size,omitempty"`

	// The type of data disk: PD_SSD or PD_HDD.
	// +kubebuilder:validation:Optional
	DiskType *string `json:"diskType,omitempty" tf:"disk_type,omitempty"`

	// +kubebuilder:validation:Optional
	IPConfiguration []IPConfigurationParameters `json:"ipConfiguration,omitempty" tf:"ip_configuration,omitempty"`

	// +kubebuilder:validation:Optional
	InsightsConfig []InsightsConfigParameters `json:"insightsConfig,omitempty" tf:"insights_config,omitempty"`

	// +kubebuilder:validation:Optional
	LocationPreference []LocationPreferenceParameters `json:"locationPreference,omitempty" tf:"location_preference,omitempty"`

	// +kubebuilder:validation:Optional
	MaintenanceWindow []MaintenanceWindowParameters `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`

	// Pricing plan for this instance, can only be PER_USE.
	// +kubebuilder:validation:Optional
	PricingPlan *string `json:"pricingPlan,omitempty" tf:"pricing_plan,omitempty"`

	// The machine type to use. See tiers
	// for more details and supported versions. Postgres supports only shared-core machine types,
	// and custom machine types such as db-custom-2-13312. See the Custom Machine Type Documentation to learn about specifying custom machine types.
	// +kubebuilder:validation:Required
	Tier *string `json:"tier" tf:"tier,omitempty"`

	// A set of key/value user label pairs to assign to the instance.
	// +kubebuilder:validation:Optional
	UserLabels map[string]*string `json:"userLabels,omitempty" tf:"user_labels,omitempty"`
}

// DatabaseInstanceSpec defines the desired state of DatabaseInstance
type DatabaseInstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatabaseInstanceParameters `json:"forProvider"`
}

// DatabaseInstanceStatus defines the observed state of DatabaseInstance.
type DatabaseInstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatabaseInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DatabaseInstance is the Schema for the DatabaseInstances API. Creates a new SQL database instance in Google Cloud SQL.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type DatabaseInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseInstanceSpec   `json:"spec"`
	Status            DatabaseInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatabaseInstanceList contains a list of DatabaseInstances
type DatabaseInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatabaseInstance `json:"items"`
}

// Repository type metadata.
var (
	DatabaseInstance_Kind             = "DatabaseInstance"
	DatabaseInstance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DatabaseInstance_Kind}.String()
	DatabaseInstance_KindAPIVersion   = DatabaseInstance_Kind + "." + CRDGroupVersion.String()
	DatabaseInstance_GroupVersionKind = CRDGroupVersion.WithKind(DatabaseInstance_Kind)
)

func init() {
	SchemeBuilder.Register(&DatabaseInstance{}, &DatabaseInstanceList{})
}
