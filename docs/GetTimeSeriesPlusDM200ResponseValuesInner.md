# GetTimeSeriesPlusDM200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**PlusDm** | **string** | plus_dm value | 

## Methods

### NewGetTimeSeriesPlusDM200ResponseValuesInner

`func NewGetTimeSeriesPlusDM200ResponseValuesInner(datetime string, plusDm string, ) *GetTimeSeriesPlusDM200ResponseValuesInner`

NewGetTimeSeriesPlusDM200ResponseValuesInner instantiates a new GetTimeSeriesPlusDM200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesPlusDM200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesPlusDM200ResponseValuesInnerWithDefaults() *GetTimeSeriesPlusDM200ResponseValuesInner`

NewGetTimeSeriesPlusDM200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesPlusDM200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesPlusDM200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesPlusDM200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesPlusDM200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetPlusDm

`func (o *GetTimeSeriesPlusDM200ResponseValuesInner) GetPlusDm() string`

GetPlusDm returns the PlusDm field if non-nil, zero value otherwise.

### GetPlusDmOk

`func (o *GetTimeSeriesPlusDM200ResponseValuesInner) GetPlusDmOk() (*string, bool)`

GetPlusDmOk returns a tuple with the PlusDm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlusDm

`func (o *GetTimeSeriesPlusDM200ResponseValuesInner) SetPlusDm(v string)`

SetPlusDm sets PlusDm field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


