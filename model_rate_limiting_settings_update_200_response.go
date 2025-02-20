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

// checks if the RateLimitingSettingsUpdate200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RateLimitingSettingsUpdate200Response{}

// RateLimitingSettingsUpdate200Response struct for RateLimitingSettingsUpdate200Response
type RateLimitingSettingsUpdate200Response struct {
	Data *RateLimitSettings `json:"data,omitempty"`
	Message NullableString `json:"message,omitempty"`
}

// NewRateLimitingSettingsUpdate200Response instantiates a new RateLimitingSettingsUpdate200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRateLimitingSettingsUpdate200Response() *RateLimitingSettingsUpdate200Response {
	this := RateLimitingSettingsUpdate200Response{}
	return &this
}

// NewRateLimitingSettingsUpdate200ResponseWithDefaults instantiates a new RateLimitingSettingsUpdate200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRateLimitingSettingsUpdate200ResponseWithDefaults() *RateLimitingSettingsUpdate200Response {
	this := RateLimitingSettingsUpdate200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RateLimitingSettingsUpdate200Response) GetData() RateLimitSettings {
	if o == nil || IsNil(o.Data) {
		var ret RateLimitSettings
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimitingSettingsUpdate200Response) GetDataOk() (*RateLimitSettings, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RateLimitingSettingsUpdate200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given RateLimitSettings and assigns it to the Data field.
func (o *RateLimitingSettingsUpdate200Response) SetData(v RateLimitSettings) {
	o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RateLimitingSettingsUpdate200Response) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RateLimitingSettingsUpdate200Response) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *RateLimitingSettingsUpdate200Response) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *RateLimitingSettingsUpdate200Response) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *RateLimitingSettingsUpdate200Response) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *RateLimitingSettingsUpdate200Response) UnsetMessage() {
	o.Message.Unset()
}

func (o RateLimitingSettingsUpdate200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RateLimitingSettingsUpdate200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	return toSerialize, nil
}

type NullableRateLimitingSettingsUpdate200Response struct {
	value *RateLimitingSettingsUpdate200Response
	isSet bool
}

func (v NullableRateLimitingSettingsUpdate200Response) Get() *RateLimitingSettingsUpdate200Response {
	return v.value
}

func (v *NullableRateLimitingSettingsUpdate200Response) Set(val *RateLimitingSettingsUpdate200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableRateLimitingSettingsUpdate200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableRateLimitingSettingsUpdate200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRateLimitingSettingsUpdate200Response(val *RateLimitingSettingsUpdate200Response) *NullableRateLimitingSettingsUpdate200Response {
	return &NullableRateLimitingSettingsUpdate200Response{value: val, isSet: true}
}

func (v NullableRateLimitingSettingsUpdate200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRateLimitingSettingsUpdate200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


