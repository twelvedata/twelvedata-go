# GetTimeSeriesVar200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesVar200ResponseMeta**](GetTimeSeriesVar200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesVar200ResponseValuesInner**](GetTimeSeriesVar200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesVar200Response

`func NewGetTimeSeriesVar200Response(meta GetTimeSeriesVar200ResponseMeta, values []GetTimeSeriesVar200ResponseValuesInner, status string, ) *GetTimeSeriesVar200Response`

NewGetTimeSeriesVar200Response instantiates a new GetTimeSeriesVar200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesVar200ResponseWithDefaults

`func NewGetTimeSeriesVar200ResponseWithDefaults() *GetTimeSeriesVar200Response`

NewGetTimeSeriesVar200ResponseWithDefaults instantiates a new GetTimeSeriesVar200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesVar200Response) GetMeta() GetTimeSeriesVar200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesVar200Response) GetMetaOk() (*GetTimeSeriesVar200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesVar200Response) SetMeta(v GetTimeSeriesVar200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesVar200Response) GetValues() []GetTimeSeriesVar200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesVar200Response) GetValuesOk() (*[]GetTimeSeriesVar200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesVar200Response) SetValues(v []GetTimeSeriesVar200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesVar200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesVar200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesVar200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


