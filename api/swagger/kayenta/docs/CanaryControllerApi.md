# \CanaryControllerApi

All URIs are relative to *https://spinnaker-kayenta.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCanaryResultsByApplicationUsingGET**](CanaryControllerApi.md#GetCanaryResultsByApplicationUsingGET) | **Get** /canary/executions | Retrieve a list of an application&#39;s canary results
[**GetCanaryResultsUsingGET**](CanaryControllerApi.md#GetCanaryResultsUsingGET) | **Get** /canary/{canaryExecutionId} | Retrieve status and results for a canary run
[**InitiateCanaryUsingPOST**](CanaryControllerApi.md#InitiateCanaryUsingPOST) | **Post** /canary/{canaryConfigId} | Initiate a canary pipeline
[**InitiateCanaryWithConfigUsingPOST**](CanaryControllerApi.md#InitiateCanaryWithConfigUsingPOST) | **Post** /canary | Initiate a canary pipeline with CanaryConfig provided


# **GetCanaryResultsByApplicationUsingGET**
> []CanaryExecutionStatusResponse GetCanaryResultsByApplicationUsingGET(ctx, optional)
Retrieve a list of an application's canary results

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **application** | **string**| application | 
 **limit** | **int32**| limit | [default to 20]
 **statuses** | **string**| statuses | 
 **storageAccountName** | **string**| storageAccountName | 

### Return type

[**[]CanaryExecutionStatusResponse**](CanaryExecutionStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCanaryResultsUsingGET**
> CanaryExecutionStatusResponse GetCanaryResultsUsingGET(ctx, canaryExecutionId, optional)
Retrieve status and results for a canary run

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **canaryExecutionId** | **string**| canaryExecutionId | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **canaryExecutionId** | **string**| canaryExecutionId | 
 **storageAccountName** | **string**| storageAccountName | 

### Return type

[**CanaryExecutionStatusResponse**](CanaryExecutionStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InitiateCanaryUsingPOST**
> CanaryExecutionResponse InitiateCanaryUsingPOST(ctx, canaryConfigId, optional)
Initiate a canary pipeline

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
 **application** | **string**| application | 
 **parentPipelineExecutionId** | **string**| parentPipelineExecutionId | 
 **metricsAccountName** | **string**| metricsAccountName | 
 **configurationAccountName** | **string**| configurationAccountName | 
 **storageAccountName** | **string**| storageAccountName | 
 **canaryExecutionRequest** | [**CanaryExecutionRequest**](CanaryExecutionRequest.md)| canaryExecutionRequest | 

### Return type

[**CanaryExecutionResponse**](CanaryExecutionResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InitiateCanaryWithConfigUsingPOST**
> CanaryExecutionResponse InitiateCanaryWithConfigUsingPOST(ctx, optional)
Initiate a canary pipeline with CanaryConfig provided

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metricsAccountName** | **string**| metricsAccountName | 
 **storageAccountName** | **string**| storageAccountName | 
 **canaryAdhocExecutionRequest** | [**CanaryAdhocExecutionRequest**](CanaryAdhocExecutionRequest.md)| canaryAdhocExecutionRequest | 

### Return type

[**CanaryExecutionResponse**](CanaryExecutionResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

