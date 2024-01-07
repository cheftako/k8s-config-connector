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

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/healthcare/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeHealthcareFHIRStores implements HealthcareFHIRStoreInterface
type FakeHealthcareFHIRStores struct {
	Fake *FakeHealthcareV1alpha1
	ns   string
}

var healthcarefhirstoresResource = v1alpha1.SchemeGroupVersion.WithResource("healthcarefhirstores")

var healthcarefhirstoresKind = v1alpha1.SchemeGroupVersion.WithKind("HealthcareFHIRStore")

// Get takes name of the healthcareFHIRStore, and returns the corresponding healthcareFHIRStore object, and an error if there is any.
func (c *FakeHealthcareFHIRStores) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.HealthcareFHIRStore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(healthcarefhirstoresResource, c.ns, name), &v1alpha1.HealthcareFHIRStore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HealthcareFHIRStore), err
}

// List takes label and field selectors, and returns the list of HealthcareFHIRStores that match those selectors.
func (c *FakeHealthcareFHIRStores) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.HealthcareFHIRStoreList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(healthcarefhirstoresResource, healthcarefhirstoresKind, c.ns, opts), &v1alpha1.HealthcareFHIRStoreList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.HealthcareFHIRStoreList{ListMeta: obj.(*v1alpha1.HealthcareFHIRStoreList).ListMeta}
	for _, item := range obj.(*v1alpha1.HealthcareFHIRStoreList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested healthcareFHIRStores.
func (c *FakeHealthcareFHIRStores) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(healthcarefhirstoresResource, c.ns, opts))

}

// Create takes the representation of a healthcareFHIRStore and creates it.  Returns the server's representation of the healthcareFHIRStore, and an error, if there is any.
func (c *FakeHealthcareFHIRStores) Create(ctx context.Context, healthcareFHIRStore *v1alpha1.HealthcareFHIRStore, opts v1.CreateOptions) (result *v1alpha1.HealthcareFHIRStore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(healthcarefhirstoresResource, c.ns, healthcareFHIRStore), &v1alpha1.HealthcareFHIRStore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HealthcareFHIRStore), err
}

// Update takes the representation of a healthcareFHIRStore and updates it. Returns the server's representation of the healthcareFHIRStore, and an error, if there is any.
func (c *FakeHealthcareFHIRStores) Update(ctx context.Context, healthcareFHIRStore *v1alpha1.HealthcareFHIRStore, opts v1.UpdateOptions) (result *v1alpha1.HealthcareFHIRStore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(healthcarefhirstoresResource, c.ns, healthcareFHIRStore), &v1alpha1.HealthcareFHIRStore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HealthcareFHIRStore), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHealthcareFHIRStores) UpdateStatus(ctx context.Context, healthcareFHIRStore *v1alpha1.HealthcareFHIRStore, opts v1.UpdateOptions) (*v1alpha1.HealthcareFHIRStore, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(healthcarefhirstoresResource, "status", c.ns, healthcareFHIRStore), &v1alpha1.HealthcareFHIRStore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HealthcareFHIRStore), err
}

// Delete takes name of the healthcareFHIRStore and deletes it. Returns an error if one occurs.
func (c *FakeHealthcareFHIRStores) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(healthcarefhirstoresResource, c.ns, name, opts), &v1alpha1.HealthcareFHIRStore{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHealthcareFHIRStores) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(healthcarefhirstoresResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.HealthcareFHIRStoreList{})
	return err
}

// Patch applies the patch and returns the patched healthcareFHIRStore.
func (c *FakeHealthcareFHIRStores) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.HealthcareFHIRStore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(healthcarefhirstoresResource, c.ns, name, pt, data, subresources...), &v1alpha1.HealthcareFHIRStore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HealthcareFHIRStore), err
}
