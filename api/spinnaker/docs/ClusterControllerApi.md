# \ClusterControllerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetClusterLoadBalancers**](ClusterControllerApi.md#GetClusterLoadBalancers) | **Get** /applications/{application}/clusters/{account}/{clusterName}/{type}/loadBalancers | Retrieve a cluster&#39;s loadbalancers
[**GetClustersDetails**](ClusterControllerApi.md#GetClustersDetails) | **Get** /applications/{application}/clusters/{account}/{clusterName} | Retrieve a cluster&#39;s details
[**GetServerGroups**](ClusterControllerApi.md#GetServerGroups) | **Get** /applications/{application}/clusters/{account}/{clusterName}/serverGroups/{serverGroupName} | Retrieve a server group&#39;s details
[**GetTargetServerGroup**](ClusterControllerApi.md#GetTargetServerGroup) | **Get** /applications/{application}/clusters/{account}/{clusterName}/{cloudProvider}/{scope}/serverGroups/target/{target} | Retrieve a server group that matches a target coordinate (e.g., newest, ancestor) relative to a cluster
[**ListClusters**](ClusterControllerApi.md#ListClusters) | **Get** /applications/{application}/clusters | Retrieve a list of cluster names for an application, grouped by account
[**ListClustersByAccount**](ClusterControllerApi.md#ListClustersByAccount) | **Get** /applications/{application}/clusters/{account} | Retrieve a list of clusters for an account
[**ListScalingActivities**](ClusterControllerApi.md#ListScalingActivities) | **Get** /applications/{application}/clusters/{account}/{clusterName}/serverGroups/{serverGroupName}/scalingActivities | Retrieve a list of scaling activities for a server group
[**ListServerGroups**](ClusterControllerApi.md#ListServerGroups) | **Get** /applications/{application}/clusters/{account}/{clusterName}/serverGroups | Retrieve a list of server groups for a cluster



## GetClusterLoadBalancers

> []map[string]interface{} GetClusterLoadBalancers(ctx, account, application, clusterName, type_).XRateLimitApp(xRateLimitApp).Execute()

Retrieve a cluster's loadbalancers

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | account | 
**application** | **string** | application | 
**clusterName** | **string** | clusterName | 
**type_** | **string** | type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterLoadBalancersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **xRateLimitApp** | **string** | X-RateLimit-App | 

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


## GetClustersDetails

> map[string]map[string]interface{} GetClustersDetails(ctx, account, application, clusterName).XRateLimitApp(xRateLimitApp).Execute()

Retrieve a cluster's details

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | account | 
**application** | **string** | application | 
**clusterName** | **string** | clusterName | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xRateLimitApp** | **string** | X-RateLimit-App | 

### Return type

[**map[string]map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerGroups

> []map[string]interface{} GetServerGroups(ctx, account, application, clusterName, serverGroupName).XRateLimitApp(xRateLimitApp).Execute()

Retrieve a server group's details

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | account | 
**application** | **string** | application | 
**clusterName** | **string** | clusterName | 
**serverGroupName** | **string** | serverGroupName | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **xRateLimitApp** | **string** | X-RateLimit-App | 

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


## GetTargetServerGroup

> map[string]map[string]interface{} GetTargetServerGroup(ctx, account, application, cloudProvider, clusterName, scope, target).XRateLimitApp(xRateLimitApp).OnlyEnabled(onlyEnabled).ValidateOldest(validateOldest).Execute()

Retrieve a server group that matches a target coordinate (e.g., newest, ancestor) relative to a cluster



### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | account | 
**application** | **string** | application | 
**cloudProvider** | **string** | cloudProvider | 
**clusterName** | **string** | clusterName | 
**scope** | **string** | scope | 
**target** | **string** | target | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTargetServerGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **xRateLimitApp** | **string** | X-RateLimit-App | 
 **onlyEnabled** | **bool** | onlyEnabled | 
 **validateOldest** | **bool** | validateOldest | 

### Return type

[**map[string]map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClusters

> map[string]map[string]interface{} ListClusters(ctx, application).XRateLimitApp(xRateLimitApp).Execute()

Retrieve a list of cluster names for an application, grouped by account

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string** | application | 

### Other Parameters

Other parameters are passed through a pointer to a apiListClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRateLimitApp** | **string** | X-RateLimit-App | 

### Return type

[**map[string]map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListClustersByAccount

> []map[string]interface{} ListClustersByAccount(ctx, account, application).XRateLimitApp(xRateLimitApp).Execute()

Retrieve a list of clusters for an account

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | account | 
**application** | **string** | application | 

### Other Parameters

Other parameters are passed through a pointer to a apiListClustersByAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRateLimitApp** | **string** | X-RateLimit-App | 

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


## ListScalingActivities

> []map[string]interface{} ListScalingActivities(ctx, account, application, clusterName, serverGroupName).XRateLimitApp(xRateLimitApp).Provider(provider).Region(region).Execute()

Retrieve a list of scaling activities for a server group

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | account | 
**application** | **string** | application | 
**clusterName** | **string** | clusterName | 
**serverGroupName** | **string** | serverGroupName | 

### Other Parameters

Other parameters are passed through a pointer to a apiListScalingActivitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **xRateLimitApp** | **string** | X-RateLimit-App | 
 **provider** | **string** | provider | [default to aws]
 **region** | **string** | region | 

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


## ListServerGroups

> []map[string]interface{} ListServerGroups(ctx, account, application, clusterName).XRateLimitApp(xRateLimitApp).Execute()

Retrieve a list of server groups for a cluster

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | account | 
**application** | **string** | application | 
**clusterName** | **string** | clusterName | 

### Other Parameters

Other parameters are passed through a pointer to a apiListServerGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xRateLimitApp** | **string** | X-RateLimit-App | 

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

