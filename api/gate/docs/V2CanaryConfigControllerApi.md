# \V2CanaryConfigControllerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCanaryConfigUsingPOST**](V2CanaryConfigControllerApi.md#CreateCanaryConfigUsingPOST) | **Post** /v2/canaryConfig | Create a canary configuration
[**DeleteCanaryConfigUsingDELETE**](V2CanaryConfigControllerApi.md#DeleteCanaryConfigUsingDELETE) | **Delete** /v2/canaryConfig/{id} | Delete a canary configuration
[**GetCanaryConfigUsingGET**](V2CanaryConfigControllerApi.md#GetCanaryConfigUsingGET) | **Get** /v2/canaryConfig/{id} | Retrieve a canary configuration by id
[**GetCanaryConfigsUsingGET**](V2CanaryConfigControllerApi.md#GetCanaryConfigsUsingGET) | **Get** /v2/canaryConfig | Retrieve a list of canary configurations
[**UpdateCanaryConfigUsingPUT**](V2CanaryConfigControllerApi.md#UpdateCanaryConfigUsingPUT) | **Put** /v2/canaryConfig/{id} | Update a canary configuration



## CreateCanaryConfigUsingPOST

> map[string]interface{} CreateCanaryConfigUsingPOST(ctx).Config(config).ConfigurationAccountName(configurationAccountName).Execute()

Create a canary configuration

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCanaryConfigUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **config** | **map[string]interface{}** | config | 
 **configurationAccountName** | **string** | configurationAccountName | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCanaryConfigUsingDELETE

> DeleteCanaryConfigUsingDELETE(ctx, id).ConfigurationAccountName(configurationAccountName).Execute()

Delete a canary configuration

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCanaryConfigUsingDELETERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **configurationAccountName** | **string** | configurationAccountName | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCanaryConfigUsingGET

> map[string]interface{} GetCanaryConfigUsingGET(ctx, id).ConfigurationAccountName(configurationAccountName).Execute()

Retrieve a canary configuration by id

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCanaryConfigUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **configurationAccountName** | **string** | configurationAccountName | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCanaryConfigsUsingGET

> []map[string]interface{} GetCanaryConfigsUsingGET(ctx).Application(application).ConfigurationAccountName(configurationAccountName).Execute()

Retrieve a list of canary configurations

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCanaryConfigsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **application** | **string** | application | 
 **configurationAccountName** | **string** | configurationAccountName | 

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


## UpdateCanaryConfigUsingPUT

> map[string]interface{} UpdateCanaryConfigUsingPUT(ctx, id).Config(config).ConfigurationAccountName(configurationAccountName).Execute()

Update a canary configuration

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCanaryConfigUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **config** | **map[string]interface{}** | config | 

 **configurationAccountName** | **string** | configurationAccountName | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

