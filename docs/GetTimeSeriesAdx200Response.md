# GetTimeSeriesAdx200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesAdx200ResponseMeta**](GetTimeSeriesAdx200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesAdx200ResponseValuesInner**](GetTimeSeriesAdx200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesAdx200Response

`func NewGetTimeSeriesAdx200Response(meta GetTimeSeriesAdx200ResponseMeta, values []GetTimeSeriesAdx200ResponseValuesInner, status string, ) *GetTimeSeriesAdx200Response`

NewGetTimeSeriesAdx200Response instantiates a new GetTimeSeriesAdx200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesAdx200ResponseWithDefaults

`func NewGetTimeSeriesAdx200ResponseWithDefaults() *GetTimeSeriesAdx200Response`

NewGetTimeSeriesAdx200ResponseWithDefaults instantiates a new GetTimeSeriesAdx200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesAdx200Response) GetMeta() GetTimeSeriesAdx200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesAdx200Response) GetMetaOk() (*GetTimeSeriesAdx200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesAdx200Response) SetMeta(v GetTimeSeriesAdx200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesAdx200Response) GetValues() []GetTimeSeriesAdx200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesAdx200Response) GetValuesOk() (*[]GetTimeSeriesAdx200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesAdx200Response) SetValues(v []GetTimeSeriesAdx200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesAdx200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesAdx200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesAdx200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


