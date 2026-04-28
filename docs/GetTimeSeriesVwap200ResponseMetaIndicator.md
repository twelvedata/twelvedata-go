# GetTimeSeriesVwap200ResponseMetaIndicator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the technical indicator | 
**SdTimePeriod** | Pointer to **int64** | Standard deviation time period | [optional] 
**Sd** | Pointer to **float64** | Standard deviation value | [optional] 

## Methods

### NewGetTimeSeriesVwap200ResponseMetaIndicator

`func NewGetTimeSeriesVwap200ResponseMetaIndicator(name string, ) *GetTimeSeriesVwap200ResponseMetaIndicator`

NewGetTimeSeriesVwap200ResponseMetaIndicator instantiates a new GetTimeSeriesVwap200ResponseMetaIndicator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesVwap200ResponseMetaIndicatorWithDefaults

`func NewGetTimeSeriesVwap200ResponseMetaIndicatorWithDefaults() *GetTimeSeriesVwap200ResponseMetaIndicator`

NewGetTimeSeriesVwap200ResponseMetaIndicatorWithDefaults instantiates a new GetTimeSeriesVwap200ResponseMetaIndicator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetTimeSeriesVwap200ResponseMetaIndicator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetTimeSeriesVwap200ResponseMetaIndicator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetTimeSeriesVwap200ResponseMetaIndicator) SetName(v string)`

SetName sets Name field to given value.


### GetSdTimePeriod

`func (o *GetTimeSeriesVwap200ResponseMetaIndicator) GetSdTimePeriod() int64`

GetSdTimePeriod returns the SdTimePeriod field if non-nil, zero value otherwise.

### GetSdTimePeriodOk

`func (o *GetTimeSeriesVwap200ResponseMetaIndicator) GetSdTimePeriodOk() (*int64, bool)`

GetSdTimePeriodOk returns a tuple with the SdTimePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdTimePeriod

`func (o *GetTimeSeriesVwap200ResponseMetaIndicator) SetSdTimePeriod(v int64)`

SetSdTimePeriod sets SdTimePeriod field to given value.

### HasSdTimePeriod

`func (o *GetTimeSeriesVwap200ResponseMetaIndicator) HasSdTimePeriod() bool`

HasSdTimePeriod returns a boolean if a field has been set.

### GetSd

`func (o *GetTimeSeriesVwap200ResponseMetaIndicator) GetSd() float64`

GetSd returns the Sd field if non-nil, zero value otherwise.

### GetSdOk

`func (o *GetTimeSeriesVwap200ResponseMetaIndicator) GetSdOk() (*float64, bool)`

GetSdOk returns a tuple with the Sd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSd

`func (o *GetTimeSeriesVwap200ResponseMetaIndicator) SetSd(v float64)`

SetSd sets Sd field to given value.

### HasSd

`func (o *GetTimeSeriesVwap200ResponseMetaIndicator) HasSd() bool`

HasSd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


