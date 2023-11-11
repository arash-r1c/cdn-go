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

// checks if the WafPresetPackagesInnerProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WafPresetPackagesInnerProvider{}

// WafPresetPackagesInnerProvider struct for WafPresetPackagesInnerProvider
type WafPresetPackagesInnerProvider struct {
	Name *string `json:"name,omitempty"`
	Logo *string `json:"logo,omitempty"`
}

// NewWafPresetPackagesInnerProvider instantiates a new WafPresetPackagesInnerProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafPresetPackagesInnerProvider() *WafPresetPackagesInnerProvider {
	this := WafPresetPackagesInnerProvider{}
	return &this
}

// NewWafPresetPackagesInnerProviderWithDefaults instantiates a new WafPresetPackagesInnerProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafPresetPackagesInnerProviderWithDefaults() *WafPresetPackagesInnerProvider {
	this := WafPresetPackagesInnerProvider{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WafPresetPackagesInnerProvider) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafPresetPackagesInnerProvider) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WafPresetPackagesInnerProvider) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WafPresetPackagesInnerProvider) SetName(v string) {
	o.Name = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *WafPresetPackagesInnerProvider) GetLogo() string {
	if o == nil || IsNil(o.Logo) {
		var ret string
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafPresetPackagesInnerProvider) GetLogoOk() (*string, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *WafPresetPackagesInnerProvider) HasLogo() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given string and assigns it to the Logo field.
func (o *WafPresetPackagesInnerProvider) SetLogo(v string) {
	o.Logo = &v
}

func (o WafPresetPackagesInnerProvider) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WafPresetPackagesInnerProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}
	return toSerialize, nil
}

type NullableWafPresetPackagesInnerProvider struct {
	value *WafPresetPackagesInnerProvider
	isSet bool
}

func (v NullableWafPresetPackagesInnerProvider) Get() *WafPresetPackagesInnerProvider {
	return v.value
}

func (v *NullableWafPresetPackagesInnerProvider) Set(val *WafPresetPackagesInnerProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableWafPresetPackagesInnerProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableWafPresetPackagesInnerProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWafPresetPackagesInnerProvider(val *WafPresetPackagesInnerProvider) *NullableWafPresetPackagesInnerProvider {
	return &NullableWafPresetPackagesInnerProvider{value: val, isSet: true}
}

func (v NullableWafPresetPackagesInnerProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWafPresetPackagesInnerProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


