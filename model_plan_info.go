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

// checks if the PlanInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlanInfo{}

// PlanInfo struct for PlanInfo
type PlanInfo struct {
	Key *string `json:"key,omitempty"`
	Name *string `json:"name,omitempty"`
	MonthlyCost *float32 `json:"monthly_cost,omitempty"`
	// between 0 to 100 is the percentage of the discount
	Discount *float32 `json:"discount,omitempty"`
	// How much balance the account needs for selected plan
	NeededBalance *float32 `json:"needed_balance,omitempty"`
}

// NewPlanInfo instantiates a new PlanInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlanInfo() *PlanInfo {
	this := PlanInfo{}
	return &this
}

// NewPlanInfoWithDefaults instantiates a new PlanInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlanInfoWithDefaults() *PlanInfo {
	this := PlanInfo{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *PlanInfo) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanInfo) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *PlanInfo) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *PlanInfo) SetKey(v string) {
	o.Key = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PlanInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PlanInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PlanInfo) SetName(v string) {
	o.Name = &v
}

// GetMonthlyCost returns the MonthlyCost field value if set, zero value otherwise.
func (o *PlanInfo) GetMonthlyCost() float32 {
	if o == nil || IsNil(o.MonthlyCost) {
		var ret float32
		return ret
	}
	return *o.MonthlyCost
}

// GetMonthlyCostOk returns a tuple with the MonthlyCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanInfo) GetMonthlyCostOk() (*float32, bool) {
	if o == nil || IsNil(o.MonthlyCost) {
		return nil, false
	}
	return o.MonthlyCost, true
}

// HasMonthlyCost returns a boolean if a field has been set.
func (o *PlanInfo) HasMonthlyCost() bool {
	if o != nil && !IsNil(o.MonthlyCost) {
		return true
	}

	return false
}

// SetMonthlyCost gets a reference to the given float32 and assigns it to the MonthlyCost field.
func (o *PlanInfo) SetMonthlyCost(v float32) {
	o.MonthlyCost = &v
}

// GetDiscount returns the Discount field value if set, zero value otherwise.
func (o *PlanInfo) GetDiscount() float32 {
	if o == nil || IsNil(o.Discount) {
		var ret float32
		return ret
	}
	return *o.Discount
}

// GetDiscountOk returns a tuple with the Discount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanInfo) GetDiscountOk() (*float32, bool) {
	if o == nil || IsNil(o.Discount) {
		return nil, false
	}
	return o.Discount, true
}

// HasDiscount returns a boolean if a field has been set.
func (o *PlanInfo) HasDiscount() bool {
	if o != nil && !IsNil(o.Discount) {
		return true
	}

	return false
}

// SetDiscount gets a reference to the given float32 and assigns it to the Discount field.
func (o *PlanInfo) SetDiscount(v float32) {
	o.Discount = &v
}

// GetNeededBalance returns the NeededBalance field value if set, zero value otherwise.
func (o *PlanInfo) GetNeededBalance() float32 {
	if o == nil || IsNil(o.NeededBalance) {
		var ret float32
		return ret
	}
	return *o.NeededBalance
}

// GetNeededBalanceOk returns a tuple with the NeededBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanInfo) GetNeededBalanceOk() (*float32, bool) {
	if o == nil || IsNil(o.NeededBalance) {
		return nil, false
	}
	return o.NeededBalance, true
}

// HasNeededBalance returns a boolean if a field has been set.
func (o *PlanInfo) HasNeededBalance() bool {
	if o != nil && !IsNil(o.NeededBalance) {
		return true
	}

	return false
}

// SetNeededBalance gets a reference to the given float32 and assigns it to the NeededBalance field.
func (o *PlanInfo) SetNeededBalance(v float32) {
	o.NeededBalance = &v
}

func (o PlanInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlanInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.MonthlyCost) {
		toSerialize["monthly_cost"] = o.MonthlyCost
	}
	if !IsNil(o.Discount) {
		toSerialize["discount"] = o.Discount
	}
	if !IsNil(o.NeededBalance) {
		toSerialize["needed_balance"] = o.NeededBalance
	}
	return toSerialize, nil
}

type NullablePlanInfo struct {
	value *PlanInfo
	isSet bool
}

func (v NullablePlanInfo) Get() *PlanInfo {
	return v.value
}

func (v *NullablePlanInfo) Set(val *PlanInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePlanInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePlanInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlanInfo(val *PlanInfo) *NullablePlanInfo {
	return &NullablePlanInfo{value: val, isSet: true}
}

func (v NullablePlanInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlanInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


