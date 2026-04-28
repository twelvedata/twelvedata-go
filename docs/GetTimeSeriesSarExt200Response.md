# GetTimeSeriesSarExt200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesSarExt200ResponseMeta**](GetTimeSeriesSarExt200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesSarExt200ResponseValuesInner**](GetTimeSeriesSarExt200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesSarExt200Response

`func NewGetTimeSeriesSarExt200Response(meta GetTimeSeriesSarExt200ResponseMeta, values []GetTimeSeriesSarExt200ResponseValuesInner, status string, ) *GetTimeSeriesSarExt200Response`

NewGetTimeSeriesSarExt200Response instantiates a new GetTimeSeriesSarExt200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesSarExt200ResponseWithDefaults

`func NewGetTimeSeriesSarExt200ResponseWithDefaults() *GetTimeSeriesSarExt200Response`

NewGetTimeSeriesSarExt200ResponseWithDefaults instantiates a new GetTimeSeriesSarExt200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesSarExt200Response) GetMeta() GetTimeSeriesSarExt200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesSarExt200Response) GetMetaOk() (*GetTimeSeriesSarExt200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesSarExt200Response) SetMeta(v GetTimeSeriesSarExt200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesSarExt200Response) GetValues() []GetTimeSeriesSarExt200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesSarExt200Response) GetValuesOk() (*[]GetTimeSeriesSarExt200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesSarExt200Response) SetValues(v []GetTimeSeriesSarExt200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesSarExt200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesSarExt200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesSarExt200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


