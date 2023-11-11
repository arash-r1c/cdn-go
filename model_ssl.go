/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.114.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package arvancloud

import (
	"encoding/json"
)

// checks if the Ssl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Ssl{}

// Ssl struct for Ssl
type Ssl struct {
	// Whether Domain is using fingerprint or not
	FingerprintStatus *bool `json:"fingerprint_status,omitempty"`
	// Whether Domain is using ssl module or not
	SslStatus *bool `json:"ssl_status,omitempty"`
	// Indicates certificate is managed by arvan, or its up to the user
	CertificateMode *string `json:"certificate_mode,omitempty"`
	// Minimum version of TLS. Empty ('') means default.
	TlsVersion *string `json:"tls_version,omitempty"`
	// Whether HSTS is enabled
	HstsStatus *bool `json:"hsts_status,omitempty"`
	// HSTS max age directive
	HstsMaxAge *string `json:"hsts_max_age,omitempty"`
	HstsSubdomain *bool `json:"hsts_subdomain,omitempty"`
	HstsPreload *bool `json:"hsts_preload,omitempty"`
	HttpsRedirect *bool `json:"https_redirect,omitempty"`
	// Replace HTTP with HTTPs in HTML and JS sources
	ReplaceHttp *bool `json:"replace_http,omitempty"`
	CertificateKeyType *string `json:"certificate_key_type,omitempty"`
	Certificates []Certificate `json:"certificates,omitempty"`
	// returns all \"certificate orders\" since the last invalid or canceled order
	Orders []CertificateOrder `json:"orders,omitempty"`
}

// NewSsl instantiates a new Ssl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSsl() *Ssl {
	this := Ssl{}
	return &this
}

// NewSslWithDefaults instantiates a new Ssl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSslWithDefaults() *Ssl {
	this := Ssl{}
	return &this
}

// GetFingerprintStatus returns the FingerprintStatus field value if set, zero value otherwise.
func (o *Ssl) GetFingerprintStatus() bool {
	if o == nil || IsNil(o.FingerprintStatus) {
		var ret bool
		return ret
	}
	return *o.FingerprintStatus
}

// GetFingerprintStatusOk returns a tuple with the FingerprintStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ssl) GetFingerprintStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.FingerprintStatus) {
		return nil, false
	}
	return o.FingerprintStatus, true
}

// HasFingerprintStatus returns a boolean if a field has been set.
func (o *Ssl) HasFingerprintStatus() bool {
	if o != nil && !IsNil(o.FingerprintStatus) {
		return true
	}

	return false
}

// SetFingerprintStatus gets a reference to the given bool and assigns it to the FingerprintStatus field.
func (o *Ssl) SetFingerprintStatus(v bool) {
	o.FingerprintStatus = &v
}

// GetSslStatus returns the SslStatus field value if set, zero value otherwise.
func (o *Ssl) GetSslStatus() bool {
	if o == nil || IsNil(o.SslStatus) {
		var ret bool
		return ret
	}
	return *o.SslStatus
}

// GetSslStatusOk returns a tuple with the SslStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ssl) GetSslStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.SslStatus) {
		return nil, false
	}
	return o.SslStatus, true
}

// HasSslStatus returns a boolean if a field has been set.
func (o *Ssl) HasSslStatus() bool {
	if o != nil && !IsNil(o.SslStatus) {
		return true
	}

	return false
}

// SetSslStatus gets a reference to the given bool and assigns it to the SslStatus field.
func (o *Ssl) SetSslStatus(v bool) {
	o.SslStatus = &v
}

// GetCertificateMode returns the CertificateMode field value if set, zero value otherwise.
func (o *Ssl) GetCertificateMode() string {
	if o == nil || IsNil(o.CertificateMode) {
		var ret string
		return ret
	}
	return *o.CertificateMode
}

// GetCertificateModeOk returns a tuple with the CertificateMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ssl) GetCertificateModeOk() (*string, bool) {
	if o == nil || IsNil(o.CertificateMode) {
		return nil, false
	}
	return o.CertificateMode, true
}

// HasCertificateMode returns a boolean if a field has been set.
func (o *Ssl) HasCertificateMode() bool {
	if o != nil && !IsNil(o.CertificateMode) {
		return true
	}

	return false
}

// SetCertificateMode gets a reference to the given string and assigns it to the CertificateMode field.
func (o *Ssl) SetCertificateMode(v string) {
	o.CertificateMode = &v
}

// GetTlsVersion returns the TlsVersion field value if set, zero value otherwise.
func (o *Ssl) GetTlsVersion() string {
	if o == nil || IsNil(o.TlsVersion) {
		var ret string
		return ret
	}
	return *o.TlsVersion
}

// GetTlsVersionOk returns a tuple with the TlsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ssl) GetTlsVersionOk() (*string, bool) {
	if o == nil || IsNil(o.TlsVersion) {
		return nil, false
	}
	return o.TlsVersion, true
}

// HasTlsVersion returns a boolean if a field has been set.
func (o *Ssl) HasTlsVersion() bool {
	if o != nil && !IsNil(o.TlsVersion) {
		return true
	}

	return false
}

