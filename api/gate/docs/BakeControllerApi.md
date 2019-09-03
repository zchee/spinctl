# \BakeControllerApi

All URIs are relative to *http://localhost:8084*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BakeOptionsUsingGET**](BakeControllerApi.md#BakeOptionsUsingGET) | **Get** /bakery/options/{cloudProvider} | Retrieve a list of available bakery base images for a given cloud provider
[**BakeOptionsUsingGET1**](BakeControllerApi.md#BakeOptionsUsingGET1) | **Get** /bakery/options | Retrieve a list of available bakery base images, grouped by cloud provider
[**LookupLogsUsingGET**](BakeControllerApi.md#LookupLogsUsingGET) | **Get** /bakery/logs/{region}/{statusId} | Retrieve the logs for a given bake



## BakeOptionsUsingGET

> map[string]interface{} BakeOptionsUsingGET(ctx, cloudProvider)
Retrieve a list of available bakery base images for a given cloud provider

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cloudProvider** | **string**| cloudProvider | 

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


## BakeOptionsUsingGET1

> map[string]interface{} BakeOptionsUsingGET1(ctx, )
Retrieve a list of available bakery base images, grouped by cloud provider

### Required Parameters

This endpoint does not need any parameter.

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


## LookupLogsUsingGET

> map[string]interface{} LookupLogsUsingGET(ctx, region, statusId)
Retrieve the logs for a given bake

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string**| region | 
**statusId** | **string**| statusId | 

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

