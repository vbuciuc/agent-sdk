/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"fmt"

	v1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/api/v1"
	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
	"github.com/Axway/agent-sdk/pkg/apic/apiserver/models/definitions/v1alpha1"
)

type ResourceDefinitionMergeFunc func(*v1alpha1.ResourceDefinition, *v1alpha1.ResourceDefinition) (*v1alpha1.ResourceDefinition, error)

// Merge builds a merge option for an update operation
func ResourceDefinitionMerge(f ResourceDefinitionMergeFunc) v1.UpdateOption {
	return v1.Merge(func(prev, new apiv1.Interface) (apiv1.Interface, error) {
		p, n := &v1alpha1.ResourceDefinition{}, &v1alpha1.ResourceDefinition{}

		switch t := prev.(type) {
		case *v1alpha1.ResourceDefinition:
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
		case *v1alpha1.ResourceDefinition:
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

// ResourceDefinitionClient -
type ResourceDefinitionClient struct {
	client v1.Scoped
}

// UnscopedResourceDefinitionClient -
type UnscopedResourceDefinitionClient struct {
	client v1.Unscoped
}

// NewResourceDefinitionClient -
func NewResourceDefinitionClient(c v1.Base) (*UnscopedResourceDefinitionClient, error) {

	client, err := c.ForKind(v1alpha1.ResourceDefinitionGVK())
	if err != nil {
		return nil, err
	}

	return &UnscopedResourceDefinitionClient{client}, nil

}

// WithScope -
func (c *UnscopedResourceDefinitionClient) WithScope(scope string) *ResourceDefinitionClient {
	return &ResourceDefinitionClient{
		c.client.WithScope(scope),
	}
}

// Get -
func (c *UnscopedResourceDefinitionClient) Get(name string) (*v1alpha1.ResourceDefinition, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &v1alpha1.ResourceDefinition{}
	service.FromInstance(ri)

	return service, nil
}

// Update -
func (c *UnscopedResourceDefinitionClient) Update(res *v1alpha1.ResourceDefinition, opts ...v1.UpdateOption) (*v1alpha1.ResourceDefinition, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &v1alpha1.ResourceDefinition{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

// List -
func (c *ResourceDefinitionClient) List(options ...v1.ListOptions) ([]*v1alpha1.ResourceDefinition, error) {
	riList, err := c.client.List(options...)
	if err != nil {
		return nil, err
	}

	result := make([]*v1alpha1.ResourceDefinition, len(riList))

	for i := range riList {
		result[i] = &v1alpha1.ResourceDefinition{}
		err := result[i].FromInstance(riList[i])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Get -
func (c *ResourceDefinitionClient) Get(name string) (*v1alpha1.ResourceDefinition, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &v1alpha1.ResourceDefinition{}
	service.FromInstance(ri)

	return service, nil
}

// Delete -
func (c *ResourceDefinitionClient) Delete(res *v1alpha1.ResourceDefinition) error {
	ri, err := res.AsInstance()

	if err != nil {
		return err
	}

	return c.client.Delete(ri)
}

// Create -
func (c *ResourceDefinitionClient) Create(res *v1alpha1.ResourceDefinition, opts ...v1.CreateOption) (*v1alpha1.ResourceDefinition, error) {
	ri, err := res.AsInstance()

	if err != nil {
		return nil, err
	}

	cri, err := c.client.Create(ri, opts...)
	if err != nil {
		return nil, err
	}

	created := &v1alpha1.ResourceDefinition{}

	err = created.FromInstance(cri)
	if err != nil {
		return nil, err
	}

	return created, err
}

// Update -
func (c *ResourceDefinitionClient) Update(res *v1alpha1.ResourceDefinition, opts ...v1.UpdateOption) (*v1alpha1.ResourceDefinition, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &v1alpha1.ResourceDefinition{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
