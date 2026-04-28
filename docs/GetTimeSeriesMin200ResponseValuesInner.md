# GetTimeSeriesMin200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Min** | **string** | Min value | 

## Methods

### NewGetTimeSeriesMin200ResponseValuesInner

`func NewGetTimeSeriesMin200ResponseValuesInner(datetime string, min string, ) *GetTimeSeriesMin200ResponseValuesInner`

NewGetTimeSeriesMin200ResponseValuesInner instantiates a new GetTimeSeriesMin200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesMin200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesMin200ResponseValuesInnerWithDefaults() *GetTimeSeriesMin200ResponseValuesInner`

NewGetTimeSeriesMin200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesMin200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesMin200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesMin200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesMin200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetMin

`func (o *GetTimeSeriesMin200ResponseValuesInner) GetMin() string`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *GetTimeSeriesMin200ResponseValuesInner) GetMinOk() (*string, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *GetTimeSeriesMin200ResponseValuesInner) SetMin(v string)`

SetMin sets Min field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


