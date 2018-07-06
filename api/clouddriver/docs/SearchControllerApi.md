# \SearchControllerApi

All URIs are relative to *https://spinnaker-clouddriver.meteoros.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchUsingGET**](SearchControllerApi.md#SearchUsingGET) | **Get** /search | search


# **SearchUsingGET**
> []SearchResultSet SearchUsingGET(ctx, type_, optional)
search

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **type_** | [**[]string**](string.md)| type | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**[]string**](string.md)| type | 
 **q** | **string**| q | 
 **platform** | **string**| platform | 
 **pageSize** | **int32**| pageSize | [default to 10000]
 **page** | **int32**| page | [default to 1]

### Return type

[**[]SearchResultSet**](SearchResultSet.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

