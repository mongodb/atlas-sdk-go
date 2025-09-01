// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

type InvoicesApi interface {

	/*
		CreateCostExplorerQueryProcess Create One Cost Explorer Query Process

		Creates a query process within the Cost Explorer for the given parameters. A token is returned that can be used to poll the status of the query and eventually retrieve the results.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param costExplorerFilterRequestBody Filter parameters for the Cost Explorer query.
		@return CreateCostExplorerQueryProcessApiRequest
	*/
	CreateCostExplorerQueryProcess(ctx context.Context, orgId string, costExplorerFilterRequestBody *CostExplorerFilterRequestBody) CreateCostExplorerQueryProcessApiRequest
	/*
		CreateCostExplorerQueryProcess Create One Cost Explorer Query Process


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateCostExplorerQueryProcessApiParams - Parameters for the request
		@return CreateCostExplorerQueryProcessApiRequest
	*/
	CreateCostExplorerQueryProcessWithParams(ctx context.Context, args *CreateCostExplorerQueryProcessApiParams) CreateCostExplorerQueryProcessApiRequest

	// Method available only for mocking purposes
	CreateCostExplorerQueryProcessExecute(r CreateCostExplorerQueryProcessApiRequest) (*CostExplorerFilterResponse, *http.Response, error)

	/*
			DownloadInvoiceCsv Return One Invoice as CSV for One Organization

			Returns one invoice that MongoDB issued to the specified organization in CSV format. A unique 24-hexadecimal digit string identifies the invoice. To use this resource, the requesting Service Account or API Key have at least the Organization Billing Viewer, Organization Billing Admin, or Organization Owner role. If you have a cross-organization setup, you can query for a linked invoice if you have the Organization Billing Admin or Organization Owner Role.
		 To compute the total owed amount of the invoice - sum up total owed amount of each payment included into the invoice. To compute payment's owed amount - use formula *totalBilledCents* * *unitPrice* + *salesTax* - *startingBalanceCents*.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
			@param invoiceId Unique 24-hexadecimal digit string that identifies the invoice submitted to the specified organization. Charges typically post the next day.
			@return DownloadInvoiceCsvApiRequest
	*/
	DownloadInvoiceCsv(ctx context.Context, orgId string, invoiceId string) DownloadInvoiceCsvApiRequest
	/*
		DownloadInvoiceCsv Return One Invoice as CSV for One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DownloadInvoiceCsvApiParams - Parameters for the request
		@return DownloadInvoiceCsvApiRequest
	*/
	DownloadInvoiceCsvWithParams(ctx context.Context, args *DownloadInvoiceCsvApiParams) DownloadInvoiceCsvApiRequest

	// Method available only for mocking purposes
	DownloadInvoiceCsvExecute(r DownloadInvoiceCsvApiRequest) (string, *http.Response, error)

	/*
		GetCostExplorerQueryProcess Return results from a given Cost Explorer query, or notify that the results are not ready yet.

		Returns the usage details for a Cost Explorer query, if the query is finished and the data is ready to be viewed. If the data is not ready, a 'processing' response willindicate that another request should be sent later to view the data.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param token Unique 64 digit string that identifies the Cost Explorer query.
		@return GetCostExplorerQueryProcessApiRequest
	*/
	GetCostExplorerQueryProcess(ctx context.Context, orgId string, token string) GetCostExplorerQueryProcessApiRequest
	/*
		GetCostExplorerQueryProcess Return results from a given Cost Explorer query, or notify that the results are not ready yet.


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetCostExplorerQueryProcessApiParams - Parameters for the request
		@return GetCostExplorerQueryProcessApiRequest
	*/
	GetCostExplorerQueryProcessWithParams(ctx context.Context, args *GetCostExplorerQueryProcessApiParams) GetCostExplorerQueryProcessApiRequest

	// Method available only for mocking purposes
	GetCostExplorerQueryProcessExecute(r GetCostExplorerQueryProcessApiRequest) (any, *http.Response, error)

	/*
			GetInvoice Return One Invoice for One Organization

			Returns one invoice that MongoDB issued to the specified organization. A unique 24-hexadecimal digit string identifies the invoice. You can choose to receive this invoice in JSON or CSV format. To use this resource, the requesting Service Account or API Key must have the Organization Billing Viewer, Organization Billing Admin, or Organization Owner role. If you have a cross-organization setup, you can query for a linked invoice if you have the Organization Billing Admin or Organization Owner role.
		To compute the total owed amount of the invoice - sum up total owed amount of each payment included into the invoice. To compute payment's owed amount - use formula *totalBilledCents* * *unitPrice* + *salesTax* - *startingBalanceCents*.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
			@param invoiceId Unique 24-hexadecimal digit string that identifies the invoice submitted to the specified organization. Charges typically post the next day.
			@return GetInvoiceApiRequest
	*/
	GetInvoice(ctx context.Context, orgId string, invoiceId string) GetInvoiceApiRequest
	/*
		GetInvoice Return One Invoice for One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetInvoiceApiParams - Parameters for the request
		@return GetInvoiceApiRequest
	*/
	GetInvoiceWithParams(ctx context.Context, args *GetInvoiceApiParams) GetInvoiceApiRequest

	// Method available only for mocking purposes
	GetInvoiceExecute(r GetInvoiceApiRequest) (*BillingInvoice, *http.Response, error)

	/*
			ListInvoices Return All Invoices for One Organization

			Returns all invoices that MongoDB issued to the specified organization. This list includes all invoices regardless of invoice status. To use this resource, the requesting Service Account or API Key must have the Organization Billing Viewer, Organization Billing Admin, or Organization Owner role. If you have a cross-organization setup, you can view linked invoices if you have the Organization Billing Admin or Organization Owner role.
		To compute the total owed amount of the invoices - sum up total owed of each invoice. It could be computed as a sum of owed amount of each payment included into the invoice. To compute payment's owed amount - use formula *totalBilledCents* * *unitPrice* + *salesTax* - *startingBalanceCents*.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
			@return ListInvoicesApiRequest
	*/
	ListInvoices(ctx context.Context, orgId string) ListInvoicesApiRequest
	/*
		ListInvoices Return All Invoices for One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListInvoicesApiParams - Parameters for the request
		@return ListInvoicesApiRequest
	*/
	ListInvoicesWithParams(ctx context.Context, args *ListInvoicesApiParams) ListInvoicesApiRequest

	// Method available only for mocking purposes
	ListInvoicesExecute(r ListInvoicesApiRequest) (*PaginatedApiInvoiceMetadata, *http.Response, error)

	/*
		ListPendingInvoices Return All Pending Invoices for One Organization

		Returns all invoices accruing charges for the current billing cycle for the specified organization. To use this resource, the requesting Service Account or API Key must have the Organization Billing Viewer, Organization Billing Admin, or Organization Owner role. If you have a cross-organization setup, you can view linked invoices if you have the Organization Billing Admin or Organization Owner Role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@return ListPendingInvoicesApiRequest
	*/
	ListPendingInvoices(ctx context.Context, orgId string) ListPendingInvoicesApiRequest
	/*
		ListPendingInvoices Return All Pending Invoices for One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListPendingInvoicesApiParams - Parameters for the request
		@return ListPendingInvoicesApiRequest
	*/
	ListPendingInvoicesWithParams(ctx context.Context, args *ListPendingInvoicesApiParams) ListPendingInvoicesApiRequest

	// Method available only for mocking purposes
	ListPendingInvoicesExecute(r ListPendingInvoicesApiRequest) (*PaginatedApiInvoice, *http.Response, error)

	/*
		QueryLineItemsFromSingleInvoice Return All Line Items for One Invoice by Invoice ID

		Query the lineItems of the specified invoice and return the result JSON. A unique 24-hexadecimal digit string identifies the invoice.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param invoiceId Unique 24-hexadecimal digit string that identifies the invoice submitted to the specified organization. Charges typically post the next day.
		@param apiPublicUsageDetailsQueryRequest Filter parameters for the lineItems query. Send a request with an empty JSON body to retrieve all line items for a given invoiceID without applying any filters.
		@return QueryLineItemsFromSingleInvoiceApiRequest
	*/
	QueryLineItemsFromSingleInvoice(ctx context.Context, orgId string, invoiceId string, apiPublicUsageDetailsQueryRequest *ApiPublicUsageDetailsQueryRequest) QueryLineItemsFromSingleInvoiceApiRequest
	/*
		QueryLineItemsFromSingleInvoice Return All Line Items for One Invoice by Invoice ID


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param QueryLineItemsFromSingleInvoiceApiParams - Parameters for the request
		@return QueryLineItemsFromSingleInvoiceApiRequest
	*/
	QueryLineItemsFromSingleInvoiceWithParams(ctx context.Context, args *QueryLineItemsFromSingleInvoiceApiParams) QueryLineItemsFromSingleInvoiceApiRequest

	// Method available only for mocking purposes
	QueryLineItemsFromSingleInvoiceExecute(r QueryLineItemsFromSingleInvoiceApiRequest) (*PaginatedPublicApiUsageDetailsLineItem, *http.Response, error)
}

