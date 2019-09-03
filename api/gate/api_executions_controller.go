/*
 * Spinnaker API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gate

import (
	_context "context"
	"fmt"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"

	"github.com/antihax/optional"
)

// Linger please

var _ _context.Context

type ExecutionsControllerApiService service

/*
ExecutionsControllerApiService Retrieves an ad-hoc collection of executions based on a number of user-supplied parameters. Either executionIds or pipelineConfigIds must be supplied in order to return any results. If both are supplied, an exception will be thrown.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *GetLatestExecutionsByConfigIdsUsingGETOpts - Optional Parameters:
 * @param "PipelineConfigIds" (optional.String) -  A comma-separated list of pipeline configuration IDs to retrieve recent executions for. Either this OR pipelineConfigIds must be supplied, but not both.
 * @param "ExecutionIds" (optional.String) -  A comma-separated list of executions to retrieve. Either this OR pipelineConfigIds must be supplied, but not both.
 * @param "Limit" (optional.Int32) -  The number of executions to return per pipeline configuration. Ignored if executionIds parameter is supplied. If this value is missing, it is defaulted to 1.
 * @param "Statuses" (optional.String) -  A comma-separated list of execution statuses to filter by. Ignored if executionIds parameter is supplied. If this value is missing, it is defaulted to all statuses.
 * @param "Expand" (optional.Bool) -  Expands each execution object in the resulting list. If this value is missing, it is defaulted to true.
@return []map[string]interface{}
*/

type GetLatestExecutionsByConfigIdsUsingGETOpts struct {
	PipelineConfigIds optional.String
	ExecutionIds      optional.String
	Limit             optional.Int32
	Statuses          optional.String
	Expand            optional.Bool
}

