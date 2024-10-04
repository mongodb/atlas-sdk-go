package test

import (
	"fmt"

	"go.mongodb.org/atlas-sdk/v20240805004/admin"
)

 
func ExampleFormatErrorMessageWithDetails_test() {
	label := "test";
	
	err := admin.ApiError{
		BadRequestDetail: &admin.BadRequestDetail{Fields: &[]admin.FieldViolation{
			{Description: &label, Field: &label},
		}},
		Detail:           new(string),
		Error:            new(int),
		ErrorCode:        new(string),
		Parameters:       &[]interface{}{},
		Reason:           new(string),
	}
	fmt.Printf("Violations: %+v\n", admin.FormatErrorMessageWithDetails("200", "/test","POST", err))
	
	// Output:
	// Violations: POST /test: HTTP 200 (Error code: "") Detail:  Reason: . Params: [], BadRequestDetail: {"fields":[{"description":"test","field":"test"}]}
}

func ExampleFormatErrorMessageWithDetails_empty() {
	err := admin.ApiError{
		BadRequestDetail: &admin.BadRequestDetail{},
		Detail:           new(string),
		Error:            new(int),
		ErrorCode:        new(string),
		Parameters:       &[]interface{}{},
		Reason:           new(string),
	}
	fmt.Printf("%+v\n", admin.FormatErrorMessageWithDetails("200", "/test","POST", err))

	// Output:
	// POST /test: HTTP 200 (Error code: "") Detail:  Reason: . Params: [], BadRequestDetail: {}
}

func ExampleFormatErrorMessageWithDetails_missing() {
	err := admin.ApiError{
		BadRequestDetail: nil,
		Detail:           new(string),
		Error:            new(int),
		ErrorCode:        new(string),
		Parameters:       &[]interface{}{},
		Reason:           new(string),
	}
	fmt.Printf("%+v\n", admin.FormatErrorMessageWithDetails("200", "/test","POST", err))
	// Output:
	// POST /test: HTTP 200 (Error code: "") Detail:  Reason: . Params: [], BadRequestDetail:
}
