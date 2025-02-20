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

// checks if the TrafficsMapCharts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrafficsMapCharts{}

// TrafficsMapCharts struct for TrafficsMapCharts
type TrafficsMapCharts struct {
	Requests *map[string]CountryRequestChart `json:"requests,omitempty"`
	Traffics *map[string]CountryTrafficChart `json:"traffics,omitempty"`
}

// NewTrafficsMapCharts instantiates a new TrafficsMapCharts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrafficsMapCharts() *TrafficsMapCharts {
	this := TrafficsMapCharts{}
	return &this
}

// NewTrafficsMapChartsWithDefaults instantiates a new TrafficsMapCharts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrafficsMapChartsWithDefaults() *TrafficsMapCharts {
	this := TrafficsMapCharts{}
	return &this
}

// GetRequests returns the Requests field value if set, zero value otherwise.
func (o *TrafficsMapCharts) GetRequests() map[string]CountryRequestChart {
	if o == nil || IsNil(o.Requests) {
		var ret map[string]CountryRequestChart
		return ret
	}
	return *o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficsMapCharts) GetRequestsOk() (*map[string]CountryRequestChart, bool) {
	if o == nil || IsNil(o.Requests) {
		return nil, false
	}
	return o.Requests, true
}

// HasRequests returns a boolean if a field has been set.
func (o *TrafficsMapCharts) HasRequests() bool {
	if o != nil && !IsNil(o.Requests) {
		return true
	}

	return false
}

// SetRequests gets a reference to the given map[string]CountryRequestChart and assigns it to the Requests field.
func (o *TrafficsMapCharts) SetRequests(v map[string]CountryRequestChart) {
	o.Requests = &v
}

// GetTraffics returns the Traffics field value if set, zero value otherwise.
func (o *TrafficsMapCharts) GetTraffics() map[string]CountryTrafficChart {
	if o == nil || IsNil(o.Traffics) {
		var ret map[string]CountryTrafficChart
		return ret
	}
	return *o.Traffics
}

// GetTrafficsOk returns a tuple with the Traffics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficsMapCharts) GetTrafficsOk() (*map[string]CountryTrafficChart, bool) {
	if o == nil || IsNil(o.Traffics) {
		return nil, false
	}
	return o.Traffics, true
}

// HasTraffics returns a boolean if a field has been set.
func (o *TrafficsMapCharts) HasTraffics() bool {
	if o != nil && !IsNil(o.Traffics) {
		return true
	}

	return false
}

// SetTraffics gets a reference to the given map[string]CountryTrafficChart and assigns it to the Traffics field.
func (o *TrafficsMapCharts) SetTraffics(v map[string]CountryTrafficChart) {
	o.Traffics = &v
}

func (o TrafficsMapCharts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrafficsMapCharts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Requests) {
		toSerialize["requests"] = o.Requests
	}
	if !IsNil(o.Traffics) {
		toSerialize["traffics"] = o.Traffics
	}
	return toSerialize, nil
}

type NullableTrafficsMapCharts struct {
	value *TrafficsMapCharts
	isSet bool
}

func (v NullableTrafficsMapCharts) Get() *TrafficsMapCharts {
	return v.value
}

func (v *NullableTrafficsMapCharts) Set(val *TrafficsMapCharts) {
	v.value = val
	v.isSet = true
}

func (v NullableTrafficsMapCharts) IsSet() bool {
	return v.isSet
}

func (v *NullableTrafficsMapCharts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrafficsMapCharts(val *TrafficsMapCharts) *NullableTrafficsMapCharts {
	return &NullableTrafficsMapCharts{value: val, isSet: true}
}

func (v NullableTrafficsMapCharts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrafficsMapCharts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


