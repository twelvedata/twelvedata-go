# GetTimeSeriesSqrt200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Sqrt** | **string** | SQRT value | 

## Methods

### NewGetTimeSeriesSqrt200ResponseValuesInner

`func NewGetTimeSeriesSqrt200ResponseValuesInner(datetime string, sqrt string, ) *GetTimeSeriesSqrt200ResponseValuesInner`

NewGetTimeSeriesSqrt200ResponseValuesInner instantiates a new GetTimeSeriesSqrt200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesSqrt200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesSqrt200ResponseValuesInnerWithDefaults() *GetTimeSeriesSqrt200ResponseValuesInner`

NewGetTimeSeriesSqrt200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesSqrt200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesSqrt200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesSqrt200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesSqrt200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetSqrt

`func (o *GetTimeSeriesSqrt200ResponseValuesInner) GetSqrt() string`

GetSqrt returns the Sqrt field if non-nil, zero value otherwise.

### GetSqrtOk

`func (o *GetTimeSeriesSqrt200ResponseValuesInner) GetSqrtOk() (*string, bool)`

GetSqrtOk returns a tuple with the Sqrt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSqrt

`func (o *GetTimeSeriesSqrt200ResponseValuesInner) SetSqrt(v string)`

SetSqrt sets Sqrt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


