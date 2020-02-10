# \CredentialsControllerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccount**](CredentialsControllerApi.md#GetAccount) | **Get** /credentials/{account} | Retrieve an account&#39;s details
[**ListAccounts**](CredentialsControllerApi.md#ListAccounts) | **Get** /credentials | Retrieve a list of accounts



## GetAccount

> AccountDetails GetAccount(ctx, account).XRateLimitApp(xRateLimitApp).AccountNonExpired(accountNonExpired).AccountNonLocked(accountNonLocked).AllowedAccounts(allowedAccounts).Authorities0Authority(authorities0Authority).CredentialsNonExpired(credentialsNonExpired).Email(email).Enabled(enabled).FirstName(firstName).LastName(lastName).Password(password).Roles(roles).Username(username).Execute()

Retrieve an account's details

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | account | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRateLimitApp** | **string** | X-RateLimit-App | 
 **accountNonExpired** | **bool** |  | 
 **accountNonLocked** | **bool** |  | 
 **allowedAccounts** | [**[]string**](string.md) |  | 
 **authorities0Authority** | **string** |  | 
 **credentialsNonExpired** | **bool** |  | 
 **email** | **string** |  | 
 **enabled** | **bool** |  | 
 **firstName** | **string** |  | 
 **lastName** | **string** |  | 
 **password** | **string** |  | 
 **roles** | [**[]string**](string.md) |  | 
 **username** | **string** |  | 

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


## ListAccounts

> []Account ListAccounts(ctx).AccountNonExpired(accountNonExpired).AccountNonLocked(accountNonLocked).AllowedAccounts(allowedAccounts).Authorities0Authority(authorities0Authority).CredentialsNonExpired(credentialsNonExpired).Email(email).Enabled(enabled).Expand(expand).FirstName(firstName).LastName(lastName).Password(password).Roles(roles).Username(username).Execute()

Retrieve a list of accounts

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountNonExpired** | **bool** |  | 
 **accountNonLocked** | **bool** |  | 
 **allowedAccounts** | [**[]string**](string.md) |  | 
 **authorities0Authority** | **string** |  | 
 **credentialsNonExpired** | **bool** |  | 
 **email** | **string** |  | 
 **enabled** | **bool** |  | 
 **expand** | **bool** | expand | 
 **firstName** | **string** |  | 
 **lastName** | **string** |  | 
 **password** | **string** |  | 
 **roles** | [**[]string**](string.md) |  | 
 **username** | **string** |  | 

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

