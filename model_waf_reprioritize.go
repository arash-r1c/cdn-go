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

// checks if the WafReprioritize type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WafReprioritize{}

// WafReprioritize struct for WafReprioritize
type WafReprioritize struct {
	// ID of the package you want to move
	PackageId string `json:"package_id"`
	// ID of the package you want to be prior to the selected package
	AfterPackageId *string `json:"after_package_id,omitempty"`
	// ID of the package you want to follow the selected package
	BeforePackageId *string `json:"before_package_id,omitempty"`
}

// NewWafReprioritize instantiates a new WafReprioritize object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafReprioritize(packageId string) *WafReprioritize {
	this := WafReprioritize{}
	this.PackageId = packageId
	return &this
}

// NewWafReprioritizeWithDefaults instantiates a new WafReprioritize object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafReprioritizeWithDefaults() *WafReprioritize {
	this := WafReprioritize{}
	return &this
}

// GetPackageId returns the PackageId field value
func (o *WafReprioritize) GetPackageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PackageId
}

// GetPackageIdOk returns a tuple with the PackageId field value
// and a boolean to check if the value has been set.
func (o *WafReprioritize) GetPackageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageId, true
}

// SetPackageId sets field value
func (o *WafReprioritize) SetPackageId(v string) {
	o.PackageId = v
}

// GetAfterPackageId returns the AfterPackageId field value if set, zero value otherwise.
func (o *WafReprioritize) GetAfterPackageId() string {
	if o == nil || IsNil(o.AfterPackageId) {
		var ret string
		return ret
	}
	return *o.AfterPackageId
}

// GetAfterPackageIdOk returns a tuple with the AfterPackageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafReprioritize) GetAfterPackageIdOk() (*string, bool) {
	if o == nil || IsNil(o.AfterPackageId) {
		return nil, false
	}
	return o.AfterPackageId, true
}

// HasAfterPackageId returns a boolean if a field has been set.
func (o *WafReprioritize) HasAfterPackageId() bool {
	if o != nil && !IsNil(o.AfterPackageId) {
		return true
	}

	return false
}

// SetAfterPackageId gets a reference to the given string and assigns it to the AfterPackageId field.
func (o *WafReprioritize) SetAfterPackageId(v string) {
	o.AfterPackageId = &v
}

// GetBeforePackageId returns the BeforePackageId field value if set, zero value otherwise.
func (o *WafReprioritize) GetBeforePackageId() string {
	if o == nil || IsNil(o.BeforePackageId) {
		var ret string
		return ret
	}
	return *o.BeforePackageId
}

// GetBeforePackageIdOk returns a tuple with the BeforePackageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafReprioritize) GetBeforePackageIdOk() (*string, bool) {
	if o == nil || IsNil(o.BeforePackageId) {
		return nil, false
	}
	return o.BeforePackageId, true
}

// HasBeforePackageId returns a boolean if a field has been set.
func (o *WafReprioritize) HasBeforePackageId() bool {
	if o != nil && !IsNil(o.BeforePackageId) {
		return true
	}

	return false
}

// SetBeforePackageId gets a reference to the given string and assigns it to the BeforePackageId field.
func (o *WafReprioritize) SetBeforePackageId(v string) {
	o.BeforePackageId = &v
}

func (o WafReprioritize) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WafReprioritize) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["package_id"] = o.PackageId
	if !IsNil(o.AfterPackageId) {
		toSerialize["after_package_id"] = o.AfterPackageId
	}
	if !IsNil(o.BeforePackageId) {
		toSerialize["before_package_id"] = o.BeforePackageId
	}
	return toSerialize, nil
}

type NullableWafReprioritize struct {
	value *WafReprioritize
	isSet bool
}

func (v NullableWafReprioritize) Get() *WafReprioritize {
	return v.value
}

func (v *NullableWafReprioritize) Set(val *WafReprioritize) {
	v.value = val
	v.isSet = true
}

func (v NullableWafReprioritize) IsSet() bool {
	return v.isSet
}

func (v *NullableWafReprioritize) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWafReprioritize(val *WafReprioritize) *NullableWafReprioritize {
	return &NullableWafReprioritize{value: val, isSet: true}
}

func (v NullableWafReprioritize) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWafReprioritize) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


