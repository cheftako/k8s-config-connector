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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/osconfig/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeOSConfigOSPolicyAssignments implements OSConfigOSPolicyAssignmentInterface
type FakeOSConfigOSPolicyAssignments struct {
	Fake *FakeOsconfigV1beta1
	ns   string
}

var osconfigospolicyassignmentsResource = v1beta1.SchemeGroupVersion.WithResource("osconfigospolicyassignments")

var osconfigospolicyassignmentsKind = v1beta1.SchemeGroupVersion.WithKind("OSConfigOSPolicyAssignment")

// Get takes name of the oSConfigOSPolicyAssignment, and returns the corresponding oSConfigOSPolicyAssignment object, and an error if there is any.
func (c *FakeOSConfigOSPolicyAssignments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.OSConfigOSPolicyAssignment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(osconfigospolicyassignmentsResource, c.ns, name), &v1beta1.OSConfigOSPolicyAssignment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.OSConfigOSPolicyAssignment), err
}

// List takes label and field selectors, and returns the list of OSConfigOSPolicyAssignments that match those selectors.
func (c *FakeOSConfigOSPolicyAssignments) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.OSConfigOSPolicyAssignmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(osconfigospolicyassignmentsResource, osconfigospolicyassignmentsKind, c.ns, opts), &v1beta1.OSConfigOSPolicyAssignmentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.OSConfigOSPolicyAssignmentList{ListMeta: obj.(*v1beta1.OSConfigOSPolicyAssignmentList).ListMeta}
	for _, item := range obj.(*v1beta1.OSConfigOSPolicyAssignmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested oSConfigOSPolicyAssignments.
func (c *FakeOSConfigOSPolicyAssignments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(osconfigospolicyassignmentsResource, c.ns, opts))

}

// Create takes the representation of a oSConfigOSPolicyAssignment and creates it.  Returns the server's representation of the oSConfigOSPolicyAssignment, and an error, if there is any.
func (c *FakeOSConfigOSPolicyAssignments) Create(ctx context.Context, oSConfigOSPolicyAssignment *v1beta1.OSConfigOSPolicyAssignment, opts v1.CreateOptions) (result *v1beta1.OSConfigOSPolicyAssignment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(osconfigospolicyassignmentsResource, c.ns, oSConfigOSPolicyAssignment), &v1beta1.OSConfigOSPolicyAssignment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.OSConfigOSPolicyAssignment), err
}

// Update takes the representation of a oSConfigOSPolicyAssignment and updates it. Returns the server's representation of the oSConfigOSPolicyAssignment, and an error, if there is any.
func (c *FakeOSConfigOSPolicyAssignments) Update(ctx context.Context, oSConfigOSPolicyAssignment *v1beta1.OSConfigOSPolicyAssignment, opts v1.UpdateOptions) (result *v1beta1.OSConfigOSPolicyAssignment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(osconfigospolicyassignmentsResource, c.ns, oSConfigOSPolicyAssignment), &v1beta1.OSConfigOSPolicyAssignment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.OSConfigOSPolicyAssignment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOSConfigOSPolicyAssignments) UpdateStatus(ctx context.Context, oSConfigOSPolicyAssignment *v1beta1.OSConfigOSPolicyAssignment, opts v1.UpdateOptions) (*v1beta1.OSConfigOSPolicyAssignment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(osconfigospolicyassignmentsResource, "status", c.ns, oSConfigOSPolicyAssignment), &v1beta1.OSConfigOSPolicyAssignment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.OSConfigOSPolicyAssignment), err
}

// Delete takes name of the oSConfigOSPolicyAssignment and deletes it. Returns an error if one occurs.
func (c *FakeOSConfigOSPolicyAssignments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(osconfigospolicyassignmentsResource, c.ns, name, opts), &v1beta1.OSConfigOSPolicyAssignment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOSConfigOSPolicyAssignments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(osconfigospolicyassignmentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.OSConfigOSPolicyAssignmentList{})
	return err
}

// Patch applies the patch and returns the patched oSConfigOSPolicyAssignment.
func (c *FakeOSConfigOSPolicyAssignments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.OSConfigOSPolicyAssignment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(osconfigospolicyassignmentsResource, c.ns, name, pt, data, subresources...), &v1beta1.OSConfigOSPolicyAssignment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.OSConfigOSPolicyAssignment), err
}
