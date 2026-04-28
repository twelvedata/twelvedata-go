# GetTimeSeriesLinearRegAngle200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Linearregangle** | **string** | Linear regression angle value | 

## Methods

### NewGetTimeSeriesLinearRegAngle200ResponseValuesInner

`func NewGetTimeSeriesLinearRegAngle200ResponseValuesInner(datetime string, linearregangle string, ) *GetTimeSeriesLinearRegAngle200ResponseValuesInner`

NewGetTimeSeriesLinearRegAngle200ResponseValuesInner instantiates a new GetTimeSeriesLinearRegAngle200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesLinearRegAngle200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesLinearRegAngle200ResponseValuesInnerWithDefaults() *GetTimeSeriesLinearRegAngle200ResponseValuesInner`

NewGetTimeSeriesLinearRegAngle200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesLinearRegAngle200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesLinearRegAngle200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesLinearRegAngle200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesLinearRegAngle200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetLinearregangle

`func (o *GetTimeSeriesLinearRegAngle200ResponseValuesInner) GetLinearregangle() string`

GetLinearregangle returns the Linearregangle field if non-nil, zero value otherwise.

### GetLinearregangleOk

`func (o *GetTimeSeriesLinearRegAngle200ResponseValuesInner) GetLinearregangleOk() (*string, bool)`

GetLinearregangleOk returns a tuple with the Linearregangle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinearregangle

`func (o *GetTimeSeriesLinearRegAngle200ResponseValuesInner) SetLinearregangle(v string)`

SetLinearregangle sets Linearregangle field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


