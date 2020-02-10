# \WebhookControllerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListPreconfiguredWebhooks**](WebhookControllerApi.md#ListPreconfiguredWebhooks) | **Get** /webhooks/preconfigured | Retrieve a list of preconfigured webhooks in Orca
[**Webhooks**](WebhookControllerApi.md#Webhooks) | **Post** /webhooks/{type}/{source} | Endpoint for posting webhooks to Spinnaker&#39;s webhook service



## ListPreconfiguredWebhooks

> []map[string]interface{} ListPreconfiguredWebhooks(ctx).Execute()

Retrieve a list of preconfigured webhooks in Orca

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPreconfiguredWebhooksRequest struct via the builder pattern


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


## Webhooks

> map[string]interface{} Webhooks(ctx, source, type_).Event(event).XEventKey(xEventKey).XHubSignature(xHubSignature).Execute()

Endpoint for posting webhooks to Spinnaker's webhook service

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**source** | **string** | source | 
**type_** | **string** | type | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **event** | **map[string]interface{}** | event | 
 **xEventKey** | **string** | X-Event-Key | 
 **xHubSignature** | **string** | X-Hub-Signature | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

