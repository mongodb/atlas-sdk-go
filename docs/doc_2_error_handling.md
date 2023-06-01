
# Error Handling

You can obtain detailed information about errors returned by the API when you use the Atlas Go SDK. Use the error code to determine the cause of the error. To learn more about API error codes, see [Atlas Administration API Error Codes](https://www.mongodb.com/docs/atlas/reference/api-errors/).

Errors are represented by [ErrorObject](https://github.com/mongodb/atlas-sdk-go/blob/main/admin/model_error.go).

## Fetching Error Object

To fetch the error object, execute the following:

```go
import "go.mongodb.org/atlas-sdk/admin"

projects, response, err := admin.ProjectsApi.ListProjects(ctx).Execute()
apiError := admin.AsError(err)
fmt.Println(apiError)
```

## Checking for the Existence of a Specific Error Code

To check for the existence of a specific error code, execute the following:

```go
import admin "go.mongodb.org/atlas-sdk/admin"

projects, response, err := admin.ProjectsApi.ListProjects(ctx).Execute()
if admin.IsErrorCode(err, "code"){
 // Do something
}
```

