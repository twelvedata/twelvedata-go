# GetTimeSeriesStochRsi200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesStochRsi200ResponseMeta**](GetTimeSeriesStochRsi200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesStochRsi200ResponseValuesInner**](GetTimeSeriesStochRsi200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesStochRsi200Response

`func NewGetTimeSeriesStochRsi200Response(meta GetTimeSeriesStochRsi200ResponseMeta, values []GetTimeSeriesStochRsi200ResponseValuesInner, status string, ) *GetTimeSeriesStochRsi200Response`

NewGetTimeSeriesStochRsi200Response instantiates a new GetTimeSeriesStochRsi200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesStochRsi200ResponseWithDefaults

`func NewGetTimeSeriesStochRsi200ResponseWithDefaults() *GetTimeSeriesStochRsi200Response`

NewGetTimeSeriesStochRsi200ResponseWithDefaults instantiates a new GetTimeSeriesStochRsi200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesStochRsi200Response) GetMeta() GetTimeSeriesStochRsi200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesStochRsi200Response) GetMetaOk() (*GetTimeSeriesStochRsi200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesStochRsi200Response) SetMeta(v GetTimeSeriesStochRsi200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesStochRsi200Response) GetValues() []GetTimeSeriesStochRsi200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesStochRsi200Response) GetValuesOk() (*[]GetTimeSeriesStochRsi200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesStochRsi200Response) SetValues(v []GetTimeSeriesStochRsi200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesStochRsi200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesStochRsi200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesStochRsi200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


