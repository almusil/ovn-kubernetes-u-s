/*


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

package v1

import (
	"context"
	"time"

	v1 "github.com/ovn-org/ovn-kubernetes/go-controller/pkg/crd/adminpolicybasedroute/v1"
	scheme "github.com/ovn-org/ovn-kubernetes/go-controller/pkg/crd/adminpolicybasedroute/v1/apis/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AdminPolicyBasedExternalRoutesGetter has a method to return a AdminPolicyBasedExternalRouteInterface.
// A group's client should implement this interface.
type AdminPolicyBasedExternalRoutesGetter interface {
	AdminPolicyBasedExternalRoutes() AdminPolicyBasedExternalRouteInterface
}

// AdminPolicyBasedExternalRouteInterface has methods to work with AdminPolicyBasedExternalRoute resources.
type AdminPolicyBasedExternalRouteInterface interface {
	Create(ctx context.Context, adminPolicyBasedExternalRoute *v1.AdminPolicyBasedExternalRoute, opts metav1.CreateOptions) (*v1.AdminPolicyBasedExternalRoute, error)
	Update(ctx context.Context, adminPolicyBasedExternalRoute *v1.AdminPolicyBasedExternalRoute, opts metav1.UpdateOptions) (*v1.AdminPolicyBasedExternalRoute, error)
	UpdateStatus(ctx context.Context, adminPolicyBasedExternalRoute *v1.AdminPolicyBasedExternalRoute, opts metav1.UpdateOptions) (*v1.AdminPolicyBasedExternalRoute, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.AdminPolicyBasedExternalRoute, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.AdminPolicyBasedExternalRouteList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.AdminPolicyBasedExternalRoute, err error)
	AdminPolicyBasedExternalRouteExpansion
}

// adminPolicyBasedExternalRoutes implements AdminPolicyBasedExternalRouteInterface
type adminPolicyBasedExternalRoutes struct {
	client rest.Interface
}

// newAdminPolicyBasedExternalRoutes returns a AdminPolicyBasedExternalRoutes
func newAdminPolicyBasedExternalRoutes(c *K8sV1Client) *adminPolicyBasedExternalRoutes {
	return &adminPolicyBasedExternalRoutes{
		client: c.RESTClient(),
	}
}

// Get takes name of the adminPolicyBasedExternalRoute, and returns the corresponding adminPolicyBasedExternalRoute object, and an error if there is any.
func (c *adminPolicyBasedExternalRoutes) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.AdminPolicyBasedExternalRoute, err error) {
	result = &v1.AdminPolicyBasedExternalRoute{}
	err = c.client.Get().
		Resource("adminpolicybasedexternalroutes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AdminPolicyBasedExternalRoutes that match those selectors.
func (c *adminPolicyBasedExternalRoutes) List(ctx context.Context, opts metav1.ListOptions) (result *v1.AdminPolicyBasedExternalRouteList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.AdminPolicyBasedExternalRouteList{}
	err = c.client.Get().
		Resource("adminpolicybasedexternalroutes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested adminPolicyBasedExternalRoutes.
func (c *adminPolicyBasedExternalRoutes) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("adminpolicybasedexternalroutes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a adminPolicyBasedExternalRoute and creates it.  Returns the server's representation of the adminPolicyBasedExternalRoute, and an error, if there is any.
func (c *adminPolicyBasedExternalRoutes) Create(ctx context.Context, adminPolicyBasedExternalRoute *v1.AdminPolicyBasedExternalRoute, opts metav1.CreateOptions) (result *v1.AdminPolicyBasedExternalRoute, err error) {
	result = &v1.AdminPolicyBasedExternalRoute{}
	err = c.client.Post().
		Resource("adminpolicybasedexternalroutes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(adminPolicyBasedExternalRoute).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a adminPolicyBasedExternalRoute and updates it. Returns the server's representation of the adminPolicyBasedExternalRoute, and an error, if there is any.
func (c *adminPolicyBasedExternalRoutes) Update(ctx context.Context, adminPolicyBasedExternalRoute *v1.AdminPolicyBasedExternalRoute, opts metav1.UpdateOptions) (result *v1.AdminPolicyBasedExternalRoute, err error) {
	result = &v1.AdminPolicyBasedExternalRoute{}
	err = c.client.Put().
		Resource("adminpolicybasedexternalroutes").
		Name(adminPolicyBasedExternalRoute.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(adminPolicyBasedExternalRoute).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *adminPolicyBasedExternalRoutes) UpdateStatus(ctx context.Context, adminPolicyBasedExternalRoute *v1.AdminPolicyBasedExternalRoute, opts metav1.UpdateOptions) (result *v1.AdminPolicyBasedExternalRoute, err error) {
	result = &v1.AdminPolicyBasedExternalRoute{}
	err = c.client.Put().
		Resource("adminpolicybasedexternalroutes").
		Name(adminPolicyBasedExternalRoute.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(adminPolicyBasedExternalRoute).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the adminPolicyBasedExternalRoute and deletes it. Returns an error if one occurs.
func (c *adminPolicyBasedExternalRoutes) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("adminpolicybasedexternalroutes").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *adminPolicyBasedExternalRoutes) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("adminpolicybasedexternalroutes").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched adminPolicyBasedExternalRoute.
func (c *adminPolicyBasedExternalRoutes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.AdminPolicyBasedExternalRoute, err error) {
	result = &v1.AdminPolicyBasedExternalRoute{}
	err = c.client.Patch(pt).
		Resource("adminpolicybasedexternalroutes").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}