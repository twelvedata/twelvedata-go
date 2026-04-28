# GetTimeSeriesCrsi200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Crsi** | **string** | crsi value | 

## Methods

### NewGetTimeSeriesCrsi200ResponseValuesInner

`func NewGetTimeSeriesCrsi200ResponseValuesInner(datetime string, crsi string, ) *GetTimeSeriesCrsi200ResponseValuesInner`

NewGetTimeSeriesCrsi200ResponseValuesInner instantiates a new GetTimeSeriesCrsi200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesCrsi200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesCrsi200ResponseValuesInnerWithDefaults() *GetTimeSeriesCrsi200ResponseValuesInner`

NewGetTimeSeriesCrsi200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesCrsi200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesCrsi200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesCrsi200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesCrsi200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetCrsi

`func (o *GetTimeSeriesCrsi200ResponseValuesInner) GetCrsi() string`

GetCrsi returns the Crsi field if non-nil, zero value otherwise.

### GetCrsiOk

`func (o *GetTimeSeriesCrsi200ResponseValuesInner) GetCrsiOk() (*string, bool)`

GetCrsiOk returns a tuple with the Crsi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrsi

`func (o *GetTimeSeriesCrsi200ResponseValuesInner) SetCrsi(v string)`

SetCrsi sets Crsi field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


