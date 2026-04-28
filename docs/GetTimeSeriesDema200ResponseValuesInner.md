# GetTimeSeriesDema200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Dema** | **string** | Dema value | 

## Methods

### NewGetTimeSeriesDema200ResponseValuesInner

`func NewGetTimeSeriesDema200ResponseValuesInner(datetime string, dema string, ) *GetTimeSeriesDema200ResponseValuesInner`

NewGetTimeSeriesDema200ResponseValuesInner instantiates a new GetTimeSeriesDema200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesDema200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesDema200ResponseValuesInnerWithDefaults() *GetTimeSeriesDema200ResponseValuesInner`

NewGetTimeSeriesDema200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesDema200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesDema200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesDema200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesDema200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetDema

`func (o *GetTimeSeriesDema200ResponseValuesInner) GetDema() string`

GetDema returns the Dema field if non-nil, zero value otherwise.

### GetDemaOk

`func (o *GetTimeSeriesDema200ResponseValuesInner) GetDemaOk() (*string, bool)`

GetDemaOk returns a tuple with the Dema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDema

`func (o *GetTimeSeriesDema200ResponseValuesInner) SetDema(v string)`

SetDema sets Dema field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


