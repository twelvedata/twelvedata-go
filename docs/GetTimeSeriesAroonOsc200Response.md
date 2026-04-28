# GetTimeSeriesAroonOsc200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesAroonOsc200ResponseMeta**](GetTimeSeriesAroonOsc200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesAroonOsc200ResponseValuesInner**](GetTimeSeriesAroonOsc200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesAroonOsc200Response

`func NewGetTimeSeriesAroonOsc200Response(meta GetTimeSeriesAroonOsc200ResponseMeta, values []GetTimeSeriesAroonOsc200ResponseValuesInner, status string, ) *GetTimeSeriesAroonOsc200Response`

NewGetTimeSeriesAroonOsc200Response instantiates a new GetTimeSeriesAroonOsc200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesAroonOsc200ResponseWithDefaults

`func NewGetTimeSeriesAroonOsc200ResponseWithDefaults() *GetTimeSeriesAroonOsc200Response`

NewGetTimeSeriesAroonOsc200ResponseWithDefaults instantiates a new GetTimeSeriesAroonOsc200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesAroonOsc200Response) GetMeta() GetTimeSeriesAroonOsc200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesAroonOsc200Response) GetMetaOk() (*GetTimeSeriesAroonOsc200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesAroonOsc200Response) SetMeta(v GetTimeSeriesAroonOsc200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesAroonOsc200Response) GetValues() []GetTimeSeriesAroonOsc200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesAroonOsc200Response) GetValuesOk() (*[]GetTimeSeriesAroonOsc200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesAroonOsc200Response) SetValues(v []GetTimeSeriesAroonOsc200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesAroonOsc200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesAroonOsc200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesAroonOsc200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


