/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.105.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package arvancloud

import (
	"encoding/json"
	"fmt"
)

// DnsRecord - struct for DnsRecord
type DnsRecord struct {
	AAAARecord *AAAARecord
	ANAMERecord *ANAMERecord
	ARecord *ARecord
	CAARecord *CAARecord
	CNAMERecord *CNAMERecord
	DKIMRecord *DKIMRecord
	MXRecord *MXRecord
	NSRecord *NSRecord
	PTRRecord *PTRRecord
	SPFRecord *SPFRecord
	SRVRecord *SRVRecord
	TLSARecord *TLSARecord
	TXTRecord *TXTRecord
}

// AAAARecordAsDnsRecord is a convenience function that returns AAAARecord wrapped in DnsRecord
func AAAARecordAsDnsRecord(v *AAAARecord) DnsRecord {
	return DnsRecord{
		AAAARecord: v,
	}
}

// ANAMERecordAsDnsRecord is a convenience function that returns ANAMERecord wrapped in DnsRecord
func ANAMERecordAsDnsRecord(v *ANAMERecord) DnsRecord {
	return DnsRecord{
		ANAMERecord: v,
	}
}

// ARecordAsDnsRecord is a convenience function that returns ARecord wrapped in DnsRecord
func ARecordAsDnsRecord(v *ARecord) DnsRecord {
	return DnsRecord{
		ARecord: v,
	}
}

// CAARecordAsDnsRecord is a convenience function that returns CAARecord wrapped in DnsRecord
func CAARecordAsDnsRecord(v *CAARecord) DnsRecord {
	return DnsRecord{
		CAARecord: v,
	}
}

// CNAMERecordAsDnsRecord is a convenience function that returns CNAMERecord wrapped in DnsRecord
func CNAMERecordAsDnsRecord(v *CNAMERecord) DnsRecord {
	return DnsRecord{
		CNAMERecord: v,
	}
}

// DKIMRecordAsDnsRecord is a convenience function that returns DKIMRecord wrapped in DnsRecord
func DKIMRecordAsDnsRecord(v *DKIMRecord) DnsRecord {
	return DnsRecord{
		DKIMRecord: v,
	}
}

// MXRecordAsDnsRecord is a convenience function that returns MXRecord wrapped in DnsRecord
func MXRecordAsDnsRecord(v *MXRecord) DnsRecord {
	return DnsRecord{
		MXRecord: v,
	}
}

// NSRecordAsDnsRecord is a convenience function that returns NSRecord wrapped in DnsRecord
func NSRecordAsDnsRecord(v *NSRecord) DnsRecord {
	return DnsRecord{
		NSRecord: v,
	}
}

// PTRRecordAsDnsRecord is a convenience function that returns PTRRecord wrapped in DnsRecord
func PTRRecordAsDnsRecord(v *PTRRecord) DnsRecord {
	return DnsRecord{
		PTRRecord: v,
	}
}

// SPFRecordAsDnsRecord is a convenience function that returns SPFRecord wrapped in DnsRecord
func SPFRecordAsDnsRecord(v *SPFRecord) DnsRecord {
	return DnsRecord{
		SPFRecord: v,
	}
}

// SRVRecordAsDnsRecord is a convenience function that returns SRVRecord wrapped in DnsRecord
func SRVRecordAsDnsRecord(v *SRVRecord) DnsRecord {
	return DnsRecord{
		SRVRecord: v,
	}
}

// TLSARecordAsDnsRecord is a convenience function that returns TLSARecord wrapped in DnsRecord
func TLSARecordAsDnsRecord(v *TLSARecord) DnsRecord {
	return DnsRecord{
		TLSARecord: v,
	}
}

