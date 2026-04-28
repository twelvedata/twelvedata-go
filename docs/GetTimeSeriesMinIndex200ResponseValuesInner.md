# GetTimeSeriesMinIndex200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Minidx** | **string** | Index of lowest value over period | 

## Methods

### NewGetTimeSeriesMinIndex200ResponseValuesInner

`func NewGetTimeSeriesMinIndex200ResponseValuesInner(datetime string, minidx string, ) *GetTimeSeriesMinIndex200ResponseValuesInner`

NewGetTimeSeriesMinIndex200ResponseValuesInner instantiates a new GetTimeSeriesMinIndex200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesMinIndex200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesMinIndex200ResponseValuesInnerWithDefaults() *GetTimeSeriesMinIndex200ResponseValuesInner`

NewGetTimeSeriesMinIndex200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesMinIndex200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesMinIndex200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesMinIndex200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesMinIndex200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetMinidx

`func (o *GetTimeSeriesMinIndex200ResponseValuesInner) GetMinidx() string`

GetMinidx returns the Minidx field if non-nil, zero value otherwise.

### GetMinidxOk

`func (o *GetTimeSeriesMinIndex200ResponseValuesInner) GetMinidxOk() (*string, bool)`

GetMinidxOk returns a tuple with the Minidx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinidx

`func (o *GetTimeSeriesMinIndex200ResponseValuesInner) SetMinidx(v string)`

SetMinidx sets Minidx field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


