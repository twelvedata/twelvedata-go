# GetTimeSeriesRocp200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesRocp200ResponseMeta**](GetTimeSeriesRocp200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesRocp200ResponseValuesInner**](GetTimeSeriesRocp200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesRocp200Response

`func NewGetTimeSeriesRocp200Response(meta GetTimeSeriesRocp200ResponseMeta, values []GetTimeSeriesRocp200ResponseValuesInner, status string, ) *GetTimeSeriesRocp200Response`

NewGetTimeSeriesRocp200Response instantiates a new GetTimeSeriesRocp200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesRocp200ResponseWithDefaults

`func NewGetTimeSeriesRocp200ResponseWithDefaults() *GetTimeSeriesRocp200Response`

NewGetTimeSeriesRocp200ResponseWithDefaults instantiates a new GetTimeSeriesRocp200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesRocp200Response) GetMeta() GetTimeSeriesRocp200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesRocp200Response) GetMetaOk() (*GetTimeSeriesRocp200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesRocp200Response) SetMeta(v GetTimeSeriesRocp200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesRocp200Response) GetValues() []GetTimeSeriesRocp200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesRocp200Response) GetValuesOk() (*[]GetTimeSeriesRocp200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesRocp200Response) SetValues(v []GetTimeSeriesRocp200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesRocp200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesRocp200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesRocp200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


