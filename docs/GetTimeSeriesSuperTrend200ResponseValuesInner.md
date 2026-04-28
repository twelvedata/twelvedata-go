# GetTimeSeriesSuperTrend200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Supertrend** | **string** | SuperTrend value | 

## Methods

### NewGetTimeSeriesSuperTrend200ResponseValuesInner

`func NewGetTimeSeriesSuperTrend200ResponseValuesInner(datetime string, supertrend string, ) *GetTimeSeriesSuperTrend200ResponseValuesInner`

NewGetTimeSeriesSuperTrend200ResponseValuesInner instantiates a new GetTimeSeriesSuperTrend200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesSuperTrend200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesSuperTrend200ResponseValuesInnerWithDefaults() *GetTimeSeriesSuperTrend200ResponseValuesInner`

NewGetTimeSeriesSuperTrend200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesSuperTrend200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesSuperTrend200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesSuperTrend200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesSuperTrend200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetSupertrend

`func (o *GetTimeSeriesSuperTrend200ResponseValuesInner) GetSupertrend() string`

GetSupertrend returns the Supertrend field if non-nil, zero value otherwise.

### GetSupertrendOk

`func (o *GetTimeSeriesSuperTrend200ResponseValuesInner) GetSupertrendOk() (*string, bool)`

GetSupertrendOk returns a tuple with the Supertrend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupertrend

`func (o *GetTimeSeriesSuperTrend200ResponseValuesInner) SetSupertrend(v string)`

SetSupertrend sets Supertrend field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


