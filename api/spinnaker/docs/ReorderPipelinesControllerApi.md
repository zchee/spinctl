# \ReorderPipelinesControllerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReorderPipelines**](ReorderPipelinesControllerApi.md#ReorderPipelines) | **Post** /actions/pipelines/reorder | Re-order pipelines



## ReorderPipelines

> map[string]interface{} ReorderPipelines(ctx).ReorderPipelinesCommand(reorderPipelinesCommand).Execute()

Re-order pipelines

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReorderPipelinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reorderPipelinesCommand** | [**ReorderPipelinesCommand**](ReorderPipelinesCommand.md) | reorderPipelinesCommand | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

