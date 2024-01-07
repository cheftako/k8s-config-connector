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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/compute/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeComputeTargetTCPProxies implements ComputeTargetTCPProxyInterface
type FakeComputeTargetTCPProxies struct {
	Fake *FakeComputeV1beta1
	ns   string
}

var computetargettcpproxiesResource = v1beta1.SchemeGroupVersion.WithResource("computetargettcpproxies")

var computetargettcpproxiesKind = v1beta1.SchemeGroupVersion.WithKind("ComputeTargetTCPProxy")

// Get takes name of the computeTargetTCPProxy, and returns the corresponding computeTargetTCPProxy object, and an error if there is any.
func (c *FakeComputeTargetTCPProxies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ComputeTargetTCPProxy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(computetargettcpproxiesResource, c.ns, name), &v1beta1.ComputeTargetTCPProxy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeTargetTCPProxy), err
}

// List takes label and field selectors, and returns the list of ComputeTargetTCPProxies that match those selectors.
func (c *FakeComputeTargetTCPProxies) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ComputeTargetTCPProxyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(computetargettcpproxiesResource, computetargettcpproxiesKind, c.ns, opts), &v1beta1.ComputeTargetTCPProxyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ComputeTargetTCPProxyList{ListMeta: obj.(*v1beta1.ComputeTargetTCPProxyList).ListMeta}
	for _, item := range obj.(*v1beta1.ComputeTargetTCPProxyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeTargetTCPProxies.
func (c *FakeComputeTargetTCPProxies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(computetargettcpproxiesResource, c.ns, opts))

}

// Create takes the representation of a computeTargetTCPProxy and creates it.  Returns the server's representation of the computeTargetTCPProxy, and an error, if there is any.
func (c *FakeComputeTargetTCPProxies) Create(ctx context.Context, computeTargetTCPProxy *v1beta1.ComputeTargetTCPProxy, opts v1.CreateOptions) (result *v1beta1.ComputeTargetTCPProxy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(computetargettcpproxiesResource, c.ns, computeTargetTCPProxy), &v1beta1.ComputeTargetTCPProxy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeTargetTCPProxy), err
}

// Update takes the representation of a computeTargetTCPProxy and updates it. Returns the server's representation of the computeTargetTCPProxy, and an error, if there is any.
func (c *FakeComputeTargetTCPProxies) Update(ctx context.Context, computeTargetTCPProxy *v1beta1.ComputeTargetTCPProxy, opts v1.UpdateOptions) (result *v1beta1.ComputeTargetTCPProxy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(computetargettcpproxiesResource, c.ns, computeTargetTCPProxy), &v1beta1.ComputeTargetTCPProxy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeTargetTCPProxy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeTargetTCPProxies) UpdateStatus(ctx context.Context, computeTargetTCPProxy *v1beta1.ComputeTargetTCPProxy, opts v1.UpdateOptions) (*v1beta1.ComputeTargetTCPProxy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(computetargettcpproxiesResource, "status", c.ns, computeTargetTCPProxy), &v1beta1.ComputeTargetTCPProxy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeTargetTCPProxy), err
}

// Delete takes name of the computeTargetTCPProxy and deletes it. Returns an error if one occurs.
func (c *FakeComputeTargetTCPProxies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(computetargettcpproxiesResource, c.ns, name, opts), &v1beta1.ComputeTargetTCPProxy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeTargetTCPProxies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(computetargettcpproxiesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ComputeTargetTCPProxyList{})
	return err
}

// Patch applies the patch and returns the patched computeTargetTCPProxy.
func (c *FakeComputeTargetTCPProxies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ComputeTargetTCPProxy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(computetargettcpproxiesResource, c.ns, name, pt, data, subresources...), &v1beta1.ComputeTargetTCPProxy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeTargetTCPProxy), err
}
