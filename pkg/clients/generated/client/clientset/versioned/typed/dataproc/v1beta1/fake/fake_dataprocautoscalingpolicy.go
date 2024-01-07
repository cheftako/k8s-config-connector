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

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/dataproc/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDataprocAutoscalingPolicies implements DataprocAutoscalingPolicyInterface
type FakeDataprocAutoscalingPolicies struct {
	Fake *FakeDataprocV1beta1
	ns   string
}

var dataprocautoscalingpoliciesResource = v1beta1.SchemeGroupVersion.WithResource("dataprocautoscalingpolicies")

var dataprocautoscalingpoliciesKind = v1beta1.SchemeGroupVersion.WithKind("DataprocAutoscalingPolicy")

// Get takes name of the dataprocAutoscalingPolicy, and returns the corresponding dataprocAutoscalingPolicy object, and an error if there is any.
func (c *FakeDataprocAutoscalingPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.DataprocAutoscalingPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dataprocautoscalingpoliciesResource, c.ns, name), &v1beta1.DataprocAutoscalingPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DataprocAutoscalingPolicy), err
}

// List takes label and field selectors, and returns the list of DataprocAutoscalingPolicies that match those selectors.
func (c *FakeDataprocAutoscalingPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.DataprocAutoscalingPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dataprocautoscalingpoliciesResource, dataprocautoscalingpoliciesKind, c.ns, opts), &v1beta1.DataprocAutoscalingPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.DataprocAutoscalingPolicyList{ListMeta: obj.(*v1beta1.DataprocAutoscalingPolicyList).ListMeta}
	for _, item := range obj.(*v1beta1.DataprocAutoscalingPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dataprocAutoscalingPolicies.
func (c *FakeDataprocAutoscalingPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dataprocautoscalingpoliciesResource, c.ns, opts))

}

// Create takes the representation of a dataprocAutoscalingPolicy and creates it.  Returns the server's representation of the dataprocAutoscalingPolicy, and an error, if there is any.
func (c *FakeDataprocAutoscalingPolicies) Create(ctx context.Context, dataprocAutoscalingPolicy *v1beta1.DataprocAutoscalingPolicy, opts v1.CreateOptions) (result *v1beta1.DataprocAutoscalingPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dataprocautoscalingpoliciesResource, c.ns, dataprocAutoscalingPolicy), &v1beta1.DataprocAutoscalingPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DataprocAutoscalingPolicy), err
}

// Update takes the representation of a dataprocAutoscalingPolicy and updates it. Returns the server's representation of the dataprocAutoscalingPolicy, and an error, if there is any.
func (c *FakeDataprocAutoscalingPolicies) Update(ctx context.Context, dataprocAutoscalingPolicy *v1beta1.DataprocAutoscalingPolicy, opts v1.UpdateOptions) (result *v1beta1.DataprocAutoscalingPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dataprocautoscalingpoliciesResource, c.ns, dataprocAutoscalingPolicy), &v1beta1.DataprocAutoscalingPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DataprocAutoscalingPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDataprocAutoscalingPolicies) UpdateStatus(ctx context.Context, dataprocAutoscalingPolicy *v1beta1.DataprocAutoscalingPolicy, opts v1.UpdateOptions) (*v1beta1.DataprocAutoscalingPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dataprocautoscalingpoliciesResource, "status", c.ns, dataprocAutoscalingPolicy), &v1beta1.DataprocAutoscalingPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DataprocAutoscalingPolicy), err
}

// Delete takes name of the dataprocAutoscalingPolicy and deletes it. Returns an error if one occurs.
func (c *FakeDataprocAutoscalingPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(dataprocautoscalingpoliciesResource, c.ns, name, opts), &v1beta1.DataprocAutoscalingPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDataprocAutoscalingPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dataprocautoscalingpoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.DataprocAutoscalingPolicyList{})
	return err
}

// Patch applies the patch and returns the patched dataprocAutoscalingPolicy.
func (c *FakeDataprocAutoscalingPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.DataprocAutoscalingPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dataprocautoscalingpoliciesResource, c.ns, name, pt, data, subresources...), &v1beta1.DataprocAutoscalingPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DataprocAutoscalingPolicy), err
}
