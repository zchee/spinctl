# \ConcourseControllerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Jobs**](ConcourseControllerApi.md#Jobs) | **Get** /concourse/{buildMaster}/teams/{team}/pipelines/{pipeline}/jobs | Retrieve the list of job names for a given pipeline available to triggers
[**ListPipelines**](ConcourseControllerApi.md#ListPipelines) | **Get** /concourse/{buildMaster}/teams/{team}/pipelines | Retrieve the list of pipeline names for a given team available to triggers
[**Resources**](ConcourseControllerApi.md#Resources) | **Get** /concourse/{buildMaster}/teams/{team}/pipelines/{pipeline}/resources | Retrieve the list of resource names for a given pipeline available to the Concourse stage



## Jobs

> []map[string]interface{} Jobs(ctx, buildMaster, pipeline, team).Execute()

Retrieve the list of job names for a given pipeline available to triggers

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildMaster** | **string** | buildMaster | 
**pipeline** | **string** | pipeline | 
**team** | **string** | team | 

### Other Parameters

Other parameters are passed through a pointer to a apiJobsRequest struct via the builder pattern


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


## ListPipelines

> []map[string]interface{} ListPipelines(ctx, buildMaster, team).Execute()

Retrieve the list of pipeline names for a given team available to triggers

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildMaster** | **string** | buildMaster | 
**team** | **string** | team | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPipelinesRequest struct via the builder pattern


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


## Resources

> []map[string]interface{} Resources(ctx, buildMaster, pipeline, team).Execute()

Retrieve the list of resource names for a given pipeline available to the Concourse stage

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildMaster** | **string** | buildMaster | 
**pipeline** | **string** | pipeline | 
**team** | **string** | team | 

### Other Parameters

Other parameters are passed through a pointer to a apiResourcesRequest struct via the builder pattern


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

