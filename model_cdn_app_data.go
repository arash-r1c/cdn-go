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

// checks if the CdnAppData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CdnAppData{}

// CdnAppData struct for CdnAppData
type CdnAppData struct {
	Data *CdnApp `json:"data,omitempty"`
}

// NewCdnAppData instantiates a new CdnAppData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCdnAppData() *CdnAppData {
	this := CdnAppData{}
	return &this
}

// NewCdnAppDataWithDefaults instantiates a new CdnAppData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCdnAppDataWithDefaults() *CdnAppData {
	this := CdnAppData{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CdnAppData) GetData() CdnApp {
	if o == nil || IsNil(o.Data) {
		var ret CdnApp
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdnAppData) GetDataOk() (*CdnApp, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CdnAppData) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given CdnApp and assigns it to the Data field.
func (o *CdnAppData) SetData(v CdnApp) {
	o.Data = &v
}

func (o CdnAppData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CdnAppData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCdnAppData struct {
	value *CdnAppData
	isSet bool
}

func (v NullableCdnAppData) Get() *CdnAppData {
	return v.value
}

func (v *NullableCdnAppData) Set(val *CdnAppData) {
	v.value = val
	v.isSet = true
}

func (v NullableCdnAppData) IsSet() bool {
	return v.isSet
}

func (v *NullableCdnAppData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCdnAppData(val *CdnAppData) *NullableCdnAppData {
	return &NullableCdnAppData{value: val, isSet: true}
}

func (v NullableCdnAppData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCdnAppData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


