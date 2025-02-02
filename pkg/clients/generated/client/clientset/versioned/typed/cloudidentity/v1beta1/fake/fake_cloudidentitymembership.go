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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/cloudidentity/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCloudIdentityMemberships implements CloudIdentityMembershipInterface
type FakeCloudIdentityMemberships struct {
	Fake *FakeCloudidentityV1beta1
	ns   string
}

var cloudidentitymembershipsResource = schema.GroupVersionResource{Group: "cloudidentity.cnrm.cloud.google.com", Version: "v1beta1", Resource: "cloudidentitymemberships"}

var cloudidentitymembershipsKind = schema.GroupVersionKind{Group: "cloudidentity.cnrm.cloud.google.com", Version: "v1beta1", Kind: "CloudIdentityMembership"}

// Get takes name of the cloudIdentityMembership, and returns the corresponding cloudIdentityMembership object, and an error if there is any.
func (c *FakeCloudIdentityMemberships) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.CloudIdentityMembership, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cloudidentitymembershipsResource, c.ns, name), &v1beta1.CloudIdentityMembership{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CloudIdentityMembership), err
}

// List takes label and field selectors, and returns the list of CloudIdentityMemberships that match those selectors.
func (c *FakeCloudIdentityMemberships) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.CloudIdentityMembershipList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cloudidentitymembershipsResource, cloudidentitymembershipsKind, c.ns, opts), &v1beta1.CloudIdentityMembershipList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.CloudIdentityMembershipList{ListMeta: obj.(*v1beta1.CloudIdentityMembershipList).ListMeta}
	for _, item := range obj.(*v1beta1.CloudIdentityMembershipList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cloudIdentityMemberships.
func (c *FakeCloudIdentityMemberships) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cloudidentitymembershipsResource, c.ns, opts))

}

// Create takes the representation of a cloudIdentityMembership and creates it.  Returns the server's representation of the cloudIdentityMembership, and an error, if there is any.
func (c *FakeCloudIdentityMemberships) Create(ctx context.Context, cloudIdentityMembership *v1beta1.CloudIdentityMembership, opts v1.CreateOptions) (result *v1beta1.CloudIdentityMembership, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cloudidentitymembershipsResource, c.ns, cloudIdentityMembership), &v1beta1.CloudIdentityMembership{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CloudIdentityMembership), err
}

// Update takes the representation of a cloudIdentityMembership and updates it. Returns the server's representation of the cloudIdentityMembership, and an error, if there is any.
func (c *FakeCloudIdentityMemberships) Update(ctx context.Context, cloudIdentityMembership *v1beta1.CloudIdentityMembership, opts v1.UpdateOptions) (result *v1beta1.CloudIdentityMembership, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cloudidentitymembershipsResource, c.ns, cloudIdentityMembership), &v1beta1.CloudIdentityMembership{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CloudIdentityMembership), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCloudIdentityMemberships) UpdateStatus(ctx context.Context, cloudIdentityMembership *v1beta1.CloudIdentityMembership, opts v1.UpdateOptions) (*v1beta1.CloudIdentityMembership, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cloudidentitymembershipsResource, "status", c.ns, cloudIdentityMembership), &v1beta1.CloudIdentityMembership{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CloudIdentityMembership), err
}

// Delete takes name of the cloudIdentityMembership and deletes it. Returns an error if one occurs.
func (c *FakeCloudIdentityMemberships) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(cloudidentitymembershipsResource, c.ns, name, opts), &v1beta1.CloudIdentityMembership{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCloudIdentityMemberships) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cloudidentitymembershipsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.CloudIdentityMembershipList{})
	return err
}

// Patch applies the patch and returns the patched cloudIdentityMembership.
func (c *FakeCloudIdentityMemberships) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.CloudIdentityMembership, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cloudidentitymembershipsResource, c.ns, name, pt, data, subresources...), &v1beta1.CloudIdentityMembership{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CloudIdentityMembership), err
}