// InvoicesApiService InvoicesApi service
type InvoicesApiService service

type CreateCostExplorerQueryProcessApiRequest struct {
	ctx                           context.Context
	ApiService                    InvoicesApi
	orgId                         string
	costExplorerFilterRequestBody *CostExplorerFilterRequestBody
}

type CreateCostExplorerQueryProcessApiParams struct {
	OrgId                         string
	CostExplorerFilterRequestBody *CostExplorerFilterRequestBody
}

func (a *InvoicesApiService) CreateCostExplorerQueryProcessWithParams(ctx context.Context, args *CreateCostExplorerQueryProcessApiParams) CreateCostExplorerQueryProcessApiRequest {
	return CreateCostExplorerQueryProcessApiRequest{
		ApiService:                    a,
		ctx:                           ctx,
		orgId:                         args.OrgId,
		costExplorerFilterRequestBody: args.CostExplorerFilterRequestBody,
	}
}

func (r CreateCostExplorerQueryProcessApiRequest) Execute() (*CostExplorerFilterResponse, *http.Response, error) {
	return r.ApiService.CreateCostExplorerQueryProcessExecute(r)
}

/*
CreateCostExplorerQueryProcess Create One Cost Explorer Query Process

Creates a query process within the Cost Explorer for the given parameters. A token is returned that can be used to poll the status of the query and eventually retrieve the results.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return CreateCostExplorerQueryProcessApiRequest
*/
func (a *InvoicesApiService) CreateCostExplorerQueryProcess(ctx context.Context, orgId string, costExplorerFilterRequestBody *CostExplorerFilterRequestBody) CreateCostExplorerQueryProcessApiRequest {
	return CreateCostExplorerQueryProcessApiRequest{
		ApiService:                    a,
		ctx:                           ctx,
		orgId:                         orgId,
		costExplorerFilterRequestBody: costExplorerFilterRequestBody,
	}
}

