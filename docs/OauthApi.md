# \OauthApi

All URIs are relative to *https://api.meltwater.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostOauth2Token**](OauthApi.md#PostOauth2Token) | **Post** /oauth2/token | Create an access token


# **PostOauth2Token**
> OAuth2Token PostOauth2Token($userKey, $authorization, $grantType, $scope)

Create an access token

Create an OAuth2 access token based on the provided `client_id` and `client_secret`


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userKey** | **string**| The &#x60;user_key&#x60; from [developer.meltwater.io](https://developer.meltwater.io/admin/applications/). | 
 **authorization** | **string**| &#x60;client_id:client_secret&#x60;  Basic Auth (RFC2617) credentials. Must contain the realm &#x60;Basic&#x60; followed by a Base64-encoded &#x60;client_id&#x60;:&#x60;client_secret&#x60; pair.   #### Example:      Basic aAlfbb1haWxDSXhhDXxxZWKJAyZXQ&#x3D; | 
 **grantType** | **string**| OAuth2 grant type, use &#x60;client_credentials&#x60; | 
 **scope** | **string**| OAuth2 scope, use &#x60;search&#x60; | 

### Return type

[**OAuth2Token**](OAuth2Token.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

