# GetTimeSeriesBBands200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**UpperBand** | **string** | Upper band value | 
**MiddleBand** | **string** | Middle band value | 
**LowerBand** | **string** | Lower band value | 

## Methods

### NewGetTimeSeriesBBands200ResponseValuesInner

`func NewGetTimeSeriesBBands200ResponseValuesInner(datetime string, upperBand string, middleBand string, lowerBand string, ) *GetTimeSeriesBBands200ResponseValuesInner`

NewGetTimeSeriesBBands200ResponseValuesInner instantiates a new GetTimeSeriesBBands200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesBBands200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesBBands200ResponseValuesInnerWithDefaults() *GetTimeSeriesBBands200ResponseValuesInner`

NewGetTimeSeriesBBands200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesBBands200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesBBands200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesBBands200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesBBands200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetUpperBand

`func (o *GetTimeSeriesBBands200ResponseValuesInner) GetUpperBand() string`

GetUpperBand returns the UpperBand field if non-nil, zero value otherwise.

### GetUpperBandOk

`func (o *GetTimeSeriesBBands200ResponseValuesInner) GetUpperBandOk() (*string, bool)`

GetUpperBandOk returns a tuple with the UpperBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperBand

`func (o *GetTimeSeriesBBands200ResponseValuesInner) SetUpperBand(v string)`

SetUpperBand sets UpperBand field to given value.


### GetMiddleBand

`func (o *GetTimeSeriesBBands200ResponseValuesInner) GetMiddleBand() string`

GetMiddleBand returns the MiddleBand field if non-nil, zero value otherwise.

### GetMiddleBandOk

`func (o *GetTimeSeriesBBands200ResponseValuesInner) GetMiddleBandOk() (*string, bool)`

GetMiddleBandOk returns a tuple with the MiddleBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleBand

`func (o *GetTimeSeriesBBands200ResponseValuesInner) SetMiddleBand(v string)`

SetMiddleBand sets MiddleBand field to given value.


### GetLowerBand

`func (o *GetTimeSeriesBBands200ResponseValuesInner) GetLowerBand() string`

GetLowerBand returns the LowerBand field if non-nil, zero value otherwise.

### GetLowerBandOk

`func (o *GetTimeSeriesBBands200ResponseValuesInner) GetLowerBandOk() (*string, bool)`

GetLowerBandOk returns a tuple with the LowerBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerBand

`func (o *GetTimeSeriesBBands200ResponseValuesInner) SetLowerBand(v string)`

SetLowerBand sets LowerBand field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


