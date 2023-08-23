/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.105.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package arvancloud

import (
	"encoding/json"
	"fmt"
)

// MonitoringStatus the model 'MonitoringStatus'
type MonitoringStatus string

// List of MonitoringStatus
const (
	OFF MonitoringStatus = "off"
	NO_DATA MonitoringStatus = "no-data"
	HEALTHY MonitoringStatus = "healthy"
	UNHEALTHY MonitoringStatus = "unhealthy"
)

// All allowed values of MonitoringStatus enum
var AllowedMonitoringStatusEnumValues = []MonitoringStatus{
	"off",
	"no-data",
	"healthy",
	"unhealthy",
}

func (v *MonitoringStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MonitoringStatus(value)
	for _, existing := range AllowedMonitoringStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MonitoringStatus", value)
}

// NewMonitoringStatusFromValue returns a pointer to a valid MonitoringStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMonitoringStatusFromValue(v string) (*MonitoringStatus, error) {
	ev := MonitoringStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MonitoringStatus: valid values are %v", v, AllowedMonitoringStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MonitoringStatus) IsValid() bool {
	for _, existing := range AllowedMonitoringStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MonitoringStatus value
func (v MonitoringStatus) Ptr() *MonitoringStatus {
	return &v
}

type NullableMonitoringStatus struct {
	value *MonitoringStatus
	isSet bool
}

func (v NullableMonitoringStatus) Get() *MonitoringStatus {
	return v.value
}

func (v *NullableMonitoringStatus) Set(val *MonitoringStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringStatus(val *MonitoringStatus) *NullableMonitoringStatus {
	return &NullableMonitoringStatus{value: val, isSet: true}
}

func (v NullableMonitoringStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

