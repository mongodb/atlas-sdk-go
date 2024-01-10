# Best Practices

Follow these best practices when using the Atlas Go SDK.

## Using Getters Instead of Direct Field Access

When accessing responses, use the getter function instead of direct field access.

For example, use `response.GetField()` instead of `response.Field`.

Using getter functions allows seamless pointer handling and can help prevent panic errors.
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

## Working with Date Fields

In the Atlas Go SDK, the `*time.Time` type represents date fields for handling time-related data. 
When you compare values based on `time.Time`, either never compare pointers directly or do the following:

1. Avoid using direct comparison operators (e.g., `myStruct.MyDateField == ""`) to check for equality. Comparing pointers directly will check if they refer to the same memory address rather than comparing the actual date values. Since each `time.Time` instance is allocated in different memory locations, direct comparisons might yield unexpected results.

2. Use the `Has` function to check non-nil pointers:
The SDK provides a dedicated `HasFieldName` or `GetFieldName` function for each model to check if a `time.Time` pointer is non-nil before accessing its value. Always use this function to ensure that the pointer is valid before you perform any operations.

3. Use `time.Time` methods to compare date values:
When you have confirmed that the `time.Time pointer` is non-nil, you can safely use `time.Time` methods to compare the actual date values. Commonly used methods for comparison include `Before`, `After`, and `Equal`:
```go
   if !sdkModel.HasSomeDateField() {
       return;
   }
    datePtr1 := sdkModel.SomeDateField;
    if datePtr1.Before(*datePtr2) {
        // datePtr1 is before datePtr2.
    } else if datePtr1.After(*datePtr2) {
        // datePtr1 is after datePtr2.
    } else {
        // datePtr1 and datePtr2 are equal.
    }
```

## Working with Pointers

SDK pointers are utilized to denote optional values in the Go programming language:

```golang
type Data struct {
    // Represents an optional name
    Name *string `json:"results,omitempty"`
}
```

In the example above, the string value is optional, and it won't be sent to the server if not explicitly set.

## Working with Arrays

All arrays in the SDK are represented as pointers:

```golang
type Data struct {
    Results *[]DataRole `json:"results,omitempty"`
}
```

Scenarios for using pointers with arrays:

1. Update request containing an empty array (resetting the field values):

If a struct property is explicitly set to an empty array, the SDK will send an empty array request to the Atlas API.

```golang
data := Data{
    // Sending an empty array
    Results: &[]DataRole{}
}
```

2. Update request without updating the array field:

When performing an update operation, we recommend the struct property not to be present.

```golang
data := Data{
    // Sending an empty array by not setting field values (value is nil)
    // Results: &[]DataRole{}
}
```

These practices ensure accurate handling of optional values and array updates in the SDK when working with pointers in Golang.

## Working with Binary Responses

In the Atlas Go SDK, the `io.ReadCloser` type is used to return binary data using the APIs. 

- Use `io.Copy` to store on a file or pass through another steam.

- Use `io.ReadAll` to read all bytes in memory.

- Call the `.Close()` function after reading the data

Note: see example in [examples/download/downloadLogs.go](../examples/download/downloadLogs.go)

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

We don't recommend using the `bodyclose` rule as it reports an many false positives with Atlas GO SDK and other libraries.

To learn more, see [bodyclose](https://github.com/timakin/bodyclose/issues/39).