func (a *ExecutionsControllerApiService) GetLatestExecutionsByConfigIdsUsingGET(ctx _context.Context, localVarOptionals *GetLatestExecutionsByConfigIdsUsingGETOpts) ([]map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHttpMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []map[string]interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/executions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.PipelineConfigIds.IsSet() {
		localVarQueryParams.Add("pipelineConfigIds", parameterToString(localVarOptionals.PipelineConfigIds.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ExecutionIds.IsSet() {
		localVarQueryParams.Add("executionIds", parameterToString(localVarOptionals.ExecutionIds.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Statuses.IsSet() {
		localVarQueryParams.Add("statuses", parameterToString(localVarOptionals.Statuses.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Expand.IsSet() {
		localVarQueryParams.Add("expand", parameterToString(localVarOptionals.Expand.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []map[string]interface{}
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
ExecutionsControllerApiService Search for pipeline executions using a combination of criteria. The returned list is sorted by buildTime (trigger time) in reverse order so that newer executions are first in the list.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param application Only includes executions that are part of this application. If this value is \"*\", results will include executions of all applications.
 * @param optional nil or *SearchForPipelineExecutionsByTriggerUsingGETOpts - Optional Parameters:
 * @param "TriggerTypes" (optional.String) -  Only includes executions that were triggered by a trigger with a type that is equal to a type provided in this field. The list of trigger types should be a comma-delimited string. If this value is missing, results will includes executions of all trigger types.
 * @param "PipelineName" (optional.String) -  Only includes executions that with this pipeline name.
 * @param "EventId" (optional.String) -  Only includes executions that were triggered by a trigger with this eventId.
 * @param "Trigger" (optional.String) -  Only includes executions that were triggered by a trigger that matches the subset of fields provided by this value. This value should be a base64-encoded string of a JSON representation of a trigger object. The comparison succeeds if the execution trigger contains all the fields of the input trigger, the fields are of the same type, and each value of the field \"matches\". The term \"matches\" is specific for each field's type: - For Strings: A String value in the execution's trigger matches the input trigger's String value if the former equals the latter (case-insensitive) OR if the former matches the latter as a regular expression. - For Maps: A Map value in the execution's trigger matches the input trigger's Map value if the former contains all keys of the latter and their values match. - For Collections: A Collection value in the execution's trigger matches the input trigger's Collection value if the former has a unique element that matches each element of the latter. - Every other value is compared using the Java \"equals\" method (Groovy \"==\" operator)
 * @param "TriggerTimeStartBoundary" (optional.Int64) -  Only includes executions that were built at or after the given time, represented as a Unix timestamp in ms (UTC). This value must be >= 0 and <= the value of [triggerTimeEndBoundary], if provided. If this value is missing, it is defaulted to 0.
 * @param "TriggerTimeEndBoundary" (optional.Int64) -  Only includes executions that were built at or before the given time, represented as a Unix timestamp in ms (UTC). This value must be <= 9223372036854775807 (Long.MAX_VALUE) and >= the value of [triggerTimeStartBoundary], if provided. If this value is missing, it is defaulted to 9223372036854775807.
 * @param "Statuses" (optional.String) -  Only includes executions with a status that is equal to a status provided in this field. The list of statuses should be given as a comma-delimited string. If this value is missing, includes executions of all statuses. Allowed statuses are: NOT_STARTED, RUNNING, PAUSED, SUSPENDED, SUCCEEDED, FAILED_CONTINUE, TERMINAL, CANCELED, REDIRECT, STOPPED, SKIPPED, BUFFERED.
 * @param "StartIndex" (optional.Int32) -  Sets the first item of the resulting list for pagination. The list is 0-indexed. This value must be >= 0. If this value is missing, it is defaulted to 0.
 * @param "Size" (optional.Int32) -  Sets the size of the resulting list for pagination. This value must be > 0. If this value is missing, it is defaulted to 10.
 * @param "Reverse" (optional.Bool) -  Reverses the resulting list before it is paginated. If this value is missing, it is defaulted to false.
 * @param "Expand" (optional.Bool) -  Expands each execution object in the resulting list. If this value is missing, it is defaulted to false.
@return []map[string]interface{}
*/

type SearchForPipelineExecutionsByTriggerUsingGETOpts struct {
	TriggerTypes             optional.String
	PipelineName             optional.String
	EventId                  optional.String
	Trigger                  optional.String
	TriggerTimeStartBoundary optional.Int64
	TriggerTimeEndBoundary   optional.Int64
	Statuses                 optional.String
	StartIndex               optional.Int32
	Size                     optional.Int32
	Reverse                  optional.Bool
	Expand                   optional.Bool
}

func (a *ExecutionsControllerApiService) SearchForPipelineExecutionsByTriggerUsingGET(ctx _context.Context, application string, localVarOptionals *SearchForPipelineExecutionsByTriggerUsingGETOpts) ([]map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHttpMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []map[string]interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/applications/{application}/executions/search"
	localVarPath = strings.Replace(localVarPath, "{"+"application"+"}", _neturl.QueryEscape(fmt.Sprintf("%v", application)), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if localVarOptionals != nil && localVarOptionals.TriggerTypes.IsSet() {
		localVarQueryParams.Add("triggerTypes", parameterToString(localVarOptionals.TriggerTypes.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PipelineName.IsSet() {
		localVarQueryParams.Add("pipelineName", parameterToString(localVarOptionals.PipelineName.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.EventId.IsSet() {
		localVarQueryParams.Add("eventId", parameterToString(localVarOptionals.EventId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Trigger.IsSet() {
		localVarQueryParams.Add("trigger", parameterToString(localVarOptionals.Trigger.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TriggerTimeStartBoundary.IsSet() {
		localVarQueryParams.Add("triggerTimeStartBoundary", parameterToString(localVarOptionals.TriggerTimeStartBoundary.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TriggerTimeEndBoundary.IsSet() {
		localVarQueryParams.Add("triggerTimeEndBoundary", parameterToString(localVarOptionals.TriggerTimeEndBoundary.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Statuses.IsSet() {
		localVarQueryParams.Add("statuses", parameterToString(localVarOptionals.Statuses.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.StartIndex.IsSet() {
		localVarQueryParams.Add("startIndex", parameterToString(localVarOptionals.StartIndex.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Size.IsSet() {
		localVarQueryParams.Add("size", parameterToString(localVarOptionals.Size.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Reverse.IsSet() {
		localVarQueryParams.Add("reverse", parameterToString(localVarOptionals.Reverse.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Expand.IsSet() {
		localVarQueryParams.Add("expand", parameterToString(localVarOptionals.Expand.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []map[string]interface{}
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
