// Copyright 2025 Google LLC
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

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ClusterSpec struct {
	/* An optional description of this cluster. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* The initial number of nodes for the cluster. */
	// +optional
	InitialNodeCount *int32 `json:"initialNodeCount,omitempty"`

	/* The node configuration for the cluster. */
	// +optional
	NodeConfig *NodeConfig `json:"nodeConfig,omitempty"`

	/* The master auth settings for controlling access to the cluster endpoint. */
	// +optional
	MasterAuth *MasterAuth `json:"masterAuth,omitempty"`

	/* The logging service the cluster should use to write logs. */
	// +optional
	LoggingService *string `json:"loggingService,omitempty"`

	/* The monitoring service the cluster should use to write metrics. */
	// +optional
	MonitoringService *string `json:"monitoringService,omitempty"`

	/* The name of the Google Compute Engine network to which the cluster is connected. */
	// +optional
	NetworkRef *v1alpha1.ResourceRef `json:"networkRef,omitempty"`

	/* The name of the Google Compute Engine subnetwork to which the cluster is connected. */
	// +optional
	SubnetworkRef *v1alpha1.ResourceRef `json:"subnetworkRef,omitempty"`

	/* The location for the cluster. */
	Location string `json:"location"`

	/* The list of node pools associated with this cluster. */
	// +optional
	NodePools []NodePool `json:"nodePools,omitempty"`

	/* Configuration for the legacy ABAC authorization mode. */
	// +optional
	LegacyAbac *LegacyAbac `json:"legacyAbac,omitempty"`

	/* Configuration for NetworkPolicy. */
	// +optional
	NetworkPolicy *NetworkPolicy `json:"networkPolicy,omitempty"`

	/* Configuration for cluster add-ons. */
	// +optional
	AddonsConfig *AddonsConfig `json:"addonsConfig,omitempty"`

	/* The private cluster config. */
	// +optional
	PrivateClusterConfig *PrivateClusterConfig `json:"privateClusterConfig,omitempty"`

	/* The master authorized networks configuration. */
	// +optional
	MasterAuthorizedNetworksConfig *MasterAuthorizedNetworksConfig `json:"masterAuthorizedNetworksConfig,omitempty"`

	/* The release channel configuration. */
	// +optional
	ReleaseChannel *ReleaseChannel `json:"releaseChannel,omitempty"`

	/* The workload identity configuration. */
	// +optional
	WorkloadIdentityConfig *WorkloadIdentityConfig `json:"workloadIdentityConfig,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type ClusterStatus struct {
	/* Conditions represent the latest available observations of the
	   ContainerCluster's current state. */
	// +optional
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	/* The IP address of this cluster's master endpoint. */
	// +optional
	Endpoint *string `json:"endpoint,omitempty"`

	/* The current software version of the master endpoint. */
	// +optional
	CurrentMasterVersion *string `json:"currentMasterVersion,omitempty"`

	/* The current version of the node software components. */
	// +optional
	CurrentNodeVersion *string `json:"currentNodeVersion,omitempty"`

	/* The time the cluster was created, in RFC3339 text format. */
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	/* The current status of this cluster. */
	// +optional
	Status *string `json:"status,omitempty"`

	/* The number of nodes currently in the cluster. */
	// +optional
	CurrentNodeCount *int32 `json:"currentNodeCount,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* The self link of the cluster. */
	// +optional
	SelfLink *string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcontainercluster;gcpcontainerclusters
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].message",type="string"
// +kubebuilder:printcolumn:name="Endpoint",JSONPath=".status.endpoint",type="string"
// +kubebuilder:printcolumn:name="MasterVersion",JSONPath=".status.currentMasterVersion",type="string"
// +kubebuilder:printcolumn:name="Location",JSONPath=".spec.location",type="string"

// ContainerCluster is the Schema for the container API
// +k8s:openapi-gen=true
type ContainerCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterSpec   `json:"spec,omitempty"`
	Status ClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ContainerClusterList contains a list of ContainerCluster
type ContainerClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContainerCluster `json:"items"`
}

