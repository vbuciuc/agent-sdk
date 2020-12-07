/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// ResourceDiscoverySpec struct for ResourceDiscoverySpec
type ResourceDiscoverySpec struct {
	Version          string                               `json:"version,omitempty"`
	Kind             string                               `json:"kind,omitempty"`
	Group            string                               `json:"group,omitempty"`
	NamespaceFilter  ResourceDiscoverySpecNamespaceFilter `json:"namespaceFilter,omitempty"`
	ResourceFilter   ResourceDiscoverySpecResourceFilter  `json:"resourceFilter,omitempty"`
	KeepSpecFields   []string                             `json:"keepSpecFields,omitempty"`
	KeepStatusFields []string                             `json:"keepStatusFields,omitempty"`
	IgnoreLabels     []string                             `json:"ignoreLabels,omitempty"`
	Tags             []string                             `json:"tags,omitempty"`
	ExtraAttributes  map[string]string                    `json:"extraAttributes,omitempty"`
}