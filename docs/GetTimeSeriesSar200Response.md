# GetTimeSeriesSar200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesSar200ResponseMeta**](GetTimeSeriesSar200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesSar200ResponseValuesInner**](GetTimeSeriesSar200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesSar200Response

`func NewGetTimeSeriesSar200Response(meta GetTimeSeriesSar200ResponseMeta, values []GetTimeSeriesSar200ResponseValuesInner, status string, ) *GetTimeSeriesSar200Response`

NewGetTimeSeriesSar200Response instantiates a new GetTimeSeriesSar200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesSar200ResponseWithDefaults

`func NewGetTimeSeriesSar200ResponseWithDefaults() *GetTimeSeriesSar200Response`

NewGetTimeSeriesSar200ResponseWithDefaults instantiates a new GetTimeSeriesSar200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesSar200Response) GetMeta() GetTimeSeriesSar200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesSar200Response) GetMetaOk() (*GetTimeSeriesSar200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesSar200Response) SetMeta(v GetTimeSeriesSar200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesSar200Response) GetValues() []GetTimeSeriesSar200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesSar200Response) GetValuesOk() (*[]GetTimeSeriesSar200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesSar200Response) SetValues(v []GetTimeSeriesSar200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesSar200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesSar200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesSar200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


