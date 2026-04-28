# GetTimeSeriesT3ma200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesT3ma200ResponseMeta**](GetTimeSeriesT3ma200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesT3ma200ResponseValuesInner**](GetTimeSeriesT3ma200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesT3ma200Response

`func NewGetTimeSeriesT3ma200Response(meta GetTimeSeriesT3ma200ResponseMeta, values []GetTimeSeriesT3ma200ResponseValuesInner, status string, ) *GetTimeSeriesT3ma200Response`

NewGetTimeSeriesT3ma200Response instantiates a new GetTimeSeriesT3ma200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesT3ma200ResponseWithDefaults

`func NewGetTimeSeriesT3ma200ResponseWithDefaults() *GetTimeSeriesT3ma200Response`

NewGetTimeSeriesT3ma200ResponseWithDefaults instantiates a new GetTimeSeriesT3ma200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesT3ma200Response) GetMeta() GetTimeSeriesT3ma200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesT3ma200Response) GetMetaOk() (*GetTimeSeriesT3ma200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesT3ma200Response) SetMeta(v GetTimeSeriesT3ma200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesT3ma200Response) GetValues() []GetTimeSeriesT3ma200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesT3ma200Response) GetValuesOk() (*[]GetTimeSeriesT3ma200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesT3ma200Response) SetValues(v []GetTimeSeriesT3ma200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesT3ma200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesT3ma200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesT3ma200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


