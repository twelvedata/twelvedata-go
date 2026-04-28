# GetTimeSeriesMa200ResponseMetaIndicator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the technical indicator | 
**SeriesType** | **string** | Price type on which technical indicator is calculated | 
**TimePeriod** | **int64** | Number of periods to average over | 
**MaType** | **string** | The type of moving average used | 

## Methods

### NewGetTimeSeriesMa200ResponseMetaIndicator

`func NewGetTimeSeriesMa200ResponseMetaIndicator(name string, seriesType string, timePeriod int64, maType string, ) *GetTimeSeriesMa200ResponseMetaIndicator`

NewGetTimeSeriesMa200ResponseMetaIndicator instantiates a new GetTimeSeriesMa200ResponseMetaIndicator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesMa200ResponseMetaIndicatorWithDefaults

`func NewGetTimeSeriesMa200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesMa200ResponseMetaIndicator`

NewGetTimeSeriesMa200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesMa200ResponseMetaIndicator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetTimeSeriesMa200ResponseMetaIndicator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetTimeSeriesMa200ResponseMetaIndicator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetTimeSeriesMa200ResponseMetaIndicator) SetName(v string)`

SetName sets Name field to given value.


### GetSeriesType

`func (o *GetTimeSeriesMa200ResponseMetaIndicator) GetSeriesType() string`

GetSeriesType returns the SeriesType field if non-nil, zero value otherwise.

### GetSeriesTypeOk

`func (o *GetTimeSeriesMa200ResponseMetaIndicator) GetSeriesTypeOk() (*string, bool)`

GetSeriesTypeOk returns a tuple with the SeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesType

`func (o *GetTimeSeriesMa200ResponseMetaIndicator) SetSeriesType(v string)`

SetSeriesType sets SeriesType field to given value.


### GetTimePeriod

`func (o *GetTimeSeriesMa200ResponseMetaIndicator) GetTimePeriod() int64`

GetTimePeriod returns the TimePeriod field if non-nil, zero value otherwise.

### GetTimePeriodOk

`func (o *GetTimeSeriesMa200ResponseMetaIndicator) GetTimePeriodOk() (*int64, bool)`

GetTimePeriodOk returns a tuple with the TimePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePeriod

`func (o *GetTimeSeriesMa200ResponseMetaIndicator) SetTimePeriod(v int64)`

SetTimePeriod sets TimePeriod field to given value.


### GetMaType

`func (o *GetTimeSeriesMa200ResponseMetaIndicator) GetMaType() string`

GetMaType returns the MaType field if non-nil, zero value otherwise.

### GetMaTypeOk

`func (o *GetTimeSeriesMa200ResponseMetaIndicator) GetMaTypeOk() (*string, bool)`

GetMaTypeOk returns a tuple with the MaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaType

`func (o *GetTimeSeriesMa200ResponseMetaIndicator) SetMaType(v string)`

SetMaType sets MaType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


