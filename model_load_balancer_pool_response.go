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

// checks if the LoadBalancerPoolResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadBalancerPoolResponse{}

// LoadBalancerPoolResponse struct for LoadBalancerPoolResponse
type LoadBalancerPoolResponse struct {
	Data []LoadBalancerPool `json:"data,omitempty"`
	Message NullableString `json:"message,omitempty"`
}

// NewLoadBalancerPoolResponse instantiates a new LoadBalancerPoolResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancerPoolResponse() *LoadBalancerPoolResponse {
	this := LoadBalancerPoolResponse{}
	return &this
}

// NewLoadBalancerPoolResponseWithDefaults instantiates a new LoadBalancerPoolResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancerPoolResponseWithDefaults() *LoadBalancerPoolResponse {
	this := LoadBalancerPoolResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *LoadBalancerPoolResponse) GetData() []LoadBalancerPool {
	if o == nil || IsNil(o.Data) {
		var ret []LoadBalancerPool
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerPoolResponse) GetDataOk() ([]LoadBalancerPool, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *LoadBalancerPoolResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []LoadBalancerPool and assigns it to the Data field.
func (o *LoadBalancerPoolResponse) SetData(v []LoadBalancerPool) {
	o.Data = v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoadBalancerPoolResponse) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoadBalancerPoolResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *LoadBalancerPoolResponse) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *LoadBalancerPoolResponse) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *LoadBalancerPoolResponse) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *LoadBalancerPoolResponse) UnsetMessage() {
	o.Message.Unset()
}

func (o LoadBalancerPoolResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadBalancerPoolResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	return toSerialize, nil
}

type NullableLoadBalancerPoolResponse struct {
	value *LoadBalancerPoolResponse
	isSet bool
}

func (v NullableLoadBalancerPoolResponse) Get() *LoadBalancerPoolResponse {
	return v.value
}

func (v *NullableLoadBalancerPoolResponse) Set(val *LoadBalancerPoolResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancerPoolResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancerPoolResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancerPoolResponse(val *LoadBalancerPoolResponse) *NullableLoadBalancerPoolResponse {
	return &NullableLoadBalancerPoolResponse{value: val, isSet: true}
}

func (v NullableLoadBalancerPoolResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancerPoolResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


