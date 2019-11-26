# \JobControllerApi

All URIs are relative to *http://localhost:8084*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetJobUsingGET**](JobControllerApi.md#GetJobUsingGET) | **Get** /applications/{applicationName}/jobs/{account}/{region}/{name} | Get job



## GetJobUsingGET

> map[string]map[string]interface{} GetJobUsingGET(ctx, account, applicationName, name, region, optional)

Get job

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string**| account | 
**applicationName** | **string**| applicationName | 
**name** | **string**| name | 
**region** | **string**| region | 
 **optional** | ***GetJobUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetJobUsingGETOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **xRateLimitApp** | **optional.String**| X-RateLimit-App | 
 **expand** | **optional.String**| expand | [default to false]

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

