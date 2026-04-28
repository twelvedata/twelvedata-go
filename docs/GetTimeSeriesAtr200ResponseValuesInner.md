# GetTimeSeriesAtr200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Atr** | **string** | ATR value | 

## Methods

### NewGetTimeSeriesAtr200ResponseValuesInner

`func NewGetTimeSeriesAtr200ResponseValuesInner(datetime string, atr string, ) *GetTimeSeriesAtr200ResponseValuesInner`

NewGetTimeSeriesAtr200ResponseValuesInner instantiates a new GetTimeSeriesAtr200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesAtr200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesAtr200ResponseValuesInnerWithDefaults() *GetTimeSeriesAtr200ResponseValuesInner`

NewGetTimeSeriesAtr200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesAtr200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesAtr200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesAtr200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesAtr200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetAtr

`func (o *GetTimeSeriesAtr200ResponseValuesInner) GetAtr() string`

GetAtr returns the Atr field if non-nil, zero value otherwise.

### GetAtrOk

`func (o *GetTimeSeriesAtr200ResponseValuesInner) GetAtrOk() (*string, bool)`

GetAtrOk returns a tuple with the Atr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtr

`func (o *GetTimeSeriesAtr200ResponseValuesInner) SetAtr(v string)`

SetAtr sets Atr field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


