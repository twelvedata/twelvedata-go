# GetTimeSeriesMacdExt200ResponseMetaIndicator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the technical indicator | 
**SeriesType** | **string** | Price type on which technical indicator is calculated | 
**FastPeriod** | **int64** | The shorter time period for calculation | 
**FastMaType** | **string** | The type of fast moving average used in the calculation | 
**SlowPeriod** | **int64** | The longer time period for calculation | 
**SlowMaType** | **string** | The type of slow moving average used in the calculation | 
**SignalPeriod** | **int64** | The time period used for generating the signal line | 
**SignalMaType** | **string** | The type of moving average used for generating the signal line | 

## Methods

### NewGetTimeSeriesMacdExt200ResponseMetaIndicator

`func NewGetTimeSeriesMacdExt200ResponseMetaIndicator(name string, seriesType string, fastPeriod int64, fastMaType string, slowPeriod int64, slowMaType string, signalPeriod int64, signalMaType string, ) *GetTimeSeriesMacdExt200ResponseMetaIndicator`

NewGetTimeSeriesMacdExt200ResponseMetaIndicator instantiates a new GetTimeSeriesMacdExt200ResponseMetaIndicator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesMacdExt200ResponseMetaIndicatorWithDefaults

`func NewGetTimeSeriesMacdExt200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesMacdExt200ResponseMetaIndicator`

NewGetTimeSeriesMacdExt200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesMacdExt200ResponseMetaIndicator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) SetName(v string)`

SetName sets Name field to given value.


### GetSeriesType

`func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) GetSeriesType() string`

GetSeriesType returns the SeriesType field if non-nil, zero value otherwise.

### GetSeriesTypeOk

`func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) GetSeriesTypeOk() (*string, bool)`

GetSeriesTypeOk returns a tuple with the SeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesType

`func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) SetSeriesType(v string)`

SetSeriesType sets SeriesType field to given value.


### GetFastPeriod

`func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) GetFastPeriod() int64`

GetFastPeriod returns the FastPeriod field if non-nil, zero value otherwise.

### GetFastPeriodOk

`func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) GetFastPeriodOk() (*int64, bool)`

GetFastPeriodOk returns a tuple with the FastPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastPeriod

`func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) SetFastPeriod(v int64)`

SetFastPeriod sets FastPeriod field to given value.


### GetFastMaType

`func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) GetFastMaType() string`

GetFastMaType returns the FastMaType field if non-nil, zero value otherwise.

### GetFastMaTypeOk

`func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) GetFastMaTypeOk() (*string, bool)`

GetFastMaTypeOk returns a tuple with the FastMaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastMaType

`func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) SetFastMaType(v string)`

SetFastMaType sets FastMaType field to given value.


### GetSlowPeriod

`func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) GetSlowPeriod() int64`

GetSlowPeriod returns the SlowPeriod field if non-nil, zero value otherwise.

### GetSlowPeriodOk

`func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) GetSlowPeriodOk() (*int64, bool)`

GetSlowPeriodOk returns a tuple with the SlowPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlowPeriod

`func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) SetSlowPeriod(v int64)`

SetSlowPeriod sets SlowPeriod field to given value.


### GetSlowMaType

`func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) GetSlowMaType() string`

GetSlowMaType returns the SlowMaType field if non-nil, zero value otherwise.

### GetSlowMaTypeOk

`func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) GetSlowMaTypeOk() (*string, bool)`

GetSlowMaTypeOk returns a tuple with the SlowMaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlowMaType

`func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) SetSlowMaType(v string)`

SetSlowMaType sets SlowMaType field to given value.


### GetSignalPeriod

`func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) GetSignalPeriod() int64`

GetSignalPeriod returns the SignalPeriod field if non-nil, zero value otherwise.

### GetSignalPeriodOk

`func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) GetSignalPeriodOk() (*int64, bool)`

GetSignalPeriodOk returns a tuple with the SignalPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalPeriod

`func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) SetSignalPeriod(v int64)`

SetSignalPeriod sets SignalPeriod field to given value.


### GetSignalMaType

`func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) GetSignalMaType() string`

GetSignalMaType returns the SignalMaType field if non-nil, zero value otherwise.

### GetSignalMaTypeOk

`func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) GetSignalMaTypeOk() (*string, bool)`

GetSignalMaTypeOk returns a tuple with the SignalMaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalMaType

`func (o *GetTimeSeriesMacdExt200ResponseMetaIndicator) SetSignalMaType(v string)`

SetSignalMaType sets SignalMaType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


