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

// checks if the CountryStatistics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CountryStatistics{}

// CountryStatistics struct for CountryStatistics
type CountryStatistics struct {
	// The name of the country
	Country *string `json:"country,omitempty"`
	// The number of requests from the country
	Requests *int32 `json:"requests,omitempty"`
	// The amount of traffic from the country
	Traffics *int32 `json:"traffics,omitempty"`
}

// NewCountryStatistics instantiates a new CountryStatistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountryStatistics() *CountryStatistics {
	this := CountryStatistics{}
	return &this
}

// NewCountryStatisticsWithDefaults instantiates a new CountryStatistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountryStatisticsWithDefaults() *CountryStatistics {
	this := CountryStatistics{}
	return &this
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *CountryStatistics) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryStatistics) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *CountryStatistics) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *CountryStatistics) SetCountry(v string) {
	o.Country = &v
}

// GetRequests returns the Requests field value if set, zero value otherwise.
func (o *CountryStatistics) GetRequests() int32 {
	if o == nil || IsNil(o.Requests) {
		var ret int32
		return ret
	}
	return *o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryStatistics) GetRequestsOk() (*int32, bool) {
	if o == nil || IsNil(o.Requests) {
		return nil, false
	}
	return o.Requests, true
}

// HasRequests returns a boolean if a field has been set.
func (o *CountryStatistics) HasRequests() bool {
	if o != nil && !IsNil(o.Requests) {
		return true
	}

	return false
}

// SetRequests gets a reference to the given int32 and assigns it to the Requests field.
func (o *CountryStatistics) SetRequests(v int32) {
	o.Requests = &v
}

// GetTraffics returns the Traffics field value if set, zero value otherwise.
func (o *CountryStatistics) GetTraffics() int32 {
	if o == nil || IsNil(o.Traffics) {
		var ret int32
		return ret
	}
	return *o.Traffics
}

// GetTrafficsOk returns a tuple with the Traffics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryStatistics) GetTrafficsOk() (*int32, bool) {
	if o == nil || IsNil(o.Traffics) {
		return nil, false
	}
	return o.Traffics, true
}

// HasTraffics returns a boolean if a field has been set.
func (o *CountryStatistics) HasTraffics() bool {
	if o != nil && !IsNil(o.Traffics) {
		return true
	}

	return false
}

// SetTraffics gets a reference to the given int32 and assigns it to the Traffics field.
func (o *CountryStatistics) SetTraffics(v int32) {
	o.Traffics = &v
}

func (o CountryStatistics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CountryStatistics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.Requests) {
		toSerialize["requests"] = o.Requests
	}
	if !IsNil(o.Traffics) {
		toSerialize["traffics"] = o.Traffics
	}
	return toSerialize, nil
}

type NullableCountryStatistics struct {
	value *CountryStatistics
	isSet bool
}

func (v NullableCountryStatistics) Get() *CountryStatistics {
	return v.value
}

func (v *NullableCountryStatistics) Set(val *CountryStatistics) {
	v.value = val
	v.isSet = true
}

func (v NullableCountryStatistics) IsSet() bool {
	return v.isSet
}

func (v *NullableCountryStatistics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountryStatistics(val *CountryStatistics) *NullableCountryStatistics {
	return &NullableCountryStatistics{value: val, isSet: true}
}

func (v NullableCountryStatistics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountryStatistics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


