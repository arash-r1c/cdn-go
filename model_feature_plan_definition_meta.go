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

// checks if the FeaturePlanDefinitionMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeaturePlanDefinitionMeta{}

// FeaturePlanDefinitionMeta struct for FeaturePlanDefinitionMeta
type FeaturePlanDefinitionMeta struct {
	Labels []FeaturePlanDefinitionMetaLabelsInner `json:"labels,omitempty"`
	Tip *string `json:"tip,omitempty"`
	AvailableParams []map[string]interface{} `json:"available_params,omitempty"`
}

// NewFeaturePlanDefinitionMeta instantiates a new FeaturePlanDefinitionMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeaturePlanDefinitionMeta() *FeaturePlanDefinitionMeta {
	this := FeaturePlanDefinitionMeta{}
	return &this
}

// NewFeaturePlanDefinitionMetaWithDefaults instantiates a new FeaturePlanDefinitionMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeaturePlanDefinitionMetaWithDefaults() *FeaturePlanDefinitionMeta {
	this := FeaturePlanDefinitionMeta{}
	return &this
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *FeaturePlanDefinitionMeta) GetLabels() []FeaturePlanDefinitionMetaLabelsInner {
	if o == nil || IsNil(o.Labels) {
		var ret []FeaturePlanDefinitionMetaLabelsInner
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeaturePlanDefinitionMeta) GetLabelsOk() ([]FeaturePlanDefinitionMetaLabelsInner, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *FeaturePlanDefinitionMeta) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []FeaturePlanDefinitionMetaLabelsInner and assigns it to the Labels field.
func (o *FeaturePlanDefinitionMeta) SetLabels(v []FeaturePlanDefinitionMetaLabelsInner) {
	o.Labels = v
}

// GetTip returns the Tip field value if set, zero value otherwise.
func (o *FeaturePlanDefinitionMeta) GetTip() string {
	if o == nil || IsNil(o.Tip) {
		var ret string
		return ret
	}
	return *o.Tip
}

// GetTipOk returns a tuple with the Tip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeaturePlanDefinitionMeta) GetTipOk() (*string, bool) {
	if o == nil || IsNil(o.Tip) {
		return nil, false
	}
	return o.Tip, true
}

// HasTip returns a boolean if a field has been set.
func (o *FeaturePlanDefinitionMeta) HasTip() bool {
	if o != nil && !IsNil(o.Tip) {
		return true
	}

	return false
}

// SetTip gets a reference to the given string and assigns it to the Tip field.
func (o *FeaturePlanDefinitionMeta) SetTip(v string) {
	o.Tip = &v
}

// GetAvailableParams returns the AvailableParams field value if set, zero value otherwise.
func (o *FeaturePlanDefinitionMeta) GetAvailableParams() []map[string]interface{} {
	if o == nil || IsNil(o.AvailableParams) {
		var ret []map[string]interface{}
		return ret
	}
	return o.AvailableParams
}

// GetAvailableParamsOk returns a tuple with the AvailableParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeaturePlanDefinitionMeta) GetAvailableParamsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.AvailableParams) {
		return nil, false
	}
	return o.AvailableParams, true
}

// HasAvailableParams returns a boolean if a field has been set.
func (o *FeaturePlanDefinitionMeta) HasAvailableParams() bool {
	if o != nil && !IsNil(o.AvailableParams) {
		return true
	}

	return false
}

// SetAvailableParams gets a reference to the given []map[string]interface{} and assigns it to the AvailableParams field.
func (o *FeaturePlanDefinitionMeta) SetAvailableParams(v []map[string]interface{}) {
	o.AvailableParams = v
}

func (o FeaturePlanDefinitionMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeaturePlanDefinitionMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Tip) {
		toSerialize["tip"] = o.Tip
	}
	if !IsNil(o.AvailableParams) {
		toSerialize["available_params"] = o.AvailableParams
	}
	return toSerialize, nil
}

type NullableFeaturePlanDefinitionMeta struct {
	value *FeaturePlanDefinitionMeta
	isSet bool
}

func (v NullableFeaturePlanDefinitionMeta) Get() *FeaturePlanDefinitionMeta {
	return v.value
}

func (v *NullableFeaturePlanDefinitionMeta) Set(val *FeaturePlanDefinitionMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableFeaturePlanDefinitionMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableFeaturePlanDefinitionMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeaturePlanDefinitionMeta(val *FeaturePlanDefinitionMeta) *NullableFeaturePlanDefinitionMeta {
	return &NullableFeaturePlanDefinitionMeta{value: val, isSet: true}
}

func (v NullableFeaturePlanDefinitionMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeaturePlanDefinitionMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


