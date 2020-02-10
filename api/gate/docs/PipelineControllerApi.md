# \PipelineControllerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelPipelineUsingPUT1**](PipelineControllerApi.md#CancelPipelineUsingPUT1) | **Put** /pipelines/{id}/cancel | Cancel a pipeline execution
[**DeletePipelineUsingDELETE**](PipelineControllerApi.md#DeletePipelineUsingDELETE) | **Delete** /pipelines/{application}/{pipelineName} | Delete a pipeline definition
[**DeletePipelineUsingDELETE1**](PipelineControllerApi.md#DeletePipelineUsingDELETE1) | **Delete** /pipelines/{id} | Delete a pipeline execution
[**EvaluateExpressionForExecutionAtStageUsingGET**](PipelineControllerApi.md#EvaluateExpressionForExecutionAtStageUsingGET) | **Get** /pipelines/{id}/{stageId}/evaluateExpression | Evaluate a pipeline expression at a specific stage using the provided execution as context
[**EvaluateExpressionForExecutionUsingGET**](PipelineControllerApi.md#EvaluateExpressionForExecutionUsingGET) | **Get** /pipelines/{id}/evaluateExpression | Evaluate a pipeline expression using the provided execution as context
[**EvaluateExpressionForExecutionViaPOSTUsingPOST1**](PipelineControllerApi.md#EvaluateExpressionForExecutionViaPOSTUsingPOST1) | **Post** /pipelines/{id}/evaluateExpression | Evaluate a pipeline expression using the provided execution as context
[**EvaluateVariablesUsingPOST**](PipelineControllerApi.md#EvaluateVariablesUsingPOST) | **Post** /pipelines/{id}/evaluateVariables | Evaluate variables same as Evaluate Variables stage using the provided execution as context
[**GetPipelineUsingGET**](PipelineControllerApi.md#GetPipelineUsingGET) | **Get** /pipelines/{id} | Retrieve a pipeline execution
[**InvokePipelineConfigUsingPOST1**](PipelineControllerApi.md#InvokePipelineConfigUsingPOST1) | **Post** /pipelines/{application}/{pipelineNameOrId} | Trigger a pipeline execution
[**InvokePipelineConfigViaEchoUsingPOST**](PipelineControllerApi.md#InvokePipelineConfigViaEchoUsingPOST) | **Post** /pipelines/v2/{application}/{pipelineNameOrId} | Trigger a pipeline execution
[**PausePipelineUsingPUT**](PipelineControllerApi.md#PausePipelineUsingPUT) | **Put** /pipelines/{id}/pause | Pause a pipeline execution
[**RenamePipelineUsingPOST**](PipelineControllerApi.md#RenamePipelineUsingPOST) | **Post** /pipelines/move | Rename a pipeline definition
[**RestartStageUsingPUT**](PipelineControllerApi.md#RestartStageUsingPUT) | **Put** /pipelines/{id}/stages/{stageId}/restart | Restart a stage execution
[**ResumePipelineUsingPUT**](PipelineControllerApi.md#ResumePipelineUsingPUT) | **Put** /pipelines/{id}/resume | Resume a pipeline execution
[**SavePipelineUsingPOST**](PipelineControllerApi.md#SavePipelineUsingPOST) | **Post** /pipelines | Save a pipeline definition
[**StartUsingPOST**](PipelineControllerApi.md#StartUsingPOST) | **Post** /pipelines/start | Initiate a pipeline execution
[**UpdatePipelineUsingPUT**](PipelineControllerApi.md#UpdatePipelineUsingPUT) | **Put** /pipelines/{id} | Update a pipeline definition
[**UpdateStageUsingPATCH**](PipelineControllerApi.md#UpdateStageUsingPATCH) | **Patch** /pipelines/{id}/stages/{stageId} | Update a stage execution



## CancelPipelineUsingPUT1

> CancelPipelineUsingPUT1(ctx, id).Force(force).Reason(reason).Execute()

Cancel a pipeline execution

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelPipelineUsingPUT1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **bool** | force | [default to false]
 **reason** | **string** | reason | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePipelineUsingDELETE

> DeletePipelineUsingDELETE(ctx, application, pipelineName).Execute()

Delete a pipeline definition

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string** | application | 
**pipelineName** | **string** | pipelineName | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePipelineUsingDELETERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePipelineUsingDELETE1

> map[string]map[string]interface{} DeletePipelineUsingDELETE1(ctx, id).Execute()

Delete a pipeline execution

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePipelineUsingDELETE1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## EvaluateExpressionForExecutionAtStageUsingGET

> map[string]map[string]interface{} EvaluateExpressionForExecutionAtStageUsingGET(ctx, id, stageId).Expression(expression).Execute()

Evaluate a pipeline expression at a specific stage using the provided execution as context

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 
**stageId** | **string** | stageId | 

### Other Parameters

Other parameters are passed through a pointer to a apiEvaluateExpressionForExecutionAtStageUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expression** | **string** | expression | 



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


## EvaluateExpressionForExecutionUsingGET

> map[string]map[string]interface{} EvaluateExpressionForExecutionUsingGET(ctx, id).Expression(expression).Execute()

Evaluate a pipeline expression using the provided execution as context

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEvaluateExpressionForExecutionUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expression** | **string** | expression | 


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


## EvaluateExpressionForExecutionViaPOSTUsingPOST1

> map[string]map[string]interface{} EvaluateExpressionForExecutionViaPOSTUsingPOST1(ctx, id).PipelineExpression(pipelineExpression).Execute()

Evaluate a pipeline expression using the provided execution as context

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEvaluateExpressionForExecutionViaPOSTUsingPOST1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pipelineExpression** | **map[string]interface{}** | pipelineExpression | 


### Return type

[**map[string]map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EvaluateVariablesUsingPOST

> map[string]map[string]interface{} EvaluateVariablesUsingPOST(ctx).Expressions(expressions).ExecutionId(executionId).RequisiteStageRefIds(requisiteStageRefIds).SpelVersion(spelVersion).Execute()

Evaluate variables same as Evaluate Variables stage using the provided execution as context

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEvaluateVariablesUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expressions** | [**[]map[string]string**](map.md) | List of variables/expressions to evaluate | 
 **executionId** | **string** | Execution id to run against | 
 **requisiteStageRefIds** | **string** | Comma separated list of requisite stage IDs for the evaluation stage | 
 **spelVersion** | **string** | Version of SpEL evaluation logic to use (v3 or v4) | 

### Return type

[**map[string]map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPipelineUsingGET

> map[string]interface{} GetPipelineUsingGET(ctx, id).Execute()

Retrieve a pipeline execution

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelineUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## InvokePipelineConfigUsingPOST1

> InvokePipelineConfigUsingPOST1(ctx, application, pipelineNameOrId).Trigger(trigger).Execute()

Trigger a pipeline execution

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string** | application | 
**pipelineNameOrId** | **string** | pipelineNameOrId | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvokePipelineConfigUsingPOST1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **trigger** | **map[string]interface{}** | trigger | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvokePipelineConfigViaEchoUsingPOST

> HttpEntity InvokePipelineConfigViaEchoUsingPOST(ctx, application, pipelineNameOrId).Trigger(trigger).Execute()

Trigger a pipeline execution

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string** | application | 
**pipelineNameOrId** | **string** | pipelineNameOrId | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvokePipelineConfigViaEchoUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **trigger** | **map[string]interface{}** | trigger | 

### Return type

[**HttpEntity**](HttpEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PausePipelineUsingPUT

> PausePipelineUsingPUT(ctx, id).Execute()

Pause a pipeline execution

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPausePipelineUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenamePipelineUsingPOST

> RenamePipelineUsingPOST(ctx).RenameCommand(renameCommand).Execute()

Rename a pipeline definition

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRenamePipelineUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **renameCommand** | **map[string]interface{}** | renameCommand | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestartStageUsingPUT

> map[string]map[string]interface{} RestartStageUsingPUT(ctx, id, stageId).Context(context).Execute()

Restart a stage execution

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 
**stageId** | **string** | stageId | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestartStageUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **context** | **map[string]interface{}** | context | 



### Return type

[**map[string]map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumePipelineUsingPUT

> map[string]map[string]interface{} ResumePipelineUsingPUT(ctx, id).Execute()

Resume a pipeline execution

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumePipelineUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## SavePipelineUsingPOST

> SavePipelineUsingPOST(ctx).Pipeline(pipeline).Execute()

Save a pipeline definition

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSavePipelineUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pipeline** | **map[string]interface{}** | pipeline | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartUsingPOST

> ResponseEntity StartUsingPOST(ctx).Map_(map_).Execute()

Initiate a pipeline execution

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **map_** | **map[string]interface{}** | map | 

### Return type

[**ResponseEntity**](ResponseEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePipelineUsingPUT

> map[string]map[string]interface{} UpdatePipelineUsingPUT(ctx, id).Pipeline(pipeline).Execute()

Update a pipeline definition

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePipelineUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pipeline** | **map[string]interface{}** | pipeline | 


### Return type

[**map[string]map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStageUsingPATCH

> map[string]map[string]interface{} UpdateStageUsingPATCH(ctx, id, stageId).Context(context).Execute()

Update a stage execution

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 
**stageId** | **string** | stageId | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStageUsingPATCHRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **context** | **map[string]interface{}** | context | 



### Return type

[**map[string]map[string]interface{}**](map[string]interface{}.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

