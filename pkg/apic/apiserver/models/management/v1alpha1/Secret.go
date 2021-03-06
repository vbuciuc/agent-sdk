/*
 * This file is automatically generated
 */

package v1alpha1

import (
	"encoding/json"

	apiv1 "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/api/v1"
)

var (
	_SecretGVK = apiv1.GroupVersionKind{
		GroupKind: apiv1.GroupKind{
			Group: "management",
			Kind:  "Secret",
		},
		APIVersion: "v1alpha1",
	}
)

const (
	SecretScope = "Integration"

	SecretResource = "secrets"
)

func SecretGVK() apiv1.GroupVersionKind {
	return _SecretGVK
}

func init() {
	apiv1.RegisterGVK(_SecretGVK, SecretScope, SecretResource)
}

// Secret Resource
type Secret struct {
	apiv1.ResourceMeta

	Spec SecretSpec `json:"spec"`
}

// FromInstance converts a ResourceInstance to a Secret
func (res *Secret) FromInstance(ri *apiv1.ResourceInstance) error {
	if ri == nil {
		res = nil
		return nil
	}

	m, err := json.Marshal(ri.Spec)
	if err != nil {
		return err
	}

	spec := &SecretSpec{}
	err = json.Unmarshal(m, spec)
	if err != nil {
		return err
	}

	*res = Secret{ResourceMeta: ri.ResourceMeta, Spec: *spec}

	return err
}

// AsInstance converts a Secret to a ResourceInstance
func (res *Secret) AsInstance() (*apiv1.ResourceInstance, error) {
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
	meta.GroupVersionKind = SecretGVK()

	return &apiv1.ResourceInstance{ResourceMeta: meta, Spec: spec}, nil
}
