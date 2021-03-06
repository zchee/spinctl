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

// JobControllerApiService JobControllerApi service
type JobControllerApiService service

type apiGetJobUsingGETRequest struct {
	ctx             _context.Context
	apiService      *JobControllerApiService
	account         string
	applicationName string
	name            string
	region          string
	xRateLimitApp   *string
	expand          *string
}

func (r apiGetJobUsingGETRequest) XRateLimitApp(xRateLimitApp string) apiGetJobUsingGETRequest {
	r.xRateLimitApp = &xRateLimitApp
	return r
}

func (r apiGetJobUsingGETRequest) Expand(expand string) apiGetJobUsingGETRequest {
	r.expand = &expand
	return r
}

/*
GetJobUsingGET Get job
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param account account
 * @param applicationName applicationName
 * @param name name
 * @param region region
@return apiGetJobUsingGETRequest
*/
func (a *JobControllerApiService) GetJobUsingGET(ctx _context.Context, account string, applicationName string, name string, region string) apiGetJobUsingGETRequest {
	return apiGetJobUsingGETRequest{
		apiService:      a,
		ctx:             ctx,
		account:         account,
		applicationName: applicationName,
		name:            name,
		region:          region,
	}
}

/*
Execute executes the request
 @return map[string]map[string]interface{}
*/
func (r apiGetJobUsingGETRequest) Execute() (map[string]map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]map[string]interface{}
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "JobControllerApiService.GetJobUsingGET")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/applications/{applicationName}/jobs/{account}/{region}/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"account"+"}", _neturl.QueryEscape(parameterToString(r.account, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"applicationName"+"}", _neturl.QueryEscape(parameterToString(r.applicationName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", _neturl.QueryEscape(parameterToString(r.name, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"region"+"}", _neturl.QueryEscape(parameterToString(r.region, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.expand != nil {
		localVarQueryParams.Add("expand", parameterToString(*r.expand, ""))
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
	if r.xRateLimitApp != nil {
		localVarHeaderParams["X-RateLimit-App"] = parameterToString(*r.xRateLimitApp, "")
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
			var v map[string]map[string]interface{}
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
