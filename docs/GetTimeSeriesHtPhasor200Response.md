# GetTimeSeriesHtPhasor200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesHtPhasor200ResponseMeta**](GetTimeSeriesHtPhasor200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesHtPhasor200ResponseValuesInner**](GetTimeSeriesHtPhasor200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesHtPhasor200Response

`func NewGetTimeSeriesHtPhasor200Response(meta GetTimeSeriesHtPhasor200ResponseMeta, values []GetTimeSeriesHtPhasor200ResponseValuesInner, status string, ) *GetTimeSeriesHtPhasor200Response`

NewGetTimeSeriesHtPhasor200Response instantiates a new GetTimeSeriesHtPhasor200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesHtPhasor200ResponseWithDefaults

`func NewGetTimeSeriesHtPhasor200ResponseWithDefaults() *GetTimeSeriesHtPhasor200Response`

NewGetTimeSeriesHtPhasor200ResponseWithDefaults instantiates a new GetTimeSeriesHtPhasor200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesHtPhasor200Response) GetMeta() GetTimeSeriesHtPhasor200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesHtPhasor200Response) GetMetaOk() (*GetTimeSeriesHtPhasor200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesHtPhasor200Response) SetMeta(v GetTimeSeriesHtPhasor200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesHtPhasor200Response) GetValues() []GetTimeSeriesHtPhasor200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesHtPhasor200Response) GetValuesOk() (*[]GetTimeSeriesHtPhasor200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesHtPhasor200Response) SetValues(v []GetTimeSeriesHtPhasor200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesHtPhasor200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesHtPhasor200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesHtPhasor200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


