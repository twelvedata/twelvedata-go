# GetTimeSeriesMacdSlope200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**MacdSlope** | **string** | MACD slope value | 
**MacdSignalSlope** | **string** | MACD signal slope value | 
**MacdHistSlope** | **string** | MACD histogram slope value | 

## Methods

### NewGetTimeSeriesMacdSlope200ResponseValuesInner

`func NewGetTimeSeriesMacdSlope200ResponseValuesInner(datetime string, macdSlope string, macdSignalSlope string, macdHistSlope string, ) *GetTimeSeriesMacdSlope200ResponseValuesInner`

NewGetTimeSeriesMacdSlope200ResponseValuesInner instantiates a new GetTimeSeriesMacdSlope200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesMacdSlope200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesMacdSlope200ResponseValuesInnerWithDefaults() *GetTimeSeriesMacdSlope200ResponseValuesInner`

NewGetTimeSeriesMacdSlope200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesMacdSlope200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesMacdSlope200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesMacdSlope200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesMacdSlope200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetMacdSlope

`func (o *GetTimeSeriesMacdSlope200ResponseValuesInner) GetMacdSlope() string`

GetMacdSlope returns the MacdSlope field if non-nil, zero value otherwise.

### GetMacdSlopeOk

`func (o *GetTimeSeriesMacdSlope200ResponseValuesInner) GetMacdSlopeOk() (*string, bool)`

GetMacdSlopeOk returns a tuple with the MacdSlope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacdSlope

`func (o *GetTimeSeriesMacdSlope200ResponseValuesInner) SetMacdSlope(v string)`

SetMacdSlope sets MacdSlope field to given value.


### GetMacdSignalSlope

`func (o *GetTimeSeriesMacdSlope200ResponseValuesInner) GetMacdSignalSlope() string`

GetMacdSignalSlope returns the MacdSignalSlope field if non-nil, zero value otherwise.

### GetMacdSignalSlopeOk

`func (o *GetTimeSeriesMacdSlope200ResponseValuesInner) GetMacdSignalSlopeOk() (*string, bool)`

GetMacdSignalSlopeOk returns a tuple with the MacdSignalSlope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacdSignalSlope

`func (o *GetTimeSeriesMacdSlope200ResponseValuesInner) SetMacdSignalSlope(v string)`

SetMacdSignalSlope sets MacdSignalSlope field to given value.


### GetMacdHistSlope

`func (o *GetTimeSeriesMacdSlope200ResponseValuesInner) GetMacdHistSlope() string`

GetMacdHistSlope returns the MacdHistSlope field if non-nil, zero value otherwise.

### GetMacdHistSlopeOk

`func (o *GetTimeSeriesMacdSlope200ResponseValuesInner) GetMacdHistSlopeOk() (*string, bool)`

GetMacdHistSlopeOk returns a tuple with the MacdHistSlope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacdHistSlope

`func (o *GetTimeSeriesMacdSlope200ResponseValuesInner) SetMacdHistSlope(v string)`

SetMacdHistSlope sets MacdHistSlope field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


