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

// checks if the TransferDomain type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransferDomain{}

// TransferDomain struct for TransferDomain
type TransferDomain struct {
	AccountId string `json:"account_id"`
}

// NewTransferDomain instantiates a new TransferDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferDomain(accountId string) *TransferDomain {
	this := TransferDomain{}
	this.AccountId = accountId
	return &this
}

// NewTransferDomainWithDefaults instantiates a new TransferDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferDomainWithDefaults() *TransferDomain {
	this := TransferDomain{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *TransferDomain) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *TransferDomain) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *TransferDomain) SetAccountId(v string) {
	o.AccountId = v
}

func (o TransferDomain) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferDomain) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_id"] = o.AccountId
	return toSerialize, nil
}

type NullableTransferDomain struct {
	value *TransferDomain
	isSet bool
}

func (v NullableTransferDomain) Get() *TransferDomain {
	return v.value
}

func (v *NullableTransferDomain) Set(val *TransferDomain) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferDomain(val *TransferDomain) *NullableTransferDomain {
	return &NullableTransferDomain{value: val, isSet: true}
}

func (v NullableTransferDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


