# \ServerGroupManagerControllerApi

All URIs are relative to *https://localhost:8084*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetServerGroupManagersForApplicationUsingGET**](ServerGroupManagerControllerApi.md#GetServerGroupManagersForApplicationUsingGET) | **Get** /applications/{application}/serverGroupManagers | Retrieve a list of server group managers for an application


# **GetServerGroupManagersForApplicationUsingGET**
> []interface{} GetServerGroupManagersForApplicationUsingGET(ctx, application)
Retrieve a list of server group managers for an application

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **application** | **string**| application | 

### Return type

[**[]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

