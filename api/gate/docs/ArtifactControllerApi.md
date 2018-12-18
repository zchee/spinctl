# \ArtifactControllerApi

All URIs are relative to *https://localhost:8084*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllUsingGET**](ArtifactControllerApi.md#AllUsingGET) | **Get** /artifacts/credentials | Retrieve the list of artifact accounts configured in Clouddriver.


# **AllUsingGET**
> []interface{} AllUsingGET(ctx, optional)
Retrieve the list of artifact accounts configured in Clouddriver.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AllUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AllUsingGETOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRateLimitApp** | **optional.String**| X-RateLimit-App | 

### Return type

[**[]interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

