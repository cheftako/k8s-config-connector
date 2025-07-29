// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1beta1

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ContainerClusterGVK = GroupVersion.WithKind("ContainerCluster")

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcontainercluster;gcpcontainerclusters
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// +kubebuilder:printcolumn:name="Endpoint",type="string",JSONPath=".status.observedState.endpoint"
// +kubebuilder:printcolumn:name="MasterVersion",type="string",JSONPath=".status.observedState.currentMasterVersion"

// ContainerCluster is the Schema for the container API
// +k8s:openapi-gen=true
type ContainerCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ContainerClusterSpec   `json:"spec,omitempty"`
	Status ContainerClusterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ContainerClusterList contains a list of ContainerCluster
type ContainerClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContainerCluster `json:"items"`
}

// +kcc:spec:proto=google.container.v1.Cluster
type ContainerClusterSpec struct {
	/* Immutable. The Project that this resource belongs to. */
	// +required
	ProjectRef *refs.ProjectRef `json:"projectRef"`

	/* Immutable. Location of the instance. */
	// +required
	Location string `json:"location"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// +optional
	Description *string `json:"description,omitempty"`
	// The number of nodes to create in this cluster. You must ensure that your
	// Compute Engine [resource
	// quota](https://cloud.google.com/compute/docs/resource-quotas) is sufficient
	// for this number of nodes. You must also have available firewall and routes
	// capacity.
	//
	// This field is deprecated and will be removed in a future version of GKE.
	// Use the `initial_node_count` field in the `node_pool` message instead.
	// +optional
	InitialNodeCount *int32 `json:"initialNodeCount,omitempty"`
	// The authentication information for accessing the master endpoint.
	// If unspecified, the defaults are used:
	// On clusters before v1.12, basic authentication is enabled, a client
	// certificate is issued, and IP-based authentication is disabled.
	// On clusters v1.12 and later, basic authentication is disabled, a client
	// certificate is issued, and IP-based authentication is disabled.
	// [IP-based authentication](https://cloud.google.com/kubernetes-engine/docs/how-to/master-authorized-networks)
	// can be enabled by specifying `master_authorized_networks_config`.
	// +optional
	MasterAuth *ContainerClusterMasterAuth `json:"masterAuth,omitempty"`
	// The logging service the cluster should use to write logs.
	// Currently available options:
	//
	// * `logging.googleapis.com/kubernetes` - The Cloud Logging service with a
	//   Kubernetes-native resource model
	// * `logging.googleapis.com` - The legacy Cloud Logging service (no longer
	//   available as of GKE 1.15).
	// * `none` - no logs will be exported from the cluster.
	//
	// If left as an empty string,`logging.googleapis.com/kubernetes` will be
	// used for GKE 1.14+ or `logging.googleapis.com` for earlier versions.
	// +optional
	LoggingService *string `json:"loggingService,omitempty"`
	// The monitoring service the cluster should use to write metrics.
	// Currently available options:
	//
	// * `monitoring.googleapis.com/kubernetes` - The Cloud Monitoring service with
	//   a Kubernetes-native resource model
	// * `monitoring.googleapis.com` - The legacy Cloud Monitoring service (no
	//   longer available as of GKE 1.15).
	// * `none` - no metrics will be exported from the cluster.
	//
	// If left as an empty string,`monitoring.googleapis.com/kubernetes` will be
	// used for GKE 1.14+ or `monitoring.googleapis.com` for earlier versions.
	// +optional
	MonitoringService *string `json:"monitoringService,omitempty"`
	// The name of the Google Compute Engine
	// [network](https://cloud.google.com/compute/docs/networks-and-firewalls#networks)
	// to which the cluster is connected. If left unspecified, the `default`
	// network will be used.
	// +optional
	Network *string `json:"network,omitempty"`
	// +optional
	NetworkRef *refs.ComputeNetworkRef `json:"networkRef,omitempty"`
	// The IP address range of the container pods in this cluster, in
	// [CIDR](http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing)
	// notation (e.g. `10.96.0.0/14`). Leave blank to have one automatically
	// chosen or specify a `/14` block in `10.0.0.0/8`.
	// +optional
	ClusterIpv4Cidr *string `json:"clusterIpv4Cidr,omitempty"`
	// Configuration for the addons that can be automatically spun up in the
	// cluster, enabling additional functionality.
	// +optional
	AddonsConfig *ContainerClusterAddonsConfig `json:"addonsConfig,omitempty"`
	// The name of the Google Compute Engine
	// [subnetwork](https://cloud.google.com/compute/docs/subnetworks) to which
	// the cluster is connected.
	// +optional
	Subnetwork *string `json:"subnetwork,omitempty"`
	// +optional
	SubnetworkRef *refs.ComputeSubnetworkRef `json:"subnetworkRef,omitempty"`
	// The node pools associated with this cluster.
	// This field should not be set if "node_config" or "initial_node_count" are
	// specified.
	// +optional
	NodePools []ContainerClusterNodePools `json:"nodePools,omitempty"`
}

type ContainerClusterStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ContainerCluster resource in GCP.
	// +optional
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// +optional
	ObservedState *ContainerClusterObservedState `json:"observedState,omitempty"`
}

// +kcc:observedstate:proto=google.container.v1.Cluster
type ContainerClusterObservedState struct {
	// The IP address of this cluster's master endpoint.
	// The endpoint can be accessed from the internet at
	// `https://username:password@endpoint/`.
	//
	// See the `masterAuth` property of this resource for username and password
	// information.
	// +optional
	Endpoint *string `json:"endpoint,omitempty"`
	// The current software version of the master endpoint.
	// +optional
	CurrentMasterVersion *string `json:"currentMasterVersion,omitempty"`
	// The current version of the node software components.
	// If they are currently at multiple versions because of a version update,
	// this field follows the lowest version among all nodes.
	// +optional
	CurrentNodeVersion *string `json:"currentNodeVersion,omitempty"`
	// The time the cluster was created, in
	// [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
	// +optional
	CreateTime *string `json:"createTime,omitempty"`
	// The current status of this cluster.
	// +optional
	Status *string `json:"status,omitempty"`
	// A human-readable description of the status of this cluster.
	// +optional
	StatusMessage *string `json:"statusMessage,omitempty"`
	// The size of the address space on each node for hosting pods.
	// +optional
	NodeIpv4CidrSize *int32 `json:"nodeIpv4CidrSize,omitempty"`
	// The IP address range of the Kubernetes services in this cluster, in
	// [CIDR](http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing)
	// notation (e.g. `1.2.3.4/29`). Service addresses are typically assigned
	// randomly from this range. If left blank, a range will be chosen with the
	// default size.
	// +optional
	ServicesIpv4Cidr *string `json:"servicesIpv4Cidr,omitempty"`
	// The default maximum number of pods per node in this cluster.
	// This doesn't reflect the maximum number of pods per node in managed nodepools.
	// +optional
	DefaultMaxPodsPerNode *int32 `json:"defaultMaxPodsPerNode,omitempty"`
	// The time the cluster will be automatically deleted in
	// [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
	// +optional
	ExpireTime *string `json:"expireTime,omitempty"`
}

