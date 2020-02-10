/*
 * Spinnaker API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.18.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gate

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// ExecutionsControllerApiService ExecutionsControllerApi service
type ExecutionsControllerApiService service

type apiGetLatestExecutionsByConfigIdsRequest struct {
	ctx               _context.Context
	apiService        *ExecutionsControllerApiService
	executionIds      *string
	expand            *bool
	limit             *int32
	pipelineConfigIds *string
	statuses          *string
}

func (r apiGetLatestExecutionsByConfigIdsRequest) ExecutionIds(executionIds string) apiGetLatestExecutionsByConfigIdsRequest {
	r.executionIds = &executionIds
	return r
}

func (r apiGetLatestExecutionsByConfigIdsRequest) Expand(expand bool) apiGetLatestExecutionsByConfigIdsRequest {
	r.expand = &expand
	return r
}

func (r apiGetLatestExecutionsByConfigIdsRequest) Limit(limit int32) apiGetLatestExecutionsByConfigIdsRequest {
	r.limit = &limit
	return r
}

func (r apiGetLatestExecutionsByConfigIdsRequest) PipelineConfigIds(pipelineConfigIds string) apiGetLatestExecutionsByConfigIdsRequest {
	r.pipelineConfigIds = &pipelineConfigIds
	return r
}

func (r apiGetLatestExecutionsByConfigIdsRequest) Statuses(statuses string) apiGetLatestExecutionsByConfigIdsRequest {
	r.statuses = &statuses
	return r
}

/*
GetLatestExecutionsByConfigIds Retrieves an ad-hoc collection of executions based on a number of user-supplied parameters. Either executionIds or pipelineConfigIds must be supplied in order to return any results. If both are supplied, an exception will be thrown.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return apiGetLatestExecutionsByConfigIdsRequest
*/
func (a *ExecutionsControllerApiService) GetLatestExecutionsByConfigIds(ctx _context.Context) apiGetLatestExecutionsByConfigIdsRequest {
	return apiGetLatestExecutionsByConfigIdsRequest{
		apiService: a,
		ctx:        ctx,
	}
}

