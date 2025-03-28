// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type ResourcePoliciesApi interface {

	/*
		CreateAtlasResourcePolicy Create one Atlas Resource Policy

		Create one Atlas Resource Policy for an org.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param apiAtlasResourcePolicyCreate Atlas Resource Policy to create.
		@return CreateAtlasResourcePolicyApiRequest
	*/
	CreateAtlasResourcePolicy(ctx context.Context, orgId string, apiAtlasResourcePolicyCreate *ApiAtlasResourcePolicyCreate) CreateAtlasResourcePolicyApiRequest
	/*
		CreateAtlasResourcePolicy Create one Atlas Resource Policy


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateAtlasResourcePolicyApiParams - Parameters for the request
		@return CreateAtlasResourcePolicyApiRequest
	*/
	CreateAtlasResourcePolicyWithParams(ctx context.Context, args *CreateAtlasResourcePolicyApiParams) CreateAtlasResourcePolicyApiRequest

	// Method available only for mocking purposes
	CreateAtlasResourcePolicyExecute(r CreateAtlasResourcePolicyApiRequest) (*ApiAtlasResourcePolicy, *http.Response, error)

	/*
		DeleteAtlasResourcePolicy Delete one Atlas Resource Policy

		Delete one Atlas Resource Policy for an org.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param resourcePolicyId Unique 24-hexadecimal digit string that identifies an atlas resource policy.
		@return DeleteAtlasResourcePolicyApiRequest
	*/
	DeleteAtlasResourcePolicy(ctx context.Context, orgId string, resourcePolicyId string) DeleteAtlasResourcePolicyApiRequest
	/*
		DeleteAtlasResourcePolicy Delete one Atlas Resource Policy


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteAtlasResourcePolicyApiParams - Parameters for the request
		@return DeleteAtlasResourcePolicyApiRequest
	*/
	DeleteAtlasResourcePolicyWithParams(ctx context.Context, args *DeleteAtlasResourcePolicyApiParams) DeleteAtlasResourcePolicyApiRequest

	// Method available only for mocking purposes
	DeleteAtlasResourcePolicyExecute(r DeleteAtlasResourcePolicyApiRequest) (*http.Response, error)

	/*
		GetAtlasResourcePolicies Return all Atlas Resource Policies

		Return all Atlas Resource Policies for the org.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@return GetAtlasResourcePoliciesApiRequest
	*/
	GetAtlasResourcePolicies(ctx context.Context, orgId string) GetAtlasResourcePoliciesApiRequest
	/*
		GetAtlasResourcePolicies Return all Atlas Resource Policies


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetAtlasResourcePoliciesApiParams - Parameters for the request
		@return GetAtlasResourcePoliciesApiRequest
	*/
	GetAtlasResourcePoliciesWithParams(ctx context.Context, args *GetAtlasResourcePoliciesApiParams) GetAtlasResourcePoliciesApiRequest

	// Method available only for mocking purposes
	GetAtlasResourcePoliciesExecute(r GetAtlasResourcePoliciesApiRequest) ([]ApiAtlasResourcePolicy, *http.Response, error)

	/*
		GetAtlasResourcePolicy Return one Atlas Resource Policy

		Return one Atlas Resource Policy for an org.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param resourcePolicyId Unique 24-hexadecimal digit string that identifies an atlas resource policy.
		@return GetAtlasResourcePolicyApiRequest
	*/
	GetAtlasResourcePolicy(ctx context.Context, orgId string, resourcePolicyId string) GetAtlasResourcePolicyApiRequest
	/*
		GetAtlasResourcePolicy Return one Atlas Resource Policy


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetAtlasResourcePolicyApiParams - Parameters for the request
		@return GetAtlasResourcePolicyApiRequest
	*/
	GetAtlasResourcePolicyWithParams(ctx context.Context, args *GetAtlasResourcePolicyApiParams) GetAtlasResourcePolicyApiRequest

	// Method available only for mocking purposes
	GetAtlasResourcePolicyExecute(r GetAtlasResourcePolicyApiRequest) (*ApiAtlasResourcePolicy, *http.Response, error)

	/*
		GetResourcesNonCompliant Return all non-compliant resources

		Return all non-compliant resources for an org.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@return GetResourcesNonCompliantApiRequest
	*/
	GetResourcesNonCompliant(ctx context.Context, orgId string) GetResourcesNonCompliantApiRequest
	/*
		GetResourcesNonCompliant Return all non-compliant resources


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetResourcesNonCompliantApiParams - Parameters for the request
		@return GetResourcesNonCompliantApiRequest
	*/
	GetResourcesNonCompliantWithParams(ctx context.Context, args *GetResourcesNonCompliantApiParams) GetResourcesNonCompliantApiRequest

	// Method available only for mocking purposes
	GetResourcesNonCompliantExecute(r GetResourcesNonCompliantApiRequest) ([]ApiAtlasNonCompliantResource, *http.Response, error)

	/*
		UpdateAtlasResourcePolicy Update one Atlas Resource Policy

		Update one Atlas Resource Policy for an org.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param resourcePolicyId Unique 24-hexadecimal digit string that identifies an atlas resource policy.
		@param apiAtlasResourcePolicyEdit Atlas Resource Policy to update.
		@return UpdateAtlasResourcePolicyApiRequest
	*/
	UpdateAtlasResourcePolicy(ctx context.Context, orgId string, resourcePolicyId string, apiAtlasResourcePolicyEdit *ApiAtlasResourcePolicyEdit) UpdateAtlasResourcePolicyApiRequest
	/*
		UpdateAtlasResourcePolicy Update one Atlas Resource Policy


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateAtlasResourcePolicyApiParams - Parameters for the request
		@return UpdateAtlasResourcePolicyApiRequest
	*/
	UpdateAtlasResourcePolicyWithParams(ctx context.Context, args *UpdateAtlasResourcePolicyApiParams) UpdateAtlasResourcePolicyApiRequest

	// Method available only for mocking purposes
	UpdateAtlasResourcePolicyExecute(r UpdateAtlasResourcePolicyApiRequest) (*ApiAtlasResourcePolicy, *http.Response, error)

	/*
		ValidateAtlasResourcePolicy Validate one Atlas Resource Policy

		Validate one Atlas Resource Policy for an org.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param apiAtlasResourcePolicyCreate Atlas Resource Policy to create.
		@return ValidateAtlasResourcePolicyApiRequest
	*/
	ValidateAtlasResourcePolicy(ctx context.Context, orgId string, apiAtlasResourcePolicyCreate *ApiAtlasResourcePolicyCreate) ValidateAtlasResourcePolicyApiRequest
	/*
		ValidateAtlasResourcePolicy Validate one Atlas Resource Policy


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ValidateAtlasResourcePolicyApiParams - Parameters for the request
		@return ValidateAtlasResourcePolicyApiRequest
	*/
	ValidateAtlasResourcePolicyWithParams(ctx context.Context, args *ValidateAtlasResourcePolicyApiParams) ValidateAtlasResourcePolicyApiRequest

	// Method available only for mocking purposes
	ValidateAtlasResourcePolicyExecute(r ValidateAtlasResourcePolicyApiRequest) (*ApiAtlasResourcePolicy, *http.Response, error)
}

