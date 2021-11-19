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

type RegionInstanceGroupManagerAutoHealingPoliciesObservation struct {
}

type RegionInstanceGroupManagerAutoHealingPoliciesParameters struct {

	// The health check resource that signals autohealing.
	// +kubebuilder:validation:Required
	HealthCheck *string `json:"healthCheck" tf:"health_check,omitempty"`

	// The number of seconds that the managed instance group waits before it applies autohealing policies to new instances or recently recreated instances. Between 0 and 3600.
	// +kubebuilder:validation:Required
	InitialDelaySec *int64 `json:"initialDelaySec" tf:"initial_delay_sec,omitempty"`
}

type RegionInstanceGroupManagerNamedPortObservation struct {
}

type RegionInstanceGroupManagerNamedPortParameters struct {

	// The name of the port.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The port number.
	// +kubebuilder:validation:Required
	Port *int64 `json:"port" tf:"port,omitempty"`
}

type RegionInstanceGroupManagerObservation struct {
	Fingerprint *string `json:"fingerprint,omitempty" tf:"fingerprint,omitempty"`

	InstanceGroup *string `json:"instanceGroup,omitempty" tf:"instance_group,omitempty"`

	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	Status []RegionInstanceGroupManagerStatusObservation `json:"status,omitempty" tf:"status,omitempty"`
}

type RegionInstanceGroupManagerParameters struct {

	// The autohealing policies for this managed instance group. You can specify only one value.
	// +kubebuilder:validation:Optional
	AutoHealingPolicies []RegionInstanceGroupManagerAutoHealingPoliciesParameters `json:"autoHealingPolicies,omitempty" tf:"auto_healing_policies,omitempty"`

	// The base instance name to use for instances in this group. The value must be a valid RFC1035 name. Supported characters are lowercase letters, numbers, and hyphens (-). Instances are named by appending a hyphen and a random four-character string to the base instance name.
	// +kubebuilder:validation:Required
	BaseInstanceName *string `json:"baseInstanceName" tf:"base_instance_name,omitempty"`

	// An optional textual description of the instance group manager.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The shape to which the group converges either proactively or on resize events (depending on the value set in updatePolicy.instanceRedistributionType).
	// +kubebuilder:validation:Optional
	DistributionPolicyTargetShape *string `json:"distributionPolicyTargetShape,omitempty" tf:"distribution_policy_target_shape,omitempty"`

	// The distribution policy for this managed instance group. You can specify one or more values.
	// +kubebuilder:validation:Optional
	DistributionPolicyZones []*string `json:"distributionPolicyZones,omitempty" tf:"distribution_policy_zones,omitempty"`

	// The name of the instance group manager. Must be 1-63 characters long and comply with RFC1035. Supported characters include lowercase letters, numbers, and hyphens.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The named port configuration.
	// +kubebuilder:validation:Optional
	NamedPort []RegionInstanceGroupManagerNamedPortParameters `json:"namedPort,omitempty" tf:"named_port,omitempty"`

	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region where the managed instance group resides.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Disks created on the instances that will be preserved on instance delete, update, etc. Structure is documented below. For more information see the official documentation. Proactive cross zone instance redistribution must be disabled before you can update stateful disks on existing instance group managers. This can be controlled via the update_policy.
	// +kubebuilder:validation:Optional
	StatefulDisk []RegionInstanceGroupManagerStatefulDiskParameters `json:"statefulDisk,omitempty" tf:"stateful_disk,omitempty"`

	// The full URL of all target pools to which new instances in the group are added. Updating the target pools attribute does not affect existing instances.
	// +kubebuilder:validation:Optional
	TargetPools []*string `json:"targetPools,omitempty" tf:"target_pools,omitempty"`

	// The target number of running instances for this managed instance group. This value should always be explicitly set unless this resource is attached to an autoscaler, in which case it should never be set. Defaults to 0.
	// +kubebuilder:validation:Optional
	TargetSize *int64 `json:"targetSize,omitempty" tf:"target_size,omitempty"`

	// The update policy for this managed instance group.
	// +kubebuilder:validation:Optional
	UpdatePolicy []RegionInstanceGroupManagerUpdatePolicyParameters `json:"updatePolicy,omitempty" tf:"update_policy,omitempty"`

	// Application versions managed by this instance group. Each version deals with a specific instance template, allowing canary release scenarios.
	// +kubebuilder:validation:Required
	Version []RegionInstanceGroupManagerVersionParameters `json:"version" tf:"version,omitempty"`

	// Whether to wait for all instances to be created/updated before returning. Note that if this is set to true and the operation does not succeed, Terraform will continue trying until it times out.
	// +kubebuilder:validation:Optional
	WaitForInstances *bool `json:"waitForInstances,omitempty" tf:"wait_for_instances,omitempty"`

	// When used with wait_for_instances specifies the status to wait for. When STABLE is specified this resource will wait until the instances are stable before returning. When UPDATED is set, it will wait for the version target to be reached and any per instance configs to be effective as well as all instances to be stable before returning.
	// +kubebuilder:validation:Optional
	WaitForInstancesStatus *string `json:"waitForInstancesStatus,omitempty" tf:"wait_for_instances_status,omitempty"`
}