// CreateCostExplorerQueryProcessExecute executes the request
//
//	@return CostExplorerFilterResponse
func (a *InvoicesApiService) CreateCostExplorerQueryProcessExecute(r CreateCostExplorerQueryProcessApiRequest) (*CostExplorerFilterResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *CostExplorerFilterResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoicesApiService.CreateCostExplorerQueryProcess")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/billing/costExplorer/usage"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.costExplorerFilterRequestBody == nil {
		return localVarReturnValue, nil, reportError("costExplorerFilterRequestBody is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2023-01-01+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.costExplorerFilterRequestBody
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

type DownloadInvoiceCsvApiRequest struct {
	ctx        context.Context
	ApiService InvoicesApi
	orgId      string
	invoiceId  string
}

type DownloadInvoiceCsvApiParams struct {
	OrgId     string
	InvoiceId string
}

func (a *InvoicesApiService) DownloadInvoiceCsvWithParams(ctx context.Context, args *DownloadInvoiceCsvApiParams) DownloadInvoiceCsvApiRequest {
	return DownloadInvoiceCsvApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      args.OrgId,
		invoiceId:  args.InvoiceId,
	}
}

func (r DownloadInvoiceCsvApiRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.DownloadInvoiceCsvExecute(r)
}

/*
DownloadInvoiceCsv Return One Invoice as CSV for One Organization

Returns one invoice that MongoDB issued to the specified organization in CSV format. A unique 24-hexadecimal digit string identifies the invoice. To use this resource, the requesting Service Account or API Key have at least the Organization Billing Viewer, Organization Billing Admin, or Organization Owner role. If you have a cross-organization setup, you can query for a linked invoice if you have the Organization Billing Admin or Organization Owner Role.

	To compute the total owed amount of the invoice - sum up total owed amount of each payment included into the invoice. To compute payment's owed amount - use formula *totalBilledCents* * *unitPrice* + *salesTax* - *startingBalanceCents*.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param invoiceId Unique 24-hexadecimal digit string that identifies the invoice submitted to the specified organization. Charges typically post the next day.
	@return DownloadInvoiceCsvApiRequest
*/
func (a *InvoicesApiService) DownloadInvoiceCsv(ctx context.Context, orgId string, invoiceId string) DownloadInvoiceCsvApiRequest {
	return DownloadInvoiceCsvApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
		invoiceId:  invoiceId,
	}
}

// DownloadInvoiceCsvExecute executes the request
//
//	@return string
func (a *InvoicesApiService) DownloadInvoiceCsvExecute(r DownloadInvoiceCsvApiRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoicesApiService.DownloadInvoiceCsv")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/invoices/{invoiceId}/csv"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	if r.invoiceId == "" {
		return localVarReturnValue, nil, reportError("invoiceId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"invoiceId"+"}", url.PathEscape(r.invoiceId), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+csv"}

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

type GetCostExplorerQueryProcessApiRequest struct {
	ctx        context.Context
	ApiService InvoicesApi
	orgId      string
	token      string
}

type GetCostExplorerQueryProcessApiParams struct {
	OrgId string
	Token string
}

func (a *InvoicesApiService) GetCostExplorerQueryProcessWithParams(ctx context.Context, args *GetCostExplorerQueryProcessApiParams) GetCostExplorerQueryProcessApiRequest {
	return GetCostExplorerQueryProcessApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      args.OrgId,
		token:      args.Token,
	}
}

func (r GetCostExplorerQueryProcessApiRequest) Execute() (any, *http.Response, error) {
	return r.ApiService.GetCostExplorerQueryProcessExecute(r)
}

/*
GetCostExplorerQueryProcess Return results from a given Cost Explorer query, or notify that the results are not ready yet.

Returns the usage details for a Cost Explorer query, if the query is finished and the data is ready to be viewed. If the data is not ready, a 'processing' response willindicate that another request should be sent later to view the data.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param token Unique 64 digit string that identifies the Cost Explorer query.
	@return GetCostExplorerQueryProcessApiRequest
*/
func (a *InvoicesApiService) GetCostExplorerQueryProcess(ctx context.Context, orgId string, token string) GetCostExplorerQueryProcessApiRequest {
	return GetCostExplorerQueryProcessApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
		token:      token,
	}
}