// ResourcePoliciesApiService ResourcePoliciesApi service
type ResourcePoliciesApiService service

type CreateAtlasResourcePolicyApiRequest struct {
	ctx                          context.Context
	ApiService                   ResourcePoliciesApi
	orgId                        string
	apiAtlasResourcePolicyCreate *ApiAtlasResourcePolicyCreate
}

type CreateAtlasResourcePolicyApiParams struct {
	OrgId                        string
	ApiAtlasResourcePolicyCreate *ApiAtlasResourcePolicyCreate
}

func (a *ResourcePoliciesApiService) CreateAtlasResourcePolicyWithParams(ctx context.Context, args *CreateAtlasResourcePolicyApiParams) CreateAtlasResourcePolicyApiRequest {
	return CreateAtlasResourcePolicyApiRequest{
		ApiService:                   a,
		ctx:                          ctx,
		orgId:                        args.OrgId,
		apiAtlasResourcePolicyCreate: args.ApiAtlasResourcePolicyCreate,
	}
}

func (r CreateAtlasResourcePolicyApiRequest) Execute() (*ApiAtlasResourcePolicy, *http.Response, error) {
	return r.ApiService.CreateAtlasResourcePolicyExecute(r)
}

/*
CreateAtlasResourcePolicy Create one Atlas Resource Policy

Create one Atlas Resource Policy for an org.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return CreateAtlasResourcePolicyApiRequest
*/
func (a *ResourcePoliciesApiService) CreateAtlasResourcePolicy(ctx context.Context, orgId string, apiAtlasResourcePolicyCreate *ApiAtlasResourcePolicyCreate) CreateAtlasResourcePolicyApiRequest {
	return CreateAtlasResourcePolicyApiRequest{
		ApiService:                   a,
		ctx:                          ctx,
		orgId:                        orgId,
		apiAtlasResourcePolicyCreate: apiAtlasResourcePolicyCreate,
	}
}

