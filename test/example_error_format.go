package test

import (
	"fmt"

	"go.mongodb.org/atlas-sdk/v20240805004/admin"
)

var testString = "test";
var testInt = 1
 
func ExampleFormatErrorMessageWithDetails_details_test() {
	label := "test";

	err := admin.ApiError{
		BadRequestDetail: &admin.BadRequestDetail{Fields: &[]admin.FieldViolation{
			{Description: &label, Field: &label},
		}},
		Detail:           &testString,
		Error:            &testInt,
		ErrorCode:        &testString,
		Parameters:       nil,
		Reason:           &testString,
	}
	fmt.Printf("Violations: %+v\n", admin.FormatErrorMessageWithDetails("200", "/test","POST", err))
	
	// Output:
	// Violations: POST /test: HTTP 200 (Error code: "") Detail:  Reason: . Params: [], BadRequestDetail: {"fields":[{"description":"test","field":"test"}]}
}

func ExampleFormatErrorMessageWithDetails_details_empty() {
	err := admin.ApiError{
		BadRequestDetail: &admin.BadRequestDetail{},
		Detail:           &testString,
		Error:            &testInt,
		ErrorCode:        &testString,
		Parameters:       nil,
		Reason:           &testString,
	}
	fmt.Printf("%+v\n", admin.FormatErrorMessageWithDetails("200", "/test","POST", err))

	// Output:
	// POST /test: HTTP 200 (Error code: "") Detail:  Reason: . Params: [], BadRequestDetail: {}
}

func ExampleFormatErrorMessageWithDetails_details_missing() {
	err := admin.ApiError{
		BadRequestDetail: nil,
		Detail:           &testString,
		Error:            &testInt,
		ErrorCode:        &testString,
		Parameters:       nil,
		Reason:           &testString,
	}
	fmt.Printf("%+v\n", admin.FormatErrorMessageWithDetails("200", "/test","POST", err))
	// Output:
	// POST /test: HTTP 200 (Error code: "") Detail:  Reason: . Params: [], BadRequestDetail:
}