// GetCostExplorerQueryProcessExecute executes the request
//
//	@return any
func (a *InvoicesApiService) GetCostExplorerQueryProcessExecute(r GetCostExplorerQueryProcessApiRequest) (any, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue any
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoicesApiService.GetCostExplorerQueryProcess")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/billing/costExplorer/usage/{token}"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	if r.token == "" {
		return localVarReturnValue, nil, reportError("token is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"token"+"}", url.PathEscape(r.token), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

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

type GetInvoiceApiRequest struct {
	ctx        context.Context
	ApiService InvoicesApi
	orgId      string
	invoiceId  string
}

type GetInvoiceApiParams struct {
	OrgId     string
	InvoiceId string
}

func (a *InvoicesApiService) GetInvoiceWithParams(ctx context.Context, args *GetInvoiceApiParams) GetInvoiceApiRequest {
	return GetInvoiceApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      args.OrgId,
		invoiceId:  args.InvoiceId,
	}
}

func (r GetInvoiceApiRequest) Execute() (*BillingInvoice, *http.Response, error) {
	return r.ApiService.GetInvoiceExecute(r)
}

/*
GetInvoice Return One Invoice for One Organization

Returns one invoice that MongoDB issued to the specified organization. A unique 24-hexadecimal digit string identifies the invoice. You can choose to receive this invoice in JSON or CSV format. To use this resource, the requesting Service Account or API Key must have the Organization Billing Viewer, Organization Billing Admin, or Organization Owner role. If you have a cross-organization setup, you can query for a linked invoice if you have the Organization Billing Admin or Organization Owner role.
To compute the total owed amount of the invoice - sum up total owed amount of each payment included into the invoice. To compute payment's owed amount - use formula *totalBilledCents* * *unitPrice* + *salesTax* - *startingBalanceCents*.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param invoiceId Unique 24-hexadecimal digit string that identifies the invoice submitted to the specified organization. Charges typically post the next day.
	@return GetInvoiceApiRequest
*/
func (a *InvoicesApiService) GetInvoice(ctx context.Context, orgId string, invoiceId string) GetInvoiceApiRequest {
	return GetInvoiceApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
		invoiceId:  invoiceId,
	}
}

