# Best Practices

An number of best practices for Atlas SDK users.

## Using Getters Instead of Direct Field Access

When accessing responses prefer the getter function instead of direct field access.

Instead of `response.Field` use `response.GetField()`.

Using Getters will allow seamless pointer handling and help you to avoid panic errors.
Additionally SDK provides `Set`, `IsSet` and `Unset` methods for safe field modifications.

## Check for Empty Strings for String Pointers

When model contains an pointer to the `string` 
we need to be aware that setting that field to the empty string (`""`)
will send that value to the server.

Instead of direct assignment:
```go
test := ""
requestBody.StringPointerValue = test
```

Users should always check for empty strings before assigning them:
```go
if test != ""   
    requestBody.StringPointerValue = test
```

> NOTE: if field is an string type and empty string (`""`) assignment 
will be valid.

## Use Method for Creating Models

You should prefer to use dedicated methods for creating new models.
Instead of using 

```
GroupInvitationUpdateRequest{
    ...
}
```
Prefer:
```
admin.NewGroupInvitationUpdateRequest(...)
```

## Use golangci-lint Validators

[golangci-lint](https://golangci-lint.run/) should be used to 
detect common errors with usage of Atlas SDK. 
We do not provide our own linter 

### Linting Issues

We do not recomend using `bodyclose` rule with the Atlas CLI 
as it cannot detect response body closed in the SDK. 

See: https://github.com/timakin/bodyclose/issues/39 for more informatiuon
