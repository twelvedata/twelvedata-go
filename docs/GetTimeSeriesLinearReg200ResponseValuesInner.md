# GetTimeSeriesLinearReg200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Linearreg** | **string** | linearreg value | 

## Methods

### NewGetTimeSeriesLinearReg200ResponseValuesInner

`func NewGetTimeSeriesLinearReg200ResponseValuesInner(datetime string, linearreg string, ) *GetTimeSeriesLinearReg200ResponseValuesInner`

NewGetTimeSeriesLinearReg200ResponseValuesInner instantiates a new GetTimeSeriesLinearReg200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesLinearReg200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesLinearReg200ResponseValuesInnerWithDefaults() *GetTimeSeriesLinearReg200ResponseValuesInner`

NewGetTimeSeriesLinearReg200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesLinearReg200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesLinearReg200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesLinearReg200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesLinearReg200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetLinearreg

`func (o *GetTimeSeriesLinearReg200ResponseValuesInner) GetLinearreg() string`

GetLinearreg returns the Linearreg field if non-nil, zero value otherwise.

### GetLinearregOk

`func (o *GetTimeSeriesLinearReg200ResponseValuesInner) GetLinearregOk() (*string, bool)`

GetLinearregOk returns a tuple with the Linearreg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinearreg

`func (o *GetTimeSeriesLinearReg200ResponseValuesInner) SetLinearreg(v string)`

SetLinearreg sets Linearreg field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


