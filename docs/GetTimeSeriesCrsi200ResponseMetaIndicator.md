# GetTimeSeriesCrsi200ResponseMetaIndicator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the technical indicator | 
**SeriesType** | **string** | Price type on which technical indicator is calculated | 
**RsiPeriod** | **int64** | Number of periods for RSI used to calculate price momentum | 
**UpDownLength** | **int64** | Number of periods for RSI used to calculate up/down trend | 
**PercentRankPeriod** | **int64** | Number of periods used to calculate PercentRank | 

## Methods

### NewGetTimeSeriesCrsi200ResponseMetaIndicator

`func NewGetTimeSeriesCrsi200ResponseMetaIndicator(name string, seriesType string, rsiPeriod int64, upDownLength int64, percentRankPeriod int64, ) *GetTimeSeriesCrsi200ResponseMetaIndicator`

NewGetTimeSeriesCrsi200ResponseMetaIndicator instantiates a new GetTimeSeriesCrsi200ResponseMetaIndicator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesCrsi200ResponseMetaIndicatorWithDefaults

`func NewGetTimeSeriesCrsi200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesCrsi200ResponseMetaIndicator`

NewGetTimeSeriesCrsi200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesCrsi200ResponseMetaIndicator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetTimeSeriesCrsi200ResponseMetaIndicator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetTimeSeriesCrsi200ResponseMetaIndicator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetTimeSeriesCrsi200ResponseMetaIndicator) SetName(v string)`

SetName sets Name field to given value.


### GetSeriesType

`func (o *GetTimeSeriesCrsi200ResponseMetaIndicator) GetSeriesType() string`

GetSeriesType returns the SeriesType field if non-nil, zero value otherwise.

### GetSeriesTypeOk

`func (o *GetTimeSeriesCrsi200ResponseMetaIndicator) GetSeriesTypeOk() (*string, bool)`

GetSeriesTypeOk returns a tuple with the SeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesType

`func (o *GetTimeSeriesCrsi200ResponseMetaIndicator) SetSeriesType(v string)`

SetSeriesType sets SeriesType field to given value.


### GetRsiPeriod

`func (o *GetTimeSeriesCrsi200ResponseMetaIndicator) GetRsiPeriod() int64`

GetRsiPeriod returns the RsiPeriod field if non-nil, zero value otherwise.

### GetRsiPeriodOk

`func (o *GetTimeSeriesCrsi200ResponseMetaIndicator) GetRsiPeriodOk() (*int64, bool)`

GetRsiPeriodOk returns a tuple with the RsiPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsiPeriod

`func (o *GetTimeSeriesCrsi200ResponseMetaIndicator) SetRsiPeriod(v int64)`

SetRsiPeriod sets RsiPeriod field to given value.


### GetUpDownLength

`func (o *GetTimeSeriesCrsi200ResponseMetaIndicator) GetUpDownLength() int64`

GetUpDownLength returns the UpDownLength field if non-nil, zero value otherwise.

### GetUpDownLengthOk

`func (o *GetTimeSeriesCrsi200ResponseMetaIndicator) GetUpDownLengthOk() (*int64, bool)`

GetUpDownLengthOk returns a tuple with the UpDownLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpDownLength

`func (o *GetTimeSeriesCrsi200ResponseMetaIndicator) SetUpDownLength(v int64)`

SetUpDownLength sets UpDownLength field to given value.


### GetPercentRankPeriod

`func (o *GetTimeSeriesCrsi200ResponseMetaIndicator) GetPercentRankPeriod() int64`

GetPercentRankPeriod returns the PercentRankPeriod field if non-nil, zero value otherwise.

### GetPercentRankPeriodOk

`func (o *GetTimeSeriesCrsi200ResponseMetaIndicator) GetPercentRankPeriodOk() (*int64, bool)`

GetPercentRankPeriodOk returns a tuple with the PercentRankPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentRankPeriod

`func (o *GetTimeSeriesCrsi200ResponseMetaIndicator) SetPercentRankPeriod(v int64)`

SetPercentRankPeriod sets PercentRankPeriod field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


