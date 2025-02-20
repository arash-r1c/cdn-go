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

// checks if the ResponseTimeStatistics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseTimeStatistics{}

// ResponseTimeStatistics struct for ResponseTimeStatistics
type ResponseTimeStatistics struct {
	ResponseTime interface{} `json:"response_time,omitempty"`
}

// NewResponseTimeStatistics instantiates a new ResponseTimeStatistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseTimeStatistics() *ResponseTimeStatistics {
	this := ResponseTimeStatistics{}
	return &this
}

// NewResponseTimeStatisticsWithDefaults instantiates a new ResponseTimeStatistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseTimeStatisticsWithDefaults() *ResponseTimeStatistics {
	this := ResponseTimeStatistics{}
	return &this
}

// GetResponseTime returns the ResponseTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResponseTimeStatistics) GetResponseTime() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ResponseTime
}

// GetResponseTimeOk returns a tuple with the ResponseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResponseTimeStatistics) GetResponseTimeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ResponseTime) {
		return nil, false
	}
	return &o.ResponseTime, true
}

// HasResponseTime returns a boolean if a field has been set.
func (o *ResponseTimeStatistics) HasResponseTime() bool {
	if o != nil && IsNil(o.ResponseTime) {
		return true
	}

	return false
}

// SetResponseTime gets a reference to the given interface{} and assigns it to the ResponseTime field.
func (o *ResponseTimeStatistics) SetResponseTime(v interface{}) {
	o.ResponseTime = v
}

func (o ResponseTimeStatistics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseTimeStatistics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ResponseTime != nil {
		toSerialize["response_time"] = o.ResponseTime
	}
	return toSerialize, nil
}

type NullableResponseTimeStatistics struct {
	value *ResponseTimeStatistics
	isSet bool
}

func (v NullableResponseTimeStatistics) Get() *ResponseTimeStatistics {
	return v.value
}

func (v *NullableResponseTimeStatistics) Set(val *ResponseTimeStatistics) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseTimeStatistics) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseTimeStatistics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseTimeStatistics(val *ResponseTimeStatistics) *NullableResponseTimeStatistics {
	return &NullableResponseTimeStatistics{value: val, isSet: true}
}

func (v NullableResponseTimeStatistics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseTimeStatistics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


