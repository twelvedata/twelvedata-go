# GetTimeSeriesMidPrice200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesMidPrice200ResponseMeta**](GetTimeSeriesMidPrice200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesMidPrice200ResponseValuesInner**](GetTimeSeriesMidPrice200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesMidPrice200Response

`func NewGetTimeSeriesMidPrice200Response(meta GetTimeSeriesMidPrice200ResponseMeta, values []GetTimeSeriesMidPrice200ResponseValuesInner, status string, ) *GetTimeSeriesMidPrice200Response`

NewGetTimeSeriesMidPrice200Response instantiates a new GetTimeSeriesMidPrice200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesMidPrice200ResponseWithDefaults

`func NewGetTimeSeriesMidPrice200ResponseWithDefaults() *GetTimeSeriesMidPrice200Response`

NewGetTimeSeriesMidPrice200ResponseWithDefaults instantiates a new GetTimeSeriesMidPrice200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesMidPrice200Response) GetMeta() GetTimeSeriesMidPrice200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesMidPrice200Response) GetMetaOk() (*GetTimeSeriesMidPrice200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesMidPrice200Response) SetMeta(v GetTimeSeriesMidPrice200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesMidPrice200Response) GetValues() []GetTimeSeriesMidPrice200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesMidPrice200Response) GetValuesOk() (*[]GetTimeSeriesMidPrice200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesMidPrice200Response) SetValues(v []GetTimeSeriesMidPrice200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesMidPrice200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesMidPrice200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesMidPrice200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


