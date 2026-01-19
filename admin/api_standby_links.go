// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type StandbyLinksApi interface {

	/*
		CreateGroupStandbyLink Create Standby Link

		Creates a disaster recovery standby link between an active cluster and a standby cluster. Both clusters must not already be part of a standby link system, must be single-region clusters, and must be in different regions.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies the project containing the clusters.
		@param apiStandbyLinkRequest Request body containing the active and standby cluster identifiers.
		@return CreateGroupStandbyLinkApiRequest
	*/
	CreateGroupStandbyLink(ctx context.Context, groupId string, apiStandbyLinkRequest *ApiStandbyLinkRequest) CreateGroupStandbyLinkApiRequest
	/*
		CreateGroupStandbyLink Create Standby Link


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateGroupStandbyLinkApiParams - Parameters for the request
		@return CreateGroupStandbyLinkApiRequest
	*/
	CreateGroupStandbyLinkWithParams(ctx context.Context, args *CreateGroupStandbyLinkApiParams) CreateGroupStandbyLinkApiRequest

	// Method available only for mocking purposes
	CreateGroupStandbyLinkExecute(r CreateGroupStandbyLinkApiRequest) (*ApiStandbyLinkResponse, *http.Response, error)

	/*
		CreateStandbyLinkFailover Create Standby Link Failover

		Creates a new failover operation to promote a standby cluster or reverse cluster roles. The standby link must be in ACTIVE status to initiate a failover.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies the project.
		@param standbyLinkId Unique 24-hexadecimal digit string that identifies the standby link.
		@param apiStandbyLinkFailoverRequest Request body containing the failover type (PLANNED or UNPLANNED).
		@return CreateStandbyLinkFailoverApiRequest
	*/
	CreateStandbyLinkFailover(ctx context.Context, groupId string, standbyLinkId string, apiStandbyLinkFailoverRequest *ApiStandbyLinkFailoverRequest) CreateStandbyLinkFailoverApiRequest
	/*
		CreateStandbyLinkFailover Create Standby Link Failover


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateStandbyLinkFailoverApiParams - Parameters for the request
		@return CreateStandbyLinkFailoverApiRequest
	*/
	CreateStandbyLinkFailoverWithParams(ctx context.Context, args *CreateStandbyLinkFailoverApiParams) CreateStandbyLinkFailoverApiRequest

	// Method available only for mocking purposes
	CreateStandbyLinkFailoverExecute(r CreateStandbyLinkFailoverApiRequest) (*ApiStandbyLinkFailoverResponse, *http.Response, error)

	/*
		DeleteGroupStandbyLink Delete Standby Link

		Terminates an existing standby link and converts the active and standby clusters back to regular clusters.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies the project.
		@param standbyLinkId Unique 24-hexadecimal digit string that identifies the standby link.
		@return DeleteGroupStandbyLinkApiRequest
	*/
	DeleteGroupStandbyLink(ctx context.Context, groupId string, standbyLinkId string) DeleteGroupStandbyLinkApiRequest
	/*
		DeleteGroupStandbyLink Delete Standby Link


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteGroupStandbyLinkApiParams - Parameters for the request
		@return DeleteGroupStandbyLinkApiRequest
	*/
	DeleteGroupStandbyLinkWithParams(ctx context.Context, args *DeleteGroupStandbyLinkApiParams) DeleteGroupStandbyLinkApiRequest

	// Method available only for mocking purposes
	DeleteGroupStandbyLinkExecute(r DeleteGroupStandbyLinkApiRequest) (*http.Response, error)

	/*
		GetGroupStandbyLink Return One Standby Link

		Returns detailed information about a specific standby link including status, group ID, and cluster unique IDs.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies the project.
		@param standbyLinkId Unique 24-hexadecimal digit string that identifies the standby link.
		@return GetGroupStandbyLinkApiRequest
	*/
	GetGroupStandbyLink(ctx context.Context, groupId string, standbyLinkId string) GetGroupStandbyLinkApiRequest
	/*
		GetGroupStandbyLink Return One Standby Link


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetGroupStandbyLinkApiParams - Parameters for the request
		@return GetGroupStandbyLinkApiRequest
	*/
	GetGroupStandbyLinkWithParams(ctx context.Context, args *GetGroupStandbyLinkApiParams) GetGroupStandbyLinkApiRequest

	// Method available only for mocking purposes
	GetGroupStandbyLinkExecute(r GetGroupStandbyLinkApiRequest) (*ApiStandbyLinkResponse, *http.Response, error)

	/*
		GetStandbyLinkFailover Return One Standby Link Failover

		Returns the current status and details of a specific failover operation for a standby link.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies the project.
		@param standbyLinkId Unique 24-hexadecimal digit string that identifies the standby link.
		@param failoverId Unique 24-hexadecimal digit string that identifies the failover operation.
		@return GetStandbyLinkFailoverApiRequest
	*/
	GetStandbyLinkFailover(ctx context.Context, groupId string, standbyLinkId string, failoverId string) GetStandbyLinkFailoverApiRequest
	/*
		GetStandbyLinkFailover Return One Standby Link Failover


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetStandbyLinkFailoverApiParams - Parameters for the request
		@return GetStandbyLinkFailoverApiRequest
	*/
	GetStandbyLinkFailoverWithParams(ctx context.Context, args *GetStandbyLinkFailoverApiParams) GetStandbyLinkFailoverApiRequest

	// Method available only for mocking purposes
	GetStandbyLinkFailoverExecute(r GetStandbyLinkFailoverApiRequest) (*ApiStandbyLinkFailoverResponse, *http.Response, error)

	/*
		ListGroupStandbyLinks Return All Standby Links for One Project

		Returns all standby links for the specified project. Each standby link represents a disaster recovery relationship between an active cluster and a standby cluster.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies the project.
		@return ListGroupStandbyLinksApiRequest
	*/
	ListGroupStandbyLinks(ctx context.Context, groupId string) ListGroupStandbyLinksApiRequest
	/*
		ListGroupStandbyLinks Return All Standby Links for One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListGroupStandbyLinksApiParams - Parameters for the request
		@return ListGroupStandbyLinksApiRequest
	*/
	ListGroupStandbyLinksWithParams(ctx context.Context, args *ListGroupStandbyLinksApiParams) ListGroupStandbyLinksApiRequest

	// Method available only for mocking purposes
	ListGroupStandbyLinksExecute(r ListGroupStandbyLinksApiRequest) (*PaginatedApiStandbyLinkResponse, *http.Response, error)

	/*
		ListStandbyLinkFailovers Return All Standby Link Failovers

		Returns all failover operations for the specified standby link. Each failover represents a planned or unplanned process to promote a standby cluster to an active cluster.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies the project.
		@param standbyLinkId Unique 24-hexadecimal digit string that identifies the standby link.
		@return ListStandbyLinkFailoversApiRequest
	*/
	ListStandbyLinkFailovers(ctx context.Context, groupId string, standbyLinkId string) ListStandbyLinkFailoversApiRequest
	/*
		ListStandbyLinkFailovers Return All Standby Link Failovers


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListStandbyLinkFailoversApiParams - Parameters for the request
		@return ListStandbyLinkFailoversApiRequest
	*/
	ListStandbyLinkFailoversWithParams(ctx context.Context, args *ListStandbyLinkFailoversApiParams) ListStandbyLinkFailoversApiRequest

	// Method available only for mocking purposes
	ListStandbyLinkFailoversExecute(r ListStandbyLinkFailoversApiRequest) (*PaginatedApiStandbyLinkFailoverResponse, *http.Response, error)
}

