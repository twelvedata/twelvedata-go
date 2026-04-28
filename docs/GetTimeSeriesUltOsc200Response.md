# GetTimeSeriesUltOsc200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesUltOsc200ResponseMeta**](GetTimeSeriesUltOsc200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesUltOsc200ResponseValuesInner**](GetTimeSeriesUltOsc200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesUltOsc200Response

`func NewGetTimeSeriesUltOsc200Response(meta GetTimeSeriesUltOsc200ResponseMeta, values []GetTimeSeriesUltOsc200ResponseValuesInner, status string, ) *GetTimeSeriesUltOsc200Response`

NewGetTimeSeriesUltOsc200Response instantiates a new GetTimeSeriesUltOsc200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesUltOsc200ResponseWithDefaults

`func NewGetTimeSeriesUltOsc200ResponseWithDefaults() *GetTimeSeriesUltOsc200Response`

NewGetTimeSeriesUltOsc200ResponseWithDefaults instantiates a new GetTimeSeriesUltOsc200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesUltOsc200Response) GetMeta() GetTimeSeriesUltOsc200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesUltOsc200Response) GetMetaOk() (*GetTimeSeriesUltOsc200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesUltOsc200Response) SetMeta(v GetTimeSeriesUltOsc200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesUltOsc200Response) GetValues() []GetTimeSeriesUltOsc200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesUltOsc200Response) GetValuesOk() (*[]GetTimeSeriesUltOsc200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesUltOsc200Response) SetValues(v []GetTimeSeriesUltOsc200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesUltOsc200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesUltOsc200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesUltOsc200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


