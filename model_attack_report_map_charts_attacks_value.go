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

// checks if the AttackReportMapChartsAttacksValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttackReportMapChartsAttacksValue{}

// AttackReportMapChartsAttacksValue struct for AttackReportMapChartsAttacksValue
type AttackReportMapChartsAttacksValue struct {
	// A numerical key used for coloring the map
	FillKey *int64 `json:"fillKey,omitempty"`
	// The name of the country
	Name *string `json:"name,omitempty"`
	// The number of attacks
	Value *int64 `json:"value,omitempty"`
}

// NewAttackReportMapChartsAttacksValue instantiates a new AttackReportMapChartsAttacksValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttackReportMapChartsAttacksValue() *AttackReportMapChartsAttacksValue {
	this := AttackReportMapChartsAttacksValue{}
	return &this
}

// NewAttackReportMapChartsAttacksValueWithDefaults instantiates a new AttackReportMapChartsAttacksValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttackReportMapChartsAttacksValueWithDefaults() *AttackReportMapChartsAttacksValue {
	this := AttackReportMapChartsAttacksValue{}
	return &this
}

// GetFillKey returns the FillKey field value if set, zero value otherwise.
func (o *AttackReportMapChartsAttacksValue) GetFillKey() int64 {
	if o == nil || IsNil(o.FillKey) {
		var ret int64
		return ret
	}
	return *o.FillKey
}

// GetFillKeyOk returns a tuple with the FillKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackReportMapChartsAttacksValue) GetFillKeyOk() (*int64, bool) {
	if o == nil || IsNil(o.FillKey) {
		return nil, false
	}
	return o.FillKey, true
}

// HasFillKey returns a boolean if a field has been set.
func (o *AttackReportMapChartsAttacksValue) HasFillKey() bool {
	if o != nil && !IsNil(o.FillKey) {
		return true
	}

	return false
}

// SetFillKey gets a reference to the given int64 and assigns it to the FillKey field.
func (o *AttackReportMapChartsAttacksValue) SetFillKey(v int64) {
	o.FillKey = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AttackReportMapChartsAttacksValue) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackReportMapChartsAttacksValue) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AttackReportMapChartsAttacksValue) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AttackReportMapChartsAttacksValue) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *AttackReportMapChartsAttacksValue) GetValue() int64 {
	if o == nil || IsNil(o.Value) {
		var ret int64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackReportMapChartsAttacksValue) GetValueOk() (*int64, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *AttackReportMapChartsAttacksValue) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int64 and assigns it to the Value field.
func (o *AttackReportMapChartsAttacksValue) SetValue(v int64) {
	o.Value = &v
}

func (o AttackReportMapChartsAttacksValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttackReportMapChartsAttacksValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FillKey) {
		toSerialize["fillKey"] = o.FillKey
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableAttackReportMapChartsAttacksValue struct {
	value *AttackReportMapChartsAttacksValue
	isSet bool
}

func (v NullableAttackReportMapChartsAttacksValue) Get() *AttackReportMapChartsAttacksValue {
	return v.value
}

func (v *NullableAttackReportMapChartsAttacksValue) Set(val *AttackReportMapChartsAttacksValue) {
	v.value = val
	v.isSet = true
}

func (v NullableAttackReportMapChartsAttacksValue) IsSet() bool {
	return v.isSet
}

func (v *NullableAttackReportMapChartsAttacksValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttackReportMapChartsAttacksValue(val *AttackReportMapChartsAttacksValue) *NullableAttackReportMapChartsAttacksValue {
	return &NullableAttackReportMapChartsAttacksValue{value: val, isSet: true}
}

func (v NullableAttackReportMapChartsAttacksValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttackReportMapChartsAttacksValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


