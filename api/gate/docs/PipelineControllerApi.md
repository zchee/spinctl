# \PipelineControllerApi

All URIs are relative to *http://localhost:8084*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelPipelineUsingPUT1**](PipelineControllerApi.md#CancelPipelineUsingPUT1) | **Put** /pipelines/{id}/cancel | Cancel a pipeline execution
[**DeletePipelineUsingDELETE**](PipelineControllerApi.md#DeletePipelineUsingDELETE) | **Delete** /pipelines/{application}/{pipelineName} | Delete a pipeline definition
[**DeletePipelineUsingDELETE1**](PipelineControllerApi.md#DeletePipelineUsingDELETE1) | **Delete** /pipelines/{id} | Delete a pipeline execution
[**EvaluateExpressionForExecutionAtStageUsingDELETE**](PipelineControllerApi.md#EvaluateExpressionForExecutionAtStageUsingDELETE) | **Delete** /pipelines/{id}/{stageId}/evaluateExpression | Evaluate a pipeline expression at a specific stage using the provided execution as context
[**EvaluateExpressionForExecutionAtStageUsingGET**](PipelineControllerApi.md#EvaluateExpressionForExecutionAtStageUsingGET) | **Get** /pipelines/{id}/{stageId}/evaluateExpression | Evaluate a pipeline expression at a specific stage using the provided execution as context
[**EvaluateExpressionForExecutionAtStageUsingHEAD**](PipelineControllerApi.md#EvaluateExpressionForExecutionAtStageUsingHEAD) | **Head** /pipelines/{id}/{stageId}/evaluateExpression | Evaluate a pipeline expression at a specific stage using the provided execution as context
[**EvaluateExpressionForExecutionAtStageUsingOPTIONS**](PipelineControllerApi.md#EvaluateExpressionForExecutionAtStageUsingOPTIONS) | **Options** /pipelines/{id}/{stageId}/evaluateExpression | Evaluate a pipeline expression at a specific stage using the provided execution as context
[**EvaluateExpressionForExecutionAtStageUsingPATCH**](PipelineControllerApi.md#EvaluateExpressionForExecutionAtStageUsingPATCH) | **Patch** /pipelines/{id}/{stageId}/evaluateExpression | Evaluate a pipeline expression at a specific stage using the provided execution as context
[**EvaluateExpressionForExecutionAtStageUsingPOST**](PipelineControllerApi.md#EvaluateExpressionForExecutionAtStageUsingPOST) | **Post** /pipelines/{id}/{stageId}/evaluateExpression | Evaluate a pipeline expression at a specific stage using the provided execution as context
[**EvaluateExpressionForExecutionAtStageUsingPUT**](PipelineControllerApi.md#EvaluateExpressionForExecutionAtStageUsingPUT) | **Put** /pipelines/{id}/{stageId}/evaluateExpression | Evaluate a pipeline expression at a specific stage using the provided execution as context
[**EvaluateExpressionForExecutionUsingDELETE**](PipelineControllerApi.md#EvaluateExpressionForExecutionUsingDELETE) | **Delete** /pipelines/{id}/evaluateExpression | Evaluate a pipeline expression using the provided execution as context
[**EvaluateExpressionForExecutionUsingGET**](PipelineControllerApi.md#EvaluateExpressionForExecutionUsingGET) | **Get** /pipelines/{id}/evaluateExpression | Evaluate a pipeline expression using the provided execution as context
[**EvaluateExpressionForExecutionUsingHEAD**](PipelineControllerApi.md#EvaluateExpressionForExecutionUsingHEAD) | **Head** /pipelines/{id}/evaluateExpression | Evaluate a pipeline expression using the provided execution as context
[**EvaluateExpressionForExecutionUsingOPTIONS**](PipelineControllerApi.md#EvaluateExpressionForExecutionUsingOPTIONS) | **Options** /pipelines/{id}/evaluateExpression | Evaluate a pipeline expression using the provided execution as context
[**EvaluateExpressionForExecutionUsingPATCH**](PipelineControllerApi.md#EvaluateExpressionForExecutionUsingPATCH) | **Patch** /pipelines/{id}/evaluateExpression | Evaluate a pipeline expression using the provided execution as context
[**EvaluateExpressionForExecutionUsingPOST**](PipelineControllerApi.md#EvaluateExpressionForExecutionUsingPOST) | **Post** /pipelines/{id}/evaluateExpression | Evaluate a pipeline expression using the provided execution as context
[**EvaluateExpressionForExecutionUsingPUT**](PipelineControllerApi.md#EvaluateExpressionForExecutionUsingPUT) | **Put** /pipelines/{id}/evaluateExpression | Evaluate a pipeline expression using the provided execution as context
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

> CancelPipelineUsingPUT1(ctx, id, optional)

Cancel a pipeline execution

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 
 **optional** | ***CancelPipelineUsingPUT1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CancelPipelineUsingPUT1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| force | [default to false]
 **reason** | **optional.String**| reason | 

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

> DeletePipelineUsingDELETE(ctx, application, pipelineName)

Delete a pipeline definition

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string**| application | 
**pipelineName** | **string**| pipelineName | 

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

> map[string]map[string]interface{} DeletePipelineUsingDELETE1(ctx, id)

Delete a pipeline execution

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 

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


## EvaluateExpressionForExecutionAtStageUsingDELETE

> map[string]map[string]interface{} EvaluateExpressionForExecutionAtStageUsingDELETE(ctx, expression, id, stageId)

Evaluate a pipeline expression at a specific stage using the provided execution as context

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**expression** | **string**| expression | 
**id** | **string**| id | 
**stageId** | **string**| stageId | 

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

> map[string]map[string]interface{} EvaluateExpressionForExecutionAtStageUsingGET(ctx, expression, id, stageId)

Evaluate a pipeline expression at a specific stage using the provided execution as context

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**expression** | **string**| expression | 
**id** | **string**| id | 
**stageId** | **string**| stageId | 

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


## EvaluateExpressionForExecutionAtStageUsingHEAD

> map[string]map[string]interface{} EvaluateExpressionForExecutionAtStageUsingHEAD(ctx, expression, id, stageId)

Evaluate a pipeline expression at a specific stage using the provided execution as context

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**expression** | **string**| expression | 
**id** | **string**| id | 
**stageId** | **string**| stageId | 

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


## EvaluateExpressionForExecutionAtStageUsingOPTIONS

> map[string]map[string]interface{} EvaluateExpressionForExecutionAtStageUsingOPTIONS(ctx, expression, id, stageId)

Evaluate a pipeline expression at a specific stage using the provided execution as context

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**expression** | **string**| expression | 
**id** | **string**| id | 
**stageId** | **string**| stageId | 

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


## EvaluateExpressionForExecutionAtStageUsingPATCH

> map[string]map[string]interface{} EvaluateExpressionForExecutionAtStageUsingPATCH(ctx, expression, id, stageId)

Evaluate a pipeline expression at a specific stage using the provided execution as context

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**expression** | **string**| expression | 
**id** | **string**| id | 
**stageId** | **string**| stageId | 

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


## EvaluateExpressionForExecutionAtStageUsingPOST

> map[string]map[string]interface{} EvaluateExpressionForExecutionAtStageUsingPOST(ctx, expression, id, stageId)

Evaluate a pipeline expression at a specific stage using the provided execution as context

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**expression** | **string**| expression | 
**id** | **string**| id | 
**stageId** | **string**| stageId | 

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


## EvaluateExpressionForExecutionAtStageUsingPUT

> map[string]map[string]interface{} EvaluateExpressionForExecutionAtStageUsingPUT(ctx, expression, id, stageId)

Evaluate a pipeline expression at a specific stage using the provided execution as context

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**expression** | **string**| expression | 
**id** | **string**| id | 
**stageId** | **string**| stageId | 

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


## EvaluateExpressionForExecutionUsingDELETE

> map[string]map[string]interface{} EvaluateExpressionForExecutionUsingDELETE(ctx, expression, id)

Evaluate a pipeline expression using the provided execution as context

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**expression** | **string**| expression | 
**id** | **string**| id | 

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

> map[string]map[string]interface{} EvaluateExpressionForExecutionUsingGET(ctx, expression, id)

Evaluate a pipeline expression using the provided execution as context

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**expression** | **string**| expression | 
**id** | **string**| id | 

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


## EvaluateExpressionForExecutionUsingHEAD

> map[string]map[string]interface{} EvaluateExpressionForExecutionUsingHEAD(ctx, expression, id)

Evaluate a pipeline expression using the provided execution as context

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**expression** | **string**| expression | 
**id** | **string**| id | 

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


## EvaluateExpressionForExecutionUsingOPTIONS

> map[string]map[string]interface{} EvaluateExpressionForExecutionUsingOPTIONS(ctx, expression, id)

Evaluate a pipeline expression using the provided execution as context

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**expression** | **string**| expression | 
**id** | **string**| id | 

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


## EvaluateExpressionForExecutionUsingPATCH

> map[string]map[string]interface{} EvaluateExpressionForExecutionUsingPATCH(ctx, expression, id)

Evaluate a pipeline expression using the provided execution as context

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**expression** | **string**| expression | 
**id** | **string**| id | 

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


## EvaluateExpressionForExecutionUsingPOST

> map[string]map[string]interface{} EvaluateExpressionForExecutionUsingPOST(ctx, expression, id)

Evaluate a pipeline expression using the provided execution as context

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**expression** | **string**| expression | 
**id** | **string**| id | 

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


## EvaluateExpressionForExecutionUsingPUT

> map[string]map[string]interface{} EvaluateExpressionForExecutionUsingPUT(ctx, expression, id)

Evaluate a pipeline expression using the provided execution as context

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**expression** | **string**| expression | 
**id** | **string**| id | 

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


## GetPipelineUsingGET

> map[string]interface{} GetPipelineUsingGET(ctx, id)

Retrieve a pipeline execution

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 

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

> HttpEntity InvokePipelineConfigUsingPOST1(ctx, application, pipelineNameOrId, optional)

Trigger a pipeline execution

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string**| application | 
**pipelineNameOrId** | **string**| pipelineNameOrId | 
 **optional** | ***InvokePipelineConfigUsingPOST1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a InvokePipelineConfigUsingPOST1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **trigger** | **optional.Interface**| trigger | 

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


## InvokePipelineConfigViaEchoUsingPOST

> HttpEntity InvokePipelineConfigViaEchoUsingPOST(ctx, application, pipelineNameOrId, optional)

Trigger a pipeline execution

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**application** | **string**| application | 
**pipelineNameOrId** | **string**| pipelineNameOrId | 
 **optional** | ***InvokePipelineConfigViaEchoUsingPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a InvokePipelineConfigViaEchoUsingPOSTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **trigger** | **optional.Interface**| trigger | 

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

> PausePipelineUsingPUT(ctx, id)

Pause a pipeline execution

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 

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

> RenamePipelineUsingPOST(ctx, renameCommand)

Rename a pipeline definition

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**renameCommand** | **map[string]interface{}**| renameCommand | 

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

> map[string]map[string]interface{} RestartStageUsingPUT(ctx, context, id, stageId)

Restart a stage execution

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**context** | **map[string]interface{}**| context | 
**id** | **string**| id | 
**stageId** | **string**| stageId | 

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

> map[string]map[string]interface{} ResumePipelineUsingPUT(ctx, id)

Resume a pipeline execution

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 

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

> SavePipelineUsingPOST(ctx, pipeline)

Save a pipeline definition

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipeline** | **map[string]interface{}**| pipeline | 

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

> ResponseEntity StartUsingPOST(ctx, map_)

Initiate a pipeline execution

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**map_** | **map[string]interface{}**| map | 

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

> map[string]map[string]interface{} UpdatePipelineUsingPUT(ctx, pipeline, id)

Update a pipeline definition

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipeline** | **map[string]interface{}**| pipeline | 
**id** | **string**| id | 

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

> map[string]map[string]interface{} UpdateStageUsingPATCH(ctx, context, id, stageId)

Update a stage execution

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**context** | **map[string]interface{}**| context | 
**id** | **string**| id | 
**stageId** | **string**| stageId | 

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

