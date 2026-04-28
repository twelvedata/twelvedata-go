# GetTimeSeriesMinusDM200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**MinusDm** | **string** | Minus Directional Movement value | 

## Methods

### NewGetTimeSeriesMinusDM200ResponseValuesInner

`func NewGetTimeSeriesMinusDM200ResponseValuesInner(datetime string, minusDm string, ) *GetTimeSeriesMinusDM200ResponseValuesInner`

NewGetTimeSeriesMinusDM200ResponseValuesInner instantiates a new GetTimeSeriesMinusDM200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesMinusDM200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesMinusDM200ResponseValuesInnerWithDefaults() *GetTimeSeriesMinusDM200ResponseValuesInner`

NewGetTimeSeriesMinusDM200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesMinusDM200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesMinusDM200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesMinusDM200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesMinusDM200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetMinusDm

`func (o *GetTimeSeriesMinusDM200ResponseValuesInner) GetMinusDm() string`

GetMinusDm returns the MinusDm field if non-nil, zero value otherwise.

### GetMinusDmOk

`func (o *GetTimeSeriesMinusDM200ResponseValuesInner) GetMinusDmOk() (*string, bool)`

GetMinusDmOk returns a tuple with the MinusDm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinusDm

`func (o *GetTimeSeriesMinusDM200ResponseValuesInner) SetMinusDm(v string)`

SetMinusDm sets MinusDm field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


