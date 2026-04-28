# GetTimeSeriesHlc3200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Hlc3** | **string** | hlc3 value | 

## Methods

### NewGetTimeSeriesHlc3200ResponseValuesInner

`func NewGetTimeSeriesHlc3200ResponseValuesInner(datetime string, hlc3 string, ) *GetTimeSeriesHlc3200ResponseValuesInner`

NewGetTimeSeriesHlc3200ResponseValuesInner instantiates a new GetTimeSeriesHlc3200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesHlc3200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesHlc3200ResponseValuesInnerWithDefaults() *GetTimeSeriesHlc3200ResponseValuesInner`

NewGetTimeSeriesHlc3200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesHlc3200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesHlc3200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesHlc3200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesHlc3200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetHlc3

`func (o *GetTimeSeriesHlc3200ResponseValuesInner) GetHlc3() string`

GetHlc3 returns the Hlc3 field if non-nil, zero value otherwise.

### GetHlc3Ok

`func (o *GetTimeSeriesHlc3200ResponseValuesInner) GetHlc3Ok() (*string, bool)`

GetHlc3Ok returns a tuple with the Hlc3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHlc3

`func (o *GetTimeSeriesHlc3200ResponseValuesInner) SetHlc3(v string)`

SetHlc3 sets Hlc3 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


