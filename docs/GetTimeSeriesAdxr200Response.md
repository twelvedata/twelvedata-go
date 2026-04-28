# GetTimeSeriesAdxr200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesAdxr200ResponseMeta**](GetTimeSeriesAdxr200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesAdxr200ResponseValuesInner**](GetTimeSeriesAdxr200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesAdxr200Response

`func NewGetTimeSeriesAdxr200Response(meta GetTimeSeriesAdxr200ResponseMeta, values []GetTimeSeriesAdxr200ResponseValuesInner, status string, ) *GetTimeSeriesAdxr200Response`

NewGetTimeSeriesAdxr200Response instantiates a new GetTimeSeriesAdxr200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesAdxr200ResponseWithDefaults

`func NewGetTimeSeriesAdxr200ResponseWithDefaults() *GetTimeSeriesAdxr200Response`

NewGetTimeSeriesAdxr200ResponseWithDefaults instantiates a new GetTimeSeriesAdxr200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesAdxr200Response) GetMeta() GetTimeSeriesAdxr200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesAdxr200Response) GetMetaOk() (*GetTimeSeriesAdxr200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesAdxr200Response) SetMeta(v GetTimeSeriesAdxr200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesAdxr200Response) GetValues() []GetTimeSeriesAdxr200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesAdxr200Response) GetValuesOk() (*[]GetTimeSeriesAdxr200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesAdxr200Response) SetValues(v []GetTimeSeriesAdxr200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesAdxr200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesAdxr200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesAdxr200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


