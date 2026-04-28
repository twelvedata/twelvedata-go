# GetTimeSeriesCorrel200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Correl** | **string** | Correl value | 

## Methods

### NewGetTimeSeriesCorrel200ResponseValuesInner

`func NewGetTimeSeriesCorrel200ResponseValuesInner(datetime string, correl string, ) *GetTimeSeriesCorrel200ResponseValuesInner`

NewGetTimeSeriesCorrel200ResponseValuesInner instantiates a new GetTimeSeriesCorrel200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesCorrel200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesCorrel200ResponseValuesInnerWithDefaults() *GetTimeSeriesCorrel200ResponseValuesInner`

NewGetTimeSeriesCorrel200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesCorrel200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesCorrel200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesCorrel200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesCorrel200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetCorrel

`func (o *GetTimeSeriesCorrel200ResponseValuesInner) GetCorrel() string`

GetCorrel returns the Correl field if non-nil, zero value otherwise.

### GetCorrelOk

`func (o *GetTimeSeriesCorrel200ResponseValuesInner) GetCorrelOk() (*string, bool)`

GetCorrelOk returns a tuple with the Correl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrel

`func (o *GetTimeSeriesCorrel200ResponseValuesInner) SetCorrel(v string)`

SetCorrel sets Correl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


