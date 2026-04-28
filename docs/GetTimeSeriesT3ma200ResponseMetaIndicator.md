# GetTimeSeriesT3ma200ResponseMetaIndicator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the technical indicator | 
**SeriesType** | **string** | Price type on which technical indicator is calculated | 
**TimePeriod** | **int64** | The time period used for calculation in the indicator | 
**VFactor** | **float64** | The factor used to adjust the indicator&#39;s volatility | 

## Methods

### NewGetTimeSeriesT3ma200ResponseMetaIndicator

`func NewGetTimeSeriesT3ma200ResponseMetaIndicator(name string, seriesType string, timePeriod int64, vFactor float64, ) *GetTimeSeriesT3ma200ResponseMetaIndicator`

NewGetTimeSeriesT3ma200ResponseMetaIndicator instantiates a new GetTimeSeriesT3ma200ResponseMetaIndicator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesT3ma200ResponseMetaIndicatorWithDefaults

`func NewGetTimeSeriesT3ma200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesT3ma200ResponseMetaIndicator`

NewGetTimeSeriesT3ma200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesT3ma200ResponseMetaIndicator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetTimeSeriesT3ma200ResponseMetaIndicator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetTimeSeriesT3ma200ResponseMetaIndicator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetTimeSeriesT3ma200ResponseMetaIndicator) SetName(v string)`

SetName sets Name field to given value.


### GetSeriesType

`func (o *GetTimeSeriesT3ma200ResponseMetaIndicator) GetSeriesType() string`

GetSeriesType returns the SeriesType field if non-nil, zero value otherwise.

### GetSeriesTypeOk

`func (o *GetTimeSeriesT3ma200ResponseMetaIndicator) GetSeriesTypeOk() (*string, bool)`

GetSeriesTypeOk returns a tuple with the SeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesType

`func (o *GetTimeSeriesT3ma200ResponseMetaIndicator) SetSeriesType(v string)`

SetSeriesType sets SeriesType field to given value.


### GetTimePeriod

`func (o *GetTimeSeriesT3ma200ResponseMetaIndicator) GetTimePeriod() int64`

GetTimePeriod returns the TimePeriod field if non-nil, zero value otherwise.

### GetTimePeriodOk

`func (o *GetTimeSeriesT3ma200ResponseMetaIndicator) GetTimePeriodOk() (*int64, bool)`

GetTimePeriodOk returns a tuple with the TimePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePeriod

`func (o *GetTimeSeriesT3ma200ResponseMetaIndicator) SetTimePeriod(v int64)`

SetTimePeriod sets TimePeriod field to given value.


### GetVFactor

`func (o *GetTimeSeriesT3ma200ResponseMetaIndicator) GetVFactor() float64`

GetVFactor returns the VFactor field if non-nil, zero value otherwise.

### GetVFactorOk

`func (o *GetTimeSeriesT3ma200ResponseMetaIndicator) GetVFactorOk() (*float64, bool)`

GetVFactorOk returns a tuple with the VFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVFactor

`func (o *GetTimeSeriesT3ma200ResponseMetaIndicator) SetVFactor(v float64)`

SetVFactor sets VFactor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


