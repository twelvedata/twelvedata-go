# GetTimeSeriesDpo200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesDpo200ResponseMeta**](GetTimeSeriesDpo200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesDpo200ResponseValuesInner**](GetTimeSeriesDpo200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesDpo200Response

`func NewGetTimeSeriesDpo200Response(meta GetTimeSeriesDpo200ResponseMeta, values []GetTimeSeriesDpo200ResponseValuesInner, status string, ) *GetTimeSeriesDpo200Response`

NewGetTimeSeriesDpo200Response instantiates a new GetTimeSeriesDpo200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesDpo200ResponseWithDefaults

`func NewGetTimeSeriesDpo200ResponseWithDefaults() *GetTimeSeriesDpo200Response`

NewGetTimeSeriesDpo200ResponseWithDefaults instantiates a new GetTimeSeriesDpo200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesDpo200Response) GetMeta() GetTimeSeriesDpo200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesDpo200Response) GetMetaOk() (*GetTimeSeriesDpo200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesDpo200Response) SetMeta(v GetTimeSeriesDpo200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesDpo200Response) GetValues() []GetTimeSeriesDpo200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesDpo200Response) GetValuesOk() (*[]GetTimeSeriesDpo200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesDpo200Response) SetValues(v []GetTimeSeriesDpo200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesDpo200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesDpo200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesDpo200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


