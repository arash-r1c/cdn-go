# BulkTrafficReportData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requests** | Pointer to **int64** |  | [optional] 
**IngressBytes** | Pointer to **int64** |  | [optional] 
**EgressBytes** | Pointer to [**BulkTrafficReportDataEgressBytes**](BulkTrafficReportDataEgressBytes.md) |  | [optional] 

## Methods

### NewBulkTrafficReportData

`func NewBulkTrafficReportData() *BulkTrafficReportData`

NewBulkTrafficReportData instantiates a new BulkTrafficReportData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkTrafficReportDataWithDefaults

`func NewBulkTrafficReportDataWithDefaults() *BulkTrafficReportData`

NewBulkTrafficReportDataWithDefaults instantiates a new BulkTrafficReportData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequests

`func (o *BulkTrafficReportData) GetRequests() int64`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *BulkTrafficReportData) GetRequestsOk() (*int64, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *BulkTrafficReportData) SetRequests(v int64)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *BulkTrafficReportData) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetIngressBytes

`func (o *BulkTrafficReportData) GetIngressBytes() int64`

GetIngressBytes returns the IngressBytes field if non-nil, zero value otherwise.

### GetIngressBytesOk

`func (o *BulkTrafficReportData) GetIngressBytesOk() (*int64, bool)`

GetIngressBytesOk returns a tuple with the IngressBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressBytes

`func (o *BulkTrafficReportData) SetIngressBytes(v int64)`

SetIngressBytes sets IngressBytes field to given value.

### HasIngressBytes

`func (o *BulkTrafficReportData) HasIngressBytes() bool`

HasIngressBytes returns a boolean if a field has been set.

### GetEgressBytes

`func (o *BulkTrafficReportData) GetEgressBytes() BulkTrafficReportDataEgressBytes`

GetEgressBytes returns the EgressBytes field if non-nil, zero value otherwise.

### GetEgressBytesOk

`func (o *BulkTrafficReportData) GetEgressBytesOk() (*BulkTrafficReportDataEgressBytes, bool)`

GetEgressBytesOk returns a tuple with the EgressBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressBytes

`func (o *BulkTrafficReportData) SetEgressBytes(v BulkTrafficReportDataEgressBytes)`

SetEgressBytes sets EgressBytes field to given value.

### HasEgressBytes

`func (o *BulkTrafficReportData) HasEgressBytes() bool`

HasEgressBytes returns a boolean if a field has been set.


[[Back to Model list]](HOW-TO.md#documentation-for-models) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints) [[Back to README]](HOW-TO.md)


