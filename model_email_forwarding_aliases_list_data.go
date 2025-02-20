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

// checks if the EmailForwardingAliasesListData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmailForwardingAliasesListData{}

// EmailForwardingAliasesListData struct for EmailForwardingAliasesListData
type EmailForwardingAliasesListData struct {
	Data []EmailForwardingAliasesListInner `json:"data,omitempty"`
}

// NewEmailForwardingAliasesListData instantiates a new EmailForwardingAliasesListData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailForwardingAliasesListData() *EmailForwardingAliasesListData {
	this := EmailForwardingAliasesListData{}
	return &this
}

// NewEmailForwardingAliasesListDataWithDefaults instantiates a new EmailForwardingAliasesListData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailForwardingAliasesListDataWithDefaults() *EmailForwardingAliasesListData {
	this := EmailForwardingAliasesListData{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *EmailForwardingAliasesListData) GetData() []EmailForwardingAliasesListInner {
	if o == nil || IsNil(o.Data) {
		var ret []EmailForwardingAliasesListInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailForwardingAliasesListData) GetDataOk() ([]EmailForwardingAliasesListInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *EmailForwardingAliasesListData) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []EmailForwardingAliasesListInner and assigns it to the Data field.
func (o *EmailForwardingAliasesListData) SetData(v []EmailForwardingAliasesListInner) {
	o.Data = v
}

func (o EmailForwardingAliasesListData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmailForwardingAliasesListData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableEmailForwardingAliasesListData struct {
	value *EmailForwardingAliasesListData
	isSet bool
}

func (v NullableEmailForwardingAliasesListData) Get() *EmailForwardingAliasesListData {
	return v.value
}

func (v *NullableEmailForwardingAliasesListData) Set(val *EmailForwardingAliasesListData) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailForwardingAliasesListData) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailForwardingAliasesListData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailForwardingAliasesListData(val *EmailForwardingAliasesListData) *NullableEmailForwardingAliasesListData {
	return &NullableEmailForwardingAliasesListData{value: val, isSet: true}
}

func (v NullableEmailForwardingAliasesListData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailForwardingAliasesListData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


