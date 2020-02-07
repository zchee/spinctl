# \ArtifactControllerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllUsingGET**](ArtifactControllerApi.md#AllUsingGET) | **Get** /artifacts/credentials | Retrieve the list of artifact accounts configured in Clouddriver.
[**ArtifactVersionsUsingGET**](ArtifactControllerApi.md#ArtifactVersionsUsingGET) | **Get** /artifacts/account/{accountName}/versions | Retrieve the list of artifact versions by account and artifact names
[**GetArtifactUsingGET**](ArtifactControllerApi.md#GetArtifactUsingGET) | **Get** /artifacts/{provider}/{packageName}/{version} | Retrieve the specified artifact version for an artifact provider and package name



## AllUsingGET

> []map[string]interface{} AllUsingGET(ctx).XRateLimitApp(xRateLimitApp).Execute()

Retrieve the list of artifact accounts configured in Clouddriver.

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAllUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRateLimitApp** | **string** | X-RateLimit-App | 

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


## ArtifactVersionsUsingGET

> []string ArtifactVersionsUsingGET(ctx, accountName).ArtifactName(artifactName).Type_(type_).XRateLimitApp(xRateLimitApp).Execute()

Retrieve the list of artifact versions by account and artifact names

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** | accountName | 

### Other Parameters

Other parameters are passed through a pointer to a apiArtifactVersionsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **artifactName** | **string** | artifactName | 
 **type_** | **string** | type | 
 **xRateLimitApp** | **string** | X-RateLimit-App | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArtifactUsingGET

> map[string]interface{} GetArtifactUsingGET(ctx, packageName, provider, version).Execute()

Retrieve the specified artifact version for an artifact provider and package name

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageName** | **string** | packageName | 
**provider** | **string** | provider | 
**version** | **string** | version | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetArtifactUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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

