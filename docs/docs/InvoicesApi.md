# \InvoicesApi

All URIs are relative to *https://cloud.mongodb.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCostExplorerQueryProcess**](InvoicesApi.md#CreateCostExplorerQueryProcess) | **Post** /api/atlas/v2/orgs/{orgId}/billing/costExplorer/usage | Create Cost Explorer query process
[**CreateCostExplorerQueryProcess1**](InvoicesApi.md#CreateCostExplorerQueryProcess1) | **Get** /api/atlas/v2/orgs/{orgId}/billing/costExplorer/usage/{token} | Return results from a given Cost Explorer query, or notify that the results are not ready yet.
[**DownloadInvoiceCSV**](InvoicesApi.md#DownloadInvoiceCSV) | **Get** /api/atlas/v2/orgs/{orgId}/invoices/{invoiceId}/csv | Return One Organization Invoice as CSV
[**GetInvoice**](InvoicesApi.md#GetInvoice) | **Get** /api/atlas/v2/orgs/{orgId}/invoices/{invoiceId} | Return One Organization Invoice
[**ListInvoices**](InvoicesApi.md#ListInvoices) | **Get** /api/atlas/v2/orgs/{orgId}/invoices | Return All Invoices for One Organization
[**ListPendingInvoices**](InvoicesApi.md#ListPendingInvoices) | **Get** /api/atlas/v2/orgs/{orgId}/invoices/pending | Return All Pending Invoices for One Organization
[**QueryLineItemsFromSingleInvoice**](InvoicesApi.md#QueryLineItemsFromSingleInvoice) | **Get** /api/atlas/v2/orgs/{orgId}/invoices/{invoiceId}/lineItems/:search | Query lineItems of the specified invoiceId



## CreateCostExplorerQueryProcess

> CostExplorerFilterResponse CreateCostExplorerQueryProcess(ctx, orgId, costExplorerFilterRequestBody CostExplorerFilterRequestBody).Execute()

Create Cost Explorer query process


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805002/admin"
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
    costExplorerFilterRequestBody := *openapiclient.NewCostExplorerFilterRequestBody(time.Now(), time.Now()) // CostExplorerFilterRequestBody | 

    resp, r, err := sdk.InvoicesApi.CreateCostExplorerQueryProcess(context.Background(), orgId, &costExplorerFilterRequestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesApi.CreateCostExplorerQueryProcess`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateCostExplorerQueryProcess`: CostExplorerFilterResponse
    fmt.Fprintf(os.Stdout, "Response from `InvoicesApi.CreateCostExplorerQueryProcess`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCostExplorerQueryProcessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **costExplorerFilterRequestBody** | [**CostExplorerFilterRequestBody**](CostExplorerFilterRequestBody.md) | Filter parameters for the Cost Explorer query. | 

### Return type

[**CostExplorerFilterResponse**](CostExplorerFilterResponse.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2023-01-01+json
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCostExplorerQueryProcess1

> string CreateCostExplorerQueryProcess1(ctx, orgId, token).Execute()

Return results from a given Cost Explorer query, or notify that the results are not ready yet.


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805002/admin"
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

    resp, r, err := sdk.InvoicesApi.CreateCostExplorerQueryProcess1(context.Background(), orgId, token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesApi.CreateCostExplorerQueryProcess1`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `CreateCostExplorerQueryProcess1`: string
    fmt.Fprintf(os.Stdout, "Response from `InvoicesApi.CreateCostExplorerQueryProcess1`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**token** | **string** | Unique 64 digit string that identifies the Cost Explorer query. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCostExplorerQueryProcess1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+csv, application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadInvoiceCSV

> string DownloadInvoiceCSV(ctx, orgId, invoiceId).Execute()

Return One Organization Invoice as CSV


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805002/admin"
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

    resp, r, err := sdk.InvoicesApi.DownloadInvoiceCSV(context.Background(), orgId, invoiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesApi.DownloadInvoiceCSV`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `DownloadInvoiceCSV`: string
    fmt.Fprintf(os.Stdout, "Response from `InvoicesApi.DownloadInvoiceCSV`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**invoiceId** | **string** | Unique 24-hexadecimal digit string that identifies the invoice submitted to the specified organization. Charges typically post the next day. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadInvoiceCSVRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+csv, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoice

> string GetInvoice(ctx, orgId, invoiceId).Execute()

Return One Organization Invoice


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805002/admin"
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

    resp, r, err := sdk.InvoicesApi.GetInvoice(context.Background(), orgId, invoiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesApi.GetInvoice`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `GetInvoice`: string
    fmt.Fprintf(os.Stdout, "Response from `InvoicesApi.GetInvoice`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**invoiceId** | **string** | Unique 24-hexadecimal digit string that identifies the invoice submitted to the specified organization. Charges typically post the next day. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+csv, application/vnd.atlas.2023-01-01+json, application/json

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

    "go.mongodb.org/atlas-sdk/v20240805002/admin"
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
    itemsPerPage := int(100) // int |  (optional) (default to 100)
    pageNum := int(1) // int |  (optional) (default to 1)
    viewLinkedInvoices := true // bool |  (optional) (default to true)
    statusNames := []string{"Inner_example"} // []string |  (optional)
    fromDate := time.Now() // string |  (optional)
    toDate := time.Now() // string |  (optional)
    sortBy := "sortBy_example" // string |  (optional) (default to "END_DATE")
    orderBy := "desc" // string |  (optional) (default to "desc")

    resp, r, err := sdk.InvoicesApi.ListInvoices(context.Background(), orgId).IncludeCount(includeCount).ItemsPerPage(itemsPerPage).PageNum(pageNum).ViewLinkedInvoices(viewLinkedInvoices).StatusNames(statusNames).FromDate(fromDate).ToDate(toDate).SortBy(sortBy).OrderBy(orderBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesApi.ListInvoices`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListInvoices`: PaginatedApiInvoiceMetadata
    fmt.Fprintf(os.Stdout, "Response from `InvoicesApi.ListInvoices`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeCount** | **bool** | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. | [default to true]
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]
 **viewLinkedInvoices** | **bool** | Flag that indicates whether to return linked invoices in the linkedInvoices field. | [default to true]
 **statusNames** | **[]string** | Statuses of the invoice to be retrieved. Omit to return invoices of all statuses. | 
 **fromDate** | **string** | Retrieve the invoices the startDates of which are greater than or equal to the fromDate. If omit, the invoices return will go back to earliest startDate. | 
 **toDate** | **string** | Retrieve the invoices the endDates of which are smaller than or equal to the toDate. If omit, the invoices return will go further to latest endDate. | 
 **sortBy** | **string** | Field used to sort the returned invoices by. Use in combination with orderBy parameter to control the order of the result. | [default to &quot;END_DATE&quot;]
 **orderBy** | **string** | Field used to order the returned invoices by. Use in combination of sortBy parameter to control the order of the result. | [default to &quot;desc&quot;]

### Return type

[**PaginatedApiInvoiceMetadata**](PaginatedApiInvoiceMetadata.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPendingInvoices

> PaginatedApiInvoice ListPendingInvoices(ctx, orgId).Execute()

Return All Pending Invoices for One Organization


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805002/admin"
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

    resp, r, err := sdk.InvoicesApi.ListPendingInvoices(context.Background(), orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesApi.ListPendingInvoices`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `ListPendingInvoices`: PaginatedApiInvoice
    fmt.Fprintf(os.Stdout, "Response from `InvoicesApi.ListPendingInvoices`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPendingInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PaginatedApiInvoice**](PaginatedApiInvoice.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.atlas.2023-01-01+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryLineItemsFromSingleInvoice

> PaginatedPublicApiUsageDetailsLineItem QueryLineItemsFromSingleInvoice(ctx, orgId, invoiceId, apiPublicUsageDetailsQueryRequest ApiPublicUsageDetailsQueryRequest).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()

Query lineItems of the specified invoiceId


### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    "go.mongodb.org/atlas-sdk/v20240805002/admin"
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
    apiPublicUsageDetailsQueryRequest := *openapiclient.NewApiPublicUsageDetailsQueryRequest() // ApiPublicUsageDetailsQueryRequest | 
    itemsPerPage := int(100) // int |  (optional) (default to 100)
    pageNum := int(1) // int |  (optional) (default to 1)

    resp, r, err := sdk.InvoicesApi.QueryLineItemsFromSingleInvoice(context.Background(), orgId, invoiceId, &apiPublicUsageDetailsQueryRequest).ItemsPerPage(itemsPerPage).PageNum(pageNum).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InvoicesApi.QueryLineItemsFromSingleInvoice`: %v (%v)\n", err, r)
        apiError, ok := admin.AsError(err)
        if ok {
            fmt.Fprintf(os.Stderr, "API error obj: %v\n", apiError)
        }
        return
    }
    // response from `QueryLineItemsFromSingleInvoice`: PaginatedPublicApiUsageDetailsLineItem
    fmt.Fprintf(os.Stdout, "Response from `InvoicesApi.QueryLineItemsFromSingleInvoice`: %v (%v)\n", resp, r)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access. | 
**invoiceId** | **string** | Unique 24-hexadecimal digit string that identifies the invoice submitted to the specified organization. Charges typically post the next day. | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryLineItemsFromSingleInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiPublicUsageDetailsQueryRequest** | [**ApiPublicUsageDetailsQueryRequest**](ApiPublicUsageDetailsQueryRequest.md) | Filter parameters for the lineItems query. Send a request with an empty JSON body to retrieve all line items for a given invoiceID without applying any filters. | 
 **itemsPerPage** | **int** | Number of items that the response returns per page. | [default to 100]
 **pageNum** | **int** | Number of the page that displays the current set of the total objects that the response returns. | [default to 1]

### Return type

[**PaginatedPublicApiUsageDetailsLineItem**](PaginatedPublicApiUsageDetailsLineItem.md)

### Authorization
[DigestAuth](../README.md#Authentication)

### HTTP request headers

- **Content-Type**: application/vnd.atlas.2024-08-05+json
- **Accept**: application/vnd.atlas.2024-08-05+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

