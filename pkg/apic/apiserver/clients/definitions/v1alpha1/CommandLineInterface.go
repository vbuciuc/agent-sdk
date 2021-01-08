/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"fmt"

	v1 "github.com/Axway/axway-sdk/pkg/apic/apiserver/clients/api/v1"
	apiv1 "github.com/Axway/axway-sdk/pkg/apic/apiserver/models/api/v1"
	"github.com/Axway/axway-sdk/pkg/apic/apiserver/models/definitions/v1alpha1"
)

type CommandLineInterfaceMergeFunc func(*v1alpha1.CommandLineInterface, *v1alpha1.CommandLineInterface) (*v1alpha1.CommandLineInterface, error)

// Merge builds a merge option for an update operation
func CommandLineInterfaceMerge(f CommandLineInterfaceMergeFunc) v1.UpdateOption {
	return v1.Merge(func(prev, new apiv1.Interface) (apiv1.Interface, error) {
		p, n := &v1alpha1.CommandLineInterface{}, &v1alpha1.CommandLineInterface{}

		switch t := prev.(type) {
		case *v1alpha1.CommandLineInterface:
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
		case *v1alpha1.CommandLineInterface:
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

// CommandLineInterfaceClient -
type CommandLineInterfaceClient struct {
	client v1.Scoped
}

// UnscopedCommandLineInterfaceClient -
type UnscopedCommandLineInterfaceClient struct {
	client v1.Unscoped
}

// NewCommandLineInterfaceClient -
func NewCommandLineInterfaceClient(c v1.Base) (*UnscopedCommandLineInterfaceClient, error) {

	client, err := c.ForKind(v1alpha1.CommandLineInterfaceGVK())
	if err != nil {
		return nil, err
	}

	return &UnscopedCommandLineInterfaceClient{client}, nil

}

// WithScope -
func (c *UnscopedCommandLineInterfaceClient) WithScope(scope string) *CommandLineInterfaceClient {
	return &CommandLineInterfaceClient{
		c.client.WithScope(scope),
	}
}

// Get -
func (c *UnscopedCommandLineInterfaceClient) Get(name string) (*v1alpha1.CommandLineInterface, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &v1alpha1.CommandLineInterface{}
	service.FromInstance(ri)

	return service, nil
}

// Update -
func (c *UnscopedCommandLineInterfaceClient) Update(res *v1alpha1.CommandLineInterface, opts ...v1.UpdateOption) (*v1alpha1.CommandLineInterface, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &v1alpha1.CommandLineInterface{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

// List -
func (c *CommandLineInterfaceClient) List(options ...v1.ListOptions) ([]*v1alpha1.CommandLineInterface, error) {
	riList, err := c.client.List(options...)
	if err != nil {
		return nil, err
	}

	result := make([]*v1alpha1.CommandLineInterface, len(riList))

	for i := range riList {
		result[i] = &v1alpha1.CommandLineInterface{}
		err := result[i].FromInstance(riList[i])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Get -
func (c *CommandLineInterfaceClient) Get(name string) (*v1alpha1.CommandLineInterface, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &v1alpha1.CommandLineInterface{}
	service.FromInstance(ri)

	return service, nil
}

// Delete -
func (c *CommandLineInterfaceClient) Delete(res *v1alpha1.CommandLineInterface) error {
	ri, err := res.AsInstance()

	if err != nil {
		return err
	}

	return c.client.Delete(ri)
}

// Create -
func (c *CommandLineInterfaceClient) Create(res *v1alpha1.CommandLineInterface, opts ...v1.CreateOption) (*v1alpha1.CommandLineInterface, error) {
	ri, err := res.AsInstance()

	if err != nil {
		return nil, err
	}

	cri, err := c.client.Create(ri, opts...)
	if err != nil {
		return nil, err
	}

	created := &v1alpha1.CommandLineInterface{}

	err = created.FromInstance(cri)
	if err != nil {
		return nil, err
	}

	return created, err
}

// Update -
func (c *CommandLineInterfaceClient) Update(res *v1alpha1.CommandLineInterface, opts ...v1.UpdateOption) (*v1alpha1.CommandLineInterface, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri, opts...)
	if err != nil {
		return nil, err
	}

	updated := &v1alpha1.CommandLineInterface{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}
