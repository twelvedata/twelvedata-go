# GetTimeSeriesObv200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesObv200ResponseMeta**](GetTimeSeriesObv200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesObv200ResponseValuesInner**](GetTimeSeriesObv200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesObv200Response

`func NewGetTimeSeriesObv200Response(meta GetTimeSeriesObv200ResponseMeta, values []GetTimeSeriesObv200ResponseValuesInner, status string, ) *GetTimeSeriesObv200Response`

NewGetTimeSeriesObv200Response instantiates a new GetTimeSeriesObv200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesObv200ResponseWithDefaults

`func NewGetTimeSeriesObv200ResponseWithDefaults() *GetTimeSeriesObv200Response`

NewGetTimeSeriesObv200ResponseWithDefaults instantiates a new GetTimeSeriesObv200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesObv200Response) GetMeta() GetTimeSeriesObv200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesObv200Response) GetMetaOk() (*GetTimeSeriesObv200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesObv200Response) SetMeta(v GetTimeSeriesObv200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesObv200Response) GetValues() []GetTimeSeriesObv200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesObv200Response) GetValuesOk() (*[]GetTimeSeriesObv200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesObv200Response) SetValues(v []GetTimeSeriesObv200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesObv200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesObv200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesObv200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


