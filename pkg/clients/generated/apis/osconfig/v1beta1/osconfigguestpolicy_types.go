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

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type GuestpolicyApt struct {
	/* Type of archive files in this repository. The default behavior is DEB. Possible values: ARCHIVE_TYPE_UNSPECIFIED, DEB, DEB_SRC */
	// +optional
	ArchiveType *string `json:"archiveType,omitempty"`

	/* Required. List of components for this repository. Must contain at least one item. */
	// +optional
	Components []string `json:"components,omitempty"`

	/* Required. Distribution of this repository. */
	Distribution string `json:"distribution"`

	/* URI of the key file for this repository. The agent maintains a keyring at `/etc/apt/trusted.gpg.d/osconfig_agent_managed.gpg` containing all the keys in any applied guest policy. */
	// +optional
	GpgKey *string `json:"gpgKey,omitempty"`

	/* Required. URI for this repository. */
	Uri string `json:"uri"`
}

type GuestpolicyArchiveExtraction struct {
	/* Required. The id of the relevant artifact in the recipe. */
	// +optional
	ArtifactId *string `json:"artifactId,omitempty"`

	/* Directory to extract archive to. Defaults to `/` on Linux or `C:` on Windows. */
	// +optional
	Destination *string `json:"destination,omitempty"`

	/* Required. The type of the archive to extract. Possible values: TYPE_UNSPECIFIED, VALIDATION, DESIRED_STATE_CHECK, DESIRED_STATE_ENFORCEMENT, DESIRED_STATE_CHECK_POST_ENFORCEMENT */
	// +optional
	Type *string `json:"type,omitempty"`
}

type GuestpolicyArtifacts struct {
	/* Defaults to false. When false, recipes are subject to validations based on the artifact type: Remote: A checksum must be specified, and only protocols with transport-layer security are permitted. GCS: An object generation number must be specified. */
	// +optional
	AllowInsecure *bool `json:"allowInsecure,omitempty"`

	/* A Google Cloud Storage artifact. */
	// +optional
	Gcs *GuestpolicyGcs `json:"gcs,omitempty"`

	/* Required. Id of the artifact, which the installation and update steps of this recipe can reference. Artifacts in a recipe cannot have the same id. */
	// +optional
	Id *string `json:"id,omitempty"`

	/* A generic remote artifact. */
	// +optional
	Remote *GuestpolicyRemote `json:"remote,omitempty"`
}

type GuestpolicyAssignment struct {
	/* Targets instances matching at least one of these label sets. This allows an assignment to target disparate groups, for example "env=prod or env=staging". */
	// +optional
	GroupLabels []GuestpolicyGroupLabels `json:"groupLabels,omitempty"`

	/* Targets VM instances whose name starts with one of these prefixes. Like labels, this is another way to group VM instances when targeting configs, for example prefix="prod-". Only supported for project-level policies. */
	// +optional
	InstanceNamePrefixes []string `json:"instanceNamePrefixes,omitempty"`

	/*  */
	// +optional
	Instances []v1alpha1.ResourceRef `json:"instances,omitempty"`

	/* Targets VM instances matching at least one of the following OS types. VM instances must match all supplied criteria for a given OsType to be included. */
	// +optional
	OsTypes []GuestpolicyOsTypes `json:"osTypes,omitempty"`

	/* Targets instances in any of these zones. Leave empty to target instances in any zone. Zonal targeting is uncommon and is supported to facilitate the management of changes by zone. */
	// +optional
	Zones []string `json:"zones,omitempty"`
}

type GuestpolicyDpkgInstallation struct {
	/* Required. The id of the relevant artifact in the recipe. */
	// +optional
	ArtifactId *string `json:"artifactId,omitempty"`
}

type GuestpolicyFileCopy struct {
	/* Required. The id of the relevant artifact in the recipe. */
	// +optional
	ArtifactId *string `json:"artifactId,omitempty"`

	/* Required. The absolute path on the instance to put the file. */
	// +optional
	Destination *string `json:"destination,omitempty"`

	/* Whether to allow this step to overwrite existing files. If this is false and the file already exists the file is not overwritten and the step is considered a success. Defaults to false. */
	// +optional
	Overwrite *bool `json:"overwrite,omitempty"`

	/* Consists of three octal digits which represent, in order, the permissions of the owner, group, and other users for the file (similarly to the numeric mode used in the linux chmod utility). Each digit represents a three bit number with the 4 bit corresponding to the read permissions, the 2 bit corresponds to the write bit, and the one bit corresponds to the execute permission. Default behavior is 755. Below are some examples of permissions and their associated values: read, write, and execute: 7 read and execute: 5 read and write: 6 read only: 4 */
	// +optional
	Permissions *string `json:"permissions,omitempty"`
}

