# GetTimeSeriesDpo200ResponseMetaIndicator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the technical indicator | 
**SeriesType** | **string** | Price type on which technical indicator is calculated | 
**TimePeriod** | **int64** | Number of periods to average over | 
**Centered** | **bool** | Specifies if there should be a shift to match the current price | 

## Methods

### NewGetTimeSeriesDpo200ResponseMetaIndicator

`func NewGetTimeSeriesDpo200ResponseMetaIndicator(name string, seriesType string, timePeriod int64, centered bool, ) *GetTimeSeriesDpo200ResponseMetaIndicator`

NewGetTimeSeriesDpo200ResponseMetaIndicator instantiates a new GetTimeSeriesDpo200ResponseMetaIndicator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesDpo200ResponseMetaIndicatorWithDefaults

`func NewGetTimeSeriesDpo200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesDpo200ResponseMetaIndicator`

NewGetTimeSeriesDpo200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesDpo200ResponseMetaIndicator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetTimeSeriesDpo200ResponseMetaIndicator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetTimeSeriesDpo200ResponseMetaIndicator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetTimeSeriesDpo200ResponseMetaIndicator) SetName(v string)`

SetName sets Name field to given value.


### GetSeriesType

`func (o *GetTimeSeriesDpo200ResponseMetaIndicator) GetSeriesType() string`

GetSeriesType returns the SeriesType field if non-nil, zero value otherwise.

### GetSeriesTypeOk

`func (o *GetTimeSeriesDpo200ResponseMetaIndicator) GetSeriesTypeOk() (*string, bool)`

GetSeriesTypeOk returns a tuple with the SeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesType

`func (o *GetTimeSeriesDpo200ResponseMetaIndicator) SetSeriesType(v string)`

SetSeriesType sets SeriesType field to given value.


### GetTimePeriod

`func (o *GetTimeSeriesDpo200ResponseMetaIndicator) GetTimePeriod() int64`

GetTimePeriod returns the TimePeriod field if non-nil, zero value otherwise.

### GetTimePeriodOk

`func (o *GetTimeSeriesDpo200ResponseMetaIndicator) GetTimePeriodOk() (*int64, bool)`

GetTimePeriodOk returns a tuple with the TimePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePeriod

`func (o *GetTimeSeriesDpo200ResponseMetaIndicator) SetTimePeriod(v int64)`

SetTimePeriod sets TimePeriod field to given value.


### GetCentered

`func (o *GetTimeSeriesDpo200ResponseMetaIndicator) GetCentered() bool`

GetCentered returns the Centered field if non-nil, zero value otherwise.

### GetCenteredOk

`func (o *GetTimeSeriesDpo200ResponseMetaIndicator) GetCenteredOk() (*bool, bool)`

GetCenteredOk returns a tuple with the Centered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCentered

`func (o *GetTimeSeriesDpo200ResponseMetaIndicator) SetCentered(v bool)`

SetCentered sets Centered field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


