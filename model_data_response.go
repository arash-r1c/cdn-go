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

// checks if the DataResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataResponse{}

// DataResponse struct for DataResponse
type DataResponse struct {
	Data map[string]interface{} `json:"data,omitempty"`
}

// NewDataResponse instantiates a new DataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataResponse() *DataResponse {
	this := DataResponse{}
	return &this
}

// NewDataResponseWithDefaults instantiates a new DataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataResponseWithDefaults() *DataResponse {
	this := DataResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DataResponse) GetData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DataResponse) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DataResponse) HasData() bool {
	if o != nil && IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *DataResponse) SetData(v map[string]interface{}) {
	o.Data = v
}

func (o DataResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableDataResponse struct {
	value *DataResponse
	isSet bool
}

func (v NullableDataResponse) Get() *DataResponse {
	return v.value
}

func (v *NullableDataResponse) Set(val *DataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataResponse(val *DataResponse) *NullableDataResponse {
	return &NullableDataResponse{value: val, isSet: true}
}

func (v NullableDataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


