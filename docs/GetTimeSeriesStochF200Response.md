# GetTimeSeriesStochF200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesStochF200ResponseMeta**](GetTimeSeriesStochF200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesStochF200ResponseValuesInner**](GetTimeSeriesStochF200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesStochF200Response

`func NewGetTimeSeriesStochF200Response(meta GetTimeSeriesStochF200ResponseMeta, values []GetTimeSeriesStochF200ResponseValuesInner, status string, ) *GetTimeSeriesStochF200Response`

NewGetTimeSeriesStochF200Response instantiates a new GetTimeSeriesStochF200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesStochF200ResponseWithDefaults

`func NewGetTimeSeriesStochF200ResponseWithDefaults() *GetTimeSeriesStochF200Response`

NewGetTimeSeriesStochF200ResponseWithDefaults instantiates a new GetTimeSeriesStochF200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesStochF200Response) GetMeta() GetTimeSeriesStochF200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesStochF200Response) GetMetaOk() (*GetTimeSeriesStochF200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesStochF200Response) SetMeta(v GetTimeSeriesStochF200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesStochF200Response) GetValues() []GetTimeSeriesStochF200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesStochF200Response) GetValuesOk() (*[]GetTimeSeriesStochF200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesStochF200Response) SetValues(v []GetTimeSeriesStochF200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesStochF200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesStochF200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesStochF200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


