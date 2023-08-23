/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.105.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package arvancloud

import (
	"encoding/json"
	"time"
)

// checks if the CertificateOrder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertificateOrder{}

// CertificateOrder struct for CertificateOrder
type CertificateOrder struct {
	Id *string `json:"id,omitempty"`
	OrderId *string `json:"order_id,omitempty"`
	// - `unprocessed` - Order is in the process queue - `canceled` - Order is canceled in favor of a new one with updated subject names - `pending` - Authorization Challenges are set, Validating authorization challenges... - `ready` - Challenges are validated, ready to issue the certificate - `processing` - Issuing Certificate... - `valid` - Certificate is issued successfully, this is the final stage - `invalid` - An Error Occurred, this order cannot proceed anymore, a new order will be created automatically - `terminated` - An Unknown Error occurred, this order cannot proceed anymore, a new order will be created automatically - `killed` - Order failed despite many retries, will not proceed anymore nor retry, needs manual intervention 
	Status *string `json:"status,omitempty"`
	DomainNames []string `json:"domain_names,omitempty"`
	Errors []map[string]interface{} `json:"errors,omitempty"`
	// Expired order is treated as invalid order
	ExpiryDate *time.Time `json:"expiry_date,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewCertificateOrder instantiates a new CertificateOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateOrder() *CertificateOrder {
	this := CertificateOrder{}
	return &this
}

// NewCertificateOrderWithDefaults instantiates a new CertificateOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateOrderWithDefaults() *CertificateOrder {
	this := CertificateOrder{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CertificateOrder) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateOrder) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CertificateOrder) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CertificateOrder) SetId(v string) {
	o.Id = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *CertificateOrder) GetOrderId() string {
	if o == nil || IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateOrder) GetOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *CertificateOrder) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *CertificateOrder) SetOrderId(v string) {
	o.OrderId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CertificateOrder) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateOrder) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CertificateOrder) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CertificateOrder) SetStatus(v string) {
	o.Status = &v
}

// GetDomainNames returns the DomainNames field value if set, zero value otherwise.
func (o *CertificateOrder) GetDomainNames() []string {
	if o == nil || IsNil(o.DomainNames) {
		var ret []string
		return ret
	}
	return o.DomainNames
}

// GetDomainNamesOk returns a tuple with the DomainNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateOrder) GetDomainNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.DomainNames) {
		return nil, false
	}
	return o.DomainNames, true
}

// HasDomainNames returns a boolean if a field has been set.
func (o *CertificateOrder) HasDomainNames() bool {
	if o != nil && !IsNil(o.DomainNames) {
		return true
	}

	return false
}

// SetDomainNames gets a reference to the given []string and assigns it to the DomainNames field.
func (o *CertificateOrder) SetDomainNames(v []string) {
	o.DomainNames = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *CertificateOrder) GetErrors() []map[string]interface{} {
	if o == nil || IsNil(o.Errors) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateOrder) GetErrorsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *CertificateOrder) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []map[string]interface{} and assigns it to the Errors field.
func (o *CertificateOrder) SetErrors(v []map[string]interface{}) {
	o.Errors = v
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *CertificateOrder) GetExpiryDate() time.Time {
	if o == nil || IsNil(o.ExpiryDate) {
		var ret time.Time
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateOrder) GetExpiryDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiryDate) {
		return nil, false
	}
	return o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *CertificateOrder) HasExpiryDate() bool {
	if o != nil && !IsNil(o.ExpiryDate) {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.
func (o *CertificateOrder) SetExpiryDate(v time.Time) {
	o.ExpiryDate = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CertificateOrder) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateOrder) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CertificateOrder) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CertificateOrder) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CertificateOrder) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateOrder) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CertificateOrder) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *CertificateOrder) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o CertificateOrder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificateOrder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	// skip: order_id is readOnly
	// skip: status is readOnly
	// skip: domain_names is readOnly
	// skip: errors is readOnly
	// skip: expiry_date is readOnly
	// skip: created_at is readOnly
	// skip: updated_at is readOnly
	return toSerialize, nil
}

type NullableCertificateOrder struct {
	value *CertificateOrder
	isSet bool
}

func (v NullableCertificateOrder) Get() *CertificateOrder {
	return v.value
}

func (v *NullableCertificateOrder) Set(val *CertificateOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateOrder(val *CertificateOrder) *NullableCertificateOrder {
	return &NullableCertificateOrder{value: val, isSet: true}
}

func (v NullableCertificateOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


