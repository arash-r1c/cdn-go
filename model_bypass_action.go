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

// checks if the BypassAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BypassAction{}

// BypassAction struct for BypassAction
type BypassAction struct {
	Rlimit *bool `json:"rlimit,omitempty"`
	Challenge *bool `json:"challenge,omitempty"`
	Waf *bool `json:"waf,omitempty"`
}

// NewBypassAction instantiates a new BypassAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBypassAction() *BypassAction {
	this := BypassAction{}
	return &this
}

// NewBypassActionWithDefaults instantiates a new BypassAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBypassActionWithDefaults() *BypassAction {
	this := BypassAction{}
	return &this
}

// GetRlimit returns the Rlimit field value if set, zero value otherwise.
func (o *BypassAction) GetRlimit() bool {
	if o == nil || IsNil(o.Rlimit) {
		var ret bool
		return ret
	}
	return *o.Rlimit
}

// GetRlimitOk returns a tuple with the Rlimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BypassAction) GetRlimitOk() (*bool, bool) {
	if o == nil || IsNil(o.Rlimit) {
		return nil, false
	}
	return o.Rlimit, true
}

// HasRlimit returns a boolean if a field has been set.
func (o *BypassAction) HasRlimit() bool {
	if o != nil && !IsNil(o.Rlimit) {
		return true
	}

	return false
}

// SetRlimit gets a reference to the given bool and assigns it to the Rlimit field.
func (o *BypassAction) SetRlimit(v bool) {
	o.Rlimit = &v
}

// GetChallenge returns the Challenge field value if set, zero value otherwise.
func (o *BypassAction) GetChallenge() bool {
	if o == nil || IsNil(o.Challenge) {
		var ret bool
		return ret
	}
	return *o.Challenge
}

// GetChallengeOk returns a tuple with the Challenge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BypassAction) GetChallengeOk() (*bool, bool) {
	if o == nil || IsNil(o.Challenge) {
		return nil, false
	}
	return o.Challenge, true
}

// HasChallenge returns a boolean if a field has been set.
func (o *BypassAction) HasChallenge() bool {
	if o != nil && !IsNil(o.Challenge) {
		return true
	}

	return false
}

// SetChallenge gets a reference to the given bool and assigns it to the Challenge field.
func (o *BypassAction) SetChallenge(v bool) {
	o.Challenge = &v
}

// GetWaf returns the Waf field value if set, zero value otherwise.
func (o *BypassAction) GetWaf() bool {
	if o == nil || IsNil(o.Waf) {
		var ret bool
		return ret
	}
	return *o.Waf
}

// GetWafOk returns a tuple with the Waf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BypassAction) GetWafOk() (*bool, bool) {
	if o == nil || IsNil(o.Waf) {
		return nil, false
	}
	return o.Waf, true
}

// HasWaf returns a boolean if a field has been set.
func (o *BypassAction) HasWaf() bool {
	if o != nil && !IsNil(o.Waf) {
		return true
	}

	return false
}

// SetWaf gets a reference to the given bool and assigns it to the Waf field.
func (o *BypassAction) SetWaf(v bool) {
	o.Waf = &v
}

func (o BypassAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BypassAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rlimit) {
		toSerialize["rlimit"] = o.Rlimit
	}
	if !IsNil(o.Challenge) {
		toSerialize["challenge"] = o.Challenge
	}
	if !IsNil(o.Waf) {
		toSerialize["waf"] = o.Waf
	}
	return toSerialize, nil
}

type NullableBypassAction struct {
	value *BypassAction
	isSet bool
}

func (v NullableBypassAction) Get() *BypassAction {
	return v.value
}

func (v *NullableBypassAction) Set(val *BypassAction) {
	v.value = val
	v.isSet = true
}

func (v NullableBypassAction) IsSet() bool {
	return v.isSet
}

func (v *NullableBypassAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBypassAction(val *BypassAction) *NullableBypassAction {
	return &NullableBypassAction{value: val, isSet: true}
}

func (v NullableBypassAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBypassAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


