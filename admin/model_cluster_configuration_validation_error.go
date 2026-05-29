// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ClusterConfigurationValidationError struct for ClusterConfigurationValidationError
type ClusterConfigurationValidationError struct {
	// Machine-readable error code identifying the type of validation failure.
	ErrorCode *string `json:"errorCode,omitempty"`
	// Description of the validation failure.
	ValidationIssue *string `json:"validationIssue,omitempty"`
}

// NewClusterConfigurationValidationError instantiates a new ClusterConfigurationValidationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterConfigurationValidationError() *ClusterConfigurationValidationError {
	this := ClusterConfigurationValidationError{}
	return &this
}

// NewClusterConfigurationValidationErrorWithDefaults instantiates a new ClusterConfigurationValidationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterConfigurationValidationErrorWithDefaults() *ClusterConfigurationValidationError {
	this := ClusterConfigurationValidationError{}
	return &this
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise
func (o *ClusterConfigurationValidationError) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterConfigurationValidationError) GetErrorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}

	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *ClusterConfigurationValidationError) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *ClusterConfigurationValidationError) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetValidationIssue returns the ValidationIssue field value if set, zero value otherwise
func (o *ClusterConfigurationValidationError) GetValidationIssue() string {
	if o == nil || IsNil(o.ValidationIssue) {
		var ret string
		return ret
	}
	return *o.ValidationIssue
}

// GetValidationIssueOk returns a tuple with the ValidationIssue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterConfigurationValidationError) GetValidationIssueOk() (*string, bool) {
	if o == nil || IsNil(o.ValidationIssue) {
		return nil, false
	}

	return o.ValidationIssue, true
}

// HasValidationIssue returns a boolean if a field has been set.
func (o *ClusterConfigurationValidationError) HasValidationIssue() bool {
	if o != nil && !IsNil(o.ValidationIssue) {
		return true
	}

	return false
}

// SetValidationIssue gets a reference to the given string and assigns it to the ValidationIssue field.
func (o *ClusterConfigurationValidationError) SetValidationIssue(v string) {
	o.ValidationIssue = &v
}
