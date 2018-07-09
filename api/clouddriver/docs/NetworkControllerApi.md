# \NetworkControllerApi

All URIs are relative to *https://spinnaker-clouddriver.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListByCloudProviderUsingGET**](NetworkControllerApi.md#ListByCloudProviderUsingGET) | **Get** /networks/{cloudProvider} | listByCloudProvider
[**ListUsingGET6**](NetworkControllerApi.md#ListUsingGET6) | **Get** /networks | list


# **ListByCloudProviderUsingGET**
> []Network ListByCloudProviderUsingGET(ctx, cloudProvider)
listByCloudProvider

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **cloudProvider** | **string**| cloudProvider | 

### Return type

[**[]Network**](Network.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUsingGET6**
> map[string][]Network ListUsingGET6(ctx, )
list

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**map[string][]Network**](array.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

