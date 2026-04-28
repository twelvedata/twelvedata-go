# GetTimeSeriesPlusDI200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**PlusDi** | **string** | plus_di value | 

## Methods

### NewGetTimeSeriesPlusDI200ResponseValuesInner

`func NewGetTimeSeriesPlusDI200ResponseValuesInner(datetime string, plusDi string, ) *GetTimeSeriesPlusDI200ResponseValuesInner`

NewGetTimeSeriesPlusDI200ResponseValuesInner instantiates a new GetTimeSeriesPlusDI200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesPlusDI200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesPlusDI200ResponseValuesInnerWithDefaults() *GetTimeSeriesPlusDI200ResponseValuesInner`

NewGetTimeSeriesPlusDI200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesPlusDI200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesPlusDI200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesPlusDI200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesPlusDI200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetPlusDi

`func (o *GetTimeSeriesPlusDI200ResponseValuesInner) GetPlusDi() string`

GetPlusDi returns the PlusDi field if non-nil, zero value otherwise.

### GetPlusDiOk

`func (o *GetTimeSeriesPlusDI200ResponseValuesInner) GetPlusDiOk() (*string, bool)`

GetPlusDiOk returns a tuple with the PlusDi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlusDi

`func (o *GetTimeSeriesPlusDI200ResponseValuesInner) SetPlusDi(v string)`

SetPlusDi sets PlusDi field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