/*
Execute executes the request
 @return []map[string]interface{}
*/
func (r apiGetLatestExecutionsByConfigIdsRequest) Execute() ([]map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "ExecutionsControllerApiService.GetLatestExecutionsByConfigIds")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/executions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.executionIds != nil {
		localVarQueryParams.Add("executionIds", parameterToString(*r.executionIds, ""))
	}
	if r.expand != nil {
		localVarQueryParams.Add("expand", parameterToString(*r.expand, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.pipelineConfigIds != nil {
		localVarQueryParams.Add("pipelineConfigIds", parameterToString(*r.pipelineConfigIds, ""))
	}
	if r.statuses != nil {
		localVarQueryParams.Add("statuses", parameterToString(*r.statuses, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v []map[string]interface{}
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = r.apiService.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type apiSearchForPipelineExecutionsByTriggerRequest struct {
	ctx                      _context.Context
	apiService               *ExecutionsControllerApiService
	application              string
	eventId                  *string
	expand                   *bool
	pipelineName             *string
	reverse                  *bool
	size                     *int32
	startIndex               *int32
	statuses                 *string
	trigger                  *string
	triggerTimeEndBoundary   *int64
	triggerTimeStartBoundary *int64
	triggerTypes             *string
}

func (r apiSearchForPipelineExecutionsByTriggerRequest) EventId(eventId string) apiSearchForPipelineExecutionsByTriggerRequest {
	r.eventId = &eventId
	return r
}

func (r apiSearchForPipelineExecutionsByTriggerRequest) Expand(expand bool) apiSearchForPipelineExecutionsByTriggerRequest {
	r.expand = &expand
	return r
}

func (r apiSearchForPipelineExecutionsByTriggerRequest) PipelineName(pipelineName string) apiSearchForPipelineExecutionsByTriggerRequest {
	r.pipelineName = &pipelineName
	return r
}

func (r apiSearchForPipelineExecutionsByTriggerRequest) Reverse(reverse bool) apiSearchForPipelineExecutionsByTriggerRequest {
	r.reverse = &reverse
	return r
}

func (r apiSearchForPipelineExecutionsByTriggerRequest) Size(size int32) apiSearchForPipelineExecutionsByTriggerRequest {
	r.size = &size
	return r
}

func (r apiSearchForPipelineExecutionsByTriggerRequest) StartIndex(startIndex int32) apiSearchForPipelineExecutionsByTriggerRequest {
	r.startIndex = &startIndex
	return r
}

func (r apiSearchForPipelineExecutionsByTriggerRequest) Statuses(statuses string) apiSearchForPipelineExecutionsByTriggerRequest {
	r.statuses = &statuses
	return r
}

func (r apiSearchForPipelineExecutionsByTriggerRequest) Trigger(trigger string) apiSearchForPipelineExecutionsByTriggerRequest {
	r.trigger = &trigger
	return r
}

func (r apiSearchForPipelineExecutionsByTriggerRequest) TriggerTimeEndBoundary(triggerTimeEndBoundary int64) apiSearchForPipelineExecutionsByTriggerRequest {
	r.triggerTimeEndBoundary = &triggerTimeEndBoundary
	return r
}

func (r apiSearchForPipelineExecutionsByTriggerRequest) TriggerTimeStartBoundary(triggerTimeStartBoundary int64) apiSearchForPipelineExecutionsByTriggerRequest {
	r.triggerTimeStartBoundary = &triggerTimeStartBoundary
	return r
}

func (r apiSearchForPipelineExecutionsByTriggerRequest) TriggerTypes(triggerTypes string) apiSearchForPipelineExecutionsByTriggerRequest {
	r.triggerTypes = &triggerTypes
	return r
}

/*
SearchForPipelineExecutionsByTrigger Search for pipeline executions using a combination of criteria. The returned list is sorted by buildTime (trigger time) in reverse order so that newer executions are first in the list.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param application Only includes executions that are part of this application. If this value is \"*\", results will include executions of all applications.
@return apiSearchForPipelineExecutionsByTriggerRequest
*/
func (a *ExecutionsControllerApiService) SearchForPipelineExecutionsByTrigger(ctx _context.Context, application string) apiSearchForPipelineExecutionsByTriggerRequest {
	return apiSearchForPipelineExecutionsByTriggerRequest{
		apiService:  a,
		ctx:         ctx,
		application: application,
	}
}

/*
Execute executes the request
 @return []map[string]interface{}
*/
func (r apiSearchForPipelineExecutionsByTriggerRequest) Execute() ([]map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "ExecutionsControllerApiService.SearchForPipelineExecutionsByTrigger")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/applications/{application}/executions/search"
	localVarPath = strings.Replace(localVarPath, "{"+"application"+"}", _neturl.QueryEscape(parameterToString(r.application, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.eventId != nil {
		localVarQueryParams.Add("eventId", parameterToString(*r.eventId, ""))
	}
	if r.expand != nil {
		localVarQueryParams.Add("expand", parameterToString(*r.expand, ""))
	}
	if r.pipelineName != nil {
		localVarQueryParams.Add("pipelineName", parameterToString(*r.pipelineName, ""))
	}
	if r.reverse != nil {
		localVarQueryParams.Add("reverse", parameterToString(*r.reverse, ""))
	}
	if r.size != nil {
		localVarQueryParams.Add("size", parameterToString(*r.size, ""))
	}
	if r.startIndex != nil {
		localVarQueryParams.Add("startIndex", parameterToString(*r.startIndex, ""))
	}
	if r.statuses != nil {
		localVarQueryParams.Add("statuses", parameterToString(*r.statuses, ""))
	}
	if r.trigger != nil {
		localVarQueryParams.Add("trigger", parameterToString(*r.trigger, ""))
	}
	if r.triggerTimeEndBoundary != nil {
		localVarQueryParams.Add("triggerTimeEndBoundary", parameterToString(*r.triggerTimeEndBoundary, ""))
	}
	if r.triggerTimeStartBoundary != nil {
		localVarQueryParams.Add("triggerTimeStartBoundary", parameterToString(*r.triggerTimeStartBoundary, ""))
	}
	if r.triggerTypes != nil {
		localVarQueryParams.Add("triggerTypes", parameterToString(*r.triggerTypes, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 200 {
			var v []map[string]interface{}
			err = r.apiService.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = r.apiService.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
