/*
 * API Server specification.
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: SNAPSHOT
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1alpha1

// ResourceDefinitionSpecScope struct for ResourceDefinitionSpecScope
type ResourceDefinitionSpecScope struct {
	// Defines the kind of the scope. The server infers this from the endpoint the client submits the request to.
	Kind string `json:"kind,omitempty"`
}
