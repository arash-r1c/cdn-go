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

// checks if the FirewallSettingsView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FirewallSettingsView{}

// FirewallSettingsView struct for FirewallSettingsView
type FirewallSettingsView struct {
	DefaultActionDetails map[string]interface{} `json:"default_action_details,omitempty"`
	IsEnabled *bool `json:"is_enabled,omitempty"`
	DefaultAction *string `json:"default_action,omitempty"`
	// True to verify that SNI and hostname are equal
	VerifySni *bool `json:"verify_sni,omitempty"`
	// Shows hether global whitelist should be skipped for the domain or not
	SkipGlobalWhitelist NullableBool `json:"skip_global_whitelist,omitempty"`
	// Shows whether global firewall should be skipped for the domain or not
	SkipGlobalFirewall NullableBool `json:"skip_global_firewall,omitempty"`
}

// NewFirewallSettingsView instantiates a new FirewallSettingsView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirewallSettingsView() *FirewallSettingsView {
	this := FirewallSettingsView{}
	var verifySni bool = false
	this.VerifySni = &verifySni
	var skipGlobalWhitelist bool = false
	this.SkipGlobalWhitelist = *NewNullableBool(&skipGlobalWhitelist)
	var skipGlobalFirewall bool = false
	this.SkipGlobalFirewall = *NewNullableBool(&skipGlobalFirewall)
	return &this
}

// NewFirewallSettingsViewWithDefaults instantiates a new FirewallSettingsView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirewallSettingsViewWithDefaults() *FirewallSettingsView {
	this := FirewallSettingsView{}
	var verifySni bool = false
	this.VerifySni = &verifySni
	var skipGlobalWhitelist bool = false
	this.SkipGlobalWhitelist = *NewNullableBool(&skipGlobalWhitelist)
	var skipGlobalFirewall bool = false
	this.SkipGlobalFirewall = *NewNullableBool(&skipGlobalFirewall)
	return &this
}

// GetDefaultActionDetails returns the DefaultActionDetails field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirewallSettingsView) GetDefaultActionDetails() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.DefaultActionDetails
}

// GetDefaultActionDetailsOk returns a tuple with the DefaultActionDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirewallSettingsView) GetDefaultActionDetailsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.DefaultActionDetails) {
		return map[string]interface{}{}, false
	}
	return o.DefaultActionDetails, true
}

// HasDefaultActionDetails returns a boolean if a field has been set.
func (o *FirewallSettingsView) HasDefaultActionDetails() bool {
	if o != nil && IsNil(o.DefaultActionDetails) {
		return true
	}

	return false
}

// SetDefaultActionDetails gets a reference to the given map[string]interface{} and assigns it to the DefaultActionDetails field.
func (o *FirewallSettingsView) SetDefaultActionDetails(v map[string]interface{}) {
	o.DefaultActionDetails = v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *FirewallSettingsView) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallSettingsView) GetIsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabled) {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *FirewallSettingsView) HasIsEnabled() bool {
	if o != nil && !IsNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *FirewallSettingsView) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetDefaultAction returns the DefaultAction field value if set, zero value otherwise.
func (o *FirewallSettingsView) GetDefaultAction() string {
	if o == nil || IsNil(o.DefaultAction) {
		var ret string
		return ret
	}
	return *o.DefaultAction
}

// GetDefaultActionOk returns a tuple with the DefaultAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallSettingsView) GetDefaultActionOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultAction) {
		return nil, false
	}
	return o.DefaultAction, true
}

// HasDefaultAction returns a boolean if a field has been set.
func (o *FirewallSettingsView) HasDefaultAction() bool {
	if o != nil && !IsNil(o.DefaultAction) {
		return true
	}

	return false
}

// SetDefaultAction gets a reference to the given string and assigns it to the DefaultAction field.
func (o *FirewallSettingsView) SetDefaultAction(v string) {
	o.DefaultAction = &v
}

// GetVerifySni returns the VerifySni field value if set, zero value otherwise.
func (o *FirewallSettingsView) GetVerifySni() bool {
	if o == nil || IsNil(o.VerifySni) {
		var ret bool
		return ret
	}
	return *o.VerifySni
}

// GetVerifySniOk returns a tuple with the VerifySni field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallSettingsView) GetVerifySniOk() (*bool, bool) {
	if o == nil || IsNil(o.VerifySni) {
		return nil, false
	}
	return o.VerifySni, true
}

// HasVerifySni returns a boolean if a field has been set.
func (o *FirewallSettingsView) HasVerifySni() bool {
	if o != nil && !IsNil(o.VerifySni) {
		return true
	}

	return false
}

// SetVerifySni gets a reference to the given bool and assigns it to the VerifySni field.
func (o *FirewallSettingsView) SetVerifySni(v bool) {
	o.VerifySni = &v
}

