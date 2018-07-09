# \MetricSetListControllerApi

All URIs are relative to *https://spinnaker-kayenta.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMetricSetListUsingDELETE**](MetricSetListControllerApi.md#DeleteMetricSetListUsingDELETE) | **Delete** /metricSetList/{metricSetListId} | Delete a metric set list
[**ListAllMetricSetListsUsingGET**](MetricSetListControllerApi.md#ListAllMetricSetListsUsingGET) | **Get** /metricSetList | Retrieve a list of metric set list ids and timestamps
[**LoadMetricSetListUsingGET**](MetricSetListControllerApi.md#LoadMetricSetListUsingGET) | **Get** /metricSetList/{metricSetListId} | Retrieve a metric set list from object storage
[**StoreMetricSetListUsingPOST**](MetricSetListControllerApi.md#StoreMetricSetListUsingPOST) | **Post** /metricSetList | Write a metric set list to object storage


# **DeleteMetricSetListUsingDELETE**
> DeleteMetricSetListUsingDELETE(ctx, metricSetListId, optional)
Delete a metric set list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **metricSetListId** | **string**| metricSetListId | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metricSetListId** | **string**| metricSetListId | 
 **accountName** | **string**| accountName | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAllMetricSetListsUsingGET**
> []Mapstringobject ListAllMetricSetListsUsingGET(ctx, optional)
Retrieve a list of metric set list ids and timestamps

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

# **LoadMetricSetListUsingGET**
> []MetricSet LoadMetricSetListUsingGET(ctx, metricSetListId, optional)
Retrieve a metric set list from object storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **metricSetListId** | **string**| metricSetListId | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metricSetListId** | **string**| metricSetListId | 
 **accountName** | **string**| accountName | 

### Return type

[**[]MetricSet**](MetricSet.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StoreMetricSetListUsingPOST**
> interface{} StoreMetricSetListUsingPOST(ctx, metricSetList, optional)
Write a metric set list to object storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **metricSetList** | [**[]MetricSet**](MetricSet.md)| metricSetList | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metricSetList** | [**[]MetricSet**](MetricSet.md)| metricSetList | 
 **accountName** | **string**| accountName | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