// StandbyLinksApiService StandbyLinksApi service
type StandbyLinksApiService service

type CreateGroupStandbyLinkApiRequest struct {
	ctx                   context.Context
	ApiService            StandbyLinksApi
	groupId               string
	apiStandbyLinkRequest *ApiStandbyLinkRequest
	envelope              *bool
}

type CreateGroupStandbyLinkApiParams struct {
	GroupId               string
	ApiStandbyLinkRequest *ApiStandbyLinkRequest
	Envelope              *bool
}

func (a *StandbyLinksApiService) CreateGroupStandbyLinkWithParams(ctx context.Context, args *CreateGroupStandbyLinkApiParams) CreateGroupStandbyLinkApiRequest {
	return CreateGroupStandbyLinkApiRequest{
		ApiService:            a,
		ctx:                   ctx,
		groupId:               args.GroupId,
		apiStandbyLinkRequest: args.ApiStandbyLinkRequest,
		envelope:              args.Envelope,
	}
}

// Flag that indicates whether to wrap the response in an envelope.
func (r CreateGroupStandbyLinkApiRequest) Envelope(envelope bool) CreateGroupStandbyLinkApiRequest {
	r.envelope = &envelope
	return r
}

func (r CreateGroupStandbyLinkApiRequest) Execute() (*ApiStandbyLinkResponse, *http.Response, error) {
	return r.ApiService.CreateGroupStandbyLinkExecute(r)
}

