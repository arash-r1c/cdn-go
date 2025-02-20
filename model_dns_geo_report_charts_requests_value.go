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

// checks if the DnsGeoReportChartsRequestsValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnsGeoReportChartsRequestsValue{}

// DnsGeoReportChartsRequestsValue struct for DnsGeoReportChartsRequestsValue
type DnsGeoReportChartsRequestsValue struct {
	FillKey *int64 `json:"fillKey,omitempty"`
	Name *string `json:"name,omitempty"`
	Value *int64 `json:"value,omitempty"`
}

// NewDnsGeoReportChartsRequestsValue instantiates a new DnsGeoReportChartsRequestsValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsGeoReportChartsRequestsValue() *DnsGeoReportChartsRequestsValue {
	this := DnsGeoReportChartsRequestsValue{}
	return &this
}

// NewDnsGeoReportChartsRequestsValueWithDefaults instantiates a new DnsGeoReportChartsRequestsValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsGeoReportChartsRequestsValueWithDefaults() *DnsGeoReportChartsRequestsValue {
	this := DnsGeoReportChartsRequestsValue{}
	return &this
}

// GetFillKey returns the FillKey field value if set, zero value otherwise.
func (o *DnsGeoReportChartsRequestsValue) GetFillKey() int64 {
	if o == nil || IsNil(o.FillKey) {
		var ret int64
		return ret
	}
	return *o.FillKey
}

// GetFillKeyOk returns a tuple with the FillKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsGeoReportChartsRequestsValue) GetFillKeyOk() (*int64, bool) {
	if o == nil || IsNil(o.FillKey) {
		return nil, false
	}
	return o.FillKey, true
}

// HasFillKey returns a boolean if a field has been set.
func (o *DnsGeoReportChartsRequestsValue) HasFillKey() bool {
	if o != nil && !IsNil(o.FillKey) {
		return true
	}

	return false
}

// SetFillKey gets a reference to the given int64 and assigns it to the FillKey field.
func (o *DnsGeoReportChartsRequestsValue) SetFillKey(v int64) {
	o.FillKey = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DnsGeoReportChartsRequestsValue) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsGeoReportChartsRequestsValue) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DnsGeoReportChartsRequestsValue) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DnsGeoReportChartsRequestsValue) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DnsGeoReportChartsRequestsValue) GetValue() int64 {
	if o == nil || IsNil(o.Value) {
		var ret int64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsGeoReportChartsRequestsValue) GetValueOk() (*int64, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DnsGeoReportChartsRequestsValue) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int64 and assigns it to the Value field.
func (o *DnsGeoReportChartsRequestsValue) SetValue(v int64) {
	o.Value = &v
}

func (o DnsGeoReportChartsRequestsValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnsGeoReportChartsRequestsValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FillKey) {
		toSerialize["fillKey"] = o.FillKey
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableDnsGeoReportChartsRequestsValue struct {
	value *DnsGeoReportChartsRequestsValue
	isSet bool
}

func (v NullableDnsGeoReportChartsRequestsValue) Get() *DnsGeoReportChartsRequestsValue {
	return v.value
}

func (v *NullableDnsGeoReportChartsRequestsValue) Set(val *DnsGeoReportChartsRequestsValue) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsGeoReportChartsRequestsValue) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsGeoReportChartsRequestsValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsGeoReportChartsRequestsValue(val *DnsGeoReportChartsRequestsValue) *NullableDnsGeoReportChartsRequestsValue {
	return &NullableDnsGeoReportChartsRequestsValue{value: val, isSet: true}
}

func (v NullableDnsGeoReportChartsRequestsValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsGeoReportChartsRequestsValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


