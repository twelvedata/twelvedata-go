# GetTimeSeriesPercentB200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**PercentB** | **string** | Percent_b value | 

## Methods

### NewGetTimeSeriesPercentB200ResponseValuesInner

`func NewGetTimeSeriesPercentB200ResponseValuesInner(datetime string, percentB string, ) *GetTimeSeriesPercentB200ResponseValuesInner`

NewGetTimeSeriesPercentB200ResponseValuesInner instantiates a new GetTimeSeriesPercentB200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesPercentB200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesPercentB200ResponseValuesInnerWithDefaults() *GetTimeSeriesPercentB200ResponseValuesInner`

NewGetTimeSeriesPercentB200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesPercentB200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesPercentB200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesPercentB200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesPercentB200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetPercentB

`func (o *GetTimeSeriesPercentB200ResponseValuesInner) GetPercentB() string`

GetPercentB returns the PercentB field if non-nil, zero value otherwise.

### GetPercentBOk

`func (o *GetTimeSeriesPercentB200ResponseValuesInner) GetPercentBOk() (*string, bool)`

GetPercentBOk returns a tuple with the PercentB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentB

`func (o *GetTimeSeriesPercentB200ResponseValuesInner) SetPercentB(v string)`

SetPercentB sets PercentB field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