// GetInvoiceExecute executes the request
//
//	@return BillingInvoice
func (a *InvoicesApiService) GetInvoiceExecute(r GetInvoiceApiRequest) (*BillingInvoice, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *BillingInvoice
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoicesApiService.GetInvoice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/invoices/{invoiceId}"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	if r.invoiceId == "" {
		return localVarReturnValue, nil, reportError("invoiceId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"invoiceId"+"}", url.PathEscape(r.invoiceId), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

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

type ListInvoicesApiRequest struct {
	ctx                context.Context
	ApiService         InvoicesApi
	orgId              string
	includeCount       *bool
	itemsPerPage       *int
	pageNum            *int
	viewLinkedInvoices *bool
	statusNames        *[]string
	fromDate           *string
	toDate             *string
	sortBy             *string
	orderBy            *string
}

type ListInvoicesApiParams struct {
	OrgId              string
	IncludeCount       *bool
	ItemsPerPage       *int
	PageNum            *int
	ViewLinkedInvoices *bool
	StatusNames        *[]string
	FromDate           *string
	ToDate             *string
	SortBy             *string
	OrderBy            *string
}

func (a *InvoicesApiService) ListInvoicesWithParams(ctx context.Context, args *ListInvoicesApiParams) ListInvoicesApiRequest {
	return ListInvoicesApiRequest{
		ApiService:         a,
		ctx:                ctx,
		orgId:              args.OrgId,
		includeCount:       args.IncludeCount,
		itemsPerPage:       args.ItemsPerPage,
		pageNum:            args.PageNum,
		viewLinkedInvoices: args.ViewLinkedInvoices,
		statusNames:        args.StatusNames,
		fromDate:           args.FromDate,
		toDate:             args.ToDate,
		sortBy:             args.SortBy,
		orderBy:            args.OrderBy,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListInvoicesApiRequest) IncludeCount(includeCount bool) ListInvoicesApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListInvoicesApiRequest) ItemsPerPage(itemsPerPage int) ListInvoicesApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListInvoicesApiRequest) PageNum(pageNum int) ListInvoicesApiRequest {
	r.pageNum = &pageNum
	return r
}

// Flag that indicates whether to return linked invoices in the linkedInvoices field.
func (r ListInvoicesApiRequest) ViewLinkedInvoices(viewLinkedInvoices bool) ListInvoicesApiRequest {
	r.viewLinkedInvoices = &viewLinkedInvoices
	return r
}

// Statuses of the invoice to be retrieved. Omit to return invoices of all statuses.
func (r ListInvoicesApiRequest) StatusNames(statusNames []string) ListInvoicesApiRequest {
	r.statusNames = &statusNames
	return r
}

// Retrieve the invoices the startDates of which are greater than or equal to the fromDate. If omit, the invoices return will go back to earliest startDate.
func (r ListInvoicesApiRequest) FromDate(fromDate string) ListInvoicesApiRequest {
	r.fromDate = &fromDate
	return r
}

// Retrieve the invoices the endDates of which are smaller than or equal to the toDate. If omit, the invoices return will go further to latest endDate.
func (r ListInvoicesApiRequest) ToDate(toDate string) ListInvoicesApiRequest {
	r.toDate = &toDate
	return r
}

