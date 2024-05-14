/*
Emby Server REST API (BETA)

Explore the Emby Server API

API version: 4.8.0.61
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package embyclient

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"os"
)


// ToneMapOptionsServiceAPIService ToneMapOptionsServiceAPI service
type ToneMapOptionsServiceAPIService service

type ToneMapOptionsServiceAPIGetEncodingFulltonemapoptionsRequest struct {
	ctx context.Context
	ApiService *ToneMapOptionsServiceAPIService
}

func (r ToneMapOptionsServiceAPIGetEncodingFulltonemapoptionsRequest) Execute() (*ModelEditObjectContainer, *http.Response, error) {
	return r.ApiService.GetEncodingFulltonemapoptionsExecute(r)
}

/*
GetEncodingFulltonemapoptions Gets the tone mapping options

Requires authentication as user

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ToneMapOptionsServiceAPIGetEncodingFulltonemapoptionsRequest
*/
func (a *ToneMapOptionsServiceAPIService) GetEncodingFulltonemapoptions(ctx context.Context) ToneMapOptionsServiceAPIGetEncodingFulltonemapoptionsRequest {
	return ToneMapOptionsServiceAPIGetEncodingFulltonemapoptionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ModelEditObjectContainer
func (a *ToneMapOptionsServiceAPIService) GetEncodingFulltonemapoptionsExecute(r ToneMapOptionsServiceAPIGetEncodingFulltonemapoptionsRequest) (*ModelEditObjectContainer, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelEditObjectContainer
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ToneMapOptionsServiceAPIService.GetEncodingFulltonemapoptions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Encoding/FullToneMapOptions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikeyauth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("api_key", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ToneMapOptionsServiceAPIGetEncodingPublictonemapoptionsRequest struct {
	ctx context.Context
	ApiService *ToneMapOptionsServiceAPIService
}

func (r ToneMapOptionsServiceAPIGetEncodingPublictonemapoptionsRequest) Execute() (*ModelEditObjectContainer, *http.Response, error) {
	return r.ApiService.GetEncodingPublictonemapoptionsExecute(r)
}

/*
GetEncodingPublictonemapoptions Gets the tone mapping options

Requires authentication as user

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ToneMapOptionsServiceAPIGetEncodingPublictonemapoptionsRequest
*/
func (a *ToneMapOptionsServiceAPIService) GetEncodingPublictonemapoptions(ctx context.Context) ToneMapOptionsServiceAPIGetEncodingPublictonemapoptionsRequest {
	return ToneMapOptionsServiceAPIGetEncodingPublictonemapoptionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ModelEditObjectContainer
func (a *ToneMapOptionsServiceAPIService) GetEncodingPublictonemapoptionsExecute(r ToneMapOptionsServiceAPIGetEncodingPublictonemapoptionsRequest) (*ModelEditObjectContainer, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelEditObjectContainer
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ToneMapOptionsServiceAPIService.GetEncodingPublictonemapoptions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Encoding/PublicToneMapOptions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikeyauth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("api_key", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ToneMapOptionsServiceAPIPostEncodingFulltonemapoptionsRequest struct {
	ctx context.Context
	ApiService *ToneMapOptionsServiceAPIService
	body *os.File
}

// Binary stream
func (r ToneMapOptionsServiceAPIPostEncodingFulltonemapoptionsRequest) Body(body *os.File) ToneMapOptionsServiceAPIPostEncodingFulltonemapoptionsRequest {
	r.body = body
	return r
}

func (r ToneMapOptionsServiceAPIPostEncodingFulltonemapoptionsRequest) Execute() (*http.Response, error) {
	return r.ApiService.PostEncodingFulltonemapoptionsExecute(r)
}

/*
PostEncodingFulltonemapoptions Updates the tone mapping options

Requires authentication as administrator

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ToneMapOptionsServiceAPIPostEncodingFulltonemapoptionsRequest
*/
func (a *ToneMapOptionsServiceAPIService) PostEncodingFulltonemapoptions(ctx context.Context) ToneMapOptionsServiceAPIPostEncodingFulltonemapoptionsRequest {
	return ToneMapOptionsServiceAPIPostEncodingFulltonemapoptionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ToneMapOptionsServiceAPIService) PostEncodingFulltonemapoptionsExecute(r ToneMapOptionsServiceAPIPostEncodingFulltonemapoptionsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ToneMapOptionsServiceAPIService.PostEncodingFulltonemapoptions")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Encoding/FullToneMapOptions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/octet-stream"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikeyauth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("api_key", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ToneMapOptionsServiceAPIPostEncodingPublictonemapoptionsRequest struct {
	ctx context.Context
	ApiService *ToneMapOptionsServiceAPIService
	body *os.File
}

// Binary stream
func (r ToneMapOptionsServiceAPIPostEncodingPublictonemapoptionsRequest) Body(body *os.File) ToneMapOptionsServiceAPIPostEncodingPublictonemapoptionsRequest {
	r.body = body
	return r
}

func (r ToneMapOptionsServiceAPIPostEncodingPublictonemapoptionsRequest) Execute() (*http.Response, error) {
	return r.ApiService.PostEncodingPublictonemapoptionsExecute(r)
}

/*
PostEncodingPublictonemapoptions Updates the tone mapping options

Requires authentication as administrator

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ToneMapOptionsServiceAPIPostEncodingPublictonemapoptionsRequest
*/
func (a *ToneMapOptionsServiceAPIService) PostEncodingPublictonemapoptions(ctx context.Context) ToneMapOptionsServiceAPIPostEncodingPublictonemapoptionsRequest {
	return ToneMapOptionsServiceAPIPostEncodingPublictonemapoptionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ToneMapOptionsServiceAPIService) PostEncodingPublictonemapoptionsExecute(r ToneMapOptionsServiceAPIPostEncodingPublictonemapoptionsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ToneMapOptionsServiceAPIService.PostEncodingPublictonemapoptions")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/Encoding/PublicToneMapOptions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/octet-stream"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apikeyauth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarQueryParams.Add("api_key", key)
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
