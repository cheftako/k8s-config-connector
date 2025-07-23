/*
Copyright 2020 Google LLC

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ContainerCluster is the Schema for the containerclusters API
// +k8s:openapi-gen=true
type ContainerCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ContainerClusterSpec   `json:"spec,omitempty"`
	Status ContainerClusterStatus `json:"status,omitempty"`
}

// ContainerClusterSpec defines the desired state of ContainerCluster
type ContainerClusterSpec struct {
	// The name of the cluster, unique within the project and location.
	Name string `json:"name"`
	// The location for the cluster.
	Location string `json:"location"`
	// The initial number of nodes for the cluster.
	InitialNodeCount int64 `json:"initialNodeCount,omitempty"`
	// The configuration for the master node.
	MasterAuth *MasterAuth `json:"masterAuth,omitempty"`
	// The configuration for the node pools in the cluster.
	NodePools []*NodePool `json:"nodePools,omitempty"`
	// The configuration for addons.
	AddonsConfig *AddonsConfig `json:"addonsConfig,omitempty"`
	// The configuration for network policy.
	NetworkPolicy *NetworkPolicy `json:"networkPolicy,omitempty"`
	// The IP allocation policy for the cluster.
	IPAllocationPolicy *IPAllocationPolicy `json:"ipAllocationPolicy,omitempty"`
	// The configuration for master authorized networks.
	MasterAuthorizedNetworksConfig *MasterAuthorizedNetworksConfig `json:"masterAuthorizedNetworksConfig,omitempty"`
	// The configuration for private clusters.
	PrivateClusterConfig *PrivateClusterConfig `json:"privateClusterConfig,omitempty"`
	// The configuration for binary authorization.
	BinaryAuthorization *BinaryAuthorization `json:"binaryAuthorization,omitempty"`
	// The configuration for shielded nodes.
	ShieldedNodes *ShieldedNodes `json:"shieldedNodes,omitempty"`
	// The release channel for the cluster.
	ReleaseChannel *ReleaseChannel `json:"releaseChannel,omitempty"`
	// The workload identity configuration for the cluster.
	WorkloadIdentityConfig *WorkloadIdentityConfig `json:"workloadIdentityConfig,omitempty"`
	// The notification configuration for the cluster.
	NotificationConfig *NotificationConfig `json:"notificationConfig,omitempty"`
	// The confidential nodes configuration for the cluster.
	ConfidentialNodes *ConfidentialNodes `json:"confidentialNodes,omitempty"`
	// The authenticator groups configuration for the cluster.
	AuthenticatorGroupsConfig *AuthenticatorGroupsConfig `json:"authenticatorGroupsConfig,omitempty"`
	// The logging service for the cluster.
	LoggingService string `json:"loggingService,omitempty"`
	// The monitoring service for the cluster.
	MonitoringService string `json:"monitoringService,omitempty"`
	// The networking configuration for the cluster.
	Network string `json:"network,omitempty"`
	// The subnetwork configuration for the cluster.
	Subnetwork string `json:"subnetwork,omitempty"`
	// The default maximum number of pods per node in the cluster.
	DefaultMaxPodsPerNode int64 `json:"defaultMaxPodsPerNode,omitempty"`
	// The resource usage export configuration for the cluster.
	ResourceUsageExportConfig *ResourceUsageExportConfig `json:"resourceUsageExportConfig,omitempty"`
	// The vertical pod autoscaling configuration for the cluster.
	VerticalPodAutoscaling *VerticalPodAutoscaling `json:"verticalPodAutoscaling,omitempty"`
	// The resource labels for the cluster.
	ResourceLabels map[string]string `json:"resourceLabels,omitempty"`
}

// ContainerClusterStatus defines the observed state of ContainerCluster
type ContainerClusterStatus struct {
	// The status of the cluster.
	Status string `json:"status,omitempty"`
	// The endpoint for the cluster.
	Endpoint string `json:"endpoint,omitempty"`
	// The current master version of the cluster.
	CurrentMasterVersion string `json:"currentMasterVersion,omitempty"`
	// The current node version of the cluster.
	CurrentNodeVersion string `json:"currentNodeVersion,omitempty"`
	// The create time of the cluster.
	CreateTime string `json:"createTime,omitempty"`
	// The expire time of the cluster.
	ExpireTime string `json:"expireTime,omitempty"`
	// The current node count of the cluster.
	CurrentNodeCount int64 `json:"currentNodeCount,omitempty"`
	// The services IPv4 CIDR block for the cluster.
	ServicesIPV4CIDR string `json:"servicesIpv4Cidr,omitempty"`
	// The TPU IPv4 CIDR block for the cluster.
	TpuIPV4CIDRBlock string `json:"tpuIpv4CidrBlock,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ContainerClusterList contains a list of ContainerCluster
type ContainerClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContainerCluster `json:"items"`
}

// MasterAuth defines the master authentication configuration.
type MasterAuth struct {
	// The username for basic authentication.
	Username string `json:"username,omitempty"`
	// The password for basic authentication.
	Password string `json:"password,omitempty"`
	// The client certificate configuration.
	ClientCertificateConfig *ClientCertificateConfig `json:"clientCertificateConfig,omitempty"`
}

// ClientCertificateConfig defines the client certificate configuration.
type ClientCertificateConfig struct {
	// Whether to issue a client certificate.
	IssueClientCertificate bool `json:"issueClientCertificate,omitempty"`
}

// NodePool defines the node pool configuration.
type NodePool struct {
	// The name of the node pool.
	Name string `json:"name"`
	// The initial number of nodes for the node pool.
	InitialNodeCount int64 `json:"initialNodeCount,omitempty"`
	// The configuration for the nodes in the node pool.
	Config *NodeConfig `json:"config,omitempty"`
	// The autoscaling configuration for the node pool.
	Autoscaling *NodePoolAutoscaling `json:"autoscaling,omitempty"`
	// The management configuration for the node pool.
	Management *NodeManagement `json:"management,omitempty"`
	// The maximum number of pods per node in the node pool.
	MaxPodsPerNode int64 `json:"maxPodsPerNode,omitempty"`
	// The network configuration for the node pool.
	NetworkConfig *NodeNetworkConfig `json:"networkConfig,omitempty"`
}

// NodeConfig defines the node configuration.
type NodeConfig struct {
	// The machine type for the nodes.
	MachineType string `json:"machineType,omitempty"`
	// The disk size for the nodes.
	DiskSizeGb int64 `json:"diskSizeGb,omitempty"`
	// The OAuth scopes for the nodes.
	OAuthScopes []string `json:"oauthScopes,omitempty"`
	// The service account for the nodes.
	ServiceAccount string `json:"serviceAccount,omitempty"`
	// The metadata for the nodes.
	Metadata map[string]string `json:"metadata,omitempty"`
	// The image type for the nodes.
	ImageType string `json:"imageType,omitempty"`
	// The labels for the nodes.
	Labels map[string]string `json:"labels,omitempty"`
	// The local SSD count for the nodes.
	LocalSsdCount int64 `json:"localSsdCount,omitempty"`
	// The tags for the nodes.
	Tags []string `json:"tags,omitempty"`
	// Whether the nodes are preemptible.
	Preemptible bool `json:"preemptible,omitempty"`
	// The accelerators for the nodes.
	Accelerators []*AcceleratorConfig `json:"accelerators,omitempty"`
	// The disk type for the nodes.
	DiskType string `json:"diskType,omitempty"`
	// The minimum CPU platform for the nodes.
	MinCpuPlatform string `json:"minCpuPlatform,omitempty"`
	// The workload metadata configuration for the nodes.
	WorkloadMetadataConfig *WorkloadMetadataConfig `json:"workloadMetadataConfig,omitempty"`
	// The taints for the nodes.
	Taints []*NodeTaint `json:"taints,omitempty"`
	// The sandbox configuration for the nodes.
	SandboxConfig *SandboxConfig `json:"sandboxConfig,omitempty"`
	// The node group for the nodes.
	NodeGroup string `json:"nodeGroup,omitempty"`
	// The reservation affinity for the nodes.
	ReservationAffinity *ReservationAffinity `json:"reservationAffinity,omitempty"`
	// The shielded instance configuration for the nodes.
	ShieldedInstanceConfig *ShieldedInstanceConfig `json:"shieldedInstanceConfig,omitempty"`
	// The Linux node configuration for the nodes.
	LinuxNodeConfig *LinuxNodeConfig `json:"linuxNodeConfig,omitempty"`
	// The kubelet configuration for the nodes.
	KubeletConfig *NodeKubeletConfig `json:"kubeletConfig,omitempty"`
	// The boot disk KMS key for the nodes.
	BootDiskKmsKey string `json:"bootDiskKmsKey,omitempty"`
	// The GCFS configuration for the nodes.
	GcfsConfig *GcfsConfig `json:"gcfsConfig,omitempty"`
}

// NodePoolAutoscaling defines the node pool autoscaling configuration.
type NodePoolAutoscaling struct {
	// Whether autoscaling is enabled.
	Enabled bool `json:"enabled,omitempty"`
	// The minimum number of nodes in the node pool.
	MinNodeCount int64 `json:"minNodeCount,omitempty"`
	// The maximum number of nodes in the node pool.
	MaxNodeCount int64 `json:"maxNodeCount,omitempty"`
}

// NodeManagement defines the node management configuration.
type NodeManagement struct {
	// Whether auto upgrade is enabled.
	AutoUpgrade bool `json:"autoUpgrade,omitempty"`
	// Whether auto repair is enabled.
	AutoRepair bool `json:"autoRepair,omitempty"`
}

// NodeNetworkConfig defines the node network configuration.
type NodeNetworkConfig struct {
	// Whether to create a pod range.
	CreatePodRange bool `json:"createPodRange,omitempty"`
	// The pod range.
	PodRange string `json:"podRange,omitempty"`
	// The pod IPv4 CIDR block.
	PodIpv4CidrBlock string `json:"podIpv4CidrBlock,omitempty"`
}

// AcceleratorConfig defines the accelerator configuration.
type AcceleratorConfig struct {
	// The number of accelerators.
	AcceleratorCount int64 `json:"acceleratorCount,omitempty"`
	// The type of accelerator.
	AcceleratorType string `json:"acceleratorType,omitempty"`
}

// WorkloadMetadataConfig defines the workload metadata configuration.
type WorkloadMetadataConfig struct {
	// The node metadata.
	NodeMetadata string `json:"nodeMetadata,omitempty"`
}

// NodeTaint defines the node taint configuration.
type NodeTaint struct {
	// The key of the taint.
	Key string `json:"key,omitempty"`
	// The value of the taint.
	Value string `json:"value,omitempty"`
	// The effect of the taint.
	Effect string `json:"effect,omitempty"`
}

// SandboxConfig defines the sandbox configuration.
type SandboxConfig struct {
	// The type of sandbox.
	Type string `json:"type,omitempty"`
}

// ReservationAffinity defines the reservation affinity configuration.
type ReservationAffinity struct {
	// The type of reservation consumption.
	ConsumeReservationType string `json:"consumeReservationType,omitempty"`
	// The key of the reservation.
	Key string `json:"key,omitempty"`
	// The values of the reservation.
	Values []string `json:"values,omitempty"`
}

// ShieldedInstanceConfig defines the shielded instance configuration.
type ShieldedInstanceConfig struct {
	// Whether to enable secure boot.
	EnableSecureBoot bool `json:"enableSecureBoot,omitempty"`
	// Whether to enable integrity monitoring.
	EnableIntegrityMonitoring bool `json:"enableIntegrityMonitoring,omitempty"`
}

// LinuxNodeConfig defines the Linux node configuration.
type LinuxNodeConfig struct {
	// The sysctls for the nodes.
	Sysctls map[string]string `json:"sysctls,omitempty"`
}

// NodeKubeletConfig defines the node kubelet configuration.
type NodeKubeletConfig struct {
	// The CPU manager policy.
	CpuManagerPolicy string `json:"cpuManagerPolicy,omitempty"`
	// Whether the CPU CFS quota is enabled.
	CpuCfsQuota bool `json:"cpuCfsQuota,omitempty"`
	// The CPU CFS quota period.
	CpuCfsQuotaPeriod string `json:"cpuCfsQuotaPeriod,omitempty"`
}

// GcfsConfig defines the GCFS configuration.
type GcfsConfig struct {
	// Whether GCFS is enabled.
	Enabled bool `json:"enabled,omitempty"`
}

// AddonsConfig defines the addons configuration.
type AddonsConfig struct {
	// The HTTP load balancing configuration.
	HttpLoadBalancing *HttpLoadBalancing `json:"httpLoadBalancing,omitempty"`
	// The horizontal pod autoscaling configuration.
	HorizontalPodAutoscaling *HorizontalPodAutoscaling `json:"horizontalPodAutoscaling,omitempty"`
	// The network policy configuration.
	NetworkPolicyConfig *NetworkPolicyConfig `json:"networkPolicyConfig,omitempty"`
	// The Cloud Run configuration.
	CloudRunConfig *CloudRunConfig `json:"cloudRunConfig,omitempty"`
	// The DNS cache configuration.
	DnsCacheConfig *DnsCacheConfig `json:"dnsCacheConfig,omitempty"`
	// The Config Connector configuration.
	ConfigConnectorConfig *ConfigConnectorConfig `json:"configConnectorConfig,omitempty"`
	// The GCE persistent disk CSI driver configuration.
	GcePersistentDiskCsiDriverConfig *GcePersistentDiskCsiDriverConfig `json:"gcePersistentDiskCsiDriverConfig,omitempty"`
	// The GKE backup agent configuration.
	GkeBackupAgentConfig *GkeBackupAgentConfig `json:"gkeBackupAgentConfig,omitempty"`
}

// HttpLoadBalancing defines the HTTP load balancing configuration.
type HttpLoadBalancing struct {
	// Whether HTTP load balancing is disabled.
	Disabled bool `json:"disabled,omitempty"`
}

// HorizontalPodAutoscaling defines the horizontal pod autoscaling configuration.
type HorizontalPodAutoscaling struct {
	// Whether horizontal pod autoscaling is disabled.
	Disabled bool `json:"disabled,omitempty"`
}

// NetworkPolicyConfig defines the network policy configuration.
type NetworkPolicyConfig struct {
	// Whether network policy is disabled.
	Disabled bool `json:"disabled,omitempty"`
}

// CloudRunConfig defines the Cloud Run configuration.
type CloudRunConfig struct {
	// Whether Cloud Run is disabled.
	Disabled bool `json:"disabled,omitempty"`
}

// DnsCacheConfig defines the DNS cache configuration.
type DnsCacheConfig struct {
	// Whether DNS cache is enabled.
	Enabled bool `json:"enabled,omitempty"`
}

// ConfigConnectorConfig defines the Config Connector configuration.
type ConfigConnectorConfig struct {
	// Whether Config Connector is enabled.
	Enabled bool `json:"enabled,omitempty"`
}

// GcePersistentDiskCsiDriverConfig defines the GCE persistent disk CSI driver configuration.
type GcePersistentDiskCsiDriverConfig struct {
	// Whether the GCE persistent disk CSI driver is enabled.
	Enabled bool `json:"enabled,omitempty"`
}

// GkeBackupAgentConfig defines the GKE backup agent configuration.
type GkeBackupAgentConfig struct {
	// Whether the GKE backup agent is enabled.
	Enabled bool `json:"enabled,omitempty"`
}

// NetworkPolicy defines the network policy configuration.
type NetworkPolicy struct {
	// The network policy provider.
	Provider string `json:"provider,omitempty"`
	// Whether network policy is enabled.
	Enabled bool `json:"enabled,omitempty"`
}

// IPAllocationPolicy defines the IP allocation policy configuration.
type IPAllocationPolicy struct {
	// Whether to use IP aliases.
	UseIpAliases bool `json:"useIpAliases,omitempty"`
	// Whether to create a subnetwork.
	CreateSubnetwork bool `json:"createSubnetwork,omitempty"`
	// The name of the subnetwork.
	SubnetworkName string `json:"subnetworkName,omitempty"`
	// The cluster IPv4 CIDR block.
	ClusterIpv4CidrBlock string `json:"clusterIpv4CidrBlock,omitempty"`
	// The services IPv4 CIDR block.
	ServicesIpv4CidrBlock string `json:"servicesIpv4CidrBlock,omitempty"`
}

// MasterAuthorizedNetworksConfig defines the master authorized networks configuration.
type MasterAuthorizedNetworksConfig struct {
	// Whether master authorized networks is enabled.
	Enabled bool `json:"enabled,omitempty"`
	// The CIDR blocks for master authorized networks.
	CidrBlocks []*CidrBlock `json:"cidrBlocks,omitempty"`
}

// CidrBlock defines the CIDR block configuration.
type CidrBlock struct {
	// The display name of the CIDR block.
	DisplayName string `json:"displayName,omitempty"`
	// The CIDR block.
	CidrBlock string `json:"cidrBlock,omitempty"`
}

// PrivateClusterConfig defines the private cluster configuration.
type PrivateClusterConfig struct {
	// Whether to enable private nodes.
	EnablePrivateNodes bool `json:"enablePrivateNodes,omitempty"`
	// Whether to enable a private endpoint.
	EnablePrivateEndpoint bool `json:"enablePrivateEndpoint,omitempty"`
	// The master IPv4 CIDR block.
	MasterIpv4CidrBlock string `json:"masterIpv4CidrBlock,omitempty"`
	// The private endpoint.
	PrivateEndpoint string `json:"privateEndpoint,omitempty"`
	// The public endpoint.
	PublicEndpoint string `json:"publicEndpoint,omitempty"`
}

// BinaryAuthorization defines the binary authorization configuration.
type BinaryAuthorization struct {
	// Whether binary authorization is enabled.
	Enabled bool `json:"enabled,omitempty"`
}

// ShieldedNodes defines the shielded nodes configuration.
type ShieldedNodes struct {
	// Whether shielded nodes is enabled.
	Enabled bool `json:"enabled,omitempty"`
}

// ReleaseChannel defines the release channel configuration.
type ReleaseChannel struct {
	// The channel of the release.
	Channel string `json:"channel,omitempty"`
}

// WorkloadIdentityConfig defines the workload identity configuration.
type WorkloadIdentityConfig struct {
	// The workload pool.
	WorkloadPool string `json:"workloadPool,omitempty"`
}

// NotificationConfig defines the notification configuration.
type NotificationConfig struct {
	// The pub/sub topic for notifications.
	Pubsub *Pubsub `json:"pubsub,omitempty"`
}

// Pubsub defines the pub/sub configuration.
type Pubsub struct {
	// Whether pub/sub is enabled.
	Enabled bool `json:"enabled,omitempty"`
	// The pub/sub topic.
	Topic string `json:"topic,omitempty"`
}

// ConfidentialNodes defines the confidential nodes configuration.
type ConfidentialNodes struct {
	// Whether confidential nodes is enabled.
	Enabled bool `json:"enabled,omitempty"`
}

// AuthenticatorGroupsConfig defines the authenticator groups configuration.
type AuthenticatorGroupsConfig struct {
	// Whether authenticator groups is enabled.
	Enabled bool `json:"enabled,omitempty"`
	// The security group for authenticator groups.
	SecurityGroup string `json:"securityGroup,omitempty"`
}

// ResourceUsageExportConfig defines the resource usage export configuration.
type ResourceUsageExportConfig struct {
	// The BigQuery destination for resource usage export.
	BigqueryDestination *BigQueryDestination `json:"bigqueryDestination,omitempty"`
	// Whether to enable network egress metering.
	EnableNetworkEgressMetering bool `json:"enableNetworkEgressMetering,omitempty"`
	// The consumption metering configuration.
	ConsumptionMeteringConfig *ConsumptionMeteringConfig `json:"consumptionMeteringConfig,omitempty"`
}

// BigQueryDestination defines the BigQuery destination configuration.
type BigQueryDestination struct {
	// The dataset ID for the BigQuery destination.
	DatasetId string `json:"datasetId,omitempty"`
}

// ConsumptionMeteringConfig defines the consumption metering configuration.
type ConsumptionMeteringConfig struct {
	// Whether consumption metering is enabled.
	Enabled bool `json:"enabled,omitempty"`
}

// VerticalPodAutoscaling defines the vertical pod autoscaling configuration.
type VerticalPodAutoscaling struct {
	// Whether vertical pod autoscaling is enabled.
	Enabled bool `json:"enabled,omitempty"`
}

func init() {
	SchemeBuilder.Register(&ContainerCluster{}, &ContainerClusterList{})
}
