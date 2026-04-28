# GetTimeSeriesEma200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesEma200ResponseMeta**](GetTimeSeriesEma200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesEma200ResponseValuesInner**](GetTimeSeriesEma200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesEma200Response

`func NewGetTimeSeriesEma200Response(meta GetTimeSeriesEma200ResponseMeta, values []GetTimeSeriesEma200ResponseValuesInner, status string, ) *GetTimeSeriesEma200Response`

NewGetTimeSeriesEma200Response instantiates a new GetTimeSeriesEma200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesEma200ResponseWithDefaults

`func NewGetTimeSeriesEma200ResponseWithDefaults() *GetTimeSeriesEma200Response`

NewGetTimeSeriesEma200ResponseWithDefaults instantiates a new GetTimeSeriesEma200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesEma200Response) GetMeta() GetTimeSeriesEma200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesEma200Response) GetMetaOk() (*GetTimeSeriesEma200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesEma200Response) SetMeta(v GetTimeSeriesEma200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesEma200Response) GetValues() []GetTimeSeriesEma200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesEma200Response) GetValuesOk() (*[]GetTimeSeriesEma200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesEma200Response) SetValues(v []GetTimeSeriesEma200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesEma200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesEma200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesEma200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


