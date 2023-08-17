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

// checks if the DomainWafPackagesData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DomainWafPackagesData{}

// DomainWafPackagesData struct for DomainWafPackagesData
type DomainWafPackagesData struct {
	Data []DomainWafPackage `json:"data,omitempty"`
}

// NewDomainWafPackagesData instantiates a new DomainWafPackagesData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainWafPackagesData() *DomainWafPackagesData {
	this := DomainWafPackagesData{}
	return &this
}

// NewDomainWafPackagesDataWithDefaults instantiates a new DomainWafPackagesData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainWafPackagesDataWithDefaults() *DomainWafPackagesData {
	this := DomainWafPackagesData{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DomainWafPackagesData) GetData() []DomainWafPackage {
	if o == nil || IsNil(o.Data) {
		var ret []DomainWafPackage
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainWafPackagesData) GetDataOk() ([]DomainWafPackage, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DomainWafPackagesData) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []DomainWafPackage and assigns it to the Data field.
func (o *DomainWafPackagesData) SetData(v []DomainWafPackage) {
	o.Data = v
}

func (o DomainWafPackagesData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DomainWafPackagesData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableDomainWafPackagesData struct {
	value *DomainWafPackagesData
	isSet bool
}

func (v NullableDomainWafPackagesData) Get() *DomainWafPackagesData {
	return v.value
}

func (v *NullableDomainWafPackagesData) Set(val *DomainWafPackagesData) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainWafPackagesData) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainWafPackagesData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainWafPackagesData(val *DomainWafPackagesData) *NullableDomainWafPackagesData {
	return &NullableDomainWafPackagesData{value: val, isSet: true}
}

func (v NullableDomainWafPackagesData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainWafPackagesData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


