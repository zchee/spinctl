# \ApplicationsControllerApi

All URIs are relative to *https://spinnaker-clouddriver.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUsingGET**](ApplicationsControllerApi.md#GetUsingGET) | **Get** /applications/{name} | get
[**ListUsingGET**](ApplicationsControllerApi.md#ListUsingGET) | **Get** /applications | list


# **GetUsingGET**
> ApplicationViewModel GetUsingGET(ctx, name)
get

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| name | 

### Return type

[**ApplicationViewModel**](ApplicationViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUsingGET**
> []Application ListUsingGET(ctx, optional)
list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expand** | **bool**| expand | [default to true]
 **restricted** | **bool**| restricted | [default to true]

### Return type

[**[]Application**](Application.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

