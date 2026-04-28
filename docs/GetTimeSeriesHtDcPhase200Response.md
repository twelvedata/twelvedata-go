# GetTimeSeriesHtDcPhase200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesHtDcPhase200ResponseMeta**](GetTimeSeriesHtDcPhase200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesHtDcPhase200ResponseValuesInner**](GetTimeSeriesHtDcPhase200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesHtDcPhase200Response

`func NewGetTimeSeriesHtDcPhase200Response(meta GetTimeSeriesHtDcPhase200ResponseMeta, values []GetTimeSeriesHtDcPhase200ResponseValuesInner, status string, ) *GetTimeSeriesHtDcPhase200Response`

NewGetTimeSeriesHtDcPhase200Response instantiates a new GetTimeSeriesHtDcPhase200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesHtDcPhase200ResponseWithDefaults

`func NewGetTimeSeriesHtDcPhase200ResponseWithDefaults() *GetTimeSeriesHtDcPhase200Response`

NewGetTimeSeriesHtDcPhase200ResponseWithDefaults instantiates a new GetTimeSeriesHtDcPhase200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesHtDcPhase200Response) GetMeta() GetTimeSeriesHtDcPhase200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesHtDcPhase200Response) GetMetaOk() (*GetTimeSeriesHtDcPhase200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesHtDcPhase200Response) SetMeta(v GetTimeSeriesHtDcPhase200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesHtDcPhase200Response) GetValues() []GetTimeSeriesHtDcPhase200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesHtDcPhase200Response) GetValuesOk() (*[]GetTimeSeriesHtDcPhase200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesHtDcPhase200Response) SetValues(v []GetTimeSeriesHtDcPhase200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesHtDcPhase200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesHtDcPhase200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesHtDcPhase200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


