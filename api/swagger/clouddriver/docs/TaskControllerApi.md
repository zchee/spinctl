# \TaskControllerApi

All URIs are relative to *https://spinnaker-clouddriver.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUsingGET2**](TaskControllerApi.md#GetUsingGET2) | **Get** /task/{id} | get
[**ListUsingGET10**](TaskControllerApi.md#ListUsingGET10) | **Get** /task | list


# **GetUsingGET2**
> Task GetUsingGET2(ctx, id)
get

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| id | 

### Return type

[**Task**](Task.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUsingGET10**
> []Task ListUsingGET10(ctx, )
list

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Task**](Task.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