type NodeConfig struct {
	/* The name of a Google Compute Engine machine type. */
	// +optional
	MachineType *string `json:"machineType,omitempty"`

	/* Size of the disk attached to each node, specified in GB. */
	// +optional
	DiskSizeGb *int32 `json:"diskSizeGb,omitempty"`

	/* The set of Google API scopes to be made available on all of the node VMs. */
	// +optional
	OAuthScopes []string `json:"oauthScopes,omitempty"`

	/* The Google Cloud Platform Service Account to be used by the node VMs. */
	// +optional
	ServiceAccountRef *v1alpha1.ResourceRef `json:"serviceAccountRef,omitempty"`

	/* The metadata key/value pairs assigned to instances. */
	// +optional
	Metadata map[string]string `json:"metadata,omitempty"`

	/* The image type to use for this node. */
	// +optional
	ImageType *string `json:"imageType,omitempty"`

	/* The map of Kubernetes labels (key/value pairs) to be applied to each node. */
	// +optional
	Labels map[string]string `json:"labels,omitempty"`

	/* The number of local SSD disks to attach to the node. */
	// +optional
	LocalSsdCount *int32 `json:"localSsdCount,omitempty"`

	/* The list of instance tags applied to all nodes. */
	// +optional
	Tags []string `json:"tags,omitempty"`

	/* Whether the nodes are created as preemptible VM instances. */
	// +optional
	Preemptible *bool `json:"preemptible,omitempty"`

	/* A list of hardware accelerators to be attached to each node. */
	// +optional
	GuestAccelerator []AcceleratorConfig `json:"guestAccelerator,omitempty"`

	/* Type of the disk attached to each node. */
	// +optional
	DiskType *string `json:"diskType,omitempty"`

	/* The name of the min CPU platform to be used by this node pool. */
	// +optional
	MinCpuPlatform *string `json:"minCpuPlatform,omitempty"`

	/* The workload metadata configuration for this node pool. */
	// +optional
	WorkloadMetadataConfig *WorkloadMetadataConfig `json:"workloadMetadataConfig,omitempty"`

	/* List of kubernetes taints to be applied to each node. */
	// +optional
	Taint []NodeTaint `json:"taint,omitempty"`

	/* The Customer Managed Encryption Key used to encrypt the boot disk attached to each node in the node pool. */
	// +optional
	BootDiskKmsKeyRef *v1alpha1.ResourceRef `json:"bootDiskKmsKeyRef,omitempty"`
}

type NodePool struct {
	/* The name of the node pool. */
	Name string `json:"name"`

	/* The node configuration for the pool. */
	// +optional
	Config *NodeConfig `json:"config,omitempty"`

	/* The initial number of nodes for the pool. */
	// +optional
	InitialNodeCount *int32 `json:"initialNodeCount,omitempty"`

	/* The version of the Kubernetes of this node. */
	// +optional
	Version *string `json:"version,omitempty"`

	/* Autoscaler configuration for this node pool. */
	// +optional
	Autoscaling *NodePoolAutoscaling `json:"autoscaling,omitempty"`

	/* Node management configuration, consisting of options for handling node upgrades. */
	// +optional
	Management *NodeManagement `json:"management,omitempty"`

	/* The constraint on the maximum number of pods that can be run simultaneously on a node in the node pool. */
	// +optional
	MaxPodsPerNode *int32 `json:"maxPodsPerNode,omitempty"`
}

type MasterAuth struct {
	/* The username to use for HTTP basic authentication to the master endpoint. */
	// +optional
	Username *string `json:"username,omitempty"`

	/* The password to use for HTTP basic authentication to the master endpoint. */
	// +optional
	Password *v1alpha1.SecretKeyRef `json:"password,omitempty"`

	/* Configuration for client certificate authentication on the cluster. */
	// +optional
	ClientCertificateConfig *ClientCertificateConfig `json:"clientCertificateConfig,omitempty"`
}

type ClientCertificateConfig struct {
	/* Whether client certificate authorization is enabled for this cluster. */
	IssueClientCertificate bool `json:"issueClientCertificate"`
}

type AddonsConfig struct {
	/* Configuration for the HTTP (L7) load balancing controller addon. */
	// +optional
	HttpLoadBalancing *HttpLoadBalancing `json:"httpLoadBalancing,omitempty"`

	/* Configuration for the horizontal pod autoscaling feature, which increases or decreases the number of replica pods a replication controller has based on the resource usage of the existing pods. */
	// +optional
	HorizontalPodAutoscaling *HorizontalPodAutoscaling `json:"horizontalPodAutoscaling,omitempty"`

	/* Configuration for the Kubernetes Dashboard. */
	// +optional
	KubernetesDashboard *KubernetesDashboard `json:"kubernetesDashboard,omitempty"`

	/* Configuration for NetworkPolicy. */
	// +optional
	NetworkPolicyConfig *NetworkPolicyConfig `json:"networkPolicyConfig,omitempty"`

	/* Configuration for the Cloud Run addon. */
	// +optional
	CloudRunConfig *CloudRunConfig `json:"cloudRunConfig,omitempty"`

	/* Configuration for NodeLocal DNSCache. */
	// +optional
	DnsCacheConfig *DnsCacheConfig `json:"dnsCacheConfig,omitempty"`

	/* Configuration for the ConfigConnector add-on. */
	// +optional
	ConfigConnectorConfig *ConfigConnectorConfig `json:"configConnectorConfig,omitempty"`

	/* Configuration for the GCE PD CSI driver addon. */
	// +optional
	GcePersistentDiskCsiDriverConfig *GcePersistentDiskCsiDriverConfig `json:"gcePersistentDiskCsiDriverConfig,omitempty"`
}

