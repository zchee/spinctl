# \TaskControllerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelTaskUsingPUT1**](TaskControllerApi.md#CancelTaskUsingPUT1) | **Put** /tasks/{id}/cancel | Cancel task
[**CancelTasksUsingPUT**](TaskControllerApi.md#CancelTasksUsingPUT) | **Put** /tasks/cancel | Cancel tasks
[**DeleteTaskUsingDELETE**](TaskControllerApi.md#DeleteTaskUsingDELETE) | **Delete** /tasks/{id} | Delete task
[**GetTaskDetailsUsingGET1**](TaskControllerApi.md#GetTaskDetailsUsingGET1) | **Get** /tasks/{id}/details/{taskDetailsId} | Get task details
[**GetTaskUsingGET1**](TaskControllerApi.md#GetTaskUsingGET1) | **Get** /tasks/{id} | Get task
[**TaskUsingPOST1**](TaskControllerApi.md#TaskUsingPOST1) | **Post** /tasks | Create task



## CancelTaskUsingPUT1

> map[string]map[string]interface{} CancelTaskUsingPUT1(ctx, id).Execute()

Cancel task

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelTaskUsingPUT1Request struct via the builder pattern


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


## CancelTasksUsingPUT

> map[string]map[string]interface{} CancelTasksUsingPUT(ctx).Ids(ids).Execute()

Cancel tasks

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelTasksUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | [**[]string**](string.md) | ids | 

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


## DeleteTaskUsingDELETE

> map[string]map[string]interface{} DeleteTaskUsingDELETE(ctx, id).Execute()

Delete task

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTaskUsingDELETERequest struct via the builder pattern


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


## GetTaskDetailsUsingGET1

> map[string]map[string]interface{} GetTaskDetailsUsingGET1(ctx, id, taskDetailsId).XRateLimitApp(xRateLimitApp).Execute()

Get task details

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 
**taskDetailsId** | **string** | taskDetailsId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskDetailsUsingGET1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRateLimitApp** | **string** | X-RateLimit-App | 

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


## GetTaskUsingGET1

> map[string]map[string]interface{} GetTaskUsingGET1(ctx, id).Execute()

Get task

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskUsingGET1Request struct via the builder pattern


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


## TaskUsingPOST1

> map[string]map[string]interface{} TaskUsingPOST1(ctx).Map_(map_).Execute()

Create task

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTaskUsingPOST1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **map_** | **map[string]interface{}** | map | 

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

