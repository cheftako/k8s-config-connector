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
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func (c *ContainerCluster) GetID() (string, error) {
	if c.Spec.ResourceID == nil {
		return "", fmt.Errorf("resourceID is empty")
	}
	return *c.Spec.ResourceID, nil
}

func (c *ContainerCluster) SetID(id string) {
	if c.Spec.ResourceID == nil {
		c.Spec.ResourceID = new(string)
	}
	*c.Spec.ResourceID = id
}

func (c *ContainerCluster) GetObservedGeneration() int64 {
	return c.Status.ObservedGeneration
}

func (c *ContainerCluster) SetObservedGeneration(generation int64) {
	c.Status.ObservedGeneration = generation
}

func (c *ContainerCluster) GetConditions() []v1alpha1.Condition {
	return c.Status.Conditions
}

func (c *ContainerCluster) SetConditions(conditions []v1alpha1.Condition) {
	c.Status.Conditions = conditions
}

func (c *ContainerCluster) GetLabels() map[string]string {
	return c.Labels
}

func (c *ContainerCluster) GetAnnotations() map[string]string {
	return c.Annotations
}

func (c *ContainerCluster) GetExternalRef() (string, bool) {
	return direct.GetExternal(c)
}

func (c *ContainerCluster) SetExternalRef(external string) {
	direct.SetExternal(c, external)
}