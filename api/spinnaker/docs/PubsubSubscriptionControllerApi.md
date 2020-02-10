# \PubsubSubscriptionControllerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListPubSubSubscriptions**](PubsubSubscriptionControllerApi.md#ListPubSubSubscriptions) | **Get** /pubsub/subscriptions | Retrieve the list of pub/sub subscriptions configured in Echo.



## ListPubSubSubscriptions

> []map[string]string ListPubSubSubscriptions(ctx).Execute()

Retrieve the list of pub/sub subscriptions configured in Echo.

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPubSubSubscriptionsRequest struct via the builder pattern


### Return type

[**[]map[string]string**](map.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

