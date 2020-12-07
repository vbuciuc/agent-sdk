/*
 * This file is automatically generated
 */

package v1alpha1

import (
	v1 "git.ecd.axway.org/apigov/apic_agents_sdk/pkg/apic/apiserver/clients/api/v1"
	"git.ecd.axway.org/apigov/apic_agents_sdk/pkg/apic/apiserver/models/management/v1alpha1"
)

// MeshDiscoveryClient -
type MeshDiscoveryClient struct {
	client v1.Scoped
}

// UnscopedMeshDiscoveryClient -
type UnscopedMeshDiscoveryClient struct {
	client v1.Unscoped
}

// NewMeshDiscoveryClient -
func NewMeshDiscoveryClient(c v1.Base) (*UnscopedMeshDiscoveryClient, error) {

	client, err := c.ForKind(v1alpha1.MeshDiscoveryGVK())
	if err != nil {
		return nil, err
	}

	return &UnscopedMeshDiscoveryClient{client}, nil

}

// WithScope -
func (c *UnscopedMeshDiscoveryClient) WithScope(scope string) *MeshDiscoveryClient {
	return &MeshDiscoveryClient{
		c.client.WithScope(scope),
	}
}

// List -
func (c *MeshDiscoveryClient) List(options ...v1.ListOptions) ([]*v1alpha1.MeshDiscovery, error) {
	riList, err := c.client.List(options...)
	if err != nil {
		return nil, err
	}

	result := make([]*v1alpha1.MeshDiscovery, len(riList))

	for i := range riList {
		result[i] = &v1alpha1.MeshDiscovery{}
		err := result[i].FromInstance(riList[i])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Get -
func (c *MeshDiscoveryClient) Get(name string) (*v1alpha1.MeshDiscovery, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &v1alpha1.MeshDiscovery{}
	service.FromInstance(ri)

	return service, nil
}

// Delete -
func (c *MeshDiscoveryClient) Delete(res *v1alpha1.MeshDiscovery) error {
	ri, err := res.AsInstance()

	if err != nil {
		return err
	}

	return c.client.Delete(ri)
}

// Create -
func (c *MeshDiscoveryClient) Create(res *v1alpha1.MeshDiscovery) (*v1alpha1.MeshDiscovery, error) {
	ri, err := res.AsInstance()

	if err != nil {
		return nil, err
	}

	cri, err := c.client.Create(ri)
	if err != nil {
		return nil, err
	}

	created := &v1alpha1.MeshDiscovery{}

	err = created.FromInstance(cri)
	if err != nil {
		return nil, err
	}

	return created, err
}

// Update -
func (c *MeshDiscoveryClient) Update(res *v1alpha1.MeshDiscovery) (*v1alpha1.MeshDiscovery, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri)
	if err != nil {
		return nil, err
	}

	updated := &v1alpha1.MeshDiscovery{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}