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

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/cloudids/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCloudIDSEndpoints implements CloudIDSEndpointInterface
type FakeCloudIDSEndpoints struct {
	Fake *FakeCloudidsV1alpha1
	ns   string
}

var cloudidsendpointsResource = v1alpha1.SchemeGroupVersion.WithResource("cloudidsendpoints")

var cloudidsendpointsKind = v1alpha1.SchemeGroupVersion.WithKind("CloudIDSEndpoint")

// Get takes name of the cloudIDSEndpoint, and returns the corresponding cloudIDSEndpoint object, and an error if there is any.
func (c *FakeCloudIDSEndpoints) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CloudIDSEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cloudidsendpointsResource, c.ns, name), &v1alpha1.CloudIDSEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudIDSEndpoint), err
}

// List takes label and field selectors, and returns the list of CloudIDSEndpoints that match those selectors.
func (c *FakeCloudIDSEndpoints) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CloudIDSEndpointList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cloudidsendpointsResource, cloudidsendpointsKind, c.ns, opts), &v1alpha1.CloudIDSEndpointList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CloudIDSEndpointList{ListMeta: obj.(*v1alpha1.CloudIDSEndpointList).ListMeta}
	for _, item := range obj.(*v1alpha1.CloudIDSEndpointList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cloudIDSEndpoints.
func (c *FakeCloudIDSEndpoints) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cloudidsendpointsResource, c.ns, opts))

}

// Create takes the representation of a cloudIDSEndpoint and creates it.  Returns the server's representation of the cloudIDSEndpoint, and an error, if there is any.
func (c *FakeCloudIDSEndpoints) Create(ctx context.Context, cloudIDSEndpoint *v1alpha1.CloudIDSEndpoint, opts v1.CreateOptions) (result *v1alpha1.CloudIDSEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cloudidsendpointsResource, c.ns, cloudIDSEndpoint), &v1alpha1.CloudIDSEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudIDSEndpoint), err
}

// Update takes the representation of a cloudIDSEndpoint and updates it. Returns the server's representation of the cloudIDSEndpoint, and an error, if there is any.
func (c *FakeCloudIDSEndpoints) Update(ctx context.Context, cloudIDSEndpoint *v1alpha1.CloudIDSEndpoint, opts v1.UpdateOptions) (result *v1alpha1.CloudIDSEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cloudidsendpointsResource, c.ns, cloudIDSEndpoint), &v1alpha1.CloudIDSEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudIDSEndpoint), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCloudIDSEndpoints) UpdateStatus(ctx context.Context, cloudIDSEndpoint *v1alpha1.CloudIDSEndpoint, opts v1.UpdateOptions) (*v1alpha1.CloudIDSEndpoint, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cloudidsendpointsResource, "status", c.ns, cloudIDSEndpoint), &v1alpha1.CloudIDSEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudIDSEndpoint), err
}

// Delete takes name of the cloudIDSEndpoint and deletes it. Returns an error if one occurs.
func (c *FakeCloudIDSEndpoints) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(cloudidsendpointsResource, c.ns, name, opts), &v1alpha1.CloudIDSEndpoint{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCloudIDSEndpoints) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cloudidsendpointsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CloudIDSEndpointList{})
	return err
}

// Patch applies the patch and returns the patched cloudIDSEndpoint.
func (c *FakeCloudIDSEndpoints) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CloudIDSEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cloudidsendpointsResource, c.ns, name, pt, data, subresources...), &v1alpha1.CloudIDSEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudIDSEndpoint), err
}
