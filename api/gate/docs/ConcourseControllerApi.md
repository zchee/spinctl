# \ConcourseControllerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**JobsUsingGET**](ConcourseControllerApi.md#JobsUsingGET) | **Get** /concourse/{buildMaster}/teams/{team}/pipelines/{pipeline}/jobs | Retrieve the list of job names for a given pipeline available to triggers
[**PipelinesUsingGET**](ConcourseControllerApi.md#PipelinesUsingGET) | **Get** /concourse/{buildMaster}/teams/{team}/pipelines | Retrieve the list of pipeline names for a given team available to triggers
[**ResourcesUsingGET**](ConcourseControllerApi.md#ResourcesUsingGET) | **Get** /concourse/{buildMaster}/teams/{team}/pipelines/{pipeline}/resources | Retrieve the list of resource names for a given pipeline available to the Concourse stage



## JobsUsingGET

> []map[string]interface{} JobsUsingGET(ctx, buildMaster, pipeline, team).Execute()

Retrieve the list of job names for a given pipeline available to triggers

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildMaster** | **string** | buildMaster | 
**pipeline** | **string** | pipeline | 
**team** | **string** | team | 

### Other Parameters

Other parameters are passed through a pointer to a apiJobsUsingGETRequest struct via the builder pattern


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


## PipelinesUsingGET

> []map[string]interface{} PipelinesUsingGET(ctx, buildMaster, team).Execute()

Retrieve the list of pipeline names for a given team available to triggers

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildMaster** | **string** | buildMaster | 
**team** | **string** | team | 

### Other Parameters

Other parameters are passed through a pointer to a apiPipelinesUsingGETRequest struct via the builder pattern


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


## ResourcesUsingGET

> []map[string]interface{} ResourcesUsingGET(ctx, buildMaster, pipeline, team).Execute()

Retrieve the list of resource names for a given pipeline available to the Concourse stage

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildMaster** | **string** | buildMaster | 
**pipeline** | **string** | pipeline | 
**team** | **string** | team | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourcesUsingGETRequest struct via the builder pattern


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

