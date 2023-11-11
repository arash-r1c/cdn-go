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

// checks if the ResponseTimeChartsIrSeriesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseTimeChartsIrSeriesInner{}

// ResponseTimeChartsIrSeriesInner struct for ResponseTimeChartsIrSeriesInner
type ResponseTimeChartsIrSeriesInner struct {
	Name *string `json:"name,omitempty"`
	Data []float32 `json:"data,omitempty"`
}

// NewResponseTimeChartsIrSeriesInner instantiates a new ResponseTimeChartsIrSeriesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseTimeChartsIrSeriesInner() *ResponseTimeChartsIrSeriesInner {
	this := ResponseTimeChartsIrSeriesInner{}
	return &this
}

// NewResponseTimeChartsIrSeriesInnerWithDefaults instantiates a new ResponseTimeChartsIrSeriesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseTimeChartsIrSeriesInnerWithDefaults() *ResponseTimeChartsIrSeriesInner {
	this := ResponseTimeChartsIrSeriesInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ResponseTimeChartsIrSeriesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTimeChartsIrSeriesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ResponseTimeChartsIrSeriesInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ResponseTimeChartsIrSeriesInner) SetName(v string) {
	o.Name = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ResponseTimeChartsIrSeriesInner) GetData() []float32 {
	if o == nil || IsNil(o.Data) {
		var ret []float32
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTimeChartsIrSeriesInner) GetDataOk() ([]float32, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ResponseTimeChartsIrSeriesInner) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []float32 and assigns it to the Data field.
func (o *ResponseTimeChartsIrSeriesInner) SetData(v []float32) {
	o.Data = v
}

func (o ResponseTimeChartsIrSeriesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseTimeChartsIrSeriesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableResponseTimeChartsIrSeriesInner struct {
	value *ResponseTimeChartsIrSeriesInner
	isSet bool
}

func (v NullableResponseTimeChartsIrSeriesInner) Get() *ResponseTimeChartsIrSeriesInner {
	return v.value
}

func (v *NullableResponseTimeChartsIrSeriesInner) Set(val *ResponseTimeChartsIrSeriesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseTimeChartsIrSeriesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseTimeChartsIrSeriesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseTimeChartsIrSeriesInner(val *ResponseTimeChartsIrSeriesInner) *NullableResponseTimeChartsIrSeriesInner {
	return &NullableResponseTimeChartsIrSeriesInner{value: val, isSet: true}
}

func (v NullableResponseTimeChartsIrSeriesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseTimeChartsIrSeriesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


