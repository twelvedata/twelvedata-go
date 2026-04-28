# GetTimeSeriesSqrt200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesSqrt200ResponseMeta**](GetTimeSeriesSqrt200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesSqrt200ResponseValuesInner**](GetTimeSeriesSqrt200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesSqrt200Response

`func NewGetTimeSeriesSqrt200Response(meta GetTimeSeriesSqrt200ResponseMeta, values []GetTimeSeriesSqrt200ResponseValuesInner, status string, ) *GetTimeSeriesSqrt200Response`

NewGetTimeSeriesSqrt200Response instantiates a new GetTimeSeriesSqrt200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesSqrt200ResponseWithDefaults

`func NewGetTimeSeriesSqrt200ResponseWithDefaults() *GetTimeSeriesSqrt200Response`

NewGetTimeSeriesSqrt200ResponseWithDefaults instantiates a new GetTimeSeriesSqrt200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesSqrt200Response) GetMeta() GetTimeSeriesSqrt200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesSqrt200Response) GetMetaOk() (*GetTimeSeriesSqrt200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesSqrt200Response) SetMeta(v GetTimeSeriesSqrt200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesSqrt200Response) GetValues() []GetTimeSeriesSqrt200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesSqrt200Response) GetValuesOk() (*[]GetTimeSeriesSqrt200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesSqrt200Response) SetValues(v []GetTimeSeriesSqrt200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesSqrt200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesSqrt200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesSqrt200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


