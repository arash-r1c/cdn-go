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

// checks if the ErrorLogChartStatistics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorLogChartStatistics{}

// ErrorLogChartStatistics struct for ErrorLogChartStatistics
type ErrorLogChartStatistics struct {
	// <key, value> where key is error and value is its count
	StatusCodes map[string]interface{} `json:"status_codes,omitempty"`
}

// NewErrorLogChartStatistics instantiates a new ErrorLogChartStatistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorLogChartStatistics() *ErrorLogChartStatistics {
	this := ErrorLogChartStatistics{}
	return &this
}

// NewErrorLogChartStatisticsWithDefaults instantiates a new ErrorLogChartStatistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorLogChartStatisticsWithDefaults() *ErrorLogChartStatistics {
	this := ErrorLogChartStatistics{}
	return &this
}

// GetStatusCodes returns the StatusCodes field value if set, zero value otherwise.
func (o *ErrorLogChartStatistics) GetStatusCodes() map[string]interface{} {
	if o == nil || IsNil(o.StatusCodes) {
		var ret map[string]interface{}
		return ret
	}
	return o.StatusCodes
}

// GetStatusCodesOk returns a tuple with the StatusCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorLogChartStatistics) GetStatusCodesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.StatusCodes) {
		return map[string]interface{}{}, false
	}
	return o.StatusCodes, true
}

// HasStatusCodes returns a boolean if a field has been set.
func (o *ErrorLogChartStatistics) HasStatusCodes() bool {
	if o != nil && !IsNil(o.StatusCodes) {
		return true
	}

	return false
}

// SetStatusCodes gets a reference to the given map[string]interface{} and assigns it to the StatusCodes field.
func (o *ErrorLogChartStatistics) SetStatusCodes(v map[string]interface{}) {
	o.StatusCodes = v
}

func (o ErrorLogChartStatistics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorLogChartStatistics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StatusCodes) {
		toSerialize["status_codes"] = o.StatusCodes
	}
	return toSerialize, nil
}

type NullableErrorLogChartStatistics struct {
	value *ErrorLogChartStatistics
	isSet bool
}

func (v NullableErrorLogChartStatistics) Get() *ErrorLogChartStatistics {
	return v.value
}

func (v *NullableErrorLogChartStatistics) Set(val *ErrorLogChartStatistics) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorLogChartStatistics) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorLogChartStatistics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorLogChartStatistics(val *ErrorLogChartStatistics) *NullableErrorLogChartStatistics {
	return &NullableErrorLogChartStatistics{value: val, isSet: true}
}

func (v NullableErrorLogChartStatistics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorLogChartStatistics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


