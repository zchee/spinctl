# \LoadBalancerControllerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllUsingGET**](LoadBalancerControllerApi.md#GetAllUsingGET) | **Get** /loadBalancers | Retrieve a list of load balancers for a given cloud provider
[**GetApplicationLoadBalancersUsingGET**](LoadBalancerControllerApi.md#GetApplicationLoadBalancersUsingGET) | **Get** /applications/{application}/loadBalancers | Retrieve a list of load balancers for a given application
[**GetLoadBalancerDetailsUsingGET**](LoadBalancerControllerApi.md#GetLoadBalancerDetailsUsingGET) | **Get** /loadBalancers/{account}/{region}/{name} | Retrieve a load balancer&#39;s details as a single element list for a given account, region, cloud provider and load balancer name
[**GetLoadBalancerUsingGET**](LoadBalancerControllerApi.md#GetLoadBalancerUsingGET) | **Get** /loadBalancers/{name} | Retrieve a load balancer for a given cloud provider



## GetAllUsingGET

> []map[string]interface{} GetAllUsingGET(ctx).XRateLimitApp(xRateLimitApp).Provider(provider).Execute()

Retrieve a list of load balancers for a given cloud provider

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRateLimitApp** | **string** | X-RateLimit-App | 
 **provider** | **string** | provider | [default to aws]

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


## GetApplicationLoadBalancersUsingGET

> []map[string]interface{} GetApplicationLoadBalancersUsingGET(ctx, application).XRateLimitApp(xRateLimitApp).Execute()

Retrieve a list of load balancers for a given application

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string** | application | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationLoadBalancersUsingGETRequest struct via the builder pattern


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


## GetLoadBalancerDetailsUsingGET

> []map[string]interface{} GetLoadBalancerDetailsUsingGET(ctx, account, name, region).XRateLimitApp(xRateLimitApp).Provider(provider).Execute()

Retrieve a load balancer's details as a single element list for a given account, region, cloud provider and load balancer name

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | account | 
**name** | **string** | name | 
**region** | **string** | region | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoadBalancerDetailsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xRateLimitApp** | **string** | X-RateLimit-App | 
 **provider** | **string** | provider | [default to aws]

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


## GetLoadBalancerUsingGET

> map[string]map[string]interface{} GetLoadBalancerUsingGET(ctx, name).XRateLimitApp(xRateLimitApp).Provider(provider).Execute()

Retrieve a load balancer for a given cloud provider

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoadBalancerUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRateLimitApp** | **string** | X-RateLimit-App | 
 **provider** | **string** | provider | [default to aws]

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