// +kcc:proto=google.container.v1.MasterAuth
type ContainerClusterMasterAuth struct {
	// The username to use for HTTP basic authentication to the master endpoint.
	// For clusters v1.6.0 and later, basic authentication can be disabled by
	// leaving username unspecified (or setting it to the empty string).
	//
	// Warning: basic authentication is deprecated, and will be removed in GKE
	// control plane versions 1.19 and newer. For a list of recommended
	// authentication methods, see:
	// https://cloud.google.com/kubernetes-engine/docs/how-to/api-server-authentication
	// +optional
	Username *string `json:"username,omitempty"`
	// The password to use for HTTP basic authentication to the master endpoint.
	// Because the master endpoint is open to the Internet, you should create a
	// strong password.  If a password is provided for cluster creation, username
	// must be non-empty.
	//
	// Warning: basic authentication is deprecated, and will be removed in GKE
	// control plane versions 1.19 and newer. For a list of recommended
	// authentication methods, see:
	// https://cloud.google.com/kubernetes-engine/docs/how-to/api-server-authentication
	// +optional
	Password *string `json:"password,omitempty"`
	// Configuration for client certificate authentication on the cluster. For
	// clusters before v1.12, if no configuration is specified, a client
	// certificate is issued.
	// +optional
	ClientCertificateConfig *ContainerClusterClientCertificateConfig `json:"clientCertificateConfig,omitempty"`
}

// +kcc:proto=google.container.v1.ClientCertificateConfig
type ContainerClusterClientCertificateConfig struct {
	// Issue a client certificate.
	IssueClientCertificate bool `json:"issueClientCertificate"`
}

// +kcc:proto=google.container.v1.AddonsConfig
type ContainerClusterAddonsConfig struct {
	// Configuration for the HTTP (L7) load balancing controller addon, which
	// makes it easy to set up HTTP load balancers for services in a cluster.
	// +optional
	HttpLoadBalancing *ContainerClusterHttpLoadBalancing `json:"httpLoadBalancing,omitempty"`
	// Configuration for the horizontal pod autoscaling feature, which
	// increases or decreases the number of replica pods a replication controller
	// has based on the resource usage of the existing pods.
	// +optional
	HorizontalPodAutoscaling *ContainerClusterHorizontalPodAutoscaling `json:"horizontalPodAutoscaling,omitempty"`
	// Configuration for NetworkPolicy. This only tracks whether the addon
	// is enabled or not on the Master, it does not track whether network policy
	// is enabled for the nodes.
	// +optional
	NetworkPolicyConfig *ContainerClusterNetworkPolicyConfig `json:"networkPolicyConfig,omitempty"`
}

// +kcc:proto=google.container.v1.HttpLoadBalancing
type ContainerClusterHttpLoadBalancing struct {
	// Whether the HTTP Load Balancing controller is enabled in the cluster.
	// When enabled, it runs a small pod in the cluster that manages the load
	// balancers.
	Disabled bool `json:"disabled"`
}

