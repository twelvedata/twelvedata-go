# GetTimeSeriesMacdExt200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesMacdExt200ResponseMeta**](GetTimeSeriesMacdExt200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesMacdExt200ResponseValuesInner**](GetTimeSeriesMacdExt200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesMacdExt200Response

`func NewGetTimeSeriesMacdExt200Response(meta GetTimeSeriesMacdExt200ResponseMeta, values []GetTimeSeriesMacdExt200ResponseValuesInner, status string, ) *GetTimeSeriesMacdExt200Response`

NewGetTimeSeriesMacdExt200Response instantiates a new GetTimeSeriesMacdExt200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesMacdExt200ResponseWithDefaults

`func NewGetTimeSeriesMacdExt200ResponseWithDefaults() *GetTimeSeriesMacdExt200Response`

NewGetTimeSeriesMacdExt200ResponseWithDefaults instantiates a new GetTimeSeriesMacdExt200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesMacdExt200Response) GetMeta() GetTimeSeriesMacdExt200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesMacdExt200Response) GetMetaOk() (*GetTimeSeriesMacdExt200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesMacdExt200Response) SetMeta(v GetTimeSeriesMacdExt200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesMacdExt200Response) GetValues() []GetTimeSeriesMacdExt200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesMacdExt200Response) GetValuesOk() (*[]GetTimeSeriesMacdExt200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesMacdExt200Response) SetValues(v []GetTimeSeriesMacdExt200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesMacdExt200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesMacdExt200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesMacdExt200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