type RegionInstanceGroupManagerStatefulDiskObservation struct {
}

type RegionInstanceGroupManagerStatefulDiskParameters struct {

	// A value that prescribes what should happen to the stateful disk when the VM instance is deleted. The available options are NEVER and ON_PERMANENT_INSTANCE_DELETION. NEVER - detach the disk when the VM is deleted, but do not delete the disk. ON_PERMANENT_INSTANCE_DELETION will delete the stateful disk when the VM is permanently deleted from the instance group. The default is NEVER.
	// +kubebuilder:validation:Optional
	DeleteRule *string `json:"deleteRule,omitempty" tf:"delete_rule,omitempty"`

	// The device name of the disk to be attached.
	// +kubebuilder:validation:Required
	DeviceName *string `json:"deviceName" tf:"device_name,omitempty"`
}

type RegionInstanceGroupManagerStatusObservation struct {
	IsStable *bool `json:"isStable,omitempty" tf:"is_stable,omitempty"`

	Stateful []StatusStatefulObservation `json:"stateful,omitempty" tf:"stateful,omitempty"`

	VersionTarget []StatusVersionTargetObservation `json:"versionTarget,omitempty" tf:"version_target,omitempty"`
}

type RegionInstanceGroupManagerStatusParameters struct {
}

type RegionInstanceGroupManagerUpdatePolicyObservation struct {
}

