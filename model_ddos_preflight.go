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

// checks if the DdosPreflight type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DdosPreflight{}

// DdosPreflight struct for DdosPreflight
type DdosPreflight struct {
	AccessOrigin *string `json:"access_origin,omitempty"`
	AccessCredentials *string `json:"access_credentials,omitempty"`
	AccessMethods []string `json:"access_methods,omitempty"`
	AccessHeaders *string `json:"access_headers,omitempty"`
	AccessExposeHeaders *string `json:"access_expose_headers,omitempty"`
}

// NewDdosPreflight instantiates a new DdosPreflight object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDdosPreflight() *DdosPreflight {
	this := DdosPreflight{}
	return &this
}

// NewDdosPreflightWithDefaults instantiates a new DdosPreflight object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDdosPreflightWithDefaults() *DdosPreflight {
	this := DdosPreflight{}
	return &this
}

// GetAccessOrigin returns the AccessOrigin field value if set, zero value otherwise.
func (o *DdosPreflight) GetAccessOrigin() string {
	if o == nil || IsNil(o.AccessOrigin) {
		var ret string
		return ret
	}
	return *o.AccessOrigin
}

// GetAccessOriginOk returns a tuple with the AccessOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosPreflight) GetAccessOriginOk() (*string, bool) {
	if o == nil || IsNil(o.AccessOrigin) {
		return nil, false
	}
	return o.AccessOrigin, true
}

// HasAccessOrigin returns a boolean if a field has been set.
func (o *DdosPreflight) HasAccessOrigin() bool {
	if o != nil && !IsNil(o.AccessOrigin) {
		return true
	}

	return false
}

// SetAccessOrigin gets a reference to the given string and assigns it to the AccessOrigin field.
func (o *DdosPreflight) SetAccessOrigin(v string) {
	o.AccessOrigin = &v
}

// GetAccessCredentials returns the AccessCredentials field value if set, zero value otherwise.
func (o *DdosPreflight) GetAccessCredentials() string {
	if o == nil || IsNil(o.AccessCredentials) {
		var ret string
		return ret
	}
	return *o.AccessCredentials
}

// GetAccessCredentialsOk returns a tuple with the AccessCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosPreflight) GetAccessCredentialsOk() (*string, bool) {
	if o == nil || IsNil(o.AccessCredentials) {
		return nil, false
	}
	return o.AccessCredentials, true
}

// HasAccessCredentials returns a boolean if a field has been set.
func (o *DdosPreflight) HasAccessCredentials() bool {
	if o != nil && !IsNil(o.AccessCredentials) {
		return true
	}

	return false
}

// SetAccessCredentials gets a reference to the given string and assigns it to the AccessCredentials field.
func (o *DdosPreflight) SetAccessCredentials(v string) {
	o.AccessCredentials = &v
}

// GetAccessMethods returns the AccessMethods field value if set, zero value otherwise.
func (o *DdosPreflight) GetAccessMethods() []string {
	if o == nil || IsNil(o.AccessMethods) {
		var ret []string
		return ret
	}
	return o.AccessMethods
}

// GetAccessMethodsOk returns a tuple with the AccessMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosPreflight) GetAccessMethodsOk() ([]string, bool) {
	if o == nil || IsNil(o.AccessMethods) {
		return nil, false
	}
	return o.AccessMethods, true
}

// HasAccessMethods returns a boolean if a field has been set.
func (o *DdosPreflight) HasAccessMethods() bool {
	if o != nil && !IsNil(o.AccessMethods) {
		return true
	}

	return false
}

// SetAccessMethods gets a reference to the given []string and assigns it to the AccessMethods field.
func (o *DdosPreflight) SetAccessMethods(v []string) {
	o.AccessMethods = v
}

// GetAccessHeaders returns the AccessHeaders field value if set, zero value otherwise.
func (o *DdosPreflight) GetAccessHeaders() string {
	if o == nil || IsNil(o.AccessHeaders) {
		var ret string
		return ret
	}
	return *o.AccessHeaders
}

// GetAccessHeadersOk returns a tuple with the AccessHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosPreflight) GetAccessHeadersOk() (*string, bool) {
	if o == nil || IsNil(o.AccessHeaders) {
		return nil, false
	}
	return o.AccessHeaders, true
}

// HasAccessHeaders returns a boolean if a field has been set.
func (o *DdosPreflight) HasAccessHeaders() bool {
	if o != nil && !IsNil(o.AccessHeaders) {
		return true
	}

	return false
}

// SetAccessHeaders gets a reference to the given string and assigns it to the AccessHeaders field.
func (o *DdosPreflight) SetAccessHeaders(v string) {
	o.AccessHeaders = &v
}

// GetAccessExposeHeaders returns the AccessExposeHeaders field value if set, zero value otherwise.
func (o *DdosPreflight) GetAccessExposeHeaders() string {
	if o == nil || IsNil(o.AccessExposeHeaders) {
		var ret string
		return ret
	}
	return *o.AccessExposeHeaders
}

// GetAccessExposeHeadersOk returns a tuple with the AccessExposeHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosPreflight) GetAccessExposeHeadersOk() (*string, bool) {
	if o == nil || IsNil(o.AccessExposeHeaders) {
		return nil, false
	}
	return o.AccessExposeHeaders, true
}

// HasAccessExposeHeaders returns a boolean if a field has been set.
func (o *DdosPreflight) HasAccessExposeHeaders() bool {
	if o != nil && !IsNil(o.AccessExposeHeaders) {
		return true
	}

	return false
}

// SetAccessExposeHeaders gets a reference to the given string and assigns it to the AccessExposeHeaders field.
func (o *DdosPreflight) SetAccessExposeHeaders(v string) {
	o.AccessExposeHeaders = &v
}

func (o DdosPreflight) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DdosPreflight) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessOrigin) {
		toSerialize["access_origin"] = o.AccessOrigin
	}
	if !IsNil(o.AccessCredentials) {
		toSerialize["access_credentials"] = o.AccessCredentials
	}
	if !IsNil(o.AccessMethods) {
		toSerialize["access_methods"] = o.AccessMethods
	}
	if !IsNil(o.AccessHeaders) {
		toSerialize["access_headers"] = o.AccessHeaders
	}
	if !IsNil(o.AccessExposeHeaders) {
		toSerialize["access_expose_headers"] = o.AccessExposeHeaders
	}
	return toSerialize, nil
}

type NullableDdosPreflight struct {
	value *DdosPreflight
	isSet bool
}

func (v NullableDdosPreflight) Get() *DdosPreflight {
	return v.value
}

func (v *NullableDdosPreflight) Set(val *DdosPreflight) {
	v.value = val
	v.isSet = true
}

func (v NullableDdosPreflight) IsSet() bool {
	return v.isSet
}

func (v *NullableDdosPreflight) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDdosPreflight(val *DdosPreflight) *NullableDdosPreflight {
	return &NullableDdosPreflight{value: val, isSet: true}
}

func (v NullableDdosPreflight) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDdosPreflight) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


