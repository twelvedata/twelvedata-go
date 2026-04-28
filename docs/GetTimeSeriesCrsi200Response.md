# GetTimeSeriesCrsi200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesCrsi200ResponseMeta**](GetTimeSeriesCrsi200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesCrsi200ResponseValuesInner**](GetTimeSeriesCrsi200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesCrsi200Response

`func NewGetTimeSeriesCrsi200Response(meta GetTimeSeriesCrsi200ResponseMeta, values []GetTimeSeriesCrsi200ResponseValuesInner, status string, ) *GetTimeSeriesCrsi200Response`

NewGetTimeSeriesCrsi200Response instantiates a new GetTimeSeriesCrsi200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesCrsi200ResponseWithDefaults

`func NewGetTimeSeriesCrsi200ResponseWithDefaults() *GetTimeSeriesCrsi200Response`

NewGetTimeSeriesCrsi200ResponseWithDefaults instantiates a new GetTimeSeriesCrsi200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesCrsi200Response) GetMeta() GetTimeSeriesCrsi200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesCrsi200Response) GetMetaOk() (*GetTimeSeriesCrsi200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesCrsi200Response) SetMeta(v GetTimeSeriesCrsi200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesCrsi200Response) GetValues() []GetTimeSeriesCrsi200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesCrsi200Response) GetValuesOk() (*[]GetTimeSeriesCrsi200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesCrsi200Response) SetValues(v []GetTimeSeriesCrsi200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesCrsi200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesCrsi200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesCrsi200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


