/* 
 * The Meltwater API
 *
 * The Meltwater API provides the needed resources for Meltwater clients to create & delete REST Hooks and stream Meltwater search results to your specified destination.
 *
 * OpenAPI spec version: 0.9.0
 * Contact: support@api.meltwater.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package swagger

// Creates a hook for one of your predefined searches.
type PostV1Hooks struct {

	// Target URL to send article results
	TargetUrl string `json:"target_url,omitempty"`

	// Search ID
	SearchId int32 `json:"search_id,omitempty"`
}