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

// checks if the ReportsAttacksAttackers200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportsAttacksAttackers200ResponseDataInner{}

// ReportsAttacksAttackers200ResponseDataInner struct for ReportsAttacksAttackers200ResponseDataInner
type ReportsAttacksAttackers200ResponseDataInner struct {
	Ip *string `json:"ip,omitempty"`
	Count *int32 `json:"count,omitempty"`
}

// NewReportsAttacksAttackers200ResponseDataInner instantiates a new ReportsAttacksAttackers200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportsAttacksAttackers200ResponseDataInner() *ReportsAttacksAttackers200ResponseDataInner {
	this := ReportsAttacksAttackers200ResponseDataInner{}
	return &this
}

// NewReportsAttacksAttackers200ResponseDataInnerWithDefaults instantiates a new ReportsAttacksAttackers200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportsAttacksAttackers200ResponseDataInnerWithDefaults() *ReportsAttacksAttackers200ResponseDataInner {
	this := ReportsAttacksAttackers200ResponseDataInner{}
	return &this
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *ReportsAttacksAttackers200ResponseDataInner) GetIp() string {
	if o == nil || IsNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsAttacksAttackers200ResponseDataInner) GetIpOk() (*string, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *ReportsAttacksAttackers200ResponseDataInner) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *ReportsAttacksAttackers200ResponseDataInner) SetIp(v string) {
	o.Ip = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ReportsAttacksAttackers200ResponseDataInner) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsAttacksAttackers200ResponseDataInner) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ReportsAttacksAttackers200ResponseDataInner) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *ReportsAttacksAttackers200ResponseDataInner) SetCount(v int32) {
	o.Count = &v
}

func (o ReportsAttacksAttackers200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportsAttacksAttackers200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return toSerialize, nil
}

type NullableReportsAttacksAttackers200ResponseDataInner struct {
	value *ReportsAttacksAttackers200ResponseDataInner
	isSet bool
}

func (v NullableReportsAttacksAttackers200ResponseDataInner) Get() *ReportsAttacksAttackers200ResponseDataInner {
	return v.value
}

func (v *NullableReportsAttacksAttackers200ResponseDataInner) Set(val *ReportsAttacksAttackers200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableReportsAttacksAttackers200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableReportsAttacksAttackers200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportsAttacksAttackers200ResponseDataInner(val *ReportsAttacksAttackers200ResponseDataInner) *NullableReportsAttacksAttackers200ResponseDataInner {
	return &NullableReportsAttacksAttackers200ResponseDataInner{value: val, isSet: true}
}

func (v NullableReportsAttacksAttackers200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportsAttacksAttackers200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


