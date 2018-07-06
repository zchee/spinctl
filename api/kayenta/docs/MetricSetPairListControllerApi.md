# \MetricSetPairListControllerApi

All URIs are relative to *https://spinnaker-kayenta.meteoros.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMetricSetPairListUsingDELETE**](MetricSetPairListControllerApi.md#DeleteMetricSetPairListUsingDELETE) | **Delete** /metricSetPairList/{metricSetPairListId} | Delete a metric set pair list
[**ListAllMetricSetPairListsUsingGET**](MetricSetPairListControllerApi.md#ListAllMetricSetPairListsUsingGET) | **Get** /metricSetPairList | Retrieve a list of metric set pair list ids and timestamps
[**LoadMetricSetPairListUsingGET**](MetricSetPairListControllerApi.md#LoadMetricSetPairListUsingGET) | **Get** /metricSetPairList/{metricSetPairListId} | Retrieve a metric set pair list from object storage
[**LoadMetricSetPairUsingGET**](MetricSetPairListControllerApi.md#LoadMetricSetPairUsingGET) | **Get** /metricSetPairList/{metricSetPairListId}/{metricSetPairId} | Retrieve a single metric set pair from a metricSetPairList from object storage
[**StoreMetricSetPairListUsingPOST**](MetricSetPairListControllerApi.md#StoreMetricSetPairListUsingPOST) | **Post** /metricSetPairList | Write a metric set pair list to object storage


# **DeleteMetricSetPairListUsingDELETE**
> DeleteMetricSetPairListUsingDELETE(ctx, metricSetPairListId, optional)
Delete a metric set pair list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **metricSetPairListId** | **string**| metricSetPairListId | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metricSetPairListId** | **string**| metricSetPairListId | 
 **accountName** | **string**| accountName | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAllMetricSetPairListsUsingGET**
> []Mapstringobject ListAllMetricSetPairListsUsingGET(ctx, optional)
Retrieve a list of metric set pair list ids and timestamps

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

# **LoadMetricSetPairListUsingGET**
> []MetricSetPair LoadMetricSetPairListUsingGET(ctx, metricSetPairListId, optional)
Retrieve a metric set pair list from object storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **metricSetPairListId** | **string**| metricSetPairListId | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metricSetPairListId** | **string**| metricSetPairListId | 
 **accountName** | **string**| accountName | 

### Return type

[**[]MetricSetPair**](MetricSetPair.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadMetricSetPairUsingGET**
> MetricSetPair LoadMetricSetPairUsingGET(ctx, metricSetPairListId, metricSetPairId, optional)
Retrieve a single metric set pair from a metricSetPairList from object storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **metricSetPairListId** | **string**| metricSetPairListId | 
  **metricSetPairId** | **string**| metricSetPairId | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metricSetPairListId** | **string**| metricSetPairListId | 
 **metricSetPairId** | **string**| metricSetPairId | 
 **accountName** | **string**| accountName | 

### Return type

[**MetricSetPair**](MetricSetPair.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StoreMetricSetPairListUsingPOST**
> interface{} StoreMetricSetPairListUsingPOST(ctx, metricSetPairList, optional)
Write a metric set pair list to object storage

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **metricSetPairList** | [**[]MetricSetPair**](MetricSetPair.md)| metricSetPairList | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metricSetPairList** | [**[]MetricSetPair**](MetricSetPair.md)| metricSetPairList | 
 **accountName** | **string**| accountName | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

