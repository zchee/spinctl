# \BuildControllerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBuildMastersUsingGET**](BuildControllerApi.md#GetBuildMastersUsingGET) | **Get** /v2/builds | Get build masters
[**GetBuildUsingGET**](BuildControllerApi.md#GetBuildUsingGET) | **Get** /v2/builds/{buildMaster}/build/{number}/** | Get build for build master
[**GetBuildsUsingGET**](BuildControllerApi.md#GetBuildsUsingGET) | **Get** /v2/builds/{buildMaster}/builds/** | Get builds for build master
[**GetJobConfigUsingGET**](BuildControllerApi.md#GetJobConfigUsingGET) | **Get** /v2/builds/{buildMaster}/jobs/** | Get job config
[**GetJobsForBuildMasterUsingGET**](BuildControllerApi.md#GetJobsForBuildMasterUsingGET) | **Get** /v2/builds/{buildMaster}/jobs | Get jobs for build master
[**V3GetBuildMastersUsingGET**](BuildControllerApi.md#V3GetBuildMastersUsingGET) | **Get** /v3/builds | Get build masters
[**V3GetBuildUsingGET**](BuildControllerApi.md#V3GetBuildUsingGET) | **Get** /v3/builds/{buildMaster}/build/{number} | Get build for build master
[**V3GetBuildsUsingGET**](BuildControllerApi.md#V3GetBuildsUsingGET) | **Get** /v3/builds/{buildMaster}/builds | Get builds for build master
[**V3GetJobConfigUsingGET**](BuildControllerApi.md#V3GetJobConfigUsingGET) | **Get** /v3/builds/{buildMaster}/job | Get job config
[**V3GetJobsForBuildMasterUsingGET**](BuildControllerApi.md#V3GetJobsForBuildMasterUsingGET) | **Get** /v3/builds/{buildMaster}/jobs | Get jobs for build master



## GetBuildMastersUsingGET

> []map[string]interface{} GetBuildMastersUsingGET(ctx).Type_(type_).Execute()

Get build masters



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBuildMastersUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | type | 

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


## GetBuildUsingGET

> map[string]map[string]interface{} GetBuildUsingGET(ctx, buildMaster, number).Execute()

Get build for build master



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildMaster** | **string** | buildMaster | 
**number** | **string** | number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuildUsingGETRequest struct via the builder pattern


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


## GetBuildsUsingGET

> []map[string]interface{} GetBuildsUsingGET(ctx, buildMaster).Execute()

Get builds for build master



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildMaster** | **string** | buildMaster | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuildsUsingGETRequest struct via the builder pattern


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


## GetJobConfigUsingGET

> map[string]map[string]interface{} GetJobConfigUsingGET(ctx, buildMaster).Execute()

Get job config



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildMaster** | **string** | buildMaster | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobConfigUsingGETRequest struct via the builder pattern


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


## GetJobsForBuildMasterUsingGET

> []map[string]interface{} GetJobsForBuildMasterUsingGET(ctx, buildMaster).Execute()

Get jobs for build master



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildMaster** | **string** | buildMaster | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobsForBuildMasterUsingGETRequest struct via the builder pattern


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


## V3GetBuildMastersUsingGET

> []map[string]interface{} V3GetBuildMastersUsingGET(ctx).Type_(type_).Execute()

Get build masters

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3GetBuildMastersUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | type | 

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


## V3GetBuildUsingGET

> map[string]map[string]interface{} V3GetBuildUsingGET(ctx, buildMaster, number).Job(job).Execute()

Get build for build master

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildMaster** | **string** | buildMaster | 
**number** | **string** | number | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetBuildUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **job** | **string** | job | 


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


## V3GetBuildsUsingGET

> []map[string]interface{} V3GetBuildsUsingGET(ctx, buildMaster).Job(job).Execute()

Get builds for build master

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildMaster** | **string** | buildMaster | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetBuildsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **job** | **string** | job | 

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


## V3GetJobConfigUsingGET

> map[string]map[string]interface{} V3GetJobConfigUsingGET(ctx, buildMaster).Job(job).Execute()

Get job config

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildMaster** | **string** | buildMaster | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetJobConfigUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **job** | **string** | job | 

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


## V3GetJobsForBuildMasterUsingGET

> []map[string]interface{} V3GetJobsForBuildMasterUsingGET(ctx, buildMaster).Execute()

Get jobs for build master

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildMaster** | **string** | buildMaster | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetJobsForBuildMasterUsingGETRequest struct via the builder pattern


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

