# GetTimeSeriesDpo200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Dpo** | **string** | DPO value | 

## Methods

### NewGetTimeSeriesDpo200ResponseValuesInner

`func NewGetTimeSeriesDpo200ResponseValuesInner(datetime string, dpo string, ) *GetTimeSeriesDpo200ResponseValuesInner`

NewGetTimeSeriesDpo200ResponseValuesInner instantiates a new GetTimeSeriesDpo200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesDpo200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesDpo200ResponseValuesInnerWithDefaults() *GetTimeSeriesDpo200ResponseValuesInner`

NewGetTimeSeriesDpo200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesDpo200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesDpo200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesDpo200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesDpo200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetDpo

`func (o *GetTimeSeriesDpo200ResponseValuesInner) GetDpo() string`

GetDpo returns the Dpo field if non-nil, zero value otherwise.

### GetDpoOk

`func (o *GetTimeSeriesDpo200ResponseValuesInner) GetDpoOk() (*string, bool)`

GetDpoOk returns a tuple with the Dpo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpo

`func (o *GetTimeSeriesDpo200ResponseValuesInner) SetDpo(v string)`

SetDpo sets Dpo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


