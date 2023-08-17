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

// checks if the TrafficStatistics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrafficStatistics{}

// TrafficStatistics struct for TrafficStatistics
type TrafficStatistics struct {
	Traffics *TrafficStatisticsTraffics `json:"traffics,omitempty"`
	Requests *TrafficStatisticsTraffics `json:"requests,omitempty"`
}

// NewTrafficStatistics instantiates a new TrafficStatistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrafficStatistics() *TrafficStatistics {
	this := TrafficStatistics{}
	return &this
}

// NewTrafficStatisticsWithDefaults instantiates a new TrafficStatistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrafficStatisticsWithDefaults() *TrafficStatistics {
	this := TrafficStatistics{}
	return &this
}

// GetTraffics returns the Traffics field value if set, zero value otherwise.
func (o *TrafficStatistics) GetTraffics() TrafficStatisticsTraffics {
	if o == nil || IsNil(o.Traffics) {
		var ret TrafficStatisticsTraffics
		return ret
	}
	return *o.Traffics
}

// GetTrafficsOk returns a tuple with the Traffics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficStatistics) GetTrafficsOk() (*TrafficStatisticsTraffics, bool) {
	if o == nil || IsNil(o.Traffics) {
		return nil, false
	}
	return o.Traffics, true
}

// HasTraffics returns a boolean if a field has been set.
func (o *TrafficStatistics) HasTraffics() bool {
	if o != nil && !IsNil(o.Traffics) {
		return true
	}

	return false
}

// SetTraffics gets a reference to the given TrafficStatisticsTraffics and assigns it to the Traffics field.
func (o *TrafficStatistics) SetTraffics(v TrafficStatisticsTraffics) {
	o.Traffics = &v
}

// GetRequests returns the Requests field value if set, zero value otherwise.
func (o *TrafficStatistics) GetRequests() TrafficStatisticsTraffics {
	if o == nil || IsNil(o.Requests) {
		var ret TrafficStatisticsTraffics
		return ret
	}
	return *o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficStatistics) GetRequestsOk() (*TrafficStatisticsTraffics, bool) {
	if o == nil || IsNil(o.Requests) {
		return nil, false
	}
	return o.Requests, true
}

// HasRequests returns a boolean if a field has been set.
func (o *TrafficStatistics) HasRequests() bool {
	if o != nil && !IsNil(o.Requests) {
		return true
	}

	return false
}

// SetRequests gets a reference to the given TrafficStatisticsTraffics and assigns it to the Requests field.
func (o *TrafficStatistics) SetRequests(v TrafficStatisticsTraffics) {
	o.Requests = &v
}

func (o TrafficStatistics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrafficStatistics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Traffics) {
		toSerialize["traffics"] = o.Traffics
	}
	if !IsNil(o.Requests) {
		toSerialize["requests"] = o.Requests
	}
	return toSerialize, nil
}

type NullableTrafficStatistics struct {
	value *TrafficStatistics
	isSet bool
}

func (v NullableTrafficStatistics) Get() *TrafficStatistics {
	return v.value
}

func (v *NullableTrafficStatistics) Set(val *TrafficStatistics) {
	v.value = val
	v.isSet = true
}

func (v NullableTrafficStatistics) IsSet() bool {
	return v.isSet
}

func (v *NullableTrafficStatistics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrafficStatistics(val *TrafficStatistics) *NullableTrafficStatistics {
	return &NullableTrafficStatistics{value: val, isSet: true}
}

func (v NullableTrafficStatistics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrafficStatistics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


