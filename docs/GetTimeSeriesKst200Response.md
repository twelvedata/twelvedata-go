# GetTimeSeriesKst200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesKst200ResponseMeta**](GetTimeSeriesKst200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesKst200ResponseValuesInner**](GetTimeSeriesKst200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesKst200Response

`func NewGetTimeSeriesKst200Response(meta GetTimeSeriesKst200ResponseMeta, values []GetTimeSeriesKst200ResponseValuesInner, status string, ) *GetTimeSeriesKst200Response`

NewGetTimeSeriesKst200Response instantiates a new GetTimeSeriesKst200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesKst200ResponseWithDefaults

`func NewGetTimeSeriesKst200ResponseWithDefaults() *GetTimeSeriesKst200Response`

NewGetTimeSeriesKst200ResponseWithDefaults instantiates a new GetTimeSeriesKst200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesKst200Response) GetMeta() GetTimeSeriesKst200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesKst200Response) GetMetaOk() (*GetTimeSeriesKst200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesKst200Response) SetMeta(v GetTimeSeriesKst200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesKst200Response) GetValues() []GetTimeSeriesKst200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesKst200Response) GetValuesOk() (*[]GetTimeSeriesKst200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesKst200Response) SetValues(v []GetTimeSeriesKst200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesKst200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesKst200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesKst200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


