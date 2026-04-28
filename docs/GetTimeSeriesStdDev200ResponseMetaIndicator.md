# GetTimeSeriesStdDev200ResponseMetaIndicator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the technical indicator | 
**SeriesType** | **string** | Price type on which technical indicator is calculated | 
**TimePeriod** | **int64** | The time period used for calculation in the indicator | 
**Sd** | **float64** | The standard deviation applied in the calculation | 

## Methods

### NewGetTimeSeriesStdDev200ResponseMetaIndicator

`func NewGetTimeSeriesStdDev200ResponseMetaIndicator(name string, seriesType string, timePeriod int64, sd float64, ) *GetTimeSeriesStdDev200ResponseMetaIndicator`

NewGetTimeSeriesStdDev200ResponseMetaIndicator instantiates a new GetTimeSeriesStdDev200ResponseMetaIndicator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesStdDev200ResponseMetaIndicatorWithDefaults

`func NewGetTimeSeriesStdDev200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesStdDev200ResponseMetaIndicator`

NewGetTimeSeriesStdDev200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesStdDev200ResponseMetaIndicator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetTimeSeriesStdDev200ResponseMetaIndicator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetTimeSeriesStdDev200ResponseMetaIndicator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetTimeSeriesStdDev200ResponseMetaIndicator) SetName(v string)`

SetName sets Name field to given value.


### GetSeriesType

`func (o *GetTimeSeriesStdDev200ResponseMetaIndicator) GetSeriesType() string`

GetSeriesType returns the SeriesType field if non-nil, zero value otherwise.

### GetSeriesTypeOk

`func (o *GetTimeSeriesStdDev200ResponseMetaIndicator) GetSeriesTypeOk() (*string, bool)`

GetSeriesTypeOk returns a tuple with the SeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesType

`func (o *GetTimeSeriesStdDev200ResponseMetaIndicator) SetSeriesType(v string)`

SetSeriesType sets SeriesType field to given value.


### GetTimePeriod

`func (o *GetTimeSeriesStdDev200ResponseMetaIndicator) GetTimePeriod() int64`

GetTimePeriod returns the TimePeriod field if non-nil, zero value otherwise.

### GetTimePeriodOk

`func (o *GetTimeSeriesStdDev200ResponseMetaIndicator) GetTimePeriodOk() (*int64, bool)`

GetTimePeriodOk returns a tuple with the TimePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePeriod

`func (o *GetTimeSeriesStdDev200ResponseMetaIndicator) SetTimePeriod(v int64)`

SetTimePeriod sets TimePeriod field to given value.


### GetSd

`func (o *GetTimeSeriesStdDev200ResponseMetaIndicator) GetSd() float64`

GetSd returns the Sd field if non-nil, zero value otherwise.

### GetSdOk

`func (o *GetTimeSeriesStdDev200ResponseMetaIndicator) GetSdOk() (*float64, bool)`

GetSdOk returns a tuple with the Sd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSd

`func (o *GetTimeSeriesStdDev200ResponseMetaIndicator) SetSd(v float64)`

SetSd sets Sd field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


