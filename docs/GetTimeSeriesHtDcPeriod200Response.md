# GetTimeSeriesHtDcPeriod200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesHtDcPeriod200ResponseMeta**](GetTimeSeriesHtDcPeriod200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesHtDcPeriod200ResponseValuesInner**](GetTimeSeriesHtDcPeriod200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesHtDcPeriod200Response

`func NewGetTimeSeriesHtDcPeriod200Response(meta GetTimeSeriesHtDcPeriod200ResponseMeta, values []GetTimeSeriesHtDcPeriod200ResponseValuesInner, status string, ) *GetTimeSeriesHtDcPeriod200Response`

NewGetTimeSeriesHtDcPeriod200Response instantiates a new GetTimeSeriesHtDcPeriod200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesHtDcPeriod200ResponseWithDefaults

`func NewGetTimeSeriesHtDcPeriod200ResponseWithDefaults() *GetTimeSeriesHtDcPeriod200Response`

NewGetTimeSeriesHtDcPeriod200ResponseWithDefaults instantiates a new GetTimeSeriesHtDcPeriod200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesHtDcPeriod200Response) GetMeta() GetTimeSeriesHtDcPeriod200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesHtDcPeriod200Response) GetMetaOk() (*GetTimeSeriesHtDcPeriod200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesHtDcPeriod200Response) SetMeta(v GetTimeSeriesHtDcPeriod200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesHtDcPeriod200Response) GetValues() []GetTimeSeriesHtDcPeriod200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesHtDcPeriod200Response) GetValuesOk() (*[]GetTimeSeriesHtDcPeriod200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesHtDcPeriod200Response) SetValues(v []GetTimeSeriesHtDcPeriod200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesHtDcPeriod200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesHtDcPeriod200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesHtDcPeriod200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


