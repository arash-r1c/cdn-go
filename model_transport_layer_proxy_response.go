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

// checks if the TransportLayerProxyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransportLayerProxyResponse{}

// TransportLayerProxyResponse struct for TransportLayerProxyResponse
type TransportLayerProxyResponse struct {
	Data *TransportLayerProxy `json:"data,omitempty"`
	Message NullableString `json:"message,omitempty"`
}

// NewTransportLayerProxyResponse instantiates a new TransportLayerProxyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransportLayerProxyResponse() *TransportLayerProxyResponse {
	this := TransportLayerProxyResponse{}
	return &this
}

// NewTransportLayerProxyResponseWithDefaults instantiates a new TransportLayerProxyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransportLayerProxyResponseWithDefaults() *TransportLayerProxyResponse {
	this := TransportLayerProxyResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TransportLayerProxyResponse) GetData() TransportLayerProxy {
	if o == nil || IsNil(o.Data) {
		var ret TransportLayerProxy
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportLayerProxyResponse) GetDataOk() (*TransportLayerProxy, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TransportLayerProxyResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given TransportLayerProxy and assigns it to the Data field.
func (o *TransportLayerProxyResponse) SetData(v TransportLayerProxy) {
	o.Data = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransportLayerProxyResponse) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransportLayerProxyResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *TransportLayerProxyResponse) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *TransportLayerProxyResponse) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *TransportLayerProxyResponse) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *TransportLayerProxyResponse) UnsetMessage() {
	o.Message.Unset()
}

func (o TransportLayerProxyResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransportLayerProxyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	return toSerialize, nil
}

type NullableTransportLayerProxyResponse struct {
	value *TransportLayerProxyResponse
	isSet bool
}

func (v NullableTransportLayerProxyResponse) Get() *TransportLayerProxyResponse {
	return v.value
}

func (v *NullableTransportLayerProxyResponse) Set(val *TransportLayerProxyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransportLayerProxyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransportLayerProxyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransportLayerProxyResponse(val *TransportLayerProxyResponse) *NullableTransportLayerProxyResponse {
	return &NullableTransportLayerProxyResponse{value: val, isSet: true}
}

func (v NullableTransportLayerProxyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransportLayerProxyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


