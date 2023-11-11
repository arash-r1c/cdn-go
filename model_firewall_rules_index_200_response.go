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

// checks if the FirewallRulesIndex200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FirewallRulesIndex200Response{}

// FirewallRulesIndex200Response struct for FirewallRulesIndex200Response
type FirewallRulesIndex200Response struct {
	Data []FirewallRuleView `json:"data,omitempty"`
	Links *PaginatedResponseLinks `json:"links,omitempty"`
	Meta *PaginatedResponseMeta `json:"meta,omitempty"`
}

// NewFirewallRulesIndex200Response instantiates a new FirewallRulesIndex200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirewallRulesIndex200Response() *FirewallRulesIndex200Response {
	this := FirewallRulesIndex200Response{}
	return &this
}

// NewFirewallRulesIndex200ResponseWithDefaults instantiates a new FirewallRulesIndex200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirewallRulesIndex200ResponseWithDefaults() *FirewallRulesIndex200Response {
	this := FirewallRulesIndex200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *FirewallRulesIndex200Response) GetData() []FirewallRuleView {
	if o == nil || IsNil(o.Data) {
		var ret []FirewallRuleView
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallRulesIndex200Response) GetDataOk() ([]FirewallRuleView, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *FirewallRulesIndex200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []FirewallRuleView and assigns it to the Data field.
func (o *FirewallRulesIndex200Response) SetData(v []FirewallRuleView) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *FirewallRulesIndex200Response) GetLinks() PaginatedResponseLinks {
	if o == nil || IsNil(o.Links) {
		var ret PaginatedResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallRulesIndex200Response) GetLinksOk() (*PaginatedResponseLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *FirewallRulesIndex200Response) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginatedResponseLinks and assigns it to the Links field.
func (o *FirewallRulesIndex200Response) SetLinks(v PaginatedResponseLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *FirewallRulesIndex200Response) GetMeta() PaginatedResponseMeta {
	if o == nil || IsNil(o.Meta) {
		var ret PaginatedResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallRulesIndex200Response) GetMetaOk() (*PaginatedResponseMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *FirewallRulesIndex200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PaginatedResponseMeta and assigns it to the Meta field.
func (o *FirewallRulesIndex200Response) SetMeta(v PaginatedResponseMeta) {
	o.Meta = &v
}

func (o FirewallRulesIndex200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FirewallRulesIndex200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableFirewallRulesIndex200Response struct {
	value *FirewallRulesIndex200Response
	isSet bool
}

func (v NullableFirewallRulesIndex200Response) Get() *FirewallRulesIndex200Response {
	return v.value
}

func (v *NullableFirewallRulesIndex200Response) Set(val *FirewallRulesIndex200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableFirewallRulesIndex200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableFirewallRulesIndex200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirewallRulesIndex200Response(val *FirewallRulesIndex200Response) *NullableFirewallRulesIndex200Response {
	return &NullableFirewallRulesIndex200Response{value: val, isSet: true}
}

func (v NullableFirewallRulesIndex200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirewallRulesIndex200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


