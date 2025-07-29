// Copyright 2024 Google LLC
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
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func init() {
	registry.RegisterModel(ContainerClusterGVK, NewContainerClusterModel)
}

func NewContainerClusterModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &containerClusterModel{config: *config}, nil
}

var _ directbase.Model = &containerClusterModel{}

type containerClusterModel struct {
	config config.ControllerConfig
}

func (m *containerClusterModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &ContainerCluster{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	return &containerClusterAdapter{
		gcpResource: obj,
		model:       m,
	}, nil
}

func (m *containerClusterModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	return nil, nil
}

type containerClusterAdapter struct {
	gcpResource *ContainerCluster
	model       *containerClusterModel
}

var _ directbase.Adapter = &containerClusterAdapter{}

func (a *containerClusterAdapter) Find(ctx context.Context) (bool, error) {
	return false, nil
}

func (a *containerClusterAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	return nil
}

func (a *containerClusterAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	return nil
}

func (a *containerClusterAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	return false, nil
}
func (a *containerClusterAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	return nil, nil
}
