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

// checks if the LoadBalancerOriginData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadBalancerOriginData{}

// LoadBalancerOriginData struct for LoadBalancerOriginData
type LoadBalancerOriginData struct {
	Data *LoadBalancerOrigin `json:"data,omitempty"`
}

// NewLoadBalancerOriginData instantiates a new LoadBalancerOriginData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancerOriginData() *LoadBalancerOriginData {
	this := LoadBalancerOriginData{}
	return &this
}

// NewLoadBalancerOriginDataWithDefaults instantiates a new LoadBalancerOriginData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancerOriginDataWithDefaults() *LoadBalancerOriginData {
	this := LoadBalancerOriginData{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *LoadBalancerOriginData) GetData() LoadBalancerOrigin {
	if o == nil || IsNil(o.Data) {
		var ret LoadBalancerOrigin
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerOriginData) GetDataOk() (*LoadBalancerOrigin, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *LoadBalancerOriginData) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given LoadBalancerOrigin and assigns it to the Data field.
func (o *LoadBalancerOriginData) SetData(v LoadBalancerOrigin) {
	o.Data = &v
}

func (o LoadBalancerOriginData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadBalancerOriginData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableLoadBalancerOriginData struct {
	value *LoadBalancerOriginData
	isSet bool
}

func (v NullableLoadBalancerOriginData) Get() *LoadBalancerOriginData {
	return v.value
}

func (v *NullableLoadBalancerOriginData) Set(val *LoadBalancerOriginData) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancerOriginData) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancerOriginData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancerOriginData(val *LoadBalancerOriginData) *NullableLoadBalancerOriginData {
	return &NullableLoadBalancerOriginData{value: val, isSet: true}
}

func (v NullableLoadBalancerOriginData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancerOriginData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


