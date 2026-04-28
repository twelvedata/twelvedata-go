# GetTimeSeriesMom200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesMom200ResponseMeta**](GetTimeSeriesMom200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesMom200ResponseValuesInner**](GetTimeSeriesMom200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesMom200Response

`func NewGetTimeSeriesMom200Response(meta GetTimeSeriesMom200ResponseMeta, values []GetTimeSeriesMom200ResponseValuesInner, status string, ) *GetTimeSeriesMom200Response`

NewGetTimeSeriesMom200Response instantiates a new GetTimeSeriesMom200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesMom200ResponseWithDefaults

`func NewGetTimeSeriesMom200ResponseWithDefaults() *GetTimeSeriesMom200Response`

NewGetTimeSeriesMom200ResponseWithDefaults instantiates a new GetTimeSeriesMom200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesMom200Response) GetMeta() GetTimeSeriesMom200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesMom200Response) GetMetaOk() (*GetTimeSeriesMom200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesMom200Response) SetMeta(v GetTimeSeriesMom200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesMom200Response) GetValues() []GetTimeSeriesMom200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesMom200Response) GetValuesOk() (*[]GetTimeSeriesMom200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesMom200Response) SetValues(v []GetTimeSeriesMom200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesMom200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesMom200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesMom200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


