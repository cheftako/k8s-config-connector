// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gkehub

import (
	"context"
	"fmt"
	"strings"

	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/gkehub/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Membership struct {
	id string
}

type Feature struct {
	id string
}

// resolveMembershipRef returns a membership that has membershipId as "projects/*/locations/*/memberships/{membershipId}".
func resolveMembershipRef(ctx context.Context, reader client.Reader, obj *krm.GKEHubFeatureMembership, projectID string) (*Membership, error) {
	name := obj.Spec.MembershipRef.Name
	namespace := obj.Spec.MembershipRef.Namespace
	external := obj.Spec.MembershipRef.External

	if external != "" {
		if name != "" {
			return nil, fmt.Errorf("cannot specify both name and external on membership reference")
		}

		tokens := strings.Split(external, "/")
		if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "memberships" {
			return &Membership{id: external}, nil
		}
		return nil, fmt.Errorf("format of membership external=%q was not known (use projects/*/locations/*/memberships/{membershipId})", external)
	}

	if name == "" {
		return nil, fmt.Errorf("must specify either name or external on membership reference")
	}

	key := types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	}
	if key.Namespace == "" {
		key.Namespace = obj.GetNamespace()
	}

	membership := &unstructured.Unstructured{}
	membership.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "gkehub.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "GKEHubMembership",
	})
	if err := reader.Get(ctx, key, membership); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced %v not found", key)
		}
		return nil, fmt.Errorf("error reading referenced membership %v: %w", key, err)
	}

	membershipName, _, err := unstructured.NestedString(membership.Object, "spec", "resourceID")
	if err != nil {
		return nil, fmt.Errorf("reading spec.resourceID from membership %v: %w", key, err)
	}
	if membershipName == "" {
		membershipName = membership.GetName()
	}
	membershipLocation := ValueOf(obj.Spec.MembershipLocation)
	if membershipLocation == "" {
		// membership location should default to global if not set.
		membershipLocation = "global"
	}
	return &Membership{
		id: fmt.Sprintf("projects/%s/locations/%s/memberships/%s", projectID, membershipLocation, membershipName),
	}, nil
}

// resolveFeatureRef returns a feature that has featureID as "projects/*/locations/*/features/{featureId}".
func resolveFeatureRef(ctx context.Context, reader client.Reader, obj *krm.GKEHubFeatureMembership, projectID string) (*Feature, error) {
	name := obj.Spec.FeatureRef.Name
	namespace := obj.Spec.FeatureRef.Namespace
	external := obj.Spec.FeatureRef.External

	if external != "" {
		if name != "" {
			return nil, fmt.Errorf("cannot specify both name and external on feature reference")
		}

		tokens := strings.Split(external, "/")
		if len(tokens) == 6 && tokens[0] == "projects" && tokens[2] == "locations" && tokens[4] == "features" {
			return &Feature{id: external}, nil
		}
		return nil, fmt.Errorf("format of feature external=%q was not known (use projects/*/locations/*/features/{featureId})", external)
	}

	if name == "" {
		return nil, fmt.Errorf("must specify either name or external on feature reference")
	}

	key := types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	}
	if key.Namespace == "" {
		key.Namespace = obj.GetNamespace()
	}

	feature := &unstructured.Unstructured{}
	feature.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "gkehub.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "GKEHubFeature",
	})
	if err := reader.Get(ctx, key, feature); err != nil {
		if apierrors.IsNotFound(err) {
			return nil, fmt.Errorf("referenced %v not found", key)
		}
		return nil, fmt.Errorf("error reading referenced feature %v: %w", key, err)
	}

	featureName, _, err := unstructured.NestedString(feature.Object, "spec", "resourceID")
	if err != nil {
		return nil, fmt.Errorf("reading spec.resourceID from feature %v: %w", key, err)
	}
	if featureName == "" {
		featureName = feature.GetName()
	}
	featureLocation := obj.Spec.Location
	if featureLocation == "" {
		featureLocation = "global"
	}
	return &Feature{
		id: fmt.Sprintf("projects/%s/locations/%s/features/%s", projectID, featureLocation, featureName),
	}, nil
}

func resolveMetricsGcpServiceAccountRef(ctx context.Context, reader client.Reader, r *v1beta1.MetricsGcpServiceAccountRef, resourceNamespace string) (string, error) {
	name := r.Name
	namespace := r.Namespace
	external := r.External
	return getSAEmailWith(ctx, reader, name, namespace, external, resourceNamespace)
}

func resolveGcpServiceAccountRef(ctx context.Context, reader client.Reader, r *v1beta1.GcpServiceAccountRef, resourceNamespace string) (string, error) {
	name := r.Name
	namespace := r.Namespace
	external := r.External
	return getSAEmailWith(ctx, reader, name, namespace, external, resourceNamespace)

}

func getSAEmailWith(ctx context.Context, reader client.Reader, name, namespace, external, resourceNamespace string) (string, error) {
	if external != "" {
		if name != "" {
			return "", fmt.Errorf("cannot specify both name and external on IAMServiceAccount reference")
		}
		return external, nil
	}
	if name == "" {
		return "", fmt.Errorf("must specify either name or external on IAMServiceAccount reference")
	}
	if namespace == "" {
		namespace = resourceNamespace
	}

	key := types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	}
	sa := &unstructured.Unstructured{}
	sa.SetGroupVersionKind(schema.GroupVersionKind{
		Group:   "iam.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "IAMServiceAccount",
	})
	if err := reader.Get(ctx, key, sa); err != nil {
		if apierrors.IsNotFound(err) {
			return "", fmt.Errorf("referenced IAMServiceAccount %v not found", key)
		}
		return "", fmt.Errorf("error reading referenced IAMServiceAccount %v: %w", key, err)
	}
	email, _, err := unstructured.NestedString(sa.Object, "status", "email")
	if err != nil {
		return "", fmt.Errorf("reading status.email from IAMServiceAccount %v: %w", key, err)
	}
	// if the status.email not populated, should we construct the email from spec.resourceID or metadata.name.
	if email == "" {
		return "", fmt.Errorf("status.email is empty from IAMServiceAccount %v, expected not-empty", key)
	}
	return email, nil
}
