# GetTimeSeriesHtTrendMode200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**HtTrendmode** | **string** | ht_trendmode value | 

## Methods

### NewGetTimeSeriesHtTrendMode200ResponseValuesInner

`func NewGetTimeSeriesHtTrendMode200ResponseValuesInner(datetime string, htTrendmode string, ) *GetTimeSeriesHtTrendMode200ResponseValuesInner`

NewGetTimeSeriesHtTrendMode200ResponseValuesInner instantiates a new GetTimeSeriesHtTrendMode200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesHtTrendMode200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesHtTrendMode200ResponseValuesInnerWithDefaults() *GetTimeSeriesHtTrendMode200ResponseValuesInner`

NewGetTimeSeriesHtTrendMode200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesHtTrendMode200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesHtTrendMode200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesHtTrendMode200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesHtTrendMode200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetHtTrendmode

`func (o *GetTimeSeriesHtTrendMode200ResponseValuesInner) GetHtTrendmode() string`

GetHtTrendmode returns the HtTrendmode field if non-nil, zero value otherwise.

### GetHtTrendmodeOk

`func (o *GetTimeSeriesHtTrendMode200ResponseValuesInner) GetHtTrendmodeOk() (*string, bool)`

GetHtTrendmodeOk returns a tuple with the HtTrendmode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtTrendmode

`func (o *GetTimeSeriesHtTrendMode200ResponseValuesInner) SetHtTrendmode(v string)`

SetHtTrendmode sets HtTrendmode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


