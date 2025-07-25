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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func (c *ContainerCluster) GetManagementRefs() []v1alpha1.ManagementRef {
	return nil
}

func (c *ContainerCluster) GetResourceRefs() []v1alpha1.ResourceRef {
	var refs []v1alpha1.ResourceRef
	if c.Spec.NetworkRef != nil {
		refs = append(refs, *c.Spec.NetworkRef)
	}
	if c.Spec.SubnetworkRef != nil {
		refs = append(refs, *c.Spec.SubnetworkRef)
	}
	if c.Spec.NodeConfig != nil {
		if c.Spec.NodeConfig.ServiceAccountRef != nil {
			refs = append(refs, *c.Spec.NodeConfig.ServiceAccountRef)
		}
		if c.Spec.NodeConfig.BootDiskKmsKeyRef != nil {
			refs = append(refs, *c.Spec.NodeConfig.BootDiskKmsKeyRef)
		}
	}
	if c.Spec.PrivateClusterConfig != nil && c.Spec.PrivateClusterConfig.PeeringRef != nil {
		refs = append(refs, *c.Spec.PrivateClusterConfig.PeeringRef)
	}
	return refs
}
