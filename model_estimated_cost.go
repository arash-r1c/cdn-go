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

// checks if the EstimatedCost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EstimatedCost{}

// EstimatedCost struct for EstimatedCost
type EstimatedCost struct {
	Period *string `json:"period,omitempty"`
	Currency *string `json:"currency,omitempty"`
	Value *float32 `json:"value,omitempty"`
}

// NewEstimatedCost instantiates a new EstimatedCost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEstimatedCost() *EstimatedCost {
	this := EstimatedCost{}
	return &this
}

// NewEstimatedCostWithDefaults instantiates a new EstimatedCost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEstimatedCostWithDefaults() *EstimatedCost {
	this := EstimatedCost{}
	return &this
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *EstimatedCost) GetPeriod() string {
	if o == nil || IsNil(o.Period) {
		var ret string
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimatedCost) GetPeriodOk() (*string, bool) {
	if o == nil || IsNil(o.Period) {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *EstimatedCost) HasPeriod() bool {
	if o != nil && !IsNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given string and assigns it to the Period field.
func (o *EstimatedCost) SetPeriod(v string) {
	o.Period = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *EstimatedCost) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimatedCost) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *EstimatedCost) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *EstimatedCost) SetCurrency(v string) {
	o.Currency = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *EstimatedCost) GetValue() float32 {
	if o == nil || IsNil(o.Value) {
		var ret float32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EstimatedCost) GetValueOk() (*float32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *EstimatedCost) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float32 and assigns it to the Value field.
func (o *EstimatedCost) SetValue(v float32) {
	o.Value = &v
}

func (o EstimatedCost) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EstimatedCost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Period) {
		toSerialize["period"] = o.Period
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableEstimatedCost struct {
	value *EstimatedCost
	isSet bool
}

func (v NullableEstimatedCost) Get() *EstimatedCost {
	return v.value
}

func (v *NullableEstimatedCost) Set(val *EstimatedCost) {
	v.value = val
	v.isSet = true
}

func (v NullableEstimatedCost) IsSet() bool {
	return v.isSet
}

func (v *NullableEstimatedCost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEstimatedCost(val *EstimatedCost) *NullableEstimatedCost {
	return &NullableEstimatedCost{value: val, isSet: true}
}

func (v NullableEstimatedCost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEstimatedCost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


