# GetTimeSeriesMama200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesMama200ResponseMeta**](GetTimeSeriesMama200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesMama200ResponseValuesInner**](GetTimeSeriesMama200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesMama200Response

`func NewGetTimeSeriesMama200Response(meta GetTimeSeriesMama200ResponseMeta, values []GetTimeSeriesMama200ResponseValuesInner, status string, ) *GetTimeSeriesMama200Response`

NewGetTimeSeriesMama200Response instantiates a new GetTimeSeriesMama200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesMama200ResponseWithDefaults

`func NewGetTimeSeriesMama200ResponseWithDefaults() *GetTimeSeriesMama200Response`

NewGetTimeSeriesMama200ResponseWithDefaults instantiates a new GetTimeSeriesMama200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesMama200Response) GetMeta() GetTimeSeriesMama200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesMama200Response) GetMetaOk() (*GetTimeSeriesMama200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesMama200Response) SetMeta(v GetTimeSeriesMama200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesMama200Response) GetValues() []GetTimeSeriesMama200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesMama200Response) GetValuesOk() (*[]GetTimeSeriesMama200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesMama200Response) SetValues(v []GetTimeSeriesMama200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesMama200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesMama200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesMama200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


