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

// SubnetControllerApiService SubnetControllerApi service
type SubnetControllerApiService service

type apiListSubnetsByCloudProviderRequest struct {
	ctx           _context.Context
	apiService    *SubnetControllerApiService
	cloudProvider string
	xRateLimitApp *string
}

func (r apiListSubnetsByCloudProviderRequest) XRateLimitApp(xRateLimitApp string) apiListSubnetsByCloudProviderRequest {
	r.xRateLimitApp = &xRateLimitApp
	return r
}

/*
ListSubnetsByCloudProvider Retrieve a list of subnets for a given cloud provider
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param cloudProvider cloudProvider
@return apiListSubnetsByCloudProviderRequest
*/
func (a *SubnetControllerApiService) ListSubnetsByCloudProvider(ctx _context.Context, cloudProvider string) apiListSubnetsByCloudProviderRequest {
	return apiListSubnetsByCloudProviderRequest{
		apiService:    a,
		ctx:           ctx,
		cloudProvider: cloudProvider,
	}
}

/*
Execute executes the request
 @return []map[string]interface{}
*/
func (r apiListSubnetsByCloudProviderRequest) Execute() ([]map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "SubnetControllerApiService.ListSubnetsByCloudProvider")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subnets/{cloudProvider}"
	localVarPath = strings.Replace(localVarPath, "{"+"cloudProvider"+"}", _neturl.QueryEscape(parameterToString(r.cloudProvider, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
