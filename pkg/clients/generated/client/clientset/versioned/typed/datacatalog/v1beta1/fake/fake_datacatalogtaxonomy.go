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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/datacatalog/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDataCatalogTaxonomies implements DataCatalogTaxonomyInterface
type FakeDataCatalogTaxonomies struct {
	Fake *FakeDatacatalogV1beta1
	ns   string
}

var datacatalogtaxonomiesResource = v1beta1.SchemeGroupVersion.WithResource("datacatalogtaxonomies")

var datacatalogtaxonomiesKind = v1beta1.SchemeGroupVersion.WithKind("DataCatalogTaxonomy")

// Get takes name of the dataCatalogTaxonomy, and returns the corresponding dataCatalogTaxonomy object, and an error if there is any.
func (c *FakeDataCatalogTaxonomies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.DataCatalogTaxonomy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(datacatalogtaxonomiesResource, c.ns, name), &v1beta1.DataCatalogTaxonomy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DataCatalogTaxonomy), err
}

// List takes label and field selectors, and returns the list of DataCatalogTaxonomies that match those selectors.
func (c *FakeDataCatalogTaxonomies) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.DataCatalogTaxonomyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(datacatalogtaxonomiesResource, datacatalogtaxonomiesKind, c.ns, opts), &v1beta1.DataCatalogTaxonomyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.DataCatalogTaxonomyList{ListMeta: obj.(*v1beta1.DataCatalogTaxonomyList).ListMeta}
	for _, item := range obj.(*v1beta1.DataCatalogTaxonomyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dataCatalogTaxonomies.
func (c *FakeDataCatalogTaxonomies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(datacatalogtaxonomiesResource, c.ns, opts))

}

// Create takes the representation of a dataCatalogTaxonomy and creates it.  Returns the server's representation of the dataCatalogTaxonomy, and an error, if there is any.
func (c *FakeDataCatalogTaxonomies) Create(ctx context.Context, dataCatalogTaxonomy *v1beta1.DataCatalogTaxonomy, opts v1.CreateOptions) (result *v1beta1.DataCatalogTaxonomy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(datacatalogtaxonomiesResource, c.ns, dataCatalogTaxonomy), &v1beta1.DataCatalogTaxonomy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DataCatalogTaxonomy), err
}

// Update takes the representation of a dataCatalogTaxonomy and updates it. Returns the server's representation of the dataCatalogTaxonomy, and an error, if there is any.
func (c *FakeDataCatalogTaxonomies) Update(ctx context.Context, dataCatalogTaxonomy *v1beta1.DataCatalogTaxonomy, opts v1.UpdateOptions) (result *v1beta1.DataCatalogTaxonomy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(datacatalogtaxonomiesResource, c.ns, dataCatalogTaxonomy), &v1beta1.DataCatalogTaxonomy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DataCatalogTaxonomy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDataCatalogTaxonomies) UpdateStatus(ctx context.Context, dataCatalogTaxonomy *v1beta1.DataCatalogTaxonomy, opts v1.UpdateOptions) (*v1beta1.DataCatalogTaxonomy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(datacatalogtaxonomiesResource, "status", c.ns, dataCatalogTaxonomy), &v1beta1.DataCatalogTaxonomy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DataCatalogTaxonomy), err
}

// Delete takes name of the dataCatalogTaxonomy and deletes it. Returns an error if one occurs.
func (c *FakeDataCatalogTaxonomies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(datacatalogtaxonomiesResource, c.ns, name, opts), &v1beta1.DataCatalogTaxonomy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDataCatalogTaxonomies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(datacatalogtaxonomiesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.DataCatalogTaxonomyList{})
	return err
}

// Patch applies the patch and returns the patched dataCatalogTaxonomy.
func (c *FakeDataCatalogTaxonomies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.DataCatalogTaxonomy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(datacatalogtaxonomiesResource, c.ns, name, pt, data, subresources...), &v1beta1.DataCatalogTaxonomy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.DataCatalogTaxonomy), err
}
