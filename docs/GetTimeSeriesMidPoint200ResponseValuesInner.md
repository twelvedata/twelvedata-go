# GetTimeSeriesMidPoint200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Midpoint** | **string** | Midpoint value | 

## Methods

### NewGetTimeSeriesMidPoint200ResponseValuesInner

`func NewGetTimeSeriesMidPoint200ResponseValuesInner(datetime string, midpoint string, ) *GetTimeSeriesMidPoint200ResponseValuesInner`

NewGetTimeSeriesMidPoint200ResponseValuesInner instantiates a new GetTimeSeriesMidPoint200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesMidPoint200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesMidPoint200ResponseValuesInnerWithDefaults() *GetTimeSeriesMidPoint200ResponseValuesInner`

NewGetTimeSeriesMidPoint200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesMidPoint200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesMidPoint200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesMidPoint200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesMidPoint200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetMidpoint

`func (o *GetTimeSeriesMidPoint200ResponseValuesInner) GetMidpoint() string`

GetMidpoint returns the Midpoint field if non-nil, zero value otherwise.

### GetMidpointOk

`func (o *GetTimeSeriesMidPoint200ResponseValuesInner) GetMidpointOk() (*string, bool)`

GetMidpointOk returns a tuple with the Midpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMidpoint

`func (o *GetTimeSeriesMidPoint200ResponseValuesInner) SetMidpoint(v string)`

SetMidpoint sets Midpoint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


