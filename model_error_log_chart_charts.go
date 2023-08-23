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

// checks if the ErrorLogChartCharts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorLogChartCharts{}

// ErrorLogChartCharts struct for ErrorLogChartCharts
type ErrorLogChartCharts struct {
	StatusCode *ErrorLogChartChartsStatusCode `json:"status_code,omitempty"`
}

// NewErrorLogChartCharts instantiates a new ErrorLogChartCharts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorLogChartCharts() *ErrorLogChartCharts {
	this := ErrorLogChartCharts{}
	return &this
}

// NewErrorLogChartChartsWithDefaults instantiates a new ErrorLogChartCharts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorLogChartChartsWithDefaults() *ErrorLogChartCharts {
	this := ErrorLogChartCharts{}
	return &this
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *ErrorLogChartCharts) GetStatusCode() ErrorLogChartChartsStatusCode {
	if o == nil || IsNil(o.StatusCode) {
		var ret ErrorLogChartChartsStatusCode
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorLogChartCharts) GetStatusCodeOk() (*ErrorLogChartChartsStatusCode, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *ErrorLogChartCharts) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given ErrorLogChartChartsStatusCode and assigns it to the StatusCode field.
func (o *ErrorLogChartCharts) SetStatusCode(v ErrorLogChartChartsStatusCode) {
	o.StatusCode = &v
}

func (o ErrorLogChartCharts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorLogChartCharts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StatusCode) {
		toSerialize["status_code"] = o.StatusCode
	}
	return toSerialize, nil
}

type NullableErrorLogChartCharts struct {
	value *ErrorLogChartCharts
	isSet bool
}

func (v NullableErrorLogChartCharts) Get() *ErrorLogChartCharts {
	return v.value
}

func (v *NullableErrorLogChartCharts) Set(val *ErrorLogChartCharts) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorLogChartCharts) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorLogChartCharts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorLogChartCharts(val *ErrorLogChartCharts) *NullableErrorLogChartCharts {
	return &NullableErrorLogChartCharts{value: val, isSet: true}
}

func (v NullableErrorLogChartCharts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorLogChartCharts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


