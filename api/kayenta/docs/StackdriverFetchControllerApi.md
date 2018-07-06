# \StackdriverFetchControllerApi

All URIs are relative to *https://spinnaker-kayenta.meteoros.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryMetricsUsingPOST1**](StackdriverFetchControllerApi.md#QueryMetricsUsingPOST1) | **Post** /fetch/stackdriver/query | Exercise the Stackdriver Metrics Service directly, without any orchestration or judging


# **QueryMetricsUsingPOST1**
> interface{} QueryMetricsUsingPOST1(ctx, optional)
Exercise the Stackdriver Metrics Service directly, without any orchestration or judging

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
 **metricType** | **string**| metricType | [default to compute.googleapis.com/instance/cpu/utilization]
 **groupByFields** | [**[]string**](string.md)| groupByFields | 
 **project** | **string**| project | 
 **resourceType** | **string**| Used to identify the type of the resource being queried, e.g. aws_ec2_instance, gce_instance. | 
 **location** | **string**| The location to use when scoping the query. Valid choices depend on what cloud platform the query relates to (could be a region, a namespace, or something else). | 
 **scope** | **string**| The name of the resource to use when scoping the query. The most common use-case is to provide a server group name. | 
 **startTimeIso** | **string**| An ISO format timestamp, e.g.: 2018-02-21T12:48:00Z | 
 **endTimeIso** | **string**| An ISO format timestamp, e.g.: 2018-02-21T12:51:00Z | 
 **step** | **int64**| seconds | [default to 60]
 **customFilter** | **string**| customFilter | 
 **extendedScopeParams** | [**interface{}**](interface{}.md)| extendedScopeParams | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

