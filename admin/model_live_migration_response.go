// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the LiveMigrationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LiveMigrationResponse{}

// LiveMigrationResponse struct for LiveMigrationResponse
type LiveMigrationResponse struct {
	// Unique 24-hexadecimal digit string that identifies the migration job.
	Id *string `json:"_id,omitempty"`
	// Replication lag between the source and destination clusters. Atlas returns this setting only during an active migration, before the cutover phase.
	LagTimeSeconds NullableInt64 `json:"lagTimeSeconds,omitempty"`
	// List of hosts running MongoDB Agents. These Agents can transfer your MongoDB data between one source and one target cluster.
	MigrationHosts []string `json:"migrationHosts,omitempty"`
	// Flag that indicates the migrated cluster can be cut over to MongoDB Atlas.
	ReadyForCutover *bool `json:"readyForCutover,omitempty"`
	// Progress made in migrating one cluster to MongoDB Atlas.  | Status   | Explanation | |----------|-------------| | NEW      | Someone scheduled a local cluster migration to MongoDB Atlas. | | FAILED   | The cluster migration to MongoDB Atlas failed.                | | COMPLETE | The cluster migration to MongoDB Atlas succeeded.             | | EXPIRED  | MongoDB Atlas prepares to begin the cut over of the migrating cluster when source and target clusters have almost synchronized. If `\"readyForCutover\" : true`, this synchronization starts a timer of 120 hours. You can extend this timer. If the timer expires, MongoDB Atlas returns this status. | | WORKING  | The cluster migration to MongoDB Atlas is performing one of the following tasks:<ul><li>Preparing connections to source and target clusters</li><li>Replicating data from source to target</li><li>Verifying MongoDB Atlas connection settings</li><li>Stopping replication after the cut over</li></ul> |
	Status *string `json:"status,omitempty"`
}

// NewLiveMigrationResponse instantiates a new LiveMigrationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiveMigrationResponse() *LiveMigrationResponse {
	this := LiveMigrationResponse{}
	return &this
}

// NewLiveMigrationResponseWithDefaults instantiates a new LiveMigrationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiveMigrationResponseWithDefaults() *LiveMigrationResponse {
	this := LiveMigrationResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LiveMigrationResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveMigrationResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LiveMigrationResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LiveMigrationResponse) SetId(v string) {
	o.Id = &v
}

// GetLagTimeSeconds returns the LagTimeSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LiveMigrationResponse) GetLagTimeSeconds() int64 {
	if o == nil || IsNil(o.LagTimeSeconds.Get()) {
		var ret int64
		return ret
	}
	return *o.LagTimeSeconds.Get()
}

// GetLagTimeSecondsOk returns a tuple with the LagTimeSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LiveMigrationResponse) GetLagTimeSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LagTimeSeconds.Get(), o.LagTimeSeconds.IsSet()
}

// HasLagTimeSeconds returns a boolean if a field has been set.
func (o *LiveMigrationResponse) HasLagTimeSeconds() bool {
	if o != nil && o.LagTimeSeconds.IsSet() {
		return true
	}

	return false
}

// SetLagTimeSeconds gets a reference to the given NullableInt64 and assigns it to the LagTimeSeconds field.
func (o *LiveMigrationResponse) SetLagTimeSeconds(v int64) {
	o.LagTimeSeconds.Set(&v)
}

// SetLagTimeSecondsNil sets the value for LagTimeSeconds to be an explicit nil
func (o *LiveMigrationResponse) SetLagTimeSecondsNil() {
	o.LagTimeSeconds.Set(nil)
}

// UnsetLagTimeSeconds ensures that no value is present for LagTimeSeconds, not even an explicit nil
func (o *LiveMigrationResponse) UnsetLagTimeSeconds() {
	o.LagTimeSeconds.Unset()
}

// GetMigrationHosts returns the MigrationHosts field value if set, zero value otherwise.
func (o *LiveMigrationResponse) GetMigrationHosts() []string {
	if o == nil || IsNil(o.MigrationHosts) {
		var ret []string
		return ret
	}
	return o.MigrationHosts
}

// GetMigrationHostsOk returns a tuple with the MigrationHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveMigrationResponse) GetMigrationHostsOk() ([]string, bool) {
	if o == nil || IsNil(o.MigrationHosts) {
		return nil, false
	}
	return o.MigrationHosts, true
}

// HasMigrationHosts returns a boolean if a field has been set.
func (o *LiveMigrationResponse) HasMigrationHosts() bool {
	if o != nil && !IsNil(o.MigrationHosts) {
		return true
	}

	return false
}

// SetMigrationHosts gets a reference to the given []string and assigns it to the MigrationHosts field.
func (o *LiveMigrationResponse) SetMigrationHosts(v []string) {
	o.MigrationHosts = v
}

// GetReadyForCutover returns the ReadyForCutover field value if set, zero value otherwise.
func (o *LiveMigrationResponse) GetReadyForCutover() bool {
	if o == nil || IsNil(o.ReadyForCutover) {
		var ret bool
		return ret
	}
	return *o.ReadyForCutover
}

// GetReadyForCutoverOk returns a tuple with the ReadyForCutover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveMigrationResponse) GetReadyForCutoverOk() (*bool, bool) {
	if o == nil || IsNil(o.ReadyForCutover) {
		return nil, false
	}
	return o.ReadyForCutover, true
}

// HasReadyForCutover returns a boolean if a field has been set.
func (o *LiveMigrationResponse) HasReadyForCutover() bool {
	if o != nil && !IsNil(o.ReadyForCutover) {
		return true
	}

	return false
}

// SetReadyForCutover gets a reference to the given bool and assigns it to the ReadyForCutover field.
func (o *LiveMigrationResponse) SetReadyForCutover(v bool) {
	o.ReadyForCutover = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LiveMigrationResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LiveMigrationResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LiveMigrationResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *LiveMigrationResponse) SetStatus(v string) {
	o.Status = &v
}

func (o LiveMigrationResponse) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o LiveMigrationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.LagTimeSeconds.IsSet() {
		toSerialize["lagTimeSeconds"] = o.LagTimeSeconds.Get()
	}
	return toSerialize, nil
}

type NullableLiveMigrationResponse struct {
	value *LiveMigrationResponse
	isSet bool
}

func (v NullableLiveMigrationResponse) Get() *LiveMigrationResponse {
	return v.value
}

func (v *NullableLiveMigrationResponse) Set(val *LiveMigrationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLiveMigrationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLiveMigrationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiveMigrationResponse(val *LiveMigrationResponse) *NullableLiveMigrationResponse {
	return &NullableLiveMigrationResponse{value: val, isSet: true}
}

func (v NullableLiveMigrationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiveMigrationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
