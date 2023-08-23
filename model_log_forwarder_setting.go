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

// LogForwarderSetting - struct for LogForwarderSetting
type LogForwarderSetting struct {
	LogForwarderDatadogConnectionType *LogForwarderDatadogConnectionType
	LogForwarderKafkaConnectionType *LogForwarderKafkaConnectionType
	LogForwarderLogglyConnectionType *LogForwarderLogglyConnectionType
	LogForwarderS3ConnectionType *LogForwarderS3ConnectionType
	LogForwarderSyslogConnectionType *LogForwarderSyslogConnectionType
}

// LogForwarderDatadogConnectionTypeAsLogForwarderSetting is a convenience function that returns LogForwarderDatadogConnectionType wrapped in LogForwarderSetting
func LogForwarderDatadogConnectionTypeAsLogForwarderSetting(v *LogForwarderDatadogConnectionType) LogForwarderSetting {
	return LogForwarderSetting{
		LogForwarderDatadogConnectionType: v,
	}
}

// LogForwarderKafkaConnectionTypeAsLogForwarderSetting is a convenience function that returns LogForwarderKafkaConnectionType wrapped in LogForwarderSetting
func LogForwarderKafkaConnectionTypeAsLogForwarderSetting(v *LogForwarderKafkaConnectionType) LogForwarderSetting {
	return LogForwarderSetting{
		LogForwarderKafkaConnectionType: v,
	}
}

// LogForwarderLogglyConnectionTypeAsLogForwarderSetting is a convenience function that returns LogForwarderLogglyConnectionType wrapped in LogForwarderSetting
func LogForwarderLogglyConnectionTypeAsLogForwarderSetting(v *LogForwarderLogglyConnectionType) LogForwarderSetting {
	return LogForwarderSetting{
		LogForwarderLogglyConnectionType: v,
	}
}

// LogForwarderS3ConnectionTypeAsLogForwarderSetting is a convenience function that returns LogForwarderS3ConnectionType wrapped in LogForwarderSetting
func LogForwarderS3ConnectionTypeAsLogForwarderSetting(v *LogForwarderS3ConnectionType) LogForwarderSetting {
	return LogForwarderSetting{
		LogForwarderS3ConnectionType: v,
	}
}

