# GetTimeSeriesLinearRegIntercept200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesLinearRegIntercept200ResponseMeta**](GetTimeSeriesLinearRegIntercept200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesLinearRegIntercept200ResponseValuesInner**](GetTimeSeriesLinearRegIntercept200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesLinearRegIntercept200Response

`func NewGetTimeSeriesLinearRegIntercept200Response(meta GetTimeSeriesLinearRegIntercept200ResponseMeta, values []GetTimeSeriesLinearRegIntercept200ResponseValuesInner, status string, ) *GetTimeSeriesLinearRegIntercept200Response`

NewGetTimeSeriesLinearRegIntercept200Response instantiates a new GetTimeSeriesLinearRegIntercept200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesLinearRegIntercept200ResponseWithDefaults

`func NewGetTimeSeriesLinearRegIntercept200ResponseWithDefaults() *GetTimeSeriesLinearRegIntercept200Response`

NewGetTimeSeriesLinearRegIntercept200ResponseWithDefaults instantiates a new GetTimeSeriesLinearRegIntercept200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesLinearRegIntercept200Response) GetMeta() GetTimeSeriesLinearRegIntercept200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesLinearRegIntercept200Response) GetMetaOk() (*GetTimeSeriesLinearRegIntercept200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesLinearRegIntercept200Response) SetMeta(v GetTimeSeriesLinearRegIntercept200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesLinearRegIntercept200Response) GetValues() []GetTimeSeriesLinearRegIntercept200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesLinearRegIntercept200Response) GetValuesOk() (*[]GetTimeSeriesLinearRegIntercept200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesLinearRegIntercept200Response) SetValues(v []GetTimeSeriesLinearRegIntercept200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesLinearRegIntercept200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesLinearRegIntercept200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesLinearRegIntercept200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


