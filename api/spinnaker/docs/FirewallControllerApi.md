# \FirewallControllerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFirewallsSecurityGroup**](FirewallControllerApi.md#GetFirewallsSecurityGroup) | **Get** /firewalls/{account}/{region}/{name} | Retrieve a firewall&#39;s details
[**ListFirewalls**](FirewallControllerApi.md#ListFirewalls) | **Get** /firewalls | Retrieve a list of firewalls, grouped by account, cloud provider, and region
[**ListFirewallsByAccount**](FirewallControllerApi.md#ListFirewallsByAccount) | **Get** /firewalls/{account} | Retrieve a list of firewalls for a given account, grouped by region
[**ListFirewallsByAccountAndRegion**](FirewallControllerApi.md#ListFirewallsByAccountAndRegion) | **Get** /firewalls/{account}/{region} | Retrieve a list of firewalls for a given account and region



## GetFirewallsSecurityGroup

> map[string]interface{} GetFirewallsSecurityGroup(ctx, account, name, region).XRateLimitApp(xRateLimitApp).Provider(provider).VpcId(vpcId).Execute()

Retrieve a firewall's details

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | account | 
**name** | **string** | name | 
**region** | **string** | region | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFirewallsSecurityGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xRateLimitApp** | **string** | X-RateLimit-App | 
 **provider** | **string** | provider | [default to aws]
 **vpcId** | **string** | vpcId | 

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


## ListFirewalls

> map[string]interface{} ListFirewalls(ctx).XRateLimitApp(xRateLimitApp).Id(id).Execute()

Retrieve a list of firewalls, grouped by account, cloud provider, and region

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFirewallsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRateLimitApp** | **string** | X-RateLimit-App | 
 **id** | **string** | id | 

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


## ListFirewallsByAccount

> map[string]interface{} ListFirewallsByAccount(ctx, account).XRateLimitApp(xRateLimitApp).Provider(provider).Execute()

Retrieve a list of firewalls for a given account, grouped by region

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | account | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFirewallsByAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRateLimitApp** | **string** | X-RateLimit-App | 
 **provider** | **string** | provider | [default to aws]

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


## ListFirewallsByAccountAndRegion

> []map[string]interface{} ListFirewallsByAccountAndRegion(ctx, account, region).XRateLimitApp(xRateLimitApp).Provider(provider).Execute()

Retrieve a list of firewalls for a given account and region

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | account | 
**region** | **string** | region | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFirewallsByAccountAndRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRateLimitApp** | **string** | X-RateLimit-App | 
 **provider** | **string** | provider | [default to aws]

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

