# GetTimeSeriesMfi200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesMfi200ResponseMeta**](GetTimeSeriesMfi200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesMfi200ResponseValuesInner**](GetTimeSeriesMfi200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesMfi200Response

`func NewGetTimeSeriesMfi200Response(meta GetTimeSeriesMfi200ResponseMeta, values []GetTimeSeriesMfi200ResponseValuesInner, status string, ) *GetTimeSeriesMfi200Response`

NewGetTimeSeriesMfi200Response instantiates a new GetTimeSeriesMfi200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesMfi200ResponseWithDefaults

`func NewGetTimeSeriesMfi200ResponseWithDefaults() *GetTimeSeriesMfi200Response`

NewGetTimeSeriesMfi200ResponseWithDefaults instantiates a new GetTimeSeriesMfi200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesMfi200Response) GetMeta() GetTimeSeriesMfi200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesMfi200Response) GetMetaOk() (*GetTimeSeriesMfi200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesMfi200Response) SetMeta(v GetTimeSeriesMfi200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesMfi200Response) GetValues() []GetTimeSeriesMfi200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesMfi200Response) GetValuesOk() (*[]GetTimeSeriesMfi200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesMfi200Response) SetValues(v []GetTimeSeriesMfi200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesMfi200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesMfi200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesMfi200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


