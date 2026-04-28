# GetTimeSeriesEma200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Ema** | **string** | EMA value | 

## Methods

### NewGetTimeSeriesEma200ResponseValuesInner

`func NewGetTimeSeriesEma200ResponseValuesInner(datetime string, ema string, ) *GetTimeSeriesEma200ResponseValuesInner`

NewGetTimeSeriesEma200ResponseValuesInner instantiates a new GetTimeSeriesEma200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesEma200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesEma200ResponseValuesInnerWithDefaults() *GetTimeSeriesEma200ResponseValuesInner`

NewGetTimeSeriesEma200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesEma200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesEma200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesEma200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesEma200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetEma

`func (o *GetTimeSeriesEma200ResponseValuesInner) GetEma() string`

GetEma returns the Ema field if non-nil, zero value otherwise.

### GetEmaOk

`func (o *GetTimeSeriesEma200ResponseValuesInner) GetEmaOk() (*string, bool)`

GetEmaOk returns a tuple with the Ema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEma

`func (o *GetTimeSeriesEma200ResponseValuesInner) SetEma(v string)`

SetEma sets Ema field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


