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

// checks if the StatusCodeSummaryChartsStatusCodeInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatusCodeSummaryChartsStatusCodeInner{}

// StatusCodeSummaryChartsStatusCodeInner struct for StatusCodeSummaryChartsStatusCodeInner
type StatusCodeSummaryChartsStatusCodeInner struct {
	Name *string `json:"name,omitempty"`
	Y *int32 `json:"y,omitempty"`
}

// NewStatusCodeSummaryChartsStatusCodeInner instantiates a new StatusCodeSummaryChartsStatusCodeInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusCodeSummaryChartsStatusCodeInner() *StatusCodeSummaryChartsStatusCodeInner {
	this := StatusCodeSummaryChartsStatusCodeInner{}
	return &this
}

// NewStatusCodeSummaryChartsStatusCodeInnerWithDefaults instantiates a new StatusCodeSummaryChartsStatusCodeInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusCodeSummaryChartsStatusCodeInnerWithDefaults() *StatusCodeSummaryChartsStatusCodeInner {
	this := StatusCodeSummaryChartsStatusCodeInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StatusCodeSummaryChartsStatusCodeInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusCodeSummaryChartsStatusCodeInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StatusCodeSummaryChartsStatusCodeInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StatusCodeSummaryChartsStatusCodeInner) SetName(v string) {
	o.Name = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *StatusCodeSummaryChartsStatusCodeInner) GetY() int32 {
	if o == nil || IsNil(o.Y) {
		var ret int32
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusCodeSummaryChartsStatusCodeInner) GetYOk() (*int32, bool) {
	if o == nil || IsNil(o.Y) {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *StatusCodeSummaryChartsStatusCodeInner) HasY() bool {
	if o != nil && !IsNil(o.Y) {
		return true
	}

	return false
}

// SetY gets a reference to the given int32 and assigns it to the Y field.
func (o *StatusCodeSummaryChartsStatusCodeInner) SetY(v int32) {
	o.Y = &v
}

func (o StatusCodeSummaryChartsStatusCodeInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatusCodeSummaryChartsStatusCodeInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Y) {
		toSerialize["y"] = o.Y
	}
	return toSerialize, nil
}

type NullableStatusCodeSummaryChartsStatusCodeInner struct {
	value *StatusCodeSummaryChartsStatusCodeInner
	isSet bool
}

func (v NullableStatusCodeSummaryChartsStatusCodeInner) Get() *StatusCodeSummaryChartsStatusCodeInner {
	return v.value
}

func (v *NullableStatusCodeSummaryChartsStatusCodeInner) Set(val *StatusCodeSummaryChartsStatusCodeInner) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusCodeSummaryChartsStatusCodeInner) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusCodeSummaryChartsStatusCodeInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusCodeSummaryChartsStatusCodeInner(val *StatusCodeSummaryChartsStatusCodeInner) *NullableStatusCodeSummaryChartsStatusCodeInner {
	return &NullableStatusCodeSummaryChartsStatusCodeInner{value: val, isSet: true}
}

func (v NullableStatusCodeSummaryChartsStatusCodeInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusCodeSummaryChartsStatusCodeInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


