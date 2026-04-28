# GetTimeSeriesAdd200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesAdd200ResponseMeta**](GetTimeSeriesAdd200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesAdd200ResponseValuesInner**](GetTimeSeriesAdd200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesAdd200Response

`func NewGetTimeSeriesAdd200Response(meta GetTimeSeriesAdd200ResponseMeta, values []GetTimeSeriesAdd200ResponseValuesInner, status string, ) *GetTimeSeriesAdd200Response`

NewGetTimeSeriesAdd200Response instantiates a new GetTimeSeriesAdd200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesAdd200ResponseWithDefaults

`func NewGetTimeSeriesAdd200ResponseWithDefaults() *GetTimeSeriesAdd200Response`

NewGetTimeSeriesAdd200ResponseWithDefaults instantiates a new GetTimeSeriesAdd200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesAdd200Response) GetMeta() GetTimeSeriesAdd200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesAdd200Response) GetMetaOk() (*GetTimeSeriesAdd200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesAdd200Response) SetMeta(v GetTimeSeriesAdd200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesAdd200Response) GetValues() []GetTimeSeriesAdd200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesAdd200Response) GetValuesOk() (*[]GetTimeSeriesAdd200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesAdd200Response) SetValues(v []GetTimeSeriesAdd200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesAdd200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesAdd200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesAdd200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


