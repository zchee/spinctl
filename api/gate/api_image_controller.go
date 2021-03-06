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

// ImageControllerApiService ImageControllerApi service
type ImageControllerApiService service

type apiFindImagesUsingGETRequest struct {
	ctx        _context.Context
	apiService *ImageControllerApiService
	account    *string
	count      *int32
	provider   *string
	q          *string
	region     *string
}

func (r apiFindImagesUsingGETRequest) Account(account string) apiFindImagesUsingGETRequest {
	r.account = &account
	return r
}

func (r apiFindImagesUsingGETRequest) Count(count int32) apiFindImagesUsingGETRequest {
	r.count = &count
	return r
}

func (r apiFindImagesUsingGETRequest) Provider(provider string) apiFindImagesUsingGETRequest {
	r.provider = &provider
	return r
}

func (r apiFindImagesUsingGETRequest) Q(q string) apiFindImagesUsingGETRequest {
	r.q = &q
	return r
}

func (r apiFindImagesUsingGETRequest) Region(region string) apiFindImagesUsingGETRequest {
	r.region = &region
	return r
}

/*
FindImagesUsingGET Retrieve a list of images, filtered by cloud provider, region, and account
The query parameter `q` filters the list of images by image name
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return apiFindImagesUsingGETRequest
*/
func (a *ImageControllerApiService) FindImagesUsingGET(ctx _context.Context) apiFindImagesUsingGETRequest {
	return apiFindImagesUsingGETRequest{
		apiService: a,
		ctx:        ctx,
	}
}

/*
Execute executes the request
 @return []map[string]interface{}
*/
func (r apiFindImagesUsingGETRequest) Execute() ([]map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "ImageControllerApiService.FindImagesUsingGET")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/images/find"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.account != nil {
		localVarQueryParams.Add("account", parameterToString(*r.account, ""))
	}
	if r.count != nil {
		localVarQueryParams.Add("count", parameterToString(*r.count, ""))
	}
	if r.provider != nil {
		localVarQueryParams.Add("provider", parameterToString(*r.provider, ""))
	}
	if r.q != nil {
		localVarQueryParams.Add("q", parameterToString(*r.q, ""))
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

type apiFindTagsUsingGETRequest struct {
	ctx           _context.Context
	apiService    *ImageControllerApiService
	account       *string
	repository    *string
	xRateLimitApp *string
	provider      *string
}

func (r apiFindTagsUsingGETRequest) Account(account string) apiFindTagsUsingGETRequest {
	r.account = &account
	return r
}

func (r apiFindTagsUsingGETRequest) Repository(repository string) apiFindTagsUsingGETRequest {
	r.repository = &repository
	return r
}

func (r apiFindTagsUsingGETRequest) XRateLimitApp(xRateLimitApp string) apiFindTagsUsingGETRequest {
	r.xRateLimitApp = &xRateLimitApp
	return r
}

func (r apiFindTagsUsingGETRequest) Provider(provider string) apiFindTagsUsingGETRequest {
	r.provider = &provider
	return r
}

/*
FindTagsUsingGET Find tags
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@return apiFindTagsUsingGETRequest
*/
func (a *ImageControllerApiService) FindTagsUsingGET(ctx _context.Context) apiFindTagsUsingGETRequest {
	return apiFindTagsUsingGETRequest{
		apiService: a,
		ctx:        ctx,
	}
}

/*
Execute executes the request
 @return []map[string]interface{}
*/
func (r apiFindTagsUsingGETRequest) Execute() ([]map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "ImageControllerApiService.FindTagsUsingGET")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/images/tags"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.account == nil {
		return localVarReturnValue, nil, reportError("account is required and must be specified")
	}

	if r.repository == nil {
		return localVarReturnValue, nil, reportError("repository is required and must be specified")
	}

	localVarQueryParams.Add("account", parameterToString(*r.account, ""))
	if r.provider != nil {
		localVarQueryParams.Add("provider", parameterToString(*r.provider, ""))
	}
	localVarQueryParams.Add("repository", parameterToString(*r.repository, ""))
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

type apiGetImageDetailsUsingGETRequest struct {
	ctx           _context.Context
	apiService    *ImageControllerApiService
	account       string
	imageId       string
	region        string
	xRateLimitApp *string
	provider      *string
}

func (r apiGetImageDetailsUsingGETRequest) XRateLimitApp(xRateLimitApp string) apiGetImageDetailsUsingGETRequest {
	r.xRateLimitApp = &xRateLimitApp
	return r
}

func (r apiGetImageDetailsUsingGETRequest) Provider(provider string) apiGetImageDetailsUsingGETRequest {
	r.provider = &provider
	return r
}

/*
GetImageDetailsUsingGET Get image details
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param account account
 * @param imageId imageId
 * @param region region
@return apiGetImageDetailsUsingGETRequest
*/
func (a *ImageControllerApiService) GetImageDetailsUsingGET(ctx _context.Context, account string, imageId string, region string) apiGetImageDetailsUsingGETRequest {
	return apiGetImageDetailsUsingGETRequest{
		apiService: a,
		ctx:        ctx,
		account:    account,
		imageId:    imageId,
		region:     region,
	}
}

/*
Execute executes the request
 @return []map[string]interface{}
*/
func (r apiGetImageDetailsUsingGETRequest) Execute() ([]map[string]interface{}, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []map[string]interface{}
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "ImageControllerApiService.GetImageDetailsUsingGET")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/images/{account}/{region}/{imageId}"
	localVarPath = strings.Replace(localVarPath, "{"+"account"+"}", _neturl.QueryEscape(parameterToString(r.account, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"imageId"+"}", _neturl.QueryEscape(parameterToString(r.imageId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"region"+"}", _neturl.QueryEscape(parameterToString(r.region, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.provider != nil {
		localVarQueryParams.Add("provider", parameterToString(*r.provider, ""))
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
