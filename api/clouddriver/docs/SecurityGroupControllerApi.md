# \SecurityGroupControllerApi

All URIs are relative to *https://spinnaker-clouddriver.meteoros.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUsingGET1**](SecurityGroupControllerApi.md#GetUsingGET1) | **Get** /securityGroups/{account}/{cloudProvider}/{region}/{securityGroupName} | get
[**ListByAccountAndCloudProviderAndNameUsingGET**](SecurityGroupControllerApi.md#ListByAccountAndCloudProviderAndNameUsingGET) | **Get** /securityGroups/{account}/{cloudProvider}/{securityGroupName} | listByAccountAndCloudProviderAndName
[**ListByAccountAndCloudProviderAndRegionUsingGET**](SecurityGroupControllerApi.md#ListByAccountAndCloudProviderAndRegionUsingGET) | **Get** /securityGroups/{account}/{cloudProvider} | listByAccountAndCloudProviderAndRegion
[**ListByAccountUsingGET2**](SecurityGroupControllerApi.md#ListByAccountUsingGET2) | **Get** /securityGroups/{account} | listByAccount
[**ListUsingGET7**](SecurityGroupControllerApi.md#ListUsingGET7) | **Get** /securityGroups | list


# **GetUsingGET1**
> SecurityGroup GetUsingGET1(ctx, account, cloudProvider, region, securityGroupName, optional)
get

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **account** | **string**| account | 
  **cloudProvider** | **string**| cloudProvider | 
  **region** | **string**| region | 
  **securityGroupName** | **string**| securityGroupName | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string**| account | 
 **cloudProvider** | **string**| cloudProvider | 
 **region** | **string**| region | 
 **securityGroupName** | **string**| securityGroupName | 
 **vpcId** | **string**| vpcId | 

### Return type

[**SecurityGroup**](SecurityGroup.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListByAccountAndCloudProviderAndNameUsingGET**
> map[string][]SecurityGroupSummary ListByAccountAndCloudProviderAndNameUsingGET(ctx, account, cloudProvider, securityGroupName)
listByAccountAndCloudProviderAndName

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **account** | **string**| account | 
  **cloudProvider** | **string**| cloudProvider | 
  **securityGroupName** | **string**| securityGroupName | 

### Return type

[**map[string][]SecurityGroupSummary**](array.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListByAccountAndCloudProviderAndRegionUsingGET**
> []SecurityGroupSummary ListByAccountAndCloudProviderAndRegionUsingGET(ctx, account, cloudProvider, region)
listByAccountAndCloudProviderAndRegion

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **account** | **string**| account | 
  **cloudProvider** | **string**| cloudProvider | 
  **region** | **string**| region | 

### Return type

[**[]SecurityGroupSummary**](SecurityGroupSummary.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListByAccountUsingGET2**
> map[string]map[string][]SecurityGroupSummary ListByAccountUsingGET2(ctx, account)
listByAccount

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **account** | **string**| account | 

### Return type

[**map[string]map[string][]SecurityGroupSummary**](map.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListUsingGET7**
> map[string]map[string]map[string][]SecurityGroupSummary ListUsingGET7(ctx, )
list

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**map[string]map[string]map[string][]SecurityGroupSummary**](map.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

