# GetTimeSeriesSub200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesSub200ResponseMeta**](GetTimeSeriesSub200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesSub200ResponseValuesInner**](GetTimeSeriesSub200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesSub200Response

`func NewGetTimeSeriesSub200Response(meta GetTimeSeriesSub200ResponseMeta, values []GetTimeSeriesSub200ResponseValuesInner, status string, ) *GetTimeSeriesSub200Response`

NewGetTimeSeriesSub200Response instantiates a new GetTimeSeriesSub200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesSub200ResponseWithDefaults

`func NewGetTimeSeriesSub200ResponseWithDefaults() *GetTimeSeriesSub200Response`

NewGetTimeSeriesSub200ResponseWithDefaults instantiates a new GetTimeSeriesSub200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesSub200Response) GetMeta() GetTimeSeriesSub200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesSub200Response) GetMetaOk() (*GetTimeSeriesSub200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesSub200Response) SetMeta(v GetTimeSeriesSub200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesSub200Response) GetValues() []GetTimeSeriesSub200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesSub200Response) GetValuesOk() (*[]GetTimeSeriesSub200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesSub200Response) SetValues(v []GetTimeSeriesSub200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesSub200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesSub200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesSub200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


