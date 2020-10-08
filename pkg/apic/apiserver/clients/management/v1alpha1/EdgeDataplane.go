/*
 * This file is automatically generated
 */

package v1alpha1

import (
	v1 "git.ecd.axway.org/apigov/apic_agents_sdk/pkg/apic/apiserver/clients/api/v1"
	"git.ecd.axway.org/apigov/apic_agents_sdk/pkg/apic/apiserver/models/management/v1alpha1"
)

// EdgeDataplaneClient -
type EdgeDataplaneClient struct {
	client *v1.Client
}

// NewEdgeDataplaneClient -
func NewEdgeDataplaneClient(cb *v1.ClientBase) (*EdgeDataplaneClient, error) {
	client, err := cb.ForKind(v1alpha1.EdgeDataplaneGVK())
	if err != nil {
		return nil, err
	}

	return &EdgeDataplaneClient{client}, nil
}

// WithScope -
func (c *EdgeDataplaneClient) WithScope(scope string) *EdgeDataplaneClient {
	return &EdgeDataplaneClient{
		c.client.WithScope(scope),
	}
}

// List -
func (c *EdgeDataplaneClient) List(options ...v1.ListOptions) ([]*v1alpha1.EdgeDataplane, error) {
	riList, err := c.client.List(options...)
	if err != nil {
		return nil, err
	}

	result := make([]*v1alpha1.EdgeDataplane, len(riList))

	for i := range riList {
		result[i] = &v1alpha1.EdgeDataplane{}
		err := result[i].FromInstance(riList[i])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Get -
func (c *EdgeDataplaneClient) Get(name string) (*v1alpha1.EdgeDataplane, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &v1alpha1.EdgeDataplane{}
	service.FromInstance(ri)

	return service, nil
}

// Delete -
func (c *EdgeDataplaneClient) Delete(res *v1alpha1.EdgeDataplane) error {
	ri, err := res.AsInstance()

	if err != nil {
		return err
	}

	return c.client.Delete(ri)
}

// Create -
func (c *EdgeDataplaneClient) Create(res *v1alpha1.EdgeDataplane) (*v1alpha1.EdgeDataplane, error) {
	ri, err := res.AsInstance()

	if err != nil {
		return nil, err
	}

	cri, err := c.client.Create(ri)
	if err != nil {
		return nil, err
	}

	created := &v1alpha1.EdgeDataplane{}

	err = created.FromInstance(cri)
	if err != nil {
		return nil, err
	}

	return created, err
}

// Update -
func (c *EdgeDataplaneClient) Update(res *v1alpha1.EdgeDataplane) (*v1alpha1.EdgeDataplane, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri)
	updated := &v1alpha1.EdgeDataplane{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}