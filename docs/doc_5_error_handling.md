
# Error Handling

SDK enables users to obtain detailed information about errors returned from backend.
Errors are represented by [ErrorObject](./admin/model_error.go).
Users should rely on the error code for detection of specific error cases.

## Fetching Error Object
```go
import "go.mongodb.org/atlas-sdk/admin"

projects, response, err := admin.ProjectsApi.ListProjects(ctx).Execute()
apiError := admin.AsError(err)
fmt.Println(apiError)
```

## Checking for existence of specific error code
```go
import admin "go.mongodb.org/atlas-sdk/admin"

projects, response, err := admin.ProjectsApi.ListProjects(ctx).Execute()
if admin.IsErrorCode(err, "code"){
 // Do something
}
```

