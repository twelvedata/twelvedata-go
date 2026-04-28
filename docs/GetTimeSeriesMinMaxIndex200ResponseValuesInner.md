# GetTimeSeriesMinMaxIndex200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Minidx** | **string** | Index of the lowest value over the specified period | 
**Maxidx** | **string** | Index of the highest value over the specified period | 

## Methods

### NewGetTimeSeriesMinMaxIndex200ResponseValuesInner

`func NewGetTimeSeriesMinMaxIndex200ResponseValuesInner(datetime string, minidx string, maxidx string, ) *GetTimeSeriesMinMaxIndex200ResponseValuesInner`

NewGetTimeSeriesMinMaxIndex200ResponseValuesInner instantiates a new GetTimeSeriesMinMaxIndex200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesMinMaxIndex200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesMinMaxIndex200ResponseValuesInnerWithDefaults() *GetTimeSeriesMinMaxIndex200ResponseValuesInner`

NewGetTimeSeriesMinMaxIndex200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesMinMaxIndex200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesMinMaxIndex200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesMinMaxIndex200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesMinMaxIndex200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetMinidx

`func (o *GetTimeSeriesMinMaxIndex200ResponseValuesInner) GetMinidx() string`

GetMinidx returns the Minidx field if non-nil, zero value otherwise.

### GetMinidxOk

`func (o *GetTimeSeriesMinMaxIndex200ResponseValuesInner) GetMinidxOk() (*string, bool)`

GetMinidxOk returns a tuple with the Minidx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinidx

`func (o *GetTimeSeriesMinMaxIndex200ResponseValuesInner) SetMinidx(v string)`

SetMinidx sets Minidx field to given value.


### GetMaxidx

`func (o *GetTimeSeriesMinMaxIndex200ResponseValuesInner) GetMaxidx() string`

GetMaxidx returns the Maxidx field if non-nil, zero value otherwise.

### GetMaxidxOk

`func (o *GetTimeSeriesMinMaxIndex200ResponseValuesInner) GetMaxidxOk() (*string, bool)`

GetMaxidxOk returns a tuple with the Maxidx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxidx

`func (o *GetTimeSeriesMinMaxIndex200ResponseValuesInner) SetMaxidx(v string)`

SetMaxidx sets Maxidx field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


