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
	"strings"
)


// ScheduledTaskServiceAPIService ScheduledTaskServiceAPI service
type ScheduledTaskServiceAPIService service

type ScheduledTaskServiceAPIDeleteScheduledtasksRunningByIdRequest struct {
	ctx context.Context
	ApiService *ScheduledTaskServiceAPIService
	id string
}

func (r ScheduledTaskServiceAPIDeleteScheduledtasksRunningByIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteScheduledtasksRunningByIdExecute(r)
}

/*
DeleteScheduledtasksRunningById Stops a scheduled task

Requires authentication as administrator

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ScheduledTaskServiceAPIDeleteScheduledtasksRunningByIdRequest
*/
func (a *ScheduledTaskServiceAPIService) DeleteScheduledtasksRunningById(ctx context.Context, id string) ScheduledTaskServiceAPIDeleteScheduledtasksRunningByIdRequest {
	return ScheduledTaskServiceAPIDeleteScheduledtasksRunningByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ScheduledTaskServiceAPIService) DeleteScheduledtasksRunningByIdExecute(r ScheduledTaskServiceAPIDeleteScheduledtasksRunningByIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ScheduledTaskServiceAPIService.DeleteScheduledtasksRunningById")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ScheduledTasks/Running/{Id}"
	localVarPath = strings.Replace(localVarPath, "{"+"Id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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

type ScheduledTaskServiceAPIGetScheduledtasksRequest struct {
	ctx context.Context
	ApiService *ScheduledTaskServiceAPIService
	isHidden *bool
	isEnabled *bool
}

// Optional filter tasks that are hidden, or not.
func (r ScheduledTaskServiceAPIGetScheduledtasksRequest) IsHidden(isHidden bool) ScheduledTaskServiceAPIGetScheduledtasksRequest {
	r.isHidden = &isHidden
	return r
}

// Optional filter tasks that are enabled, or not.
func (r ScheduledTaskServiceAPIGetScheduledtasksRequest) IsEnabled(isEnabled bool) ScheduledTaskServiceAPIGetScheduledtasksRequest {
	r.isEnabled = &isEnabled
	return r
}

func (r ScheduledTaskServiceAPIGetScheduledtasksRequest) Execute() ([]ModelTaskInfo, *http.Response, error) {
	return r.ApiService.GetScheduledtasksExecute(r)
}

/*
GetScheduledtasks Gets scheduled tasks

Requires authentication as administrator

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ScheduledTaskServiceAPIGetScheduledtasksRequest
*/
func (a *ScheduledTaskServiceAPIService) GetScheduledtasks(ctx context.Context) ScheduledTaskServiceAPIGetScheduledtasksRequest {
	return ScheduledTaskServiceAPIGetScheduledtasksRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ModelTaskInfo
func (a *ScheduledTaskServiceAPIService) GetScheduledtasksExecute(r ScheduledTaskServiceAPIGetScheduledtasksRequest) ([]ModelTaskInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ModelTaskInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ScheduledTaskServiceAPIService.GetScheduledtasks")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ScheduledTasks"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.isHidden != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "IsHidden", r.isHidden, "")
	}
	if r.isEnabled != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "IsEnabled", r.isEnabled, "")
	}
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

type ScheduledTaskServiceAPIGetScheduledtasksByIdRequest struct {
	ctx context.Context
	ApiService *ScheduledTaskServiceAPIService
	id string
}

func (r ScheduledTaskServiceAPIGetScheduledtasksByIdRequest) Execute() (*ModelTaskInfo, *http.Response, error) {
	return r.ApiService.GetScheduledtasksByIdExecute(r)
}

