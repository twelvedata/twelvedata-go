# GetTimeSeriesDema200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesDema200ResponseMeta**](GetTimeSeriesDema200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesDema200ResponseValuesInner**](GetTimeSeriesDema200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesDema200Response

`func NewGetTimeSeriesDema200Response(meta GetTimeSeriesDema200ResponseMeta, values []GetTimeSeriesDema200ResponseValuesInner, status string, ) *GetTimeSeriesDema200Response`

NewGetTimeSeriesDema200Response instantiates a new GetTimeSeriesDema200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesDema200ResponseWithDefaults

`func NewGetTimeSeriesDema200ResponseWithDefaults() *GetTimeSeriesDema200Response`

NewGetTimeSeriesDema200ResponseWithDefaults instantiates a new GetTimeSeriesDema200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesDema200Response) GetMeta() GetTimeSeriesDema200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesDema200Response) GetMetaOk() (*GetTimeSeriesDema200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesDema200Response) SetMeta(v GetTimeSeriesDema200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesDema200Response) GetValues() []GetTimeSeriesDema200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesDema200Response) GetValuesOk() (*[]GetTimeSeriesDema200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesDema200Response) SetValues(v []GetTimeSeriesDema200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesDema200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesDema200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesDema200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


