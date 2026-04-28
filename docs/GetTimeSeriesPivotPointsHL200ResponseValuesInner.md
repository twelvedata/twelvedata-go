# GetTimeSeriesPivotPointsHL200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**PivotPointH** | **int64** | &#x60;1&#x60; if it is a high pivot point, otherwise &#x60;0&#x60; | 
**PivotPointL** | **int64** | &#x60;1&#x60; if it is a low pivot point, otherwise &#x60;0&#x60; | 

## Methods

### NewGetTimeSeriesPivotPointsHL200ResponseValuesInner

`func NewGetTimeSeriesPivotPointsHL200ResponseValuesInner(datetime string, pivotPointH int64, pivotPointL int64, ) *GetTimeSeriesPivotPointsHL200ResponseValuesInner`

NewGetTimeSeriesPivotPointsHL200ResponseValuesInner instantiates a new GetTimeSeriesPivotPointsHL200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesPivotPointsHL200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesPivotPointsHL200ResponseValuesInnerWithDefaults() *GetTimeSeriesPivotPointsHL200ResponseValuesInner`

NewGetTimeSeriesPivotPointsHL200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesPivotPointsHL200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesPivotPointsHL200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesPivotPointsHL200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesPivotPointsHL200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetPivotPointH

`func (o *GetTimeSeriesPivotPointsHL200ResponseValuesInner) GetPivotPointH() int64`

GetPivotPointH returns the PivotPointH field if non-nil, zero value otherwise.

### GetPivotPointHOk

`func (o *GetTimeSeriesPivotPointsHL200ResponseValuesInner) GetPivotPointHOk() (*int64, bool)`

GetPivotPointHOk returns a tuple with the PivotPointH field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPivotPointH

`func (o *GetTimeSeriesPivotPointsHL200ResponseValuesInner) SetPivotPointH(v int64)`

SetPivotPointH sets PivotPointH field to given value.


### GetPivotPointL

`func (o *GetTimeSeriesPivotPointsHL200ResponseValuesInner) GetPivotPointL() int64`

GetPivotPointL returns the PivotPointL field if non-nil, zero value otherwise.

### GetPivotPointLOk

`func (o *GetTimeSeriesPivotPointsHL200ResponseValuesInner) GetPivotPointLOk() (*int64, bool)`

GetPivotPointLOk returns a tuple with the PivotPointL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPivotPointL

`func (o *GetTimeSeriesPivotPointsHL200ResponseValuesInner) SetPivotPointL(v int64)`

SetPivotPointL sets PivotPointL field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


