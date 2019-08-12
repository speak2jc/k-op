/*
Copyright The Kubernetes Authors.

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

package v1alpha1

import (
	scheme "speak2jc/k-op/pkg/generated/clientset/versioned/scheme"
	"time"

	v1alpha1 "github.com/speak2jc/k-op/pkg/apis/example/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// KeevakindListsGetter has a method to return a KeevakindListInterface.
// A group's client should implement this interface.
type KeevakindListsGetter interface {
	KeevakindLists(namespace string) KeevakindListInterface
}

// KeevakindListInterface has methods to work with KeevakindList resources.
type KeevakindListInterface interface {
	Create(*v1alpha1.KeevakindList) (*v1alpha1.KeevakindList, error)
	Update(*v1alpha1.KeevakindList) (*v1alpha1.KeevakindList, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.KeevakindList, error)
	List(opts v1.ListOptions) (*v1alpha1.KeevakindListList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KeevakindList, err error)
	KeevakindListExpansion
}

// keevakindLists implements KeevakindListInterface
type keevakindLists struct {
	client rest.Interface
	ns     string
}

// newKeevakindLists returns a KeevakindLists
func newKeevakindLists(c *ExampleV1alpha1Client, namespace string) *keevakindLists {
	return &keevakindLists{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the keevakindList, and returns the corresponding keevakindList object, and an error if there is any.
func (c *keevakindLists) Get(name string, options v1.GetOptions) (result *v1alpha1.KeevakindList, err error) {
	result = &v1alpha1.KeevakindList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("keevakindlists").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KeevakindLists that match those selectors.
func (c *keevakindLists) List(opts v1.ListOptions) (result *v1alpha1.KeevakindListList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.KeevakindListList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("keevakindlists").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested keevakindLists.
func (c *keevakindLists) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("keevakindlists").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a keevakindList and creates it.  Returns the server's representation of the keevakindList, and an error, if there is any.
func (c *keevakindLists) Create(keevakindList *v1alpha1.KeevakindList) (result *v1alpha1.KeevakindList, err error) {
	result = &v1alpha1.KeevakindList{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("keevakindlists").
		Body(keevakindList).
		Do().
		Into(result)
	return
}

// Update takes the representation of a keevakindList and updates it. Returns the server's representation of the keevakindList, and an error, if there is any.
func (c *keevakindLists) Update(keevakindList *v1alpha1.KeevakindList) (result *v1alpha1.KeevakindList, err error) {
	result = &v1alpha1.KeevakindList{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("keevakindlists").
		Name(keevakindList.Name).
		Body(keevakindList).
		Do().
		Into(result)
	return
}

// Delete takes name of the keevakindList and deletes it. Returns an error if one occurs.
func (c *keevakindLists) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("keevakindlists").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *keevakindLists) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("keevakindlists").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched keevakindList.
func (c *keevakindLists) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.KeevakindList, err error) {
	result = &v1alpha1.KeevakindList{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("keevakindlists").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
