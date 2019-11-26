# \AmazonInfrastructureControllerApi

All URIs are relative to *http://localhost:8084*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplicationFunctionsUsingGET**](AmazonInfrastructureControllerApi.md#ApplicationFunctionsUsingGET) | **Get** /applications/{application}/functions | Get application functions
[**FunctionsUsingGET**](AmazonInfrastructureControllerApi.md#FunctionsUsingGET) | **Get** /functions | Get functions
[**InstanceTypesUsingGET**](AmazonInfrastructureControllerApi.md#InstanceTypesUsingGET) | **Get** /instanceTypes | Get instance types
[**SubnetsUsingGET**](AmazonInfrastructureControllerApi.md#SubnetsUsingGET) | **Get** /subnets | Get subnets
[**VpcsUsingGET**](AmazonInfrastructureControllerApi.md#VpcsUsingGET) | **Get** /vpcs | Get VPCs



## ApplicationFunctionsUsingGET

> []map[string]interface{} ApplicationFunctionsUsingGET(ctx, application)

Get application functions

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string**| application | 

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

> []map[string]interface{} FunctionsUsingGET(ctx, optional)

Get functions

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FunctionsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FunctionsUsingGETOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **optional.String**| account | 
 **functionName** | **optional.String**| functionName | 
 **region** | **optional.String**| region | 

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

> []map[string]interface{} InstanceTypesUsingGET(ctx, )

Get instance types

### Required Parameters

This endpoint does not need any parameter.

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

> []map[string]interface{} SubnetsUsingGET(ctx, )

Get subnets

### Required Parameters

This endpoint does not need any parameter.

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

> []map[string]interface{} VpcsUsingGET(ctx, )

Get VPCs

### Required Parameters

This endpoint does not need any parameter.

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

