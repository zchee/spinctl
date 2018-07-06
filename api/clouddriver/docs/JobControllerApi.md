# \JobControllerApi

All URIs are relative to *https://spinnaker-clouddriver.meteoros.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelJobUsingDELETE**](JobControllerApi.md#CancelJobUsingDELETE) | **Delete** /applications/{application}/jobs/{account}/{location}/{id} | Collect a JobStatus
[**CollectJobUsingPOST**](JobControllerApi.md#CollectJobUsingPOST) | **Post** /applications/{application}/jobs/{account}/{location}/{id} | Collect a JobStatus
[**GetFileContentsUsingGET**](JobControllerApi.md#GetFileContentsUsingGET) | **Get** /applications/{application}/jobs/{account}/{location}/{id}/{fileName} | Collect a file from a job


# **CancelJobUsingDELETE**
> CancelJobUsingDELETE(ctx, application, account, location, id)
Collect a JobStatus

Collects the output of the job, may modify the job.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **application** | **string**| Application name | 
  **account** | **string**| Account job was created by | 
  **location** | **string**| Namespace, region, or zone job is running in | 
  **id** | **string**| Unique identifier of job being looked up | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CollectJobUsingPOST**
> JobStatus CollectJobUsingPOST(ctx, application, account, location, id)
Collect a JobStatus

Collects the output of the job, may modify the job.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **application** | **string**| Application name | 
  **account** | **string**| Account job was created by | 
  **location** | **string**| Namespace, region, or zone job is running in | 
  **id** | **string**| Unique identifier of job being looked up | 

### Return type

[**JobStatus**](JobStatus.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFileContentsUsingGET**
> interface{} GetFileContentsUsingGET(ctx, application, account, location, id, fileName)
Collect a file from a job

Collects the file result of a job.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **application** | **string**| Application name | 
  **account** | **string**| Account job was created by | 
  **location** | **string**| Namespace, region, or zone job is running in | 
  **id** | **string**| Unique identifier of job being looked up | 
  **fileName** | **string**| File name to look up | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

