/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.114.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package arvancloud

import (
	"encoding/json"
)

// checks if the HighRequestedIp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HighRequestedIp{}

// HighRequestedIp struct for HighRequestedIp
type HighRequestedIp struct {
	Ip *HighRequestedIpIp `json:"ip,omitempty"`
	RequestCount *int32 `json:"request_count,omitempty"`
}

// NewHighRequestedIp instantiates a new HighRequestedIp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHighRequestedIp() *HighRequestedIp {
	this := HighRequestedIp{}
	return &this
}

// NewHighRequestedIpWithDefaults instantiates a new HighRequestedIp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHighRequestedIpWithDefaults() *HighRequestedIp {
	this := HighRequestedIp{}
	return &this
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *HighRequestedIp) GetIp() HighRequestedIpIp {
	if o == nil || IsNil(o.Ip) {
		var ret HighRequestedIpIp
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HighRequestedIp) GetIpOk() (*HighRequestedIpIp, bool) {
	if o == nil || IsNil(o.Ip) {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *HighRequestedIp) HasIp() bool {
	if o != nil && !IsNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given HighRequestedIpIp and assigns it to the Ip field.
func (o *HighRequestedIp) SetIp(v HighRequestedIpIp) {
	o.Ip = &v
}

// GetRequestCount returns the RequestCount field value if set, zero value otherwise.
func (o *HighRequestedIp) GetRequestCount() int32 {
	if o == nil || IsNil(o.RequestCount) {
		var ret int32
		return ret
	}
	return *o.RequestCount
}

// GetRequestCountOk returns a tuple with the RequestCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HighRequestedIp) GetRequestCountOk() (*int32, bool) {
	if o == nil || IsNil(o.RequestCount) {
		return nil, false
	}
	return o.RequestCount, true
}

// HasRequestCount returns a boolean if a field has been set.
func (o *HighRequestedIp) HasRequestCount() bool {
	if o != nil && !IsNil(o.RequestCount) {
		return true
	}

	return false
}

// SetRequestCount gets a reference to the given int32 and assigns it to the RequestCount field.
func (o *HighRequestedIp) SetRequestCount(v int32) {
	o.RequestCount = &v
}

func (o HighRequestedIp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HighRequestedIp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !IsNil(o.RequestCount) {
		toSerialize["request_count"] = o.RequestCount
	}
	return toSerialize, nil
}

type NullableHighRequestedIp struct {
	value *HighRequestedIp
	isSet bool
}

func (v NullableHighRequestedIp) Get() *HighRequestedIp {
	return v.value
}

func (v *NullableHighRequestedIp) Set(val *HighRequestedIp) {
	v.value = val
	v.isSet = true
}

func (v NullableHighRequestedIp) IsSet() bool {
	return v.isSet
}

func (v *NullableHighRequestedIp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHighRequestedIp(val *HighRequestedIp) *NullableHighRequestedIp {
	return &NullableHighRequestedIp{value: val, isSet: true}
}

func (v NullableHighRequestedIp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHighRequestedIp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