// TXTRecordAsDnsRecord is a convenience function that returns TXTRecord wrapped in DnsRecord
func TXTRecordAsDnsRecord(v *TXTRecord) DnsRecord {
	return DnsRecord{
		TXTRecord: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *DnsRecord) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AAAARecord
	err = newStrictDecoder(data).Decode(&dst.AAAARecord)
	if err == nil {
		jsonAAAARecord, _ := json.Marshal(dst.AAAARecord)
		if string(jsonAAAARecord) == "{}" { // empty struct
			dst.AAAARecord = nil
		} else {
			match++
		}
	} else {
		dst.AAAARecord = nil
	}

	// try to unmarshal data into ANAMERecord
	err = newStrictDecoder(data).Decode(&dst.ANAMERecord)
	if err == nil {
		jsonANAMERecord, _ := json.Marshal(dst.ANAMERecord)
		if string(jsonANAMERecord) == "{}" { // empty struct
			dst.ANAMERecord = nil
		} else {
			match++
		}
	} else {
		dst.ANAMERecord = nil
	}

	// try to unmarshal data into ARecord
	err = newStrictDecoder(data).Decode(&dst.ARecord)
	if err == nil {
		jsonARecord, _ := json.Marshal(dst.ARecord)
		if string(jsonARecord) == "{}" { // empty struct
			dst.ARecord = nil
		} else {
			match++
		}
	} else {
		dst.ARecord = nil
	}

	// try to unmarshal data into CAARecord
	err = newStrictDecoder(data).Decode(&dst.CAARecord)
	if err == nil {
		jsonCAARecord, _ := json.Marshal(dst.CAARecord)
		if string(jsonCAARecord) == "{}" { // empty struct
			dst.CAARecord = nil
		} else {
			match++
		}
	} else {
		dst.CAARecord = nil
	}

	// try to unmarshal data into CNAMERecord
	err = newStrictDecoder(data).Decode(&dst.CNAMERecord)
	if err == nil {
		jsonCNAMERecord, _ := json.Marshal(dst.CNAMERecord)
		if string(jsonCNAMERecord) == "{}" { // empty struct
			dst.CNAMERecord = nil
		} else {
			match++
		}
	} else {
		dst.CNAMERecord = nil
	}

	// try to unmarshal data into DKIMRecord
	err = newStrictDecoder(data).Decode(&dst.DKIMRecord)
	if err == nil {
		jsonDKIMRecord, _ := json.Marshal(dst.DKIMRecord)
		if string(jsonDKIMRecord) == "{}" { // empty struct
			dst.DKIMRecord = nil
		} else {
			match++
		}
	} else {
		dst.DKIMRecord = nil
	}

	// try to unmarshal data into MXRecord
	err = newStrictDecoder(data).Decode(&dst.MXRecord)
	if err == nil {
		jsonMXRecord, _ := json.Marshal(dst.MXRecord)
		if string(jsonMXRecord) == "{}" { // empty struct
			dst.MXRecord = nil
		} else {
			match++
		}
	} else {
		dst.MXRecord = nil
	}

	// try to unmarshal data into NSRecord
	err = newStrictDecoder(data).Decode(&dst.NSRecord)
	if err == nil {
		jsonNSRecord, _ := json.Marshal(dst.NSRecord)
		if string(jsonNSRecord) == "{}" { // empty struct
			dst.NSRecord = nil
		} else {
			match++
		}
	} else {
		dst.NSRecord = nil
	}

	// try to unmarshal data into PTRRecord
	err = newStrictDecoder(data).Decode(&dst.PTRRecord)
	if err == nil {
		jsonPTRRecord, _ := json.Marshal(dst.PTRRecord)
		if string(jsonPTRRecord) == "{}" { // empty struct
			dst.PTRRecord = nil
		} else {
			match++
		}
	} else {
		dst.PTRRecord = nil
	}

	// try to unmarshal data into SPFRecord
	err = newStrictDecoder(data).Decode(&dst.SPFRecord)
	if err == nil {
		jsonSPFRecord, _ := json.Marshal(dst.SPFRecord)
		if string(jsonSPFRecord) == "{}" { // empty struct
			dst.SPFRecord = nil
		} else {
			match++
		}
	} else {
		dst.SPFRecord = nil
	}

	// try to unmarshal data into SRVRecord
	err = newStrictDecoder(data).Decode(&dst.SRVRecord)
	if err == nil {
		jsonSRVRecord, _ := json.Marshal(dst.SRVRecord)
		if string(jsonSRVRecord) == "{}" { // empty struct
			dst.SRVRecord = nil
		} else {
			match++
		}
	} else {
		dst.SRVRecord = nil
	}

	// try to unmarshal data into TLSARecord
	err = newStrictDecoder(data).Decode(&dst.TLSARecord)
	if err == nil {
		jsonTLSARecord, _ := json.Marshal(dst.TLSARecord)
		if string(jsonTLSARecord) == "{}" { // empty struct
			dst.TLSARecord = nil
		} else {
			match++
		}
	} else {
		dst.TLSARecord = nil
	}

	// try to unmarshal data into TXTRecord
	err = newStrictDecoder(data).Decode(&dst.TXTRecord)
	if err == nil {
		jsonTXTRecord, _ := json.Marshal(dst.TXTRecord)
		if string(jsonTXTRecord) == "{}" { // empty struct
			dst.TXTRecord = nil
		} else {
			match++
		}
	} else {
		dst.TXTRecord = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AAAARecord = nil
		dst.ANAMERecord = nil
		dst.ARecord = nil
		dst.CAARecord = nil
		dst.CNAMERecord = nil
		dst.DKIMRecord = nil
		dst.MXRecord = nil
		dst.NSRecord = nil
		dst.PTRRecord = nil
		dst.SPFRecord = nil
		dst.SRVRecord = nil
		dst.TLSARecord = nil
		dst.TXTRecord = nil

		return fmt.Errorf("data matches more than one schema in oneOf(DnsRecord)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(DnsRecord)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DnsRecord) MarshalJSON() ([]byte, error) {
	if src.AAAARecord != nil {
		return json.Marshal(&src.AAAARecord)
	}

	if src.ANAMERecord != nil {
		return json.Marshal(&src.ANAMERecord)
	}

	if src.ARecord != nil {
		return json.Marshal(&src.ARecord)
	}

	if src.CAARecord != nil {
		return json.Marshal(&src.CAARecord)
	}

	if src.CNAMERecord != nil {
		return json.Marshal(&src.CNAMERecord)
	}

	if src.DKIMRecord != nil {
		return json.Marshal(&src.DKIMRecord)
	}

	if src.MXRecord != nil {
		return json.Marshal(&src.MXRecord)
	}

	if src.NSRecord != nil {
		return json.Marshal(&src.NSRecord)
	}

	if src.PTRRecord != nil {
		return json.Marshal(&src.PTRRecord)
	}

	if src.SPFRecord != nil {
		return json.Marshal(&src.SPFRecord)
	}

	if src.SRVRecord != nil {
		return json.Marshal(&src.SRVRecord)
	}

	if src.TLSARecord != nil {
		return json.Marshal(&src.TLSARecord)
	}

	if src.TXTRecord != nil {
		return json.Marshal(&src.TXTRecord)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DnsRecord) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AAAARecord != nil {
		return obj.AAAARecord
	}

	if obj.ANAMERecord != nil {
		return obj.ANAMERecord
	}

	if obj.ARecord != nil {
		return obj.ARecord
	}

	if obj.CAARecord != nil {
		return obj.CAARecord
	}

	if obj.CNAMERecord != nil {
		return obj.CNAMERecord
	}

	if obj.DKIMRecord != nil {
		return obj.DKIMRecord
	}

	if obj.MXRecord != nil {
		return obj.MXRecord
	}

	if obj.NSRecord != nil {
		return obj.NSRecord
	}

	if obj.PTRRecord != nil {
		return obj.PTRRecord
	}

	if obj.SPFRecord != nil {
		return obj.SPFRecord
	}

	if obj.SRVRecord != nil {
		return obj.SRVRecord
	}

	if obj.TLSARecord != nil {
		return obj.TLSARecord
	}

	if obj.TXTRecord != nil {
		return obj.TXTRecord
	}

	// all schemas are nil
	return nil
}

type NullableDnsRecord struct {
	value *DnsRecord
	isSet bool
}

func (v NullableDnsRecord) Get() *DnsRecord {
	return v.value
}

func (v *NullableDnsRecord) Set(val *DnsRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsRecord(val *DnsRecord) *NullableDnsRecord {
	return &NullableDnsRecord{value: val, isSet: true}
}

func (v NullableDnsRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


