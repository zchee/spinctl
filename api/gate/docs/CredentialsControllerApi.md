# \CredentialsControllerApi

All URIs are relative to *http://localhost:8084*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountUsingGET**](CredentialsControllerApi.md#GetAccountUsingGET) | **Get** /credentials/{account} | Retrieve an account&#39;s details
[**GetAccountsUsingGET**](CredentialsControllerApi.md#GetAccountsUsingGET) | **Get** /credentials | Retrieve a list of accounts



## GetAccountUsingGET

> AccountDetails GetAccountUsingGET(ctx, account, optional)
Retrieve an account's details

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string**| account | 
 **optional** | ***GetAccountUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAccountUsingGETOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **roles** | [**optional.Interface of []string**](string.md)|  | 
 **allowedAccounts** | [**optional.Interface of []string**](string.md)|  | 
 **email** | **optional.String**|  | 
 **username** | **optional.String**|  | 
 **firstName** | **optional.String**|  | 
 **lastName** | **optional.String**|  | 
 **xRateLimitApp** | **optional.String**| X-RateLimit-App | 

### Return type

[**AccountDetails**](AccountDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountsUsingGET

> []Account GetAccountsUsingGET(ctx, optional)
Retrieve a list of accounts

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAccountsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAccountsUsingGETOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roles** | [**optional.Interface of []string**](string.md)|  | 
 **allowedAccounts** | [**optional.Interface of []string**](string.md)|  | 
 **email** | **optional.String**|  | 
 **username** | **optional.String**|  | 
 **firstName** | **optional.String**|  | 
 **lastName** | **optional.String**|  | 
 **expand** | **optional.Bool**| expand | 

### Return type

[**[]Account**](Account.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

