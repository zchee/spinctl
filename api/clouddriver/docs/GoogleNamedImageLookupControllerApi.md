# \GoogleNamedImageLookupControllerApi

All URIs are relative to *https://spinnaker-clouddriver.meteoros.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListUsingGET3**](GoogleNamedImageLookupControllerApi.md#ListUsingGET3) | **Get** /gce/images/find | list


# **ListUsingGET3**
> []NamedImage ListUsingGET3(ctx, optional)
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
 **q** | **string**| q | 
 **account** | **string**| account | 

### Return type

[**[]NamedImage**](NamedImage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

