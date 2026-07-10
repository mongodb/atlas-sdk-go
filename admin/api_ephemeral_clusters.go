// Code based on the AtlasAPI V2 OpenAPI file
package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
)

type EphemeralClustersAPI interface {

	/*
		CreateEphemeralCluster Create an Ephemeral Atlas Cluster

		Creates an ephemeral Atlas cluster for users or agents.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param createEphemeralClusterRequest Ephemeral cluster to create.
		@return CreateEphemeralClusterApiRequest
	*/
	CreateEphemeralCluster(ctx context.Context, createEphemeralClusterRequest *CreateEphemeralClusterRequest) CreateEphemeralClusterApiRequest
	/*
		CreateEphemeralCluster Create an Ephemeral Atlas Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateEphemeralClusterApiParams - Parameters for the request
		@return CreateEphemeralClusterApiRequest
	*/
	CreateEphemeralClusterWithParams(ctx context.Context, args *CreateEphemeralClusterApiParams) CreateEphemeralClusterApiRequest

	// Method available only for mocking purposes
	CreateEphemeralClusterExecute(r CreateEphemeralClusterApiRequest) (*EphemeralClusterCreated, *http.Response, error)
}

// EphemeralClustersAPIService EphemeralClustersAPI service
type EphemeralClustersAPIService service

type CreateEphemeralClusterApiRequest struct {
	ctx                           context.Context
	ApiService                    EphemeralClustersAPI
	createEphemeralClusterRequest *CreateEphemeralClusterRequest
}

type CreateEphemeralClusterApiParams struct {
	CreateEphemeralClusterRequest *CreateEphemeralClusterRequest
}

func (a *EphemeralClustersAPIService) CreateEphemeralClusterWithParams(ctx context.Context, args *CreateEphemeralClusterApiParams) CreateEphemeralClusterApiRequest {
	return CreateEphemeralClusterApiRequest{
		ApiService:                    a,
		ctx:                           ctx,
		createEphemeralClusterRequest: args.CreateEphemeralClusterRequest,
	}
}

func (r CreateEphemeralClusterApiRequest) Execute() (*EphemeralClusterCreated, *http.Response, error) {
	return r.ApiService.CreateEphemeralClusterExecute(r)
}

/*
CreateEphemeralCluster Create an Ephemeral Atlas Cluster

Creates an ephemeral Atlas cluster for users or agents.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CreateEphemeralClusterApiRequest
*/
func (a *EphemeralClustersAPIService) CreateEphemeralCluster(ctx context.Context, createEphemeralClusterRequest *CreateEphemeralClusterRequest) CreateEphemeralClusterApiRequest {
	return CreateEphemeralClusterApiRequest{
		ApiService:                    a,
		ctx:                           ctx,
		createEphemeralClusterRequest: createEphemeralClusterRequest,
	}
}

// CreateEphemeralClusterExecute executes the request
//
//	@return EphemeralClusterCreated
func (a *EphemeralClustersAPIService) CreateEphemeralClusterExecute(r CreateEphemeralClusterApiRequest) (*EphemeralClusterCreated, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *EphemeralClusterCreated
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EphemeralClustersAPIService.CreateEphemeralCluster")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/unauth/ephemeralClusters:create"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createEphemeralClusterRequest == nil {
		return localVarReturnValue, nil, reportError("createEphemeralClusterRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2025-03-12+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2025-03-12+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.createEphemeralClusterRequest
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
