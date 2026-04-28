# GetTimeSeriesSuperTrend200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesSuperTrend200ResponseMeta**](GetTimeSeriesSuperTrend200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesSuperTrend200ResponseValuesInner**](GetTimeSeriesSuperTrend200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesSuperTrend200Response

`func NewGetTimeSeriesSuperTrend200Response(meta GetTimeSeriesSuperTrend200ResponseMeta, values []GetTimeSeriesSuperTrend200ResponseValuesInner, status string, ) *GetTimeSeriesSuperTrend200Response`

NewGetTimeSeriesSuperTrend200Response instantiates a new GetTimeSeriesSuperTrend200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesSuperTrend200ResponseWithDefaults

`func NewGetTimeSeriesSuperTrend200ResponseWithDefaults() *GetTimeSeriesSuperTrend200Response`

NewGetTimeSeriesSuperTrend200ResponseWithDefaults instantiates a new GetTimeSeriesSuperTrend200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesSuperTrend200Response) GetMeta() GetTimeSeriesSuperTrend200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesSuperTrend200Response) GetMetaOk() (*GetTimeSeriesSuperTrend200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesSuperTrend200Response) SetMeta(v GetTimeSeriesSuperTrend200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesSuperTrend200Response) GetValues() []GetTimeSeriesSuperTrend200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesSuperTrend200Response) GetValuesOk() (*[]GetTimeSeriesSuperTrend200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesSuperTrend200Response) SetValues(v []GetTimeSeriesSuperTrend200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesSuperTrend200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesSuperTrend200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesSuperTrend200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


