# GetTimeSeriesWma200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Wma** | **string** | WMA value | 

## Methods

### NewGetTimeSeriesWma200ResponseValuesInner

`func NewGetTimeSeriesWma200ResponseValuesInner(datetime string, wma string, ) *GetTimeSeriesWma200ResponseValuesInner`

NewGetTimeSeriesWma200ResponseValuesInner instantiates a new GetTimeSeriesWma200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesWma200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesWma200ResponseValuesInnerWithDefaults() *GetTimeSeriesWma200ResponseValuesInner`

NewGetTimeSeriesWma200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesWma200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesWma200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesWma200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesWma200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetWma

`func (o *GetTimeSeriesWma200ResponseValuesInner) GetWma() string`

GetWma returns the Wma field if non-nil, zero value otherwise.

### GetWmaOk

`func (o *GetTimeSeriesWma200ResponseValuesInner) GetWmaOk() (*string, bool)`

GetWmaOk returns a tuple with the Wma field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWma

`func (o *GetTimeSeriesWma200ResponseValuesInner) SetWma(v string)`

SetWma sets Wma field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


