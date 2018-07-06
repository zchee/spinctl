# \CacheControllerApi

All URIs are relative to *https://spinnaker-clouddriver.meteoros.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HandleOnDemandUsingPOST**](CacheControllerApi.md#HandleOnDemandUsingPOST) | **Post** /cache/{cloudProvider}/{type} | handleOnDemand
[**PendingOnDemandsUsingGET**](CacheControllerApi.md#PendingOnDemandsUsingGET) | **Get** /cache/{cloudProvider}/{type} | pendingOnDemands


# **HandleOnDemandUsingPOST**
> ResponseEntity HandleOnDemandUsingPOST(ctx, cloudProvider, type_, data)
handleOnDemand

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **cloudProvider** | **string**| cloudProvider | 
  **type_** | **string**| type | 
  **data** | [**interface{}**](interface{}.md)| data | 

### Return type

[**ResponseEntity**](ResponseEntity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PendingOnDemandsUsingGET**
> []ModelMap PendingOnDemandsUsingGET(ctx, cloudProvider, type_, optional)
pendingOnDemands

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **cloudProvider** | **string**| cloudProvider | 
  **type_** | **string**| type | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudProvider** | **string**| cloudProvider | 
 **type_** | **string**| type | 
 **id** | **string**| id | 

### Return type

[**[]ModelMap**](Map.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

