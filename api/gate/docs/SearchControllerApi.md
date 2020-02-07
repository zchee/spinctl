# \SearchControllerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchUsingGET**](SearchControllerApi.md#SearchUsingGET) | **Get** /search | Search infrastructure



## SearchUsingGET

> []map[string]interface{} SearchUsingGET(ctx).Type_(type_).XRateLimitApp(xRateLimitApp).AllowShortQuery(allowShortQuery).Page(page).PageSize(pageSize).Platform(platform).Q(q).Execute()

Search infrastructure

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | type | 
 **xRateLimitApp** | **string** | X-RateLimit-App | 
 **allowShortQuery** | **bool** | allowShortQuery | [default to false]
 **page** | **int32** | page | [default to 1]
 **pageSize** | **int32** | pageSize | [default to 10000]
 **platform** | **string** | platform | 
 **q** | **string** | q | 

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

