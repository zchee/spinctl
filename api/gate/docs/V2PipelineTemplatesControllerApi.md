# \V2PipelineTemplatesControllerApi

All URIs are relative to *https://localhost:8084*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUsingPOST1**](V2PipelineTemplatesControllerApi.md#CreateUsingPOST1) | **Post** /v2/pipelineTemplates | Create a pipeline template.
[**DeleteUsingDELETE1**](V2PipelineTemplatesControllerApi.md#DeleteUsingDELETE1) | **Delete** /v2/pipelineTemplates/{id} | Delete a pipeline template.
[**GetUsingGET2**](V2PipelineTemplatesControllerApi.md#GetUsingGET2) | **Get** /v2/pipelineTemplates/{id} | Get a pipeline template.
[**ListPipelineTemplateDependentsUsingGET1**](V2PipelineTemplatesControllerApi.md#ListPipelineTemplateDependentsUsingGET1) | **Get** /v2/pipelineTemplates/{id}/dependents | List all pipelines that implement a pipeline template
[**ListUsingGET1**](V2PipelineTemplatesControllerApi.md#ListUsingGET1) | **Get** /v2/pipelineTemplates | List pipeline templates.
[**PlanUsingPOST**](V2PipelineTemplatesControllerApi.md#PlanUsingPOST) | **Post** /v2/pipelineTemplates/plan | Plan a pipeline template configuration.
[**UpdateUsingPOST1**](V2PipelineTemplatesControllerApi.md#UpdateUsingPOST1) | **Post** /v2/pipelineTemplates/{id} | Update a pipeline template.


# **CreateUsingPOST1**
> CreateUsingPOST1(ctx, pipelineTemplate)
Create a pipeline template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pipelineTemplate** | [**interface{}**](interface{}.md)| pipelineTemplate | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUsingDELETE1**
> map[string]interface{} DeleteUsingDELETE1(ctx, id, optional)
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

 **application** | **optional.String**| application | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsingGET2**
> map[string]interface{} GetUsingGET2(ctx, id)
Get a pipeline template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| id | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPipelineTemplateDependentsUsingGET1**
> []interface{} ListPipelineTemplateDependentsUsingGET1(ctx, id)
List all pipelines that implement a pipeline template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| id | 

### Return type

[**[]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUsingGET1**
> []interface{} ListUsingGET1(ctx, optional)
List pipeline templates.

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

[**[]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlanUsingPOST**
> map[string]interface{} PlanUsingPOST(ctx, pipeline)
Plan a pipeline template configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **pipeline** | [**interface{}**](interface{}.md)| pipeline | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUsingPOST1**
> UpdateUsingPOST1(ctx, id, pipelineTemplate, optional)
Update a pipeline template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| id | 
  **pipelineTemplate** | [**interface{}**](interface{}.md)| pipelineTemplate | 
 **optional** | ***UpdateUsingPOST1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateUsingPOST1Opts struct

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

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

