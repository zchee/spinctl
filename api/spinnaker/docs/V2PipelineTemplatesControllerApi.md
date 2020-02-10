# \V2PipelineTemplatesControllerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePipelineTemplateV2**](V2PipelineTemplatesControllerApi.md#CreatePipelineTemplateV2) | **Post** /v2/pipelineTemplates/create | (ALPHA) Create a pipeline template.
[**DeletePipelineTemplateV2**](V2PipelineTemplatesControllerApi.md#DeletePipelineTemplateV2) | **Delete** /v2/pipelineTemplates/{id} | Delete a pipeline template.
[**GetPipelineTemplateV2**](V2PipelineTemplatesControllerApi.md#GetPipelineTemplateV2) | **Get** /v2/pipelineTemplates/{id} | (ALPHA) Get a pipeline template.
[**ListPipelineTemplateDependentsV2**](V2PipelineTemplatesControllerApi.md#ListPipelineTemplateDependentsV2) | **Get** /v2/pipelineTemplates/{id}/dependents | (ALPHA) List all pipelines that implement a pipeline template
[**ListPipelineTemplatesV2**](V2PipelineTemplatesControllerApi.md#ListPipelineTemplatesV2) | **Get** /v2/pipelineTemplates | (ALPHA) List pipeline templates.
[**ListVersions**](V2PipelineTemplatesControllerApi.md#ListVersions) | **Get** /v2/pipelineTemplates/versions | List pipeline templates with versions
[**Plan**](V2PipelineTemplatesControllerApi.md#Plan) | **Post** /v2/pipelineTemplates/plan | (ALPHA) Plan a pipeline template configuration.
[**UpdatePipelineTemplateV2**](V2PipelineTemplatesControllerApi.md#UpdatePipelineTemplateV2) | **Post** /v2/pipelineTemplates/update/{id} | (ALPHA) Update a pipeline template.



## CreatePipelineTemplateV2

> CreatePipelineTemplateV2(ctx).PipelineTemplate(pipelineTemplate).Tag(tag).Execute()

(ALPHA) Create a pipeline template.

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePipelineTemplateV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pipelineTemplate** | **map[string]interface{}** | pipelineTemplate | 
 **tag** | **string** | tag | 

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


## DeletePipelineTemplateV2

> map[string]map[string]interface{} DeletePipelineTemplateV2(ctx, id).Application(application).Digest(digest).Tag(tag).Execute()

Delete a pipeline template.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePipelineTemplateV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **application** | **string** | application | 
 **digest** | **string** | digest | 
 **tag** | **string** | tag | 

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


## GetPipelineTemplateV2

> map[string]map[string]interface{} GetPipelineTemplateV2(ctx, id).Digest(digest).Tag(tag).Execute()

(ALPHA) Get a pipeline template.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelineTemplateV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **digest** | **string** | digest | 
 **tag** | **string** | tag | 

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


## ListPipelineTemplateDependentsV2

> []map[string]interface{} ListPipelineTemplateDependentsV2(ctx, id).Execute()

(ALPHA) List all pipelines that implement a pipeline template

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPipelineTemplateDependentsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ListPipelineTemplatesV2

> []map[string]interface{} ListPipelineTemplatesV2(ctx).Scopes(scopes).Execute()

(ALPHA) List pipeline templates.

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPipelineTemplatesV2Request struct via the builder pattern


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


## ListVersions

> map[string]interface{} ListVersions(ctx).Scopes(scopes).Execute()

List pipeline templates with versions

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scopes** | [**[]string**](string.md) | scopes | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Plan

> map[string]map[string]interface{} Plan(ctx).Pipeline(pipeline).Execute()

(ALPHA) Plan a pipeline template configuration.

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pipeline** | **map[string]interface{}** | pipeline | 

### Return type

[**map[string]map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePipelineTemplateV2

> UpdatePipelineTemplateV2(ctx, id).PipelineTemplate(pipelineTemplate).SkipPlanDependents(skipPlanDependents).Tag(tag).Execute()

(ALPHA) Update a pipeline template.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePipelineTemplateV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pipelineTemplate** | **map[string]interface{}** | pipelineTemplate | 

 **skipPlanDependents** | **bool** | skipPlanDependents | [default to false]
 **tag** | **string** | tag | 

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

