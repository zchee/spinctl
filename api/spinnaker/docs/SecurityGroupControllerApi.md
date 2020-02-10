# \SecurityGroupControllerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSecurityGroup**](SecurityGroupControllerApi.md#GetSecurityGroup) | **Get** /securityGroups/{account}/{region}/{name} | Retrieve a security group&#39;s details
[**ListSecurityGroups**](SecurityGroupControllerApi.md#ListSecurityGroups) | **Get** /securityGroups | Retrieve a list of security groups, grouped by account, cloud provider, and region
[**ListSecurityGroupsByAccount**](SecurityGroupControllerApi.md#ListSecurityGroupsByAccount) | **Get** /securityGroups/{account} | Retrieve a list of security groups for a given account, grouped by region



## GetSecurityGroup

> map[string]interface{} GetSecurityGroup(ctx, account, name, region).XRateLimitApp(xRateLimitApp).Provider(provider).VpcId(vpcId).Execute()

Retrieve a security group's details

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | account | 
**name** | **string** | name | 
**region** | **string** | region | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityGroupRequest struct via the builder pattern


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


## ListSecurityGroups

> map[string]interface{} ListSecurityGroups(ctx).XRateLimitApp(xRateLimitApp).Id(id).Execute()

Retrieve a list of security groups, grouped by account, cloud provider, and region

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSecurityGroupsRequest struct via the builder pattern


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


## ListSecurityGroupsByAccount

> map[string]interface{} ListSecurityGroupsByAccount(ctx, account).XRateLimitApp(xRateLimitApp).Provider(provider).Execute()

Retrieve a list of security groups for a given account, grouped by region

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | account | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSecurityGroupsByAccountRequest struct via the builder pattern


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

