# GetTimeSeriesRsi200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Rsi** | **string** | RSI value | 

## Methods

### NewGetTimeSeriesRsi200ResponseValuesInner

`func NewGetTimeSeriesRsi200ResponseValuesInner(datetime string, rsi string, ) *GetTimeSeriesRsi200ResponseValuesInner`

NewGetTimeSeriesRsi200ResponseValuesInner instantiates a new GetTimeSeriesRsi200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesRsi200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesRsi200ResponseValuesInnerWithDefaults() *GetTimeSeriesRsi200ResponseValuesInner`

NewGetTimeSeriesRsi200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesRsi200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesRsi200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesRsi200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesRsi200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetRsi

`func (o *GetTimeSeriesRsi200ResponseValuesInner) GetRsi() string`

GetRsi returns the Rsi field if non-nil, zero value otherwise.

### GetRsiOk

`func (o *GetTimeSeriesRsi200ResponseValuesInner) GetRsiOk() (*string, bool)`

GetRsiOk returns a tuple with the Rsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsi

`func (o *GetTimeSeriesRsi200ResponseValuesInner) SetRsi(v string)`

SetRsi sets Rsi field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


