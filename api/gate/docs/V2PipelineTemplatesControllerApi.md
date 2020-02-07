# \V2PipelineTemplatesControllerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUsingPOST1**](V2PipelineTemplatesControllerApi.md#CreateUsingPOST1) | **Post** /v2/pipelineTemplates/create | (ALPHA) Create a pipeline template.
[**DeleteUsingDELETE1**](V2PipelineTemplatesControllerApi.md#DeleteUsingDELETE1) | **Delete** /v2/pipelineTemplates/{id} | Delete a pipeline template.
[**GetUsingGET2**](V2PipelineTemplatesControllerApi.md#GetUsingGET2) | **Get** /v2/pipelineTemplates/{id} | (ALPHA) Get a pipeline template.
[**ListPipelineTemplateDependentsUsingGET1**](V2PipelineTemplatesControllerApi.md#ListPipelineTemplateDependentsUsingGET1) | **Get** /v2/pipelineTemplates/{id}/dependents | (ALPHA) List all pipelines that implement a pipeline template
[**ListUsingGET1**](V2PipelineTemplatesControllerApi.md#ListUsingGET1) | **Get** /v2/pipelineTemplates | (ALPHA) List pipeline templates.
[**ListVersionsUsingGET**](V2PipelineTemplatesControllerApi.md#ListVersionsUsingGET) | **Get** /v2/pipelineTemplates/versions | List pipeline templates with versions
[**PlanUsingPOST**](V2PipelineTemplatesControllerApi.md#PlanUsingPOST) | **Post** /v2/pipelineTemplates/plan | (ALPHA) Plan a pipeline template configuration.
[**UpdateUsingPOST1**](V2PipelineTemplatesControllerApi.md#UpdateUsingPOST1) | **Post** /v2/pipelineTemplates/update/{id} | (ALPHA) Update a pipeline template.



## CreateUsingPOST1

> CreateUsingPOST1(ctx).PipelineTemplate(pipelineTemplate).Tag(tag).Execute()

(ALPHA) Create a pipeline template.

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUsingPOST1Request struct via the builder pattern


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


## DeleteUsingDELETE1

> map[string]map[string]interface{} DeleteUsingDELETE1(ctx, id).Application(application).Digest(digest).Tag(tag).Execute()

Delete a pipeline template.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUsingDELETE1Request struct via the builder pattern


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


## GetUsingGET2

> map[string]map[string]interface{} GetUsingGET2(ctx, id).Digest(digest).Tag(tag).Execute()

(ALPHA) Get a pipeline template.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsingGET2Request struct via the builder pattern


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


## ListPipelineTemplateDependentsUsingGET1

> []map[string]interface{} ListPipelineTemplateDependentsUsingGET1(ctx, id).Execute()

(ALPHA) List all pipelines that implement a pipeline template

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPipelineTemplateDependentsUsingGET1Request struct via the builder pattern


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


## ListUsingGET1

> []map[string]interface{} ListUsingGET1(ctx).Scopes(scopes).Execute()

(ALPHA) List pipeline templates.

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsingGET1Request struct via the builder pattern


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


## ListVersionsUsingGET

> map[string]interface{} ListVersionsUsingGET(ctx).Scopes(scopes).Execute()

List pipeline templates with versions

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVersionsUsingGETRequest struct via the builder pattern


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


## PlanUsingPOST

> map[string]map[string]interface{} PlanUsingPOST(ctx).Pipeline(pipeline).Execute()

(ALPHA) Plan a pipeline template configuration.

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlanUsingPOSTRequest struct via the builder pattern


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


## UpdateUsingPOST1

> UpdateUsingPOST1(ctx, id).PipelineTemplate(pipelineTemplate).SkipPlanDependents(skipPlanDependents).Tag(tag).Execute()

(ALPHA) Update a pipeline template.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUsingPOST1Request struct via the builder pattern


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