// CreateAtlasResourcePolicyExecute executes the request
//
//	@return ApiAtlasResourcePolicy
func (a *ResourcePoliciesApiService) CreateAtlasResourcePolicyExecute(r CreateAtlasResourcePolicyApiRequest) (*ApiAtlasResourcePolicy, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ApiAtlasResourcePolicy
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourcePoliciesApiService.CreateAtlasResourcePolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/resourcePolicies"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiAtlasResourcePolicyCreate == nil {
		return localVarReturnValue, nil, reportError("apiAtlasResourcePolicyCreate is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2024-08-05+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-08-05+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.apiAtlasResourcePolicyCreate
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

type DeleteAtlasResourcePolicyApiRequest struct {
	ctx              context.Context
	ApiService       ResourcePoliciesApi
	orgId            string
	resourcePolicyId string
}

type DeleteAtlasResourcePolicyApiParams struct {
	OrgId            string
	ResourcePolicyId string
}

func (a *ResourcePoliciesApiService) DeleteAtlasResourcePolicyWithParams(ctx context.Context, args *DeleteAtlasResourcePolicyApiParams) DeleteAtlasResourcePolicyApiRequest {
	return DeleteAtlasResourcePolicyApiRequest{
		ApiService:       a,
		ctx:              ctx,
		orgId:            args.OrgId,
		resourcePolicyId: args.ResourcePolicyId,
	}
}

func (r DeleteAtlasResourcePolicyApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteAtlasResourcePolicyExecute(r)
}

/*
DeleteAtlasResourcePolicy Delete one Atlas Resource Policy

Delete one Atlas Resource Policy for an org.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param resourcePolicyId Unique 24-hexadecimal digit string that identifies an atlas resource policy.
	@return DeleteAtlasResourcePolicyApiRequest
*/
func (a *ResourcePoliciesApiService) DeleteAtlasResourcePolicy(ctx context.Context, orgId string, resourcePolicyId string) DeleteAtlasResourcePolicyApiRequest {
	return DeleteAtlasResourcePolicyApiRequest{
		ApiService:       a,
		ctx:              ctx,
		orgId:            orgId,
		resourcePolicyId: resourcePolicyId,
	}
}

// DeleteAtlasResourcePolicyExecute executes the request
func (a *ResourcePoliciesApiService) DeleteAtlasResourcePolicyExecute(r DeleteAtlasResourcePolicyApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourcePoliciesApiService.DeleteAtlasResourcePolicy")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/resourcePolicies/{resourcePolicyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resourcePolicyId"+"}", url.PathEscape(r.resourcePolicyId), -1)

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

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-08-05+json"}

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

type GetAtlasResourcePoliciesApiRequest struct {
	ctx        context.Context
	ApiService ResourcePoliciesApi
	orgId      string
}

type GetAtlasResourcePoliciesApiParams struct {
	OrgId string
}

func (a *ResourcePoliciesApiService) GetAtlasResourcePoliciesWithParams(ctx context.Context, args *GetAtlasResourcePoliciesApiParams) GetAtlasResourcePoliciesApiRequest {
	return GetAtlasResourcePoliciesApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      args.OrgId,
	}
}

func (r GetAtlasResourcePoliciesApiRequest) Execute() ([]ApiAtlasResourcePolicy, *http.Response, error) {
	return r.ApiService.GetAtlasResourcePoliciesExecute(r)
}

/*
GetAtlasResourcePolicies Return all Atlas Resource Policies

Return all Atlas Resource Policies for the org.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return GetAtlasResourcePoliciesApiRequest
*/
func (a *ResourcePoliciesApiService) GetAtlasResourcePolicies(ctx context.Context, orgId string) GetAtlasResourcePoliciesApiRequest {
	return GetAtlasResourcePoliciesApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
	}
}

// GetAtlasResourcePoliciesExecute executes the request
//
//	@return []ApiAtlasResourcePolicy
func (a *ResourcePoliciesApiService) GetAtlasResourcePoliciesExecute(r GetAtlasResourcePoliciesApiRequest) ([]ApiAtlasResourcePolicy, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue []ApiAtlasResourcePolicy
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourcePoliciesApiService.GetAtlasResourcePolicies")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/resourcePolicies"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)

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

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-08-05+json"}

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

type GetAtlasResourcePolicyApiRequest struct {
	ctx              context.Context
	ApiService       ResourcePoliciesApi
	orgId            string
	resourcePolicyId string
}

