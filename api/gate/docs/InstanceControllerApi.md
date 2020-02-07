# \InstanceControllerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConsoleOutputUsingGET**](InstanceControllerApi.md#GetConsoleOutputUsingGET) | **Get** /instances/{account}/{region}/{instanceId}/console | Retrieve an instance&#39;s console output
[**GetInstanceDetailsUsingGET**](InstanceControllerApi.md#GetInstanceDetailsUsingGET) | **Get** /instances/{account}/{region}/{instanceId} | Retrieve an instance&#39;s details



## GetConsoleOutputUsingGET

> map[string]interface{} GetConsoleOutputUsingGET(ctx, account, instanceId, region).XRateLimitApp(xRateLimitApp).Provider(provider).Execute()

Retrieve an instance's console output

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | account | 
**instanceId** | **string** | instanceId | 
**region** | **string** | region | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConsoleOutputUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xRateLimitApp** | **string** | X-RateLimit-App | 
 **provider** | **string** | provider | [default to aws]

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


## GetInstanceDetailsUsingGET

> map[string]interface{} GetInstanceDetailsUsingGET(ctx, account, instanceId, region).XRateLimitApp(xRateLimitApp).Execute()

Retrieve an instance's details

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | account | 
**instanceId** | **string** | instanceId | 
**region** | **string** | region | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceDetailsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xRateLimitApp** | **string** | X-RateLimit-App | 

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

