# \ApplicationControllerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelApplicationTask**](ApplicationControllerApi.md#CancelApplicationTask) | **Put** /applications/{application}/tasks/{id}/cancel | Cancel task
[**CancelPipeline**](ApplicationControllerApi.md#CancelPipeline) | **Put** /applications/{application}/pipelines/{id}/cancel | Cancel pipeline
[**CreateApplicationTask**](ApplicationControllerApi.md#CreateApplicationTask) | **Post** /applications/{application}/tasks | Create task
[**GetApplication**](ApplicationControllerApi.md#GetApplication) | **Get** /applications/{application} | Retrieve an application&#39;s details
[**GetApplicationTask**](ApplicationControllerApi.md#GetApplicationTask) | **Get** /applications/{application}/tasks/{id} | Get task
[**GetApplicationTaskDetails**](ApplicationControllerApi.md#GetApplicationTaskDetails) | **Get** /applications/{application}/tasks/{id}/details/{taskDetailsId} | Get task details
[**GetPipelineConfig**](ApplicationControllerApi.md#GetPipelineConfig) | **Get** /applications/{application}/pipelineConfigs/{pipelineName} | Retrieve a pipeline configuration
[**GetStrategyConfig**](ApplicationControllerApi.md#GetStrategyConfig) | **Get** /applications/{application}/strategyConfigs/{strategyName} | Retrieve a pipeline strategy configuration
[**InvokePipelineConfig**](ApplicationControllerApi.md#InvokePipelineConfig) | **Post** /applications/{application}/pipelineConfigs/{pipelineName} | Invoke pipeline config
[**ListAllApplications**](ApplicationControllerApi.md#ListAllApplications) | **Get** /applications | Retrieve a list of applications
[**ListApplicationHistory**](ApplicationControllerApi.md#ListApplicationHistory) | **Get** /applications/{application}/history | Retrieve a list of an application&#39;s configuration revision history
[**ListApplicationPipelines**](ApplicationControllerApi.md#ListApplicationPipelines) | **Get** /applications/{application}/pipelines | Retrieve a list of an application&#39;s pipeline executions
[**ListPipelineConfigsForApplication**](ApplicationControllerApi.md#ListPipelineConfigsForApplication) | **Get** /applications/{application}/pipelineConfigs | Retrieve a list of an application&#39;s pipeline configurations
[**ListStrategyConfigsForApplication**](ApplicationControllerApi.md#ListStrategyConfigsForApplication) | **Get** /applications/{application}/strategyConfigs | Retrieve a list of an application&#39;s pipeline strategy configurations
[**ListTasks**](ApplicationControllerApi.md#ListTasks) | **Get** /applications/{application}/tasks | Retrieve a list of an application&#39;s tasks



## CancelApplicationTask

> map[string]map[string]interface{} CancelApplicationTask(ctx, application, id).Execute()

Cancel task

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string** | application | 
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelApplicationTaskRequest struct via the builder pattern


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


## CancelPipeline

> map[string]map[string]interface{} CancelPipeline(ctx, application, id).Reason(reason).Execute()

Cancel pipeline

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string** | application | 
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelPipelineRequest struct via the builder pattern


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


## CreateApplicationTask

> map[string]map[string]interface{} CreateApplicationTask(ctx, application).Map_(map_).Execute()

Create task

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string** | application | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationTaskRequest struct via the builder pattern


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


## GetApplication

> map[string]map[string]interface{} GetApplication(ctx, application).Expand(expand).Execute()

Retrieve an application's details

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string** | application | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationRequest struct via the builder pattern


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


## GetApplicationTask

> map[string]map[string]interface{} GetApplicationTask(ctx, application, id).Execute()

Get task

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string** | application | 
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationTaskRequest struct via the builder pattern


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


## GetApplicationTaskDetails

> map[string]map[string]interface{} GetApplicationTaskDetails(ctx, application, id, taskDetailsId).XRateLimitApp(xRateLimitApp).Execute()

Get task details

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string** | application | 
**id** | **string** | id | 
**taskDetailsId** | **string** | taskDetailsId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationTaskDetailsRequest struct via the builder pattern


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


## GetPipelineConfig

> map[string]map[string]interface{} GetPipelineConfig(ctx, application, pipelineName).Execute()

Retrieve a pipeline configuration

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string** | application | 
**pipelineName** | **string** | pipelineName | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelineConfigRequest struct via the builder pattern


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


## GetStrategyConfig

> map[string]map[string]interface{} GetStrategyConfig(ctx, application, strategyName).Execute()

Retrieve a pipeline strategy configuration

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string** | application | 
**strategyName** | **string** | strategyName | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStrategyConfigRequest struct via the builder pattern


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


## InvokePipelineConfig

> HttpEntity InvokePipelineConfig(ctx, application, pipelineName).Trigger(trigger).User(user).Execute()

Invoke pipeline config

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string** | application | 
**pipelineName** | **string** | pipelineName | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvokePipelineConfigRequest struct via the builder pattern


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


## ListAllApplications

> []map[string]interface{} ListAllApplications(ctx).Account(account).Owner(owner).Execute()

Retrieve a list of applications

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAllApplicationsRequest struct via the builder pattern


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


## ListApplicationHistory

> []map[string]interface{} ListApplicationHistory(ctx, application).Limit(limit).Execute()

Retrieve a list of an application's configuration revision history

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string** | application | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationHistoryRequest struct via the builder pattern


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


## ListApplicationPipelines

> []map[string]interface{} ListApplicationPipelines(ctx, application).Expand(expand).Limit(limit).Statuses(statuses).Execute()

Retrieve a list of an application's pipeline executions

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string** | application | 

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationPipelinesRequest struct via the builder pattern


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


## ListPipelineConfigsForApplication

> []map[string]interface{} ListPipelineConfigsForApplication(ctx, application).Execute()

Retrieve a list of an application's pipeline configurations

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string** | application | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPipelineConfigsForApplicationRequest struct via the builder pattern


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


## ListStrategyConfigsForApplication

> []map[string]interface{} ListStrategyConfigsForApplication(ctx, application).Execute()

Retrieve a list of an application's pipeline strategy configurations

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string** | application | 

### Other Parameters

Other parameters are passed through a pointer to a apiListStrategyConfigsForApplicationRequest struct via the builder pattern


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


## ListTasks

> []map[string]interface{} ListTasks(ctx, application).Limit(limit).Page(page).Statuses(statuses).Execute()

Retrieve a list of an application's tasks

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string** | application | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTasksRequest struct via the builder pattern


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

