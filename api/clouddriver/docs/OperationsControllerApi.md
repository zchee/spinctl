# \OperationsControllerApi

All URIs are relative to *https://spinnaker-clouddriver.meteoros.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudProviderOperationUsingPOST**](OperationsControllerApi.md#CloudProviderOperationUsingPOST) | **Post** /{cloudProvider}/ops/{name} | cloudProviderOperation
[**CloudProviderOperationsUsingPOST**](OperationsControllerApi.md#CloudProviderOperationsUsingPOST) | **Post** /{cloudProvider}/ops | cloudProviderOperations
[**OperationUsingPOST**](OperationsControllerApi.md#OperationUsingPOST) | **Post** /ops/{name} | operation
[**OperationsUsingPOST**](OperationsControllerApi.md#OperationsUsingPOST) | **Post** /ops | operations


# **CloudProviderOperationUsingPOST**
> map[string]string CloudProviderOperationUsingPOST(ctx, cloudProvider, name, requestBody, optional)
cloudProviderOperation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **cloudProvider** | **string**| cloudProvider | 
  **name** | **string**| name | 
  **requestBody** | [**interface{}**](interface{}.md)| requestBody | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudProvider** | **string**| cloudProvider | 
 **name** | **string**| name | 
 **requestBody** | [**interface{}**](interface{}.md)| requestBody | 
 **clientRequestId** | **string**| clientRequestId | 

### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CloudProviderOperationsUsingPOST**
> map[string]string CloudProviderOperationsUsingPOST(ctx, cloudProvider, requestBody, optional)
cloudProviderOperations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **cloudProvider** | **string**| cloudProvider | 
  **requestBody** | [**[]MapstringMap**](Map«string,Map».md)| requestBody | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudProvider** | **string**| cloudProvider | 
 **requestBody** | [**[]MapstringMap**](Map«string,Map».md)| requestBody | 
 **clientRequestId** | **string**| clientRequestId | 

### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationUsingPOST**
> map[string]string OperationUsingPOST(ctx, name, requestBody, optional)
operation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| name | 
  **requestBody** | [**interface{}**](interface{}.md)| requestBody | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string**| name | 
 **requestBody** | [**interface{}**](interface{}.md)| requestBody | 
 **clientRequestId** | **string**| clientRequestId | 

### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OperationsUsingPOST**
> map[string]string OperationsUsingPOST(ctx, requestBody, optional)
operations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **requestBody** | [**[]MapstringMap**](Map«string,Map».md)| requestBody | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | [**[]MapstringMap**](Map«string,Map».md)| requestBody | 
 **clientRequestId** | **string**| clientRequestId | 

### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