// +kcc:proto=google.container.v1.HorizontalPodAutoscaling
type ContainerClusterHorizontalPodAutoscaling struct {
	// Whether the Horizontal Pod Autoscaling feature is enabled in the cluster.
	// When enabled, it ensures that metrics are collected into Stackdriver
	// Monitoring.
	Disabled bool `json:"disabled"`
}

// +kcc:proto=google.container.v1.NetworkPolicyConfig
type ContainerClusterNetworkPolicyConfig struct {
	// Whether NetworkPolicy is enabled for this cluster.
	Disabled bool `json:"disabled"`
}

// +kcc:proto=google.container.v1.NodePool
type ContainerClusterNodePools struct {
	// The name of the node pool.
	Name string `json:"name"`
	// The number of nodes in the node pool.
	// +optional
	InitialNodeCount *int32 `json:"initialNodeCount,omitempty"`
	// The node configuration for the node pool.
	// +optional
	Config *ContainerClusterNodeConfig `json:"config,omitempty"`
	// The version of the node software components.
	// +optional
	Version *string `json:"version,omitempty"`
	// The list of Google Compute Engine
	// [zones](https://cloud.google.com/compute/docs/zones#available) in which the
	// node pool's nodes should be located.
	// +optional
	Locations []string `json:"locations,omitempty"`
}

// +kcc:proto=google.container.v1.NodeConfig
type ContainerClusterNodeConfig struct {
	// The name of a Google Compute Engine [machine
	// type](https://cloud.google.com/compute/docs/machine-types)
	//
	// If unspecified, the default machine type is `e2-medium`.
	// +optional
	MachineType *string `json:"machineType,omitempty"`
	// Size of the disk attached to each node, specified in GB.
	// The smallest allowed disk size is 10GB.
	//
	// If unspecified, the default disk size is 100GB.
	// +optional
	DiskSizeGb *int32 `json:"diskSizeGb,omitempty"`
	// The set of Google API scopes to be made available on all of the
	// node VMs under the "default" service account.
	//
	// The following scopes are recommended, but not required, and by default are
	// not included:
	//
	// * `https://www.googleapis.com/auth/compute` is required for mounting
	// persistent storage on your nodes.
	// * `https://www.googleapis.com/auth/devstorage.read_only` is required for
	// communicating with **gcr.io**
	// (the [Google Container
	// Registry](https://cloud.google.com/container-registry/)).
	//
	// If unspecified, no scopes are added, unless Cloud Logging or Cloud
	// Monitoring are enabled, in which case their required scopes will be added.
	// +optional
	OauthScopes []string `json:"oauthScopes,omitempty"`
	// The Google Cloud Platform Service Account to be used by the node VMs.
	// Specify the email address of the Service Account; otherwise, if no Service
	// Account is specified, the "default" service account is used.
	// +optional
	ServiceAccount *string `json:"serviceAccount,omitempty"`
	// The metadata key/value pairs assigned to instances in the cluster.
	// +optional
	Metadata map[string]string `json:"metadata,omitempty"`
	// The image type to use for this node. Note that for a given image type,
	// the latest version of it will be used. Please see
	// https://cloud.google.com/kubernetes-engine/docs/concepts/node-images for
	// available image types.
	// +optional
	ImageType *string `json:"imageType,omitempty"`
	// The map of Kubernetes labels (key/value pairs) to be applied to each node.
	// +optional
	Labels map[string]string `json:"labels,omitempty"`
	// The number of local SSD disks to be attached to the node.
	// +optional
	LocalSsdCount *int32 `json:"localSsdCount,omitempty"`
	// The list of instance tags applied to all nodes.
	// +optional
	Tags []string `json:"tags,omitempty"`
	// Whether the nodes are created as preemptible VM instances.
	// +optional
	Preemptible *bool `json:"preemptible,omitempty"`
	// A list of hardware accelerators to be attached to each node.
	// +optional
	Accelerators []ContainerClusterAcceleratorConfig `json:"accelerators,omitempty"`
	// Type of the disk attached to each node (e.g. 'pd-standard', 'pd-ssd' or
	// 'pd-balanced')
	//
	// If unspecified, the default disk type is 'pd-standard'
	// +optional
	DiskType *string `json:"diskType,omitempty"`
	// Minimum CPU platform to be used by this instance.
	// +optional
	MinCpuPlatform *string `json:"minCpuPlatform,omitempty"`
}

// +kcc:proto=google.container.v1.AcceleratorConfig
type ContainerClusterAcceleratorConfig struct {
	// The number of the accelerator cards exposed to an instance.
	AcceleratorCount int64 `json:"acceleratorCount"`
	// The accelerator type resource name.
	AcceleratorType string `json:"acceleratorType"`
}

type ContainerClusterRef struct {
	// A reference to an externally managed ContainerCluster resource.
	// Should be in the format `projects/{{project}}/locations/{{location}}/clusters/{{cluster}}`.
	External string `json:"external,omitempty"`
	/* The `name` field of a `ContainerCluster` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `ContainerCluster` resource. */
	Namespace string `json:"namespace,omitempty"`
}

func init() {
	SchemeBuilder.Register(&ContainerCluster{}, &ContainerClusterList{})
}
