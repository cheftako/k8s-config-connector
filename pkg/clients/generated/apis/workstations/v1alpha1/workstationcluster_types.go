// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type WorkstationclusterAllowedProjects struct {
	/* The `projectID` field of a project, when not managed by Config Connector. */
	// +optional
	External *string `json:"external,omitempty"`

	/* The kind of the Project resource; optional but must be `Project` if provided. */
	// +optional
	Kind *string `json:"kind,omitempty"`

	/* The `name` field of a `Project` resource. */
	// +optional
	Name *string `json:"name,omitempty"`

	/* The `namespace` field of a `Project` resource. */
	// +optional
	Namespace *string `json:"namespace,omitempty"`
}

type WorkstationclusterAnnotations struct {
	/* Key for the annotation. */
	// +optional
	Key *string `json:"key,omitempty"`

	/* Value for the annotation. */
	// +optional
	Value *string `json:"value,omitempty"`
}

type WorkstationclusterLabels struct {
	/* Key for the annotation. */
	// +optional
	Key *string `json:"key,omitempty"`

	/* Value for the annotation. */
	// +optional
	Value *string `json:"value,omitempty"`
}

type WorkstationclusterPrivateClusterConfig struct {
	/* Optional. Additional projects that are allowed to attach to the workstation cluster's service attachment. By default, the workstation cluster's project and the VPC host project (if different) are allowed. */
	// +optional
	AllowedProjects []WorkstationclusterAllowedProjects `json:"allowedProjects,omitempty"`

	/* Immutable. Whether Workstations endpoint is private. */
	// +optional
	EnablePrivateEndpoint *bool `json:"enablePrivateEndpoint,omitempty"`
}

type WorkstationClusterSpec struct {
	/* Optional. Client-specified annotations. */
	// +optional
	Annotations []WorkstationclusterAnnotations `json:"annotations,omitempty"`

	/* Optional. Human-readable name for this workstation cluster. */
	// +optional
	DisplayName *string `json:"displayName,omitempty"`

	/* Optional. [Labels](https://cloud.google.com/workstations/docs/label-resources) that are applied to the workstation cluster and that are also propagated to the underlying Compute Engine resources. */
	// +optional
	Labels []WorkstationclusterLabels `json:"labels,omitempty"`

	/* The location of the cluster. */
	// +optional
	Location *string `json:"location,omitempty"`

	/* Immutable. Reference to the Compute Engine network in which instances associated with this workstation cluster will be created. */
	NetworkRef v1alpha1.ResourceRef `json:"networkRef"`

	/* Optional. Configuration for private workstation cluster. */
	// +optional
	PrivateClusterConfig *WorkstationclusterPrivateClusterConfig `json:"privateClusterConfig,omitempty"`

	/* Immutable. The Project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. The WorkstationCluster name. If not given, the metadata.name will be used. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Immutable. Reference to the Compute Engine subnetwork in which instances associated with this workstation cluster will be created. Must be part of the subnetwork specified for this workstation cluster. */
	SubnetworkRef v1alpha1.ResourceRef `json:"subnetworkRef"`
}

type WorkstationclusterGcpConditionsStatus struct {
	/* The status code, which should be an enum value of [google.rpc.Code][google.rpc.Code]. */
	// +optional
	Code *int64 `json:"code,omitempty"`

	/* A developer-facing error message, which should be in English. Any user-facing error message should be localized and sent in the [google.rpc.Status.details][google.rpc.Status.details] field, or localized by the client. */
	// +optional
	Message *string `json:"message,omitempty"`
}

type WorkstationclusterObservedStateStatus struct {
	/* Output only. Hostname for the workstation cluster. This field will be populated only when private endpoint is enabled. To access workstations in the workstation cluster, create a new DNS zone mapping this domain name to an internal IP address and a forwarding rule mapping that address to the service attachment. */
	// +optional
	ClusterHostname *string `json:"clusterHostname,omitempty"`

	/* Output only. The private IP address of the control plane for this workstation cluster. Workstation VMs need access to this IP address to work with the service, so make sure that your firewall rules allow egress from the workstation VMs to this address. */
	// +optional
	ControlPlaneIP *string `json:"controlPlaneIP,omitempty"`

	/* Output only. Time when this workstation cluster was created. */
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	/* Output only. Whether this workstation cluster is in degraded mode, in which case it may require user action to restore full functionality. Details can be found in [conditions][google.cloud.workstations.v1.WorkstationCluster.conditions]. */
	// +optional
	Degraded *bool `json:"degraded,omitempty"`

	/* Output only. Time when this workstation cluster was soft-deleted. */
	// +optional
	DeleteTime *string `json:"deleteTime,omitempty"`

	/* Optional. Checksum computed by the server. May be sent on update and delete requests to make sure that the client has an up-to-date value before proceeding. */
	// +optional
	Etag *string `json:"etag,omitempty"`

	/* Output only. Status conditions describing the workstation cluster's current state. */
	// +optional
	GcpConditions []WorkstationclusterGcpConditionsStatus `json:"gcpConditions,omitempty"`

	/* Output only. Indicates whether this workstation cluster is currently being updated to match its intended state. */
	// +optional
	Reconciling *bool `json:"reconciling,omitempty"`

	/* Output only. Service attachment URI for the workstation cluster. The service attachment is created when private endpoint is enabled. To access workstations in the workstation cluster, configure access to the managed service using [Private Service Connect](https://cloud.google.com/vpc/docs/configure-private-service-connect-services). */
	// +optional
	ServiceAttachmentUri *string `json:"serviceAttachmentUri,omitempty"`

	/* Output only. A system-assigned unique identifier for this workstation cluster. */
	// +optional
	Uid *string `json:"uid,omitempty"`

	/* Output only. Time when this workstation cluster was most recently updated. */
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`
}

type WorkstationClusterStatus struct {
	/* Conditions represent the latest available observations of the
	   WorkstationCluster's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* A unique specifier for the WorkstationCluster resource in GCP. */
	// +optional
	ExternalRef *string `json:"externalRef,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* ObservedState is the state of the resource as most recently observed in GCP. */
	// +optional
	ObservedState *WorkstationclusterObservedStateStatus `json:"observedState,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// WorkstationCluster is the Schema for the workstations API
// +k8s:openapi-gen=true
type WorkstationCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WorkstationClusterSpec   `json:"spec,omitempty"`
	Status WorkstationClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WorkstationClusterList contains a list of WorkstationCluster
type WorkstationClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorkstationCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&WorkstationCluster{}, &WorkstationClusterList{})
}
