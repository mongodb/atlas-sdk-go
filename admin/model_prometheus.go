// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// checks if the Prometheus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Prometheus{}

// Prometheus Details to integrate one Prometheus account with one MongoDB Cloud project.
type Prometheus struct {
	// Flag that indicates whether someone has activated the Prometheus integration.
	Enabled bool `json:"enabled"`
	// Combination of IPv4 address and Internet Assigned Numbers Authority (IANA) port or the IANA port alone to which Prometheus binds to ingest MongoDB metrics.
	ListenAddress     *string `json:"listenAddress,omitempty"`
	Password          *string `json:"password,omitempty"`
	RateLimitInterval *int    `json:"rateLimitInterval,omitempty"`
	// Security Scheme to apply to HyperText Transfer Protocol (HTTP) traffic between Prometheus and MongoDB Cloud.
	Scheme string `json:"scheme"`
	// Desired method to discover the Prometheus service.
	ServiceDiscovery string `json:"serviceDiscovery"`
	// Root-relative path to the Transport Layer Security (TLS) Privacy Enhanced Mail (PEM) key and certificate file on the host.
	TlsPemPath *string `json:"tlsPemPath,omitempty"`
	// Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud. The value must match the third-party service integration type.
	Type *string `json:"type,omitempty"`
	// Human-readable label that identifies your Prometheus incoming webhook.
	Username string `json:"username"`
}

// NewPrometheus instantiates a new Prometheus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrometheus(enabled bool, scheme string, serviceDiscovery string, username string) *Prometheus {
	this := Prometheus{}
	this.Enabled = enabled
	var listenAddress string = ":9216"
	this.ListenAddress = &listenAddress
	this.Scheme = scheme
	this.ServiceDiscovery = serviceDiscovery
	this.Username = username
	return &this
}

// NewPrometheusWithDefaults instantiates a new Prometheus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrometheusWithDefaults() *Prometheus {
	this := Prometheus{}
	var listenAddress string = ":9216"
	this.ListenAddress = &listenAddress
	return &this
}

// GetEnabled returns the Enabled field value
func (o *Prometheus) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *Prometheus) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *Prometheus) SetEnabled(v bool) {
	o.Enabled = v
}

// GetListenAddress returns the ListenAddress field value if set, zero value otherwise.
func (o *Prometheus) GetListenAddress() string {
	if o == nil || IsNil(o.ListenAddress) {
		var ret string
		return ret
	}
	return *o.ListenAddress
}

// GetListenAddressOk returns a tuple with the ListenAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prometheus) GetListenAddressOk() (*string, bool) {
	if o == nil || IsNil(o.ListenAddress) {
		return nil, false
	}
	return o.ListenAddress, true
}

// HasListenAddress returns a boolean if a field has been set.
func (o *Prometheus) HasListenAddress() bool {
	if o != nil && !IsNil(o.ListenAddress) {
		return true
	}

	return false
}

// SetListenAddress gets a reference to the given string and assigns it to the ListenAddress field.
func (o *Prometheus) SetListenAddress(v string) {
	o.ListenAddress = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *Prometheus) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prometheus) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *Prometheus) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *Prometheus) SetPassword(v string) {
	o.Password = &v
}

// GetRateLimitInterval returns the RateLimitInterval field value if set, zero value otherwise.
func (o *Prometheus) GetRateLimitInterval() int {
	if o == nil || IsNil(o.RateLimitInterval) {
		var ret int
		return ret
	}
	return *o.RateLimitInterval
}

// GetRateLimitIntervalOk returns a tuple with the RateLimitInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prometheus) GetRateLimitIntervalOk() (*int, bool) {
	if o == nil || IsNil(o.RateLimitInterval) {
		return nil, false
	}
	return o.RateLimitInterval, true
}

// HasRateLimitInterval returns a boolean if a field has been set.
func (o *Prometheus) HasRateLimitInterval() bool {
	if o != nil && !IsNil(o.RateLimitInterval) {
		return true
	}

	return false
}

// SetRateLimitInterval gets a reference to the given int and assigns it to the RateLimitInterval field.
func (o *Prometheus) SetRateLimitInterval(v int) {
	o.RateLimitInterval = &v
}

// GetScheme returns the Scheme field value
func (o *Prometheus) GetScheme() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scheme
}

// GetSchemeOk returns a tuple with the Scheme field value
// and a boolean to check if the value has been set.
func (o *Prometheus) GetSchemeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scheme, true
}

// SetScheme sets field value
func (o *Prometheus) SetScheme(v string) {
	o.Scheme = v
}

// GetServiceDiscovery returns the ServiceDiscovery field value
func (o *Prometheus) GetServiceDiscovery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceDiscovery
}

// GetServiceDiscoveryOk returns a tuple with the ServiceDiscovery field value
// and a boolean to check if the value has been set.
func (o *Prometheus) GetServiceDiscoveryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceDiscovery, true
}

// SetServiceDiscovery sets field value
func (o *Prometheus) SetServiceDiscovery(v string) {
	o.ServiceDiscovery = v
}

// GetTlsPemPath returns the TlsPemPath field value if set, zero value otherwise.
func (o *Prometheus) GetTlsPemPath() string {
	if o == nil || IsNil(o.TlsPemPath) {
		var ret string
		return ret
	}
	return *o.TlsPemPath
}

// GetTlsPemPathOk returns a tuple with the TlsPemPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prometheus) GetTlsPemPathOk() (*string, bool) {
	if o == nil || IsNil(o.TlsPemPath) {
		return nil, false
	}
	return o.TlsPemPath, true
}

// HasTlsPemPath returns a boolean if a field has been set.
func (o *Prometheus) HasTlsPemPath() bool {
	if o != nil && !IsNil(o.TlsPemPath) {
		return true
	}

	return false
}

// SetTlsPemPath gets a reference to the given string and assigns it to the TlsPemPath field.
func (o *Prometheus) SetTlsPemPath(v string) {
	o.TlsPemPath = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Prometheus) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Prometheus) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Prometheus) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Prometheus) SetType(v string) {
	o.Type = &v
}

// GetUsername returns the Username field value
func (o *Prometheus) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *Prometheus) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *Prometheus) SetUsername(v string) {
	o.Username = v
}

func (o Prometheus) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o Prometheus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.ListenAddress) {
		toSerialize["listenAddress"] = o.ListenAddress
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.RateLimitInterval) {
		toSerialize["rateLimitInterval"] = o.RateLimitInterval
	}
	toSerialize["scheme"] = o.Scheme
	toSerialize["serviceDiscovery"] = o.ServiceDiscovery
	if !IsNil(o.TlsPemPath) {
		toSerialize["tlsPemPath"] = o.TlsPemPath
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["username"] = o.Username
	return toSerialize, nil
}

type NullablePrometheus struct {
	value *Prometheus
	isSet bool
}

func (v NullablePrometheus) Get() *Prometheus {
	return v.value
}

func (v *NullablePrometheus) Set(val *Prometheus) {
	v.value = val
	v.isSet = true
}

func (v NullablePrometheus) IsSet() bool {
	return v.isSet
}

func (v *NullablePrometheus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrometheus(val *Prometheus) *NullablePrometheus {
	return &NullablePrometheus{value: val, isSet: true}
}

func (v NullablePrometheus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrometheus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
