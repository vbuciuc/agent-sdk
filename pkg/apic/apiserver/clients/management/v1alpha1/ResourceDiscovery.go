/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"fmt"

	v1 "github.com/Axway/axway-sdk/pkg/apic/apiserver/clients/api/v1"
	apiv1 "github.com/Axway/axway-sdk/pkg/apic/apiserver/models/api/v1"
	"github.com/Axway/axway-sdk/pkg/apic/apiserver/models/management/v1alpha1"
)

type ResourceDiscoveryMergeFunc func(*v1alpha1.ResourceDiscovery, *v1alpha1.ResourceDiscovery) (*v1alpha1.ResourceDiscovery, error)

// Merge builds a merge option for an update operation
func ResourceDiscoveryMerge(f ResourceDiscoveryMergeFunc) v1.UpdateOption {
	return v1.Merge(func(prev, new apiv1.Interface) (apiv1.Interface, error) {
		p, n := &v1alpha1.ResourceDiscovery{}, &v1alpha1.ResourceDiscovery{}

		switch t := prev.(type) {
		case *v1alpha1.ResourceDiscovery:
			p = t
		case *apiv1.ResourceInstance:
			err := p.FromInstance(t)
			if err != nil {
				return nil, fmt.Errorf("merge: failed to unserialise prev resource: %w", err)
			}
		default:
			return nil, fmt.Errorf("merge: failed to unserialise prev resource, unxexpected resource type: %T", t)
		}

		switch t := new.(type) {
		case *v1alpha1.ResourceDiscovery:
			n = t
		case *apiv1.ResourceInstance:
			err := n.FromInstance(t)
			if err != nil {
				return nil, fmt.Errorf("merge: failed to unserialize new resource: %w", err)
			}
		default:
			return nil, fmt.Errorf("merge: failed to unserialise new resource, unxexpected resource type: %T", t)
		}

		return f(p, n)
	})
}

// ResourceDiscoveryClient -
type ResourceDiscoveryClient struct {
	client v1.Scoped
}

// UnscopedResourceDiscoveryClient -
type UnscopedResourceDiscoveryClient struct {
	client v1.Unscoped
}

// NewResourceDiscoveryClient -
func NewResourceDiscoveryClient(c v1.Base) (*UnscopedResourceDiscoveryClient, error) {

	client, err := c.ForKind(v1alpha1.ResourceDiscoveryGVK())
	if err != nil {
		return nil, err
	}

	return &UnscopedResourceDiscoveryClient{client}, nil

}

// WithScope -
func (c *UnscopedResourceDiscoveryClient) WithScope(scope string) *ResourceDiscoveryClient {
	return &ResourceDiscoveryClient{
		c.client.WithScope(scope),
	}
}

// Get -
func (c *UnscopedResourceDiscoveryClient) Get(name string) (*v1alpha1.ResourceDiscovery, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &v1alpha1.ResourceDiscovery{}
	service.FromInstance(ri)

	return service, nil
}

// Update -
func (c *UnscopedResourceDiscoveryClient) Update(res *v1alpha1.ResourceDiscovery, opts ...v1.UpdateOption) (*v1alpha1.ResourceDiscovery, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &v1alpha1.ResourceDiscovery{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

// List -
func (c *ResourceDiscoveryClient) List(options ...v1.ListOptions) ([]*v1alpha1.ResourceDiscovery, error) {
	riList, err := c.client.List(options...)
	if err != nil {
		return nil, err
	}

	result := make([]*v1alpha1.ResourceDiscovery, len(riList))

	for i := range riList {
		result[i] = &v1alpha1.ResourceDiscovery{}
		err := result[i].FromInstance(riList[i])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Get -
func (c *ResourceDiscoveryClient) Get(name string) (*v1alpha1.ResourceDiscovery, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &v1alpha1.ResourceDiscovery{}
	service.FromInstance(ri)

	return service, nil
}

// Delete -
func (c *ResourceDiscoveryClient) Delete(res *v1alpha1.ResourceDiscovery) error {
	ri, err := res.AsInstance()

	if err != nil {
		return err
	}

	return c.client.Delete(ri)
}

// Create -
func (c *ResourceDiscoveryClient) Create(res *v1alpha1.ResourceDiscovery, opts ...v1.CreateOption) (*v1alpha1.ResourceDiscovery, error) {
	ri, err := res.AsInstance()

	if err != nil {
		return nil, err
	}

	cri, err := c.client.Create(ri, opts...)
	if err != nil {
		return nil, err
	}

	created := &v1alpha1.ResourceDiscovery{}

	err = created.FromInstance(cri)
	if err != nil {
		return nil, err
	}

	return created, err
}

// Update -
func (c *ResourceDiscoveryClient) Update(res *v1alpha1.ResourceDiscovery, opts ...v1.UpdateOption) (*v1alpha1.ResourceDiscovery, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &v1alpha1.ResourceDiscovery{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