type GuestpolicyFileExec struct {
	/* Defaults to [0]. A list of possible return values that the program can return to indicate a success. */
	// +optional
	AllowedExitCodes []int `json:"allowedExitCodes,omitempty"`

	/* Arguments to be passed to the provided executable. */
	// +optional
	Args []string `json:"args,omitempty"`

	/* The id of the relevant artifact in the recipe. */
	// +optional
	ArtifactId *string `json:"artifactId,omitempty"`

	/* The absolute path of the file on the local filesystem. */
	// +optional
	LocalPath *string `json:"localPath,omitempty"`
}

type GuestpolicyGcs struct {
	/*  */
	// +optional
	BucketRef *v1alpha1.ResourceRef `json:"bucketRef,omitempty"`

	/* Must be provided if allow_insecure is false. Generation number of the Google Cloud Storage object. `https://storage.googleapis.com/my-bucket/foo/bar#1234567` this value would be `1234567`. */
	// +optional
	Generation *int `json:"generation,omitempty"`

	/* Name of the Google Cloud Storage object. As specified [here] (https://cloud.google.com/storage/docs/naming#objectnames) Given an example URL: `https://storage.googleapis.com/my-bucket/foo/bar#1234567` this value would be `foo/bar`. */
	// +optional
	Object *string `json:"object,omitempty"`
}

type GuestpolicyGoo struct {
	/* Required. The name of the repository. */
	Name string `json:"name"`

	/* Required. The url of the repository. */
	Url string `json:"url"`
}

type GuestpolicyGroupLabels struct {
	/* Google Compute Engine instance labels that must be present for an instance to be included in this assignment group. */
	// +optional
	Labels map[string]string `json:"labels,omitempty"`
}

type GuestpolicyInstallSteps struct {
	/* Extracts an archive into the specified directory. */
	// +optional
	ArchiveExtraction *GuestpolicyArchiveExtraction `json:"archiveExtraction,omitempty"`

	/* Installs a deb file via dpkg. */
	// +optional
	DpkgInstallation *GuestpolicyDpkgInstallation `json:"dpkgInstallation,omitempty"`

	/* Copies a file onto the instance. */
	// +optional
	FileCopy *GuestpolicyFileCopy `json:"fileCopy,omitempty"`

	/* Executes an artifact or local file. */
	// +optional
	FileExec *GuestpolicyFileExec `json:"fileExec,omitempty"`

	/* Installs an MSI file. */
	// +optional
	MsiInstallation *GuestpolicyMsiInstallation `json:"msiInstallation,omitempty"`

	/* Installs an rpm file via the rpm utility. */
	// +optional
	RpmInstallation *GuestpolicyRpmInstallation `json:"rpmInstallation,omitempty"`

	/* Runs commands in a shell. */
	// +optional
	ScriptRun *GuestpolicyScriptRun `json:"scriptRun,omitempty"`
}

type GuestpolicyMsiInstallation struct {
	/* Return codes that indicate that the software installed or updated successfully. Behaviour defaults to [0] */
	// +optional
	AllowedExitCodes []int `json:"allowedExitCodes,omitempty"`

	/* Required. The id of the relevant artifact in the recipe. */
	// +optional
	ArtifactId *string `json:"artifactId,omitempty"`

	/* The flags to use when installing the MSI defaults to ["/i"] (i.e. the install flag). */
	// +optional
	Flags []string `json:"flags,omitempty"`
}

type GuestpolicyOsTypes struct {
	/* Targets VM instances with OS Inventory enabled and having the following OS architecture. */
	// +optional
	OsArchitecture *string `json:"osArchitecture,omitempty"`

	/* Targets VM instances with OS Inventory enabled and having the following OS short name, for example "debian" or "windows". */
	// +optional
	OsShortName *string `json:"osShortName,omitempty"`

	/* Targets VM instances with OS Inventory enabled and having the following following OS version. */
	// +optional
	OsVersion *string `json:"osVersion,omitempty"`
}

type GuestpolicyPackageRepositories struct {
	/* An Apt Repository. */
	// +optional
	Apt *GuestpolicyApt `json:"apt,omitempty"`

	/* A Goo Repository. */
	// +optional
	Goo *GuestpolicyGoo `json:"goo,omitempty"`

	/* A Yum Repository. */
	// +optional
	Yum *GuestpolicyYum `json:"yum,omitempty"`

	/* A Zypper Repository. */
	// +optional
	Zypper *GuestpolicyZypper `json:"zypper,omitempty"`
}

