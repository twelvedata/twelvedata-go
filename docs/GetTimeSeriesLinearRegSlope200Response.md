# GetTimeSeriesLinearRegSlope200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesLinearRegSlope200ResponseMeta**](GetTimeSeriesLinearRegSlope200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesLinearRegSlope200ResponseValuesInner**](GetTimeSeriesLinearRegSlope200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesLinearRegSlope200Response

`func NewGetTimeSeriesLinearRegSlope200Response(meta GetTimeSeriesLinearRegSlope200ResponseMeta, values []GetTimeSeriesLinearRegSlope200ResponseValuesInner, status string, ) *GetTimeSeriesLinearRegSlope200Response`

NewGetTimeSeriesLinearRegSlope200Response instantiates a new GetTimeSeriesLinearRegSlope200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesLinearRegSlope200ResponseWithDefaults

`func NewGetTimeSeriesLinearRegSlope200ResponseWithDefaults() *GetTimeSeriesLinearRegSlope200Response`

NewGetTimeSeriesLinearRegSlope200ResponseWithDefaults instantiates a new GetTimeSeriesLinearRegSlope200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesLinearRegSlope200Response) GetMeta() GetTimeSeriesLinearRegSlope200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesLinearRegSlope200Response) GetMetaOk() (*GetTimeSeriesLinearRegSlope200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesLinearRegSlope200Response) SetMeta(v GetTimeSeriesLinearRegSlope200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesLinearRegSlope200Response) GetValues() []GetTimeSeriesLinearRegSlope200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesLinearRegSlope200Response) GetValuesOk() (*[]GetTimeSeriesLinearRegSlope200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesLinearRegSlope200Response) SetValues(v []GetTimeSeriesLinearRegSlope200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesLinearRegSlope200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesLinearRegSlope200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesLinearRegSlope200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


