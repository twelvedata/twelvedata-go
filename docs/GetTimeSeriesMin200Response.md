# GetTimeSeriesMin200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesMin200ResponseMeta**](GetTimeSeriesMin200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesMin200ResponseValuesInner**](GetTimeSeriesMin200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesMin200Response

`func NewGetTimeSeriesMin200Response(meta GetTimeSeriesMin200ResponseMeta, values []GetTimeSeriesMin200ResponseValuesInner, status string, ) *GetTimeSeriesMin200Response`

NewGetTimeSeriesMin200Response instantiates a new GetTimeSeriesMin200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesMin200ResponseWithDefaults

`func NewGetTimeSeriesMin200ResponseWithDefaults() *GetTimeSeriesMin200Response`

NewGetTimeSeriesMin200ResponseWithDefaults instantiates a new GetTimeSeriesMin200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesMin200Response) GetMeta() GetTimeSeriesMin200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesMin200Response) GetMetaOk() (*GetTimeSeriesMin200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesMin200Response) SetMeta(v GetTimeSeriesMin200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesMin200Response) GetValues() []GetTimeSeriesMin200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesMin200Response) GetValuesOk() (*[]GetTimeSeriesMin200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesMin200Response) SetValues(v []GetTimeSeriesMin200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesMin200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesMin200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesMin200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


