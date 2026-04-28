# GetTimeSeriesLinearReg200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesLinearReg200ResponseMeta**](GetTimeSeriesLinearReg200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesLinearReg200ResponseValuesInner**](GetTimeSeriesLinearReg200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesLinearReg200Response

`func NewGetTimeSeriesLinearReg200Response(meta GetTimeSeriesLinearReg200ResponseMeta, values []GetTimeSeriesLinearReg200ResponseValuesInner, status string, ) *GetTimeSeriesLinearReg200Response`

NewGetTimeSeriesLinearReg200Response instantiates a new GetTimeSeriesLinearReg200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesLinearReg200ResponseWithDefaults

`func NewGetTimeSeriesLinearReg200ResponseWithDefaults() *GetTimeSeriesLinearReg200Response`

NewGetTimeSeriesLinearReg200ResponseWithDefaults instantiates a new GetTimeSeriesLinearReg200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesLinearReg200Response) GetMeta() GetTimeSeriesLinearReg200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesLinearReg200Response) GetMetaOk() (*GetTimeSeriesLinearReg200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesLinearReg200Response) SetMeta(v GetTimeSeriesLinearReg200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesLinearReg200Response) GetValues() []GetTimeSeriesLinearReg200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesLinearReg200Response) GetValuesOk() (*[]GetTimeSeriesLinearReg200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesLinearReg200Response) SetValues(v []GetTimeSeriesLinearReg200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesLinearReg200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesLinearReg200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesLinearReg200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


