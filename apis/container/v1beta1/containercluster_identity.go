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
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
)

func init() {
	registry.RegisterModel(v1alpha1.SchemeGroupVersion.WithKind("ContainerCluster"), &containerClusterModel{})
}

type containerClusterModel struct {
	directbase.CoreModel
}

func (m *containerClusterModel) Build(c direct.Context, b *direct.Builder, resource *v1alpha1.Resource) (interface{}, error) {
	gcp, err := m.buildGCP(c, b, resource)
	if err != nil {
		return nil, err
	}
	return gcp, nil
}

func (m *containerClusterModel) buildGCP(c direct.Context, b *direct.Builder, resource *v1alpha1.Resource) (interface{}, error) {
	gcp := &ContainerCluster{}
	if err := resource.Spec.UnmarshalTo(gcp); err != nil {
		return nil, fmt.Errorf("error unmarshalling to ContainerCluster: %w", err)
	}
	return gcp, nil
}