// GetSkipGlobalWhitelist returns the SkipGlobalWhitelist field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirewallSettingsView) GetSkipGlobalWhitelist() bool {
	if o == nil || IsNil(o.SkipGlobalWhitelist.Get()) {
		var ret bool
		return ret
	}
	return *o.SkipGlobalWhitelist.Get()
}

// GetSkipGlobalWhitelistOk returns a tuple with the SkipGlobalWhitelist field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirewallSettingsView) GetSkipGlobalWhitelistOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SkipGlobalWhitelist.Get(), o.SkipGlobalWhitelist.IsSet()
}

// HasSkipGlobalWhitelist returns a boolean if a field has been set.
func (o *FirewallSettingsView) HasSkipGlobalWhitelist() bool {
	if o != nil && o.SkipGlobalWhitelist.IsSet() {
		return true
	}

	return false
}

// SetSkipGlobalWhitelist gets a reference to the given NullableBool and assigns it to the SkipGlobalWhitelist field.
func (o *FirewallSettingsView) SetSkipGlobalWhitelist(v bool) {
	o.SkipGlobalWhitelist.Set(&v)
}
// SetSkipGlobalWhitelistNil sets the value for SkipGlobalWhitelist to be an explicit nil
func (o *FirewallSettingsView) SetSkipGlobalWhitelistNil() {
	o.SkipGlobalWhitelist.Set(nil)
}

// UnsetSkipGlobalWhitelist ensures that no value is present for SkipGlobalWhitelist, not even an explicit nil
func (o *FirewallSettingsView) UnsetSkipGlobalWhitelist() {
	o.SkipGlobalWhitelist.Unset()
}

// GetSkipGlobalFirewall returns the SkipGlobalFirewall field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirewallSettingsView) GetSkipGlobalFirewall() bool {
	if o == nil || IsNil(o.SkipGlobalFirewall.Get()) {
		var ret bool
		return ret
	}
	return *o.SkipGlobalFirewall.Get()
}

// GetSkipGlobalFirewallOk returns a tuple with the SkipGlobalFirewall field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirewallSettingsView) GetSkipGlobalFirewallOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SkipGlobalFirewall.Get(), o.SkipGlobalFirewall.IsSet()
}

// HasSkipGlobalFirewall returns a boolean if a field has been set.
func (o *FirewallSettingsView) HasSkipGlobalFirewall() bool {
	if o != nil && o.SkipGlobalFirewall.IsSet() {
		return true
	}

	return false
}

// SetSkipGlobalFirewall gets a reference to the given NullableBool and assigns it to the SkipGlobalFirewall field.
func (o *FirewallSettingsView) SetSkipGlobalFirewall(v bool) {
	o.SkipGlobalFirewall.Set(&v)
}
// SetSkipGlobalFirewallNil sets the value for SkipGlobalFirewall to be an explicit nil
func (o *FirewallSettingsView) SetSkipGlobalFirewallNil() {
	o.SkipGlobalFirewall.Set(nil)
}

// UnsetSkipGlobalFirewall ensures that no value is present for SkipGlobalFirewall, not even an explicit nil
func (o *FirewallSettingsView) UnsetSkipGlobalFirewall() {
	o.SkipGlobalFirewall.Unset()
}

func (o FirewallSettingsView) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FirewallSettingsView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultActionDetails != nil {
		toSerialize["default_action_details"] = o.DefaultActionDetails
	}
	// skip: is_enabled is readOnly
	if !IsNil(o.DefaultAction) {
		toSerialize["default_action"] = o.DefaultAction
	}
	if !IsNil(o.VerifySni) {
		toSerialize["verify_sni"] = o.VerifySni
	}
	if o.SkipGlobalWhitelist.IsSet() {
		toSerialize["skip_global_whitelist"] = o.SkipGlobalWhitelist.Get()
	}
	if o.SkipGlobalFirewall.IsSet() {
		toSerialize["skip_global_firewall"] = o.SkipGlobalFirewall.Get()
	}
	return toSerialize, nil
}

type NullableFirewallSettingsView struct {
	value *FirewallSettingsView
	isSet bool
}

func (v NullableFirewallSettingsView) Get() *FirewallSettingsView {
	return v.value
}

func (v *NullableFirewallSettingsView) Set(val *FirewallSettingsView) {
	v.value = val
	v.isSet = true
}

func (v NullableFirewallSettingsView) IsSet() bool {
	return v.isSet
}

func (v *NullableFirewallSettingsView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirewallSettingsView(val *FirewallSettingsView) *NullableFirewallSettingsView {
	return &NullableFirewallSettingsView{value: val, isSet: true}
}

func (v NullableFirewallSettingsView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirewallSettingsView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


