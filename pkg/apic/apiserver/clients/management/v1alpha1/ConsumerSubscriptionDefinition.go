/*
 * This file is automatically generated
 */

package v1alpha1

import (
	v1 "git.ecd.axway.int/apigov/apic_agents_sdk/pkg/apic/apiserver/clients/api/v1"
	"git.ecd.axway.int/apigov/apic_agents_sdk/pkg/apic/apiserver/models/management/v1alpha1"
)

// ConsumerSubscriptionDefinitionClient -
type ConsumerSubscriptionDefinitionClient struct {
	client *v1.Client
}

// NewConsumerSubscriptionDefinitionClient -
func NewConsumerSubscriptionDefinitionClient(cb *v1.ClientBase) (*ConsumerSubscriptionDefinitionClient, error) {
	client, err := cb.ForKind(v1alpha1.ConsumerSubscriptionDefinitionGVK())
	if err != nil {
		return nil, err
	}

	return &ConsumerSubscriptionDefinitionClient{client}, nil
}

// WithScope -
func (c *ConsumerSubscriptionDefinitionClient) WithScope(scope string) *ConsumerSubscriptionDefinitionClient {
	return &ConsumerSubscriptionDefinitionClient{
		c.client.WithScope(scope),
	}
}

// List -
func (c *ConsumerSubscriptionDefinitionClient) List(options ...v1.ListOptions) ([]*v1alpha1.ConsumerSubscriptionDefinition, error) {
	riList, err := c.client.List(options...)
	if err != nil {
		return nil, err
	}

	result := make([]*v1alpha1.ConsumerSubscriptionDefinition, len(riList))

	for i := range riList {
		result[i] = &v1alpha1.ConsumerSubscriptionDefinition{}
		err := result[i].FromInstance(riList[i])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Get -
func (c *ConsumerSubscriptionDefinitionClient) Get(name string) (*v1alpha1.ConsumerSubscriptionDefinition, error) {
	ri, err := c.client.Get(name)
	if err != nil {
		return nil, err
	}

	service := &v1alpha1.ConsumerSubscriptionDefinition{}
	service.FromInstance(ri)

	return service, nil
}

// Delete -
func (c *ConsumerSubscriptionDefinitionClient) Delete(res *v1alpha1.ConsumerSubscriptionDefinition) error {
	ri, err := res.AsInstance()

	if err != nil {
		return err
	}

	return c.client.Delete(ri)
}

// Create -
func (c *ConsumerSubscriptionDefinitionClient) Create(res *v1alpha1.ConsumerSubscriptionDefinition) (*v1alpha1.ConsumerSubscriptionDefinition, error) {
	ri, err := res.AsInstance()

	if err != nil {
		return nil, err
	}

	cri, err := c.client.Create(ri)
	if err != nil {
		return nil, err
	}

	created := &v1alpha1.ConsumerSubscriptionDefinition{}

	err = created.FromInstance(cri)
	if err != nil {
		return nil, err
	}

	return created, err
}

// Update -
func (c *ConsumerSubscriptionDefinitionClient) Update(res *v1alpha1.ConsumerSubscriptionDefinition) (*v1alpha1.ConsumerSubscriptionDefinition, error) {
	ri, err := res.AsInstance()
	if err != nil {
		return nil, err
	}
	resource, err := c.client.Update(ri)
	updated := &v1alpha1.ConsumerSubscriptionDefinition{}

	// Updates the resource in place
	err = updated.FromInstance(resource)
	if err != nil {
		return nil, err
	}

	return updated, nil
}