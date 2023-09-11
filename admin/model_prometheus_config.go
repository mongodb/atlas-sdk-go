// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// PrometheusConfig struct for PrometheusConfig
type PrometheusConfig struct {
	Enabled                  *bool   `json:"enabled,omitempty"`
	ListenAddress            *string `json:"listenAddress,omitempty"`
	MetricsPath              *string `json:"metricsPath,omitempty"`
	Mode                     *string `json:"mode,omitempty"`
	RateLimitIntervalSeconds *int    `json:"rateLimitIntervalSeconds,omitempty"`
	Scheme                   *string `json:"scheme,omitempty"`
	ServiceDiscovery         *string `json:"serviceDiscovery,omitempty"`
	TlsPemPassword           *string `json:"tlsPemPassword,omitempty"`
	TlsPemPath               *string `json:"tlsPemPath,omitempty"`
	Username                 *string `json:"username,omitempty"`
}

// NewPrometheusConfig instantiates a new PrometheusConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrometheusConfig() *PrometheusConfig {
	this := PrometheusConfig{}
	return &this
}

// NewPrometheusConfigWithDefaults instantiates a new PrometheusConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrometheusConfigWithDefaults() *PrometheusConfig {
	this := PrometheusConfig{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise
func (o *PrometheusConfig) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusConfig) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}

	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PrometheusConfig) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PrometheusConfig) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetListenAddress returns the ListenAddress field value if set, zero value otherwise
func (o *PrometheusConfig) GetListenAddress() string {
	if o == nil || IsNil(o.ListenAddress) {
		var ret string
		return ret
	}
	return *o.ListenAddress
}

// GetListenAddressOk returns a tuple with the ListenAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusConfig) GetListenAddressOk() (*string, bool) {
	if o == nil || IsNil(o.ListenAddress) {
		return nil, false
	}

	return o.ListenAddress, true
}

// HasListenAddress returns a boolean if a field has been set.
func (o *PrometheusConfig) HasListenAddress() bool {
	if o != nil && !IsNil(o.ListenAddress) {
		return true
	}

	return false
}

// SetListenAddress gets a reference to the given string and assigns it to the ListenAddress field.
func (o *PrometheusConfig) SetListenAddress(v string) {
	o.ListenAddress = &v
}

// GetMetricsPath returns the MetricsPath field value if set, zero value otherwise
func (o *PrometheusConfig) GetMetricsPath() string {
	if o == nil || IsNil(o.MetricsPath) {
		var ret string
		return ret
	}
	return *o.MetricsPath
}

// GetMetricsPathOk returns a tuple with the MetricsPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusConfig) GetMetricsPathOk() (*string, bool) {
	if o == nil || IsNil(o.MetricsPath) {
		return nil, false
	}

	return o.MetricsPath, true
}

// HasMetricsPath returns a boolean if a field has been set.
func (o *PrometheusConfig) HasMetricsPath() bool {
	if o != nil && !IsNil(o.MetricsPath) {
		return true
	}

	return false
}

// SetMetricsPath gets a reference to the given string and assigns it to the MetricsPath field.
func (o *PrometheusConfig) SetMetricsPath(v string) {
	o.MetricsPath = &v
}

// GetMode returns the Mode field value if set, zero value otherwise
func (o *PrometheusConfig) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusConfig) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}

	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *PrometheusConfig) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *PrometheusConfig) SetMode(v string) {
	o.Mode = &v
}

// GetRateLimitIntervalSeconds returns the RateLimitIntervalSeconds field value if set, zero value otherwise
func (o *PrometheusConfig) GetRateLimitIntervalSeconds() int {
	if o == nil || IsNil(o.RateLimitIntervalSeconds) {
		var ret int
		return ret
	}
	return *o.RateLimitIntervalSeconds
}

// GetRateLimitIntervalSecondsOk returns a tuple with the RateLimitIntervalSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusConfig) GetRateLimitIntervalSecondsOk() (*int, bool) {
	if o == nil || IsNil(o.RateLimitIntervalSeconds) {
		return nil, false
	}

	return o.RateLimitIntervalSeconds, true
}

// HasRateLimitIntervalSeconds returns a boolean if a field has been set.
func (o *PrometheusConfig) HasRateLimitIntervalSeconds() bool {
	if o != nil && !IsNil(o.RateLimitIntervalSeconds) {
		return true
	}

	return false
}

// SetRateLimitIntervalSeconds gets a reference to the given int and assigns it to the RateLimitIntervalSeconds field.
func (o *PrometheusConfig) SetRateLimitIntervalSeconds(v int) {
	o.RateLimitIntervalSeconds = &v
}

