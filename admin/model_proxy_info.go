// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"encoding/json"
)

// ProxyInfo VPC Proxy Information.
type ProxyInfo struct {
	// Authentication key for the proxy.
	AuthKey *string `json:"authKey,omitempty"`
	// DNS name to use to reach the proxy/s.
	DnsName *string `json:"dnsName,omitempty"`
}

// NewProxyInfo instantiates a new ProxyInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProxyInfo() *ProxyInfo {
	this := ProxyInfo{}
	return &this
}

// NewProxyInfoWithDefaults instantiates a new ProxyInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProxyInfoWithDefaults() *ProxyInfo {
	this := ProxyInfo{}
	return &this
}

// GetAuthKey returns the AuthKey field value if set, zero value otherwise
func (o *ProxyInfo) GetAuthKey() string {
	if o == nil || IsNil(o.AuthKey) {
		var ret string
		return ret
	}
	return *o.AuthKey
}

// GetAuthKeyOk returns a tuple with the AuthKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyInfo) GetAuthKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AuthKey) {
		return nil, false
	}

	return o.AuthKey, true
}

// HasAuthKey returns a boolean if a field has been set.
func (o *ProxyInfo) HasAuthKey() bool {
	if o != nil && !IsNil(o.AuthKey) {
		return true
	}

	return false
}

// SetAuthKey gets a reference to the given string and assigns it to the AuthKey field.
func (o *ProxyInfo) SetAuthKey(v string) {
	o.AuthKey = &v
}

// GetDnsName returns the DnsName field value if set, zero value otherwise
func (o *ProxyInfo) GetDnsName() string {
	if o == nil || IsNil(o.DnsName) {
		var ret string
		return ret
	}
	return *o.DnsName
}

// GetDnsNameOk returns a tuple with the DnsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProxyInfo) GetDnsNameOk() (*string, bool) {
	if o == nil || IsNil(o.DnsName) {
		return nil, false
	}

	return o.DnsName, true
}

// HasDnsName returns a boolean if a field has been set.
func (o *ProxyInfo) HasDnsName() bool {
	if o != nil && !IsNil(o.DnsName) {
		return true
	}

	return false
}

// SetDnsName gets a reference to the given string and assigns it to the DnsName field.
func (o *ProxyInfo) SetDnsName(v string) {
	o.DnsName = &v
}

func (o ProxyInfo) MarshalJSONWithoutReadOnly() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}
func (o ProxyInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthKey) {
		toSerialize["authKey"] = o.AuthKey
	}
	if !IsNil(o.DnsName) {
		toSerialize["dnsName"] = o.DnsName
	}
	return toSerialize, nil
}
