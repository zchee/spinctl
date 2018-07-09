# \PipelineControllerApi

All URIs are relative to *https://spinnaker-kayenta.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelUsingPUT**](PipelineControllerApi.md#CancelUsingPUT) | **Put** /pipelines/{executionId}/cancel | Cancel a pipeline execution
[**GetPipelineUsingGET**](PipelineControllerApi.md#GetPipelineUsingGET) | **Get** /pipelines/{executionId} | Retrieve a pipeline execution
[**StartUsingPOST**](PipelineControllerApi.md#StartUsingPOST) | **Post** /pipelines/start | Initiate a pipeline execution


# **CancelUsingPUT**
> CancelUsingPUT(ctx, executionId)
Cancel a pipeline execution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **executionId** | **string**| executionId | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPipelineUsingGET**
> Execution GetPipelineUsingGET(ctx, executionId)
Retrieve a pipeline execution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **executionId** | **string**| executionId | 

### Return type

[**Execution**](Execution.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartUsingPOST**
> string StartUsingPOST(ctx, map_)
Initiate a pipeline execution

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **map_** | [**interface{}**](interface{}.md)| map | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

