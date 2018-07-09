# \EntityTagsControllerApi

All URIs are relative to *https://spinnaker-front50.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchDeleteUsingPOST**](EntityTagsControllerApi.md#BatchDeleteUsingPOST) | **Post** /v2/tags/batchDelete | batchDelete
[**BatchUpdateUsingPOST1**](EntityTagsControllerApi.md#BatchUpdateUsingPOST1) | **Post** /v2/tags/batchUpdate | batchUpdate
[**CreateUsingPOST1**](EntityTagsControllerApi.md#CreateUsingPOST1) | **Post** /v2/tags | create
[**DeleteUsingDELETE1**](EntityTagsControllerApi.md#DeleteUsingDELETE1) | **Delete** /v2/tags/** | delete
[**TagUsingGET**](EntityTagsControllerApi.md#TagUsingGET) | **Get** /v2/tags/** | tag
[**TagsUsingGET**](EntityTagsControllerApi.md#TagsUsingGET) | **Get** /v2/tags | tags


# **BatchDeleteUsingPOST**
> BatchDeleteUsingPOST(ctx, ids)
batchDelete

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **ids** | **[]string**| ids | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BatchUpdateUsingPOST1**
> []EntityTags BatchUpdateUsingPOST1(ctx, tags)
batchUpdate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **tags** | [**[]EntityTags**](EntityTags.md)| tags | 

### Return type

[**[]EntityTags**](EntityTags.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUsingPOST1**
> EntityTags CreateUsingPOST1(ctx, tag)
create

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **tag** | [**EntityTags**](EntityTags.md)| tag | 

### Return type

[**EntityTags**](EntityTags.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUsingDELETE1**
> DeleteUsingDELETE1(ctx, )
delete

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TagUsingGET**
> EntityTags TagUsingGET(ctx, )
tag

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**EntityTags**](EntityTags.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TagsUsingGET**
> []EntityTags TagsUsingGET(ctx, optional)
tags

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **prefix** | **string**| prefix | 
 **ids** | [**[]string**](string.md)| ids | 
 **refresh** | **bool**| refresh | 

### Return type

[**[]EntityTags**](EntityTags.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