// LogForwarderSyslogConnectionTypeAsLogForwarderSetting is a convenience function that returns LogForwarderSyslogConnectionType wrapped in LogForwarderSetting
func LogForwarderSyslogConnectionTypeAsLogForwarderSetting(v *LogForwarderSyslogConnectionType) LogForwarderSetting {
	return LogForwarderSetting{
		LogForwarderSyslogConnectionType: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *LogForwarderSetting) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into LogForwarderDatadogConnectionType
	err = newStrictDecoder(data).Decode(&dst.LogForwarderDatadogConnectionType)
	if err == nil {
		jsonLogForwarderDatadogConnectionType, _ := json.Marshal(dst.LogForwarderDatadogConnectionType)
		if string(jsonLogForwarderDatadogConnectionType) == "{}" { // empty struct
			dst.LogForwarderDatadogConnectionType = nil
		} else {
			match++
		}
	} else {
		dst.LogForwarderDatadogConnectionType = nil
	}

	// try to unmarshal data into LogForwarderKafkaConnectionType
	err = newStrictDecoder(data).Decode(&dst.LogForwarderKafkaConnectionType)
	if err == nil {
		jsonLogForwarderKafkaConnectionType, _ := json.Marshal(dst.LogForwarderKafkaConnectionType)
		if string(jsonLogForwarderKafkaConnectionType) == "{}" { // empty struct
			dst.LogForwarderKafkaConnectionType = nil
		} else {
			match++
		}
	} else {
		dst.LogForwarderKafkaConnectionType = nil
	}

	// try to unmarshal data into LogForwarderLogglyConnectionType
	err = newStrictDecoder(data).Decode(&dst.LogForwarderLogglyConnectionType)
	if err == nil {
		jsonLogForwarderLogglyConnectionType, _ := json.Marshal(dst.LogForwarderLogglyConnectionType)
		if string(jsonLogForwarderLogglyConnectionType) == "{}" { // empty struct
			dst.LogForwarderLogglyConnectionType = nil
		} else {
			match++
		}
	} else {
		dst.LogForwarderLogglyConnectionType = nil
	}

	// try to unmarshal data into LogForwarderS3ConnectionType
	err = newStrictDecoder(data).Decode(&dst.LogForwarderS3ConnectionType)
	if err == nil {
		jsonLogForwarderS3ConnectionType, _ := json.Marshal(dst.LogForwarderS3ConnectionType)
		if string(jsonLogForwarderS3ConnectionType) == "{}" { // empty struct
			dst.LogForwarderS3ConnectionType = nil
		} else {
			match++
		}
	} else {
		dst.LogForwarderS3ConnectionType = nil
	}

	// try to unmarshal data into LogForwarderSyslogConnectionType
	err = newStrictDecoder(data).Decode(&dst.LogForwarderSyslogConnectionType)
	if err == nil {
		jsonLogForwarderSyslogConnectionType, _ := json.Marshal(dst.LogForwarderSyslogConnectionType)
		if string(jsonLogForwarderSyslogConnectionType) == "{}" { // empty struct
			dst.LogForwarderSyslogConnectionType = nil
		} else {
			match++
		}
	} else {
		dst.LogForwarderSyslogConnectionType = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.LogForwarderDatadogConnectionType = nil
		dst.LogForwarderKafkaConnectionType = nil
		dst.LogForwarderLogglyConnectionType = nil
		dst.LogForwarderS3ConnectionType = nil
		dst.LogForwarderSyslogConnectionType = nil

		return fmt.Errorf("data matches more than one schema in oneOf(LogForwarderSetting)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(LogForwarderSetting)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src LogForwarderSetting) MarshalJSON() ([]byte, error) {
	if src.LogForwarderDatadogConnectionType != nil {
		return json.Marshal(&src.LogForwarderDatadogConnectionType)
	}

	if src.LogForwarderKafkaConnectionType != nil {
		return json.Marshal(&src.LogForwarderKafkaConnectionType)
	}

	if src.LogForwarderLogglyConnectionType != nil {
		return json.Marshal(&src.LogForwarderLogglyConnectionType)
	}

	if src.LogForwarderS3ConnectionType != nil {
		return json.Marshal(&src.LogForwarderS3ConnectionType)
	}

	if src.LogForwarderSyslogConnectionType != nil {
		return json.Marshal(&src.LogForwarderSyslogConnectionType)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *LogForwarderSetting) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.LogForwarderDatadogConnectionType != nil {
		return obj.LogForwarderDatadogConnectionType
	}

	if obj.LogForwarderKafkaConnectionType != nil {
		return obj.LogForwarderKafkaConnectionType
	}

	if obj.LogForwarderLogglyConnectionType != nil {
		return obj.LogForwarderLogglyConnectionType
	}

	if obj.LogForwarderS3ConnectionType != nil {
		return obj.LogForwarderS3ConnectionType
	}

	if obj.LogForwarderSyslogConnectionType != nil {
		return obj.LogForwarderSyslogConnectionType
	}

	// all schemas are nil
	return nil
}

type NullableLogForwarderSetting struct {
	value *LogForwarderSetting
	isSet bool
}

func (v NullableLogForwarderSetting) Get() *LogForwarderSetting {
	return v.value
}

func (v *NullableLogForwarderSetting) Set(val *LogForwarderSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableLogForwarderSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableLogForwarderSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogForwarderSetting(val *LogForwarderSetting) *NullableLogForwarderSetting {
	return &NullableLogForwarderSetting{value: val, isSet: true}
}

func (v NullableLogForwarderSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogForwarderSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


