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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (ref *ContainerClusterRef) Normalize(ctx context.Context, reader client.Reader, namespace string) error {
	if ref == nil {
		return nil
	}

	if ref.External != "" && ref.Name != "" {
		return fmt.Errorf("cannot specify both name and external on containercluster reference")
	}

	if ref.External != "" {
		return nil
	}

	if ref.Name == "" {
		return fmt.Errorf("must specify either name or external on containercluster reference")
	}

	key := types.NamespacedName{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
	if key.Namespace == "" {
		key.Namespace = namespace
	}

	containerCluster := &unstructured.Unstructured{}
	containerCluster.SetGroupVersionKind(ContainerClusterGVK)
	if err := reader.Get(ctx, key, containerCluster); err != nil {
		if apierrors.IsNotFound(err) {
			return k8s.NewReferenceNotFoundError(containerCluster.GroupVersionKind(), key)
		}
		return fmt.Errorf("error reading referenced ContainerCluster %v: %w", key, err)
	}

	external, _, err := unstructured.NestedString(containerCluster.Object, "status", "externalRef")
	if err != nil {
		return err
	}
	if external == "" {
		return k8s.NewReferenceNotFoundError(containerCluster.GroupVersionKind(), key)
	}
	ref.External = external
	return nil
}

func (ref *ContainerClusterRef) NormalizedExternal(ctx context.Context, reader client.Reader, namespace string) (string, error) {
	if ref == nil {
		return "", nil
	}
	if err := ref.Normalize(ctx, reader, namespace); err != nil {
		return "", err
	}
	if ref.External == "" {
		return "", fmt.Errorf("external field must be set for containercluster reference")
	}
	return ref.External, nil
}
