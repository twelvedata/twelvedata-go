# GetTimeSeriesExp200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesExp200ResponseMeta**](GetTimeSeriesExp200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesExp200ResponseValuesInner**](GetTimeSeriesExp200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesExp200Response

`func NewGetTimeSeriesExp200Response(meta GetTimeSeriesExp200ResponseMeta, values []GetTimeSeriesExp200ResponseValuesInner, status string, ) *GetTimeSeriesExp200Response`

NewGetTimeSeriesExp200Response instantiates a new GetTimeSeriesExp200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesExp200ResponseWithDefaults

`func NewGetTimeSeriesExp200ResponseWithDefaults() *GetTimeSeriesExp200Response`

NewGetTimeSeriesExp200ResponseWithDefaults instantiates a new GetTimeSeriesExp200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesExp200Response) GetMeta() GetTimeSeriesExp200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesExp200Response) GetMetaOk() (*GetTimeSeriesExp200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesExp200Response) SetMeta(v GetTimeSeriesExp200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesExp200Response) GetValues() []GetTimeSeriesExp200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesExp200Response) GetValuesOk() (*[]GetTimeSeriesExp200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesExp200Response) SetValues(v []GetTimeSeriesExp200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesExp200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesExp200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesExp200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


