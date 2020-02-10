# \V2CanaryControllerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCanaryResult**](V2CanaryControllerApi.md#GetCanaryResult) | **Get** /v2/canaries/canary/{canaryConfigIdOrExecutionId} | Retrieve a canary result
[**GetCanaryResultDeprecated**](V2CanaryControllerApi.md#GetCanaryResultDeprecated) | **Get** /v2/canaries/canary/{canaryConfigId}/{canaryExecutionId} | (DEPRECATED) Retrieve a canary result
[**GetMetricSetPairList**](V2CanaryControllerApi.md#GetMetricSetPairList) | **Get** /v2/canaries/metricSetPairList/{metricSetPairListId} | Retrieve a metric set pair list
[**InitiateCanary**](V2CanaryControllerApi.md#InitiateCanary) | **Post** /v2/canaries/canary/{canaryConfigIdOrExecutionId} | Start a canary execution
[**InitiateCanaryWithConfig**](V2CanaryControllerApi.md#InitiateCanaryWithConfig) | **Post** /v2/canaries/canary | Start a canary execution with the supplied canary config
[**ListCanaryResultsByApplication**](V2CanaryControllerApi.md#ListCanaryResultsByApplication) | **Get** /v2/canaries/{application}/executions | Retrieve a list of an application&#39;s canary results
[**ListCredentials**](V2CanaryControllerApi.md#ListCredentials) | **Get** /v2/canaries/credentials | Retrieve a list of configured Kayenta accounts
[**ListJudges**](V2CanaryControllerApi.md#ListJudges) | **Get** /v2/canaries/judges | Retrieve a list of all configured canary judges
[**ListMetricsServiceMetadata**](V2CanaryControllerApi.md#ListMetricsServiceMetadata) | **Get** /v2/canaries/metadata/metricsService | Retrieve a list of descriptors for use in populating the canary config ui



## GetCanaryResult

> map[string]interface{} GetCanaryResult(ctx, canaryConfigIdOrExecutionId).StorageAccountName(storageAccountName).Execute()

Retrieve a canary result

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**canaryConfigIdOrExecutionId** | **string** | canaryConfigId or canaryExecutionId. In get request, should canaryExecutionId. In post request, should canaryConfigId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCanaryResultRequest struct via the builder pattern


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


## GetCanaryResultDeprecated

> map[string]interface{} GetCanaryResultDeprecated(ctx, canaryConfigId, canaryExecutionId).StorageAccountName(storageAccountName).Execute()

(DEPRECATED) Retrieve a canary result

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**canaryConfigId** | **string** | canaryConfigId | 
**canaryExecutionId** | **string** | canaryExecutionId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCanaryResultDeprecatedRequest struct via the builder pattern


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


## GetMetricSetPairList

> []map[string]interface{} GetMetricSetPairList(ctx, metricSetPairListId).StorageAccountName(storageAccountName).Execute()

Retrieve a metric set pair list

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricSetPairListId** | **string** | metricSetPairListId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricSetPairListRequest struct via the builder pattern


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


## InitiateCanary

> map[string]interface{} InitiateCanary(ctx, canaryConfigIdOrExecutionId).ExecutionRequest(executionRequest).Application(application).ConfigurationAccountName(configurationAccountName).MetricsAccountName(metricsAccountName).ParentPipelineExecutionId(parentPipelineExecutionId).StorageAccountName(storageAccountName).Execute()

Start a canary execution

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**canaryConfigIdOrExecutionId** | **string** | canaryConfigId or canaryExecutionId. In get request, should canaryExecutionId. In post request, should canaryConfigId | 

### Other Parameters

Other parameters are passed through a pointer to a apiInitiateCanaryRequest struct via the builder pattern


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


## InitiateCanaryWithConfig

> map[string]interface{} InitiateCanaryWithConfig(ctx).AdhocExecutionRequest(adhocExecutionRequest).Application(application).MetricsAccountName(metricsAccountName).ParentPipelineExecutionId(parentPipelineExecutionId).StorageAccountName(storageAccountName).Execute()

Start a canary execution with the supplied canary config

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInitiateCanaryWithConfigRequest struct via the builder pattern


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


## ListCanaryResultsByApplication

> []map[string]interface{} ListCanaryResultsByApplication(ctx, application).Limit(limit).Statuses(statuses).StorageAccountName(storageAccountName).Execute()

Retrieve a list of an application's canary results

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string** | application | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCanaryResultsByApplicationRequest struct via the builder pattern


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


## ListCredentials

> []map[string]interface{} ListCredentials(ctx).Execute()

Retrieve a list of configured Kayenta accounts

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCredentialsRequest struct via the builder pattern


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


## ListJudges

> []map[string]interface{} ListJudges(ctx).Execute()

Retrieve a list of all configured canary judges

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListJudgesRequest struct via the builder pattern


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


## ListMetricsServiceMetadata

> []map[string]interface{} ListMetricsServiceMetadata(ctx).Filter(filter).MetricsAccountName(metricsAccountName).Execute()

Retrieve a list of descriptors for use in populating the canary config ui

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMetricsServiceMetadataRequest struct via the builder pattern


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

