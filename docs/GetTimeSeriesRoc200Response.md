# GetTimeSeriesRoc200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesRoc200ResponseMeta**](GetTimeSeriesRoc200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesRoc200ResponseValuesInner**](GetTimeSeriesRoc200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesRoc200Response

`func NewGetTimeSeriesRoc200Response(meta GetTimeSeriesRoc200ResponseMeta, values []GetTimeSeriesRoc200ResponseValuesInner, status string, ) *GetTimeSeriesRoc200Response`

NewGetTimeSeriesRoc200Response instantiates a new GetTimeSeriesRoc200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesRoc200ResponseWithDefaults

`func NewGetTimeSeriesRoc200ResponseWithDefaults() *GetTimeSeriesRoc200Response`

NewGetTimeSeriesRoc200ResponseWithDefaults instantiates a new GetTimeSeriesRoc200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesRoc200Response) GetMeta() GetTimeSeriesRoc200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesRoc200Response) GetMetaOk() (*GetTimeSeriesRoc200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesRoc200Response) SetMeta(v GetTimeSeriesRoc200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesRoc200Response) GetValues() []GetTimeSeriesRoc200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesRoc200Response) GetValuesOk() (*[]GetTimeSeriesRoc200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesRoc200Response) SetValues(v []GetTimeSeriesRoc200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesRoc200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesRoc200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesRoc200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


