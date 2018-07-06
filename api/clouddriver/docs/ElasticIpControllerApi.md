# \ElasticIpControllerApi

All URIs are relative to *https://spinnaker-clouddriver.meteoros.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListByAccountAndRegionUsingGET**](ElasticIpControllerApi.md#ListByAccountAndRegionUsingGET) | **Get** /elasticIps/{account} | listByAccountAndRegion


# **ListByAccountAndRegionUsingGET**
> []ElasticIp ListByAccountAndRegionUsingGET(ctx, account, region)
listByAccountAndRegion

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **account** | **string**| account | 
  **region** | **string**| region | 

### Return type

[**[]ElasticIp**](ElasticIp.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