type HttpLoadBalancing struct {
	/* Whether the HTTP Load Balancing controller is enabled in the cluster. */
	Disabled bool `json:"disabled"`
}

type HorizontalPodAutoscaling struct {
	/* Whether the Horizontal Pod Autoscaling feature is enabled in the cluster. */
	Disabled bool `json:"disabled"`
}

type KubernetesDashboard struct {
	/* Whether the Kubernetes Dashboard is enabled for this cluster. */
	Disabled bool `json:"disabled"`
}

type NetworkPolicyConfig struct {
	/* Whether NetworkPolicy is enabled for this cluster. */
	Disabled bool `json:"disabled"`
}

type CloudRunConfig struct {
	/* Whether Cloud Run is enabled for this cluster. */
	Disabled bool `json:"disabled"`
}

type DnsCacheConfig struct {
	/* Whether NodeLocal DNSCache is enabled for this cluster. */
	Enabled bool `json:"enabled"`
}

type ConfigConnectorConfig struct {
	/* Whether ConfigConnector is enabled for this cluster. */
	Enabled bool `json:"enabled"`
}

type GcePersistentDiskCsiDriverConfig struct {
	/* Whether the GCE PD CSI driver is enabled for this cluster. */
	Enabled bool `json:"enabled"`
}

type PrivateClusterConfig struct {
	/* Whether the master's internal IP address is used as the cluster endpoint. */
	// +optional
	EnablePrivateEndpoint *bool `json:"enablePrivateEndpoint,omitempty"`

	/* The IP range in CIDR notation to use for the master network. */
	// +optional
	MasterIpv4CidrBlock *string `json:"masterIpv4CidrBlock,omitempty"`

	/* The peering name and project of the VPC that is peered with the cluster's master network. */
	// +optional
	PeeringRef *v1alpha1.ResourceRef `json:"peeringRef,omitempty"`
}

type MasterAuthorizedNetworksConfig struct {
	/* Whether or not master authorized networks is enabled. */
	Enabled bool `json:"enabled"`

	/* CIDR blocks to be configured as authorized networks. */
	// +optional
	CidrBlocks []MasterAuthorizedNetworksConfigCidrBlock `json:"cidrBlocks,omitempty"`
}

type MasterAuthorizedNetworksConfigCidrBlock struct {
	/* Field for users to identify CIDR blocks. */
	// +optional
	DisplayName *string `json:"displayName,omitempty"`

	/* The CIDR block. */
	CidrBlock string `json:"cidrBlock"`
}

type LegacyAbac struct {
	/* Whether the ABAC authorizer is enabled for this cluster. */
	Enabled bool `json:"enabled"`
}

type NetworkPolicy struct {
	/* The selected network policy provider. */
	Provider string `json:"provider"`

	/* Whether network policy is enabled on the cluster. */
	Enabled bool `json:"enabled"`
}

type ReleaseChannel struct {
	/* The selected release channel. */
	Channel string `json:"channel"`
}

type WorkloadIdentityConfig struct {
	/* The workload pool to attach all Kubernetes service accounts to. */
	WorkloadPool string `json:"workloadPool"`
}

type AcceleratorConfig struct {
	/* The number of the accelerator cards exposed to an instance. */
	AcceleratorCount int64 `json:"acceleratorCount"`

	/* The accelerator type resource name. */
	AcceleratorType string `json:"acceleratorType"`
}

type WorkloadMetadataConfig struct {
	/* Mode is the configuration for how to expose metadata to workloads running on the node. */
	Mode string `json:"mode"`
}

type NodeTaint struct {
	/* The taint key. */
	Key string `json:"key"`

	/* The taint value. */
	Value string `json:"value"`

	/* The taint effect. */
	Effect string `json:"effect"`
}

type NodePoolAutoscaling struct {
	/* Is autoscaling enabled for this node pool. */
	// +optional
	Enabled *bool `json:"enabled,omitempty"`

	/* Minimum number of nodes in the NodePool. */
	// +optional
	MinNodeCount *int32 `json:"minNodeCount,omitempty"`

	/* Maximum number of nodes in the NodePool. */
	// +optional
	MaxNodeCount *int32 `json:"maxNodeCount,omitempty"`
}

type NodeManagement struct {
	/* Whether the nodes will be automatically upgraded. */
	// +optional
	AutoUpgrade *bool `json:"autoUpgrade,omitempty"`

	/* Whether the nodes will be automatically repaired. */
	// +optional
	AutoRepair *bool `json:"autoRepair,omitempty"`
}

func init() {
	SchemeBuilder.Register(&ContainerCluster{}, &ContainerClusterList{})
}