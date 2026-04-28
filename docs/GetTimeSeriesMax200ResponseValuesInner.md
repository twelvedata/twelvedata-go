# GetTimeSeriesMax200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Max** | **string** | Max value | 

## Methods

### NewGetTimeSeriesMax200ResponseValuesInner

`func NewGetTimeSeriesMax200ResponseValuesInner(datetime string, max string, ) *GetTimeSeriesMax200ResponseValuesInner`

NewGetTimeSeriesMax200ResponseValuesInner instantiates a new GetTimeSeriesMax200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesMax200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesMax200ResponseValuesInnerWithDefaults() *GetTimeSeriesMax200ResponseValuesInner`

NewGetTimeSeriesMax200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesMax200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesMax200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesMax200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesMax200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetMax

`func (o *GetTimeSeriesMax200ResponseValuesInner) GetMax() string`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *GetTimeSeriesMax200ResponseValuesInner) GetMaxOk() (*string, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *GetTimeSeriesMax200ResponseValuesInner) SetMax(v string)`

SetMax sets Max field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


