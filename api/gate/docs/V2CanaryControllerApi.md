# \V2CanaryControllerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCanaryResultUsingGET**](V2CanaryControllerApi.md#GetCanaryResultUsingGET) | **Get** /v2/canaries/canary/{canaryConfigId}/{canaryExecutionId} | (DEPRECATED) Retrieve a canary result
[**GetCanaryResultUsingGET1**](V2CanaryControllerApi.md#GetCanaryResultUsingGET1) | **Get** /v2/canaries/canary/{canaryExecutionId} | Retrieve a canary result
[**GetCanaryResultsByApplicationUsingGET**](V2CanaryControllerApi.md#GetCanaryResultsByApplicationUsingGET) | **Get** /v2/canaries/{application}/executions | Retrieve a list of an application&#39;s canary results
[**GetMetricSetPairListUsingGET**](V2CanaryControllerApi.md#GetMetricSetPairListUsingGET) | **Get** /v2/canaries/metricSetPairList/{metricSetPairListId} | Retrieve a metric set pair list
[**InitiateCanaryUsingPOST**](V2CanaryControllerApi.md#InitiateCanaryUsingPOST) | **Post** /v2/canaries/canary/{canaryConfigId} | Start a canary execution
[**InitiateCanaryWithConfigUsingPOST**](V2CanaryControllerApi.md#InitiateCanaryWithConfigUsingPOST) | **Post** /v2/canaries/canary | Start a canary execution with the supplied canary config
[**ListCredentialsUsingGET**](V2CanaryControllerApi.md#ListCredentialsUsingGET) | **Get** /v2/canaries/credentials | Retrieve a list of configured Kayenta accounts
[**ListJudgesUsingGET**](V2CanaryControllerApi.md#ListJudgesUsingGET) | **Get** /v2/canaries/judges | Retrieve a list of all configured canary judges
[**ListMetricsServiceMetadataUsingGET**](V2CanaryControllerApi.md#ListMetricsServiceMetadataUsingGET) | **Get** /v2/canaries/metadata/metricsService | Retrieve a list of descriptors for use in populating the canary config ui



## GetCanaryResultUsingGET

> map[string]interface{} GetCanaryResultUsingGET(ctx, canaryConfigId, canaryExecutionId).StorageAccountName(storageAccountName).Execute()

(DEPRECATED) Retrieve a canary result

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**canaryConfigId** | **string** | canaryConfigId | 
**canaryExecutionId** | **string** | canaryExecutionId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCanaryResultUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **storageAccountName** | **string** | storageAccountName | 

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


## GetCanaryResultUsingGET1

> map[string]interface{} GetCanaryResultUsingGET1(ctx, canaryExecutionId).StorageAccountName(storageAccountName).Execute()

Retrieve a canary result

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**canaryExecutionId** | **string** | canaryExecutionId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCanaryResultUsingGET1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storageAccountName** | **string** | storageAccountName | 

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


## GetCanaryResultsByApplicationUsingGET

> []map[string]interface{} GetCanaryResultsByApplicationUsingGET(ctx, application).Limit(limit).Statuses(statuses).StorageAccountName(storageAccountName).Execute()

Retrieve a list of an application's canary results

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string** | application | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCanaryResultsByApplicationUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | limit | 
 **statuses** | **string** | Comma-separated list of statuses, e.g.: RUNNING, SUCCEEDED, TERMINAL | 
 **storageAccountName** | **string** | storageAccountName | 

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


## GetMetricSetPairListUsingGET

> []map[string]interface{} GetMetricSetPairListUsingGET(ctx, metricSetPairListId).StorageAccountName(storageAccountName).Execute()

Retrieve a metric set pair list

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricSetPairListId** | **string** | metricSetPairListId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricSetPairListUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storageAccountName** | **string** | storageAccountName | 

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


## InitiateCanaryUsingPOST

> map[string]interface{} InitiateCanaryUsingPOST(ctx, canaryConfigId).ExecutionRequest(executionRequest).Application(application).ConfigurationAccountName(configurationAccountName).MetricsAccountName(metricsAccountName).ParentPipelineExecutionId(parentPipelineExecutionId).StorageAccountName(storageAccountName).Execute()

Start a canary execution

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**canaryConfigId** | **string** | canaryConfigId | 

### Other Parameters

Other parameters are passed through a pointer to a apiInitiateCanaryUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **executionRequest** | **map[string]interface{}** | executionRequest | 

 **application** | **string** | application | 
 **configurationAccountName** | **string** | configurationAccountName | 
 **metricsAccountName** | **string** | metricsAccountName | 
 **parentPipelineExecutionId** | **string** | parentPipelineExecutionId | 
 **storageAccountName** | **string** | storageAccountName | 

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


## InitiateCanaryWithConfigUsingPOST

> map[string]interface{} InitiateCanaryWithConfigUsingPOST(ctx).AdhocExecutionRequest(adhocExecutionRequest).Application(application).MetricsAccountName(metricsAccountName).ParentPipelineExecutionId(parentPipelineExecutionId).StorageAccountName(storageAccountName).Execute()

Start a canary execution with the supplied canary config

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInitiateCanaryWithConfigUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adhocExecutionRequest** | **map[string]interface{}** | adhocExecutionRequest | 
 **application** | **string** | application | 
 **metricsAccountName** | **string** | metricsAccountName | 
 **parentPipelineExecutionId** | **string** | parentPipelineExecutionId | 
 **storageAccountName** | **string** | storageAccountName | 

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


## ListCredentialsUsingGET

> []map[string]interface{} ListCredentialsUsingGET(ctx).Execute()

Retrieve a list of configured Kayenta accounts

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCredentialsUsingGETRequest struct via the builder pattern


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


## ListJudgesUsingGET

> []map[string]interface{} ListJudgesUsingGET(ctx).Execute()

Retrieve a list of all configured canary judges

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListJudgesUsingGETRequest struct via the builder pattern


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


## ListMetricsServiceMetadataUsingGET

> []map[string]interface{} ListMetricsServiceMetadataUsingGET(ctx).Filter(filter).MetricsAccountName(metricsAccountName).Execute()

Retrieve a list of descriptors for use in populating the canary config ui

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMetricsServiceMetadataUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | filter | 
 **metricsAccountName** | **string** | metricsAccountName | 

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

