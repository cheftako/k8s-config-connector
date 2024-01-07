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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/networkservices/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNetworkServicesTCPRoutes implements NetworkServicesTCPRouteInterface
type FakeNetworkServicesTCPRoutes struct {
	Fake *FakeNetworkservicesV1beta1
	ns   string
}

var networkservicestcproutesResource = v1beta1.SchemeGroupVersion.WithResource("networkservicestcproutes")

var networkservicestcproutesKind = v1beta1.SchemeGroupVersion.WithKind("NetworkServicesTCPRoute")

// Get takes name of the networkServicesTCPRoute, and returns the corresponding networkServicesTCPRoute object, and an error if there is any.
func (c *FakeNetworkServicesTCPRoutes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.NetworkServicesTCPRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networkservicestcproutesResource, c.ns, name), &v1beta1.NetworkServicesTCPRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkServicesTCPRoute), err
}

// List takes label and field selectors, and returns the list of NetworkServicesTCPRoutes that match those selectors.
func (c *FakeNetworkServicesTCPRoutes) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.NetworkServicesTCPRouteList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(networkservicestcproutesResource, networkservicestcproutesKind, c.ns, opts), &v1beta1.NetworkServicesTCPRouteList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.NetworkServicesTCPRouteList{ListMeta: obj.(*v1beta1.NetworkServicesTCPRouteList).ListMeta}
	for _, item := range obj.(*v1beta1.NetworkServicesTCPRouteList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkServicesTCPRoutes.
func (c *FakeNetworkServicesTCPRoutes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networkservicestcproutesResource, c.ns, opts))

}

// Create takes the representation of a networkServicesTCPRoute and creates it.  Returns the server's representation of the networkServicesTCPRoute, and an error, if there is any.
func (c *FakeNetworkServicesTCPRoutes) Create(ctx context.Context, networkServicesTCPRoute *v1beta1.NetworkServicesTCPRoute, opts v1.CreateOptions) (result *v1beta1.NetworkServicesTCPRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(networkservicestcproutesResource, c.ns, networkServicesTCPRoute), &v1beta1.NetworkServicesTCPRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkServicesTCPRoute), err
}

// Update takes the representation of a networkServicesTCPRoute and updates it. Returns the server's representation of the networkServicesTCPRoute, and an error, if there is any.
func (c *FakeNetworkServicesTCPRoutes) Update(ctx context.Context, networkServicesTCPRoute *v1beta1.NetworkServicesTCPRoute, opts v1.UpdateOptions) (result *v1beta1.NetworkServicesTCPRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(networkservicestcproutesResource, c.ns, networkServicesTCPRoute), &v1beta1.NetworkServicesTCPRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkServicesTCPRoute), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetworkServicesTCPRoutes) UpdateStatus(ctx context.Context, networkServicesTCPRoute *v1beta1.NetworkServicesTCPRoute, opts v1.UpdateOptions) (*v1beta1.NetworkServicesTCPRoute, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(networkservicestcproutesResource, "status", c.ns, networkServicesTCPRoute), &v1beta1.NetworkServicesTCPRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkServicesTCPRoute), err
}

// Delete takes name of the networkServicesTCPRoute and deletes it. Returns an error if one occurs.
func (c *FakeNetworkServicesTCPRoutes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(networkservicestcproutesResource, c.ns, name, opts), &v1beta1.NetworkServicesTCPRoute{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkServicesTCPRoutes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(networkservicestcproutesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.NetworkServicesTCPRouteList{})
	return err
}

// Patch applies the patch and returns the patched networkServicesTCPRoute.
func (c *FakeNetworkServicesTCPRoutes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.NetworkServicesTCPRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(networkservicestcproutesResource, c.ns, name, pt, data, subresources...), &v1beta1.NetworkServicesTCPRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.NetworkServicesTCPRoute), err
}
