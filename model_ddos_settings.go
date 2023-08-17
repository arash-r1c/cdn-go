/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.105.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package r1cdn

import (
	"encoding/json"
)

// checks if the DdosSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DdosSettings{}

// DdosSettings struct for DdosSettings
type DdosSettings struct {
	IsEnabled *bool `json:"is_enabled,omitempty"`
	ProtectionMode *string `json:"protection_mode,omitempty"`
	CaptchaService *string `json:"captcha_service,omitempty"`
	// Time in seconds for cookie max-age
	Ttl *int32 `json:"ttl,omitempty"`
	// Adds \"SameSite=None; Secure\" to set-cookie header
	HttpsOnly *bool `json:"https_only,omitempty"`
	Preflight *DdosPreflight `json:"preflight,omitempty"`
}

// NewDdosSettings instantiates a new DdosSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDdosSettings() *DdosSettings {
	this := DdosSettings{}
	return &this
}

// NewDdosSettingsWithDefaults instantiates a new DdosSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDdosSettingsWithDefaults() *DdosSettings {
	this := DdosSettings{}
	return &this
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *DdosSettings) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosSettings) GetIsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabled) {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *DdosSettings) HasIsEnabled() bool {
	if o != nil && !IsNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *DdosSettings) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetProtectionMode returns the ProtectionMode field value if set, zero value otherwise.
func (o *DdosSettings) GetProtectionMode() string {
	if o == nil || IsNil(o.ProtectionMode) {
		var ret string
		return ret
	}
	return *o.ProtectionMode
}

// GetProtectionModeOk returns a tuple with the ProtectionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosSettings) GetProtectionModeOk() (*string, bool) {
	if o == nil || IsNil(o.ProtectionMode) {
		return nil, false
	}
	return o.ProtectionMode, true
}

// HasProtectionMode returns a boolean if a field has been set.
func (o *DdosSettings) HasProtectionMode() bool {
	if o != nil && !IsNil(o.ProtectionMode) {
		return true
	}

	return false
}

// SetProtectionMode gets a reference to the given string and assigns it to the ProtectionMode field.
func (o *DdosSettings) SetProtectionMode(v string) {
	o.ProtectionMode = &v
}

// GetCaptchaService returns the CaptchaService field value if set, zero value otherwise.
func (o *DdosSettings) GetCaptchaService() string {
	if o == nil || IsNil(o.CaptchaService) {
		var ret string
		return ret
	}
	return *o.CaptchaService
}

// GetCaptchaServiceOk returns a tuple with the CaptchaService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosSettings) GetCaptchaServiceOk() (*string, bool) {
	if o == nil || IsNil(o.CaptchaService) {
		return nil, false
	}
	return o.CaptchaService, true
}

// HasCaptchaService returns a boolean if a field has been set.
func (o *DdosSettings) HasCaptchaService() bool {
	if o != nil && !IsNil(o.CaptchaService) {
		return true
	}

	return false
}

// SetCaptchaService gets a reference to the given string and assigns it to the CaptchaService field.
func (o *DdosSettings) SetCaptchaService(v string) {
	o.CaptchaService = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *DdosSettings) GetTtl() int32 {
	if o == nil || IsNil(o.Ttl) {
		var ret int32
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosSettings) GetTtlOk() (*int32, bool) {
	if o == nil || IsNil(o.Ttl) {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *DdosSettings) HasTtl() bool {
	if o != nil && !IsNil(o.Ttl) {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int32 and assigns it to the Ttl field.
func (o *DdosSettings) SetTtl(v int32) {
	o.Ttl = &v
}

// GetHttpsOnly returns the HttpsOnly field value if set, zero value otherwise.
func (o *DdosSettings) GetHttpsOnly() bool {
	if o == nil || IsNil(o.HttpsOnly) {
		var ret bool
		return ret
	}
	return *o.HttpsOnly
}

// GetHttpsOnlyOk returns a tuple with the HttpsOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosSettings) GetHttpsOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.HttpsOnly) {
		return nil, false
	}
	return o.HttpsOnly, true
}

// HasHttpsOnly returns a boolean if a field has been set.
func (o *DdosSettings) HasHttpsOnly() bool {
	if o != nil && !IsNil(o.HttpsOnly) {
		return true
	}

	return false
}

// SetHttpsOnly gets a reference to the given bool and assigns it to the HttpsOnly field.
func (o *DdosSettings) SetHttpsOnly(v bool) {
	o.HttpsOnly = &v
}

// GetPreflight returns the Preflight field value if set, zero value otherwise.
func (o *DdosSettings) GetPreflight() DdosPreflight {
	if o == nil || IsNil(o.Preflight) {
		var ret DdosPreflight
		return ret
	}
	return *o.Preflight
}

// GetPreflightOk returns a tuple with the Preflight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosSettings) GetPreflightOk() (*DdosPreflight, bool) {
	if o == nil || IsNil(o.Preflight) {
		return nil, false
	}
	return o.Preflight, true
}

// HasPreflight returns a boolean if a field has been set.
func (o *DdosSettings) HasPreflight() bool {
	if o != nil && !IsNil(o.Preflight) {
		return true
	}

	return false
}

// SetPreflight gets a reference to the given DdosPreflight and assigns it to the Preflight field.
func (o *DdosSettings) SetPreflight(v DdosPreflight) {
	o.Preflight = &v
}

func (o DdosSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DdosSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: is_enabled is readOnly
	if !IsNil(o.ProtectionMode) {
		toSerialize["protection_mode"] = o.ProtectionMode
	}
	if !IsNil(o.CaptchaService) {
		toSerialize["captcha_service"] = o.CaptchaService
	}
	if !IsNil(o.Ttl) {
		toSerialize["ttl"] = o.Ttl
	}
	if !IsNil(o.HttpsOnly) {
		toSerialize["https_only"] = o.HttpsOnly
	}
	if !IsNil(o.Preflight) {
		toSerialize["preflight"] = o.Preflight
	}
	return toSerialize, nil
}

type NullableDdosSettings struct {
	value *DdosSettings
	isSet bool
}

func (v NullableDdosSettings) Get() *DdosSettings {
	return v.value
}

func (v *NullableDdosSettings) Set(val *DdosSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableDdosSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableDdosSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDdosSettings(val *DdosSettings) *NullableDdosSettings {
	return &NullableDdosSettings{value: val, isSet: true}
}

func (v NullableDdosSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDdosSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