/*
CreateGroupStandbyLink Create Standby Link

Creates a disaster recovery standby link between an active cluster and a standby cluster. Both clusters must not already be part of a standby link system, must be single-region clusters, and must be in different regions.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies the project containing the clusters.
	@return CreateGroupStandbyLinkApiRequest
*/
func (a *StandbyLinksApiService) CreateGroupStandbyLink(ctx context.Context, groupId string, apiStandbyLinkRequest *ApiStandbyLinkRequest) CreateGroupStandbyLinkApiRequest {
	return CreateGroupStandbyLinkApiRequest{
		ApiService:            a,
		ctx:                   ctx,
		groupId:               groupId,
		apiStandbyLinkRequest: apiStandbyLinkRequest,
	}
}

// CreateGroupStandbyLinkExecute executes the request
//
//	@return ApiStandbyLinkResponse
func (a *StandbyLinksApiService) CreateGroupStandbyLinkExecute(r CreateGroupStandbyLinkApiRequest) (*ApiStandbyLinkResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ApiStandbyLinkResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StandbyLinksApiService.CreateGroupStandbyLink")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/standbyLinks"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiStandbyLinkRequest == nil {
		return localVarReturnValue, nil, reportError("apiStandbyLinkRequest is required and must be specified")
	}

	if r.envelope != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "envelope", r.envelope, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.preview+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.preview+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.apiStandbyLinkRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type CreateStandbyLinkFailoverApiRequest struct {
	ctx                           context.Context
	ApiService                    StandbyLinksApi
	groupId                       string
	standbyLinkId                 string
	apiStandbyLinkFailoverRequest *ApiStandbyLinkFailoverRequest
	envelope                      *bool
}

type CreateStandbyLinkFailoverApiParams struct {
	GroupId                       string
	StandbyLinkId                 string
	ApiStandbyLinkFailoverRequest *ApiStandbyLinkFailoverRequest
	Envelope                      *bool
}

func (a *StandbyLinksApiService) CreateStandbyLinkFailoverWithParams(ctx context.Context, args *CreateStandbyLinkFailoverApiParams) CreateStandbyLinkFailoverApiRequest {
	return CreateStandbyLinkFailoverApiRequest{
		ApiService:                    a,
		ctx:                           ctx,
		groupId:                       args.GroupId,
		standbyLinkId:                 args.StandbyLinkId,
		apiStandbyLinkFailoverRequest: args.ApiStandbyLinkFailoverRequest,
		envelope:                      args.Envelope,
	}
}

// Flag that indicates whether to wrap the response in an envelope.
func (r CreateStandbyLinkFailoverApiRequest) Envelope(envelope bool) CreateStandbyLinkFailoverApiRequest {
	r.envelope = &envelope
	return r
}

func (r CreateStandbyLinkFailoverApiRequest) Execute() (*ApiStandbyLinkFailoverResponse, *http.Response, error) {
	return r.ApiService.CreateStandbyLinkFailoverExecute(r)
}

/*
CreateStandbyLinkFailover Create Standby Link Failover

Creates a new failover operation to promote a standby cluster or reverse cluster roles. The standby link must be in ACTIVE status to initiate a failover.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies the project.
	@param standbyLinkId Unique 24-hexadecimal digit string that identifies the standby link.
	@return CreateStandbyLinkFailoverApiRequest
*/
func (a *StandbyLinksApiService) CreateStandbyLinkFailover(ctx context.Context, groupId string, standbyLinkId string, apiStandbyLinkFailoverRequest *ApiStandbyLinkFailoverRequest) CreateStandbyLinkFailoverApiRequest {
	return CreateStandbyLinkFailoverApiRequest{
		ApiService:                    a,
		ctx:                           ctx,
		groupId:                       groupId,
		standbyLinkId:                 standbyLinkId,
		apiStandbyLinkFailoverRequest: apiStandbyLinkFailoverRequest,
	}
}

// CreateStandbyLinkFailoverExecute executes the request
//
//	@return ApiStandbyLinkFailoverResponse
func (a *StandbyLinksApiService) CreateStandbyLinkFailoverExecute(r CreateStandbyLinkFailoverApiRequest) (*ApiStandbyLinkFailoverResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ApiStandbyLinkFailoverResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StandbyLinksApiService.CreateStandbyLinkFailover")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/standbyLinks/{standbyLinkId}/failovers"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.standbyLinkId == "" {
		return localVarReturnValue, nil, reportError("standbyLinkId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"standbyLinkId"+"}", url.PathEscape(r.standbyLinkId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiStandbyLinkFailoverRequest == nil {
		return localVarReturnValue, nil, reportError("apiStandbyLinkFailoverRequest is required and must be specified")
	}

	if r.envelope != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "envelope", r.envelope, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.preview+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.preview+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.apiStandbyLinkFailoverRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type DeleteGroupStandbyLinkApiRequest struct {
	ctx           context.Context
	ApiService    StandbyLinksApi
	groupId       string
	standbyLinkId string
	envelope      *bool
}

type DeleteGroupStandbyLinkApiParams struct {
	GroupId       string
	StandbyLinkId string
	Envelope      *bool
}

func (a *StandbyLinksApiService) DeleteGroupStandbyLinkWithParams(ctx context.Context, args *DeleteGroupStandbyLinkApiParams) DeleteGroupStandbyLinkApiRequest {
	return DeleteGroupStandbyLinkApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       args.GroupId,
		standbyLinkId: args.StandbyLinkId,
		envelope:      args.Envelope,
	}
}

// Flag that indicates whether to wrap the response in an envelope.
func (r DeleteGroupStandbyLinkApiRequest) Envelope(envelope bool) DeleteGroupStandbyLinkApiRequest {
	r.envelope = &envelope
	return r
}

func (r DeleteGroupStandbyLinkApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteGroupStandbyLinkExecute(r)
}

/*
DeleteGroupStandbyLink Delete Standby Link

Terminates an existing standby link and converts the active and standby clusters back to regular clusters.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies the project.
	@param standbyLinkId Unique 24-hexadecimal digit string that identifies the standby link.
	@return DeleteGroupStandbyLinkApiRequest
*/
func (a *StandbyLinksApiService) DeleteGroupStandbyLink(ctx context.Context, groupId string, standbyLinkId string) DeleteGroupStandbyLinkApiRequest {
	return DeleteGroupStandbyLinkApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       groupId,
		standbyLinkId: standbyLinkId,
	}
}