type GuestpolicyPackages struct {
	/* The desired_state the agent should maintain for this package. The default is to ensure the package is installed. Possible values: DESIRED_STATE_UNSPECIFIED, INSTALLED, REMOVED */
	// +optional
	DesiredState *string `json:"desiredState,omitempty"`

	/* Type of package manager that can be used to install this package. If a system does not have the package manager, the package is not installed or removed no error message is returned. By default, or if you specify `ANY`, the agent attempts to install and remove this package using the default package manager. This is useful when creating a policy that applies to different types of systems. The default behavior is ANY. Possible values: MANAGER_UNSPECIFIED, ANY, APT, YUM, ZYPPER, GOO */
	// +optional
	Manager *string `json:"manager,omitempty"`

	/* Required. The name of the package. A package is uniquely identified for conflict validation by checking the package name and the manager(s) that the package targets. */
	// +optional
	Name *string `json:"name,omitempty"`
}

type GuestpolicyRecipes struct {
	/* Resources available to be used in the steps in the recipe. */
	// +optional
	Artifacts []GuestpolicyArtifacts `json:"artifacts,omitempty"`

	/* Default is INSTALLED. The desired state the agent should maintain for this recipe. INSTALLED: The software recipe is installed on the instance but won't be updated to new versions. UPDATED: The software recipe is installed on the instance. The recipe is updated to a higher version, if a higher version of the recipe is assigned to this instance. REMOVE: Remove is unsupported for software recipes and attempts to create or update a recipe to the REMOVE state is rejected. Possible values: DESIRED_STATE_UNSPECIFIED, INSTALLED, REMOVED */
	// +optional
	DesiredState *string `json:"desiredState,omitempty"`

	/* Actions to be taken for installing this recipe. On failure it stops executing steps and does not attempt another installation. Any steps taken (including partially completed steps) are not rolled back. */
	// +optional
	InstallSteps []GuestpolicyInstallSteps `json:"installSteps,omitempty"`

	/* Required. Unique identifier for the recipe. Only one recipe with a given name is installed on an instance. Names are also used to identify resources which helps to determine whether guest policies have conflicts. This means that requests to create multiple recipes with the same name and version are rejected since they could potentially have conflicting assignments. */
	// +optional
	Name *string `json:"name,omitempty"`

	/* Actions to be taken for updating this recipe. On failure it stops executing steps and does not attempt another update for this recipe. Any steps taken (including partially completed steps) are not rolled back. */
	// +optional
	UpdateSteps []GuestpolicyUpdateSteps `json:"updateSteps,omitempty"`

	/* The version of this software recipe. Version can be up to 4 period separated numbers (e.g. 12.34.56.78). */
	// +optional
	Version *string `json:"version,omitempty"`
}

type GuestpolicyRemote struct {
	/* Must be provided if `allow_insecure` is `false`. SHA256 checksum in hex format, to compare to the checksum of the artifact. If the checksum is not empty and it doesn't match the artifact then the recipe installation fails before running any of the steps. */
	// +optional
	Checksum *string `json:"checksum,omitempty"`

	/* URI from which to fetch the object. It should contain both the protocol and path following the format {protocol}://{location}. */
	// +optional
	Uri *string `json:"uri,omitempty"`
}

type GuestpolicyRpmInstallation struct {
	/* Required. The id of the relevant artifact in the recipe. */
	// +optional
	ArtifactId *string `json:"artifactId,omitempty"`
}

type GuestpolicyScriptRun struct {
	/* Return codes that indicate that the software installed or updated successfully. Behaviour defaults to [0] */
	// +optional
	AllowedExitCodes []int `json:"allowedExitCodes,omitempty"`

	/* The script interpreter to use to run the script. If no interpreter is specified the script is executed directly, which likely only succeed for scripts with [shebang lines](https://en.wikipedia.org/wiki/Shebang_(Unix)). Possible values: INTERPRETER_UNSPECIFIED, NONE, SHELL, POWERSHELL */
	// +optional
	Interpreter *string `json:"interpreter,omitempty"`

	/* Required. The shell script to be executed. */
	// +optional
	Script *string `json:"script,omitempty"`
}

