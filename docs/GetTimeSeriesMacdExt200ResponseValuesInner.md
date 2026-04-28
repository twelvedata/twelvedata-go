# GetTimeSeriesMacdExt200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Macd** | **string** | MACD value | 
**MacdSignal** | **string** | MACD signal line value | 
**MacdHist** | **string** | MACD histogram value | 

## Methods

### NewGetTimeSeriesMacdExt200ResponseValuesInner

`func NewGetTimeSeriesMacdExt200ResponseValuesInner(datetime string, macd string, macdSignal string, macdHist string, ) *GetTimeSeriesMacdExt200ResponseValuesInner`

NewGetTimeSeriesMacdExt200ResponseValuesInner instantiates a new GetTimeSeriesMacdExt200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesMacdExt200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesMacdExt200ResponseValuesInnerWithDefaults() *GetTimeSeriesMacdExt200ResponseValuesInner`

NewGetTimeSeriesMacdExt200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesMacdExt200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesMacdExt200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesMacdExt200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesMacdExt200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetMacd

`func (o *GetTimeSeriesMacdExt200ResponseValuesInner) GetMacd() string`

GetMacd returns the Macd field if non-nil, zero value otherwise.

### GetMacdOk

`func (o *GetTimeSeriesMacdExt200ResponseValuesInner) GetMacdOk() (*string, bool)`

GetMacdOk returns a tuple with the Macd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacd

`func (o *GetTimeSeriesMacdExt200ResponseValuesInner) SetMacd(v string)`

SetMacd sets Macd field to given value.


### GetMacdSignal

`func (o *GetTimeSeriesMacdExt200ResponseValuesInner) GetMacdSignal() string`

GetMacdSignal returns the MacdSignal field if non-nil, zero value otherwise.

### GetMacdSignalOk

`func (o *GetTimeSeriesMacdExt200ResponseValuesInner) GetMacdSignalOk() (*string, bool)`

GetMacdSignalOk returns a tuple with the MacdSignal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacdSignal

`func (o *GetTimeSeriesMacdExt200ResponseValuesInner) SetMacdSignal(v string)`

SetMacdSignal sets MacdSignal field to given value.


### GetMacdHist

`func (o *GetTimeSeriesMacdExt200ResponseValuesInner) GetMacdHist() string`

GetMacdHist returns the MacdHist field if non-nil, zero value otherwise.

### GetMacdHistOk

`func (o *GetTimeSeriesMacdExt200ResponseValuesInner) GetMacdHistOk() (*string, bool)`

GetMacdHistOk returns a tuple with the MacdHist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacdHist

`func (o *GetTimeSeriesMacdExt200ResponseValuesInner) SetMacdHist(v string)`

SetMacdHist sets MacdHist field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


