# GetTimeSeriesMinMax200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Min** | **string** | Min value | 
**Max** | **string** | Max value | 

## Methods

### NewGetTimeSeriesMinMax200ResponseValuesInner

`func NewGetTimeSeriesMinMax200ResponseValuesInner(datetime string, min string, max string, ) *GetTimeSeriesMinMax200ResponseValuesInner`

NewGetTimeSeriesMinMax200ResponseValuesInner instantiates a new GetTimeSeriesMinMax200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesMinMax200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesMinMax200ResponseValuesInnerWithDefaults() *GetTimeSeriesMinMax200ResponseValuesInner`

NewGetTimeSeriesMinMax200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesMinMax200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesMinMax200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesMinMax200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesMinMax200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetMin

`func (o *GetTimeSeriesMinMax200ResponseValuesInner) GetMin() string`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *GetTimeSeriesMinMax200ResponseValuesInner) GetMinOk() (*string, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *GetTimeSeriesMinMax200ResponseValuesInner) SetMin(v string)`

SetMin sets Min field to given value.


### GetMax

`func (o *GetTimeSeriesMinMax200ResponseValuesInner) GetMax() string`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *GetTimeSeriesMinMax200ResponseValuesInner) GetMaxOk() (*string, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *GetTimeSeriesMinMax200ResponseValuesInner) SetMax(v string)`

SetMax sets Max field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


