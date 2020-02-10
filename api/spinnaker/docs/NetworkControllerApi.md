# \NetworkControllerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListNetworks**](NetworkControllerApi.md#ListNetworks) | **Get** /networks | Retrieve a list of networks, grouped by cloud provider
[**ListNetworksByCloudProvider**](NetworkControllerApi.md#ListNetworksByCloudProvider) | **Get** /networks/{cloudProvider} | Retrieve a list of networks for a given cloud provider



## ListNetworks

> map[string]map[string]interface{} ListNetworks(ctx).XRateLimitApp(xRateLimitApp).Execute()

Retrieve a list of networks, grouped by cloud provider

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRateLimitApp** | **string** | X-RateLimit-App | 

### Return type

[**map[string]map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworksByCloudProvider

> []map[string]interface{} ListNetworksByCloudProvider(ctx, cloudProvider).XRateLimitApp(xRateLimitApp).Execute()

Retrieve a list of networks for a given cloud provider

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProvider** | **string** | cloudProvider | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNetworksByCloudProviderRequest struct via the builder pattern


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

