# \ClusterControllerApi

All URIs are relative to *https://spinnaker-clouddriver.example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetForAccountAndNameAndTypeUsingGET**](ClusterControllerApi.md#GetForAccountAndNameAndTypeUsingGET) | **Get** /applications/{application}/clusters/{account}/{name}/{type} | getForAccountAndNameAndType
[**GetForAccountAndNameUsingGET**](ClusterControllerApi.md#GetForAccountAndNameUsingGET) | **Get** /applications/{application}/clusters/{account}/{name} | getForAccountAndName
[**GetForAccountUsingGET**](ClusterControllerApi.md#GetForAccountUsingGET) | **Get** /applications/{application}/clusters/{account} | getForAccount
[**GetServerGroupSummaryUsingGET**](ClusterControllerApi.md#GetServerGroupSummaryUsingGET) | **Get** /applications/{application}/clusters/{account}/{clusterName}/{cloudProvider}/{scope}/serverGroups/target/{target}/{summaryType} | getServerGroupSummary
[**GetServerGroupUsingGET**](ClusterControllerApi.md#GetServerGroupUsingGET) | **Get** /applications/{application}/clusters/{account}/{clusterName}/{type}/serverGroups/{serverGroupName} | getServerGroup
[**GetServerGroupsUsingGET**](ClusterControllerApi.md#GetServerGroupsUsingGET) | **Get** /applications/{application}/clusters/{account}/{clusterName}/{type}/serverGroups | getServerGroups
[**GetTargetServerGroupUsingGET**](ClusterControllerApi.md#GetTargetServerGroupUsingGET) | **Get** /applications/{application}/clusters/{account}/{clusterName}/{cloudProvider}/{scope}/serverGroups/target/{target} | getTargetServerGroup
[**ListByAccountUsingGET**](ClusterControllerApi.md#ListByAccountUsingGET) | **Get** /applications/{application}/clusters | listByAccount


# **GetForAccountAndNameAndTypeUsingGET**
> Cluster GetForAccountAndNameAndTypeUsingGET(ctx, application, account, name, type_, optional)
getForAccountAndNameAndType

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **application** | **string**| application | 
  **account** | **string**| account | 
  **name** | **string**| name | 
  **type_** | **string**| type | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **application** | **string**| application | 
 **account** | **string**| account | 
 **name** | **string**| name | 
 **type_** | **string**| type | 
 **expand** | **bool**| expand | [default to true]

### Return type

[**Cluster**](Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetForAccountAndNameUsingGET**
> []Cluster GetForAccountAndNameUsingGET(ctx, application, account, name, optional)
getForAccountAndName

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **application** | **string**| application | 
  **account** | **string**| account | 
  **name** | **string**| name | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **application** | **string**| application | 
 **account** | **string**| account | 
 **name** | **string**| name | 
 **expand** | **bool**| expand | [default to true]

### Return type

[**[]Cluster**](Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetForAccountUsingGET**
> []ClusterViewModel GetForAccountUsingGET(ctx, application, account)
getForAccount

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **application** | **string**| application | 
  **account** | **string**| account | 

### Return type

[**[]ClusterViewModel**](ClusterViewModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServerGroupSummaryUsingGET**
> Summary GetServerGroupSummaryUsingGET(ctx, application, account, clusterName, cloudProvider, scope, target, summaryType, optional)
getServerGroupSummary

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **application** | **string**| application | 
  **account** | **string**| account | 
  **clusterName** | **string**| clusterName | 
  **cloudProvider** | **string**| cloudProvider | 
  **scope** | **string**| scope | 
  **target** | **string**| target | 
  **summaryType** | **string**| summaryType | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **application** | **string**| application | 
 **account** | **string**| account | 
 **clusterName** | **string**| clusterName | 
 **cloudProvider** | **string**| cloudProvider | 
 **scope** | **string**| scope | 
 **target** | **string**| target | 
 **summaryType** | **string**| summaryType | 
 **onlyEnabled** | **string**| onlyEnabled | [default to false]

### Return type

[**Summary**](Summary.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServerGroupUsingGET**
> interface{} GetServerGroupUsingGET(ctx, application, account, clusterName, type_, serverGroupName, optional)
getServerGroup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **application** | **string**| application | 
  **account** | **string**| account | 
  **clusterName** | **string**| clusterName | 
  **type_** | **string**| type | 
  **serverGroupName** | **string**| serverGroupName | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **application** | **string**| application | 
 **account** | **string**| account | 
 **clusterName** | **string**| clusterName | 
 **type_** | **string**| type | 
 **serverGroupName** | **string**| serverGroupName | 
 **region** | **string**| region | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetServerGroupsUsingGET**
> []ServerGroup GetServerGroupsUsingGET(ctx, application, account, clusterName, type_, optional)
getServerGroups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **application** | **string**| application | 
  **account** | **string**| account | 
  **clusterName** | **string**| clusterName | 
  **type_** | **string**| type | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **application** | **string**| application | 
 **account** | **string**| account | 
 **clusterName** | **string**| clusterName | 
 **type_** | **string**| type | 
 **region** | **string**| region | 
 **expand** | **bool**| expand | [default to true]

### Return type

[**[]ServerGroup**](ServerGroup.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTargetServerGroupUsingGET**
> ServerGroup GetTargetServerGroupUsingGET(ctx, application, account, clusterName, cloudProvider, scope, target, optional)
getTargetServerGroup

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **application** | **string**| application | 
  **account** | **string**| account | 
  **clusterName** | **string**| clusterName | 
  **cloudProvider** | **string**| cloudProvider | 
  **scope** | **string**| scope | 
  **target** | **string**| target | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **application** | **string**| application | 
 **account** | **string**| account | 
 **clusterName** | **string**| clusterName | 
 **cloudProvider** | **string**| cloudProvider | 
 **scope** | **string**| scope | 
 **target** | **string**| target | 
 **onlyEnabled** | **string**| onlyEnabled | [default to false]
 **validateOldest** | **string**| validateOldest | [default to true]

### Return type

[**ServerGroup**](ServerGroup.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListByAccountUsingGET**
> map[string][]string ListByAccountUsingGET(ctx, application)
listByAccount

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **application** | **string**| application | 

### Return type

[**map[string][]string**](array.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

