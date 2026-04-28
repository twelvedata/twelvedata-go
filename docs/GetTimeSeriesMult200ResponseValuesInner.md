# GetTimeSeriesMult200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Mult** | **string** | Mult value | 

## Methods

### NewGetTimeSeriesMult200ResponseValuesInner

`func NewGetTimeSeriesMult200ResponseValuesInner(datetime string, mult string, ) *GetTimeSeriesMult200ResponseValuesInner`

NewGetTimeSeriesMult200ResponseValuesInner instantiates a new GetTimeSeriesMult200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesMult200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesMult200ResponseValuesInnerWithDefaults() *GetTimeSeriesMult200ResponseValuesInner`

NewGetTimeSeriesMult200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesMult200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesMult200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesMult200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesMult200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetMult

`func (o *GetTimeSeriesMult200ResponseValuesInner) GetMult() string`

GetMult returns the Mult field if non-nil, zero value otherwise.

### GetMultOk

`func (o *GetTimeSeriesMult200ResponseValuesInner) GetMultOk() (*string, bool)`

GetMultOk returns a tuple with the Mult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMult

`func (o *GetTimeSeriesMult200ResponseValuesInner) SetMult(v string)`

SetMult sets Mult field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


