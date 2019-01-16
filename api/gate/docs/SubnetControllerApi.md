# \SubnetControllerApi

All URIs are relative to *https://localhost:8084*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllByCloudProviderUsingGET1**](SubnetControllerApi.md#AllByCloudProviderUsingGET1) | **Get** /subnets/{cloudProvider} | Retrieve a list of subnets for a given cloud provider


# **AllByCloudProviderUsingGET1**
> []HashMap AllByCloudProviderUsingGET1(ctx, cloudProvider, optional)
Retrieve a list of subnets for a given cloud provider

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cloudProvider** | **string**| cloudProvider | 
 **optional** | ***AllByCloudProviderUsingGET1Opts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllByCloudProviderUsingGET1Opts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRateLimitApp** | **optional.String**| X-RateLimit-App | 

### Return type

[**[]HashMap**](HashMap.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

