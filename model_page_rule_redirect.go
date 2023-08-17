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

// checks if the PageRuleRedirect type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PageRuleRedirect{}

// PageRuleRedirect struct for PageRuleRedirect
type PageRuleRedirect struct {
	Enable *bool `json:"enable,omitempty"`
	StatusCode *int32 `json:"status_code,omitempty"`
	Url NullableString `json:"url,omitempty"`
}

// NewPageRuleRedirect instantiates a new PageRuleRedirect object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPageRuleRedirect() *PageRuleRedirect {
	this := PageRuleRedirect{}
	var enable bool = false
	this.Enable = &enable
	var statusCode int32 = 301
	this.StatusCode = &statusCode
	return &this
}

// NewPageRuleRedirectWithDefaults instantiates a new PageRuleRedirect object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageRuleRedirectWithDefaults() *PageRuleRedirect {
	this := PageRuleRedirect{}
	var enable bool = false
	this.Enable = &enable
	var statusCode int32 = 301
	this.StatusCode = &statusCode
	return &this
}

// GetEnable returns the Enable field value if set, zero value otherwise.
func (o *PageRuleRedirect) GetEnable() bool {
	if o == nil || IsNil(o.Enable) {
		var ret bool
		return ret
	}
	return *o.Enable
}

// GetEnableOk returns a tuple with the Enable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageRuleRedirect) GetEnableOk() (*bool, bool) {
	if o == nil || IsNil(o.Enable) {
		return nil, false
	}
	return o.Enable, true
}

// HasEnable returns a boolean if a field has been set.
func (o *PageRuleRedirect) HasEnable() bool {
	if o != nil && !IsNil(o.Enable) {
		return true
	}

	return false
}

// SetEnable gets a reference to the given bool and assigns it to the Enable field.
func (o *PageRuleRedirect) SetEnable(v bool) {
	o.Enable = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *PageRuleRedirect) GetStatusCode() int32 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageRuleRedirect) GetStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *PageRuleRedirect) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *PageRuleRedirect) SetStatusCode(v int32) {
	o.StatusCode = &v
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PageRuleRedirect) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PageRuleRedirect) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *PageRuleRedirect) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *PageRuleRedirect) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *PageRuleRedirect) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *PageRuleRedirect) UnsetUrl() {
	o.Url.Unset()
}

func (o PageRuleRedirect) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PageRuleRedirect) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enable) {
		toSerialize["enable"] = o.Enable
	}
	if !IsNil(o.StatusCode) {
		toSerialize["status_code"] = o.StatusCode
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	return toSerialize, nil
}

type NullablePageRuleRedirect struct {
	value *PageRuleRedirect
	isSet bool
}

func (v NullablePageRuleRedirect) Get() *PageRuleRedirect {
	return v.value
}

func (v *NullablePageRuleRedirect) Set(val *PageRuleRedirect) {
	v.value = val
	v.isSet = true
}

func (v NullablePageRuleRedirect) IsSet() bool {
	return v.isSet
}

func (v *NullablePageRuleRedirect) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageRuleRedirect(val *PageRuleRedirect) *NullablePageRuleRedirect {
	return &NullablePageRuleRedirect{value: val, isSet: true}
}

func (v NullablePageRuleRedirect) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageRuleRedirect) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


