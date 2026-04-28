# GetTimeSeriesTypPrice200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesTypPrice200ResponseMeta**](GetTimeSeriesTypPrice200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesTypPrice200ResponseValuesInner**](GetTimeSeriesTypPrice200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesTypPrice200Response

`func NewGetTimeSeriesTypPrice200Response(meta GetTimeSeriesTypPrice200ResponseMeta, values []GetTimeSeriesTypPrice200ResponseValuesInner, status string, ) *GetTimeSeriesTypPrice200Response`

NewGetTimeSeriesTypPrice200Response instantiates a new GetTimeSeriesTypPrice200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesTypPrice200ResponseWithDefaults

`func NewGetTimeSeriesTypPrice200ResponseWithDefaults() *GetTimeSeriesTypPrice200Response`

NewGetTimeSeriesTypPrice200ResponseWithDefaults instantiates a new GetTimeSeriesTypPrice200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesTypPrice200Response) GetMeta() GetTimeSeriesTypPrice200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesTypPrice200Response) GetMetaOk() (*GetTimeSeriesTypPrice200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesTypPrice200Response) SetMeta(v GetTimeSeriesTypPrice200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesTypPrice200Response) GetValues() []GetTimeSeriesTypPrice200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesTypPrice200Response) GetValuesOk() (*[]GetTimeSeriesTypPrice200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesTypPrice200Response) SetValues(v []GetTimeSeriesTypPrice200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesTypPrice200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesTypPrice200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesTypPrice200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


