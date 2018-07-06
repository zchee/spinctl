# \NotificationControllerApi

All URIs are relative to *https://spinnaker-front50.meteoros.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchUpdateUsingPOST2**](NotificationControllerApi.md#BatchUpdateUsingPOST2) | **Post** /notifications/batchUpdate | batchUpdate
[**DeleteUsingDELETE2**](NotificationControllerApi.md#DeleteUsingDELETE2) | **Delete** /notifications/{type}/{name} | delete
[**GetGlobalUsingGET**](NotificationControllerApi.md#GetGlobalUsingGET) | **Get** /notifications/global | getGlobal
[**ListByApplicationUsingGET**](NotificationControllerApi.md#ListByApplicationUsingGET) | **Get** /notifications/{type}/{name} | listByApplication
[**ListUsingGET1**](NotificationControllerApi.md#ListUsingGET1) | **Get** /notifications | list
[**SaveGlobalUsingPOST**](NotificationControllerApi.md#SaveGlobalUsingPOST) | **Post** /notifications/global | saveGlobal
[**SaveUsingPOST**](NotificationControllerApi.md#SaveUsingPOST) | **Post** /notifications/{type}/{name} | save


# **BatchUpdateUsingPOST2**
> BatchUpdateUsingPOST2(ctx, notifications)
batchUpdate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **notifications** | [**[]Notification**](Notification.md)| notifications | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUsingDELETE2**
> DeleteUsingDELETE2(ctx, type_, name)
delete

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **type_** | **string**| type | 
  **name** | **string**| name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGlobalUsingGET**
> map[string]interface{} GetGlobalUsingGET(ctx, )
getGlobal

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListByApplicationUsingGET**
> map[string]interface{} ListByApplicationUsingGET(ctx, type_, name)
listByApplication

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **type_** | **string**| type | 
  **name** | **string**| name | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUsingGET1**
> []Notification ListUsingGET1(ctx, )
list

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Notification**](Notification.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SaveGlobalUsingPOST**
> SaveGlobalUsingPOST(ctx, notification)
saveGlobal

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **notification** | [**interface{}**](interface{}.md)| notification | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SaveUsingPOST**
> SaveUsingPOST(ctx, type_, name, notification)
save

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **type_** | **string**| type | 
  **name** | **string**| name | 
  **notification** | [**interface{}**](interface{}.md)| notification | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

