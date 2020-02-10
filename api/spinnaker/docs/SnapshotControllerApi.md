# \SnapshotControllerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCurrentSnapshot**](SnapshotControllerApi.md#GetCurrentSnapshot) | **Get** /applications/{application}/snapshots/{account} | Get current snapshot
[**GetSnapshotHistory**](SnapshotControllerApi.md#GetSnapshotHistory) | **Get** /applications/{application}/snapshots/{account}/history | Get snapshot history



## GetCurrentSnapshot

> map[string]map[string]interface{} GetCurrentSnapshot(ctx, account, application).Execute()

Get current snapshot

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | account | 
**application** | **string** | application | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentSnapshotRequest struct via the builder pattern


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


## GetSnapshotHistory

> []map[string]interface{} GetSnapshotHistory(ctx, account, application).Limit(limit).Execute()

Get snapshot history

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | account | 
**application** | **string** | application | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSnapshotHistoryRequest struct via the builder pattern


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

