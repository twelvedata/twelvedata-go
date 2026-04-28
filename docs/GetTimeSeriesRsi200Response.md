# GetTimeSeriesRsi200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesRsi200ResponseMeta**](GetTimeSeriesRsi200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesRsi200ResponseValuesInner**](GetTimeSeriesRsi200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesRsi200Response

`func NewGetTimeSeriesRsi200Response(meta GetTimeSeriesRsi200ResponseMeta, values []GetTimeSeriesRsi200ResponseValuesInner, status string, ) *GetTimeSeriesRsi200Response`

NewGetTimeSeriesRsi200Response instantiates a new GetTimeSeriesRsi200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesRsi200ResponseWithDefaults

`func NewGetTimeSeriesRsi200ResponseWithDefaults() *GetTimeSeriesRsi200Response`

NewGetTimeSeriesRsi200ResponseWithDefaults instantiates a new GetTimeSeriesRsi200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesRsi200Response) GetMeta() GetTimeSeriesRsi200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesRsi200Response) GetMetaOk() (*GetTimeSeriesRsi200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesRsi200Response) SetMeta(v GetTimeSeriesRsi200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesRsi200Response) GetValues() []GetTimeSeriesRsi200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesRsi200Response) GetValuesOk() (*[]GetTimeSeriesRsi200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesRsi200Response) SetValues(v []GetTimeSeriesRsi200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesRsi200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesRsi200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesRsi200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


