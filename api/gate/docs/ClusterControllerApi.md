# \ClusterControllerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetClusterLoadBalancersUsingGET**](ClusterControllerApi.md#GetClusterLoadBalancersUsingGET) | **Get** /applications/{application}/clusters/{account}/{clusterName}/{type}/loadBalancers | Retrieve a cluster&#39;s loadbalancers
[**GetClustersUsingGET**](ClusterControllerApi.md#GetClustersUsingGET) | **Get** /applications/{application}/clusters/{account}/{clusterName} | Retrieve a cluster&#39;s details
[**GetClustersUsingGET1**](ClusterControllerApi.md#GetClustersUsingGET1) | **Get** /applications/{application}/clusters/{account} | Retrieve a list of clusters for an account
[**GetClustersUsingGET2**](ClusterControllerApi.md#GetClustersUsingGET2) | **Get** /applications/{application}/clusters | Retrieve a list of cluster names for an application, grouped by account
[**GetScalingActivitiesUsingGET**](ClusterControllerApi.md#GetScalingActivitiesUsingGET) | **Get** /applications/{application}/clusters/{account}/{clusterName}/serverGroups/{serverGroupName}/scalingActivities | Retrieve a list of scaling activities for a server group
[**GetServerGroupsUsingGET**](ClusterControllerApi.md#GetServerGroupsUsingGET) | **Get** /applications/{application}/clusters/{account}/{clusterName}/serverGroups/{serverGroupName} | Retrieve a server group&#39;s details
[**GetServerGroupsUsingGET1**](ClusterControllerApi.md#GetServerGroupsUsingGET1) | **Get** /applications/{application}/clusters/{account}/{clusterName}/serverGroups | Retrieve a list of server groups for a cluster
[**GetTargetServerGroupUsingGET**](ClusterControllerApi.md#GetTargetServerGroupUsingGET) | **Get** /applications/{application}/clusters/{account}/{clusterName}/{cloudProvider}/{scope}/serverGroups/target/{target} | Retrieve a server group that matches a target coordinate (e.g., newest, ancestor) relative to a cluster



## GetClusterLoadBalancersUsingGET

> []map[string]interface{} GetClusterLoadBalancersUsingGET(ctx, account, applicationName, clusterName, type_).XRateLimitApp(xRateLimitApp).Execute()

Retrieve a cluster's loadbalancers

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | account | 
**applicationName** | **string** | applicationName | 
**clusterName** | **string** | clusterName | 
**type_** | **string** | type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClusterLoadBalancersUsingGETRequest struct via the builder pattern


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


## GetClustersUsingGET

> map[string]map[string]interface{} GetClustersUsingGET(ctx, account, application, clusterName).XRateLimitApp(xRateLimitApp).Execute()

Retrieve a cluster's details

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | account | 
**application** | **string** | application | 
**clusterName** | **string** | clusterName | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersUsingGETRequest struct via the builder pattern


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


## GetClustersUsingGET1

> []map[string]interface{} GetClustersUsingGET1(ctx, account, application).XRateLimitApp(xRateLimitApp).Execute()

Retrieve a list of clusters for an account

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | account | 
**application** | **string** | application | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersUsingGET1Request struct via the builder pattern


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


## GetClustersUsingGET2

> map[string]map[string]interface{} GetClustersUsingGET2(ctx, application).XRateLimitApp(xRateLimitApp).Execute()

Retrieve a list of cluster names for an application, grouped by account

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string** | application | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersUsingGET2Request struct via the builder pattern


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


## GetScalingActivitiesUsingGET

> []map[string]interface{} GetScalingActivitiesUsingGET(ctx, account, application, clusterName, serverGroupName).XRateLimitApp(xRateLimitApp).Provider(provider).Region(region).Execute()

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

Other parameters are passed through a pointer to a apiGetScalingActivitiesUsingGETRequest struct via the builder pattern


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


## GetServerGroupsUsingGET

> []map[string]interface{} GetServerGroupsUsingGET(ctx, account, application, clusterName, serverGroupName).XRateLimitApp(xRateLimitApp).Execute()

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

Other parameters are passed through a pointer to a apiGetServerGroupsUsingGETRequest struct via the builder pattern


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


## GetServerGroupsUsingGET1

> []map[string]interface{} GetServerGroupsUsingGET1(ctx, account, application, clusterName).XRateLimitApp(xRateLimitApp).Execute()

Retrieve a list of server groups for a cluster

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **string** | account | 
**application** | **string** | application | 
**clusterName** | **string** | clusterName | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerGroupsUsingGET1Request struct via the builder pattern


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


## GetTargetServerGroupUsingGET

> map[string]map[string]interface{} GetTargetServerGroupUsingGET(ctx, account, application, cloudProvider, clusterName, scope, target).XRateLimitApp(xRateLimitApp).OnlyEnabled(onlyEnabled).ValidateOldest(validateOldest).Execute()

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

Other parameters are passed through a pointer to a apiGetTargetServerGroupUsingGETRequest struct via the builder pattern


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

