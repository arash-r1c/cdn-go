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

// checks if the TransportLayerProxyFirewall type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransportLayerProxyFirewall{}

// TransportLayerProxyFirewall struct for TransportLayerProxyFirewall
type TransportLayerProxyFirewall struct {
	Access *string `json:"access,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Type *string `json:"type,omitempty"`
	Match *TransportLayerProxyMatch `json:"match,omitempty"`
	Active *bool `json:"active,omitempty"`
}

// NewTransportLayerProxyFirewall instantiates a new TransportLayerProxyFirewall object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransportLayerProxyFirewall() *TransportLayerProxyFirewall {
	this := TransportLayerProxyFirewall{}
	return &this
}

// NewTransportLayerProxyFirewallWithDefaults instantiates a new TransportLayerProxyFirewall object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransportLayerProxyFirewallWithDefaults() *TransportLayerProxyFirewall {
	this := TransportLayerProxyFirewall{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *TransportLayerProxyFirewall) GetAccess() string {
	if o == nil || IsNil(o.Access) {
		var ret string
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportLayerProxyFirewall) GetAccessOk() (*string, bool) {
	if o == nil || IsNil(o.Access) {
		return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *TransportLayerProxyFirewall) HasAccess() bool {
	if o != nil && !IsNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given string and assigns it to the Access field.
func (o *TransportLayerProxyFirewall) SetAccess(v string) {
	o.Access = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TransportLayerProxyFirewall) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportLayerProxyFirewall) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TransportLayerProxyFirewall) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TransportLayerProxyFirewall) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TransportLayerProxyFirewall) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportLayerProxyFirewall) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TransportLayerProxyFirewall) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TransportLayerProxyFirewall) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TransportLayerProxyFirewall) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportLayerProxyFirewall) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TransportLayerProxyFirewall) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TransportLayerProxyFirewall) SetType(v string) {
	o.Type = &v
}

// GetMatch returns the Match field value if set, zero value otherwise.
func (o *TransportLayerProxyFirewall) GetMatch() TransportLayerProxyMatch {
	if o == nil || IsNil(o.Match) {
		var ret TransportLayerProxyMatch
		return ret
	}
	return *o.Match
}

// GetMatchOk returns a tuple with the Match field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportLayerProxyFirewall) GetMatchOk() (*TransportLayerProxyMatch, bool) {
	if o == nil || IsNil(o.Match) {
		return nil, false
	}
	return o.Match, true
}

// HasMatch returns a boolean if a field has been set.
func (o *TransportLayerProxyFirewall) HasMatch() bool {
	if o != nil && !IsNil(o.Match) {
		return true
	}

	return false
}

// SetMatch gets a reference to the given TransportLayerProxyMatch and assigns it to the Match field.
func (o *TransportLayerProxyFirewall) SetMatch(v TransportLayerProxyMatch) {
	o.Match = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *TransportLayerProxyFirewall) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportLayerProxyFirewall) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *TransportLayerProxyFirewall) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *TransportLayerProxyFirewall) SetActive(v bool) {
	o.Active = &v
}

func (o TransportLayerProxyFirewall) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransportLayerProxyFirewall) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Match) {
		toSerialize["match"] = o.Match
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	return toSerialize, nil
}

type NullableTransportLayerProxyFirewall struct {
	value *TransportLayerProxyFirewall
	isSet bool
}

func (v NullableTransportLayerProxyFirewall) Get() *TransportLayerProxyFirewall {
	return v.value
}

func (v *NullableTransportLayerProxyFirewall) Set(val *TransportLayerProxyFirewall) {
	v.value = val
	v.isSet = true
}

func (v NullableTransportLayerProxyFirewall) IsSet() bool {
	return v.isSet
}

func (v *NullableTransportLayerProxyFirewall) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransportLayerProxyFirewall(val *TransportLayerProxyFirewall) *NullableTransportLayerProxyFirewall {
	return &NullableTransportLayerProxyFirewall{value: val, isSet: true}
}

func (v NullableTransportLayerProxyFirewall) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransportLayerProxyFirewall) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


