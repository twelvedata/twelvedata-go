# GetTimeSeriesSum200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Sum** | **string** | Sum value | 

## Methods

### NewGetTimeSeriesSum200ResponseValuesInner

`func NewGetTimeSeriesSum200ResponseValuesInner(datetime string, sum string, ) *GetTimeSeriesSum200ResponseValuesInner`

NewGetTimeSeriesSum200ResponseValuesInner instantiates a new GetTimeSeriesSum200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesSum200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesSum200ResponseValuesInnerWithDefaults() *GetTimeSeriesSum200ResponseValuesInner`

NewGetTimeSeriesSum200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesSum200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesSum200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesSum200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesSum200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetSum

`func (o *GetTimeSeriesSum200ResponseValuesInner) GetSum() string`

GetSum returns the Sum field if non-nil, zero value otherwise.

### GetSumOk

`func (o *GetTimeSeriesSum200ResponseValuesInner) GetSumOk() (*string, bool)`

GetSumOk returns a tuple with the Sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSum

`func (o *GetTimeSeriesSum200ResponseValuesInner) SetSum(v string)`

SetSum sets Sum field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


