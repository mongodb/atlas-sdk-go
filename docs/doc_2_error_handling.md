# Error Handling

You can obtain detailed information about errors returned by the API when you use the Atlas Go SDK. Use the error code to determine the cause of the error. To learn more about API error codes, see [Atlas Administration API Error Codes](https://www.mongodb.com/docs/atlas/reference/api-errors/).

Errors are represented by [ApiErrorObject](https://github.com/mongodb/atlas-sdk-go/blob/main/admin/model_api_error.go).

## Fetching Error Object

To fetch the error object, execute the following:

```go
// Surrounding code omitted for brevity

projects, response, err := admin.ProjectsApi.ListProjects(ctx).Execute()
apiError, ok := admin.AsError(err)
fmt.Println(apiError)
```

## Checking for the Existence of a Specific Error Code

To check for the existence of a specific error code (e.g. `MAXIMUM_INDEXES_FOR_TENANT_EXCEEDED`), execute the following:

```go
// Surrounding code omitted for brevity

projects, response, err := admin.ProjectsApi.ListProjects(ctx).Execute()
if admin.IsErrorCode(err, "code"){
 // Do something
}
```

## Checking for the Existence of a Specific Response Status Code

To check for the existence of a specific HTTP response error code, execute the following:

```go
// Surrounding code omitted for brevity

projects, response, err := admin.ProjectsApi.ListProjects(ctx).Execute()
apiError, ok := admin.AsError(err)
if ok && apiError.GetError() == 404 {
 // Do something
}
```

## Mocking Errors

SDK errors can be mocked by creating instance of `GenericOpenAPIError` struct. 

```go
// Surrounding code omitted for brevity

apiError := admin.GenericOpenAPIError{}
apiError.SetModel(admin.ApiError{
    Detail: admin.PtrString("Error when listing clusters"),
    Error: admin.PtrInt(400),
    ErrorCode: admin.PtrString("CLUSTERS_UNREACHABLE"),
    Reason: admin.PtrString("Clusters unreachable"),
    
})
apiError.SetError("Mocked error")
```

Struct can be passed as `error` for the all SDK mocked methods. 
