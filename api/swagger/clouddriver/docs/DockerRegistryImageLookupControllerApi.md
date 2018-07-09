# \DockerRegistryImageLookupControllerApi

All URIs are relative to *https://spinnaker-clouddriver.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindUsingGET**](DockerRegistryImageLookupControllerApi.md#FindUsingGET) | **Get** /dockerRegistry/images/find | find
[**GetTagsUsingGET**](DockerRegistryImageLookupControllerApi.md#GetTagsUsingGET) | **Get** /dockerRegistry/images/tags | getTags


# **FindUsingGET**
> []ModelMap FindUsingGET(ctx, optional)
find

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string**|  | 
 **account** | **string**|  | 
 **region** | **string**|  | 
 **count** | **int32**|  | 

### Return type

[**[]ModelMap**](Map.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTagsUsingGET**
> []string GetTagsUsingGET(ctx, account, repository)
getTags

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **account** | **string**| account | 
  **repository** | **string**| repository | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