type RegionInstanceGroupManagerUpdatePolicyParameters struct {

	// The instance redistribution policy for regional managed instance groups. Valid values are: "PROACTIVE", "NONE". If PROACTIVE (default), the group attempts to maintain an even distribution of VM instances across zones in the region. If NONE, proactive redistribution is disabled.
	// +kubebuilder:validation:Optional
	InstanceRedistributionType *string `json:"instanceRedistributionType,omitempty" tf:"instance_redistribution_type,omitempty"`

	// The maximum number of instances that can be created above the specified targetSize during the update process. Conflicts with max_surge_percent. It has to be either 0 or at least equal to the number of zones. If fixed values are used, at least one of max_unavailable_fixed or max_surge_fixed must be greater than 0.
	// +kubebuilder:validation:Optional
	MaxSurgeFixed *int64 `json:"maxSurgeFixed,omitempty" tf:"max_surge_fixed,omitempty"`

	// The maximum number of instances(calculated as percentage) that can be created above the specified targetSize during the update process. Conflicts with max_surge_fixed. Percent value is only allowed for regional managed instance groups with size at least 10.
	// +kubebuilder:validation:Optional
	MaxSurgePercent *int64 `json:"maxSurgePercent,omitempty" tf:"max_surge_percent,omitempty"`

	// The maximum number of instances that can be unavailable during the update process. Conflicts with max_unavailable_percent. It has to be either 0 or at least equal to the number of zones. If fixed values are used, at least one of max_unavailable_fixed or max_surge_fixed must be greater than 0.
	// +kubebuilder:validation:Optional
	MaxUnavailableFixed *int64 `json:"maxUnavailableFixed,omitempty" tf:"max_unavailable_fixed,omitempty"`

	// The maximum number of instances(calculated as percentage) that can be unavailable during the update process. Conflicts with max_unavailable_fixed. Percent value is only allowed for regional managed instance groups with size at least 10.
	// +kubebuilder:validation:Optional
	MaxUnavailablePercent *int64 `json:"maxUnavailablePercent,omitempty" tf:"max_unavailable_percent,omitempty"`

	// Minimal action to be taken on an instance. You can specify either RESTART to restart existing instances or REPLACE to delete and create new instances from the target template. If you specify a RESTART, the Updater will attempt to perform that action only. However, if the Updater determines that the minimal action you specify is not enough to perform the update, it might perform a more disruptive action.
	// +kubebuilder:validation:Required
	MinimalAction *string `json:"minimalAction" tf:"minimal_action,omitempty"`

	// The instance replacement method for regional managed instance groups. Valid values are: "RECREATE", "SUBSTITUTE". If SUBSTITUTE (default), the group replaces VM instances with new instances that have randomly generated names. If RECREATE, instance names are preserved.  You must also set max_unavailable_fixed or max_unavailable_percent to be greater than 0.
	// +kubebuilder:validation:Optional
	ReplacementMethod *string `json:"replacementMethod,omitempty" tf:"replacement_method,omitempty"`

	// The type of update process. You can specify either PROACTIVE so that the instance group manager proactively executes actions in order to bring instances to their target versions or OPPORTUNISTIC so that no action is proactively executed but the update will be performed as part of other actions (for example, resizes or recreateInstances calls).
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type RegionInstanceGroupManagerVersionObservation struct {
}

type RegionInstanceGroupManagerVersionParameters struct {

	// The full URL to an instance template from which all new instances of this version will be created.
	// +kubebuilder:validation:Required
	InstanceTemplate *string `json:"instanceTemplate" tf:"instance_template,omitempty"`

	// Version name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The number of instances calculated as a fixed number or a percentage depending on the settings.
	// +kubebuilder:validation:Optional
	TargetSize []VersionTargetSizeParameters `json:"targetSize,omitempty" tf:"target_size,omitempty"`
}

type StatefulPerInstanceConfigsObservation struct {
	AllEffective *bool `json:"allEffective,omitempty" tf:"all_effective,omitempty"`
}

type StatefulPerInstanceConfigsParameters struct {
}

type StatusStatefulObservation struct {
	HasStatefulConfig *bool `json:"hasStatefulConfig,omitempty" tf:"has_stateful_config,omitempty"`

	PerInstanceConfigs []StatefulPerInstanceConfigsObservation `json:"perInstanceConfigs,omitempty" tf:"per_instance_configs,omitempty"`
}

type StatusStatefulParameters struct {
}

type StatusVersionTargetObservation struct {
	IsReached *bool `json:"isReached,omitempty" tf:"is_reached,omitempty"`
}

type StatusVersionTargetParameters struct {
}

type VersionTargetSizeObservation struct {
}

type VersionTargetSizeParameters struct {

	// The number of instances which are managed for this version. Conflicts with percent.
	// +kubebuilder:validation:Optional
	Fixed *int64 `json:"fixed,omitempty" tf:"fixed,omitempty"`

	// The number of instances (calculated as percentage) which are managed for this version. Conflicts with fixed. Note that when using percent, rounding will be in favor of explicitly set target_size values; a managed instance group with 2 instances and 2 versions, one of which has a target_size.percent of 60 will create 2 instances of that version.
	// +kubebuilder:validation:Optional
	Percent *int64 `json:"percent,omitempty" tf:"percent,omitempty"`
}

// RegionInstanceGroupManagerSpec defines the desired state of RegionInstanceGroupManager
type RegionInstanceGroupManagerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RegionInstanceGroupManagerParameters `json:"forProvider"`
}

// RegionInstanceGroupManagerStatus defines the observed state of RegionInstanceGroupManager.
type RegionInstanceGroupManagerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RegionInstanceGroupManagerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RegionInstanceGroupManager is the Schema for the RegionInstanceGroupManagers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type RegionInstanceGroupManager struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RegionInstanceGroupManagerSpec   `json:"spec"`
	Status            RegionInstanceGroupManagerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RegionInstanceGroupManagerList contains a list of RegionInstanceGroupManagers
type RegionInstanceGroupManagerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RegionInstanceGroupManager `json:"items"`
}

// Repository type metadata.
var (
	RegionInstanceGroupManager_Kind             = "RegionInstanceGroupManager"
	RegionInstanceGroupManager_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RegionInstanceGroupManager_Kind}.String()
	RegionInstanceGroupManager_KindAPIVersion   = RegionInstanceGroupManager_Kind + "." + CRDGroupVersion.String()
	RegionInstanceGroupManager_GroupVersionKind = CRDGroupVersion.WithKind(RegionInstanceGroupManager_Kind)
)

func init() {
	SchemeBuilder.Register(&RegionInstanceGroupManager{}, &RegionInstanceGroupManagerList{})
}