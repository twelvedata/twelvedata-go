# GetTimeSeriesVwap200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesVwap200ResponseMeta**](GetTimeSeriesVwap200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesVwap200ResponseValuesInner**](GetTimeSeriesVwap200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesVwap200Response

`func NewGetTimeSeriesVwap200Response(meta GetTimeSeriesVwap200ResponseMeta, values []GetTimeSeriesVwap200ResponseValuesInner, status string, ) *GetTimeSeriesVwap200Response`

NewGetTimeSeriesVwap200Response instantiates a new GetTimeSeriesVwap200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesVwap200ResponseWithDefaults

`func NewGetTimeSeriesVwap200ResponseWithDefaults() *GetTimeSeriesVwap200Response`

NewGetTimeSeriesVwap200ResponseWithDefaults instantiates a new GetTimeSeriesVwap200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesVwap200Response) GetMeta() GetTimeSeriesVwap200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesVwap200Response) GetMetaOk() (*GetTimeSeriesVwap200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesVwap200Response) SetMeta(v GetTimeSeriesVwap200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesVwap200Response) GetValues() []GetTimeSeriesVwap200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesVwap200Response) GetValuesOk() (*[]GetTimeSeriesVwap200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesVwap200Response) SetValues(v []GetTimeSeriesVwap200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesVwap200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesVwap200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesVwap200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


