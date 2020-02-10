# \ServerGroupManagerControllerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListServerGroupManagersForApplication**](ServerGroupManagerControllerApi.md#ListServerGroupManagersForApplication) | **Get** /applications/{application}/serverGroupManagers | Retrieve a list of server group managers for an application



## ListServerGroupManagersForApplication

> []map[string]interface{} ListServerGroupManagersForApplication(ctx, application).Execute()

Retrieve a list of server group managers for an application

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string** | application | 

### Other Parameters

Other parameters are passed through a pointer to a apiListServerGroupManagersForApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

