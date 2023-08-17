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

// checks if the ChallengeAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChallengeAction{}

// ChallengeAction struct for ChallengeAction
type ChallengeAction struct {
	// The mode of mitigation (1: Cookie, 2: Javascript, 3: Captcha)
	Mode *int32 `json:"mode,omitempty"`
	Ttl *int32 `json:"ttl,omitempty"`
	HttpsOnly *bool `json:"https_only,omitempty"`
}

// NewChallengeAction instantiates a new ChallengeAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChallengeAction() *ChallengeAction {
	this := ChallengeAction{}
	return &this
}

// NewChallengeActionWithDefaults instantiates a new ChallengeAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChallengeActionWithDefaults() *ChallengeAction {
	this := ChallengeAction{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *ChallengeAction) GetMode() int32 {
	if o == nil || IsNil(o.Mode) {
		var ret int32
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChallengeAction) GetModeOk() (*int32, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *ChallengeAction) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given int32 and assigns it to the Mode field.
func (o *ChallengeAction) SetMode(v int32) {
	o.Mode = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *ChallengeAction) GetTtl() int32 {
	if o == nil || IsNil(o.Ttl) {
		var ret int32
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChallengeAction) GetTtlOk() (*int32, bool) {
	if o == nil || IsNil(o.Ttl) {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *ChallengeAction) HasTtl() bool {
	if o != nil && !IsNil(o.Ttl) {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int32 and assigns it to the Ttl field.
func (o *ChallengeAction) SetTtl(v int32) {
	o.Ttl = &v
}

// GetHttpsOnly returns the HttpsOnly field value if set, zero value otherwise.
func (o *ChallengeAction) GetHttpsOnly() bool {
	if o == nil || IsNil(o.HttpsOnly) {
		var ret bool
		return ret
	}
	return *o.HttpsOnly
}

// GetHttpsOnlyOk returns a tuple with the HttpsOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChallengeAction) GetHttpsOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.HttpsOnly) {
		return nil, false
	}
	return o.HttpsOnly, true
}

// HasHttpsOnly returns a boolean if a field has been set.
func (o *ChallengeAction) HasHttpsOnly() bool {
	if o != nil && !IsNil(o.HttpsOnly) {
		return true
	}

	return false
}

// SetHttpsOnly gets a reference to the given bool and assigns it to the HttpsOnly field.
func (o *ChallengeAction) SetHttpsOnly(v bool) {
	o.HttpsOnly = &v
}

func (o ChallengeAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChallengeAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.Ttl) {
		toSerialize["ttl"] = o.Ttl
	}
	if !IsNil(o.HttpsOnly) {
		toSerialize["https_only"] = o.HttpsOnly
	}
	return toSerialize, nil
}

type NullableChallengeAction struct {
	value *ChallengeAction
	isSet bool
}

func (v NullableChallengeAction) Get() *ChallengeAction {
	return v.value
}

func (v *NullableChallengeAction) Set(val *ChallengeAction) {
	v.value = val
	v.isSet = true
}

func (v NullableChallengeAction) IsSet() bool {
	return v.isSet
}

func (v *NullableChallengeAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChallengeAction(val *ChallengeAction) *NullableChallengeAction {
	return &NullableChallengeAction{value: val, isSet: true}
}

func (v NullableChallengeAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChallengeAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