type GuestpolicyUpdateSteps struct {
	/* Extracts an archive into the specified directory. */
	// +optional
	ArchiveExtraction *GuestpolicyArchiveExtraction `json:"archiveExtraction,omitempty"`

	/* Installs a deb file via dpkg. */
	// +optional
	DpkgInstallation *GuestpolicyDpkgInstallation `json:"dpkgInstallation,omitempty"`

	/* Copies a file onto the instance. */
	// +optional
	FileCopy *GuestpolicyFileCopy `json:"fileCopy,omitempty"`

	/* Executes an artifact or local file. */
	// +optional
	FileExec *GuestpolicyFileExec `json:"fileExec,omitempty"`

	/* Installs an MSI file. */
	// +optional
	MsiInstallation *GuestpolicyMsiInstallation `json:"msiInstallation,omitempty"`

	/* Installs an rpm file via the rpm utility. */
	// +optional
	RpmInstallation *GuestpolicyRpmInstallation `json:"rpmInstallation,omitempty"`

	/* Runs commands in a shell. */
	// +optional
	ScriptRun *GuestpolicyScriptRun `json:"scriptRun,omitempty"`
}

type GuestpolicyYum struct {
	/* Required. The location of the repository directory. */
	BaseUrl string `json:"baseUrl"`

	/* The display name of the repository. */
	// +optional
	DisplayName *string `json:"displayName,omitempty"`

	/* URIs of GPG keys. */
	// +optional
	GpgKeys []string `json:"gpgKeys,omitempty"`

	/* Required. A one word, unique name for this repository. This is the `repo id` in the Yum config file and also the `display_name` if `display_name` is omitted. This id is also used as the unique identifier when checking for guest policy conflicts. */
	Id string `json:"id"`
}

type GuestpolicyZypper struct {
	/* Required. The location of the repository directory. */
	BaseUrl string `json:"baseUrl"`

	/* The display name of the repository. */
	// +optional
	DisplayName *string `json:"displayName,omitempty"`

	/* URIs of GPG keys. */
	// +optional
	GpgKeys []string `json:"gpgKeys,omitempty"`

	/* Required. A one word, unique name for this repository. This is the `repo id` in the zypper config file and also the `display_name` if `display_name` is omitted. This id is also used as the unique identifier when checking for guest policy conflicts. */
	Id string `json:"id"`
}

type OSConfigGuestPolicySpec struct {
	/* Specifies the VMs that are assigned this policy. This allows you to target sets or groups of VMs by different parameters such as labels, names, OS, or zones. Empty assignments will target ALL VMs underneath this policy. Conflict Management Policies that exist higher up in the resource hierarchy (closer to the Org) will override those lower down if there is a conflict. At the same level in the resource hierarchy (ie. within a project), the service will prevent the creation of multiple policies that conflict with each other. If there are multiple policies that specify the same config (eg. package, software recipe, repository, etc.), the service will ensure that no VM could potentially receive instructions from both policies. To create multiple policies that specify different versions of a package or different configs for different Operating Systems, each policy must be mutually exclusive in their targeting according to labels, OS, or other criteria. Different configs are identified for conflicts in different ways. Packages are identified by their name and the package manager(s) they target. Package repositories are identified by their unique id where applicable. Some package managers don't have a unique identifier for repositories and where that's the case, no uniqueness is validated by the service. Note that if OS Inventory is disabled, a VM will not be assigned a policy that targets by OS because the service will see this VM's OS as unknown. */
	// +optional
	Assignment *GuestpolicyAssignment `json:"assignment,omitempty"`

	/* Description of the GuestPolicy. Length of the description is limited to 1024 characters. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* List of package repository configurations assigned to the VM instance. */
	// +optional
	PackageRepositories []GuestpolicyPackageRepositories `json:"packageRepositories,omitempty"`

	/* List of package configurations assigned to the VM instance. */
	// +optional
	Packages []GuestpolicyPackages `json:"packages,omitempty"`

	/* Optional. A list of Recipes to install on the VM. */
	// +optional
	Recipes []GuestpolicyRecipes `json:"recipes,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type OSConfigGuestPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   OSConfigGuestPolicy's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* Output only. Time this GuestPolicy was created. */
	CreateTime string `json:"createTime,omitempty"`
	/* The etag for this GuestPolicy. If this is provided on update, it must match the server's etag. */
	Etag string `json:"etag,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	ObservedGeneration int `json:"observedGeneration,omitempty"`
	/* Output only. Last time this GuestPolicy was updated. */
	UpdateTime string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OSConfigGuestPolicy is the Schema for the osconfig API
// +k8s:openapi-gen=true
type OSConfigGuestPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OSConfigGuestPolicySpec   `json:"spec,omitempty"`
	Status OSConfigGuestPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OSConfigGuestPolicyList contains a list of OSConfigGuestPolicy
type OSConfigGuestPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OSConfigGuestPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OSConfigGuestPolicy{}, &OSConfigGuestPolicyList{})
}
