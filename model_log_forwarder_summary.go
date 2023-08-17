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

// checks if the LogForwarderSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogForwarderSummary{}

// LogForwarderSummary struct for LogForwarderSummary
type LogForwarderSummary struct {
	Id *string `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
	ConnectionType *string `json:"connection_type,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Status *bool `json:"status,omitempty"`
}

// NewLogForwarderSummary instantiates a new LogForwarderSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogForwarderSummary() *LogForwarderSummary {
	this := LogForwarderSummary{}
	return &this
}

// NewLogForwarderSummaryWithDefaults instantiates a new LogForwarderSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogForwarderSummaryWithDefaults() *LogForwarderSummary {
	this := LogForwarderSummary{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LogForwarderSummary) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderSummary) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LogForwarderSummary) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LogForwarderSummary) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LogForwarderSummary) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderSummary) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LogForwarderSummary) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *LogForwarderSummary) SetType(v string) {
	o.Type = &v
}

// GetConnectionType returns the ConnectionType field value if set, zero value otherwise.
func (o *LogForwarderSummary) GetConnectionType() string {
	if o == nil || IsNil(o.ConnectionType) {
		var ret string
		return ret
	}
	return *o.ConnectionType
}

// GetConnectionTypeOk returns a tuple with the ConnectionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderSummary) GetConnectionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionType) {
		return nil, false
	}
	return o.ConnectionType, true
}

// HasConnectionType returns a boolean if a field has been set.
func (o *LogForwarderSummary) HasConnectionType() bool {
	if o != nil && !IsNil(o.ConnectionType) {
		return true
	}

	return false
}

// SetConnectionType gets a reference to the given string and assigns it to the ConnectionType field.
func (o *LogForwarderSummary) SetConnectionType(v string) {
	o.ConnectionType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LogForwarderSummary) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderSummary) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LogForwarderSummary) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LogForwarderSummary) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LogForwarderSummary) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderSummary) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LogForwarderSummary) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LogForwarderSummary) SetDescription(v string) {
	o.Description = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LogForwarderSummary) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogForwarderSummary) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LogForwarderSummary) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *LogForwarderSummary) SetStatus(v bool) {
	o.Status = &v
}

func (o LogForwarderSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogForwarderSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.ConnectionType) {
		toSerialize["connection_type"] = o.ConnectionType
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableLogForwarderSummary struct {
	value *LogForwarderSummary
	isSet bool
}

func (v NullableLogForwarderSummary) Get() *LogForwarderSummary {
	return v.value
}

func (v *NullableLogForwarderSummary) Set(val *LogForwarderSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableLogForwarderSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableLogForwarderSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogForwarderSummary(val *LogForwarderSummary) *NullableLogForwarderSummary {
	return &NullableLogForwarderSummary{value: val, isSet: true}
}

func (v NullableLogForwarderSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogForwarderSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


