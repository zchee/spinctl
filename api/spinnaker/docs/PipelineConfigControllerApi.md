# \PipelineConfigControllerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConvertPipelineConfigToPipelineTemplate**](PipelineConfigControllerApi.md#ConvertPipelineConfigToPipelineTemplate) | **Get** /pipelineConfigs/{pipelineConfigId}/convertToTemplate | Convert a pipeline config to a pipeline template.
[**GetAllPipelineConfigs**](PipelineConfigControllerApi.md#GetAllPipelineConfigs) | **Get** /pipelineConfigs | Get all pipeline configs.
[**GetPipelineConfigHistory**](PipelineConfigControllerApi.md#GetPipelineConfigHistory) | **Get** /pipelineConfigs/{pipelineConfigId}/history | Get pipeline config history.



## ConvertPipelineConfigToPipelineTemplate

> string ConvertPipelineConfigToPipelineTemplate(ctx, pipelineConfigId).Execute()

Convert a pipeline config to a pipeline template.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipelineConfigId** | **string** | pipelineConfigId | 

### Other Parameters

Other parameters are passed through a pointer to a apiConvertPipelineConfigToPipelineTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllPipelineConfigs

> []map[string]interface{} GetAllPipelineConfigs(ctx).Execute()

Get all pipeline configs.

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllPipelineConfigsRequest struct via the builder pattern


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


## GetPipelineConfigHistory

> []map[string]interface{} GetPipelineConfigHistory(ctx, pipelineConfigId).Limit(limit).Execute()

Get pipeline config history.

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipelineConfigId** | **string** | pipelineConfigId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelineConfigHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | limit | [default to 20]

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

