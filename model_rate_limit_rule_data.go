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

// checks if the RateLimitRuleData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RateLimitRuleData{}

// RateLimitRuleData struct for RateLimitRuleData
type RateLimitRuleData struct {
	Data *RateLimitRuleView `json:"data,omitempty"`
}

// NewRateLimitRuleData instantiates a new RateLimitRuleData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRateLimitRuleData() *RateLimitRuleData {
	this := RateLimitRuleData{}
	return &this
}

// NewRateLimitRuleDataWithDefaults instantiates a new RateLimitRuleData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRateLimitRuleDataWithDefaults() *RateLimitRuleData {
	this := RateLimitRuleData{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RateLimitRuleData) GetData() RateLimitRuleView {
	if o == nil || IsNil(o.Data) {
		var ret RateLimitRuleView
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimitRuleData) GetDataOk() (*RateLimitRuleView, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RateLimitRuleData) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given RateLimitRuleView and assigns it to the Data field.
func (o *RateLimitRuleData) SetData(v RateLimitRuleView) {
	o.Data = &v
}

func (o RateLimitRuleData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RateLimitRuleData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableRateLimitRuleData struct {
	value *RateLimitRuleData
	isSet bool
}

func (v NullableRateLimitRuleData) Get() *RateLimitRuleData {
	return v.value
}

func (v *NullableRateLimitRuleData) Set(val *RateLimitRuleData) {
	v.value = val
	v.isSet = true
}

func (v NullableRateLimitRuleData) IsSet() bool {
	return v.isSet
}

func (v *NullableRateLimitRuleData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRateLimitRuleData(val *RateLimitRuleData) *NullableRateLimitRuleData {
	return &NullableRateLimitRuleData{value: val, isSet: true}
}

func (v NullableRateLimitRuleData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRateLimitRuleData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


