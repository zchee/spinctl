# \PipelineTemplatesControllerApi

All URIs are relative to *http://localhost:8084*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUsingPOST**](PipelineTemplatesControllerApi.md#CreateUsingPOST) | **Post** /pipelineTemplates | Create a pipeline template.
[**DeleteUsingDELETE**](PipelineTemplatesControllerApi.md#DeleteUsingDELETE) | **Delete** /pipelineTemplates/{id} | Delete a pipeline template.
[**GetUsingGET**](PipelineTemplatesControllerApi.md#GetUsingGET) | **Get** /pipelineTemplates/{id} | Get a pipeline template.
[**ListPipelineTemplateDependentsUsingGET**](PipelineTemplatesControllerApi.md#ListPipelineTemplateDependentsUsingGET) | **Get** /pipelineTemplates/{id}/dependents | List all pipelines that implement a pipeline template
[**ListUsingGET**](PipelineTemplatesControllerApi.md#ListUsingGET) | **Get** /pipelineTemplates | List pipeline templates.
[**ResolveTemplatesUsingGET**](PipelineTemplatesControllerApi.md#ResolveTemplatesUsingGET) | **Get** /pipelineTemplates/resolve | Resolve a pipeline template.
[**UpdateUsingPOST**](PipelineTemplatesControllerApi.md#UpdateUsingPOST) | **Post** /pipelineTemplates/{id} | Update a pipeline template.



## CreateUsingPOST

> CreateUsingPOST(ctx, pipelineTemplate)
Create a pipeline template.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipelineTemplate** | **map[string]interface{}**| pipelineTemplate | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUsingDELETE

> map[string]map[string]interface{} DeleteUsingDELETE(ctx, id, optional)
Delete a pipeline template.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 
 **optional** | ***DeleteUsingDELETEOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteUsingDELETEOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **application** | **optional.String**| application | 

### Return type

[**map[string]map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsingGET

> map[string]map[string]interface{} GetUsingGET(ctx, id)
Get a pipeline template.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 

### Return type

[**map[string]map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPipelineTemplateDependentsUsingGET

> []map[string]interface{} ListPipelineTemplateDependentsUsingGET(ctx, id, optional)
List all pipelines that implement a pipeline template

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 
 **optional** | ***ListPipelineTemplateDependentsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListPipelineTemplateDependentsUsingGETOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recursive** | **optional.Bool**| recursive | 

### Return type

[**[]map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsingGET

> []map[string]interface{} ListUsingGET(ctx, optional)
List pipeline templates.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListUsingGETOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scopes** | [**optional.Interface of []string**](string.md)| scopes | 

### Return type

[**[]map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResolveTemplatesUsingGET

> map[string]map[string]interface{} ResolveTemplatesUsingGET(ctx, source, optional)
Resolve a pipeline template.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**source** | **string**| source | 
 **optional** | ***ResolveTemplatesUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ResolveTemplatesUsingGETOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **executionId** | **optional.String**| executionId | 
 **pipelineConfigId** | **optional.String**| pipelineConfigId | 

### Return type

[**map[string]map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUsingPOST

> UpdateUsingPOST(ctx, pipelineTemplate, id, optional)
Update a pipeline template.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipelineTemplate** | **map[string]interface{}**| pipelineTemplate | 
**id** | **string**| id | 
 **optional** | ***UpdateUsingPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateUsingPOSTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **skipPlanDependents** | **optional.Bool**| skipPlanDependents | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

