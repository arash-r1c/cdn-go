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

// checks if the ImageResize type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageResize{}

// ImageResize struct for ImageResize
type ImageResize struct {
	Status *string `json:"status,omitempty"`
	HeightBy *string `json:"height_by,omitempty"`
	WidthBy *string `json:"width_by,omitempty"`
}

// NewImageResize instantiates a new ImageResize object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageResize() *ImageResize {
	this := ImageResize{}
	var status string = "off"
	this.Status = &status
	var heightBy string = "height"
	this.HeightBy = &heightBy
	var widthBy string = "width"
	this.WidthBy = &widthBy
	return &this
}

// NewImageResizeWithDefaults instantiates a new ImageResize object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageResizeWithDefaults() *ImageResize {
	this := ImageResize{}
	var status string = "off"
	this.Status = &status
	var heightBy string = "height"
	this.HeightBy = &heightBy
	var widthBy string = "width"
	this.WidthBy = &widthBy
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ImageResize) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageResize) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ImageResize) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ImageResize) SetStatus(v string) {
	o.Status = &v
}

// GetHeightBy returns the HeightBy field value if set, zero value otherwise.
func (o *ImageResize) GetHeightBy() string {
	if o == nil || IsNil(o.HeightBy) {
		var ret string
		return ret
	}
	return *o.HeightBy
}

// GetHeightByOk returns a tuple with the HeightBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageResize) GetHeightByOk() (*string, bool) {
	if o == nil || IsNil(o.HeightBy) {
		return nil, false
	}
	return o.HeightBy, true
}

// HasHeightBy returns a boolean if a field has been set.
func (o *ImageResize) HasHeightBy() bool {
	if o != nil && !IsNil(o.HeightBy) {
		return true
	}

	return false
}

// SetHeightBy gets a reference to the given string and assigns it to the HeightBy field.
func (o *ImageResize) SetHeightBy(v string) {
	o.HeightBy = &v
}

// GetWidthBy returns the WidthBy field value if set, zero value otherwise.
func (o *ImageResize) GetWidthBy() string {
	if o == nil || IsNil(o.WidthBy) {
		var ret string
		return ret
	}
	return *o.WidthBy
}

// GetWidthByOk returns a tuple with the WidthBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageResize) GetWidthByOk() (*string, bool) {
	if o == nil || IsNil(o.WidthBy) {
		return nil, false
	}
	return o.WidthBy, true
}

// HasWidthBy returns a boolean if a field has been set.
func (o *ImageResize) HasWidthBy() bool {
	if o != nil && !IsNil(o.WidthBy) {
		return true
	}

	return false
}

// SetWidthBy gets a reference to the given string and assigns it to the WidthBy field.
func (o *ImageResize) SetWidthBy(v string) {
	o.WidthBy = &v
}

func (o ImageResize) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageResize) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.HeightBy) {
		toSerialize["height_by"] = o.HeightBy
	}
	if !IsNil(o.WidthBy) {
		toSerialize["width_by"] = o.WidthBy
	}
	return toSerialize, nil
}

type NullableImageResize struct {
	value *ImageResize
	isSet bool
}

func (v NullableImageResize) Get() *ImageResize {
	return v.value
}

func (v *NullableImageResize) Set(val *ImageResize) {
	v.value = val
	v.isSet = true
}

func (v NullableImageResize) IsSet() bool {
	return v.isSet
}

func (v *NullableImageResize) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageResize(val *ImageResize) *NullableImageResize {
	return &NullableImageResize{value: val, isSet: true}
}

func (v NullableImageResize) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageResize) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