// Field used to sort the returned invoices by. Use in combination with orderBy parameter to control the order of the result.
func (r ListInvoicesApiRequest) SortBy(sortBy string) ListInvoicesApiRequest {
	r.sortBy = &sortBy
	return r
}

// Field used to order the returned invoices by. Use in combination of sortBy parameter to control the order of the result.
func (r ListInvoicesApiRequest) OrderBy(orderBy string) ListInvoicesApiRequest {
	r.orderBy = &orderBy
	return r
}

func (r ListInvoicesApiRequest) Execute() (*PaginatedApiInvoiceMetadata, *http.Response, error) {
	return r.ApiService.ListInvoicesExecute(r)
}

/*
ListInvoices Return All Invoices for One Organization

Returns all invoices that MongoDB issued to the specified organization. This list includes all invoices regardless of invoice status. To use this resource, the requesting Service Account or API Key must have the Organization Billing Viewer, Organization Billing Admin, or Organization Owner role. If you have a cross-organization setup, you can view linked invoices if you have the Organization Billing Admin or Organization Owner role.
To compute the total owed amount of the invoices - sum up total owed of each invoice. It could be computed as a sum of owed amount of each payment included into the invoice. To compute payment's owed amount - use formula *totalBilledCents* * *unitPrice* + *salesTax* - *startingBalanceCents*.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return ListInvoicesApiRequest
*/
func (a *InvoicesApiService) ListInvoices(ctx context.Context, orgId string) ListInvoicesApiRequest {
	return ListInvoicesApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
	}
}

// ListInvoicesExecute executes the request
//
//	@return PaginatedApiInvoiceMetadata
func (a *InvoicesApiService) ListInvoicesExecute(r ListInvoicesApiRequest) (*PaginatedApiInvoiceMetadata, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedApiInvoiceMetadata
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoicesApiService.ListInvoices")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/invoices"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.includeCount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	} else {
		var defaultValue bool = true
		r.includeCount = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	}
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
	if r.viewLinkedInvoices != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "viewLinkedInvoices", r.viewLinkedInvoices, "")
	} else {
		var defaultValue bool = true
		r.viewLinkedInvoices = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "viewLinkedInvoices", r.viewLinkedInvoices, "")
	}
	if r.statusNames != nil {
		t := *r.statusNames
		// Workaround for unused import
		_ = reflect.Append
		parameterAddToHeaderOrQuery(localVarQueryParams, "statusNames", t, "multi")

	}
	if r.fromDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fromDate", r.fromDate, "")
	}
	if r.toDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "toDate", r.toDate, "")
	}
	if r.sortBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sortBy", r.sortBy, "")
	} else {
		var defaultValue string = "END_DATE"
		r.sortBy = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "sortBy", r.sortBy, "")
	}
	if r.orderBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "orderBy", r.orderBy, "")
	} else {
		var defaultValue string = "desc"
		r.orderBy = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "orderBy", r.orderBy, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

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

type ListPendingInvoicesApiRequest struct {
	ctx        context.Context
	ApiService InvoicesApi
	orgId      string
}

type ListPendingInvoicesApiParams struct {
	OrgId string
}

func (a *InvoicesApiService) ListPendingInvoicesWithParams(ctx context.Context, args *ListPendingInvoicesApiParams) ListPendingInvoicesApiRequest {
	return ListPendingInvoicesApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      args.OrgId,
	}
}

func (r ListPendingInvoicesApiRequest) Execute() (*PaginatedApiInvoice, *http.Response, error) {
	return r.ApiService.ListPendingInvoicesExecute(r)
}

/*
ListPendingInvoices Return All Pending Invoices for One Organization

Returns all invoices accruing charges for the current billing cycle for the specified organization. To use this resource, the requesting Service Account or API Key must have the Organization Billing Viewer, Organization Billing Admin, or Organization Owner role. If you have a cross-organization setup, you can view linked invoices if you have the Organization Billing Admin or Organization Owner Role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return ListPendingInvoicesApiRequest
*/
func (a *InvoicesApiService) ListPendingInvoices(ctx context.Context, orgId string) ListPendingInvoicesApiRequest {
	return ListPendingInvoicesApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
	}
}

