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

// checks if the DomainsNameserversDeprecatedCheck200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DomainsNameserversDeprecatedCheck200Response{}

// DomainsNameserversDeprecatedCheck200Response struct for DomainsNameserversDeprecatedCheck200Response
type DomainsNameserversDeprecatedCheck200Response struct {
	Data *DeprecatedNs `json:"data,omitempty"`
	Message NullableString `json:"message,omitempty"`
}

// NewDomainsNameserversDeprecatedCheck200Response instantiates a new DomainsNameserversDeprecatedCheck200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainsNameserversDeprecatedCheck200Response() *DomainsNameserversDeprecatedCheck200Response {
	this := DomainsNameserversDeprecatedCheck200Response{}
	return &this
}

// NewDomainsNameserversDeprecatedCheck200ResponseWithDefaults instantiates a new DomainsNameserversDeprecatedCheck200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainsNameserversDeprecatedCheck200ResponseWithDefaults() *DomainsNameserversDeprecatedCheck200Response {
	this := DomainsNameserversDeprecatedCheck200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DomainsNameserversDeprecatedCheck200Response) GetData() DeprecatedNs {
	if o == nil || IsNil(o.Data) {
		var ret DeprecatedNs
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainsNameserversDeprecatedCheck200Response) GetDataOk() (*DeprecatedNs, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DomainsNameserversDeprecatedCheck200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given DeprecatedNs and assigns it to the Data field.
func (o *DomainsNameserversDeprecatedCheck200Response) SetData(v DeprecatedNs) {
	o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DomainsNameserversDeprecatedCheck200Response) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DomainsNameserversDeprecatedCheck200Response) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *DomainsNameserversDeprecatedCheck200Response) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *DomainsNameserversDeprecatedCheck200Response) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *DomainsNameserversDeprecatedCheck200Response) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *DomainsNameserversDeprecatedCheck200Response) UnsetMessage() {
	o.Message.Unset()
}

func (o DomainsNameserversDeprecatedCheck200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DomainsNameserversDeprecatedCheck200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	return toSerialize, nil
}

type NullableDomainsNameserversDeprecatedCheck200Response struct {
	value *DomainsNameserversDeprecatedCheck200Response
	isSet bool
}

func (v NullableDomainsNameserversDeprecatedCheck200Response) Get() *DomainsNameserversDeprecatedCheck200Response {
	return v.value
}

func (v *NullableDomainsNameserversDeprecatedCheck200Response) Set(val *DomainsNameserversDeprecatedCheck200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainsNameserversDeprecatedCheck200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainsNameserversDeprecatedCheck200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainsNameserversDeprecatedCheck200Response(val *DomainsNameserversDeprecatedCheck200Response) *NullableDomainsNameserversDeprecatedCheck200Response {
	return &NullableDomainsNameserversDeprecatedCheck200Response{value: val, isSet: true}
}

func (v NullableDomainsNameserversDeprecatedCheck200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainsNameserversDeprecatedCheck200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


