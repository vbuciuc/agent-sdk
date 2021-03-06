/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_ResourceDefinitionGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "definitions",
			Kind:  "ResourceDefinition",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	ResourceDefinitionScope = "ResourceGroup"

	ResourceDefinitionResource = "resources"
)

func ResourceDefinitionGVK() apiv1.GroupVersionKind {
	return _ResourceDefinitionGVK
}

func init() {
	apiv1.RegisterGVK(_ResourceDefinitionGVK, ResourceDefinitionScope, ResourceDefinitionResource)
}

// ResourceDefinition Resource
type ResourceDefinition struct {
	apiv1.ResourceMeta

	Spec ResourceDefinitionSpec `json:"spec"`
}

// FromInstance converts a ResourceInstance to a ResourceDefinition
func (res *ResourceDefinition) FromInstance(ri *apiv1.ResourceInstance) error {
	if ri == nil {
		res = nil
		return nil
	}

	m, err := json.Marshal(ri.Spec)
	if err != nil {
		return err
	}

	spec := &ResourceDefinitionSpec{}
	err = json.Unmarshal(m, spec)
	if err != nil {
		return err
	}

	*res = ResourceDefinition{ResourceMeta: ri.ResourceMeta, Spec: *spec}

	return err
}

// AsInstance converts a ResourceDefinition to a ResourceInstance
func (res *ResourceDefinition) AsInstance() (*apiv1.ResourceInstance, error) {
	m, err := json.Marshal(res.Spec)
	if err != nil {
		return nil, err
	}

	spec := map[string]interface{}{}
	err = json.Unmarshal(m, &spec)
	if err != nil {
		return nil, err
	}

	meta := res.ResourceMeta
	meta.GroupVersionKind = ResourceDefinitionGVK()

	return &apiv1.ResourceInstance{ResourceMeta: meta, Spec: spec}, nil
}
