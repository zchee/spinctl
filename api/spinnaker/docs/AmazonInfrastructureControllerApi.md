# \AmazonInfrastructureControllerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplicationFunctions**](AmazonInfrastructureControllerApi.md#ApplicationFunctions) | **Get** /applications/{application}/functions | Get application functions
[**Functions**](AmazonInfrastructureControllerApi.md#Functions) | **Get** /functions | Get functions
[**InstanceTypes**](AmazonInfrastructureControllerApi.md#InstanceTypes) | **Get** /instanceTypes | Get instance types
[**Subnets**](AmazonInfrastructureControllerApi.md#Subnets) | **Get** /subnets | Get subnets
[**Vpcs**](AmazonInfrastructureControllerApi.md#Vpcs) | **Get** /vpcs | Get VPCs



## ApplicationFunctions

> []map[string]interface{} ApplicationFunctions(ctx, application).Execute()

Get application functions

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string** | application | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplicationFunctionsRequest struct via the builder pattern


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


## Functions

> []map[string]interface{} Functions(ctx).Account(account).FunctionName(functionName).Region(region).Execute()

Get functions

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFunctionsRequest struct via the builder pattern


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


## InstanceTypes

> []map[string]interface{} InstanceTypes(ctx).Execute()

Get instance types

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInstanceTypesRequest struct via the builder pattern


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


## Subnets

> []map[string]interface{} Subnets(ctx).Execute()

Get subnets

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSubnetsRequest struct via the builder pattern


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


## Vpcs

> []map[string]interface{} Vpcs(ctx).Execute()

Get VPCs

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVpcsRequest struct via the builder pattern


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

