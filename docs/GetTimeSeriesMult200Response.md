# GetTimeSeriesMult200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesMult200ResponseMeta**](GetTimeSeriesMult200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesMult200ResponseValuesInner**](GetTimeSeriesMult200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesMult200Response

`func NewGetTimeSeriesMult200Response(meta GetTimeSeriesMult200ResponseMeta, values []GetTimeSeriesMult200ResponseValuesInner, status string, ) *GetTimeSeriesMult200Response`

NewGetTimeSeriesMult200Response instantiates a new GetTimeSeriesMult200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesMult200ResponseWithDefaults

`func NewGetTimeSeriesMult200ResponseWithDefaults() *GetTimeSeriesMult200Response`

NewGetTimeSeriesMult200ResponseWithDefaults instantiates a new GetTimeSeriesMult200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesMult200Response) GetMeta() GetTimeSeriesMult200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesMult200Response) GetMetaOk() (*GetTimeSeriesMult200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesMult200Response) SetMeta(v GetTimeSeriesMult200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesMult200Response) GetValues() []GetTimeSeriesMult200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesMult200Response) GetValuesOk() (*[]GetTimeSeriesMult200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesMult200Response) SetValues(v []GetTimeSeriesMult200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesMult200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesMult200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesMult200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


