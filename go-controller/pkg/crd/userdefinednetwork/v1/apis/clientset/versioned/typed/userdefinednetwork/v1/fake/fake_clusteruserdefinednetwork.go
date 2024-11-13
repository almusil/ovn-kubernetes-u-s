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

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "github.com/ovn-org/ovn-kubernetes/go-controller/pkg/crd/userdefinednetwork/v1"
	userdefinednetworkv1 "github.com/ovn-org/ovn-kubernetes/go-controller/pkg/crd/userdefinednetwork/v1/apis/applyconfiguration/userdefinednetwork/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterUserDefinedNetworks implements ClusterUserDefinedNetworkInterface
type FakeClusterUserDefinedNetworks struct {
	Fake *FakeK8sV1
}

var clusteruserdefinednetworksResource = v1.SchemeGroupVersion.WithResource("clusteruserdefinednetworks")

var clusteruserdefinednetworksKind = v1.SchemeGroupVersion.WithKind("ClusterUserDefinedNetwork")

// Get takes name of the clusterUserDefinedNetwork, and returns the corresponding clusterUserDefinedNetwork object, and an error if there is any.
func (c *FakeClusterUserDefinedNetworks) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ClusterUserDefinedNetwork, err error) {
	emptyResult := &v1.ClusterUserDefinedNetwork{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(clusteruserdefinednetworksResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ClusterUserDefinedNetwork), err
}

// List takes label and field selectors, and returns the list of ClusterUserDefinedNetworks that match those selectors.
func (c *FakeClusterUserDefinedNetworks) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ClusterUserDefinedNetworkList, err error) {
	emptyResult := &v1.ClusterUserDefinedNetworkList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(clusteruserdefinednetworksResource, clusteruserdefinednetworksKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ClusterUserDefinedNetworkList{ListMeta: obj.(*v1.ClusterUserDefinedNetworkList).ListMeta}
	for _, item := range obj.(*v1.ClusterUserDefinedNetworkList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterUserDefinedNetworks.
func (c *FakeClusterUserDefinedNetworks) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(clusteruserdefinednetworksResource, opts))
}

// Create takes the representation of a clusterUserDefinedNetwork and creates it.  Returns the server's representation of the clusterUserDefinedNetwork, and an error, if there is any.
func (c *FakeClusterUserDefinedNetworks) Create(ctx context.Context, clusterUserDefinedNetwork *v1.ClusterUserDefinedNetwork, opts metav1.CreateOptions) (result *v1.ClusterUserDefinedNetwork, err error) {
	emptyResult := &v1.ClusterUserDefinedNetwork{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(clusteruserdefinednetworksResource, clusterUserDefinedNetwork, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ClusterUserDefinedNetwork), err
}

// Update takes the representation of a clusterUserDefinedNetwork and updates it. Returns the server's representation of the clusterUserDefinedNetwork, and an error, if there is any.
func (c *FakeClusterUserDefinedNetworks) Update(ctx context.Context, clusterUserDefinedNetwork *v1.ClusterUserDefinedNetwork, opts metav1.UpdateOptions) (result *v1.ClusterUserDefinedNetwork, err error) {
	emptyResult := &v1.ClusterUserDefinedNetwork{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(clusteruserdefinednetworksResource, clusterUserDefinedNetwork, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ClusterUserDefinedNetwork), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterUserDefinedNetworks) UpdateStatus(ctx context.Context, clusterUserDefinedNetwork *v1.ClusterUserDefinedNetwork, opts metav1.UpdateOptions) (result *v1.ClusterUserDefinedNetwork, err error) {
	emptyResult := &v1.ClusterUserDefinedNetwork{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceActionWithOptions(clusteruserdefinednetworksResource, "status", clusterUserDefinedNetwork, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ClusterUserDefinedNetwork), err
}

// Delete takes name of the clusterUserDefinedNetwork and deletes it. Returns an error if one occurs.
func (c *FakeClusterUserDefinedNetworks) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(clusteruserdefinednetworksResource, name, opts), &v1.ClusterUserDefinedNetwork{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterUserDefinedNetworks) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(clusteruserdefinednetworksResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.ClusterUserDefinedNetworkList{})
	return err
}

// Patch applies the patch and returns the patched clusterUserDefinedNetwork.
func (c *FakeClusterUserDefinedNetworks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ClusterUserDefinedNetwork, err error) {
	emptyResult := &v1.ClusterUserDefinedNetwork{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(clusteruserdefinednetworksResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ClusterUserDefinedNetwork), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied clusterUserDefinedNetwork.
func (c *FakeClusterUserDefinedNetworks) Apply(ctx context.Context, clusterUserDefinedNetwork *userdefinednetworkv1.ClusterUserDefinedNetworkApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ClusterUserDefinedNetwork, err error) {
	if clusterUserDefinedNetwork == nil {
		return nil, fmt.Errorf("clusterUserDefinedNetwork provided to Apply must not be nil")
	}
	data, err := json.Marshal(clusterUserDefinedNetwork)
	if err != nil {
		return nil, err
	}
	name := clusterUserDefinedNetwork.Name
	if name == nil {
		return nil, fmt.Errorf("clusterUserDefinedNetwork.Name must be provided to Apply")
	}
	emptyResult := &v1.ClusterUserDefinedNetwork{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(clusteruserdefinednetworksResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ClusterUserDefinedNetwork), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeClusterUserDefinedNetworks) ApplyStatus(ctx context.Context, clusterUserDefinedNetwork *userdefinednetworkv1.ClusterUserDefinedNetworkApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ClusterUserDefinedNetwork, err error) {
	if clusterUserDefinedNetwork == nil {
		return nil, fmt.Errorf("clusterUserDefinedNetwork provided to Apply must not be nil")
	}
	data, err := json.Marshal(clusterUserDefinedNetwork)
	if err != nil {
		return nil, err
	}
	name := clusterUserDefinedNetwork.Name
	if name == nil {
		return nil, fmt.Errorf("clusterUserDefinedNetwork.Name must be provided to Apply")
	}
	emptyResult := &v1.ClusterUserDefinedNetwork{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(clusteruserdefinednetworksResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions(), "status"), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ClusterUserDefinedNetwork), err
}