# GetTimeSeriesRocp200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Rocp** | **string** | ROCP value | 

## Methods

### NewGetTimeSeriesRocp200ResponseValuesInner

`func NewGetTimeSeriesRocp200ResponseValuesInner(datetime string, rocp string, ) *GetTimeSeriesRocp200ResponseValuesInner`

NewGetTimeSeriesRocp200ResponseValuesInner instantiates a new GetTimeSeriesRocp200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesRocp200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesRocp200ResponseValuesInnerWithDefaults() *GetTimeSeriesRocp200ResponseValuesInner`

NewGetTimeSeriesRocp200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesRocp200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesRocp200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesRocp200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesRocp200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetRocp

`func (o *GetTimeSeriesRocp200ResponseValuesInner) GetRocp() string`

GetRocp returns the Rocp field if non-nil, zero value otherwise.

### GetRocpOk

`func (o *GetTimeSeriesRocp200ResponseValuesInner) GetRocpOk() (*string, bool)`

GetRocpOk returns a tuple with the Rocp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRocp

`func (o *GetTimeSeriesRocp200ResponseValuesInner) SetRocp(v string)`

SetRocp sets Rocp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


