# \ServerGroupManagerControllerApi

All URIs are relative to *http://localhost:8084*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetServerGroupManagersForApplicationUsingGET**](ServerGroupManagerControllerApi.md#GetServerGroupManagersForApplicationUsingGET) | **Get** /applications/{application}/serverGroupManagers | Retrieve a list of server group managers for an application



## GetServerGroupManagersForApplicationUsingGET

> []map[string]interface{} GetServerGroupManagersForApplicationUsingGET(ctx, application)

Retrieve a list of server group managers for an application

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string**| application | 

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