// DeleteGroupStandbyLinkExecute executes the request
func (a *StandbyLinksApiService) DeleteGroupStandbyLinkExecute(r DeleteGroupStandbyLinkApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StandbyLinksApiService.DeleteGroupStandbyLink")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/standbyLinks/{standbyLinkId}"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.standbyLinkId == "" {
		return nil, reportError("standbyLinkId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"standbyLinkId"+"}", url.PathEscape(r.standbyLinkId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.envelope != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "envelope", r.envelope, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.preview+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type GetGroupStandbyLinkApiRequest struct {
	ctx           context.Context
	ApiService    StandbyLinksApi
	groupId       string
	standbyLinkId string
	envelope      *bool
}

type GetGroupStandbyLinkApiParams struct {
	GroupId       string
	StandbyLinkId string
	Envelope      *bool
}

func (a *StandbyLinksApiService) GetGroupStandbyLinkWithParams(ctx context.Context, args *GetGroupStandbyLinkApiParams) GetGroupStandbyLinkApiRequest {
	return GetGroupStandbyLinkApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       args.GroupId,
		standbyLinkId: args.StandbyLinkId,
		envelope:      args.Envelope,
	}
}

// Flag that indicates whether to wrap the response in an envelope.
func (r GetGroupStandbyLinkApiRequest) Envelope(envelope bool) GetGroupStandbyLinkApiRequest {
	r.envelope = &envelope
	return r
}

func (r GetGroupStandbyLinkApiRequest) Execute() (*ApiStandbyLinkResponse, *http.Response, error) {
	return r.ApiService.GetGroupStandbyLinkExecute(r)
}

