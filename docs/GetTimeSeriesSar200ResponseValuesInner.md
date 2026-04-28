# GetTimeSeriesSar200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Sar** | **string** | SAR value | 

## Methods

### NewGetTimeSeriesSar200ResponseValuesInner

`func NewGetTimeSeriesSar200ResponseValuesInner(datetime string, sar string, ) *GetTimeSeriesSar200ResponseValuesInner`

NewGetTimeSeriesSar200ResponseValuesInner instantiates a new GetTimeSeriesSar200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesSar200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesSar200ResponseValuesInnerWithDefaults() *GetTimeSeriesSar200ResponseValuesInner`

NewGetTimeSeriesSar200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesSar200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesSar200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesSar200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesSar200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetSar

`func (o *GetTimeSeriesSar200ResponseValuesInner) GetSar() string`

GetSar returns the Sar field if non-nil, zero value otherwise.

### GetSarOk

`func (o *GetTimeSeriesSar200ResponseValuesInner) GetSarOk() (*string, bool)`

GetSarOk returns a tuple with the Sar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSar

`func (o *GetTimeSeriesSar200ResponseValuesInner) SetSar(v string)`

SetSar sets Sar field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


