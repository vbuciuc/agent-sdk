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

type K8SResourceMergeFunc func(*v1alpha1.K8SResource, *v1alpha1.K8SResource) (*v1alpha1.K8SResource, error)

// Merge builds a merge option for an update operation
func K8SResourceMerge(f K8SResourceMergeFunc) v1.UpdateOption {
	return v1.Merge(func(prev, new apiv1.Interface) (apiv1.Interface, error) {
		p, n := &v1alpha1.K8SResource{}, &v1alpha1.K8SResource{}

		switch t := prev.(type) {
		case *v1alpha1.K8SResource:
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
		case *v1alpha1.K8SResource:
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

// K8SResourceClient -
type K8SResourceClient struct {
	client v1.Scoped
}

// UnscopedK8SResourceClient -
type UnscopedK8SResourceClient struct {
	client v1.Unscoped
}

// NewK8SResourceClient -
func NewK8SResourceClient(c v1.Base) (*UnscopedK8SResourceClient, error) {

	client, err := c.ForKind(v1alpha1.K8SResourceGVK())
	if err != nil {
		return nil, err
	}

	return &UnscopedK8SResourceClient{client}, nil

}

// WithScope -
func (c *UnscopedK8SResourceClient) WithScope(scope string) *K8SResourceClient {
	return &K8SResourceClient{
		c.client.WithScope(scope),
	}
}

// Get -
func (c *UnscopedK8SResourceClient) Get(name string) (*v1alpha1.K8SResource, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &v1alpha1.K8SResource{}
	service.FromInstance(ri)

	return service, nil
}

// Update -
func (c *UnscopedK8SResourceClient) Update(res *v1alpha1.K8SResource, opts ...v1.UpdateOption) (*v1alpha1.K8SResource, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &v1alpha1.K8SResource{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

// List -
func (c *K8SResourceClient) List(options ...v1.ListOptions) ([]*v1alpha1.K8SResource, error) {
	riList, err := c.client.List(options...)
	if err != nil {
		return nil, err
	}

	result := make([]*v1alpha1.K8SResource, len(riList))

	for i := range riList {
		result[i] = &v1alpha1.K8SResource{}
		err := result[i].FromInstance(riList[i])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Get -
func (c *K8SResourceClient) Get(name string) (*v1alpha1.K8SResource, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &v1alpha1.K8SResource{}
	service.FromInstance(ri)

	return service, nil
}

// Delete -
func (c *K8SResourceClient) Delete(res *v1alpha1.K8SResource) error {
	ri, err := res.AsInstance()

	if err != nil {
		return err
	}

	return c.client.Delete(ri)
}

// Create -
func (c *K8SResourceClient) Create(res *v1alpha1.K8SResource, opts ...v1.CreateOption) (*v1alpha1.K8SResource, error) {
	ri, err := res.AsInstance()

	if err != nil {
		return nil, err
	}

	cri, err := c.client.Create(ri, opts...)
	if err != nil {
		return nil, err
	}

	created := &v1alpha1.K8SResource{}

	err = created.FromInstance(cri)
	if err != nil {
		return nil, err
	}

	return created, err
}

// Update -
func (c *K8SResourceClient) Update(res *v1alpha1.K8SResource, opts ...v1.UpdateOption) (*v1alpha1.K8SResource, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &v1alpha1.K8SResource{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
