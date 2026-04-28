# GetTimeSeriesSub200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Sub** | **string** | SUB value | 

## Methods

### NewGetTimeSeriesSub200ResponseValuesInner

`func NewGetTimeSeriesSub200ResponseValuesInner(datetime string, sub string, ) *GetTimeSeriesSub200ResponseValuesInner`

NewGetTimeSeriesSub200ResponseValuesInner instantiates a new GetTimeSeriesSub200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesSub200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesSub200ResponseValuesInnerWithDefaults() *GetTimeSeriesSub200ResponseValuesInner`

NewGetTimeSeriesSub200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesSub200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesSub200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesSub200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesSub200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetSub

`func (o *GetTimeSeriesSub200ResponseValuesInner) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *GetTimeSeriesSub200ResponseValuesInner) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *GetTimeSeriesSub200ResponseValuesInner) SetSub(v string)`

SetSub sets Sub field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


