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

// checks if the DnsRequestReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnsRequestReport{}

// DnsRequestReport struct for DnsRequestReport
type DnsRequestReport struct {
	Statistics *DnsRequestReportStatistics `json:"statistics,omitempty"`
	Charts *DnsRequestReportCharts `json:"charts,omitempty"`
}

// NewDnsRequestReport instantiates a new DnsRequestReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsRequestReport() *DnsRequestReport {
	this := DnsRequestReport{}
	return &this
}

// NewDnsRequestReportWithDefaults instantiates a new DnsRequestReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsRequestReportWithDefaults() *DnsRequestReport {
	this := DnsRequestReport{}
	return &this
}

// GetStatistics returns the Statistics field value if set, zero value otherwise.
func (o *DnsRequestReport) GetStatistics() DnsRequestReportStatistics {
	if o == nil || IsNil(o.Statistics) {
		var ret DnsRequestReportStatistics
		return ret
	}
	return *o.Statistics
}

// GetStatisticsOk returns a tuple with the Statistics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRequestReport) GetStatisticsOk() (*DnsRequestReportStatistics, bool) {
	if o == nil || IsNil(o.Statistics) {
		return nil, false
	}
	return o.Statistics, true
}

// HasStatistics returns a boolean if a field has been set.
func (o *DnsRequestReport) HasStatistics() bool {
	if o != nil && !IsNil(o.Statistics) {
		return true
	}

	return false
}

// SetStatistics gets a reference to the given DnsRequestReportStatistics and assigns it to the Statistics field.
func (o *DnsRequestReport) SetStatistics(v DnsRequestReportStatistics) {
	o.Statistics = &v
}

// GetCharts returns the Charts field value if set, zero value otherwise.
func (o *DnsRequestReport) GetCharts() DnsRequestReportCharts {
	if o == nil || IsNil(o.Charts) {
		var ret DnsRequestReportCharts
		return ret
	}
	return *o.Charts
}

// GetChartsOk returns a tuple with the Charts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRequestReport) GetChartsOk() (*DnsRequestReportCharts, bool) {
	if o == nil || IsNil(o.Charts) {
		return nil, false
	}
	return o.Charts, true
}

// HasCharts returns a boolean if a field has been set.
func (o *DnsRequestReport) HasCharts() bool {
	if o != nil && !IsNil(o.Charts) {
		return true
	}

	return false
}

// SetCharts gets a reference to the given DnsRequestReportCharts and assigns it to the Charts field.
func (o *DnsRequestReport) SetCharts(v DnsRequestReportCharts) {
	o.Charts = &v
}

func (o DnsRequestReport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnsRequestReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Statistics) {
		toSerialize["statistics"] = o.Statistics
	}
	if !IsNil(o.Charts) {
		toSerialize["charts"] = o.Charts
	}
	return toSerialize, nil
}

type NullableDnsRequestReport struct {
	value *DnsRequestReport
	isSet bool
}

func (v NullableDnsRequestReport) Get() *DnsRequestReport {
	return v.value
}

func (v *NullableDnsRequestReport) Set(val *DnsRequestReport) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsRequestReport) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsRequestReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsRequestReport(val *DnsRequestReport) *NullableDnsRequestReport {
	return &NullableDnsRequestReport{value: val, isSet: true}
}

func (v NullableDnsRequestReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsRequestReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


