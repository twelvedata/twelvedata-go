# GetTimeSeriesKama200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesKama200ResponseMeta**](GetTimeSeriesKama200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesKama200ResponseValuesInner**](GetTimeSeriesKama200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesKama200Response

`func NewGetTimeSeriesKama200Response(meta GetTimeSeriesKama200ResponseMeta, values []GetTimeSeriesKama200ResponseValuesInner, status string, ) *GetTimeSeriesKama200Response`

NewGetTimeSeriesKama200Response instantiates a new GetTimeSeriesKama200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesKama200ResponseWithDefaults

`func NewGetTimeSeriesKama200ResponseWithDefaults() *GetTimeSeriesKama200Response`

NewGetTimeSeriesKama200ResponseWithDefaults instantiates a new GetTimeSeriesKama200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesKama200Response) GetMeta() GetTimeSeriesKama200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesKama200Response) GetMetaOk() (*GetTimeSeriesKama200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesKama200Response) SetMeta(v GetTimeSeriesKama200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesKama200Response) GetValues() []GetTimeSeriesKama200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesKama200Response) GetValuesOk() (*[]GetTimeSeriesKama200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesKama200Response) SetValues(v []GetTimeSeriesKama200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesKama200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesKama200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesKama200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


