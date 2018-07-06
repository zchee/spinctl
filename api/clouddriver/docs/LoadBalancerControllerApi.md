# \LoadBalancerControllerApi

All URIs are relative to *https://spinnaker-clouddriver.meteoros.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListUsingGET5**](LoadBalancerControllerApi.md#ListUsingGET5) | **Get** /applications/{application}/loadBalancers | list


# **ListUsingGET5**
> []LoadBalancer ListUsingGET5(ctx, application)
list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **application** | **string**| application | 

### Return type

[**[]LoadBalancer**](LoadBalancer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

