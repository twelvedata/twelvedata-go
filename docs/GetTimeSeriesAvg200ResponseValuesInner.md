# GetTimeSeriesAvg200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Avg** | **string** | Avg value | 

## Methods

### NewGetTimeSeriesAvg200ResponseValuesInner

`func NewGetTimeSeriesAvg200ResponseValuesInner(datetime string, avg string, ) *GetTimeSeriesAvg200ResponseValuesInner`

NewGetTimeSeriesAvg200ResponseValuesInner instantiates a new GetTimeSeriesAvg200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesAvg200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesAvg200ResponseValuesInnerWithDefaults() *GetTimeSeriesAvg200ResponseValuesInner`

NewGetTimeSeriesAvg200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesAvg200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesAvg200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesAvg200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesAvg200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetAvg

`func (o *GetTimeSeriesAvg200ResponseValuesInner) GetAvg() string`

GetAvg returns the Avg field if non-nil, zero value otherwise.

### GetAvgOk

`func (o *GetTimeSeriesAvg200ResponseValuesInner) GetAvgOk() (*string, bool)`

GetAvgOk returns a tuple with the Avg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvg

`func (o *GetTimeSeriesAvg200ResponseValuesInner) SetAvg(v string)`

SetAvg sets Avg field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


