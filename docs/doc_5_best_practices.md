# Best Practices

Follow these best practices when using the Atlas Go SDK.

## Using Getters Instead of Direct Field Access

When accessing responses, use the getter function instead of direct field access.

For example, use `response.GetField()` instead of `response.Field`.

Using geter functions allows seamless pointer handling and can help prevent panic errors.
Additionally, the Atlas Go SDK provides `Set`, `IsSet` and `Unset` methods for safe field modifications.

## Check for Empty Strings for String Pointers

When a model contains a pointer to a `string`, the Atlas Go SDK sends that value to the server,
even if it is set to an empty string (`""`).

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


## Use Method for Creating Models

Use dedicated methods for creating new models.

For example, instead of using the following:

```
GroupInvitationUpdateRequest{
    ...
}
```

Use the following dedicated method:
```
admin.NewGroupInvitationUpdateRequest(...)
```

## Use golangci-lint Validators

Use [golangci-lint](https://golangci-lint.run/) to detect common errors when using the Atlas Go SDK. 
The Atlas Go SDK doesn't provide its own linter.

### Linting Issues

We don't recommend using the `bodyclose` rule with the Atlas CLI 
as it can't detect the response body close when using the Atlas Go SDK. 

To learn more, see [bodyclose](https://github.com/timakin/bodyclose/issues/39).