// SetTlsVersion gets a reference to the given string and assigns it to the TlsVersion field.
func (o *Ssl) SetTlsVersion(v string) {
	o.TlsVersion = &v
}

// GetHstsStatus returns the HstsStatus field value if set, zero value otherwise.
func (o *Ssl) GetHstsStatus() bool {
	if o == nil || IsNil(o.HstsStatus) {
		var ret bool
		return ret
	}
	return *o.HstsStatus
}

// GetHstsStatusOk returns a tuple with the HstsStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ssl) GetHstsStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.HstsStatus) {
		return nil, false
	}
	return o.HstsStatus, true
}

// HasHstsStatus returns a boolean if a field has been set.
func (o *Ssl) HasHstsStatus() bool {
	if o != nil && !IsNil(o.HstsStatus) {
		return true
	}

	return false
}

// SetHstsStatus gets a reference to the given bool and assigns it to the HstsStatus field.
func (o *Ssl) SetHstsStatus(v bool) {
	o.HstsStatus = &v
}

// GetHstsMaxAge returns the HstsMaxAge field value if set, zero value otherwise.
func (o *Ssl) GetHstsMaxAge() string {
	if o == nil || IsNil(o.HstsMaxAge) {
		var ret string
		return ret
	}
	return *o.HstsMaxAge
}

// GetHstsMaxAgeOk returns a tuple with the HstsMaxAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ssl) GetHstsMaxAgeOk() (*string, bool) {
	if o == nil || IsNil(o.HstsMaxAge) {
		return nil, false
	}
	return o.HstsMaxAge, true
}

// HasHstsMaxAge returns a boolean if a field has been set.
func (o *Ssl) HasHstsMaxAge() bool {
	if o != nil && !IsNil(o.HstsMaxAge) {
		return true
	}

	return false
}

// SetHstsMaxAge gets a reference to the given string and assigns it to the HstsMaxAge field.
func (o *Ssl) SetHstsMaxAge(v string) {
	o.HstsMaxAge = &v
}

// GetHstsSubdomain returns the HstsSubdomain field value if set, zero value otherwise.
func (o *Ssl) GetHstsSubdomain() bool {
	if o == nil || IsNil(o.HstsSubdomain) {
		var ret bool
		return ret
	}
	return *o.HstsSubdomain
}

// GetHstsSubdomainOk returns a tuple with the HstsSubdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ssl) GetHstsSubdomainOk() (*bool, bool) {
	if o == nil || IsNil(o.HstsSubdomain) {
		return nil, false
	}
	return o.HstsSubdomain, true
}

// HasHstsSubdomain returns a boolean if a field has been set.
func (o *Ssl) HasHstsSubdomain() bool {
	if o != nil && !IsNil(o.HstsSubdomain) {
		return true
	}

	return false
}

// SetHstsSubdomain gets a reference to the given bool and assigns it to the HstsSubdomain field.
func (o *Ssl) SetHstsSubdomain(v bool) {
	o.HstsSubdomain = &v
}

// GetHstsPreload returns the HstsPreload field value if set, zero value otherwise.
func (o *Ssl) GetHstsPreload() bool {
	if o == nil || IsNil(o.HstsPreload) {
		var ret bool
		return ret
	}
	return *o.HstsPreload
}

// GetHstsPreloadOk returns a tuple with the HstsPreload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ssl) GetHstsPreloadOk() (*bool, bool) {
	if o == nil || IsNil(o.HstsPreload) {
		return nil, false
	}
	return o.HstsPreload, true
}

// HasHstsPreload returns a boolean if a field has been set.
func (o *Ssl) HasHstsPreload() bool {
	if o != nil && !IsNil(o.HstsPreload) {
		return true
	}

	return false
}

// SetHstsPreload gets a reference to the given bool and assigns it to the HstsPreload field.
func (o *Ssl) SetHstsPreload(v bool) {
	o.HstsPreload = &v
}

// GetHttpsRedirect returns the HttpsRedirect field value if set, zero value otherwise.
func (o *Ssl) GetHttpsRedirect() bool {
	if o == nil || IsNil(o.HttpsRedirect) {
		var ret bool
		return ret
	}
	return *o.HttpsRedirect
}

// GetHttpsRedirectOk returns a tuple with the HttpsRedirect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ssl) GetHttpsRedirectOk() (*bool, bool) {
	if o == nil || IsNil(o.HttpsRedirect) {
		return nil, false
	}
	return o.HttpsRedirect, true
}

// HasHttpsRedirect returns a boolean if a field has been set.
func (o *Ssl) HasHttpsRedirect() bool {
	if o != nil && !IsNil(o.HttpsRedirect) {
		return true
	}

	return false
}

// SetHttpsRedirect gets a reference to the given bool and assigns it to the HttpsRedirect field.
func (o *Ssl) SetHttpsRedirect(v bool) {
	o.HttpsRedirect = &v
}

