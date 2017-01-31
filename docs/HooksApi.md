# \HooksApi

All URIs are relative to *https://api.meltwater.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHook**](HooksApi.md#CreateHook) | **Post** /v1/hooks | Creates a hook for one of your predefined searches.
[**DeleteHook**](HooksApi.md#DeleteHook) | **Delete** /v1/hooks/{id} | Delete an existing hook.
[**GetAllHooks**](HooksApi.md#GetAllHooks) | **Get** /v1/hooks | List all hooks.


# **CreateHook**
> Hook CreateHook($userKey, $authorization, $v1Hooks)

Creates a hook for one of your predefined searches.

Creates a hook for one of your predefined searches.  Delivers search results for the query referenced by thesearch_id to the target_url via HTTP POST. Note that a hook id will be returned on successful creation, this id is needed to delete the hook.     Requires an OAuth access token.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userKey** | **string**| The &#x60;user_key&#x60; from [developer.meltwater.io](https://developer.meltwater.io/admin/applications/). | 
 **authorization** | **string**| &#x60;Oauth Access Token&#x60;    OAuth access token (RFC 6749). Must contain the access token type &#x60;Bearer&#x60;  followed by an OAuth access token.    #### Example:        Bearer KKwmfHwxsEoeMDTMAfxOpO... | 
 **v1Hooks** | [**PostV1Hooks**](PostV1Hooks.md)|  | 

### Return type

[**Hook**](Hook.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteHook**
> DeleteHook($userKey, $authorization, $id)

Delete an existing hook.

Delete an existing hook.  Removes the hook and stops sending any search results to the target_url.    Requires an OAuth access token.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userKey** | **string**| The &#x60;user_key&#x60; from [developer.meltwater.io](https://developer.meltwater.io/admin/applications/). | 
 **authorization** | **string**| &#x60;Oauth Access Token&#x60;    OAuth access token (RFC 6749). Must contain the access token type &#x60;Bearer&#x60;  followed by an OAuth access token.    #### Example:        Bearer KKwmfHwxsEoeMDTMAfxOpO... | 
 **id** | **string**| Hook ID received from creating a hook | 

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
 **userKey** | **string**| The &#x60;user_key&#x60; from [developer.meltwater.io](https://developer.meltwater.io/admin/applications/). | 
 **authorization** | **string**| &#x60;Oauth Access Token&#x60;    OAuth access token (RFC 6749). Must contain the access token type &#x60;Bearer&#x60;  followed by an OAuth access token.    #### Example:        Bearer KKwmfHwxsEoeMDTMAfxOpO... | 

### Return type

[**HooksCollection**](HooksCollection.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

