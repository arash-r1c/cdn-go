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

// checks if the FeaturePlanDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeaturePlanDefinition{}

// FeaturePlanDefinition struct for FeaturePlanDefinition
type FeaturePlanDefinition struct {
	Meta *FeaturePlanDefinitionMeta `json:"meta,omitempty"`
	UsageLimit NullableUsageLimit `json:"usage_limit,omitempty"`
	Pricing NullableFeaturePricing `json:"pricing,omitempty"`
}

// NewFeaturePlanDefinition instantiates a new FeaturePlanDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeaturePlanDefinition() *FeaturePlanDefinition {
	this := FeaturePlanDefinition{}
	return &this
}

// NewFeaturePlanDefinitionWithDefaults instantiates a new FeaturePlanDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeaturePlanDefinitionWithDefaults() *FeaturePlanDefinition {
	this := FeaturePlanDefinition{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *FeaturePlanDefinition) GetMeta() FeaturePlanDefinitionMeta {
	if o == nil || IsNil(o.Meta) {
		var ret FeaturePlanDefinitionMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeaturePlanDefinition) GetMetaOk() (*FeaturePlanDefinitionMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *FeaturePlanDefinition) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given FeaturePlanDefinitionMeta and assigns it to the Meta field.
func (o *FeaturePlanDefinition) SetMeta(v FeaturePlanDefinitionMeta) {
	o.Meta = &v
}

// GetUsageLimit returns the UsageLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeaturePlanDefinition) GetUsageLimit() UsageLimit {
	if o == nil || IsNil(o.UsageLimit.Get()) {
		var ret UsageLimit
		return ret
	}
	return *o.UsageLimit.Get()
}

// GetUsageLimitOk returns a tuple with the UsageLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeaturePlanDefinition) GetUsageLimitOk() (*UsageLimit, bool) {
	if o == nil {
		return nil, false
	}
	return o.UsageLimit.Get(), o.UsageLimit.IsSet()
}

// HasUsageLimit returns a boolean if a field has been set.
func (o *FeaturePlanDefinition) HasUsageLimit() bool {
	if o != nil && o.UsageLimit.IsSet() {
		return true
	}

	return false
}

// SetUsageLimit gets a reference to the given NullableUsageLimit and assigns it to the UsageLimit field.
func (o *FeaturePlanDefinition) SetUsageLimit(v UsageLimit) {
	o.UsageLimit.Set(&v)
}
// SetUsageLimitNil sets the value for UsageLimit to be an explicit nil
func (o *FeaturePlanDefinition) SetUsageLimitNil() {
	o.UsageLimit.Set(nil)
}

// UnsetUsageLimit ensures that no value is present for UsageLimit, not even an explicit nil
func (o *FeaturePlanDefinition) UnsetUsageLimit() {
	o.UsageLimit.Unset()
}

// GetPricing returns the Pricing field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeaturePlanDefinition) GetPricing() FeaturePricing {
	if o == nil || IsNil(o.Pricing.Get()) {
		var ret FeaturePricing
		return ret
	}
	return *o.Pricing.Get()
}

// GetPricingOk returns a tuple with the Pricing field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeaturePlanDefinition) GetPricingOk() (*FeaturePricing, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pricing.Get(), o.Pricing.IsSet()
}

// HasPricing returns a boolean if a field has been set.
func (o *FeaturePlanDefinition) HasPricing() bool {
	if o != nil && o.Pricing.IsSet() {
		return true
	}

	return false
}

// SetPricing gets a reference to the given NullableFeaturePricing and assigns it to the Pricing field.
func (o *FeaturePlanDefinition) SetPricing(v FeaturePricing) {
	o.Pricing.Set(&v)
}
// SetPricingNil sets the value for Pricing to be an explicit nil
func (o *FeaturePlanDefinition) SetPricingNil() {
	o.Pricing.Set(nil)
}

// UnsetPricing ensures that no value is present for Pricing, not even an explicit nil
func (o *FeaturePlanDefinition) UnsetPricing() {
	o.Pricing.Unset()
}

func (o FeaturePlanDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeaturePlanDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if o.UsageLimit.IsSet() {
		toSerialize["usage_limit"] = o.UsageLimit.Get()
	}
	if o.Pricing.IsSet() {
		toSerialize["pricing"] = o.Pricing.Get()
	}
	return toSerialize, nil
}

type NullableFeaturePlanDefinition struct {
	value *FeaturePlanDefinition
	isSet bool
}

func (v NullableFeaturePlanDefinition) Get() *FeaturePlanDefinition {
	return v.value
}

func (v *NullableFeaturePlanDefinition) Set(val *FeaturePlanDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableFeaturePlanDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableFeaturePlanDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeaturePlanDefinition(val *FeaturePlanDefinition) *NullableFeaturePlanDefinition {
	return &NullableFeaturePlanDefinition{value: val, isSet: true}
}

func (v NullableFeaturePlanDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeaturePlanDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


