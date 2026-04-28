# GetTimeSeriesApo200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Apo** | **string** | APO value | 

## Methods

### NewGetTimeSeriesApo200ResponseValuesInner

`func NewGetTimeSeriesApo200ResponseValuesInner(datetime string, apo string, ) *GetTimeSeriesApo200ResponseValuesInner`

NewGetTimeSeriesApo200ResponseValuesInner instantiates a new GetTimeSeriesApo200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesApo200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesApo200ResponseValuesInnerWithDefaults() *GetTimeSeriesApo200ResponseValuesInner`

NewGetTimeSeriesApo200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesApo200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesApo200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesApo200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesApo200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetApo

`func (o *GetTimeSeriesApo200ResponseValuesInner) GetApo() string`

GetApo returns the Apo field if non-nil, zero value otherwise.

### GetApoOk

`func (o *GetTimeSeriesApo200ResponseValuesInner) GetApoOk() (*string, bool)`

GetApoOk returns a tuple with the Apo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApo

`func (o *GetTimeSeriesApo200ResponseValuesInner) SetApo(v string)`

SetApo sets Apo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


