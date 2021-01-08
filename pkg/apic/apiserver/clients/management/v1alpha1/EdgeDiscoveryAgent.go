/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"fmt"

	v1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/clients/api/v1"
	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
	"github.com/Axway/agent-sdk/pkg/apic/apiserver/models/management/v1alpha1"
)

type EdgeDiscoveryAgentMergeFunc func(*v1alpha1.EdgeDiscoveryAgent, *v1alpha1.EdgeDiscoveryAgent) (*v1alpha1.EdgeDiscoveryAgent, error)

// Merge builds a merge option for an update operation
func EdgeDiscoveryAgentMerge(f EdgeDiscoveryAgentMergeFunc) v1.UpdateOption {
	return v1.Merge(func(prev, new apiv1.Interface) (apiv1.Interface, error) {
		p, n := &v1alpha1.EdgeDiscoveryAgent{}, &v1alpha1.EdgeDiscoveryAgent{}

		switch t := prev.(type) {
		case *v1alpha1.EdgeDiscoveryAgent:
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
		case *v1alpha1.EdgeDiscoveryAgent:
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

// EdgeDiscoveryAgentClient -
type EdgeDiscoveryAgentClient struct {
	client v1.Scoped
}

// UnscopedEdgeDiscoveryAgentClient -
type UnscopedEdgeDiscoveryAgentClient struct {
	client v1.Unscoped
}

// NewEdgeDiscoveryAgentClient -
func NewEdgeDiscoveryAgentClient(c v1.Base) (*UnscopedEdgeDiscoveryAgentClient, error) {

	client, err := c.ForKind(v1alpha1.EdgeDiscoveryAgentGVK())
	if err != nil {
		return nil, err
	}

	return &UnscopedEdgeDiscoveryAgentClient{client}, nil

}

// WithScope -
func (c *UnscopedEdgeDiscoveryAgentClient) WithScope(scope string) *EdgeDiscoveryAgentClient {
	return &EdgeDiscoveryAgentClient{
		c.client.WithScope(scope),
	}
}

// Get -
func (c *UnscopedEdgeDiscoveryAgentClient) Get(name string) (*v1alpha1.EdgeDiscoveryAgent, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &v1alpha1.EdgeDiscoveryAgent{}
	service.FromInstance(ri)

	return service, nil
}

// Update -
func (c *UnscopedEdgeDiscoveryAgentClient) Update(res *v1alpha1.EdgeDiscoveryAgent, opts ...v1.UpdateOption) (*v1alpha1.EdgeDiscoveryAgent, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &v1alpha1.EdgeDiscoveryAgent{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

// List -
func (c *EdgeDiscoveryAgentClient) List(options ...v1.ListOptions) ([]*v1alpha1.EdgeDiscoveryAgent, error) {
	riList, err := c.client.List(options...)
	if err != nil {
		return nil, err
	}

	result := make([]*v1alpha1.EdgeDiscoveryAgent, len(riList))

	for i := range riList {
		result[i] = &v1alpha1.EdgeDiscoveryAgent{}
		err := result[i].FromInstance(riList[i])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Get -
func (c *EdgeDiscoveryAgentClient) Get(name string) (*v1alpha1.EdgeDiscoveryAgent, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &v1alpha1.EdgeDiscoveryAgent{}
	service.FromInstance(ri)

	return service, nil
}

// Delete -
func (c *EdgeDiscoveryAgentClient) Delete(res *v1alpha1.EdgeDiscoveryAgent) error {
	ri, err := res.AsInstance()

	if err != nil {
		return err
	}

	return c.client.Delete(ri)
}

// Create -
func (c *EdgeDiscoveryAgentClient) Create(res *v1alpha1.EdgeDiscoveryAgent, opts ...v1.CreateOption) (*v1alpha1.EdgeDiscoveryAgent, error) {
	ri, err := res.AsInstance()

	if err != nil {
		return nil, err
	}

	cri, err := c.client.Create(ri, opts...)
	if err != nil {
		return nil, err
	}

	created := &v1alpha1.EdgeDiscoveryAgent{}

	err = created.FromInstance(cri)
	if err != nil {
		return nil, err
	}

	return created, err
}

// Update -
func (c *EdgeDiscoveryAgentClient) Update(res *v1alpha1.EdgeDiscoveryAgent, opts ...v1.UpdateOption) (*v1alpha1.EdgeDiscoveryAgent, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &v1alpha1.EdgeDiscoveryAgent{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
