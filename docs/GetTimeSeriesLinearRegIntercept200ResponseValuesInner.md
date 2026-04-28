# GetTimeSeriesLinearRegIntercept200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Linearregintercept** | **string** | Linear Regression Intercept value | 

## Methods

### NewGetTimeSeriesLinearRegIntercept200ResponseValuesInner

`func NewGetTimeSeriesLinearRegIntercept200ResponseValuesInner(datetime string, linearregintercept string, ) *GetTimeSeriesLinearRegIntercept200ResponseValuesInner`

NewGetTimeSeriesLinearRegIntercept200ResponseValuesInner instantiates a new GetTimeSeriesLinearRegIntercept200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesLinearRegIntercept200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesLinearRegIntercept200ResponseValuesInnerWithDefaults() *GetTimeSeriesLinearRegIntercept200ResponseValuesInner`

NewGetTimeSeriesLinearRegIntercept200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesLinearRegIntercept200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesLinearRegIntercept200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesLinearRegIntercept200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesLinearRegIntercept200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetLinearregintercept

`func (o *GetTimeSeriesLinearRegIntercept200ResponseValuesInner) GetLinearregintercept() string`

GetLinearregintercept returns the Linearregintercept field if non-nil, zero value otherwise.

### GetLinearreginterceptOk

`func (o *GetTimeSeriesLinearRegIntercept200ResponseValuesInner) GetLinearreginterceptOk() (*string, bool)`

GetLinearreginterceptOk returns a tuple with the Linearregintercept field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinearregintercept

`func (o *GetTimeSeriesLinearRegIntercept200ResponseValuesInner) SetLinearregintercept(v string)`

SetLinearregintercept sets Linearregintercept field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


