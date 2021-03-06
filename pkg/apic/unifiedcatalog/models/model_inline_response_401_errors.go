/*
 * Amplify Unified Catalog APIs
 *
 * APIs for Amplify Unified Catalog
 *
 * API version: 1.43.0
 * Contact: support@axway.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package unifiedcatalog
// InlineResponse401Errors struct for InlineResponse401Errors
type InlineResponse401Errors struct {
	Status int32 `json:"status,omitempty"`
	Title string `json:"title,omitempty"`
	Detail string `json:"detail,omitempty"`
	Meta InlineResponse401Meta `json:"meta,omitempty"`
}