/*
GetGroupStandbyLink Return One Standby Link

Returns detailed information about a specific standby link including status, group ID, and cluster unique IDs.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies the project.
	@param standbyLinkId Unique 24-hexadecimal digit string that identifies the standby link.
	@return GetGroupStandbyLinkApiRequest
*/
func (a *StandbyLinksApiService) GetGroupStandbyLink(ctx context.Context, groupId string, standbyLinkId string) GetGroupStandbyLinkApiRequest {
	return GetGroupStandbyLinkApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       groupId,
		standbyLinkId: standbyLinkId,
	}
}

// GetGroupStandbyLinkExecute executes the request
//
//	@return ApiStandbyLinkResponse
func (a *StandbyLinksApiService) GetGroupStandbyLinkExecute(r GetGroupStandbyLinkApiRequest) (*ApiStandbyLinkResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ApiStandbyLinkResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StandbyLinksApiService.GetGroupStandbyLink")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/standbyLinks/{standbyLinkId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.standbyLinkId == "" {
		return localVarReturnValue, nil, reportError("standbyLinkId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"standbyLinkId"+"}", url.PathEscape(r.standbyLinkId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.envelope != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "envelope", r.envelope, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.preview+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetStandbyLinkFailoverApiRequest struct {
	ctx           context.Context
	ApiService    StandbyLinksApi
	groupId       string
	standbyLinkId string
	failoverId    string
	envelope      *bool
}

type GetStandbyLinkFailoverApiParams struct {
	GroupId       string
	StandbyLinkId string
	FailoverId    string
	Envelope      *bool
}

func (a *StandbyLinksApiService) GetStandbyLinkFailoverWithParams(ctx context.Context, args *GetStandbyLinkFailoverApiParams) GetStandbyLinkFailoverApiRequest {
	return GetStandbyLinkFailoverApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       args.GroupId,
		standbyLinkId: args.StandbyLinkId,
		failoverId:    args.FailoverId,
		envelope:      args.Envelope,
	}
}

// Flag that indicates whether to wrap the response in an envelope.
func (r GetStandbyLinkFailoverApiRequest) Envelope(envelope bool) GetStandbyLinkFailoverApiRequest {
	r.envelope = &envelope
	return r
}

func (r GetStandbyLinkFailoverApiRequest) Execute() (*ApiStandbyLinkFailoverResponse, *http.Response, error) {
	return r.ApiService.GetStandbyLinkFailoverExecute(r)
}

/*
GetStandbyLinkFailover Return One Standby Link Failover

Returns the current status and details of a specific failover operation for a standby link.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies the project.
	@param standbyLinkId Unique 24-hexadecimal digit string that identifies the standby link.
	@param failoverId Unique 24-hexadecimal digit string that identifies the failover operation.
	@return GetStandbyLinkFailoverApiRequest
*/
func (a *StandbyLinksApiService) GetStandbyLinkFailover(ctx context.Context, groupId string, standbyLinkId string, failoverId string) GetStandbyLinkFailoverApiRequest {
	return GetStandbyLinkFailoverApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       groupId,
		standbyLinkId: standbyLinkId,
		failoverId:    failoverId,
	}
}

// GetStandbyLinkFailoverExecute executes the request
//
//	@return ApiStandbyLinkFailoverResponse
func (a *StandbyLinksApiService) GetStandbyLinkFailoverExecute(r GetStandbyLinkFailoverApiRequest) (*ApiStandbyLinkFailoverResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ApiStandbyLinkFailoverResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StandbyLinksApiService.GetStandbyLinkFailover")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/standbyLinks/{standbyLinkId}/failovers/{failoverId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.standbyLinkId == "" {
		return localVarReturnValue, nil, reportError("standbyLinkId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"standbyLinkId"+"}", url.PathEscape(r.standbyLinkId), -1)
	if r.failoverId == "" {
		return localVarReturnValue, nil, reportError("failoverId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"failoverId"+"}", url.PathEscape(r.failoverId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.envelope != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "envelope", r.envelope, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.preview+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ListGroupStandbyLinksApiRequest struct {
	ctx          context.Context
	ApiService   StandbyLinksApi
	groupId      string
	itemsPerPage *int
	pageNum      *int
}

type ListGroupStandbyLinksApiParams struct {
	GroupId      string
	ItemsPerPage *int
	PageNum      *int
}

func (a *StandbyLinksApiService) ListGroupStandbyLinksWithParams(ctx context.Context, args *ListGroupStandbyLinksApiParams) ListGroupStandbyLinksApiRequest {
	return ListGroupStandbyLinksApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Number of items that the response returns per page.
func (r ListGroupStandbyLinksApiRequest) ItemsPerPage(itemsPerPage int) ListGroupStandbyLinksApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListGroupStandbyLinksApiRequest) PageNum(pageNum int) ListGroupStandbyLinksApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListGroupStandbyLinksApiRequest) Execute() (*PaginatedApiStandbyLinkResponse, *http.Response, error) {
	return r.ApiService.ListGroupStandbyLinksExecute(r)
}

