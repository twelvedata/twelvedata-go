# GetTimeSeriesAtr200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesAtr200ResponseMeta**](GetTimeSeriesAtr200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesAtr200ResponseValuesInner**](GetTimeSeriesAtr200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesAtr200Response

`func NewGetTimeSeriesAtr200Response(meta GetTimeSeriesAtr200ResponseMeta, values []GetTimeSeriesAtr200ResponseValuesInner, status string, ) *GetTimeSeriesAtr200Response`

NewGetTimeSeriesAtr200Response instantiates a new GetTimeSeriesAtr200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesAtr200ResponseWithDefaults

`func NewGetTimeSeriesAtr200ResponseWithDefaults() *GetTimeSeriesAtr200Response`

NewGetTimeSeriesAtr200ResponseWithDefaults instantiates a new GetTimeSeriesAtr200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesAtr200Response) GetMeta() GetTimeSeriesAtr200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesAtr200Response) GetMetaOk() (*GetTimeSeriesAtr200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesAtr200Response) SetMeta(v GetTimeSeriesAtr200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesAtr200Response) GetValues() []GetTimeSeriesAtr200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesAtr200Response) GetValuesOk() (*[]GetTimeSeriesAtr200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesAtr200Response) SetValues(v []GetTimeSeriesAtr200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesAtr200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesAtr200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesAtr200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


