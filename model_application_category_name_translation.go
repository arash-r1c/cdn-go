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

// checks if the ApplicationCategoryNameTranslation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationCategoryNameTranslation{}

// ApplicationCategoryNameTranslation struct for ApplicationCategoryNameTranslation
type ApplicationCategoryNameTranslation struct {
	En *ApplicationCategoryNameTranslationEn `json:"en,omitempty"`
	Fa *ApplicationCategoryNameTranslationEn `json:"fa,omitempty"`
}

// NewApplicationCategoryNameTranslation instantiates a new ApplicationCategoryNameTranslation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationCategoryNameTranslation() *ApplicationCategoryNameTranslation {
	this := ApplicationCategoryNameTranslation{}
	return &this
}

// NewApplicationCategoryNameTranslationWithDefaults instantiates a new ApplicationCategoryNameTranslation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationCategoryNameTranslationWithDefaults() *ApplicationCategoryNameTranslation {
	this := ApplicationCategoryNameTranslation{}
	return &this
}

// GetEn returns the En field value if set, zero value otherwise.
func (o *ApplicationCategoryNameTranslation) GetEn() ApplicationCategoryNameTranslationEn {
	if o == nil || IsNil(o.En) {
		var ret ApplicationCategoryNameTranslationEn
		return ret
	}
	return *o.En
}

// GetEnOk returns a tuple with the En field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCategoryNameTranslation) GetEnOk() (*ApplicationCategoryNameTranslationEn, bool) {
	if o == nil || IsNil(o.En) {
		return nil, false
	}
	return o.En, true
}

// HasEn returns a boolean if a field has been set.
func (o *ApplicationCategoryNameTranslation) HasEn() bool {
	if o != nil && !IsNil(o.En) {
		return true
	}

	return false
}

// SetEn gets a reference to the given ApplicationCategoryNameTranslationEn and assigns it to the En field.
func (o *ApplicationCategoryNameTranslation) SetEn(v ApplicationCategoryNameTranslationEn) {
	o.En = &v
}

// GetFa returns the Fa field value if set, zero value otherwise.
func (o *ApplicationCategoryNameTranslation) GetFa() ApplicationCategoryNameTranslationEn {
	if o == nil || IsNil(o.Fa) {
		var ret ApplicationCategoryNameTranslationEn
		return ret
	}
	return *o.Fa
}

// GetFaOk returns a tuple with the Fa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCategoryNameTranslation) GetFaOk() (*ApplicationCategoryNameTranslationEn, bool) {
	if o == nil || IsNil(o.Fa) {
		return nil, false
	}
	return o.Fa, true
}

// HasFa returns a boolean if a field has been set.
func (o *ApplicationCategoryNameTranslation) HasFa() bool {
	if o != nil && !IsNil(o.Fa) {
		return true
	}

	return false
}

// SetFa gets a reference to the given ApplicationCategoryNameTranslationEn and assigns it to the Fa field.
func (o *ApplicationCategoryNameTranslation) SetFa(v ApplicationCategoryNameTranslationEn) {
	o.Fa = &v
}

func (o ApplicationCategoryNameTranslation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationCategoryNameTranslation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.En) {
		toSerialize["en"] = o.En
	}
	if !IsNil(o.Fa) {
		toSerialize["fa"] = o.Fa
	}
	return toSerialize, nil
}

type NullableApplicationCategoryNameTranslation struct {
	value *ApplicationCategoryNameTranslation
	isSet bool
}

func (v NullableApplicationCategoryNameTranslation) Get() *ApplicationCategoryNameTranslation {
	return v.value
}

func (v *NullableApplicationCategoryNameTranslation) Set(val *ApplicationCategoryNameTranslation) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationCategoryNameTranslation) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationCategoryNameTranslation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationCategoryNameTranslation(val *ApplicationCategoryNameTranslation) *NullableApplicationCategoryNameTranslation {
	return &NullableApplicationCategoryNameTranslation{value: val, isSet: true}
}

func (v NullableApplicationCategoryNameTranslation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationCategoryNameTranslation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


