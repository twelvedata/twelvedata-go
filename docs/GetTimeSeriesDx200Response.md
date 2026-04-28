# GetTimeSeriesDx200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesDx200ResponseMeta**](GetTimeSeriesDx200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesDx200ResponseValuesInner**](GetTimeSeriesDx200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesDx200Response

`func NewGetTimeSeriesDx200Response(meta GetTimeSeriesDx200ResponseMeta, values []GetTimeSeriesDx200ResponseValuesInner, status string, ) *GetTimeSeriesDx200Response`

NewGetTimeSeriesDx200Response instantiates a new GetTimeSeriesDx200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesDx200ResponseWithDefaults

`func NewGetTimeSeriesDx200ResponseWithDefaults() *GetTimeSeriesDx200Response`

NewGetTimeSeriesDx200ResponseWithDefaults instantiates a new GetTimeSeriesDx200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesDx200Response) GetMeta() GetTimeSeriesDx200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesDx200Response) GetMetaOk() (*GetTimeSeriesDx200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesDx200Response) SetMeta(v GetTimeSeriesDx200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesDx200Response) GetValues() []GetTimeSeriesDx200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesDx200Response) GetValuesOk() (*[]GetTimeSeriesDx200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesDx200Response) SetValues(v []GetTimeSeriesDx200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesDx200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesDx200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesDx200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


