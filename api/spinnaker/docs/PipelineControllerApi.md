# \PipelineControllerApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelPipelineExecution**](PipelineControllerApi.md#CancelPipelineExecution) | **Put** /pipelines/{id}/cancel | Cancel a pipeline execution
[**DeletePipelineDefinition**](PipelineControllerApi.md#DeletePipelineDefinition) | **Delete** /pipelines/{application}/{pipelineName} | Delete a pipeline definition
[**DeletePipelineExecution**](PipelineControllerApi.md#DeletePipelineExecution) | **Delete** /pipelines/{id} | Delete a pipeline execution
[**EvaluateExpressionForExecution**](PipelineControllerApi.md#EvaluateExpressionForExecution) | **Get** /pipelines/{id}/evaluateExpression | Evaluate a pipeline expression using the provided execution as context
[**EvaluateExpressionForExecutionAtStage**](PipelineControllerApi.md#EvaluateExpressionForExecutionAtStage) | **Get** /pipelines/{id}/{stageId}/evaluateExpression | Evaluate a pipeline expression at a specific stage using the provided execution as context
[**EvaluateExpressionForExecutionViaPOST**](PipelineControllerApi.md#EvaluateExpressionForExecutionViaPOST) | **Post** /pipelines/{id}/evaluateExpression | Evaluate a pipeline expression using the provided execution as context
[**EvaluateVariables**](PipelineControllerApi.md#EvaluateVariables) | **Post** /pipelines/{id}/evaluateVariables | Evaluate variables same as Evaluate Variables stage using the provided execution as context
[**GetPipeline**](PipelineControllerApi.md#GetPipeline) | **Get** /pipelines/{id} | Retrieve a pipeline execution
[**InvokePipelineConfigViaEcho**](PipelineControllerApi.md#InvokePipelineConfigViaEcho) | **Post** /pipelines/v2/{application}/{pipelineNameOrId} | Trigger a pipeline execution
[**PausePipeline**](PipelineControllerApi.md#PausePipeline) | **Put** /pipelines/{id}/pause | Pause a pipeline execution
[**RenamePipeline**](PipelineControllerApi.md#RenamePipeline) | **Post** /pipelines/move | Rename a pipeline definition
[**RestartStage**](PipelineControllerApi.md#RestartStage) | **Put** /pipelines/{id}/stages/{stageId}/restart | Restart a stage execution
[**ResumePipeline**](PipelineControllerApi.md#ResumePipeline) | **Put** /pipelines/{id}/resume | Resume a pipeline execution
[**SavePipeline**](PipelineControllerApi.md#SavePipeline) | **Post** /pipelines | Save a pipeline definition
[**Start**](PipelineControllerApi.md#Start) | **Post** /pipelines/start | Initiate a pipeline execution
[**TriggerPipeline**](PipelineControllerApi.md#TriggerPipeline) | **Post** /pipelines/{application}/{pipelineName} | Trigger a pipeline execution
[**UpdatePipeline**](PipelineControllerApi.md#UpdatePipeline) | **Put** /pipelines/{id} | Update a pipeline definition
[**UpdateStage**](PipelineControllerApi.md#UpdateStage) | **Patch** /pipelines/{id}/stages/{stageId} | Update a stage execution



## CancelPipelineExecution

> CancelPipelineExecution(ctx, id).Force(force).Reason(reason).Execute()

Cancel a pipeline execution

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelPipelineExecutionRequest struct via the builder pattern


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


## DeletePipelineDefinition

> DeletePipelineDefinition(ctx, application, pipelineName).Execute()

Delete a pipeline definition

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string** | application | 
**pipelineName** | **string** | pipelineName | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePipelineDefinitionRequest struct via the builder pattern


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


## DeletePipelineExecution

> map[string]map[string]interface{} DeletePipelineExecution(ctx, id).Execute()

Delete a pipeline execution

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePipelineExecutionRequest struct via the builder pattern


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


## EvaluateExpressionForExecution

> map[string]map[string]interface{} EvaluateExpressionForExecution(ctx, id).Expression(expression).Execute()

Evaluate a pipeline expression using the provided execution as context

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEvaluateExpressionForExecutionRequest struct via the builder pattern


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


## EvaluateExpressionForExecutionAtStage

> map[string]map[string]interface{} EvaluateExpressionForExecutionAtStage(ctx, id, stageId).Expression(expression).Execute()

Evaluate a pipeline expression at a specific stage using the provided execution as context

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 
**stageId** | **string** | stageId | 

### Other Parameters

Other parameters are passed through a pointer to a apiEvaluateExpressionForExecutionAtStageRequest struct via the builder pattern


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


## EvaluateExpressionForExecutionViaPOST

> map[string]map[string]interface{} EvaluateExpressionForExecutionViaPOST(ctx, id).PipelineExpression(pipelineExpression).Execute()

Evaluate a pipeline expression using the provided execution as context

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEvaluateExpressionForExecutionViaPOSTRequest struct via the builder pattern


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


## EvaluateVariables

> map[string]map[string]interface{} EvaluateVariables(ctx, id).Expressions(expressions).ExecutionId(executionId).RequisiteStageRefIds(requisiteStageRefIds).SpelVersion(spelVersion).Execute()

Evaluate variables same as Evaluate Variables stage using the provided execution as context

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEvaluateVariablesRequest struct via the builder pattern


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


## GetPipeline

> map[string]interface{} GetPipeline(ctx, id).Execute()

Retrieve a pipeline execution

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelineRequest struct via the builder pattern


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


## InvokePipelineConfigViaEcho

> HttpEntity InvokePipelineConfigViaEcho(ctx, application, pipelineNameOrId).Trigger(trigger).Execute()

Trigger a pipeline execution

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string** | application | 
**pipelineNameOrId** | **string** | pipelineNameOrId | 

### Other Parameters

Other parameters are passed through a pointer to a apiInvokePipelineConfigViaEchoRequest struct via the builder pattern


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


## PausePipeline

> PausePipeline(ctx, id).Execute()

Pause a pipeline execution

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPausePipelineRequest struct via the builder pattern


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


## RenamePipeline

> RenamePipeline(ctx).RenameCommand(renameCommand).Execute()

Rename a pipeline definition

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRenamePipelineRequest struct via the builder pattern


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


## RestartStage

> map[string]map[string]interface{} RestartStage(ctx, id, stageId).Context(context).Execute()

Restart a stage execution

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 
**stageId** | **string** | stageId | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestartStageRequest struct via the builder pattern


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


## ResumePipeline

> map[string]map[string]interface{} ResumePipeline(ctx, id).Execute()

Resume a pipeline execution

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumePipelineRequest struct via the builder pattern


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


## SavePipeline

> SavePipeline(ctx).Pipeline(pipeline).Execute()

Save a pipeline definition

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSavePipelineRequest struct via the builder pattern


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


## Start

> ResponseEntity Start(ctx).Map_(map_).Execute()

Initiate a pipeline execution

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartRequest struct via the builder pattern


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


## TriggerPipeline

> TriggerPipeline(ctx, application, pipelineName).Trigger(trigger).Execute()

Trigger a pipeline execution

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string** | application | 
**pipelineName** | **string** | pipelineName | 

### Other Parameters

Other parameters are passed through a pointer to a apiTriggerPipelineRequest struct via the builder pattern


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


## UpdatePipeline

> map[string]map[string]interface{} UpdatePipeline(ctx, id).Pipeline(pipeline).Execute()

Update a pipeline definition

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePipelineRequest struct via the builder pattern


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


## UpdateStage

> map[string]map[string]interface{} UpdateStage(ctx, id, stageId).Context(context).Execute()

Update a stage execution

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 
**stageId** | **string** | stageId | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStageRequest struct via the builder pattern


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

