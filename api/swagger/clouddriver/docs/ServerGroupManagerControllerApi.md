# \ServerGroupManagerControllerApi

All URIs are relative to *https://spinnaker-clouddriver.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetForApplicationUsingGET**](ServerGroupManagerControllerApi.md#GetForApplicationUsingGET) | **Get** /applications/{application}/serverGroupManagers | getForApplication


# **GetForApplicationUsingGET**
> []ServerGroupManager GetForApplicationUsingGET(ctx, application)
getForApplication

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **application** | **string**| application | 

### Return type

[**[]ServerGroupManager**](ServerGroupManager.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