type GetAtlasResourcePolicyApiParams struct {
	OrgId            string
	ResourcePolicyId string
}

func (a *ResourcePoliciesApiService) GetAtlasResourcePolicyWithParams(ctx context.Context, args *GetAtlasResourcePolicyApiParams) GetAtlasResourcePolicyApiRequest {
	return GetAtlasResourcePolicyApiRequest{
		ApiService:       a,
		ctx:              ctx,
		orgId:            args.OrgId,
		resourcePolicyId: args.ResourcePolicyId,
	}
}

func (r GetAtlasResourcePolicyApiRequest) Execute() (*ApiAtlasResourcePolicy, *http.Response, error) {
	return r.ApiService.GetAtlasResourcePolicyExecute(r)
}

/*
GetAtlasResourcePolicy Return one Atlas Resource Policy

Return one Atlas Resource Policy for an org.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param resourcePolicyId Unique 24-hexadecimal digit string that identifies an atlas resource policy.
	@return GetAtlasResourcePolicyApiRequest
*/
func (a *ResourcePoliciesApiService) GetAtlasResourcePolicy(ctx context.Context, orgId string, resourcePolicyId string) GetAtlasResourcePolicyApiRequest {
	return GetAtlasResourcePolicyApiRequest{
		ApiService:       a,
		ctx:              ctx,
		orgId:            orgId,
		resourcePolicyId: resourcePolicyId,
	}
}

// GetAtlasResourcePolicyExecute executes the request
//
//	@return ApiAtlasResourcePolicy
func (a *ResourcePoliciesApiService) GetAtlasResourcePolicyExecute(r GetAtlasResourcePolicyApiRequest) (*ApiAtlasResourcePolicy, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ApiAtlasResourcePolicy
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourcePoliciesApiService.GetAtlasResourcePolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/resourcePolicies/{resourcePolicyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resourcePolicyId"+"}", url.PathEscape(r.resourcePolicyId), -1)

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

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-08-05+json"}

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

type GetResourcesNonCompliantApiRequest struct {
	ctx        context.Context
	ApiService ResourcePoliciesApi
	orgId      string
}

type GetResourcesNonCompliantApiParams struct {
	OrgId string
}

func (a *ResourcePoliciesApiService) GetResourcesNonCompliantWithParams(ctx context.Context, args *GetResourcesNonCompliantApiParams) GetResourcesNonCompliantApiRequest {
	return GetResourcesNonCompliantApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      args.OrgId,
	}
}

func (r GetResourcesNonCompliantApiRequest) Execute() ([]ApiAtlasNonCompliantResource, *http.Response, error) {
	return r.ApiService.GetResourcesNonCompliantExecute(r)
}

/*
GetResourcesNonCompliant Return all non-compliant resources

Return all non-compliant resources for an org.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return GetResourcesNonCompliantApiRequest
*/
func (a *ResourcePoliciesApiService) GetResourcesNonCompliant(ctx context.Context, orgId string) GetResourcesNonCompliantApiRequest {
	return GetResourcesNonCompliantApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
	}
}

// GetResourcesNonCompliantExecute executes the request
//
//	@return []ApiAtlasNonCompliantResource
func (a *ResourcePoliciesApiService) GetResourcesNonCompliantExecute(r GetResourcesNonCompliantApiRequest) ([]ApiAtlasNonCompliantResource, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue []ApiAtlasNonCompliantResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourcePoliciesApiService.GetResourcesNonCompliant")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/nonCompliantResources"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)

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

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-08-05+json"}

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

type UpdateAtlasResourcePolicyApiRequest struct {
	ctx                        context.Context
	ApiService                 ResourcePoliciesApi
	orgId                      string
	resourcePolicyId           string
	apiAtlasResourcePolicyEdit *ApiAtlasResourcePolicyEdit
}

type UpdateAtlasResourcePolicyApiParams struct {
	OrgId                      string
	ResourcePolicyId           string
	ApiAtlasResourcePolicyEdit *ApiAtlasResourcePolicyEdit
}

func (a *ResourcePoliciesApiService) UpdateAtlasResourcePolicyWithParams(ctx context.Context, args *UpdateAtlasResourcePolicyApiParams) UpdateAtlasResourcePolicyApiRequest {
	return UpdateAtlasResourcePolicyApiRequest{
		ApiService:                 a,
		ctx:                        ctx,
		orgId:                      args.OrgId,
		resourcePolicyId:           args.ResourcePolicyId,
		apiAtlasResourcePolicyEdit: args.ApiAtlasResourcePolicyEdit,
	}
}

