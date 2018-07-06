# \ServerGroupControllerApi

All URIs are relative to *https://spinnaker-clouddriver.meteoros.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetServerGroupByApplicationUsingGET**](ServerGroupControllerApi.md#GetServerGroupByApplicationUsingGET) | **Get** /applications/{application}/serverGroups/{account}/{region}/{name} | getServerGroupByApplication
[**ListUsingGET8**](ServerGroupControllerApi.md#ListUsingGET8) | **Get** /applications/{application}/serverGroups | list


# **GetServerGroupByApplicationUsingGET**
> ServerGroup GetServerGroupByApplicationUsingGET(ctx, application, account, region, name, optional)
getServerGroupByApplication

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **application** | **string**| application | 
  **account** | **string**| account | 
  **region** | **string**| region | 
  **name** | **string**| name | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **application** | **string**| application | 
 **account** | **string**| account | 
 **region** | **string**| region | 
 **name** | **string**| name | 
 **includeDetails** | **string**| includeDetails | [default to true]

### Return type

[**ServerGroup**](ServerGroup.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUsingGET8**
> []interface{} ListUsingGET8(ctx, application, optional)
list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **application** | **string**| application | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **application** | **string**| application | 
 **expand** | **string**| expand | [default to false]
 **cloudProvider** | **string**| cloudProvider | 
 **clusters** | [**[]string**](string.md)| clusters | 

### Return type

[**[]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

