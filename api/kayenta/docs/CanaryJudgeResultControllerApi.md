# \CanaryJudgeResultControllerApi

All URIs are relative to *https://spinnaker-kayenta.meteoros.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCanaryJudgeResultUsingDELETE**](CanaryJudgeResultControllerApi.md#DeleteCanaryJudgeResultUsingDELETE) | **Delete** /canaryJudgeResult/{canaryJudgeResultId} | Delete a canary judge result
[**ListAllCanaryJudgeResultsUsingGET**](CanaryJudgeResultControllerApi.md#ListAllCanaryJudgeResultsUsingGET) | **Get** /canaryJudgeResult | Retrieve a list of canary judge result ids and timestamps
[**LoadCanaryJudgeResultUsingGET**](CanaryJudgeResultControllerApi.md#LoadCanaryJudgeResultUsingGET) | **Get** /canaryJudgeResult/{canaryJudgeResultId} | Retrieve a canary judge result from object storage
[**StoreCanaryJudgeResultUsingPOST**](CanaryJudgeResultControllerApi.md#StoreCanaryJudgeResultUsingPOST) | **Post** /canaryJudgeResult | Write a canary judge result to object storage


# **DeleteCanaryJudgeResultUsingDELETE**
> DeleteCanaryJudgeResultUsingDELETE(ctx, canaryJudgeResultId, optional)
Delete a canary judge result

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **canaryJudgeResultId** | **string**| canaryJudgeResultId | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **canaryJudgeResultId** | **string**| canaryJudgeResultId | 
 **accountName** | **string**| accountName | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAllCanaryJudgeResultsUsingGET**
> []Mapstringobject ListAllCanaryJudgeResultsUsingGET(ctx, optional)
Retrieve a list of canary judge result ids and timestamps

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountName** | **string**| accountName | 

### Return type

[**[]Mapstringobject**](Map«string,object».md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadCanaryJudgeResultUsingGET**
> CanaryResult LoadCanaryJudgeResultUsingGET(ctx, canaryJudgeResultId, optional)
Retrieve a canary judge result from object storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **canaryJudgeResultId** | **string**| canaryJudgeResultId | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **canaryJudgeResultId** | **string**| canaryJudgeResultId | 
 **accountName** | **string**| accountName | 

### Return type

[**CanaryResult**](CanaryResult.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StoreCanaryJudgeResultUsingPOST**
> interface{} StoreCanaryJudgeResultUsingPOST(ctx, canaryResult, optional)
Write a canary judge result to object storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **canaryResult** | [**CanaryResult**](CanaryResult.md)| canaryResult | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **canaryResult** | [**CanaryResult**](CanaryResult.md)| canaryResult | 
 **accountName** | **string**| accountName | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

