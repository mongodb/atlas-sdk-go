// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// SearchIndexUpdateRequest struct for SearchIndexUpdateRequest
type SearchIndexUpdateRequest struct {
	Definition SearchIndexUpdateRequestDefinition `json:"definition"`
}

// NewSearchIndexUpdateRequest instantiates a new SearchIndexUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchIndexUpdateRequest(definition SearchIndexUpdateRequestDefinition) *SearchIndexUpdateRequest {
	this := SearchIndexUpdateRequest{}
	this.Definition = definition
	return &this
}

// NewSearchIndexUpdateRequestWithDefaults instantiates a new SearchIndexUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchIndexUpdateRequestWithDefaults() *SearchIndexUpdateRequest {
	this := SearchIndexUpdateRequest{}
	return &this
}

// GetDefinition returns the Definition field value
func (o *SearchIndexUpdateRequest) GetDefinition() SearchIndexUpdateRequestDefinition {
	if o == nil {
		var ret SearchIndexUpdateRequestDefinition
		return ret
	}

	return o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value
// and a boolean to check if the value has been set.
func (o *SearchIndexUpdateRequest) GetDefinitionOk() (*SearchIndexUpdateRequestDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Definition, true
}

// SetDefinition sets field value
func (o *SearchIndexUpdateRequest) SetDefinition(v SearchIndexUpdateRequestDefinition) {
	o.Definition = v
}

func (o SearchIndexUpdateRequest) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o SearchIndexUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["definition"] = o.Definition
	return toSerialize, nil
}
