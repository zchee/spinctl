# \PipelineTemplatesControllerApi

All URIs are relative to *http://localhost*

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

> CreateUsingPOST(ctx).PipelineTemplate(pipelineTemplate).Execute()

Create a pipeline template.

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pipelineTemplate** | **map[string]interface{}** | pipelineTemplate | 

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

> map[string]map[string]interface{} DeleteUsingDELETE(ctx, id).Application(application).Execute()

Delete a pipeline template.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUsingDELETERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **application** | **string** | application | 

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

> map[string]map[string]interface{} GetUsingGET(ctx, id).Execute()

Get a pipeline template.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> []map[string]interface{} ListPipelineTemplateDependentsUsingGET(ctx, id).Recursive(recursive).Execute()

List all pipelines that implement a pipeline template

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPipelineTemplateDependentsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recursive** | **bool** | recursive | 

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

> []map[string]interface{} ListUsingGET(ctx).Scopes(scopes).Execute()

List pipeline templates.

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scopes** | [**[]string**](string.md) | scopes | 

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

> map[string]map[string]interface{} ResolveTemplatesUsingGET(ctx).Source(source).ExecutionId(executionId).PipelineConfigId(pipelineConfigId).Execute()

Resolve a pipeline template.

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResolveTemplatesUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **source** | **string** | source | 
 **executionId** | **string** | executionId | 
 **pipelineConfigId** | **string** | pipelineConfigId | 

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

> UpdateUsingPOST(ctx, id).PipelineTemplate(pipelineTemplate).SkipPlanDependents(skipPlanDependents).Execute()

Update a pipeline template.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pipelineTemplate** | **map[string]interface{}** | pipelineTemplate | 

 **skipPlanDependents** | **bool** | skipPlanDependents | [default to false]

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