// GetScheme returns the Scheme field value if set, zero value otherwise
func (o *PrometheusConfig) GetScheme() string {
	if o == nil || IsNil(o.Scheme) {
		var ret string
		return ret
	}
	return *o.Scheme
}

// GetSchemeOk returns a tuple with the Scheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusConfig) GetSchemeOk() (*string, bool) {
	if o == nil || IsNil(o.Scheme) {
		return nil, false
	}

	return o.Scheme, true
}

// HasScheme returns a boolean if a field has been set.
func (o *PrometheusConfig) HasScheme() bool {
	if o != nil && !IsNil(o.Scheme) {
		return true
	}

	return false
}

// SetScheme gets a reference to the given string and assigns it to the Scheme field.
func (o *PrometheusConfig) SetScheme(v string) {
	o.Scheme = &v
}

// GetServiceDiscovery returns the ServiceDiscovery field value if set, zero value otherwise
func (o *PrometheusConfig) GetServiceDiscovery() string {
	if o == nil || IsNil(o.ServiceDiscovery) {
		var ret string
		return ret
	}
	return *o.ServiceDiscovery
}

// GetServiceDiscoveryOk returns a tuple with the ServiceDiscovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusConfig) GetServiceDiscoveryOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceDiscovery) {
		return nil, false
	}

	return o.ServiceDiscovery, true
}

// HasServiceDiscovery returns a boolean if a field has been set.
func (o *PrometheusConfig) HasServiceDiscovery() bool {
	if o != nil && !IsNil(o.ServiceDiscovery) {
		return true
	}

	return false
}

// SetServiceDiscovery gets a reference to the given string and assigns it to the ServiceDiscovery field.
func (o *PrometheusConfig) SetServiceDiscovery(v string) {
	o.ServiceDiscovery = &v
}

// GetTlsPemPassword returns the TlsPemPassword field value if set, zero value otherwise
func (o *PrometheusConfig) GetTlsPemPassword() string {
	if o == nil || IsNil(o.TlsPemPassword) {
		var ret string
		return ret
	}
	return *o.TlsPemPassword
}

// GetTlsPemPasswordOk returns a tuple with the TlsPemPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusConfig) GetTlsPemPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.TlsPemPassword) {
		return nil, false
	}

	return o.TlsPemPassword, true
}

// HasTlsPemPassword returns a boolean if a field has been set.
func (o *PrometheusConfig) HasTlsPemPassword() bool {
	if o != nil && !IsNil(o.TlsPemPassword) {
		return true
	}

	return false
}

// SetTlsPemPassword gets a reference to the given string and assigns it to the TlsPemPassword field.
func (o *PrometheusConfig) SetTlsPemPassword(v string) {
	o.TlsPemPassword = &v
}

// GetTlsPemPath returns the TlsPemPath field value if set, zero value otherwise
func (o *PrometheusConfig) GetTlsPemPath() string {
	if o == nil || IsNil(o.TlsPemPath) {
		var ret string
		return ret
	}
	return *o.TlsPemPath
}

// GetTlsPemPathOk returns a tuple with the TlsPemPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusConfig) GetTlsPemPathOk() (*string, bool) {
	if o == nil || IsNil(o.TlsPemPath) {
		return nil, false
	}

	return o.TlsPemPath, true
}

// HasTlsPemPath returns a boolean if a field has been set.
func (o *PrometheusConfig) HasTlsPemPath() bool {
	if o != nil && !IsNil(o.TlsPemPath) {
		return true
	}

	return false
}

// SetTlsPemPath gets a reference to the given string and assigns it to the TlsPemPath field.
func (o *PrometheusConfig) SetTlsPemPath(v string) {
	o.TlsPemPath = &v
}

// GetUsername returns the Username field value if set, zero value otherwise
func (o *PrometheusConfig) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrometheusConfig) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}

	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *PrometheusConfig) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *PrometheusConfig) SetUsername(v string) {
	o.Username = &v
}

func (o PrometheusConfig) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o PrometheusConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.ListenAddress) {
		toSerialize["listenAddress"] = o.ListenAddress
	}
	if !IsNil(o.MetricsPath) {
		toSerialize["metricsPath"] = o.MetricsPath
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.RateLimitIntervalSeconds) {
		toSerialize["rateLimitIntervalSeconds"] = o.RateLimitIntervalSeconds
	}
	if !IsNil(o.Scheme) {
		toSerialize["scheme"] = o.Scheme
	}
	if !IsNil(o.ServiceDiscovery) {
		toSerialize["serviceDiscovery"] = o.ServiceDiscovery
	}
	if !IsNil(o.TlsPemPassword) {
		toSerialize["tlsPemPassword"] = o.TlsPemPassword
	}
	if !IsNil(o.TlsPemPath) {
		toSerialize["tlsPemPath"] = o.TlsPemPath
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}
