# \PipelineControllerApi

All URIs are relative to *https://spinnaker-front50.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchUpdateUsingPOST3**](PipelineControllerApi.md#BatchUpdateUsingPOST3) | **Post** /pipelines/batchUpdate | batchUpdate
[**DeleteUsingDELETE3**](PipelineControllerApi.md#DeleteUsingDELETE3) | **Delete** /pipelines/{application}/{pipeline} | delete
[**GetHistoryUsingGET1**](PipelineControllerApi.md#GetHistoryUsingGET1) | **Get** /pipelines/{id}/history | getHistory
[**ListByApplicationUsingGET1**](PipelineControllerApi.md#ListByApplicationUsingGET1) | **Get** /pipelines/{application} | listByApplication
[**ListUsingGET2**](PipelineControllerApi.md#ListUsingGET2) | **Get** /pipelines | list
[**SaveUsingPOST1**](PipelineControllerApi.md#SaveUsingPOST1) | **Post** /pipelines | save
[**UpdateUsingPUT**](PipelineControllerApi.md#UpdateUsingPUT) | **Put** /pipelines/{id} | update


# **BatchUpdateUsingPOST3**
> BatchUpdateUsingPOST3(ctx, pipelines)
batchUpdate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **pipelines** | [**[]Pipeline**](Pipeline.md)| pipelines | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUsingDELETE3**
> DeleteUsingDELETE3(ctx, application, pipeline)
delete

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **application** | **string**| application | 
  **pipeline** | **string**| pipeline | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHistoryUsingGET1**
> []Pipeline GetHistoryUsingGET1(ctx, id, optional)
getHistory

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| id | 
 **limit** | **int32**| limit | [default to 20]

### Return type

[**[]Pipeline**](Pipeline.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListByApplicationUsingGET1**
> []Pipeline ListByApplicationUsingGET1(ctx, application, optional)
listByApplication

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **application** | **string**| application | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **application** | **string**| application | 
 **refresh** | **bool**| refresh | [default to true]

### Return type

[**[]Pipeline**](Pipeline.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUsingGET2**
> []Pipeline ListUsingGET2(ctx, optional)
list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restricted** | **bool**| restricted | [default to true]
 **refresh** | **bool**| refresh | [default to true]

### Return type

[**[]Pipeline**](Pipeline.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SaveUsingPOST1**
> SaveUsingPOST1(ctx, pipeline)
save

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **pipeline** | [**interface{}**](interface{}.md)| pipeline | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUsingPUT**
> map[string]interface{} UpdateUsingPUT(ctx, id, pipeline)
update

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| id | 
  **pipeline** | [**interface{}**](interface{}.md)| pipeline | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

