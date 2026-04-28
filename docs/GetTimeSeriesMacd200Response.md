# GetTimeSeriesMacd200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesMacd200ResponseMeta**](GetTimeSeriesMacd200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesMacd200ResponseValuesInner**](GetTimeSeriesMacd200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesMacd200Response

`func NewGetTimeSeriesMacd200Response(meta GetTimeSeriesMacd200ResponseMeta, values []GetTimeSeriesMacd200ResponseValuesInner, status string, ) *GetTimeSeriesMacd200Response`

NewGetTimeSeriesMacd200Response instantiates a new GetTimeSeriesMacd200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesMacd200ResponseWithDefaults

`func NewGetTimeSeriesMacd200ResponseWithDefaults() *GetTimeSeriesMacd200Response`

NewGetTimeSeriesMacd200ResponseWithDefaults instantiates a new GetTimeSeriesMacd200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesMacd200Response) GetMeta() GetTimeSeriesMacd200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesMacd200Response) GetMetaOk() (*GetTimeSeriesMacd200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesMacd200Response) SetMeta(v GetTimeSeriesMacd200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesMacd200Response) GetValues() []GetTimeSeriesMacd200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesMacd200Response) GetValuesOk() (*[]GetTimeSeriesMacd200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesMacd200Response) SetValues(v []GetTimeSeriesMacd200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesMacd200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesMacd200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesMacd200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


