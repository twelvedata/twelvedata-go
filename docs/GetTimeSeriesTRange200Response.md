# GetTimeSeriesTRange200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesTRange200ResponseMeta**](GetTimeSeriesTRange200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesTRange200ResponseValuesInner**](GetTimeSeriesTRange200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesTRange200Response

`func NewGetTimeSeriesTRange200Response(meta GetTimeSeriesTRange200ResponseMeta, values []GetTimeSeriesTRange200ResponseValuesInner, status string, ) *GetTimeSeriesTRange200Response`

NewGetTimeSeriesTRange200Response instantiates a new GetTimeSeriesTRange200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesTRange200ResponseWithDefaults

`func NewGetTimeSeriesTRange200ResponseWithDefaults() *GetTimeSeriesTRange200Response`

NewGetTimeSeriesTRange200ResponseWithDefaults instantiates a new GetTimeSeriesTRange200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesTRange200Response) GetMeta() GetTimeSeriesTRange200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesTRange200Response) GetMetaOk() (*GetTimeSeriesTRange200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesTRange200Response) SetMeta(v GetTimeSeriesTRange200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesTRange200Response) GetValues() []GetTimeSeriesTRange200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesTRange200Response) GetValuesOk() (*[]GetTimeSeriesTRange200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesTRange200Response) SetValues(v []GetTimeSeriesTRange200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesTRange200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesTRange200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesTRange200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


