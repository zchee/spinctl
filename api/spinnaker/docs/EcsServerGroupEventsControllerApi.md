# \EcsServerGroupEventsControllerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListEvents**](EcsServerGroupEventsControllerApi.md#ListEvents) | **Get** /applications/{application}/serverGroups/{account}/{serverGroupName}/events | Retrieves a list of events for a server group



## ListEvents

> []map[string]interface{} ListEvents(ctx, account, application, serverGroupName).Provider(provider).Region(region).Execute()

Retrieves a list of events for a server group

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | account | 
**application** | **string** | application | 
**serverGroupName** | **string** | serverGroupName | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **provider** | **string** | provider | 
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

