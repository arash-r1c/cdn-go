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

// checks if the LoadBalancerStore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadBalancerStore{}

// LoadBalancerStore struct for LoadBalancerStore
type LoadBalancerStore struct {
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	Status bool `json:"status"`
	Method string `json:"method"`
}

// NewLoadBalancerStore instantiates a new LoadBalancerStore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancerStore(name string, status bool, method string) *LoadBalancerStore {
	this := LoadBalancerStore{}
	this.Name = name
	this.Status = status
	this.Method = method
	return &this
}

// NewLoadBalancerStoreWithDefaults instantiates a new LoadBalancerStore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancerStoreWithDefaults() *LoadBalancerStore {
	this := LoadBalancerStore{}
	return &this
}

// GetName returns the Name field value
func (o *LoadBalancerStore) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerStore) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LoadBalancerStore) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LoadBalancerStore) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerStore) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LoadBalancerStore) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LoadBalancerStore) SetDescription(v string) {
	o.Description = &v
}

// GetStatus returns the Status field value
func (o *LoadBalancerStore) GetStatus() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerStore) GetStatusOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *LoadBalancerStore) SetStatus(v bool) {
	o.Status = v
}

// GetMethod returns the Method field value
func (o *LoadBalancerStore) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *LoadBalancerStore) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *LoadBalancerStore) SetMethod(v string) {
	o.Method = v
}

func (o LoadBalancerStore) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadBalancerStore) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["status"] = o.Status
	toSerialize["method"] = o.Method
	return toSerialize, nil
}

type NullableLoadBalancerStore struct {
	value *LoadBalancerStore
	isSet bool
}

func (v NullableLoadBalancerStore) Get() *LoadBalancerStore {
	return v.value
}

func (v *NullableLoadBalancerStore) Set(val *LoadBalancerStore) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancerStore) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancerStore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancerStore(val *LoadBalancerStore) *NullableLoadBalancerStore {
	return &NullableLoadBalancerStore{value: val, isSet: true}
}

func (v NullableLoadBalancerStore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancerStore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


