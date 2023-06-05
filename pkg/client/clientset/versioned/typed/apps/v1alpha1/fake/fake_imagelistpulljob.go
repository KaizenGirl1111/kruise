/*
Copyright 2021 The Kruise Authors.

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

	v1alpha1 "github.com/openkruise/kruise/apis/apps/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeImageListPullJobs implements ImageListPullJobInterface
type FakeImageListPullJobs struct {
	Fake *FakeAppsV1alpha1
	ns   string
}

var imagelistpulljobsResource = schema.GroupVersionResource{Group: "apps.kruise.io", Version: "v1alpha1", Resource: "imagelistpulljobs"}

var imagelistpulljobsKind = schema.GroupVersionKind{Group: "apps.kruise.io", Version: "v1alpha1", Kind: "ImageListPullJob"}

// Get takes name of the imageListPullJob, and returns the corresponding imageListPullJob object, and an error if there is any.
func (c *FakeImageListPullJobs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ImageListPullJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(imagelistpulljobsResource, c.ns, name), &v1alpha1.ImageListPullJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ImageListPullJob), err
}

// List takes label and field selectors, and returns the list of ImageListPullJobs that match those selectors.
func (c *FakeImageListPullJobs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ImageListPullJobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(imagelistpulljobsResource, imagelistpulljobsKind, c.ns, opts), &v1alpha1.ImageListPullJobList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ImageListPullJobList{ListMeta: obj.(*v1alpha1.ImageListPullJobList).ListMeta}
	for _, item := range obj.(*v1alpha1.ImageListPullJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested imageListPullJobs.
func (c *FakeImageListPullJobs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(imagelistpulljobsResource, c.ns, opts))

}

// Create takes the representation of a imageListPullJob and creates it.  Returns the server's representation of the imageListPullJob, and an error, if there is any.
func (c *FakeImageListPullJobs) Create(ctx context.Context, imageListPullJob *v1alpha1.ImageListPullJob, opts v1.CreateOptions) (result *v1alpha1.ImageListPullJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(imagelistpulljobsResource, c.ns, imageListPullJob), &v1alpha1.ImageListPullJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ImageListPullJob), err
}

// Update takes the representation of a imageListPullJob and updates it. Returns the server's representation of the imageListPullJob, and an error, if there is any.
func (c *FakeImageListPullJobs) Update(ctx context.Context, imageListPullJob *v1alpha1.ImageListPullJob, opts v1.UpdateOptions) (result *v1alpha1.ImageListPullJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(imagelistpulljobsResource, c.ns, imageListPullJob), &v1alpha1.ImageListPullJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ImageListPullJob), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeImageListPullJobs) UpdateStatus(ctx context.Context, imageListPullJob *v1alpha1.ImageListPullJob, opts v1.UpdateOptions) (*v1alpha1.ImageListPullJob, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(imagelistpulljobsResource, "status", c.ns, imageListPullJob), &v1alpha1.ImageListPullJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ImageListPullJob), err
}

// Delete takes name of the imageListPullJob and deletes it. Returns an error if one occurs.
func (c *FakeImageListPullJobs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(imagelistpulljobsResource, c.ns, name), &v1alpha1.ImageListPullJob{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeImageListPullJobs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(imagelistpulljobsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ImageListPullJobList{})
	return err
}

// Patch applies the patch and returns the patched imageListPullJob.
func (c *FakeImageListPullJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ImageListPullJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(imagelistpulljobsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ImageListPullJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ImageListPullJob), err
}