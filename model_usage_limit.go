/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.115.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package arvancloud

import (
	"encoding/json"
)

// checks if the UsageLimit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsageLimit{}

// UsageLimit struct for UsageLimit
type UsageLimit struct {
	Min *int32 `json:"min,omitempty"`
	Max *int32 `json:"max,omitempty"`
}

// NewUsageLimit instantiates a new UsageLimit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageLimit() *UsageLimit {
	this := UsageLimit{}
	return &this
}

// NewUsageLimitWithDefaults instantiates a new UsageLimit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageLimitWithDefaults() *UsageLimit {
	this := UsageLimit{}
	return &this
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *UsageLimit) GetMin() int32 {
	if o == nil || IsNil(o.Min) {
		var ret int32
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageLimit) GetMinOk() (*int32, bool) {
	if o == nil || IsNil(o.Min) {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *UsageLimit) HasMin() bool {
	if o != nil && !IsNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given int32 and assigns it to the Min field.
func (o *UsageLimit) SetMin(v int32) {
	o.Min = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *UsageLimit) GetMax() int32 {
	if o == nil || IsNil(o.Max) {
		var ret int32
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageLimit) GetMaxOk() (*int32, bool) {
	if o == nil || IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *UsageLimit) HasMax() bool {
	if o != nil && !IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given int32 and assigns it to the Max field.
func (o *UsageLimit) SetMax(v int32) {
	o.Max = &v
}

func (o UsageLimit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsageLimit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Min) {
		toSerialize["min"] = o.Min
	}
	if !IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	return toSerialize, nil
}

type NullableUsageLimit struct {
	value *UsageLimit
	isSet bool
}

func (v NullableUsageLimit) Get() *UsageLimit {
	return v.value
}

func (v *NullableUsageLimit) Set(val *UsageLimit) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageLimit) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageLimit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageLimit(val *UsageLimit) *NullableUsageLimit {
	return &NullableUsageLimit{value: val, isSet: true}
}

func (v NullableUsageLimit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageLimit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


