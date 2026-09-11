# \InvoicesAPI

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCostExplorerProcess**](InvoicesAPI.md#CreateCostExplorerProcess) | **Post** /api/atlas/v2/orgs/{orgId}/billing/costExplorer/usage | Create One Cost Explorer Query Process
[**CreateOrgInvoiceReport**](InvoicesAPI.md#CreateOrgInvoiceReport) | **Post** /api/atlas/v2/orgs/{orgId}/invoices/{invoiceId}/reports | Create One Invoice Report
[**GetCostExplorerUsage**](InvoicesAPI.md#GetCostExplorerUsage) | **Get** /api/atlas/v2/orgs/{orgId}/billing/costExplorer/usage/{token} | Return Usage Details for One Cost Explorer Query
[**GetInvoice**](InvoicesAPI.md#GetInvoice) | **Get** /api/atlas/v2/orgs/{orgId}/invoices/{invoiceId} | Return One Invoice for One Organization
[**GetInvoiceCsv**](InvoicesAPI.md#GetInvoiceCsv) | **Get** /api/atlas/v2/orgs/{orgId}/invoices/{invoiceId}/csv | Return One Invoice as CSV for One Organization
[**GetOrgAssociatedInvoices**](InvoicesAPI.md#GetOrgAssociatedInvoices) | **Get** /api/atlas/v2/orgs/{orgId}/associatedInvoices | Return Associated Invoices
[**GetOrgInvoiceReport**](InvoicesAPI.md#GetOrgInvoiceReport) | **Get** /api/atlas/v2/orgs/{orgId}/invoices/{invoiceId}/reports/{reportId} | Return One Invoice Report
[**GetSku**](InvoicesAPI.md#GetSku) | **Get** /api/atlas/v2/skus/{skuId} | Return One Stock Keeping Unit
[**ListInvoicePending**](InvoicesAPI.md#ListInvoicePending) | **Get** /api/atlas/v2/orgs/{orgId}/invoices/pending | Return All Pending Invoices for One Organization
[**ListInvoices**](InvoicesAPI.md#ListInvoices) | **Get** /api/atlas/v2/orgs/{orgId}/invoices | Return All Invoices for One Organization
[**ListOrgInvoiceReports**](InvoicesAPI.md#ListOrgInvoiceReports) | **Get** /api/atlas/v2/orgs/{orgId}/invoices/{invoiceId}/reports | Return All Invoice Reports for One Invoice
[**ListSkus**](InvoicesAPI.md#ListSkus) | **Get** /api/atlas/v2/skus | Return All Stock Keeping Units
[**SearchInvoiceLineItems**](InvoicesAPI.md#SearchInvoiceLineItems) | **Get** /api/atlas/v2/orgs/{orgId}/invoices/{invoiceId}/lineItems:search | Return All Line Items for One Invoice by Invoice ID



## CreateCostExplorerProcess

> CostExplorerFilterResponse CreateCostExplorerProcess(ctx, orgId, costExplorerFilterRequestBody CostExplorerFilterRequestBody).Execute()

Create One Cost Explorer Query Process


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312025/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    orgId := "4888442a3354817a7320eb61" // string | 
    costExplorerFilterRequestBody := *admin.NewCostExplorerFilterRequestBody(time.Now(), time.Now()) // CostExplorerFilterRequestBody | 

    resp, r, err := sdk.InvoicesAPI.CreateCostExplorerProcess(context.Background(), orgId, &costExplorerFilterRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.CreateCostExplorerProcess`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateCostExplorerProcess`: CostExplorerFilterResponse
    fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.CreateCostExplorerProcess`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [&#x60;/orgs&#x60;](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCostExplorerProcessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **costExplorerFilterRequestBody** | [**CostExplorerFilterRequestBody**](CostExplorerFilterRequestBody.md) | Filter parameters for the Cost Explorer query. | 

### Return type

[**CostExplorerFilterResponse**](CostExplorerFilterResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrgInvoiceReport

> InvoiceReportResponse CreateOrgInvoiceReport(ctx, orgId, invoiceId, invoiceReportRequest InvoiceReportRequest).Execute()

Create One Invoice Report


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312025/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    orgId := "4888442a3354817a7320eb61" // string | 
    invoiceId := "invoiceId_example" // string | 
    invoiceReportRequest := *admin.NewInvoiceReportRequest("ReportFormat_example", "ReportType_example") // InvoiceReportRequest | 

    resp, r, err := sdk.InvoicesAPI.CreateOrgInvoiceReport(context.Background(), orgId, invoiceId, &invoiceReportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.CreateOrgInvoiceReport`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateOrgInvoiceReport`: InvoiceReportResponse
    fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.CreateOrgInvoiceReport`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [&#x60;/orgs&#x60;](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**invoiceId** | **string** | Unique string that identifies the invoice for which to generate the report. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgInvoiceReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **invoiceReportRequest** | [**InvoiceReportRequest**](InvoiceReportRequest.md) | Create Invoice Report request body. | 

### Return type

[**InvoiceReportResponse**](InvoiceReportResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2025-03-12+json
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCostExplorerUsage

> any GetCostExplorerUsage(ctx, orgId, token).Execute()

Return Usage Details for One Cost Explorer Query


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312025/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    orgId := "4888442a3354817a7320eb61" // string | 
    token := "4ABBE973862346D40F3AE859D4BE96E0F895764EB14EAB039E7B82F9D638C05C" // string | 

    resp, r, err := sdk.InvoicesAPI.GetCostExplorerUsage(context.Background(), orgId, token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.GetCostExplorerUsage`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetCostExplorerUsage`: any
    fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.GetCostExplorerUsage`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [&#x60;/orgs&#x60;](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**token** | **string** | Unique 64 digit string that identifies the Cost Explorer query. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCostExplorerUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**any**

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoice

> BillingInvoice GetInvoice(ctx, orgId, invoiceId).Execute()

Return One Invoice for One Organization


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312025/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    orgId := "4888442a3354817a7320eb61" // string | 
    invoiceId := "invoiceId_example" // string | 

    resp, r, err := sdk.InvoicesAPI.GetInvoice(context.Background(), orgId, invoiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.GetInvoice`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetInvoice`: BillingInvoice
    fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.GetInvoice`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [&#x60;/orgs&#x60;](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**invoiceId** | **string** | Unique 24-hexadecimal digit string that identifies the invoice submitted to the specified organization. Charges typically post the next day. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BillingInvoice**](BillingInvoice.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoiceCsv

> string GetInvoiceCsv(ctx, orgId, invoiceId).Execute()

Return One Invoice as CSV for One Organization


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312025/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    orgId := "4888442a3354817a7320eb61" // string | 
    invoiceId := "invoiceId_example" // string | 

    resp, r, err := sdk.InvoicesAPI.GetInvoiceCsv(context.Background(), orgId, invoiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.GetInvoiceCsv`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetInvoiceCsv`: string
    fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.GetInvoiceCsv`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [&#x60;/orgs&#x60;](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**invoiceId** | **string** | Unique 24-hexadecimal digit string that identifies the invoice submitted to the specified organization. Charges typically post the next day. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgAssociatedInvoices

> OrgAssociatedInvoiceResponse GetOrgAssociatedInvoices(ctx, orgId).Month(month).Year(year).IncludeLinkedOrgs(includeLinkedOrgs).Execute()

Return Associated Invoices


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312025/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    orgId := "4888442a3354817a7320eb61" // string | 
    month := int64(789) // int64 |  (optional)
    year := int64(789) // int64 |  (optional)
    includeLinkedOrgs := true // bool |  (optional)

    resp, r, err := sdk.InvoicesAPI.GetOrgAssociatedInvoices(context.Background(), orgId).Month(month).Year(year).IncludeLinkedOrgs(includeLinkedOrgs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.GetOrgAssociatedInvoices`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetOrgAssociatedInvoices`: OrgAssociatedInvoiceResponse
    fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.GetOrgAssociatedInvoices`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [&#x60;/orgs&#x60;](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgAssociatedInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **month** | **int64** | The month for which to retrieve invoices (1-12). Defaults to current month. | 
 **year** | **int64** | The year for which to retrieve invoices. Defaults to current year. | 
 **includeLinkedOrgs** | **bool** | Whether to include invoices from linked organizations. Defaults to false. | 

### Return type

[**OrgAssociatedInvoiceResponse**](OrgAssociatedInvoiceResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgInvoiceReport

> InvoiceReportResponse GetOrgInvoiceReport(ctx, orgId, invoiceId, reportId).Execute()

Return One Invoice Report


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312025/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    orgId := "4888442a3354817a7320eb61" // string | 
    invoiceId := "invoiceId_example" // string | 
    reportId := "reportId_example" // string | 

    resp, r, err := sdk.InvoicesAPI.GetOrgInvoiceReport(context.Background(), orgId, invoiceId, reportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.GetOrgInvoiceReport`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetOrgInvoiceReport`: InvoiceReportResponse
    fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.GetOrgInvoiceReport`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [&#x60;/orgs&#x60;](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**invoiceId** | **string** | Unique string that identifies the invoice the report was generated for. | 
**reportId** | **string** | Unique string that identifies the report to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgInvoiceReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**InvoiceReportResponse**](InvoiceReportResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSku

> SkuResponse GetSku(ctx, skuId).Execute()

Return One Stock Keeping Unit


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312025/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    skuId := "ATLAS_AWS_INSTANCE_M10" // string | 

    resp, r, err := sdk.InvoicesAPI.GetSku(context.Background(), skuId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.GetSku`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetSku`: SkuResponse
    fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.GetSku`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**skuId** | **string** | Unique identifier of the SKU to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSkuRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SkuResponse**](SkuResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInvoicePending

> PaginatedApiInvoice ListInvoicePending(ctx, orgId).Execute()

Return All Pending Invoices for One Organization


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312025/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    orgId := "4888442a3354817a7320eb61" // string | 

    resp, r, err := sdk.InvoicesAPI.ListInvoicePending(context.Background(), orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.ListInvoicePending`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListInvoicePending`: PaginatedApiInvoice
    fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.ListInvoicePending`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [&#x60;/orgs&#x60;](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListInvoicePendingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaginatedApiInvoice**](PaginatedApiInvoice.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInvoices

> PaginatedApiInvoiceMetadata ListInvoices(ctx, orgId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).ViewLinkedInvoices(viewLinkedInvoices).StatusNames(statusNames).FromDate(fromDate).ToDate(toDate).SortBy(sortBy).OrderBy(orderBy).Execute()

Return All Invoices for One Organization


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312025/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    orgId := "4888442a3354817a7320eb61" // string | 
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)
    viewLinkedInvoices := true // bool |  (optional) (default to true)
    statusNames := []string{"Inner_example"} // []string |  (optional)
    fromDate := time.Now() // string |  (optional)
    toDate := time.Now() // string |  (optional)
    sortBy := "sortBy_example" // string |  (optional) (default to "END_DATE")
    orderBy := "desc" // string |  (optional) (default to "desc")

    resp, r, err := sdk.InvoicesAPI.ListInvoices(context.Background(), orgId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).ViewLinkedInvoices(viewLinkedInvoices).StatusNames(statusNames).FromDate(fromDate).ToDate(toDate).SortBy(sortBy).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.ListInvoices`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListInvoices`: PaginatedApiInvoiceMetadata
    fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.ListInvoices`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [&#x60;/orgs&#x60;](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (&#x60;totalCount&#x60;) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **viewLinkedInvoices** | **bool** | Flag that indicates whether to return linked invoices in the &#x60;linkedInvoices&#x60; field. | [default to true]
 **statusNames** | **[]string** | Statuses of the invoice to be retrieved. Omit to return invoices of all statuses. | 
 **fromDate** | **string** | Retrieve the invoices the &#x60;startDates&#x60; of which are greater than or equal to the &#x60;fromDate&#x60;. If omit, the invoices return will go back to earliest &#x60;startDate&#x60;. | 
 **toDate** | **string** | Retrieve the invoices the &#x60;endDates&#x60; of which are smaller than or equal to the &#x60;toDate&#x60;. If omit, the invoices return will go further to latest &#x60;endDate&#x60;. | 
 **sortBy** | **string** | Field used to sort the returned invoices by. Use in combination with &#x60;orderBy&#x60; parameter to control the order of the result. | [default to &quot;END_DATE&quot;]
 **orderBy** | **string** | Field used to order the returned invoices by. Use in combination of &#x60;sortBy&#x60; parameter to control the order of the result. | [default to &quot;desc&quot;]

### Return type

[**PaginatedApiInvoiceMetadata**](PaginatedApiInvoiceMetadata.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgInvoiceReports

> PaginatedInvoiceReport ListOrgInvoiceReports(ctx, orgId, invoiceId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Invoice Reports for One Invoice


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312025/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    orgId := "4888442a3354817a7320eb61" // string | 
    invoiceId := "invoiceId_example" // string | 
    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)

    resp, r, err := sdk.InvoicesAPI.ListOrgInvoiceReports(context.Background(), orgId, invoiceId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.ListOrgInvoiceReports`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListOrgInvoiceReports`: PaginatedInvoiceReport
    fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.ListOrgInvoiceReports`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [&#x60;/orgs&#x60;](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**invoiceId** | **string** | Unique string that identifies the invoice to list reports for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgInvoiceReportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (&#x60;totalCount&#x60;) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedInvoiceReport**](PaginatedInvoiceReport.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSkus

> PaginatedApiSKU ListSkus(ctx).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Stock Keeping Units


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312025/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    includeCount := true // bool |  (optional) (default to true)
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)

    resp, r, err := sdk.InvoicesAPI.ListSkus(context.Background()).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.ListSkus`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListSkus`: PaginatedApiSKU
    fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.ListSkus`: %v (%v)\n", resp, r)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSkusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (&#x60;totalCount&#x60;) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedApiSKU**](PaginatedApiSKU.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2025-03-12+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchInvoiceLineItems

> PaginatedPublicApiUsageDetailsLineItem SearchInvoiceLineItems(ctx, orgId, invoiceId, apiPublicUsageDetailsQueryRequest ApiPublicUsageDetailsQueryRequest).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Return All Line Items for One Invoice by Invoice ID


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20250312025/admin"
)

func main() {
    apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
    apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")

    sdk, err := admin.NewClient(admin.UseDigestAuth(apiKey, apiSecret))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error initializing SDK: %v\n", err)
        return
    }

    orgId := "4888442a3354817a7320eb61" // string | 
    invoiceId := "invoiceId_example" // string | 
    apiPublicUsageDetailsQueryRequest := *admin.NewApiPublicUsageDetailsQueryRequest() // ApiPublicUsageDetailsQueryRequest | 
    itemsPerPage := int(56) // int |  (optional) (default to 100)
    pageNum := int(56) // int |  (optional) (default to 1)

    resp, r, err := sdk.InvoicesAPI.SearchInvoiceLineItems(context.Background(), orgId, invoiceId, &apiPublicUsageDetailsQueryRequest).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesAPI.SearchInvoiceLineItems`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `SearchInvoiceLineItems`: PaginatedPublicApiUsageDetailsLineItem
    fmt.Fprintf(os.Stdout, "Response from `InvoicesAPI.SearchInvoiceLineItems`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [&#x60;/orgs&#x60;](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**invoiceId** | **string** | Unique 24-hexadecimal digit string that identifies the invoice submitted to the specified organization. Charges typically post the next day. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchInvoiceLineItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiPublicUsageDetailsQueryRequest** | [**ApiPublicUsageDetailsQueryRequest**](ApiPublicUsageDetailsQueryRequest.md) | Filter parameters for the &#x60;lineItems&#x60; query. Send a request with an empty JSON body to retrieve all line items for a given &#x60;invoiceID&#x60; without applying any filters. | 
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedPublicApiUsageDetailsLineItem**](PaginatedPublicApiUsageDetailsLineItem.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2024-08-05+json
- **Accept**: application/vnd.atlas.2024-08-05+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

