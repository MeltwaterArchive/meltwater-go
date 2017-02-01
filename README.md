# Go API client for Meltwater

_**Disclaimer: This client was generated automatically based on our Swagger Spec. We are sharing it for purely demonstrative purposes. We hope it it helps. If you have any comments please open an issue, we would love to hear from you!**_

The Meltwater API provides the needed resources for Meltwater clients to create & delete REST Hooks and stream Meltwater search results to your specified destination.

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build date: 2017-02-01T12:44:19.404+01:00
- Build package: class io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```
    "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.meltwater.io*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ClientsApi* | [**CreateClientCredentials**](docs/ClientsApi.md#createclientcredentials) | **Post** /v1/clients | Register new client
*ClientsApi* | [**DeleteClientCredentials**](docs/ClientsApi.md#deleteclientcredentials) | **Delete** /v1/clients/{client_id} | Delete client.
*HooksApi* | [**CreateHook**](docs/HooksApi.md#createhook) | **Post** /v1/hooks | Creates a hook for one of your predefined searches.
*HooksApi* | [**DeleteHook**](docs/HooksApi.md#deletehook) | **Delete** /v1/hooks/{id} | Delete an existing hook.
*HooksApi* | [**GetAllHooks**](docs/HooksApi.md#getallhooks) | **Get** /v1/hooks | List all hooks.
*OauthApi* | [**CreateToken**](docs/OauthApi.md#createtoken) | **Post** /oauth2/token | Create an access token
*SearchesApi* | [**GetAllSearches**](docs/SearchesApi.md#getallsearches) | **Get** /v1/searches | List your saved searches.


## Documentation For Models

 - [ClientCredentials](docs/ClientCredentials.md)
 - [ErrorsCollection](docs/ErrorsCollection.md)
 - [Hook](docs/Hook.md)
 - [HooksCollection](docs/HooksCollection.md)
 - [ModelError](docs/ModelError.md)
 - [OAuth2Token](docs/OAuth2Token.md)
 - [PostV1Hooks](docs/PostV1Hooks.md)
 - [Search](docs/Search.md)
 - [SearchesCollection](docs/SearchesCollection.md)


## Documentation For Authorization

 All endpoints do not require authorization.


## Author

support@api.meltwater.com