/*
GetScheduledtasksById Gets a scheduled task, by Id

Requires authentication as administrator

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ScheduledTaskServiceAPIGetScheduledtasksByIdRequest
*/
func (a *ScheduledTaskServiceAPIService) GetScheduledtasksById(ctx context.Context, id string) ScheduledTaskServiceAPIGetScheduledtasksByIdRequest {
	return ScheduledTaskServiceAPIGetScheduledtasksByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ModelTaskInfo
func (a *ScheduledTaskServiceAPIService) GetScheduledtasksByIdExecute(r ScheduledTaskServiceAPIGetScheduledtasksByIdRequest) (*ModelTaskInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ModelTaskInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ScheduledTaskServiceAPIService.GetScheduledtasksById")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ScheduledTasks/{Id}"
	localVarPath = strings.Replace(localVarPath, "{"+"Id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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

type ScheduledTaskServiceAPIPostScheduledtasksByIdTriggersRequest struct {
	ctx context.Context
	ApiService *ScheduledTaskServiceAPIService
	id string
	modelTaskTriggerInfo *[]ModelTaskTriggerInfo
}

// List&#x60;1: 
func (r ScheduledTaskServiceAPIPostScheduledtasksByIdTriggersRequest) ModelTaskTriggerInfo(modelTaskTriggerInfo []ModelTaskTriggerInfo) ScheduledTaskServiceAPIPostScheduledtasksByIdTriggersRequest {
	r.modelTaskTriggerInfo = &modelTaskTriggerInfo
	return r
}

func (r ScheduledTaskServiceAPIPostScheduledtasksByIdTriggersRequest) Execute() (*http.Response, error) {
	return r.ApiService.PostScheduledtasksByIdTriggersExecute(r)
}

/*
PostScheduledtasksByIdTriggers Updates the triggers for a scheduled task

Requires authentication as administrator

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ScheduledTaskServiceAPIPostScheduledtasksByIdTriggersRequest
*/
func (a *ScheduledTaskServiceAPIService) PostScheduledtasksByIdTriggers(ctx context.Context, id string) ScheduledTaskServiceAPIPostScheduledtasksByIdTriggersRequest {
	return ScheduledTaskServiceAPIPostScheduledtasksByIdTriggersRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ScheduledTaskServiceAPIService) PostScheduledtasksByIdTriggersExecute(r ScheduledTaskServiceAPIPostScheduledtasksByIdTriggersRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ScheduledTaskServiceAPIService.PostScheduledtasksByIdTriggers")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ScheduledTasks/{Id}/Triggers"
	localVarPath = strings.Replace(localVarPath, "{"+"Id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.modelTaskTriggerInfo == nil {
		return nil, reportError("modelTaskTriggerInfo is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/xml"}

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
	localVarPostBody = r.modelTaskTriggerInfo
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

type ScheduledTaskServiceAPIPostScheduledtasksRunningByIdRequest struct {
	ctx context.Context
	ApiService *ScheduledTaskServiceAPIService
	id string
}

func (r ScheduledTaskServiceAPIPostScheduledtasksRunningByIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.PostScheduledtasksRunningByIdExecute(r)
}

/*
PostScheduledtasksRunningById Starts a scheduled task

Requires authentication as administrator

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ScheduledTaskServiceAPIPostScheduledtasksRunningByIdRequest
*/
func (a *ScheduledTaskServiceAPIService) PostScheduledtasksRunningById(ctx context.Context, id string) ScheduledTaskServiceAPIPostScheduledtasksRunningByIdRequest {
	return ScheduledTaskServiceAPIPostScheduledtasksRunningByIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ScheduledTaskServiceAPIService) PostScheduledtasksRunningByIdExecute(r ScheduledTaskServiceAPIPostScheduledtasksRunningByIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ScheduledTaskServiceAPIService.PostScheduledtasksRunningById")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ScheduledTasks/Running/{Id}"
	localVarPath = strings.Replace(localVarPath, "{"+"Id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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

type ScheduledTaskServiceAPIPostScheduledtasksRunningByIdDeleteRequest struct {
	ctx context.Context
	ApiService *ScheduledTaskServiceAPIService
	id string
}

func (r ScheduledTaskServiceAPIPostScheduledtasksRunningByIdDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.PostScheduledtasksRunningByIdDeleteExecute(r)
}

/*
PostScheduledtasksRunningByIdDelete Stops a scheduled task

Requires authentication as administrator

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ScheduledTaskServiceAPIPostScheduledtasksRunningByIdDeleteRequest
*/
func (a *ScheduledTaskServiceAPIService) PostScheduledtasksRunningByIdDelete(ctx context.Context, id string) ScheduledTaskServiceAPIPostScheduledtasksRunningByIdDeleteRequest {
	return ScheduledTaskServiceAPIPostScheduledtasksRunningByIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ScheduledTaskServiceAPIService) PostScheduledtasksRunningByIdDeleteExecute(r ScheduledTaskServiceAPIPostScheduledtasksRunningByIdDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ScheduledTaskServiceAPIService.PostScheduledtasksRunningByIdDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ScheduledTasks/Running/{Id}/Delete"
	localVarPath = strings.Replace(localVarPath, "{"+"Id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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
