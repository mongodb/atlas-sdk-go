// Code based on the AtlasAPI V2 OpenAPI file

package admin

// ClusterConfigurationValidationResult struct for ClusterConfigurationValidationResult
type ClusterConfigurationValidationResult struct {
	// List of validation errors, present only when valid is false.
	Errors *[]ClusterConfigurationValidationError `json:"errors,omitempty"`
	// Whether the cluster configuration is valid.
	Valid *bool `json:"valid,omitempty"`
}

// NewClusterConfigurationValidationResult instantiates a new ClusterConfigurationValidationResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterConfigurationValidationResult() *ClusterConfigurationValidationResult {
	this := ClusterConfigurationValidationResult{}
	return &this
}

// NewClusterConfigurationValidationResultWithDefaults instantiates a new ClusterConfigurationValidationResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterConfigurationValidationResultWithDefaults() *ClusterConfigurationValidationResult {
	this := ClusterConfigurationValidationResult{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise
func (o *ClusterConfigurationValidationResult) GetErrors() []ClusterConfigurationValidationError {
	if o == nil || IsNil(o.Errors) {
		var ret []ClusterConfigurationValidationError
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterConfigurationValidationResult) GetErrorsOk() (*[]ClusterConfigurationValidationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}

	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ClusterConfigurationValidationResult) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []ClusterConfigurationValidationError and assigns it to the Errors field.
func (o *ClusterConfigurationValidationResult) SetErrors(v []ClusterConfigurationValidationError) {
	o.Errors = &v
}

// GetValid returns the Valid field value if set, zero value otherwise
func (o *ClusterConfigurationValidationResult) GetValid() bool {
	if o == nil || IsNil(o.Valid) {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterConfigurationValidationResult) GetValidOk() (*bool, bool) {
	if o == nil || IsNil(o.Valid) {
		return nil, false
	}

	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *ClusterConfigurationValidationResult) HasValid() bool {
	if o != nil && !IsNil(o.Valid) {
		return true
	}

	return false
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *ClusterConfigurationValidationResult) SetValid(v bool) {
	o.Valid = &v
}
