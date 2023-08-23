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

// checks if the CountryList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CountryList{}

// CountryList struct for CountryList
type CountryList struct {
	// The name of the country
	Country *string `json:"country,omitempty"`
	// The 2-letter country code
	Code *string `json:"code,omitempty"`
	// The number of requests from the country
	Requests *int32 `json:"requests,omitempty"`
	// The amount of traffic from the country
	Traffics *int32 `json:"traffics,omitempty"`
}

// NewCountryList instantiates a new CountryList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountryList() *CountryList {
	this := CountryList{}
	return &this
}

// NewCountryListWithDefaults instantiates a new CountryList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountryListWithDefaults() *CountryList {
	this := CountryList{}
	return &this
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *CountryList) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryList) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *CountryList) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *CountryList) SetCountry(v string) {
	o.Country = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CountryList) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryList) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CountryList) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *CountryList) SetCode(v string) {
	o.Code = &v
}

// GetRequests returns the Requests field value if set, zero value otherwise.
func (o *CountryList) GetRequests() int32 {
	if o == nil || IsNil(o.Requests) {
		var ret int32
		return ret
	}
	return *o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryList) GetRequestsOk() (*int32, bool) {
	if o == nil || IsNil(o.Requests) {
		return nil, false
	}
	return o.Requests, true
}

// HasRequests returns a boolean if a field has been set.
func (o *CountryList) HasRequests() bool {
	if o != nil && !IsNil(o.Requests) {
		return true
	}

	return false
}

// SetRequests gets a reference to the given int32 and assigns it to the Requests field.
func (o *CountryList) SetRequests(v int32) {
	o.Requests = &v
}

// GetTraffics returns the Traffics field value if set, zero value otherwise.
func (o *CountryList) GetTraffics() int32 {
	if o == nil || IsNil(o.Traffics) {
		var ret int32
		return ret
	}
	return *o.Traffics
}

// GetTrafficsOk returns a tuple with the Traffics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryList) GetTrafficsOk() (*int32, bool) {
	if o == nil || IsNil(o.Traffics) {
		return nil, false
	}
	return o.Traffics, true
}

// HasTraffics returns a boolean if a field has been set.
func (o *CountryList) HasTraffics() bool {
	if o != nil && !IsNil(o.Traffics) {
		return true
	}

	return false
}

// SetTraffics gets a reference to the given int32 and assigns it to the Traffics field.
func (o *CountryList) SetTraffics(v int32) {
	o.Traffics = &v
}

func (o CountryList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CountryList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Requests) {
		toSerialize["requests"] = o.Requests
	}
	if !IsNil(o.Traffics) {
		toSerialize["traffics"] = o.Traffics
	}
	return toSerialize, nil
}

type NullableCountryList struct {
	value *CountryList
	isSet bool
}

func (v NullableCountryList) Get() *CountryList {
	return v.value
}

func (v *NullableCountryList) Set(val *CountryList) {
	v.value = val
	v.isSet = true
}

func (v NullableCountryList) IsSet() bool {
	return v.isSet
}

func (v *NullableCountryList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountryList(val *CountryList) *NullableCountryList {
	return &NullableCountryList{value: val, isSet: true}
}

func (v NullableCountryList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountryList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


