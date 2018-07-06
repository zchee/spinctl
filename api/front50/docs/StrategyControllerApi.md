# \StrategyControllerApi

All URIs are relative to *https://spinnaker-front50.meteoros.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchUpdateUsingPOST5**](StrategyControllerApi.md#BatchUpdateUsingPOST5) | **Post** /strategies/batchUpdate | batchUpdate
[**DeleteUsingDELETE5**](StrategyControllerApi.md#DeleteUsingDELETE5) | **Delete** /strategies/{application}/{strategy} | delete
[**GetHistoryUsingGET2**](StrategyControllerApi.md#GetHistoryUsingGET2) | **Get** /strategies/{id}/history | getHistory
[**ListByApplicationUsingGET2**](StrategyControllerApi.md#ListByApplicationUsingGET2) | **Get** /strategies/{application} | listByApplication
[**ListUsingGET3**](StrategyControllerApi.md#ListUsingGET3) | **Get** /strategies | list
[**SaveUsingPOST2**](StrategyControllerApi.md#SaveUsingPOST2) | **Post** /strategies | save
[**UpdateUsingPUT1**](StrategyControllerApi.md#UpdateUsingPUT1) | **Put** /strategies/{id} | update


# **BatchUpdateUsingPOST5**
> BatchUpdateUsingPOST5(ctx, strategies)
batchUpdate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **strategies** | [**[]Pipeline**](Pipeline.md)| strategies | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUsingDELETE5**
> DeleteUsingDELETE5(ctx, application, strategy)
delete

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **application** | **string**| application | 
  **strategy** | **string**| strategy | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHistoryUsingGET2**
> []Pipeline GetHistoryUsingGET2(ctx, id, optional)
getHistory

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| id | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string**| id | 
 **limit** | **int32**| limit | [default to 20]

### Return type

[**[]Pipeline**](Pipeline.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListByApplicationUsingGET2**
> []Pipeline ListByApplicationUsingGET2(ctx, application)
listByApplication

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **application** | **string**| application | 

### Return type

[**[]Pipeline**](Pipeline.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUsingGET3**
> []Pipeline ListUsingGET3(ctx, )
list

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Pipeline**](Pipeline.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SaveUsingPOST2**
> SaveUsingPOST2(ctx, strategy)
save

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **strategy** | [**interface{}**](interface{}.md)| strategy | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUsingPUT1**
> map[string]interface{} UpdateUsingPUT1(ctx, id, strategy)
update

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| id | 
  **strategy** | [**interface{}**](interface{}.md)| strategy | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

