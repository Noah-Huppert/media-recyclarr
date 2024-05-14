/*
Overseerr API

This is the documentation for the Overseerr API backend.  Two primary authentication methods are supported:  - **Cookie Authentication**: A valid sign-in to the `/auth/plex` or `/auth/local` will generate a valid authentication cookie. - **API Key Authentication**: Sign-in is also possible by passing an `X-Api-Key` header along with a valid API Key generated by Overseerr. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package jellyseerrclient

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// MediaAPIService MediaAPI service
type MediaAPIService service

type MediaAPIMediaGetRequest struct {
	ctx context.Context
	ApiService *MediaAPIService
	take *float32
	skip *float32
	filter *string
	sort *string
}

func (r MediaAPIMediaGetRequest) Take(take float32) MediaAPIMediaGetRequest {
	r.take = &take
	return r
}

func (r MediaAPIMediaGetRequest) Skip(skip float32) MediaAPIMediaGetRequest {
	r.skip = &skip
	return r
}

func (r MediaAPIMediaGetRequest) Filter(filter string) MediaAPIMediaGetRequest {
	r.filter = &filter
	return r
}

func (r MediaAPIMediaGetRequest) Sort(sort string) MediaAPIMediaGetRequest {
	r.sort = &sort
	return r
}

func (r MediaAPIMediaGetRequest) Execute() (*MediaGet200Response, *http.Response, error) {
	return r.ApiService.MediaGetExecute(r)
}

/*
MediaGet Get media

Returns all media (can be filtered and limited) in a JSON object.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return MediaAPIMediaGetRequest
*/
func (a *MediaAPIService) MediaGet(ctx context.Context) MediaAPIMediaGetRequest {
	return MediaAPIMediaGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MediaGet200Response
func (a *MediaAPIService) MediaGetExecute(r MediaAPIMediaGetRequest) (*MediaGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MediaGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MediaAPIService.MediaGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/media"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.take != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "take", r.take, "")
	}
	if r.skip != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skip", r.skip, "")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "")
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "")
	} else {
		var defaultValue string = "added"
		r.sort = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
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

type MediaAPIMediaMediaIdDeleteRequest struct {
	ctx context.Context
	ApiService *MediaAPIService
	mediaId string
}

func (r MediaAPIMediaMediaIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.MediaMediaIdDeleteExecute(r)
}

/*
MediaMediaIdDelete Delete media item

Removes a media item. The `MANAGE_REQUESTS` permission is required to perform this action.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param mediaId Media ID
 @return MediaAPIMediaMediaIdDeleteRequest
*/
func (a *MediaAPIService) MediaMediaIdDelete(ctx context.Context, mediaId string) MediaAPIMediaMediaIdDeleteRequest {
	return MediaAPIMediaMediaIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		mediaId: mediaId,
	}
}

