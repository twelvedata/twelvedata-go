# GetTimeSeriesMa200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesMa200ResponseMeta**](GetTimeSeriesMa200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesMa200ResponseValuesInner**](GetTimeSeriesMa200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesMa200Response

`func NewGetTimeSeriesMa200Response(meta GetTimeSeriesMa200ResponseMeta, values []GetTimeSeriesMa200ResponseValuesInner, status string, ) *GetTimeSeriesMa200Response`

NewGetTimeSeriesMa200Response instantiates a new GetTimeSeriesMa200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesMa200ResponseWithDefaults

`func NewGetTimeSeriesMa200ResponseWithDefaults() *GetTimeSeriesMa200Response`

NewGetTimeSeriesMa200ResponseWithDefaults instantiates a new GetTimeSeriesMa200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesMa200Response) GetMeta() GetTimeSeriesMa200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesMa200Response) GetMetaOk() (*GetTimeSeriesMa200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesMa200Response) SetMeta(v GetTimeSeriesMa200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesMa200Response) GetValues() []GetTimeSeriesMa200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesMa200Response) GetValuesOk() (*[]GetTimeSeriesMa200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesMa200Response) SetValues(v []GetTimeSeriesMa200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesMa200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesMa200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesMa200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


