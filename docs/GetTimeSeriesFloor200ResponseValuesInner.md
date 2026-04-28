# GetTimeSeriesFloor200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Floor** | **string** | Floor value | 

## Methods

### NewGetTimeSeriesFloor200ResponseValuesInner

`func NewGetTimeSeriesFloor200ResponseValuesInner(datetime string, floor string, ) *GetTimeSeriesFloor200ResponseValuesInner`

NewGetTimeSeriesFloor200ResponseValuesInner instantiates a new GetTimeSeriesFloor200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesFloor200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesFloor200ResponseValuesInnerWithDefaults() *GetTimeSeriesFloor200ResponseValuesInner`

NewGetTimeSeriesFloor200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesFloor200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesFloor200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesFloor200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesFloor200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetFloor

`func (o *GetTimeSeriesFloor200ResponseValuesInner) GetFloor() string`

GetFloor returns the Floor field if non-nil, zero value otherwise.

### GetFloorOk

`func (o *GetTimeSeriesFloor200ResponseValuesInner) GetFloorOk() (*string, bool)`

GetFloorOk returns a tuple with the Floor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloor

`func (o *GetTimeSeriesFloor200ResponseValuesInner) SetFloor(v string)`

SetFloor sets Floor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


