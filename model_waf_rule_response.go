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

// checks if the WafRuleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WafRuleResponse{}

// WafRuleResponse struct for WafRuleResponse
type WafRuleResponse struct {
	Data *WafRule `json:"data,omitempty"`
	Message NullableString `json:"message,omitempty"`
}

// NewWafRuleResponse instantiates a new WafRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafRuleResponse() *WafRuleResponse {
	this := WafRuleResponse{}
	return &this
}

// NewWafRuleResponseWithDefaults instantiates a new WafRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafRuleResponseWithDefaults() *WafRuleResponse {
	this := WafRuleResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *WafRuleResponse) GetData() WafRule {
	if o == nil || IsNil(o.Data) {
		var ret WafRule
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRuleResponse) GetDataOk() (*WafRule, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *WafRuleResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given WafRule and assigns it to the Data field.
func (o *WafRuleResponse) SetData(v WafRule) {
	o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WafRuleResponse) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WafRuleResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *WafRuleResponse) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *WafRuleResponse) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *WafRuleResponse) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *WafRuleResponse) UnsetMessage() {
	o.Message.Unset()
}

func (o WafRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WafRuleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	return toSerialize, nil
}

type NullableWafRuleResponse struct {
	value *WafRuleResponse
	isSet bool
}

func (v NullableWafRuleResponse) Get() *WafRuleResponse {
	return v.value
}

func (v *NullableWafRuleResponse) Set(val *WafRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWafRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWafRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWafRuleResponse(val *WafRuleResponse) *NullableWafRuleResponse {
	return &NullableWafRuleResponse{value: val, isSet: true}
}

func (v NullableWafRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWafRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


