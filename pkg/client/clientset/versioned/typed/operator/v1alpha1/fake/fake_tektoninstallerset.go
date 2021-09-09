/*
Copyright 2020 The Tekton Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/tektoncd/operator/pkg/apis/operator/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTektonInstallerSets implements TektonInstallerSetInterface
type FakeTektonInstallerSets struct {
	Fake *FakeOperatorV1alpha1
}

var tektoninstallersetsResource = schema.GroupVersionResource{Group: "operator.tekton.dev", Version: "v1alpha1", Resource: "tektoninstallersets"}

var tektoninstallersetsKind = schema.GroupVersionKind{Group: "operator.tekton.dev", Version: "v1alpha1", Kind: "TektonInstallerSet"}

// Get takes name of the tektonInstallerSet, and returns the corresponding tektonInstallerSet object, and an error if there is any.
func (c *FakeTektonInstallerSets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.TektonInstallerSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(tektoninstallersetsResource, name), &v1alpha1.TektonInstallerSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TektonInstallerSet), err
}

// List takes label and field selectors, and returns the list of TektonInstallerSets that match those selectors.
func (c *FakeTektonInstallerSets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TektonInstallerSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(tektoninstallersetsResource, tektoninstallersetsKind, opts), &v1alpha1.TektonInstallerSetList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TektonInstallerSetList{ListMeta: obj.(*v1alpha1.TektonInstallerSetList).ListMeta}
	for _, item := range obj.(*v1alpha1.TektonInstallerSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tektonInstallerSets.
func (c *FakeTektonInstallerSets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(tektoninstallersetsResource, opts))
}

// Create takes the representation of a tektonInstallerSet and creates it.  Returns the server's representation of the tektonInstallerSet, and an error, if there is any.
func (c *FakeTektonInstallerSets) Create(ctx context.Context, tektonInstallerSet *v1alpha1.TektonInstallerSet, opts v1.CreateOptions) (result *v1alpha1.TektonInstallerSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(tektoninstallersetsResource, tektonInstallerSet), &v1alpha1.TektonInstallerSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TektonInstallerSet), err
}

// Update takes the representation of a tektonInstallerSet and updates it. Returns the server's representation of the tektonInstallerSet, and an error, if there is any.
func (c *FakeTektonInstallerSets) Update(ctx context.Context, tektonInstallerSet *v1alpha1.TektonInstallerSet, opts v1.UpdateOptions) (result *v1alpha1.TektonInstallerSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(tektoninstallersetsResource, tektonInstallerSet), &v1alpha1.TektonInstallerSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TektonInstallerSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTektonInstallerSets) UpdateStatus(ctx context.Context, tektonInstallerSet *v1alpha1.TektonInstallerSet, opts v1.UpdateOptions) (*v1alpha1.TektonInstallerSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(tektoninstallersetsResource, "status", tektonInstallerSet), &v1alpha1.TektonInstallerSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TektonInstallerSet), err
}

// Delete takes name of the tektonInstallerSet and deletes it. Returns an error if one occurs.
func (c *FakeTektonInstallerSets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(tektoninstallersetsResource, name), &v1alpha1.TektonInstallerSet{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTektonInstallerSets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(tektoninstallersetsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.TektonInstallerSetList{})
	return err
}

// Patch applies the patch and returns the patched tektonInstallerSet.
func (c *FakeTektonInstallerSets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TektonInstallerSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(tektoninstallersetsResource, name, pt, data, subresources...), &v1alpha1.TektonInstallerSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TektonInstallerSet), err
}