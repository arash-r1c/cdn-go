/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.114.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package arvancloud

import (
	"encoding/json"
	"time"
)

// checks if the LoadBalancerUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LoadBalancerUpdate{}

// LoadBalancerUpdate struct for LoadBalancerUpdate
type LoadBalancerUpdate struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Status *bool `json:"status,omitempty"`
	Method *string `json:"method,omitempty"`
	// Human friendly time duration for which a pool will uninterruptedly be selected in cluster_rr strategy, i.e. pools will switch once every time slice.
	TimeSlice *string `json:"time_slice,omitempty"`
	Pools []LoadBalancerUpdatePoolsInner `json:"pools,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewLoadBalancerUpdate instantiates a new LoadBalancerUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadBalancerUpdate() *LoadBalancerUpdate {
	this := LoadBalancerUpdate{}
	var timeSlice string = "0s"
	this.TimeSlice = &timeSlice
	return &this
}

// NewLoadBalancerUpdateWithDefaults instantiates a new LoadBalancerUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadBalancerUpdateWithDefaults() *LoadBalancerUpdate {
	this := LoadBalancerUpdate{}
	var timeSlice string = "0s"
	this.TimeSlice = &timeSlice
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LoadBalancerUpdate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerUpdate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LoadBalancerUpdate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LoadBalancerUpdate) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LoadBalancerUpdate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerUpdate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LoadBalancerUpdate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LoadBalancerUpdate) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LoadBalancerUpdate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LoadBalancerUpdate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LoadBalancerUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *LoadBalancerUpdate) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerUpdate) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *LoadBalancerUpdate) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *LoadBalancerUpdate) SetStatus(v bool) {
	o.Status = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *LoadBalancerUpdate) GetMethod() string {
	if o == nil || IsNil(o.Method) {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerUpdate) GetMethodOk() (*string, bool) {
	if o == nil || IsNil(o.Method) {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *LoadBalancerUpdate) HasMethod() bool {
	if o != nil && !IsNil(o.Method) {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *LoadBalancerUpdate) SetMethod(v string) {
	o.Method = &v
}

// GetTimeSlice returns the TimeSlice field value if set, zero value otherwise.
func (o *LoadBalancerUpdate) GetTimeSlice() string {
	if o == nil || IsNil(o.TimeSlice) {
		var ret string
		return ret
	}
	return *o.TimeSlice
}

// GetTimeSliceOk returns a tuple with the TimeSlice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerUpdate) GetTimeSliceOk() (*string, bool) {
	if o == nil || IsNil(o.TimeSlice) {
		return nil, false
	}
	return o.TimeSlice, true
}

// HasTimeSlice returns a boolean if a field has been set.
func (o *LoadBalancerUpdate) HasTimeSlice() bool {
	if o != nil && !IsNil(o.TimeSlice) {
		return true
	}

	return false
}

// SetTimeSlice gets a reference to the given string and assigns it to the TimeSlice field.
func (o *LoadBalancerUpdate) SetTimeSlice(v string) {
	o.TimeSlice = &v
}

// GetPools returns the Pools field value if set, zero value otherwise.
func (o *LoadBalancerUpdate) GetPools() []LoadBalancerUpdatePoolsInner {
	if o == nil || IsNil(o.Pools) {
		var ret []LoadBalancerUpdatePoolsInner
		return ret
	}
	return o.Pools
}

// GetPoolsOk returns a tuple with the Pools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerUpdate) GetPoolsOk() ([]LoadBalancerUpdatePoolsInner, bool) {
	if o == nil || IsNil(o.Pools) {
		return nil, false
	}
	return o.Pools, true
}

// HasPools returns a boolean if a field has been set.
func (o *LoadBalancerUpdate) HasPools() bool {
	if o != nil && !IsNil(o.Pools) {
		return true
	}

	return false
}

// SetPools gets a reference to the given []LoadBalancerUpdatePoolsInner and assigns it to the Pools field.
func (o *LoadBalancerUpdate) SetPools(v []LoadBalancerUpdatePoolsInner) {
	o.Pools = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *LoadBalancerUpdate) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerUpdate) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *LoadBalancerUpdate) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *LoadBalancerUpdate) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *LoadBalancerUpdate) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadBalancerUpdate) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *LoadBalancerUpdate) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *LoadBalancerUpdate) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o LoadBalancerUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LoadBalancerUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Method) {
		toSerialize["method"] = o.Method
	}
	if !IsNil(o.TimeSlice) {
		toSerialize["time_slice"] = o.TimeSlice
	}
	if !IsNil(o.Pools) {
		toSerialize["pools"] = o.Pools
	}
	// skip: created_at is readOnly
	// skip: updated_at is readOnly
	return toSerialize, nil
}

type NullableLoadBalancerUpdate struct {
	value *LoadBalancerUpdate
	isSet bool
}

func (v NullableLoadBalancerUpdate) Get() *LoadBalancerUpdate {
	return v.value
}

func (v *NullableLoadBalancerUpdate) Set(val *LoadBalancerUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadBalancerUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadBalancerUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadBalancerUpdate(val *LoadBalancerUpdate) *NullableLoadBalancerUpdate {
	return &NullableLoadBalancerUpdate{value: val, isSet: true}
}

func (v NullableLoadBalancerUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadBalancerUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


