# GetTimeSeriesTema200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesTema200ResponseMeta**](GetTimeSeriesTema200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesTema200ResponseValuesInner**](GetTimeSeriesTema200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesTema200Response

`func NewGetTimeSeriesTema200Response(meta GetTimeSeriesTema200ResponseMeta, values []GetTimeSeriesTema200ResponseValuesInner, status string, ) *GetTimeSeriesTema200Response`

NewGetTimeSeriesTema200Response instantiates a new GetTimeSeriesTema200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesTema200ResponseWithDefaults

`func NewGetTimeSeriesTema200ResponseWithDefaults() *GetTimeSeriesTema200Response`

NewGetTimeSeriesTema200ResponseWithDefaults instantiates a new GetTimeSeriesTema200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesTema200Response) GetMeta() GetTimeSeriesTema200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesTema200Response) GetMetaOk() (*GetTimeSeriesTema200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesTema200Response) SetMeta(v GetTimeSeriesTema200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesTema200Response) GetValues() []GetTimeSeriesTema200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesTema200Response) GetValuesOk() (*[]GetTimeSeriesTema200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesTema200Response) SetValues(v []GetTimeSeriesTema200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesTema200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesTema200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesTema200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


