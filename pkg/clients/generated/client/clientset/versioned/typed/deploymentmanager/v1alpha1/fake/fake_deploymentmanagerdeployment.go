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

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/deploymentmanager/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDeploymentManagerDeployments implements DeploymentManagerDeploymentInterface
type FakeDeploymentManagerDeployments struct {
	Fake *FakeDeploymentmanagerV1alpha1
	ns   string
}

var deploymentmanagerdeploymentsResource = v1alpha1.SchemeGroupVersion.WithResource("deploymentmanagerdeployments")

var deploymentmanagerdeploymentsKind = v1alpha1.SchemeGroupVersion.WithKind("DeploymentManagerDeployment")

// Get takes name of the deploymentManagerDeployment, and returns the corresponding deploymentManagerDeployment object, and an error if there is any.
func (c *FakeDeploymentManagerDeployments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DeploymentManagerDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(deploymentmanagerdeploymentsResource, c.ns, name), &v1alpha1.DeploymentManagerDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DeploymentManagerDeployment), err
}

// List takes label and field selectors, and returns the list of DeploymentManagerDeployments that match those selectors.
func (c *FakeDeploymentManagerDeployments) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DeploymentManagerDeploymentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(deploymentmanagerdeploymentsResource, deploymentmanagerdeploymentsKind, c.ns, opts), &v1alpha1.DeploymentManagerDeploymentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DeploymentManagerDeploymentList{ListMeta: obj.(*v1alpha1.DeploymentManagerDeploymentList).ListMeta}
	for _, item := range obj.(*v1alpha1.DeploymentManagerDeploymentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested deploymentManagerDeployments.
func (c *FakeDeploymentManagerDeployments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(deploymentmanagerdeploymentsResource, c.ns, opts))

}

// Create takes the representation of a deploymentManagerDeployment and creates it.  Returns the server's representation of the deploymentManagerDeployment, and an error, if there is any.
func (c *FakeDeploymentManagerDeployments) Create(ctx context.Context, deploymentManagerDeployment *v1alpha1.DeploymentManagerDeployment, opts v1.CreateOptions) (result *v1alpha1.DeploymentManagerDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(deploymentmanagerdeploymentsResource, c.ns, deploymentManagerDeployment), &v1alpha1.DeploymentManagerDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DeploymentManagerDeployment), err
}

// Update takes the representation of a deploymentManagerDeployment and updates it. Returns the server's representation of the deploymentManagerDeployment, and an error, if there is any.
func (c *FakeDeploymentManagerDeployments) Update(ctx context.Context, deploymentManagerDeployment *v1alpha1.DeploymentManagerDeployment, opts v1.UpdateOptions) (result *v1alpha1.DeploymentManagerDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(deploymentmanagerdeploymentsResource, c.ns, deploymentManagerDeployment), &v1alpha1.DeploymentManagerDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DeploymentManagerDeployment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDeploymentManagerDeployments) UpdateStatus(ctx context.Context, deploymentManagerDeployment *v1alpha1.DeploymentManagerDeployment, opts v1.UpdateOptions) (*v1alpha1.DeploymentManagerDeployment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(deploymentmanagerdeploymentsResource, "status", c.ns, deploymentManagerDeployment), &v1alpha1.DeploymentManagerDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DeploymentManagerDeployment), err
}

// Delete takes name of the deploymentManagerDeployment and deletes it. Returns an error if one occurs.
func (c *FakeDeploymentManagerDeployments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(deploymentmanagerdeploymentsResource, c.ns, name, opts), &v1alpha1.DeploymentManagerDeployment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDeploymentManagerDeployments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(deploymentmanagerdeploymentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DeploymentManagerDeploymentList{})
	return err
}

// Patch applies the patch and returns the patched deploymentManagerDeployment.
func (c *FakeDeploymentManagerDeployments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DeploymentManagerDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(deploymentmanagerdeploymentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DeploymentManagerDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DeploymentManagerDeployment), err
}
