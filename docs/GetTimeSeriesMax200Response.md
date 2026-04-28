# GetTimeSeriesMax200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesMax200ResponseMeta**](GetTimeSeriesMax200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesMax200ResponseValuesInner**](GetTimeSeriesMax200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesMax200Response

`func NewGetTimeSeriesMax200Response(meta GetTimeSeriesMax200ResponseMeta, values []GetTimeSeriesMax200ResponseValuesInner, status string, ) *GetTimeSeriesMax200Response`

NewGetTimeSeriesMax200Response instantiates a new GetTimeSeriesMax200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesMax200ResponseWithDefaults

`func NewGetTimeSeriesMax200ResponseWithDefaults() *GetTimeSeriesMax200Response`

NewGetTimeSeriesMax200ResponseWithDefaults instantiates a new GetTimeSeriesMax200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesMax200Response) GetMeta() GetTimeSeriesMax200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesMax200Response) GetMetaOk() (*GetTimeSeriesMax200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesMax200Response) SetMeta(v GetTimeSeriesMax200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesMax200Response) GetValues() []GetTimeSeriesMax200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesMax200Response) GetValuesOk() (*[]GetTimeSeriesMax200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesMax200Response) SetValues(v []GetTimeSeriesMax200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesMax200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesMax200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesMax200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


