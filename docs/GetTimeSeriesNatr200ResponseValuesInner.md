# GetTimeSeriesNatr200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Natr** | **string** | natr value | 

## Methods

### NewGetTimeSeriesNatr200ResponseValuesInner

`func NewGetTimeSeriesNatr200ResponseValuesInner(datetime string, natr string, ) *GetTimeSeriesNatr200ResponseValuesInner`

NewGetTimeSeriesNatr200ResponseValuesInner instantiates a new GetTimeSeriesNatr200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesNatr200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesNatr200ResponseValuesInnerWithDefaults() *GetTimeSeriesNatr200ResponseValuesInner`

NewGetTimeSeriesNatr200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesNatr200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesNatr200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesNatr200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesNatr200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetNatr

`func (o *GetTimeSeriesNatr200ResponseValuesInner) GetNatr() string`

GetNatr returns the Natr field if non-nil, zero value otherwise.

### GetNatrOk

`func (o *GetTimeSeriesNatr200ResponseValuesInner) GetNatrOk() (*string, bool)`

GetNatrOk returns a tuple with the Natr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatr

`func (o *GetTimeSeriesNatr200ResponseValuesInner) SetNatr(v string)`

SetNatr sets Natr field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


