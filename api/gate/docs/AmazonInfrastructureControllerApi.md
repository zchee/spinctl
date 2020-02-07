# \AmazonInfrastructureControllerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplicationFunctionsUsingGET**](AmazonInfrastructureControllerApi.md#ApplicationFunctionsUsingGET) | **Get** /applications/{application}/functions | Get application functions
[**FunctionsUsingGET**](AmazonInfrastructureControllerApi.md#FunctionsUsingGET) | **Get** /functions | Get functions
[**InstanceTypesUsingGET**](AmazonInfrastructureControllerApi.md#InstanceTypesUsingGET) | **Get** /instanceTypes | Get instance types
[**SubnetsUsingGET**](AmazonInfrastructureControllerApi.md#SubnetsUsingGET) | **Get** /subnets | Get subnets
[**VpcsUsingGET**](AmazonInfrastructureControllerApi.md#VpcsUsingGET) | **Get** /vpcs | Get VPCs



## ApplicationFunctionsUsingGET

> []map[string]interface{} ApplicationFunctionsUsingGET(ctx, application).Execute()

Get application functions

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string** | application | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationFunctionsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## FunctionsUsingGET

> []map[string]interface{} FunctionsUsingGET(ctx).Account(account).FunctionName(functionName).Region(region).Execute()

Get functions

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFunctionsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | account | 
 **functionName** | **string** | functionName | 
 **region** | **string** | region | 

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


## InstanceTypesUsingGET

> []map[string]interface{} InstanceTypesUsingGET(ctx).Execute()

Get instance types

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInstanceTypesUsingGETRequest struct via the builder pattern


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


## SubnetsUsingGET

> []map[string]interface{} SubnetsUsingGET(ctx).Execute()

Get subnets

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSubnetsUsingGETRequest struct via the builder pattern


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


## VpcsUsingGET

> []map[string]interface{} VpcsUsingGET(ctx).Execute()

Get VPCs

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVpcsUsingGETRequest struct via the builder pattern


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

