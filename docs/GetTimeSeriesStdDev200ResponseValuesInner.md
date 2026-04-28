# GetTimeSeriesStdDev200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Stddev** | **string** | Standard Deviation value | 

## Methods

### NewGetTimeSeriesStdDev200ResponseValuesInner

`func NewGetTimeSeriesStdDev200ResponseValuesInner(datetime string, stddev string, ) *GetTimeSeriesStdDev200ResponseValuesInner`

NewGetTimeSeriesStdDev200ResponseValuesInner instantiates a new GetTimeSeriesStdDev200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesStdDev200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesStdDev200ResponseValuesInnerWithDefaults() *GetTimeSeriesStdDev200ResponseValuesInner`

NewGetTimeSeriesStdDev200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesStdDev200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesStdDev200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesStdDev200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesStdDev200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetStddev

`func (o *GetTimeSeriesStdDev200ResponseValuesInner) GetStddev() string`

GetStddev returns the Stddev field if non-nil, zero value otherwise.

### GetStddevOk

`func (o *GetTimeSeriesStdDev200ResponseValuesInner) GetStddevOk() (*string, bool)`

GetStddevOk returns a tuple with the Stddev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStddev

`func (o *GetTimeSeriesStdDev200ResponseValuesInner) SetStddev(v string)`

SetStddev sets Stddev field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


