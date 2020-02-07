# \SubnetControllerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllByCloudProviderUsingGET1**](SubnetControllerApi.md#AllByCloudProviderUsingGET1) | **Get** /subnets/{cloudProvider} | Retrieve a list of subnets for a given cloud provider



## AllByCloudProviderUsingGET1

> []map[string]interface{} AllByCloudProviderUsingGET1(ctx, cloudProvider).XRateLimitApp(xRateLimitApp).Execute()

Retrieve a list of subnets for a given cloud provider

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProvider** | **string** | cloudProvider | 

### Other Parameters

Other parameters are passed through a pointer to a apiAllByCloudProviderUsingGET1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRateLimitApp** | **string** | X-RateLimit-App | 

### Return type

[**[]map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

