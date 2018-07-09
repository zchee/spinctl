# \ApplicationsControllerApi

All URIs are relative to *https://spinnaker-front50.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplicationsUsingGET**](ApplicationsControllerApi.md#ApplicationsUsingGET) | **Get** /v2/applications | applications
[**BatchUpdateUsingPOST**](ApplicationsControllerApi.md#BatchUpdateUsingPOST) | **Post** /v2/applications/batch/applications | batchUpdate
[**CreateUsingPOST**](ApplicationsControllerApi.md#CreateUsingPOST) | **Post** /v2/applications | create
[**DeleteUsingDELETE**](ApplicationsControllerApi.md#DeleteUsingDELETE) | **Delete** /v2/applications/{applicationName} | delete
[**GetHistoryUsingGET**](ApplicationsControllerApi.md#GetHistoryUsingGET) | **Get** /v2/applications/{applicationName}/history | getHistory
[**GetUsingGET**](ApplicationsControllerApi.md#GetUsingGET) | **Get** /v2/applications/{applicationName} | get
[**UpdateUsingPATCH**](ApplicationsControllerApi.md#UpdateUsingPATCH) | **Patch** /v2/applications/{applicationName} | update


# **ApplicationsUsingGET**
> []Application ApplicationsUsingGET(ctx, params, optional)
applications

Fetch all applications.      Supports filtering by one or more attributes:     - ?email=my@email.com     - ?email=my@email.com&name=flex

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **params** | **string**| params | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **params** | **string**| params | 
 **pageSize** | **int32**| pageSize | 

### Return type

[**[]Application**](Application.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BatchUpdateUsingPOST**
> BatchUpdateUsingPOST(ctx, applications)
batchUpdate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **applications** | [**[]Application**](Application.md)| applications | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUsingPOST**
> Application CreateUsingPOST(ctx, app)
create

Create an application

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **app** | [**Application**](Application.md)| app | 

### Return type

[**Application**](Application.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUsingDELETE**
> DeleteUsingDELETE(ctx, applicationName)
delete

Delete an application

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **applicationName** | **string**| applicationName | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHistoryUsingGET**
> []Application GetHistoryUsingGET(ctx, applicationName, optional)
getHistory

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **applicationName** | **string**| applicationName | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationName** | **string**| applicationName | 
 **limit** | **int32**| limit | [default to 20]

### Return type

[**[]Application**](Application.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUsingGET**
> Application GetUsingGET(ctx, applicationName)
get

Fetch a single application by name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **applicationName** | **string**| applicationName | 

### Return type

[**Application**](Application.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUsingPATCH**
> Application UpdateUsingPATCH(ctx, applicationName, app)
update

Update an existing application

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **applicationName** | **string**| applicationName | 
  **app** | [**Application**](Application.md)| app | 

### Return type

[**Application**](Application.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

