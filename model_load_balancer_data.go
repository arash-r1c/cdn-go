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

// checks if the LoadBalancerData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadBalancerData{}

// LoadBalancerData struct for LoadBalancerData
type LoadBalancerData struct {
	Data *LoadBalancer `json:"data,omitempty"`
}

// NewLoadBalancerData instantiates a new LoadBalancerData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancerData() *LoadBalancerData {
	this := LoadBalancerData{}
	return &this
}

// NewLoadBalancerDataWithDefaults instantiates a new LoadBalancerData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancerDataWithDefaults() *LoadBalancerData {
	this := LoadBalancerData{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *LoadBalancerData) GetData() LoadBalancer {
	if o == nil || IsNil(o.Data) {
		var ret LoadBalancer
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerData) GetDataOk() (*LoadBalancer, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *LoadBalancerData) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given LoadBalancer and assigns it to the Data field.
func (o *LoadBalancerData) SetData(v LoadBalancer) {
	o.Data = &v
}

func (o LoadBalancerData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadBalancerData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableLoadBalancerData struct {
	value *LoadBalancerData
	isSet bool
}

func (v NullableLoadBalancerData) Get() *LoadBalancerData {
	return v.value
}

func (v *NullableLoadBalancerData) Set(val *LoadBalancerData) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancerData) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancerData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancerData(val *LoadBalancerData) *NullableLoadBalancerData {
	return &NullableLoadBalancerData{value: val, isSet: true}
}

func (v NullableLoadBalancerData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancerData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


