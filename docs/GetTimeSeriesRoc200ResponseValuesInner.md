# GetTimeSeriesRoc200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Roc** | **string** | roc value | 

## Methods

### NewGetTimeSeriesRoc200ResponseValuesInner

`func NewGetTimeSeriesRoc200ResponseValuesInner(datetime string, roc string, ) *GetTimeSeriesRoc200ResponseValuesInner`

NewGetTimeSeriesRoc200ResponseValuesInner instantiates a new GetTimeSeriesRoc200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesRoc200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesRoc200ResponseValuesInnerWithDefaults() *GetTimeSeriesRoc200ResponseValuesInner`

NewGetTimeSeriesRoc200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesRoc200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesRoc200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesRoc200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesRoc200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetRoc

`func (o *GetTimeSeriesRoc200ResponseValuesInner) GetRoc() string`

GetRoc returns the Roc field if non-nil, zero value otherwise.

### GetRocOk

`func (o *GetTimeSeriesRoc200ResponseValuesInner) GetRocOk() (*string, bool)`

GetRocOk returns a tuple with the Roc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoc

`func (o *GetTimeSeriesRoc200ResponseValuesInner) SetRoc(v string)`

SetRoc sets Roc field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


