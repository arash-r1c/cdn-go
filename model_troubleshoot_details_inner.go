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

// checks if the TroubleshootDetailsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TroubleshootDetailsInner{}

// TroubleshootDetailsInner struct for TroubleshootDetailsInner
type TroubleshootDetailsInner struct {
	Id *string `json:"id,omitempty"`
	Status *string `json:"status,omitempty"`
	Details *string `json:"details,omitempty"`
}

// NewTroubleshootDetailsInner instantiates a new TroubleshootDetailsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTroubleshootDetailsInner() *TroubleshootDetailsInner {
	this := TroubleshootDetailsInner{}
	return &this
}

// NewTroubleshootDetailsInnerWithDefaults instantiates a new TroubleshootDetailsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTroubleshootDetailsInnerWithDefaults() *TroubleshootDetailsInner {
	this := TroubleshootDetailsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TroubleshootDetailsInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TroubleshootDetailsInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TroubleshootDetailsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TroubleshootDetailsInner) SetId(v string) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TroubleshootDetailsInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TroubleshootDetailsInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TroubleshootDetailsInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TroubleshootDetailsInner) SetStatus(v string) {
	o.Status = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *TroubleshootDetailsInner) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TroubleshootDetailsInner) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *TroubleshootDetailsInner) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *TroubleshootDetailsInner) SetDetails(v string) {
	o.Details = &v
}

func (o TroubleshootDetailsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TroubleshootDetailsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableTroubleshootDetailsInner struct {
	value *TroubleshootDetailsInner
	isSet bool
}

func (v NullableTroubleshootDetailsInner) Get() *TroubleshootDetailsInner {
	return v.value
}

func (v *NullableTroubleshootDetailsInner) Set(val *TroubleshootDetailsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTroubleshootDetailsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTroubleshootDetailsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTroubleshootDetailsInner(val *TroubleshootDetailsInner) *NullableTroubleshootDetailsInner {
	return &NullableTroubleshootDetailsInner{value: val, isSet: true}
}

func (v NullableTroubleshootDetailsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTroubleshootDetailsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