// GetReplaceHttp returns the ReplaceHttp field value if set, zero value otherwise.
func (o *Ssl) GetReplaceHttp() bool {
	if o == nil || IsNil(o.ReplaceHttp) {
		var ret bool
		return ret
	}
	return *o.ReplaceHttp
}

// GetReplaceHttpOk returns a tuple with the ReplaceHttp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ssl) GetReplaceHttpOk() (*bool, bool) {
	if o == nil || IsNil(o.ReplaceHttp) {
		return nil, false
	}
	return o.ReplaceHttp, true
}

// HasReplaceHttp returns a boolean if a field has been set.
func (o *Ssl) HasReplaceHttp() bool {
	if o != nil && !IsNil(o.ReplaceHttp) {
		return true
	}

	return false
}

// SetReplaceHttp gets a reference to the given bool and assigns it to the ReplaceHttp field.
func (o *Ssl) SetReplaceHttp(v bool) {
	o.ReplaceHttp = &v
}

// GetCertificateKeyType returns the CertificateKeyType field value if set, zero value otherwise.
func (o *Ssl) GetCertificateKeyType() string {
	if o == nil || IsNil(o.CertificateKeyType) {
		var ret string
		return ret
	}
	return *o.CertificateKeyType
}

// GetCertificateKeyTypeOk returns a tuple with the CertificateKeyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ssl) GetCertificateKeyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CertificateKeyType) {
		return nil, false
	}
	return o.CertificateKeyType, true
}

// HasCertificateKeyType returns a boolean if a field has been set.
func (o *Ssl) HasCertificateKeyType() bool {
	if o != nil && !IsNil(o.CertificateKeyType) {
		return true
	}

	return false
}

// SetCertificateKeyType gets a reference to the given string and assigns it to the CertificateKeyType field.
func (o *Ssl) SetCertificateKeyType(v string) {
	o.CertificateKeyType = &v
}

// GetCertificates returns the Certificates field value if set, zero value otherwise.
func (o *Ssl) GetCertificates() []Certificate {
	if o == nil || IsNil(o.Certificates) {
		var ret []Certificate
		return ret
	}
	return o.Certificates
}

// GetCertificatesOk returns a tuple with the Certificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ssl) GetCertificatesOk() ([]Certificate, bool) {
	if o == nil || IsNil(o.Certificates) {
		return nil, false
	}
	return o.Certificates, true
}

// HasCertificates returns a boolean if a field has been set.
func (o *Ssl) HasCertificates() bool {
	if o != nil && !IsNil(o.Certificates) {
		return true
	}

	return false
}

// SetCertificates gets a reference to the given []Certificate and assigns it to the Certificates field.
func (o *Ssl) SetCertificates(v []Certificate) {
	o.Certificates = v
}

// GetOrders returns the Orders field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Ssl) GetOrders() []CertificateOrder {
	if o == nil {
		var ret []CertificateOrder
		return ret
	}
	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Ssl) GetOrdersOk() ([]CertificateOrder, bool) {
	if o == nil || IsNil(o.Orders) {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *Ssl) HasOrders() bool {
	if o != nil && IsNil(o.Orders) {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []CertificateOrder and assigns it to the Orders field.
func (o *Ssl) SetOrders(v []CertificateOrder) {
	o.Orders = v
}

func (o Ssl) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Ssl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FingerprintStatus) {
		toSerialize["fingerprint_status"] = o.FingerprintStatus
	}
	if !IsNil(o.SslStatus) {
		toSerialize["ssl_status"] = o.SslStatus
	}
	// skip: certificate_mode is readOnly
	if !IsNil(o.TlsVersion) {
		toSerialize["tls_version"] = o.TlsVersion
	}
	if !IsNil(o.HstsStatus) {
		toSerialize["hsts_status"] = o.HstsStatus
	}
	if !IsNil(o.HstsMaxAge) {
		toSerialize["hsts_max_age"] = o.HstsMaxAge
	}
	if !IsNil(o.HstsSubdomain) {
		toSerialize["hsts_subdomain"] = o.HstsSubdomain
	}
	if !IsNil(o.HstsPreload) {
		toSerialize["hsts_preload"] = o.HstsPreload
	}
	if !IsNil(o.HttpsRedirect) {
		toSerialize["https_redirect"] = o.HttpsRedirect
	}
	if !IsNil(o.ReplaceHttp) {
		toSerialize["replace_http"] = o.ReplaceHttp
	}
	if !IsNil(o.CertificateKeyType) {
		toSerialize["certificate_key_type"] = o.CertificateKeyType
	}
	// skip: certificates is readOnly
	if o.Orders != nil {
		toSerialize["orders"] = o.Orders
	}
	return toSerialize, nil
}

type NullableSsl struct {
	value *Ssl
	isSet bool
}

func (v NullableSsl) Get() *Ssl {
	return v.value
}

func (v *NullableSsl) Set(val *Ssl) {
	v.value = val
	v.isSet = true
}

func (v NullableSsl) IsSet() bool {
	return v.isSet
}

func (v *NullableSsl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsl(val *Ssl) *NullableSsl {
	return &NullableSsl{value: val, isSet: true}
}

func (v NullableSsl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


