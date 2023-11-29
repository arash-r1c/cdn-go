# TrafficStatisticsTraffics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Saved** | **int64** |  | 
**Bypass** | Pointer to **int64** |  | [optional] 
**Top** | **time.Time** |  | 
**Total** | **int64** |  | 

## Methods

### NewTrafficStatisticsTraffics

`func NewTrafficStatisticsTraffics(saved int64, top time.Time, total int64, ) *TrafficStatisticsTraffics`

NewTrafficStatisticsTraffics instantiates a new TrafficStatisticsTraffics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficStatisticsTrafficsWithDefaults

`func NewTrafficStatisticsTrafficsWithDefaults() *TrafficStatisticsTraffics`

NewTrafficStatisticsTrafficsWithDefaults instantiates a new TrafficStatisticsTraffics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSaved

`func (o *TrafficStatisticsTraffics) GetSaved() int64`

GetSaved returns the Saved field if non-nil, zero value otherwise.

### GetSavedOk

`func (o *TrafficStatisticsTraffics) GetSavedOk() (*int64, bool)`

GetSavedOk returns a tuple with the Saved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaved

`func (o *TrafficStatisticsTraffics) SetSaved(v int64)`

SetSaved sets Saved field to given value.


### GetBypass

`func (o *TrafficStatisticsTraffics) GetBypass() int64`

GetBypass returns the Bypass field if non-nil, zero value otherwise.

### GetBypassOk

`func (o *TrafficStatisticsTraffics) GetBypassOk() (*int64, bool)`

GetBypassOk returns a tuple with the Bypass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypass

`func (o *TrafficStatisticsTraffics) SetBypass(v int64)`

SetBypass sets Bypass field to given value.

### HasBypass

`func (o *TrafficStatisticsTraffics) HasBypass() bool`

HasBypass returns a boolean if a field has been set.

### GetTop

`func (o *TrafficStatisticsTraffics) GetTop() time.Time`

GetTop returns the Top field if non-nil, zero value otherwise.

### GetTopOk

`func (o *TrafficStatisticsTraffics) GetTopOk() (*time.Time, bool)`

GetTopOk returns a tuple with the Top field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTop

`func (o *TrafficStatisticsTraffics) SetTop(v time.Time)`

SetTop sets Top field to given value.


### GetTotal

`func (o *TrafficStatisticsTraffics) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *TrafficStatisticsTraffics) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *TrafficStatisticsTraffics) SetTotal(v int64)`

SetTotal sets Total field to given value.



[[Back to Model list]](HOW-TO.md#documentation-for-models) [[Back to API list]](HOW-TO.md#documentation-for-api-endpoints) [[Back to README]](HOW-TO.md)


