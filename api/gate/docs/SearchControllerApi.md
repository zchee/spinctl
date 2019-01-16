# \SearchControllerApi

All URIs are relative to *https://localhost:8084*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchUsingGET**](SearchControllerApi.md#SearchUsingGET) | **Get** /search | Search infrastructure


# **SearchUsingGET**
> []HashMap SearchUsingGET(ctx, type_, optional)
Search infrastructure

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**| type | 
 **optional** | ***SearchUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **optional.String**| q | 
 **platform** | **optional.String**| platform | 
 **pageSize** | **optional.Int32**| pageSize | [default to 10000]
 **page** | **optional.Int32**| page | [default to 1]
 **allowShortQuery** | **optional.Bool**| allowShortQuery | [default to false]
 **xRateLimitApp** | **optional.String**| X-RateLimit-App | 

### Return type

[**[]HashMap**](HashMap.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

