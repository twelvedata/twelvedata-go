# GetTimeSeriesPpo200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Ppo** | **string** | PPO value | 

## Methods

### NewGetTimeSeriesPpo200ResponseValuesInner

`func NewGetTimeSeriesPpo200ResponseValuesInner(datetime string, ppo string, ) *GetTimeSeriesPpo200ResponseValuesInner`

NewGetTimeSeriesPpo200ResponseValuesInner instantiates a new GetTimeSeriesPpo200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesPpo200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesPpo200ResponseValuesInnerWithDefaults() *GetTimeSeriesPpo200ResponseValuesInner`

NewGetTimeSeriesPpo200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesPpo200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesPpo200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesPpo200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesPpo200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetPpo

`func (o *GetTimeSeriesPpo200ResponseValuesInner) GetPpo() string`

GetPpo returns the Ppo field if non-nil, zero value otherwise.

### GetPpoOk

`func (o *GetTimeSeriesPpo200ResponseValuesInner) GetPpoOk() (*string, bool)`

GetPpoOk returns a tuple with the Ppo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpo

`func (o *GetTimeSeriesPpo200ResponseValuesInner) SetPpo(v string)`

SetPpo sets Ppo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


