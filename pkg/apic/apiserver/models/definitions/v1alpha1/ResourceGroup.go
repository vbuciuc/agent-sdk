/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "git.ecd.axway.int/apigov/apic_agents_sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_ResourceGroupGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "definitions",
			Kind:  "ResourceGroup",
		},
		ApiVersion: "v1alpha1",
	}
)

const (
	ResourceGroupScope = ""

	ResourceGroupResource = "groups"
)

func ResourceGroupGVK() apiv1.GroupVersionKind {
	return _ResourceGroupGVK
}

func init() {
	apiv1.RegisterGVK(_ResourceGroupGVK, ResourceGroupScope, ResourceGroupResource)
}

type ResourceGroup struct {
	apiv1.ResourceMeta

	Spec struct{} `json:"spec"`
}

func (res *ResourceGroup) FromInstance(ri *apiv1.ResourceInstance) error {
	// TODO this needs to be better
	m, err := json.Marshal(ri.Spec)
	if err != nil {
		return err
	}

	spec := &struct{}{}
	err = json.Unmarshal(m, spec)
	if err != nil {
		return err
	}

	*res = ResourceGroup{ResourceMeta: ri.ResourceMeta, Spec: *spec}

	return err
}

func (res *ResourceGroup) AsInstance() (*apiv1.ResourceInstance, error) {
	m, err := json.Marshal(res.Spec)
	if err != nil {
		return nil, err
	}

	spec := map[string]interface{}{}
	err = json.Unmarshal(m, &spec)
	if err != nil {
		return nil, err
	}

	return &apiv1.ResourceInstance{ResourceMeta: res.ResourceMeta, Spec: spec}, nil
}