func (r UpdateAtlasResourcePolicyApiRequest) Execute() (*ApiAtlasResourcePolicy, *http.Response, error) {
	return r.ApiService.UpdateAtlasResourcePolicyExecute(r)
}

/*
UpdateAtlasResourcePolicy Update one Atlas Resource Policy

Update one Atlas Resource Policy for an org.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param resourcePolicyId Unique 24-hexadecimal digit string that identifies an atlas resource policy.
	@return UpdateAtlasResourcePolicyApiRequest
*/
func (a *ResourcePoliciesApiService) UpdateAtlasResourcePolicy(ctx context.Context, orgId string, resourcePolicyId string, apiAtlasResourcePolicyEdit *ApiAtlasResourcePolicyEdit) UpdateAtlasResourcePolicyApiRequest {
	return UpdateAtlasResourcePolicyApiRequest{
		ApiService:                 a,
		ctx:                        ctx,
		orgId:                      orgId,
		resourcePolicyId:           resourcePolicyId,
		apiAtlasResourcePolicyEdit: apiAtlasResourcePolicyEdit,
	}
}

// UpdateAtlasResourcePolicyExecute executes the request
//
//	@return ApiAtlasResourcePolicy
func (a *ResourcePoliciesApiService) UpdateAtlasResourcePolicyExecute(r UpdateAtlasResourcePolicyApiRequest) (*ApiAtlasResourcePolicy, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ApiAtlasResourcePolicy
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourcePoliciesApiService.UpdateAtlasResourcePolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/resourcePolicies/{resourcePolicyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"resourcePolicyId"+"}", url.PathEscape(r.resourcePolicyId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiAtlasResourcePolicyEdit == nil {
		return localVarReturnValue, nil, reportError("apiAtlasResourcePolicyEdit is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2024-08-05+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-08-05+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.apiAtlasResourcePolicyEdit
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

type ValidateAtlasResourcePolicyApiRequest struct {
	ctx                          context.Context
	ApiService                   ResourcePoliciesApi
	orgId                        string
	apiAtlasResourcePolicyCreate *ApiAtlasResourcePolicyCreate
}

type ValidateAtlasResourcePolicyApiParams struct {
	OrgId                        string
	ApiAtlasResourcePolicyCreate *ApiAtlasResourcePolicyCreate
}

func (a *ResourcePoliciesApiService) ValidateAtlasResourcePolicyWithParams(ctx context.Context, args *ValidateAtlasResourcePolicyApiParams) ValidateAtlasResourcePolicyApiRequest {
	return ValidateAtlasResourcePolicyApiRequest{
		ApiService:                   a,
		ctx:                          ctx,
		orgId:                        args.OrgId,
		apiAtlasResourcePolicyCreate: args.ApiAtlasResourcePolicyCreate,
	}
}

func (r ValidateAtlasResourcePolicyApiRequest) Execute() (*ApiAtlasResourcePolicy, *http.Response, error) {
	return r.ApiService.ValidateAtlasResourcePolicyExecute(r)
}

/*
ValidateAtlasResourcePolicy Validate one Atlas Resource Policy

Validate one Atlas Resource Policy for an org.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return ValidateAtlasResourcePolicyApiRequest
*/
func (a *ResourcePoliciesApiService) ValidateAtlasResourcePolicy(ctx context.Context, orgId string, apiAtlasResourcePolicyCreate *ApiAtlasResourcePolicyCreate) ValidateAtlasResourcePolicyApiRequest {
	return ValidateAtlasResourcePolicyApiRequest{
		ApiService:                   a,
		ctx:                          ctx,
		orgId:                        orgId,
		apiAtlasResourcePolicyCreate: apiAtlasResourcePolicyCreate,
	}
}

// ValidateAtlasResourcePolicyExecute executes the request
//
//	@return ApiAtlasResourcePolicy
func (a *ResourcePoliciesApiService) ValidateAtlasResourcePolicyExecute(r ValidateAtlasResourcePolicyApiRequest) (*ApiAtlasResourcePolicy, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ApiAtlasResourcePolicy
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ResourcePoliciesApiService.ValidateAtlasResourcePolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/resourcePolicies:validate"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiAtlasResourcePolicyCreate == nil {
		return localVarReturnValue, nil, reportError("apiAtlasResourcePolicyCreate is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2024-08-05+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-08-05+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.apiAtlasResourcePolicyCreate
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
