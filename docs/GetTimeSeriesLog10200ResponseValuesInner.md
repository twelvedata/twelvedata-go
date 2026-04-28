# GetTimeSeriesLog10200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Log10** | **string** | Log10 value | 

## Methods

### NewGetTimeSeriesLog10200ResponseValuesInner

`func NewGetTimeSeriesLog10200ResponseValuesInner(datetime string, log10 string, ) *GetTimeSeriesLog10200ResponseValuesInner`

NewGetTimeSeriesLog10200ResponseValuesInner instantiates a new GetTimeSeriesLog10200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesLog10200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesLog10200ResponseValuesInnerWithDefaults() *GetTimeSeriesLog10200ResponseValuesInner`

NewGetTimeSeriesLog10200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesLog10200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesLog10200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesLog10200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesLog10200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetLog10

`func (o *GetTimeSeriesLog10200ResponseValuesInner) GetLog10() string`

GetLog10 returns the Log10 field if non-nil, zero value otherwise.

### GetLog10Ok

`func (o *GetTimeSeriesLog10200ResponseValuesInner) GetLog10Ok() (*string, bool)`

GetLog10Ok returns a tuple with the Log10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLog10

`func (o *GetTimeSeriesLog10200ResponseValuesInner) SetLog10(v string)`

SetLog10 sets Log10 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


