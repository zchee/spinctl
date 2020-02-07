# \ReorderPipelinesControllerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReorderPipelinesUsingPOST**](ReorderPipelinesControllerApi.md#ReorderPipelinesUsingPOST) | **Post** /actions/pipelines/reorder | Re-order pipelines



## ReorderPipelinesUsingPOST

> map[string]interface{} ReorderPipelinesUsingPOST(ctx).ReorderPipelinesCommand(reorderPipelinesCommand).Execute()

Re-order pipelines

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReorderPipelinesUsingPOSTRequest struct via the builder pattern


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

