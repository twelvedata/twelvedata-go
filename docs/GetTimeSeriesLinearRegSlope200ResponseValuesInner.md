# GetTimeSeriesLinearRegSlope200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Linearregslope** | **string** | linearregslope value | 

## Methods

### NewGetTimeSeriesLinearRegSlope200ResponseValuesInner

`func NewGetTimeSeriesLinearRegSlope200ResponseValuesInner(datetime string, linearregslope string, ) *GetTimeSeriesLinearRegSlope200ResponseValuesInner`

NewGetTimeSeriesLinearRegSlope200ResponseValuesInner instantiates a new GetTimeSeriesLinearRegSlope200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesLinearRegSlope200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesLinearRegSlope200ResponseValuesInnerWithDefaults() *GetTimeSeriesLinearRegSlope200ResponseValuesInner`

NewGetTimeSeriesLinearRegSlope200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesLinearRegSlope200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesLinearRegSlope200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesLinearRegSlope200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesLinearRegSlope200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetLinearregslope

`func (o *GetTimeSeriesLinearRegSlope200ResponseValuesInner) GetLinearregslope() string`

GetLinearregslope returns the Linearregslope field if non-nil, zero value otherwise.

### GetLinearregslopeOk

`func (o *GetTimeSeriesLinearRegSlope200ResponseValuesInner) GetLinearregslopeOk() (*string, bool)`

GetLinearregslopeOk returns a tuple with the Linearregslope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinearregslope

`func (o *GetTimeSeriesLinearRegSlope200ResponseValuesInner) SetLinearregslope(v string)`

SetLinearregslope sets Linearregslope field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


