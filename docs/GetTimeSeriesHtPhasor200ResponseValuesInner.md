# GetTimeSeriesHtPhasor200ResponseValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **string** | Datetime in local market time for equities and in UTC for forex and cryptocurrencies referring to when the bar with specified interval was opened | 
**InPhase** | **string** | In_phase value | 
**Quadrature** | **string** | Quadrature value | 

## Methods

### NewGetTimeSeriesHtPhasor200ResponseValuesInner

`func NewGetTimeSeriesHtPhasor200ResponseValuesInner(datetime string, inPhase string, quadrature string, ) *GetTimeSeriesHtPhasor200ResponseValuesInner`

NewGetTimeSeriesHtPhasor200ResponseValuesInner instantiates a new GetTimeSeriesHtPhasor200ResponseValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesHtPhasor200ResponseValuesInnerWithDefaults

`func NewGetTimeSeriesHtPhasor200ResponseValuesInnerWithDefaults() *GetTimeSeriesHtPhasor200ResponseValuesInner`

NewGetTimeSeriesHtPhasor200ResponseValuesInnerWithDefaults instantiates a new GetTimeSeriesHtPhasor200ResponseValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetTimeSeriesHtPhasor200ResponseValuesInner) GetDatetime() string`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetTimeSeriesHtPhasor200ResponseValuesInner) GetDatetimeOk() (*string, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetTimeSeriesHtPhasor200ResponseValuesInner) SetDatetime(v string)`

SetDatetime sets Datetime field to given value.


### GetInPhase

`func (o *GetTimeSeriesHtPhasor200ResponseValuesInner) GetInPhase() string`

GetInPhase returns the InPhase field if non-nil, zero value otherwise.

### GetInPhaseOk

`func (o *GetTimeSeriesHtPhasor200ResponseValuesInner) GetInPhaseOk() (*string, bool)`

GetInPhaseOk returns a tuple with the InPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInPhase

`func (o *GetTimeSeriesHtPhasor200ResponseValuesInner) SetInPhase(v string)`

SetInPhase sets InPhase field to given value.


### GetQuadrature

`func (o *GetTimeSeriesHtPhasor200ResponseValuesInner) GetQuadrature() string`

GetQuadrature returns the Quadrature field if non-nil, zero value otherwise.

### GetQuadratureOk

`func (o *GetTimeSeriesHtPhasor200ResponseValuesInner) GetQuadratureOk() (*string, bool)`

GetQuadratureOk returns a tuple with the Quadrature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuadrature

`func (o *GetTimeSeriesHtPhasor200ResponseValuesInner) SetQuadrature(v string)`

SetQuadrature sets Quadrature field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


