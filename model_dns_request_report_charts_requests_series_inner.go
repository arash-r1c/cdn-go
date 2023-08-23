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

// checks if the DnsRequestReportChartsRequestsSeriesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnsRequestReportChartsRequestsSeriesInner{}

// DnsRequestReportChartsRequestsSeriesInner struct for DnsRequestReportChartsRequestsSeriesInner
type DnsRequestReportChartsRequestsSeriesInner struct {
	Name *string `json:"name,omitempty"`
	Data []int32 `json:"data,omitempty"`
}

// NewDnsRequestReportChartsRequestsSeriesInner instantiates a new DnsRequestReportChartsRequestsSeriesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsRequestReportChartsRequestsSeriesInner() *DnsRequestReportChartsRequestsSeriesInner {
	this := DnsRequestReportChartsRequestsSeriesInner{}
	return &this
}

// NewDnsRequestReportChartsRequestsSeriesInnerWithDefaults instantiates a new DnsRequestReportChartsRequestsSeriesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsRequestReportChartsRequestsSeriesInnerWithDefaults() *DnsRequestReportChartsRequestsSeriesInner {
	this := DnsRequestReportChartsRequestsSeriesInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DnsRequestReportChartsRequestsSeriesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRequestReportChartsRequestsSeriesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DnsRequestReportChartsRequestsSeriesInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DnsRequestReportChartsRequestsSeriesInner) SetName(v string) {
	o.Name = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DnsRequestReportChartsRequestsSeriesInner) GetData() []int32 {
	if o == nil || IsNil(o.Data) {
		var ret []int32
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRequestReportChartsRequestsSeriesInner) GetDataOk() ([]int32, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DnsRequestReportChartsRequestsSeriesInner) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []int32 and assigns it to the Data field.
func (o *DnsRequestReportChartsRequestsSeriesInner) SetData(v []int32) {
	o.Data = v
}

func (o DnsRequestReportChartsRequestsSeriesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnsRequestReportChartsRequestsSeriesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableDnsRequestReportChartsRequestsSeriesInner struct {
	value *DnsRequestReportChartsRequestsSeriesInner
	isSet bool
}

func (v NullableDnsRequestReportChartsRequestsSeriesInner) Get() *DnsRequestReportChartsRequestsSeriesInner {
	return v.value
}

func (v *NullableDnsRequestReportChartsRequestsSeriesInner) Set(val *DnsRequestReportChartsRequestsSeriesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsRequestReportChartsRequestsSeriesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsRequestReportChartsRequestsSeriesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsRequestReportChartsRequestsSeriesInner(val *DnsRequestReportChartsRequestsSeriesInner) *NullableDnsRequestReportChartsRequestsSeriesInner {
	return &NullableDnsRequestReportChartsRequestsSeriesInner{value: val, isSet: true}
}

func (v NullableDnsRequestReportChartsRequestsSeriesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsRequestReportChartsRequestsSeriesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


