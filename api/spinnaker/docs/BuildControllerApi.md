# \BuildControllerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBuild**](BuildControllerApi.md#GetBuild) | **Get** /v2/builds/{buildMaster}/build/{number}/** | Get build for build master
[**GetBuildMasters**](BuildControllerApi.md#GetBuildMasters) | **Get** /v2/builds | Get build masters
[**GetBuilds**](BuildControllerApi.md#GetBuilds) | **Get** /v2/builds/{buildMaster}/builds/** | Get builds for build master
[**GetJobConfig**](BuildControllerApi.md#GetJobConfig) | **Get** /v2/builds/{buildMaster}/jobs/** | Get job config
[**GetJobsForBuildMaster**](BuildControllerApi.md#GetJobsForBuildMaster) | **Get** /v2/builds/{buildMaster}/jobs | Get jobs for build master
[**V3GetBuild**](BuildControllerApi.md#V3GetBuild) | **Get** /v3/builds/{buildMaster}/build/{number} | Get build for build master
[**V3GetBuildMasters**](BuildControllerApi.md#V3GetBuildMasters) | **Get** /v3/builds | Get build masters
[**V3GetBuilds**](BuildControllerApi.md#V3GetBuilds) | **Get** /v3/builds/{buildMaster}/builds | Get builds for build master
[**V3GetJobConfig**](BuildControllerApi.md#V3GetJobConfig) | **Get** /v3/builds/{buildMaster}/job | Get job config
[**V3GetJobsForBuildMaster**](BuildControllerApi.md#V3GetJobsForBuildMaster) | **Get** /v3/builds/{buildMaster}/jobs | Get jobs for build master



## GetBuild

> map[string]map[string]interface{} GetBuild(ctx, buildMaster, number).Execute()

Get build for build master



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildMaster** | **string** | buildMaster | 
**number** | **string** | number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuildRequest struct via the builder pattern


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


## GetBuildMasters

> []map[string]interface{} GetBuildMasters(ctx).Type_(type_).Execute()

Get build masters



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBuildMastersRequest struct via the builder pattern


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


## GetBuilds

> []map[string]interface{} GetBuilds(ctx, buildMaster).Execute()

Get builds for build master



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildMaster** | **string** | buildMaster | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuildsRequest struct via the builder pattern


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


## GetJobConfig

> map[string]map[string]interface{} GetJobConfig(ctx, buildMaster).Execute()

Get job config



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildMaster** | **string** | buildMaster | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobConfigRequest struct via the builder pattern


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


## GetJobsForBuildMaster

> []map[string]interface{} GetJobsForBuildMaster(ctx, buildMaster).Execute()

Get jobs for build master



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildMaster** | **string** | buildMaster | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobsForBuildMasterRequest struct via the builder pattern


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


## V3GetBuild

> map[string]map[string]interface{} V3GetBuild(ctx, buildMaster, number).Job(job).Execute()

Get build for build master

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildMaster** | **string** | buildMaster | 
**number** | **string** | number | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetBuildRequest struct via the builder pattern


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


## V3GetBuildMasters

> []map[string]interface{} V3GetBuildMasters(ctx).Type_(type_).Execute()

Get build masters

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV3GetBuildMastersRequest struct via the builder pattern


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


## V3GetBuilds

> []map[string]interface{} V3GetBuilds(ctx, buildMaster).Job(job).Execute()

Get builds for build master

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildMaster** | **string** | buildMaster | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetBuildsRequest struct via the builder pattern


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


## V3GetJobConfig

> map[string]map[string]interface{} V3GetJobConfig(ctx, buildMaster).Job(job).Execute()

Get job config

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildMaster** | **string** | buildMaster | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetJobConfigRequest struct via the builder pattern


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


## V3GetJobsForBuildMaster

> []map[string]interface{} V3GetJobsForBuildMaster(ctx, buildMaster).Execute()

Get jobs for build master

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildMaster** | **string** | buildMaster | 

### Other Parameters

Other parameters are passed through a pointer to a apiV3GetJobsForBuildMasterRequest struct via the builder pattern


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

