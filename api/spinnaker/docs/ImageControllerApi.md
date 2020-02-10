# \ImageControllerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindImages**](ImageControllerApi.md#FindImages) | **Get** /images/find | Retrieve a list of images, filtered by cloud provider, region, and account
[**FindTags**](ImageControllerApi.md#FindTags) | **Get** /images/tags | Find tags
[**GetImageDetails**](ImageControllerApi.md#GetImageDetails) | **Get** /images/{account}/{region}/{imageId} | Get image details



## FindImages

> []map[string]interface{} FindImages(ctx).Account(account).Count(count).Provider(provider).Q(q).Region(region).Execute()

Retrieve a list of images, filtered by cloud provider, region, and account



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | account | 
 **count** | **int32** | count | 
 **provider** | **string** | provider | [default to aws]
 **q** | **string** | q | 
 **region** | **string** | region | 

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


## FindTags

> []map[string]interface{} FindTags(ctx).Account(account).Repository(repository).XRateLimitApp(xRateLimitApp).Provider(provider).Execute()

Find tags

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | account | 
 **repository** | **string** | repository | 
 **xRateLimitApp** | **string** | X-RateLimit-App | 
 **provider** | **string** | provider | [default to aws]

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


## GetImageDetails

> []map[string]interface{} GetImageDetails(ctx, account, imageId, region).XRateLimitApp(xRateLimitApp).Provider(provider).Execute()

Get image details

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | account | 
**imageId** | **string** | imageId | 
**region** | **string** | region | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xRateLimitApp** | **string** | X-RateLimit-App | 
 **provider** | **string** | provider | [default to aws]

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

