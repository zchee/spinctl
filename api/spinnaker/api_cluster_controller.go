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

// ClusterControllerApiService ClusterControllerApi service
type ClusterControllerApiService service

type apiGetClusterLoadBalancersRequest struct {
	ctx           _context.Context
	apiService    *ClusterControllerApiService
	account       string
	application   string
	clusterName   string
	type_         string
	xRateLimitApp *string
}

func (r apiGetClusterLoadBalancersRequest) XRateLimitApp(xRateLimitApp string) apiGetClusterLoadBalancersRequest {
	r.xRateLimitApp = &xRateLimitApp
	return r
}

/*
GetClusterLoadBalancers Retrieve a cluster's loadbalancers
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param account account
 * @param application application
 * @param clusterName clusterName
 * @param type_ type
@return apiGetClusterLoadBalancersRequest
*/
func (a *ClusterControllerApiService) GetClusterLoadBalancers(ctx _context.Context, account string, application string, clusterName string, type_ string) apiGetClusterLoadBalancersRequest {
	return apiGetClusterLoadBalancersRequest{
		apiService:  a,
		ctx:         ctx,
		account:     account,
		application: application,
		clusterName: clusterName,
		type_:       type_,
	}
}

/*
Execute executes the request
 @return []map[string]interface{}
*/
func (r apiGetClusterLoadBalancersRequest) Execute() ([]map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "ClusterControllerApiService.GetClusterLoadBalancers")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/applications/{application}/clusters/{account}/{clusterName}/{type}/loadBalancers"
	localVarPath = strings.Replace(localVarPath, "{"+"account"+"}", _neturl.QueryEscape(parameterToString(r.account, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"application"+"}", _neturl.QueryEscape(parameterToString(r.application, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.QueryEscape(parameterToString(r.clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"type"+"}", _neturl.QueryEscape(parameterToString(r.type_, "")), -1)

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

type apiGetClustersDetailsRequest struct {
	ctx           _context.Context
	apiService    *ClusterControllerApiService
	account       string
	application   string
	clusterName   string
	xRateLimitApp *string
}

func (r apiGetClustersDetailsRequest) XRateLimitApp(xRateLimitApp string) apiGetClustersDetailsRequest {
	r.xRateLimitApp = &xRateLimitApp
	return r
}

/*
GetClustersDetails Retrieve a cluster's details
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param account account
 * @param application application
 * @param clusterName clusterName
@return apiGetClustersDetailsRequest
*/
func (a *ClusterControllerApiService) GetClustersDetails(ctx _context.Context, account string, application string, clusterName string) apiGetClustersDetailsRequest {
	return apiGetClustersDetailsRequest{
		apiService:  a,
		ctx:         ctx,
		account:     account,
		application: application,
		clusterName: clusterName,
	}
}

/*
Execute executes the request
 @return map[string]map[string]interface{}
*/
func (r apiGetClustersDetailsRequest) Execute() (map[string]map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]map[string]interface{}
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "ClusterControllerApiService.GetClustersDetails")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/applications/{application}/clusters/{account}/{clusterName}"
	localVarPath = strings.Replace(localVarPath, "{"+"account"+"}", _neturl.QueryEscape(parameterToString(r.account, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"application"+"}", _neturl.QueryEscape(parameterToString(r.application, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.QueryEscape(parameterToString(r.clusterName, "")), -1)

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

type apiGetServerGroupsRequest struct {
	ctx             _context.Context
	apiService      *ClusterControllerApiService
	account         string
	application     string
	clusterName     string
	serverGroupName string
	xRateLimitApp   *string
}

func (r apiGetServerGroupsRequest) XRateLimitApp(xRateLimitApp string) apiGetServerGroupsRequest {
	r.xRateLimitApp = &xRateLimitApp
	return r
}

/*
GetServerGroups Retrieve a server group's details
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param account account
 * @param application application
 * @param clusterName clusterName
 * @param serverGroupName serverGroupName
@return apiGetServerGroupsRequest
*/
func (a *ClusterControllerApiService) GetServerGroups(ctx _context.Context, account string, application string, clusterName string, serverGroupName string) apiGetServerGroupsRequest {
	return apiGetServerGroupsRequest{
		apiService:      a,
		ctx:             ctx,
		account:         account,
		application:     application,
		clusterName:     clusterName,
		serverGroupName: serverGroupName,
	}
}

/*
Execute executes the request
 @return []map[string]interface{}
*/
func (r apiGetServerGroupsRequest) Execute() ([]map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "ClusterControllerApiService.GetServerGroups")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/applications/{application}/clusters/{account}/{clusterName}/serverGroups/{serverGroupName}"
	localVarPath = strings.Replace(localVarPath, "{"+"account"+"}", _neturl.QueryEscape(parameterToString(r.account, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"application"+"}", _neturl.QueryEscape(parameterToString(r.application, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.QueryEscape(parameterToString(r.clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"serverGroupName"+"}", _neturl.QueryEscape(parameterToString(r.serverGroupName, "")), -1)

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

type apiGetTargetServerGroupRequest struct {
	ctx            _context.Context
	apiService     *ClusterControllerApiService
	account        string
	application    string
	cloudProvider  string
	clusterName    string
	scope          string
	target         string
	xRateLimitApp  *string
	onlyEnabled    *bool
	validateOldest *bool
}

func (r apiGetTargetServerGroupRequest) XRateLimitApp(xRateLimitApp string) apiGetTargetServerGroupRequest {
	r.xRateLimitApp = &xRateLimitApp
	return r
}

func (r apiGetTargetServerGroupRequest) OnlyEnabled(onlyEnabled bool) apiGetTargetServerGroupRequest {
	r.onlyEnabled = &onlyEnabled
	return r
}

func (r apiGetTargetServerGroupRequest) ValidateOldest(validateOldest bool) apiGetTargetServerGroupRequest {
	r.validateOldest = &validateOldest
	return r
}

/*
GetTargetServerGroup Retrieve a server group that matches a target coordinate (e.g., newest, ancestor) relative to a cluster
`scope` is either a zone or a region
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param account account
 * @param application application
 * @param cloudProvider cloudProvider
 * @param clusterName clusterName
 * @param scope scope
 * @param target target
@return apiGetTargetServerGroupRequest
*/
func (a *ClusterControllerApiService) GetTargetServerGroup(ctx _context.Context, account string, application string, cloudProvider string, clusterName string, scope string, target string) apiGetTargetServerGroupRequest {
	return apiGetTargetServerGroupRequest{
		apiService:    a,
		ctx:           ctx,
		account:       account,
		application:   application,
		cloudProvider: cloudProvider,
		clusterName:   clusterName,
		scope:         scope,
		target:        target,
	}
}

/*
Execute executes the request
 @return map[string]map[string]interface{}
*/
func (r apiGetTargetServerGroupRequest) Execute() (map[string]map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]map[string]interface{}
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "ClusterControllerApiService.GetTargetServerGroup")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/applications/{application}/clusters/{account}/{clusterName}/{cloudProvider}/{scope}/serverGroups/target/{target}"
	localVarPath = strings.Replace(localVarPath, "{"+"account"+"}", _neturl.QueryEscape(parameterToString(r.account, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"application"+"}", _neturl.QueryEscape(parameterToString(r.application, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"cloudProvider"+"}", _neturl.QueryEscape(parameterToString(r.cloudProvider, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.QueryEscape(parameterToString(r.clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"scope"+"}", _neturl.QueryEscape(parameterToString(r.scope, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"target"+"}", _neturl.QueryEscape(parameterToString(r.target, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.onlyEnabled != nil {
		localVarQueryParams.Add("onlyEnabled", parameterToString(*r.onlyEnabled, ""))
	}
	if r.validateOldest != nil {
		localVarQueryParams.Add("validateOldest", parameterToString(*r.validateOldest, ""))
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

type apiListClustersRequest struct {
	ctx           _context.Context
	apiService    *ClusterControllerApiService
	application   string
	xRateLimitApp *string
}

func (r apiListClustersRequest) XRateLimitApp(xRateLimitApp string) apiListClustersRequest {
	r.xRateLimitApp = &xRateLimitApp
	return r
}

/*
ListClusters Retrieve a list of cluster names for an application, grouped by account
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param application application
@return apiListClustersRequest
*/
func (a *ClusterControllerApiService) ListClusters(ctx _context.Context, application string) apiListClustersRequest {
	return apiListClustersRequest{
		apiService:  a,
		ctx:         ctx,
		application: application,
	}
}

/*
Execute executes the request
 @return map[string]map[string]interface{}
*/
func (r apiListClustersRequest) Execute() (map[string]map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]map[string]interface{}
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "ClusterControllerApiService.ListClusters")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/applications/{application}/clusters"
	localVarPath = strings.Replace(localVarPath, "{"+"application"+"}", _neturl.QueryEscape(parameterToString(r.application, "")), -1)

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

type apiListClustersByAccountRequest struct {
	ctx           _context.Context
	apiService    *ClusterControllerApiService
	account       string
	application   string
	xRateLimitApp *string
}

func (r apiListClustersByAccountRequest) XRateLimitApp(xRateLimitApp string) apiListClustersByAccountRequest {
	r.xRateLimitApp = &xRateLimitApp
	return r
}

/*
ListClustersByAccount Retrieve a list of clusters for an account
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param account account
 * @param application application
@return apiListClustersByAccountRequest
*/
func (a *ClusterControllerApiService) ListClustersByAccount(ctx _context.Context, account string, application string) apiListClustersByAccountRequest {
	return apiListClustersByAccountRequest{
		apiService:  a,
		ctx:         ctx,
		account:     account,
		application: application,
	}
}

/*
Execute executes the request
 @return []map[string]interface{}
*/
func (r apiListClustersByAccountRequest) Execute() ([]map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "ClusterControllerApiService.ListClustersByAccount")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/applications/{application}/clusters/{account}"
	localVarPath = strings.Replace(localVarPath, "{"+"account"+"}", _neturl.QueryEscape(parameterToString(r.account, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"application"+"}", _neturl.QueryEscape(parameterToString(r.application, "")), -1)

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

type apiListScalingActivitiesRequest struct {
	ctx             _context.Context
	apiService      *ClusterControllerApiService
	account         string
	application     string
	clusterName     string
	serverGroupName string
	xRateLimitApp   *string
	provider        *string
	region          *string
}

func (r apiListScalingActivitiesRequest) XRateLimitApp(xRateLimitApp string) apiListScalingActivitiesRequest {
	r.xRateLimitApp = &xRateLimitApp
	return r
}

func (r apiListScalingActivitiesRequest) Provider(provider string) apiListScalingActivitiesRequest {
	r.provider = &provider
	return r
}

func (r apiListScalingActivitiesRequest) Region(region string) apiListScalingActivitiesRequest {
	r.region = &region
	return r
}

/*
ListScalingActivities Retrieve a list of scaling activities for a server group
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param account account
 * @param application application
 * @param clusterName clusterName
 * @param serverGroupName serverGroupName
@return apiListScalingActivitiesRequest
*/
func (a *ClusterControllerApiService) ListScalingActivities(ctx _context.Context, account string, application string, clusterName string, serverGroupName string) apiListScalingActivitiesRequest {
	return apiListScalingActivitiesRequest{
		apiService:      a,
		ctx:             ctx,
		account:         account,
		application:     application,
		clusterName:     clusterName,
		serverGroupName: serverGroupName,
	}
}

/*
Execute executes the request
 @return []map[string]interface{}
*/
func (r apiListScalingActivitiesRequest) Execute() ([]map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "ClusterControllerApiService.ListScalingActivities")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/applications/{application}/clusters/{account}/{clusterName}/serverGroups/{serverGroupName}/scalingActivities"
	localVarPath = strings.Replace(localVarPath, "{"+"account"+"}", _neturl.QueryEscape(parameterToString(r.account, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"application"+"}", _neturl.QueryEscape(parameterToString(r.application, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.QueryEscape(parameterToString(r.clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"serverGroupName"+"}", _neturl.QueryEscape(parameterToString(r.serverGroupName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.provider != nil {
		localVarQueryParams.Add("provider", parameterToString(*r.provider, ""))
	}
	if r.region != nil {
		localVarQueryParams.Add("region", parameterToString(*r.region, ""))
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

type apiListServerGroupsRequest struct {
	ctx           _context.Context
	apiService    *ClusterControllerApiService
	account       string
	application   string
	clusterName   string
	xRateLimitApp *string
}

func (r apiListServerGroupsRequest) XRateLimitApp(xRateLimitApp string) apiListServerGroupsRequest {
	r.xRateLimitApp = &xRateLimitApp
	return r
}

/*
ListServerGroups Retrieve a list of server groups for a cluster
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param account account
 * @param application application
 * @param clusterName clusterName
@return apiListServerGroupsRequest
*/
func (a *ClusterControllerApiService) ListServerGroups(ctx _context.Context, account string, application string, clusterName string) apiListServerGroupsRequest {
	return apiListServerGroupsRequest{
		apiService:  a,
		ctx:         ctx,
		account:     account,
		application: application,
		clusterName: clusterName,
	}
}

/*
Execute executes the request
 @return []map[string]interface{}
*/
func (r apiListServerGroupsRequest) Execute() ([]map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "ClusterControllerApiService.ListServerGroups")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/applications/{application}/clusters/{account}/{clusterName}/serverGroups"
	localVarPath = strings.Replace(localVarPath, "{"+"account"+"}", _neturl.QueryEscape(parameterToString(r.account, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"application"+"}", _neturl.QueryEscape(parameterToString(r.application, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", _neturl.QueryEscape(parameterToString(r.clusterName, "")), -1)

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
