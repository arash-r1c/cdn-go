/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.105.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package arvancloud

import (
	"encoding/json"
)

// checks if the DomainsAppsStore200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DomainsAppsStore200Response{}

// DomainsAppsStore200Response struct for DomainsAppsStore200Response
type DomainsAppsStore200Response struct {
	Data *DomainCdnApp `json:"data,omitempty"`
	Message NullableString `json:"message,omitempty"`
}

// NewDomainsAppsStore200Response instantiates a new DomainsAppsStore200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainsAppsStore200Response() *DomainsAppsStore200Response {
	this := DomainsAppsStore200Response{}
	return &this
}

// NewDomainsAppsStore200ResponseWithDefaults instantiates a new DomainsAppsStore200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainsAppsStore200ResponseWithDefaults() *DomainsAppsStore200Response {
	this := DomainsAppsStore200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DomainsAppsStore200Response) GetData() DomainCdnApp {
	if o == nil || IsNil(o.Data) {
		var ret DomainCdnApp
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainsAppsStore200Response) GetDataOk() (*DomainCdnApp, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DomainsAppsStore200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given DomainCdnApp and assigns it to the Data field.
func (o *DomainsAppsStore200Response) SetData(v DomainCdnApp) {
	o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DomainsAppsStore200Response) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DomainsAppsStore200Response) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *DomainsAppsStore200Response) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *DomainsAppsStore200Response) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *DomainsAppsStore200Response) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *DomainsAppsStore200Response) UnsetMessage() {
	o.Message.Unset()
}

func (o DomainsAppsStore200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DomainsAppsStore200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	return toSerialize, nil
}

type NullableDomainsAppsStore200Response struct {
	value *DomainsAppsStore200Response
	isSet bool
}

func (v NullableDomainsAppsStore200Response) Get() *DomainsAppsStore200Response {
	return v.value
}

func (v *NullableDomainsAppsStore200Response) Set(val *DomainsAppsStore200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainsAppsStore200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainsAppsStore200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainsAppsStore200Response(val *DomainsAppsStore200Response) *NullableDomainsAppsStore200Response {
	return &NullableDomainsAppsStore200Response{value: val, isSet: true}
}

func (v NullableDomainsAppsStore200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainsAppsStore200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


