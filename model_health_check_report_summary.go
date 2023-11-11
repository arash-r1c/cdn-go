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

// checks if the HealthCheckReportSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HealthCheckReportSummary{}

// HealthCheckReportSummary struct for HealthCheckReportSummary
type HealthCheckReportSummary struct {
	Zone *string `json:"zone,omitempty"`
	Status *bool `json:"status,omitempty"`
	Total *int32 `json:"total,omitempty"`
	Failed *int32 `json:"failed,omitempty"`
	Details []HealthCheckReportSummaryDetail `json:"details,omitempty"`
}

// NewHealthCheckReportSummary instantiates a new HealthCheckReportSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthCheckReportSummary() *HealthCheckReportSummary {
	this := HealthCheckReportSummary{}
	return &this
}

// NewHealthCheckReportSummaryWithDefaults instantiates a new HealthCheckReportSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthCheckReportSummaryWithDefaults() *HealthCheckReportSummary {
	this := HealthCheckReportSummary{}
	return &this
}

// GetZone returns the Zone field value if set, zero value otherwise.
func (o *HealthCheckReportSummary) GetZone() string {
	if o == nil || IsNil(o.Zone) {
		var ret string
		return ret
	}
	return *o.Zone
}

// GetZoneOk returns a tuple with the Zone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheckReportSummary) GetZoneOk() (*string, bool) {
	if o == nil || IsNil(o.Zone) {
		return nil, false
	}
	return o.Zone, true
}

// HasZone returns a boolean if a field has been set.
func (o *HealthCheckReportSummary) HasZone() bool {
	if o != nil && !IsNil(o.Zone) {
		return true
	}

	return false
}

// SetZone gets a reference to the given string and assigns it to the Zone field.
func (o *HealthCheckReportSummary) SetZone(v string) {
	o.Zone = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *HealthCheckReportSummary) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheckReportSummary) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *HealthCheckReportSummary) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *HealthCheckReportSummary) SetStatus(v bool) {
	o.Status = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *HealthCheckReportSummary) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheckReportSummary) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *HealthCheckReportSummary) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *HealthCheckReportSummary) SetTotal(v int32) {
	o.Total = &v
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *HealthCheckReportSummary) GetFailed() int32 {
	if o == nil || IsNil(o.Failed) {
		var ret int32
		return ret
	}
	return *o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheckReportSummary) GetFailedOk() (*int32, bool) {
	if o == nil || IsNil(o.Failed) {
		return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *HealthCheckReportSummary) HasFailed() bool {
	if o != nil && !IsNil(o.Failed) {
		return true
	}

	return false
}

// SetFailed gets a reference to the given int32 and assigns it to the Failed field.
func (o *HealthCheckReportSummary) SetFailed(v int32) {
	o.Failed = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *HealthCheckReportSummary) GetDetails() []HealthCheckReportSummaryDetail {
	if o == nil || IsNil(o.Details) {
		var ret []HealthCheckReportSummaryDetail
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheckReportSummary) GetDetailsOk() ([]HealthCheckReportSummaryDetail, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *HealthCheckReportSummary) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []HealthCheckReportSummaryDetail and assigns it to the Details field.
func (o *HealthCheckReportSummary) SetDetails(v []HealthCheckReportSummaryDetail) {
	o.Details = v
}

func (o HealthCheckReportSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HealthCheckReportSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Zone) {
		toSerialize["zone"] = o.Zone
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.Failed) {
		toSerialize["failed"] = o.Failed
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableHealthCheckReportSummary struct {
	value *HealthCheckReportSummary
	isSet bool
}

func (v NullableHealthCheckReportSummary) Get() *HealthCheckReportSummary {
	return v.value
}

func (v *NullableHealthCheckReportSummary) Set(val *HealthCheckReportSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthCheckReportSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthCheckReportSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthCheckReportSummary(val *HealthCheckReportSummary) *NullableHealthCheckReportSummary {
	return &NullableHealthCheckReportSummary{value: val, isSet: true}
}

func (v NullableHealthCheckReportSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthCheckReportSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


