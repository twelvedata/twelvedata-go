# GetTimeSeriesHtDcPhase200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**HtDcphase** | **string** | HT_DCPHASE value | 

## Methods

### NewGetTimeSeriesHtDcPhase200ResponseValuesInner

`func NewGetTimeSeriesHtDcPhase200ResponseValuesInner(datetime string, htDcphase string, ) *GetTimeSeriesHtDcPhase200ResponseValuesInner`

NewGetTimeSeriesHtDcPhase200ResponseValuesInner instantiates a new GetTimeSeriesHtDcPhase200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesHtDcPhase200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesHtDcPhase200ResponseValuesInnerWithDefaults() *GetTimeSeriesHtDcPhase200ResponseValuesInner`

NewGetTimeSeriesHtDcPhase200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesHtDcPhase200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesHtDcPhase200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesHtDcPhase200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesHtDcPhase200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetHtDcphase

`func (o *GetTimeSeriesHtDcPhase200ResponseValuesInner) GetHtDcphase() string`

GetHtDcphase returns the HtDcphase field if non-nil, zero value otherwise.

### GetHtDcphaseOk

`func (o *GetTimeSeriesHtDcPhase200ResponseValuesInner) GetHtDcphaseOk() (*string, bool)`

GetHtDcphaseOk returns a tuple with the HtDcphase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtDcphase

`func (o *GetTimeSeriesHtDcPhase200ResponseValuesInner) SetHtDcphase(v string)`

SetHtDcphase sets HtDcphase field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


