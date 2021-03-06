/*
 * Uber API
 *
 * Move your app forward with the Uber API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Activities struct {

	// Position in pagination.
	Offset int32 `json:"offset,omitempty"`

	// Number of items to retrieve (100 max).
	Limit int32 `json:"limit,omitempty"`

	// Total number of items available.
	Count int32 `json:"count,omitempty"`

	History []Activity `json:"history,omitempty"`
}
