/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.105.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package arvancloud

import (
	"encoding/json"
)

// checks if the RateLimitSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RateLimitSettings{}

// RateLimitSettings struct for RateLimitSettings
type RateLimitSettings struct {
	DdosDetection *bool `json:"ddos_detection,omitempty"`
	ExcludeSources []string `json:"exclude_sources,omitempty"`
}

// NewRateLimitSettings instantiates a new RateLimitSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRateLimitSettings() *RateLimitSettings {
	this := RateLimitSettings{}
	return &this
}

// NewRateLimitSettingsWithDefaults instantiates a new RateLimitSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRateLimitSettingsWithDefaults() *RateLimitSettings {
	this := RateLimitSettings{}
	return &this
}

// GetDdosDetection returns the DdosDetection field value if set, zero value otherwise.
func (o *RateLimitSettings) GetDdosDetection() bool {
	if o == nil || IsNil(o.DdosDetection) {
		var ret bool
		return ret
	}
	return *o.DdosDetection
}

// GetDdosDetectionOk returns a tuple with the DdosDetection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimitSettings) GetDdosDetectionOk() (*bool, bool) {
	if o == nil || IsNil(o.DdosDetection) {
		return nil, false
	}
	return o.DdosDetection, true
}

// HasDdosDetection returns a boolean if a field has been set.
func (o *RateLimitSettings) HasDdosDetection() bool {
	if o != nil && !IsNil(o.DdosDetection) {
		return true
	}

	return false
}

// SetDdosDetection gets a reference to the given bool and assigns it to the DdosDetection field.
func (o *RateLimitSettings) SetDdosDetection(v bool) {
	o.DdosDetection = &v
}

// GetExcludeSources returns the ExcludeSources field value if set, zero value otherwise.
func (o *RateLimitSettings) GetExcludeSources() []string {
	if o == nil || IsNil(o.ExcludeSources) {
		var ret []string
		return ret
	}
	return o.ExcludeSources
}

// GetExcludeSourcesOk returns a tuple with the ExcludeSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimitSettings) GetExcludeSourcesOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludeSources) {
		return nil, false
	}
	return o.ExcludeSources, true
}

// HasExcludeSources returns a boolean if a field has been set.
func (o *RateLimitSettings) HasExcludeSources() bool {
	if o != nil && !IsNil(o.ExcludeSources) {
		return true
	}

	return false
}

// SetExcludeSources gets a reference to the given []string and assigns it to the ExcludeSources field.
func (o *RateLimitSettings) SetExcludeSources(v []string) {
	o.ExcludeSources = v
}

func (o RateLimitSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RateLimitSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DdosDetection) {
		toSerialize["ddos_detection"] = o.DdosDetection
	}
	if !IsNil(o.ExcludeSources) {
		toSerialize["exclude_sources"] = o.ExcludeSources
	}
	return toSerialize, nil
}

type NullableRateLimitSettings struct {
	value *RateLimitSettings
	isSet bool
}

func (v NullableRateLimitSettings) Get() *RateLimitSettings {
	return v.value
}

func (v *NullableRateLimitSettings) Set(val *RateLimitSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableRateLimitSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableRateLimitSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRateLimitSettings(val *RateLimitSettings) *NullableRateLimitSettings {
	return &NullableRateLimitSettings{value: val, isSet: true}
}

func (v NullableRateLimitSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRateLimitSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


