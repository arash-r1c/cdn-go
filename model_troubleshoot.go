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

// checks if the Troubleshoot type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Troubleshoot{}

// Troubleshoot struct for Troubleshoot
type Troubleshoot struct {
	Id *string `json:"id,omitempty"`
	Details []TroubleshootDetailsInner `json:"details,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
}

// NewTroubleshoot instantiates a new Troubleshoot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTroubleshoot() *Troubleshoot {
	this := Troubleshoot{}
	return &this
}

// NewTroubleshootWithDefaults instantiates a new Troubleshoot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTroubleshootWithDefaults() *Troubleshoot {
	this := Troubleshoot{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Troubleshoot) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Troubleshoot) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Troubleshoot) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Troubleshoot) SetId(v string) {
	o.Id = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *Troubleshoot) GetDetails() []TroubleshootDetailsInner {
	if o == nil || IsNil(o.Details) {
		var ret []TroubleshootDetailsInner
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Troubleshoot) GetDetailsOk() ([]TroubleshootDetailsInner, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *Troubleshoot) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []TroubleshootDetailsInner and assigns it to the Details field.
func (o *Troubleshoot) SetDetails(v []TroubleshootDetailsInner) {
	o.Details = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Troubleshoot) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Troubleshoot) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Troubleshoot) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *Troubleshoot) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

func (o Troubleshoot) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Troubleshoot) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	return toSerialize, nil
}

type NullableTroubleshoot struct {
	value *Troubleshoot
	isSet bool
}

func (v NullableTroubleshoot) Get() *Troubleshoot {
	return v.value
}

func (v *NullableTroubleshoot) Set(val *Troubleshoot) {
	v.value = val
	v.isSet = true
}

func (v NullableTroubleshoot) IsSet() bool {
	return v.isSet
}

func (v *NullableTroubleshoot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTroubleshoot(val *Troubleshoot) *NullableTroubleshoot {
	return &NullableTroubleshoot{value: val, isSet: true}
}

func (v NullableTroubleshoot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTroubleshoot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


