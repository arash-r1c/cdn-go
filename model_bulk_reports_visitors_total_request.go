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

// checks if the BulkReportsVisitorsTotalRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkReportsVisitorsTotalRequest{}

// BulkReportsVisitorsTotalRequest struct for BulkReportsVisitorsTotalRequest
type BulkReportsVisitorsTotalRequest struct {
	// List of domains' IDs
	Domains []string `json:"domains,omitempty"`
	// Whether to include sub-domains or report only root domain traffic
	ExcludeSubdomains interface{} `json:"excludeSubdomains,omitempty"`
}

// NewBulkReportsVisitorsTotalRequest instantiates a new BulkReportsVisitorsTotalRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkReportsVisitorsTotalRequest() *BulkReportsVisitorsTotalRequest {
	this := BulkReportsVisitorsTotalRequest{}
	return &this
}

// NewBulkReportsVisitorsTotalRequestWithDefaults instantiates a new BulkReportsVisitorsTotalRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkReportsVisitorsTotalRequestWithDefaults() *BulkReportsVisitorsTotalRequest {
	this := BulkReportsVisitorsTotalRequest{}
	return &this
}

// GetDomains returns the Domains field value if set, zero value otherwise.
func (o *BulkReportsVisitorsTotalRequest) GetDomains() []string {
	if o == nil || IsNil(o.Domains) {
		var ret []string
		return ret
	}
	return o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkReportsVisitorsTotalRequest) GetDomainsOk() ([]string, bool) {
	if o == nil || IsNil(o.Domains) {
		return nil, false
	}
	return o.Domains, true
}

// HasDomains returns a boolean if a field has been set.
func (o *BulkReportsVisitorsTotalRequest) HasDomains() bool {
	if o != nil && !IsNil(o.Domains) {
		return true
	}

	return false
}

// SetDomains gets a reference to the given []string and assigns it to the Domains field.
func (o *BulkReportsVisitorsTotalRequest) SetDomains(v []string) {
	o.Domains = v
}

// GetExcludeSubdomains returns the ExcludeSubdomains field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BulkReportsVisitorsTotalRequest) GetExcludeSubdomains() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.ExcludeSubdomains
}

// GetExcludeSubdomainsOk returns a tuple with the ExcludeSubdomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkReportsVisitorsTotalRequest) GetExcludeSubdomainsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.ExcludeSubdomains) {
		return nil, false
	}
	return &o.ExcludeSubdomains, true
}

// HasExcludeSubdomains returns a boolean if a field has been set.
func (o *BulkReportsVisitorsTotalRequest) HasExcludeSubdomains() bool {
	if o != nil && IsNil(o.ExcludeSubdomains) {
		return true
	}

	return false
}

// SetExcludeSubdomains gets a reference to the given interface{} and assigns it to the ExcludeSubdomains field.
func (o *BulkReportsVisitorsTotalRequest) SetExcludeSubdomains(v interface{}) {
	o.ExcludeSubdomains = v
}

func (o BulkReportsVisitorsTotalRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkReportsVisitorsTotalRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Domains) {
		toSerialize["domains"] = o.Domains
	}
	if o.ExcludeSubdomains != nil {
		toSerialize["excludeSubdomains"] = o.ExcludeSubdomains
	}
	return toSerialize, nil
}

type NullableBulkReportsVisitorsTotalRequest struct {
	value *BulkReportsVisitorsTotalRequest
	isSet bool
}

func (v NullableBulkReportsVisitorsTotalRequest) Get() *BulkReportsVisitorsTotalRequest {
	return v.value
}

func (v *NullableBulkReportsVisitorsTotalRequest) Set(val *BulkReportsVisitorsTotalRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkReportsVisitorsTotalRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkReportsVisitorsTotalRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkReportsVisitorsTotalRequest(val *BulkReportsVisitorsTotalRequest) *NullableBulkReportsVisitorsTotalRequest {
	return &NullableBulkReportsVisitorsTotalRequest{value: val, isSet: true}
}

func (v NullableBulkReportsVisitorsTotalRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkReportsVisitorsTotalRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