// ListPendingInvoicesExecute executes the request
//
//	@return PaginatedApiInvoice
func (a *InvoicesApiService) ListPendingInvoicesExecute(r ListPendingInvoicesApiRequest) (*PaginatedApiInvoice, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedApiInvoice
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoicesApiService.ListPendingInvoices")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/invoices/pending"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

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

type QueryLineItemsFromSingleInvoiceApiRequest struct {
	ctx                               context.Context
	ApiService                        InvoicesApi
	orgId                             string
	invoiceId                         string
	apiPublicUsageDetailsQueryRequest *ApiPublicUsageDetailsQueryRequest
	itemsPerPage                      *int
	pageNum                           *int
}

type QueryLineItemsFromSingleInvoiceApiParams struct {
	OrgId                             string
	InvoiceId                         string
	ApiPublicUsageDetailsQueryRequest *ApiPublicUsageDetailsQueryRequest
	ItemsPerPage                      *int
	PageNum                           *int
}

func (a *InvoicesApiService) QueryLineItemsFromSingleInvoiceWithParams(ctx context.Context, args *QueryLineItemsFromSingleInvoiceApiParams) QueryLineItemsFromSingleInvoiceApiRequest {
	return QueryLineItemsFromSingleInvoiceApiRequest{
		ApiService:                        a,
		ctx:                               ctx,
		orgId:                             args.OrgId,
		invoiceId:                         args.InvoiceId,
		apiPublicUsageDetailsQueryRequest: args.ApiPublicUsageDetailsQueryRequest,
		itemsPerPage:                      args.ItemsPerPage,
		pageNum:                           args.PageNum,
	}
}

// Number of items that the response returns per page.
func (r QueryLineItemsFromSingleInvoiceApiRequest) ItemsPerPage(itemsPerPage int) QueryLineItemsFromSingleInvoiceApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r QueryLineItemsFromSingleInvoiceApiRequest) PageNum(pageNum int) QueryLineItemsFromSingleInvoiceApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r QueryLineItemsFromSingleInvoiceApiRequest) Execute() (*PaginatedPublicApiUsageDetailsLineItem, *http.Response, error) {
	return r.ApiService.QueryLineItemsFromSingleInvoiceExecute(r)
}

/*
QueryLineItemsFromSingleInvoice Return All Line Items for One Invoice by Invoice ID

Query the lineItems of the specified invoice and return the result JSON. A unique 24-hexadecimal digit string identifies the invoice.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param invoiceId Unique 24-hexadecimal digit string that identifies the invoice submitted to the specified organization. Charges typically post the next day.
	@return QueryLineItemsFromSingleInvoiceApiRequest
*/
func (a *InvoicesApiService) QueryLineItemsFromSingleInvoice(ctx context.Context, orgId string, invoiceId string, apiPublicUsageDetailsQueryRequest *ApiPublicUsageDetailsQueryRequest) QueryLineItemsFromSingleInvoiceApiRequest {
	return QueryLineItemsFromSingleInvoiceApiRequest{
		ApiService:                        a,
		ctx:                               ctx,
		orgId:                             orgId,
		invoiceId:                         invoiceId,
		apiPublicUsageDetailsQueryRequest: apiPublicUsageDetailsQueryRequest,
	}
}

// QueryLineItemsFromSingleInvoiceExecute executes the request
//
//	@return PaginatedPublicApiUsageDetailsLineItem
func (a *InvoicesApiService) QueryLineItemsFromSingleInvoiceExecute(r QueryLineItemsFromSingleInvoiceApiRequest) (*PaginatedPublicApiUsageDetailsLineItem, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedPublicApiUsageDetailsLineItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoicesApiService.QueryLineItemsFromSingleInvoice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/invoices/{invoiceId}/lineItems:search"
	if r.orgId == "" {
		return localVarReturnValue, nil, reportError("orgId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(r.orgId), -1)
	if r.invoiceId == "" {
		return localVarReturnValue, nil, reportError("invoiceId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"invoiceId"+"}", url.PathEscape(r.invoiceId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiPublicUsageDetailsQueryRequest == nil {
		return localVarReturnValue, nil, reportError("apiPublicUsageDetailsQueryRequest is required and must be specified")
	}

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
	localVarPostBody = r.apiPublicUsageDetailsQueryRequest
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