/*
ListGroupStandbyLinks Return All Standby Links for One Project

Returns all standby links for the specified project. Each standby link represents a disaster recovery relationship between an active cluster and a standby cluster.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies the project.
	@return ListGroupStandbyLinksApiRequest
*/
func (a *StandbyLinksApiService) ListGroupStandbyLinks(ctx context.Context, groupId string) ListGroupStandbyLinksApiRequest {
	return ListGroupStandbyLinksApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// ListGroupStandbyLinksExecute executes the request
//
//	@return PaginatedApiStandbyLinkResponse
func (a *StandbyLinksApiService) ListGroupStandbyLinksExecute(r ListGroupStandbyLinksApiRequest) (*PaginatedApiStandbyLinkResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedApiStandbyLinkResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StandbyLinksApiService.ListGroupStandbyLinks")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/standbyLinks"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.itemsPerPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	} else {
		var defaultValue int = 100
		r.itemsPerPage = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	}
	if r.pageNum != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	} else {
		var defaultValue int = 1
		r.pageNum = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.preview+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ListStandbyLinkFailoversApiRequest struct {
	ctx           context.Context
	ApiService    StandbyLinksApi
	groupId       string
	standbyLinkId string
	itemsPerPage  *int
	pageNum       *int
}

type ListStandbyLinkFailoversApiParams struct {
	GroupId       string
	StandbyLinkId string
	ItemsPerPage  *int
	PageNum       *int
}

func (a *StandbyLinksApiService) ListStandbyLinkFailoversWithParams(ctx context.Context, args *ListStandbyLinkFailoversApiParams) ListStandbyLinkFailoversApiRequest {
	return ListStandbyLinkFailoversApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       args.GroupId,
		standbyLinkId: args.StandbyLinkId,
		itemsPerPage:  args.ItemsPerPage,
		pageNum:       args.PageNum,
	}
}

// Number of items that the response returns per page.
func (r ListStandbyLinkFailoversApiRequest) ItemsPerPage(itemsPerPage int) ListStandbyLinkFailoversApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListStandbyLinkFailoversApiRequest) PageNum(pageNum int) ListStandbyLinkFailoversApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListStandbyLinkFailoversApiRequest) Execute() (*PaginatedApiStandbyLinkFailoverResponse, *http.Response, error) {
	return r.ApiService.ListStandbyLinkFailoversExecute(r)
}

/*
ListStandbyLinkFailovers Return All Standby Link Failovers

Returns all failover operations for the specified standby link. Each failover represents a planned or unplanned process to promote a standby cluster to an active cluster.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies the project.
	@param standbyLinkId Unique 24-hexadecimal digit string that identifies the standby link.
	@return ListStandbyLinkFailoversApiRequest
*/
func (a *StandbyLinksApiService) ListStandbyLinkFailovers(ctx context.Context, groupId string, standbyLinkId string) ListStandbyLinkFailoversApiRequest {
	return ListStandbyLinkFailoversApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       groupId,
		standbyLinkId: standbyLinkId,
	}
}

// ListStandbyLinkFailoversExecute executes the request
//
//	@return PaginatedApiStandbyLinkFailoverResponse
func (a *StandbyLinksApiService) ListStandbyLinkFailoversExecute(r ListStandbyLinkFailoversApiRequest) (*PaginatedApiStandbyLinkFailoverResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedApiStandbyLinkFailoverResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StandbyLinksApiService.ListStandbyLinkFailovers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/standbyLinks/{standbyLinkId}/failovers"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.standbyLinkId == "" {
		return localVarReturnValue, nil, reportError("standbyLinkId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"standbyLinkId"+"}", url.PathEscape(r.standbyLinkId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.itemsPerPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	} else {
		var defaultValue int = 100
		r.itemsPerPage = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	}
	if r.pageNum != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	} else {
		var defaultValue int = 1
		r.pageNum = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.preview+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
