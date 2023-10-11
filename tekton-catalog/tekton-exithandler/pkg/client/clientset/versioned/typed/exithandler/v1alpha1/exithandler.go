// Copyright 2023 kubeflow.org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/kubeflow/kfp-tekton/tekton-catalog/tekton-exithandler/pkg/apis/exithandler/v1alpha1"
	scheme "github.com/kubeflow/kfp-tekton/tekton-catalog/tekton-exithandler/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ExitHandlersGetter has a method to return a ExitHandlerInterface.
// A group's client should implement this interface.
type ExitHandlersGetter interface {
	ExitHandlers(namespace string) ExitHandlerInterface
}

// ExitHandlerInterface has methods to work with ExitHandler resources.
type ExitHandlerInterface interface {
	Create(ctx context.Context, exitHandler *v1alpha1.ExitHandler, opts v1.CreateOptions) (*v1alpha1.ExitHandler, error)
	Update(ctx context.Context, exitHandler *v1alpha1.ExitHandler, opts v1.UpdateOptions) (*v1alpha1.ExitHandler, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ExitHandler, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ExitHandlerList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ExitHandler, err error)
	ExitHandlerExpansion
}

// exitHandlers implements ExitHandlerInterface
type exitHandlers struct {
	client rest.Interface
	ns     string
}

// newExitHandlers returns a ExitHandlers
func newExitHandlers(c *CustomV1alpha1Client, namespace string) *exitHandlers {
	return &exitHandlers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the exitHandler, and returns the corresponding exitHandler object, and an error if there is any.
func (c *exitHandlers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ExitHandler, err error) {
	result = &v1alpha1.ExitHandler{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("exithandlers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ExitHandlers that match those selectors.
func (c *exitHandlers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ExitHandlerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ExitHandlerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("exithandlers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested exitHandlers.
func (c *exitHandlers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("exithandlers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a exitHandler and creates it.  Returns the server's representation of the exitHandler, and an error, if there is any.
func (c *exitHandlers) Create(ctx context.Context, exitHandler *v1alpha1.ExitHandler, opts v1.CreateOptions) (result *v1alpha1.ExitHandler, err error) {
	result = &v1alpha1.ExitHandler{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("exithandlers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(exitHandler).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a exitHandler and updates it. Returns the server's representation of the exitHandler, and an error, if there is any.
func (c *exitHandlers) Update(ctx context.Context, exitHandler *v1alpha1.ExitHandler, opts v1.UpdateOptions) (result *v1alpha1.ExitHandler, err error) {
	result = &v1alpha1.ExitHandler{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("exithandlers").
		Name(exitHandler.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(exitHandler).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the exitHandler and deletes it. Returns an error if one occurs.
func (c *exitHandlers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("exithandlers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *exitHandlers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("exithandlers").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched exitHandler.
func (c *exitHandlers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ExitHandler, err error) {
	result = &v1alpha1.ExitHandler{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("exithandlers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}