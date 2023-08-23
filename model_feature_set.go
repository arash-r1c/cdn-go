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

// checks if the FeatureSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeatureSet{}

// FeatureSet struct for FeatureSet
type FeatureSet struct {
	Id *string `json:"id,omitempty"`
	Label *string `json:"label,omitempty"`
	Features []FeatureDefinition `json:"features,omitempty"`
}

// NewFeatureSet instantiates a new FeatureSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureSet() *FeatureSet {
	this := FeatureSet{}
	return &this
}

// NewFeatureSetWithDefaults instantiates a new FeatureSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureSetWithDefaults() *FeatureSet {
	this := FeatureSet{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FeatureSet) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureSet) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FeatureSet) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FeatureSet) SetId(v string) {
	o.Id = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *FeatureSet) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureSet) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *FeatureSet) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *FeatureSet) SetLabel(v string) {
	o.Label = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *FeatureSet) GetFeatures() []FeatureDefinition {
	if o == nil || IsNil(o.Features) {
		var ret []FeatureDefinition
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureSet) GetFeaturesOk() ([]FeatureDefinition, bool) {
	if o == nil || IsNil(o.Features) {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *FeatureSet) HasFeatures() bool {
	if o != nil && !IsNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []FeatureDefinition and assigns it to the Features field.
func (o *FeatureSet) SetFeatures(v []FeatureDefinition) {
	o.Features = v
}

func (o FeatureSet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeatureSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Features) {
		toSerialize["features"] = o.Features
	}
	return toSerialize, nil
}

type NullableFeatureSet struct {
	value *FeatureSet
	isSet bool
}

func (v NullableFeatureSet) Get() *FeatureSet {
	return v.value
}

func (v *NullableFeatureSet) Set(val *FeatureSet) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureSet) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureSet(val *FeatureSet) *NullableFeatureSet {
	return &NullableFeatureSet{value: val, isSet: true}
}

func (v NullableFeatureSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatureSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


