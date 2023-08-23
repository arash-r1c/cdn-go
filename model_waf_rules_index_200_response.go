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

// checks if the WafRulesIndex200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WafRulesIndex200Response{}

// WafRulesIndex200Response struct for WafRulesIndex200Response
type WafRulesIndex200Response struct {
	Data []WafRule `json:"data,omitempty"`
	Links *PaginatedResponseLinks `json:"links,omitempty"`
	Meta *PaginatedResponseMeta `json:"meta,omitempty"`
}

// NewWafRulesIndex200Response instantiates a new WafRulesIndex200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafRulesIndex200Response() *WafRulesIndex200Response {
	this := WafRulesIndex200Response{}
	return &this
}

// NewWafRulesIndex200ResponseWithDefaults instantiates a new WafRulesIndex200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafRulesIndex200ResponseWithDefaults() *WafRulesIndex200Response {
	this := WafRulesIndex200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *WafRulesIndex200Response) GetData() []WafRule {
	if o == nil || IsNil(o.Data) {
		var ret []WafRule
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRulesIndex200Response) GetDataOk() ([]WafRule, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *WafRulesIndex200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []WafRule and assigns it to the Data field.
func (o *WafRulesIndex200Response) SetData(v []WafRule) {
	o.Data = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *WafRulesIndex200Response) GetLinks() PaginatedResponseLinks {
	if o == nil || IsNil(o.Links) {
		var ret PaginatedResponseLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRulesIndex200Response) GetLinksOk() (*PaginatedResponseLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *WafRulesIndex200Response) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginatedResponseLinks and assigns it to the Links field.
func (o *WafRulesIndex200Response) SetLinks(v PaginatedResponseLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *WafRulesIndex200Response) GetMeta() PaginatedResponseMeta {
	if o == nil || IsNil(o.Meta) {
		var ret PaginatedResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRulesIndex200Response) GetMetaOk() (*PaginatedResponseMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *WafRulesIndex200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PaginatedResponseMeta and assigns it to the Meta field.
func (o *WafRulesIndex200Response) SetMeta(v PaginatedResponseMeta) {
	o.Meta = &v
}

func (o WafRulesIndex200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WafRulesIndex200Response) ToMap() (map[string]interface{}, error) {
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

type NullableWafRulesIndex200Response struct {
	value *WafRulesIndex200Response
	isSet bool
}

func (v NullableWafRulesIndex200Response) Get() *WafRulesIndex200Response {
	return v.value
}

func (v *NullableWafRulesIndex200Response) Set(val *WafRulesIndex200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableWafRulesIndex200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableWafRulesIndex200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWafRulesIndex200Response(val *WafRulesIndex200Response) *NullableWafRulesIndex200Response {
	return &NullableWafRulesIndex200Response{value: val, isSet: true}
}

func (v NullableWafRulesIndex200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWafRulesIndex200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


