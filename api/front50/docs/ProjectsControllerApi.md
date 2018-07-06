# \ProjectsControllerApi

All URIs are relative to *https://spinnaker-front50.meteoros.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchUpdateUsingPOST4**](ProjectsControllerApi.md#BatchUpdateUsingPOST4) | **Post** /v2/projects/batchUpdate | batchUpdate
[**CreateUsingPOST2**](ProjectsControllerApi.md#CreateUsingPOST2) | **Post** /v2/projects | create
[**DeleteUsingDELETE4**](ProjectsControllerApi.md#DeleteUsingDELETE4) | **Delete** /v2/projects/{projectId} | delete
[**ProjectUsingGET**](ProjectsControllerApi.md#ProjectUsingGET) | **Get** /v2/projects/{projectId} | project
[**ProjectsUsingGET**](ProjectsControllerApi.md#ProjectsUsingGET) | **Get** /v2/projects | projects
[**PutUsingPUT**](ProjectsControllerApi.md#PutUsingPUT) | **Put** /v2/projects/{projectId} | put
[**SearchUsingGET**](ProjectsControllerApi.md#SearchUsingGET) | **Get** /v2/projects/search | search


# **BatchUpdateUsingPOST4**
> BatchUpdateUsingPOST4(ctx, projects)
batchUpdate

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projects** | [**[]Project**](Project.md)| projects | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateUsingPOST2**
> Project CreateUsingPOST2(ctx, project)
create

Create a project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **project** | [**Project**](Project.md)| project | 

### Return type

[**Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUsingDELETE4**
> DeleteUsingDELETE4(ctx, projectId)
delete

Delete a project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectId** | **string**| projectId | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectUsingGET**
> Project ProjectUsingGET(ctx, projectId)
project

Fetch a single project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectId** | **string**| projectId | 

### Return type

[**Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsUsingGET**
> []Project ProjectsUsingGET(ctx, params, optional)
projects

Fetch all projects.      Support filtering by one or more attributes:     - ?name=projectName     - ?email=my@email.com

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **params** | **string**| params | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **params** | **string**| params | 
 **pageSize** | **int32**| pageSize | 

### Return type

[**[]Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutUsingPUT**
> Project PutUsingPUT(ctx, projectId, project)
put

Update an existing project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **projectId** | **string**| projectId | 
  **project** | [**Project**](Project.md)| project | 

### Return type

[**Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchUsingGET**
> []Project SearchUsingGET(ctx, q)
search

Search for projects given one or more attributes.  - /search?q=ProjectName - /search?q=ApplicationName 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **q** | **string**| q | 

### Return type

[**[]Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

