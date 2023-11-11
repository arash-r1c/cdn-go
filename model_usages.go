/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.114.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package arvancloud

import (
	"encoding/json"
)

// checks if the Usages type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Usages{}

// Usages struct for Usages
type Usages struct {
	FeatureUsages []FeatureUsage `json:"feature_usages,omitempty"`
	EstimatedCost *EstimatedCost `json:"estimated_cost,omitempty"`
}

// NewUsages instantiates a new Usages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsages() *Usages {
	this := Usages{}
	return &this
}

// NewUsagesWithDefaults instantiates a new Usages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsagesWithDefaults() *Usages {
	this := Usages{}
	return &this
}

// GetFeatureUsages returns the FeatureUsages field value if set, zero value otherwise.
func (o *Usages) GetFeatureUsages() []FeatureUsage {
	if o == nil || IsNil(o.FeatureUsages) {
		var ret []FeatureUsage
		return ret
	}
	return o.FeatureUsages
}

// GetFeatureUsagesOk returns a tuple with the FeatureUsages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Usages) GetFeatureUsagesOk() ([]FeatureUsage, bool) {
	if o == nil || IsNil(o.FeatureUsages) {
		return nil, false
	}
	return o.FeatureUsages, true
}

// HasFeatureUsages returns a boolean if a field has been set.
func (o *Usages) HasFeatureUsages() bool {
	if o != nil && !IsNil(o.FeatureUsages) {
		return true
	}

	return false
}

// SetFeatureUsages gets a reference to the given []FeatureUsage and assigns it to the FeatureUsages field.
func (o *Usages) SetFeatureUsages(v []FeatureUsage) {
	o.FeatureUsages = v
}

// GetEstimatedCost returns the EstimatedCost field value if set, zero value otherwise.
func (o *Usages) GetEstimatedCost() EstimatedCost {
	if o == nil || IsNil(o.EstimatedCost) {
		var ret EstimatedCost
		return ret
	}
	return *o.EstimatedCost
}

// GetEstimatedCostOk returns a tuple with the EstimatedCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Usages) GetEstimatedCostOk() (*EstimatedCost, bool) {
	if o == nil || IsNil(o.EstimatedCost) {
		return nil, false
	}
	return o.EstimatedCost, true
}

// HasEstimatedCost returns a boolean if a field has been set.
func (o *Usages) HasEstimatedCost() bool {
	if o != nil && !IsNil(o.EstimatedCost) {
		return true
	}

	return false
}

// SetEstimatedCost gets a reference to the given EstimatedCost and assigns it to the EstimatedCost field.
func (o *Usages) SetEstimatedCost(v EstimatedCost) {
	o.EstimatedCost = &v
}

func (o Usages) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Usages) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FeatureUsages) {
		toSerialize["feature_usages"] = o.FeatureUsages
	}
	if !IsNil(o.EstimatedCost) {
		toSerialize["estimated_cost"] = o.EstimatedCost
	}
	return toSerialize, nil
}

type NullableUsages struct {
	value *Usages
	isSet bool
}

func (v NullableUsages) Get() *Usages {
	return v.value
}

func (v *NullableUsages) Set(val *Usages) {
	v.value = val
	v.isSet = true
}

func (v NullableUsages) IsSet() bool {
	return v.isSet
}

func (v *NullableUsages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsages(val *Usages) *NullableUsages {
	return &NullableUsages{value: val, isSet: true}
}

func (v NullableUsages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


