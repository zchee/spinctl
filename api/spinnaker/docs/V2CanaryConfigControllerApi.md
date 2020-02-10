# \V2CanaryConfigControllerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCanaryConfig**](V2CanaryConfigControllerApi.md#CreateCanaryConfig) | **Post** /v2/canaryConfig | Create a canary configuration
[**DeleteCanaryConfig**](V2CanaryConfigControllerApi.md#DeleteCanaryConfig) | **Delete** /v2/canaryConfig/{id} | Delete a canary configuration
[**GetCanaryConfig**](V2CanaryConfigControllerApi.md#GetCanaryConfig) | **Get** /v2/canaryConfig/{id} | Retrieve a canary configuration by id
[**ListCanaryConfigs**](V2CanaryConfigControllerApi.md#ListCanaryConfigs) | **Get** /v2/canaryConfig | Retrieve a list of canary configurations
[**UpdateCanaryConfig**](V2CanaryConfigControllerApi.md#UpdateCanaryConfig) | **Put** /v2/canaryConfig/{id} | Update a canary configuration



## CreateCanaryConfig

> map[string]interface{} CreateCanaryConfig(ctx).Config(config).ConfigurationAccountName(configurationAccountName).Execute()

Create a canary configuration

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCanaryConfigRequest struct via the builder pattern


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


## DeleteCanaryConfig

> DeleteCanaryConfig(ctx, id).ConfigurationAccountName(configurationAccountName).Execute()

Delete a canary configuration

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCanaryConfigRequest struct via the builder pattern


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


## GetCanaryConfig

> map[string]interface{} GetCanaryConfig(ctx, id).ConfigurationAccountName(configurationAccountName).Execute()

Retrieve a canary configuration by id

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCanaryConfigRequest struct via the builder pattern


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


## ListCanaryConfigs

> []map[string]interface{} ListCanaryConfigs(ctx).Application(application).ConfigurationAccountName(configurationAccountName).Execute()

Retrieve a list of canary configurations

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCanaryConfigsRequest struct via the builder pattern


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


## UpdateCanaryConfig

> map[string]interface{} UpdateCanaryConfig(ctx, id).Config(config).ConfigurationAccountName(configurationAccountName).Execute()

Update a canary configuration

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCanaryConfigRequest struct via the builder pattern


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

