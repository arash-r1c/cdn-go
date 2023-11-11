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

// checks if the PTRRecordValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PTRRecordValue{}

// PTRRecordValue struct for PTRRecordValue
type PTRRecordValue struct {
	Domain *string `json:"domain,omitempty"`
}

// NewPTRRecordValue instantiates a new PTRRecordValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPTRRecordValue() *PTRRecordValue {
	this := PTRRecordValue{}
	return &this
}

// NewPTRRecordValueWithDefaults instantiates a new PTRRecordValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPTRRecordValueWithDefaults() *PTRRecordValue {
	this := PTRRecordValue{}
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *PTRRecordValue) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PTRRecordValue) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *PTRRecordValue) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *PTRRecordValue) SetDomain(v string) {
	o.Domain = &v
}

func (o PTRRecordValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PTRRecordValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	return toSerialize, nil
}

type NullablePTRRecordValue struct {
	value *PTRRecordValue
	isSet bool
}

func (v NullablePTRRecordValue) Get() *PTRRecordValue {
	return v.value
}

func (v *NullablePTRRecordValue) Set(val *PTRRecordValue) {
	v.value = val
	v.isSet = true
}

func (v NullablePTRRecordValue) IsSet() bool {
	return v.isSet
}

func (v *NullablePTRRecordValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePTRRecordValue(val *PTRRecordValue) *NullablePTRRecordValue {
	return &NullablePTRRecordValue{value: val, isSet: true}
}

func (v NullablePTRRecordValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePTRRecordValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


