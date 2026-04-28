# GetTimeSeriesExp200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Exp** | **string** | Exp value | 

## Methods

### NewGetTimeSeriesExp200ResponseValuesInner

`func NewGetTimeSeriesExp200ResponseValuesInner(datetime string, exp string, ) *GetTimeSeriesExp200ResponseValuesInner`

NewGetTimeSeriesExp200ResponseValuesInner instantiates a new GetTimeSeriesExp200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesExp200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesExp200ResponseValuesInnerWithDefaults() *GetTimeSeriesExp200ResponseValuesInner`

NewGetTimeSeriesExp200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesExp200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesExp200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesExp200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesExp200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetExp

`func (o *GetTimeSeriesExp200ResponseValuesInner) GetExp() string`

GetExp returns the Exp field if non-nil, zero value otherwise.

### GetExpOk

`func (o *GetTimeSeriesExp200ResponseValuesInner) GetExpOk() (*string, bool)`

GetExpOk returns a tuple with the Exp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExp

`func (o *GetTimeSeriesExp200ResponseValuesInner) SetExp(v string)`

SetExp sets Exp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


