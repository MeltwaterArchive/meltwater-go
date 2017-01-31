/* 
 * The Meltwater API
 *
 * The Meltwater API provides the needed resources for Meltwater clients to create & delete REST Hooks and stream Meltwater search results to your specified destination.
 *
 * OpenAPI spec version: 1.0.0
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

import (
	"net/url"
	"encoding/json"
)

type OauthApi struct {
	Configuration Configuration
}

func NewOauthApi() *OauthApi {
	configuration := NewConfiguration()
	return &OauthApi{
		Configuration: *configuration,
	}
}

func NewOauthApiWithBasePath(basePath string) *OauthApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &OauthApi{
		Configuration: *configuration,
	}
}

/**
 * Create an access token
 * Create an OAuth2 access token based on the provided &#x60;client_id&#x60; and &#x60;client_secret&#x60;
 *
 * @param userKey The &#x60;user_key&#x60; from [developer.meltwater.io](https://developer.meltwater.io/admin/applications/).
 * @param authorization &#x60;client_id:client_secret&#x60;  Basic Auth (RFC2617) credentials. Must contain the realm &#x60;Basic&#x60; followed by a Base64-encoded &#x60;client_id&#x60;:&#x60;client_secret&#x60; pair.   #### Example:      Basic aAlfbb1haWxDSXhhDXxxZWKJAyZXQ&#x3D;
 * @param grantType OAuth2 grant type, use &#x60;client_credentials&#x60;
 * @param scope OAuth2 scope, use &#x60;search&#x60;
 * @return *OAuth2Token
 */
func (a OauthApi) CreateToken(userKey string, authorization string, grantType string, scope string) (*OAuth2Token, *APIResponse, error) {

	var httpMethod = "Post"
	// create path and map variables
	path := a.Configuration.BasePath + "/oauth2/token"


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/x-www-form-urlencoded",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}

	// header params "user-key"
	headerParams["user-key"] = userKey
	// header params "Authorization"
	headerParams["Authorization"] = authorization


	formParams["grantType"] = grantType
	formParams["scope"] = scope
	var successPayload = new(OAuth2Token)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

