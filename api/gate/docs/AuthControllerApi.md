# \AuthControllerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetServiceAccountsUsingGET**](AuthControllerApi.md#GetServiceAccountsUsingGET) | **Get** /auth/user/serviceAccounts | Get service accounts
[**LoggedOutUsingGET**](AuthControllerApi.md#LoggedOutUsingGET) | **Get** /auth/loggedOut | Get logged out message
[**RedirectUsingGET**](AuthControllerApi.md#RedirectUsingGET) | **Get** /auth/redirect | Redirect to Deck
[**SyncUsingPOST**](AuthControllerApi.md#SyncUsingPOST) | **Post** /auth/roles/sync | Sync user roles
[**UserUsingGET**](AuthControllerApi.md#UserUsingGET) | **Get** /auth/user | Get user



## GetServiceAccountsUsingGET

> []map[string]interface{} GetServiceAccountsUsingGET(ctx).Execute()

Get service accounts

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceAccountsUsingGETRequest struct via the builder pattern


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


## LoggedOutUsingGET

> string LoggedOutUsingGET(ctx).Execute()

Get logged out message

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLoggedOutUsingGETRequest struct via the builder pattern


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RedirectUsingGET

> RedirectUsingGET(ctx).To(to).Execute()

Redirect to Deck

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRedirectUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **to** | **string** | to | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncUsingPOST

> SyncUsingPOST(ctx).Execute()

Sync user roles

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSyncUsingPOSTRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserUsingGET

> User UserUsingGET(ctx).Execute()

Get user

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUserUsingGETRequest struct via the builder pattern


### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

