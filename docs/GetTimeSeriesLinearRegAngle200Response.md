# GetTimeSeriesLinearRegAngle200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesLinearRegAngle200ResponseMeta**](GetTimeSeriesLinearRegAngle200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesLinearRegAngle200ResponseValuesInner**](GetTimeSeriesLinearRegAngle200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesLinearRegAngle200Response

`func NewGetTimeSeriesLinearRegAngle200Response(meta GetTimeSeriesLinearRegAngle200ResponseMeta, values []GetTimeSeriesLinearRegAngle200ResponseValuesInner, status string, ) *GetTimeSeriesLinearRegAngle200Response`

NewGetTimeSeriesLinearRegAngle200Response instantiates a new GetTimeSeriesLinearRegAngle200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesLinearRegAngle200ResponseWithDefaults

`func NewGetTimeSeriesLinearRegAngle200ResponseWithDefaults() *GetTimeSeriesLinearRegAngle200Response`

NewGetTimeSeriesLinearRegAngle200ResponseWithDefaults instantiates a new GetTimeSeriesLinearRegAngle200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesLinearRegAngle200Response) GetMeta() GetTimeSeriesLinearRegAngle200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesLinearRegAngle200Response) GetMetaOk() (*GetTimeSeriesLinearRegAngle200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesLinearRegAngle200Response) SetMeta(v GetTimeSeriesLinearRegAngle200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesLinearRegAngle200Response) GetValues() []GetTimeSeriesLinearRegAngle200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesLinearRegAngle200Response) GetValuesOk() (*[]GetTimeSeriesLinearRegAngle200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesLinearRegAngle200Response) SetValues(v []GetTimeSeriesLinearRegAngle200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesLinearRegAngle200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesLinearRegAngle200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesLinearRegAngle200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