// Execute executes the request
func (a *MediaAPIService) MediaMediaIdDeleteExecute(r MediaAPIMediaMediaIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MediaAPIService.MediaMediaIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/media/{mediaId}"
	localVarPath = strings.Replace(localVarPath, "{"+"mediaId"+"}", url.PathEscape(parameterValueToString(r.mediaId, "mediaId")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
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

type MediaAPIMediaMediaIdFileDeleteRequest struct {
	ctx context.Context
	ApiService *MediaAPIService
	mediaId string
}

func (r MediaAPIMediaMediaIdFileDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.MediaMediaIdFileDeleteExecute(r)
}

/*
MediaMediaIdFileDelete Delete media file

Removes a media file from radarr/sonarr. The `ADMIN` permission is required to perform this action.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param mediaId Media ID
 @return MediaAPIMediaMediaIdFileDeleteRequest
*/
func (a *MediaAPIService) MediaMediaIdFileDelete(ctx context.Context, mediaId string) MediaAPIMediaMediaIdFileDeleteRequest {
	return MediaAPIMediaMediaIdFileDeleteRequest{
		ApiService: a,
		ctx: ctx,
		mediaId: mediaId,
	}
}

// Execute executes the request
func (a *MediaAPIService) MediaMediaIdFileDeleteExecute(r MediaAPIMediaMediaIdFileDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MediaAPIService.MediaMediaIdFileDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/media/{mediaId}/file"
	localVarPath = strings.Replace(localVarPath, "{"+"mediaId"+"}", url.PathEscape(parameterValueToString(r.mediaId, "mediaId")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
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

type MediaAPIMediaMediaIdStatusPostRequest struct {
	ctx context.Context
	ApiService *MediaAPIService
	mediaId string
	status string
	mediaMediaIdStatusPostRequest *MediaMediaIdStatusPostRequest
}

func (r MediaAPIMediaMediaIdStatusPostRequest) MediaMediaIdStatusPostRequest(mediaMediaIdStatusPostRequest MediaMediaIdStatusPostRequest) MediaAPIMediaMediaIdStatusPostRequest {
	r.mediaMediaIdStatusPostRequest = &mediaMediaIdStatusPostRequest
	return r
}

func (r MediaAPIMediaMediaIdStatusPostRequest) Execute() (*MediaInfo, *http.Response, error) {
	return r.ApiService.MediaMediaIdStatusPostExecute(r)
}

/*
MediaMediaIdStatusPost Update media status

Updates a media item's status and returns the media in JSON format

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param mediaId Media ID
 @param status New status
 @return MediaAPIMediaMediaIdStatusPostRequest
*/
func (a *MediaAPIService) MediaMediaIdStatusPost(ctx context.Context, mediaId string, status string) MediaAPIMediaMediaIdStatusPostRequest {
	return MediaAPIMediaMediaIdStatusPostRequest{
		ApiService: a,
		ctx: ctx,
		mediaId: mediaId,
		status: status,
	}
}

// Execute executes the request
//  @return MediaInfo
func (a *MediaAPIService) MediaMediaIdStatusPostExecute(r MediaAPIMediaMediaIdStatusPostRequest) (*MediaInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MediaInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MediaAPIService.MediaMediaIdStatusPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/media/{mediaId}/{status}"
	localVarPath = strings.Replace(localVarPath, "{"+"mediaId"+"}", url.PathEscape(parameterValueToString(r.mediaId, "mediaId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"status"+"}", url.PathEscape(parameterValueToString(r.status, "status")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.mediaMediaIdStatusPostRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
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

type MediaAPIMediaMediaIdWatchDataGetRequest struct {
	ctx context.Context
	ApiService *MediaAPIService
	mediaId string
}

func (r MediaAPIMediaMediaIdWatchDataGetRequest) Execute() (*MediaMediaIdWatchDataGet200Response, *http.Response, error) {
	return r.ApiService.MediaMediaIdWatchDataGetExecute(r)
}

/*
MediaMediaIdWatchDataGet Get watch data

Returns play count, play duration, and users who have watched the media.

Requires the `ADMIN` permission.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param mediaId Media ID
 @return MediaAPIMediaMediaIdWatchDataGetRequest
*/
func (a *MediaAPIService) MediaMediaIdWatchDataGet(ctx context.Context, mediaId string) MediaAPIMediaMediaIdWatchDataGetRequest {
	return MediaAPIMediaMediaIdWatchDataGetRequest{
		ApiService: a,
		ctx: ctx,
		mediaId: mediaId,
	}
}

// Execute executes the request
//  @return MediaMediaIdWatchDataGet200Response
func (a *MediaAPIService) MediaMediaIdWatchDataGetExecute(r MediaAPIMediaMediaIdWatchDataGetRequest) (*MediaMediaIdWatchDataGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MediaMediaIdWatchDataGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MediaAPIService.MediaMediaIdWatchDataGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/media/{mediaId}/watch_data"
	localVarPath = strings.Replace(localVarPath, "{"+"mediaId"+"}", url.PathEscape(parameterValueToString(r.mediaId, "mediaId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Api-Key"] = key
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
