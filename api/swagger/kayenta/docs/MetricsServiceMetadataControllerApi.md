# \MetricsServiceMetadataControllerApi

All URIs are relative to *https://spinnaker-kayenta.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListMetadataUsingGET**](MetricsServiceMetadataControllerApi.md#ListMetadataUsingGET) | **Get** /metadata/metricsService | Retrieve a list of descriptors for use in populating the canary config ui


# **ListMetadataUsingGET**
> []ModelMap ListMetadataUsingGET(ctx, optional)
Retrieve a list of descriptors for use in populating the canary config ui

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metricsAccountName** | **string**| metricsAccountName | 
 **filter** | **string**| filter | 

### Return type

[**[]ModelMap**](Map.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

