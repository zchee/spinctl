# \ImageControllerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindImagesUsingGET**](ImageControllerApi.md#FindImagesUsingGET) | **Get** /images/find | Retrieve a list of images, filtered by cloud provider, region, and account
[**FindTagsUsingGET**](ImageControllerApi.md#FindTagsUsingGET) | **Get** /images/tags | Find tags
[**GetImageDetailsUsingGET**](ImageControllerApi.md#GetImageDetailsUsingGET) | **Get** /images/{account}/{region}/{imageId} | Get image details



## FindImagesUsingGET

> []map[string]interface{} FindImagesUsingGET(ctx).Account(account).Count(count).Provider(provider).Q(q).Region(region).Execute()

Retrieve a list of images, filtered by cloud provider, region, and account



### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindImagesUsingGETRequest struct via the builder pattern


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


## FindTagsUsingGET

> []map[string]interface{} FindTagsUsingGET(ctx).Account(account).Repository(repository).XRateLimitApp(xRateLimitApp).Provider(provider).Execute()

Find tags

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindTagsUsingGETRequest struct via the builder pattern


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


## GetImageDetailsUsingGET

> []map[string]interface{} GetImageDetailsUsingGET(ctx, account, imageId, region).XRateLimitApp(xRateLimitApp).Provider(provider).Execute()

Get image details

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | account | 
**imageId** | **string** | imageId | 
**region** | **string** | region | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageDetailsUsingGETRequest struct via the builder pattern


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

