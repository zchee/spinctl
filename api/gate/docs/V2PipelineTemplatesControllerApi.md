# \V2PipelineTemplatesControllerApi

All URIs are relative to *http://localhost:8084*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUsingPOST1**](V2PipelineTemplatesControllerApi.md#CreateUsingPOST1) | **Post** /v2/pipelineTemplates/create | (ALPHA) Create a pipeline template.
[**DeleteUsingDELETE1**](V2PipelineTemplatesControllerApi.md#DeleteUsingDELETE1) | **Delete** /v2/pipelineTemplates/{id} | Delete a pipeline template.
[**GetUsingGET2**](V2PipelineTemplatesControllerApi.md#GetUsingGET2) | **Get** /v2/pipelineTemplates/{id} | (ALPHA) Get a pipeline template.
[**ListPipelineTemplateDependentsUsingGET1**](V2PipelineTemplatesControllerApi.md#ListPipelineTemplateDependentsUsingGET1) | **Get** /v2/pipelineTemplates/{id}/dependents | (ALPHA) List all pipelines that implement a pipeline template
[**ListUsingGET1**](V2PipelineTemplatesControllerApi.md#ListUsingGET1) | **Get** /v2/pipelineTemplates | (ALPHA) List pipeline templates.
[**PlanUsingPOST**](V2PipelineTemplatesControllerApi.md#PlanUsingPOST) | **Post** /v2/pipelineTemplates/plan | (ALPHA) Plan a pipeline template configuration.
[**UpdateUsingPOST1**](V2PipelineTemplatesControllerApi.md#UpdateUsingPOST1) | **Post** /v2/pipelineTemplates/update/{id} | (ALPHA) Update a pipeline template.



## CreateUsingPOST1

> CreateUsingPOST1(ctx, pipelineTemplate, optional)
(ALPHA) Create a pipeline template.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipelineTemplate** | **map[string]interface{}**| pipelineTemplate | 
 **optional** | ***CreateUsingPOST1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateUsingPOST1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tag** | **optional.String**| tag | 

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

> map[string]map[string]interface{} DeleteUsingDELETE1(ctx, id, optional)
Delete a pipeline template.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 
 **optional** | ***DeleteUsingDELETE1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteUsingDELETE1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tag** | **optional.String**| tag | 
 **digest** | **optional.String**| digest | 
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


## GetUsingGET2

> map[string]map[string]interface{} GetUsingGET2(ctx, id, optional)
(ALPHA) Get a pipeline template.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 
 **optional** | ***GetUsingGET2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetUsingGET2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tag** | **optional.String**| tag | 
 **digest** | **optional.String**| digest | 

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

> []map[string]interface{} ListPipelineTemplateDependentsUsingGET1(ctx, id)
(ALPHA) List all pipelines that implement a pipeline template

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 

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

> []map[string]interface{} ListUsingGET1(ctx, optional)
(ALPHA) List pipeline templates.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ListUsingGET1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListUsingGET1Opts struct


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


## PlanUsingPOST

> map[string]map[string]interface{} PlanUsingPOST(ctx, pipeline)
(ALPHA) Plan a pipeline template configuration.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipeline** | **map[string]interface{}**| pipeline | 

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

> UpdateUsingPOST1(ctx, pipelineTemplate, id, optional)
(ALPHA) Update a pipeline template.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipelineTemplate** | **map[string]interface{}**| pipelineTemplate | 
**id** | **string**| id | 
 **optional** | ***UpdateUsingPOST1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateUsingPOST1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tag** | **optional.String**| tag | 
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

