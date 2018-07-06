# \InstanceControllerApi

All URIs are relative to *https://spinnaker-clouddriver.meteoros.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConsoleOutputUsingGET**](InstanceControllerApi.md#GetConsoleOutputUsingGET) | **Get** /instances/{account}/{region}/{id}/console | getConsoleOutput
[**GetInstanceUsingGET**](InstanceControllerApi.md#GetInstanceUsingGET) | **Get** /instances/{account}/{region}/{id} | getInstance


# **GetConsoleOutputUsingGET**
> interface{} GetConsoleOutputUsingGET(ctx, account, region, id, optional)
getConsoleOutput

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **account** | **string**| account | 
  **region** | **string**| region | 
  **id** | **string**| id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string**| account | 
 **region** | **string**| region | 
 **id** | **string**| id | 
 **provider** | **string**| provider | 
 **cloudProvider** | **string**| cloudProvider | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInstanceUsingGET**
> Instance GetInstanceUsingGET(ctx, account, region, id)
getInstance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **account** | **string**| account | 
  **region** | **string**| region | 
  **id** | **string**| id | 

### Return type

[**Instance**](Instance.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

