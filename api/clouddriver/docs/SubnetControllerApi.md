# \SubnetControllerApi

All URIs are relative to *https://spinnaker-clouddriver.meteoros.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListByCloudProviderUsingGET1**](SubnetControllerApi.md#ListByCloudProviderUsingGET1) | **Get** /subnets/{cloudProvider} | listByCloudProvider
[**ListUsingGET9**](SubnetControllerApi.md#ListUsingGET9) | **Get** /subnets | list


# **ListByCloudProviderUsingGET1**
> []Subnet ListByCloudProviderUsingGET1(ctx, cloudProvider)
listByCloudProvider

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **cloudProvider** | **string**| cloudProvider | 

### Return type

[**[]Subnet**](Subnet.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUsingGET9**
> []Subnet ListUsingGET9(ctx, )
list

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Subnet**](Subnet.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

