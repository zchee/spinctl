# \ServerGroupControllerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetServerGroupDetailsUsingGET**](ServerGroupControllerApi.md#GetServerGroupDetailsUsingGET) | **Get** /applications/{applicationName}/serverGroups/{account}/{region}/{serverGroupName} | Retrieve a server group&#39;s details
[**GetServerGroupsForApplicationUsingGET**](ServerGroupControllerApi.md#GetServerGroupsForApplicationUsingGET) | **Get** /applications/{applicationName}/serverGroups | Retrieve a list of server groups for a given application



## GetServerGroupDetailsUsingGET

> map[string]interface{} GetServerGroupDetailsUsingGET(ctx, account, applicationName, region, serverGroupName).XRateLimitApp(xRateLimitApp).IncludeDetails(includeDetails).Execute()

Retrieve a server group's details

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | account | 
**applicationName** | **string** | applicationName | 
**region** | **string** | region | 
**serverGroupName** | **string** | serverGroupName | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerGroupDetailsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **xRateLimitApp** | **string** | X-RateLimit-App | 
 **includeDetails** | **string** | includeDetails | [default to true]

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


## GetServerGroupsForApplicationUsingGET

> []map[string]interface{} GetServerGroupsForApplicationUsingGET(ctx, applicationName).XRateLimitApp(xRateLimitApp).CloudProvider(cloudProvider).Clusters(clusters).Expand(expand).Execute()

Retrieve a list of server groups for a given application

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationName** | **string** | applicationName | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerGroupsForApplicationUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRateLimitApp** | **string** | X-RateLimit-App | 
 **cloudProvider** | **string** | cloudProvider | 
 **clusters** | **string** | clusters | 
 **expand** | **string** | expand | [default to false]

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

