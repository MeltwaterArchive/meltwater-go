# \HooksApi

All URIs are relative to *https://api.meltwater.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHook**](HooksApi.md#CreateHook) | **Post** /v2/hooks | Creates a hook for one of your predefined searches.
[**DeleteHook**](HooksApi.md#DeleteHook) | **Delete** /v2/hooks/{hook_id} | Delete an existing hook.
[**GetAllHooks**](HooksApi.md#GetAllHooks) | **Get** /v2/hooks | List all hooks.


# **CreateHook**
> Hook CreateHook($userKey, $authorization, $v2Hooks, $xHookSecret)

Creates a hook for one of your predefined searches.

Creates a hook for one of your predefined searches.  Delivers search results for the query referenced by the `search_id` to the `target_url` via HTTP POST. Note that a `hook_id` will be returned on successful creation, this id is needed to delete the hook.   We are also returning the header: `X-Hook-Secret`, a shared secret used to sign the document body pushed to your `target_url`.    Requires an OAuth access token.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userKey** | **string**| The &#x60;user_key&#x60; from [developer.meltwater.com](https://developer.meltwater.com/admin/applications/). | 
 **authorization** | **string**| &#x60;Oauth Access Token&#x60;    OAuth access token (RFC 6749). Must contain the access token type &#x60;Bearer&#x60;  followed by an OAuth access token.    #### Example:        Bearer KKwmfHwxsEoeMDTMAfxOpO... | 
 **v2Hooks** | [**PostV2Hooks**](PostV2Hooks.md)|  | 
 **xHookSecret** | **string**| Shared secret for content signing/verification.    The shared secret header is optional and can be provided by the user or will  be set by the API. Must be between 16 and 64 characters.  Obtain the shared secret from the response header &#x60;X-Hook-Secret&#x60;.    #### Example:        e2d264b524240b9572ebc2fc7eebd980 | [optional] 

### Return type

[**Hook**](Hook.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteHook**
> DeleteHook($userKey, $authorization, $hookId)

Delete an existing hook.

Delete an existing hook.  Removes the hook and stops sending any search results to the target_url.    Requires an OAuth access token.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userKey** | **string**| The &#x60;user_key&#x60; from [developer.meltwater.com](https://developer.meltwater.com/admin/applications/). | 
 **authorization** | **string**| &#x60;Oauth Access Token&#x60;    OAuth access token (RFC 6749). Must contain the access token type &#x60;Bearer&#x60;  followed by an OAuth access token.    #### Example:        Bearer KKwmfHwxsEoeMDTMAfxOpO... | 
 **hookId** | **string**| Hook ID received from creating a hook | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllHooks**
> HooksCollection GetAllHooks($userKey, $authorization)

List all hooks.

List all hooks.     Delivers all previously generated hooks.    Requires an OAuth access token.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userKey** | **string**| The &#x60;user_key&#x60; from [developer.meltwater.com](https://developer.meltwater.com/admin/applications/). | 
 **authorization** | **string**| &#x60;Oauth Access Token&#x60;    OAuth access token (RFC 6749). Must contain the access token type &#x60;Bearer&#x60;  followed by an OAuth access token.    #### Example:        Bearer KKwmfHwxsEoeMDTMAfxOpO... | 

### Return type

[**HooksCollection**](HooksCollection.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

