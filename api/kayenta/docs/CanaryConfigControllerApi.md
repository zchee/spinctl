# \CanaryConfigControllerApi

All URIs are relative to *https://spinnaker-kayenta.meteoros.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCanaryConfigUsingDELETE**](CanaryConfigControllerApi.md#DeleteCanaryConfigUsingDELETE) | **Delete** /canaryConfig/{canaryConfigId} | Delete a canary config
[**ListAllCanaryConfigsUsingGET**](CanaryConfigControllerApi.md#ListAllCanaryConfigsUsingGET) | **Get** /canaryConfig | Retrieve a list of canary config ids and timestamps
[**LoadCanaryConfigUsingGET**](CanaryConfigControllerApi.md#LoadCanaryConfigUsingGET) | **Get** /canaryConfig/{canaryConfigId} | Retrieve a canary config from object storage
[**StoreCanaryConfigUsingPOST**](CanaryConfigControllerApi.md#StoreCanaryConfigUsingPOST) | **Post** /canaryConfig | Write a canary config to object storage
[**UpdateCanaryConfigUsingPUT**](CanaryConfigControllerApi.md#UpdateCanaryConfigUsingPUT) | **Put** /canaryConfig/{canaryConfigId} | Update a canary config


# **DeleteCanaryConfigUsingDELETE**
> DeleteCanaryConfigUsingDELETE(ctx, canaryConfigId, optional)
Delete a canary config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **canaryConfigId** | **string**| canaryConfigId | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **canaryConfigId** | **string**| canaryConfigId | 
 **configurationAccountName** | **string**| configurationAccountName | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAllCanaryConfigsUsingGET**
> []Mapstringobject ListAllCanaryConfigsUsingGET(ctx, optional)
Retrieve a list of canary config ids and timestamps

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configurationAccountName** | **string**| configurationAccountName | 
 **application** | [**[]string**](string.md)| application | 

### Return type

[**[]Mapstringobject**](Map«string,object».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadCanaryConfigUsingGET**
> CanaryConfig LoadCanaryConfigUsingGET(ctx, canaryConfigId, optional)
Retrieve a canary config from object storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **canaryConfigId** | **string**| canaryConfigId | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **canaryConfigId** | **string**| canaryConfigId | 
 **configurationAccountName** | **string**| configurationAccountName | 

### Return type

[**CanaryConfig**](CanaryConfig.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StoreCanaryConfigUsingPOST**
> CanaryConfigUpdateResponse StoreCanaryConfigUsingPOST(ctx, canaryConfig, optional)
Write a canary config to object storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **canaryConfig** | [**CanaryConfig**](CanaryConfig.md)| canaryConfig | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **canaryConfig** | [**CanaryConfig**](CanaryConfig.md)| canaryConfig | 
 **configurationAccountName** | **string**| configurationAccountName | 

### Return type

[**CanaryConfigUpdateResponse**](CanaryConfigUpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCanaryConfigUsingPUT**
> CanaryConfigUpdateResponse UpdateCanaryConfigUsingPUT(ctx, canaryConfigId, canaryConfig, optional)
Update a canary config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **canaryConfigId** | **string**| canaryConfigId | 
  **canaryConfig** | [**CanaryConfig**](CanaryConfig.md)| canaryConfig | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **canaryConfigId** | **string**| canaryConfigId | 
 **canaryConfig** | [**CanaryConfig**](CanaryConfig.md)| canaryConfig | 
 **configurationAccountName** | **string**| configurationAccountName | 

### Return type

[**CanaryConfigUpdateResponse**](CanaryConfigUpdateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

