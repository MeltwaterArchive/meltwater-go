/* 
 * Meltwater Streaming API v2
 *
 * The Meltwater Streaming API provides the needed resources for Meltwater clients to create & delete REST Hooks and stream Meltwater search results to your specified destination.
 *
 * OpenAPI spec version: 2.0.0
 * Contact: support@api.meltwater.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

// Creates a hook for one of your predefined searches.
type PostV2Hooks struct {

	// Target URL to send article results
	TargetUrl string `json:"target_url"`

	// Search ID
	SearchId int32 `json:"search_id"`
}
