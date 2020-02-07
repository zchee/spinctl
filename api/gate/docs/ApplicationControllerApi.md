# \ApplicationControllerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelPipelineUsingPUT**](ApplicationControllerApi.md#CancelPipelineUsingPUT) | **Put** /applications/{application}/pipelines/{id}/cancel | Cancel pipeline
[**CancelTaskUsingPUT**](ApplicationControllerApi.md#CancelTaskUsingPUT) | **Put** /applications/{application}/tasks/{id}/cancel | Cancel task
[**GetAllApplicationsUsingGET**](ApplicationControllerApi.md#GetAllApplicationsUsingGET) | **Get** /applications | Retrieve a list of applications
[**GetApplicationHistoryUsingGET**](ApplicationControllerApi.md#GetApplicationHistoryUsingGET) | **Get** /applications/{application}/history | Retrieve a list of an application&#39;s configuration revision history
[**GetApplicationUsingGET**](ApplicationControllerApi.md#GetApplicationUsingGET) | **Get** /applications/{application} | Retrieve an application&#39;s details
[**GetPipelineConfigUsingGET**](ApplicationControllerApi.md#GetPipelineConfigUsingGET) | **Get** /applications/{application}/pipelineConfigs/{pipelineName} | Retrieve a pipeline configuration
[**GetPipelineConfigsForApplicationUsingGET**](ApplicationControllerApi.md#GetPipelineConfigsForApplicationUsingGET) | **Get** /applications/{application}/pipelineConfigs | Retrieve a list of an application&#39;s pipeline configurations
[**GetPipelinesUsingGET**](ApplicationControllerApi.md#GetPipelinesUsingGET) | **Get** /applications/{application}/pipelines | Retrieve a list of an application&#39;s pipeline executions
[**GetStrategyConfigUsingGET**](ApplicationControllerApi.md#GetStrategyConfigUsingGET) | **Get** /applications/{application}/strategyConfigs/{strategyName} | Retrieve a pipeline strategy configuration
[**GetStrategyConfigsForApplicationUsingGET**](ApplicationControllerApi.md#GetStrategyConfigsForApplicationUsingGET) | **Get** /applications/{application}/strategyConfigs | Retrieve a list of an application&#39;s pipeline strategy configurations
[**GetTaskDetailsUsingGET**](ApplicationControllerApi.md#GetTaskDetailsUsingGET) | **Get** /applications/{application}/tasks/{id}/details/{taskDetailsId} | Get task details
[**GetTaskUsingGET**](ApplicationControllerApi.md#GetTaskUsingGET) | **Get** /applications/{application}/tasks/{id} | Get task
[**GetTasksUsingGET**](ApplicationControllerApi.md#GetTasksUsingGET) | **Get** /applications/{application}/tasks | Retrieve a list of an application&#39;s tasks
[**InvokePipelineConfigUsingPOST**](ApplicationControllerApi.md#InvokePipelineConfigUsingPOST) | **Post** /applications/{application}/pipelineConfigs/{pipelineName} | Invoke pipeline config
[**TaskUsingPOST**](ApplicationControllerApi.md#TaskUsingPOST) | **Post** /applications/{application}/tasks | Create task



## CancelPipelineUsingPUT

> map[string]map[string]interface{} CancelPipelineUsingPUT(ctx, id).Reason(reason).Execute()

Cancel pipeline

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelPipelineUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reason** | **string** | reason | 

### Return type

[**map[string]map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelTaskUsingPUT

> map[string]map[string]interface{} CancelTaskUsingPUT(ctx, id).Execute()

Cancel task

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelTaskUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**map[string]map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllApplicationsUsingGET

> []map[string]interface{} GetAllApplicationsUsingGET(ctx).Account(account).Owner(owner).Execute()

Retrieve a list of applications

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllApplicationsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | filters results to only include applications deployed in the specified account | 
 **owner** | **string** | filteres results to only include applications owned by the specified email | 

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


## GetApplicationHistoryUsingGET

> []map[string]interface{} GetApplicationHistoryUsingGET(ctx, application).Limit(limit).Execute()

Retrieve a list of an application's configuration revision history

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string** | application | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationHistoryUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | limit | [default to 20]

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


## GetApplicationUsingGET

> map[string]map[string]interface{} GetApplicationUsingGET(ctx, application).Expand(expand).Execute()

Retrieve an application's details

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string** | application | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **bool** | expand | [default to true]

### Return type

[**map[string]map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPipelineConfigUsingGET

> map[string]map[string]interface{} GetPipelineConfigUsingGET(ctx, application, pipelineName).Execute()

Retrieve a pipeline configuration

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string** | application | 
**pipelineName** | **string** | pipelineName | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelineConfigUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**map[string]map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPipelineConfigsForApplicationUsingGET

> []map[string]interface{} GetPipelineConfigsForApplicationUsingGET(ctx, application).Execute()

Retrieve a list of an application's pipeline configurations

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string** | application | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelineConfigsForApplicationUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetPipelinesUsingGET

> []map[string]interface{} GetPipelinesUsingGET(ctx, application).Expand(expand).Limit(limit).Statuses(statuses).Execute()

Retrieve a list of an application's pipeline executions

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string** | application | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelinesUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **bool** | expand | 
 **limit** | **int32** | limit | 
 **statuses** | **string** | statuses | 

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


## GetStrategyConfigUsingGET

> map[string]map[string]interface{} GetStrategyConfigUsingGET(ctx, application, strategyName).Execute()

Retrieve a pipeline strategy configuration

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string** | application | 
**strategyName** | **string** | strategyName | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStrategyConfigUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**map[string]map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStrategyConfigsForApplicationUsingGET

> []map[string]interface{} GetStrategyConfigsForApplicationUsingGET(ctx, application).Execute()

Retrieve a list of an application's pipeline strategy configurations

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string** | application | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStrategyConfigsForApplicationUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetTaskDetailsUsingGET

> map[string]map[string]interface{} GetTaskDetailsUsingGET(ctx, id, taskDetailsId).XRateLimitApp(xRateLimitApp).Execute()

Get task details

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 
**taskDetailsId** | **string** | taskDetailsId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskDetailsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRateLimitApp** | **string** | X-RateLimit-App | 

### Return type

[**map[string]map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskUsingGET

> map[string]map[string]interface{} GetTaskUsingGET(ctx, id).Execute()

Get task

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**map[string]map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTasksUsingGET

> []map[string]interface{} GetTasksUsingGET(ctx, application).Limit(limit).Page(page).Statuses(statuses).Execute()

Retrieve a list of an application's tasks

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string** | application | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTasksUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | limit | 
 **page** | **int32** | page | 
 **statuses** | **string** | statuses | 

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


## InvokePipelineConfigUsingPOST

> HttpEntity InvokePipelineConfigUsingPOST(ctx, application, pipelineName).Trigger(trigger).User(user).Execute()

Invoke pipeline config

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string** | application | 
**pipelineName** | **string** | pipelineName | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvokePipelineConfigUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **trigger** | **map[string]interface{}** | trigger | 
 **user** | **string** | user | 

### Return type

[**HttpEntity**](HttpEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TaskUsingPOST

> map[string]map[string]interface{} TaskUsingPOST(ctx, application).Map_(map_).Execute()

Create task

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string** | application | 

### Other Parameters

Other parameters are passed through a pointer to a apiTaskUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **map_** | **map[string]interface{}** | map | 


### Return type

[**map[string]map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

