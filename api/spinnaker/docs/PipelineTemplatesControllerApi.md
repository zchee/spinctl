# \PipelineTemplatesControllerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePipelineTemplate**](PipelineTemplatesControllerApi.md#CreatePipelineTemplate) | **Post** /pipelineTemplates | Create a pipeline template.
[**DeletePipelineTemplate**](PipelineTemplatesControllerApi.md#DeletePipelineTemplate) | **Delete** /pipelineTemplates/{id} | Delete a pipeline template.
[**GetPipelineTemplate**](PipelineTemplatesControllerApi.md#GetPipelineTemplate) | **Get** /pipelineTemplates/{id} | Get a pipeline template.
[**ListPipelineTemplateDependents**](PipelineTemplatesControllerApi.md#ListPipelineTemplateDependents) | **Get** /pipelineTemplates/{id}/dependents | List all pipelines that implement a pipeline template
[**ListPipelineTemplates**](PipelineTemplatesControllerApi.md#ListPipelineTemplates) | **Get** /pipelineTemplates | List pipeline templates.
[**ResolveTemplates**](PipelineTemplatesControllerApi.md#ResolveTemplates) | **Get** /pipelineTemplates/resolve | Resolve a pipeline template.
[**UpdatePipelineTemplate**](PipelineTemplatesControllerApi.md#UpdatePipelineTemplate) | **Post** /pipelineTemplates/{id} | Update a pipeline template.



## CreatePipelineTemplate

> CreatePipelineTemplate(ctx).PipelineTemplate(pipelineTemplate).Execute()

Create a pipeline template.

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePipelineTemplateRequest struct via the builder pattern


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


## DeletePipelineTemplate

> map[string]map[string]interface{} DeletePipelineTemplate(ctx, id).Application(application).Execute()

Delete a pipeline template.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePipelineTemplateRequest struct via the builder pattern


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


## GetPipelineTemplate

> map[string]map[string]interface{} GetPipelineTemplate(ctx, id).Execute()

Get a pipeline template.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelineTemplateRequest struct via the builder pattern


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


## ListPipelineTemplateDependents

> []map[string]interface{} ListPipelineTemplateDependents(ctx, id).Recursive(recursive).Execute()

List all pipelines that implement a pipeline template

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPipelineTemplateDependentsRequest struct via the builder pattern


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


## ListPipelineTemplates

> []map[string]interface{} ListPipelineTemplates(ctx).Scopes(scopes).Execute()

List pipeline templates.

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPipelineTemplatesRequest struct via the builder pattern


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


## ResolveTemplates

> map[string]map[string]interface{} ResolveTemplates(ctx).Source(source).ExecutionId(executionId).PipelineConfigId(pipelineConfigId).Execute()

Resolve a pipeline template.

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResolveTemplatesRequest struct via the builder pattern


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


## UpdatePipelineTemplate

> UpdatePipelineTemplate(ctx, id).PipelineTemplate(pipelineTemplate).SkipPlanDependents(skipPlanDependents).Execute()

Update a pipeline template.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePipelineTemplateRequest struct via the builder pattern


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

