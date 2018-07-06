# \DatadogFetchControllerApi

All URIs are relative to *https://spinnaker-kayenta.meteoros.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryMetricsUsingPOST**](DatadogFetchControllerApi.md#QueryMetricsUsingPOST) | **Post** /fetch/datadog/query | queryMetrics


# **QueryMetricsUsingPOST**
> interface{} QueryMetricsUsingPOST(ctx, optional)
queryMetrics

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metricsAccountName** | **string**| metricsAccountName | 
 **storageAccountName** | **string**| storageAccountName | 
 **metricSetName** | **string**| metricSetName | [default to cpu]
 **metricName** | **string**| metricName | [default to avg:system.cpu.user]
 **scope** | **string**| The scope of the Datadog query. e.g. autoscaling_group:myapp-prod-v002 | 
 **start** | **string**| An ISO format timestamp, e.g.: 2018-03-15T01:23:45Z | 
 **end** | **string**| An ISO format timestamp, e.g.: 2018-03-15T01:23:45Z | 
 **step** | **int64**| seconds | [default to 60]

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

