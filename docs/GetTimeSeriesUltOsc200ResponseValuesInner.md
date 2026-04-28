# GetTimeSeriesUltOsc200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**Ultosc** | **string** | Ultimate Oscillator value | 

## Methods

### NewGetTimeSeriesUltOsc200ResponseValuesInner

`func NewGetTimeSeriesUltOsc200ResponseValuesInner(datetime string, ultosc string, ) *GetTimeSeriesUltOsc200ResponseValuesInner`

NewGetTimeSeriesUltOsc200ResponseValuesInner instantiates a new GetTimeSeriesUltOsc200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesUltOsc200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesUltOsc200ResponseValuesInnerWithDefaults() *GetTimeSeriesUltOsc200ResponseValuesInner`

NewGetTimeSeriesUltOsc200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesUltOsc200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesUltOsc200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesUltOsc200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesUltOsc200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetUltosc

`func (o *GetTimeSeriesUltOsc200ResponseValuesInner) GetUltosc() string`

GetUltosc returns the Ultosc field if non-nil, zero value otherwise.

### GetUltoscOk

`func (o *GetTimeSeriesUltOsc200ResponseValuesInner) GetUltoscOk() (*string, bool)`

GetUltoscOk returns a tuple with the Ultosc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUltosc

`func (o *GetTimeSeriesUltOsc200ResponseValuesInner) SetUltosc(v string)`

SetUltosc sets Ultosc field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


