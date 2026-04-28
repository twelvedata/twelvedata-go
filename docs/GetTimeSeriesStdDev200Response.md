# GetTimeSeriesStdDev200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesStdDev200ResponseMeta**](GetTimeSeriesStdDev200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesStdDev200ResponseValuesInner**](GetTimeSeriesStdDev200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesStdDev200Response

`func NewGetTimeSeriesStdDev200Response(meta GetTimeSeriesStdDev200ResponseMeta, values []GetTimeSeriesStdDev200ResponseValuesInner, status string, ) *GetTimeSeriesStdDev200Response`

NewGetTimeSeriesStdDev200Response instantiates a new GetTimeSeriesStdDev200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesStdDev200ResponseWithDefaults

`func NewGetTimeSeriesStdDev200ResponseWithDefaults() *GetTimeSeriesStdDev200Response`

NewGetTimeSeriesStdDev200ResponseWithDefaults instantiates a new GetTimeSeriesStdDev200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesStdDev200Response) GetMeta() GetTimeSeriesStdDev200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesStdDev200Response) GetMetaOk() (*GetTimeSeriesStdDev200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesStdDev200Response) SetMeta(v GetTimeSeriesStdDev200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesStdDev200Response) GetValues() []GetTimeSeriesStdDev200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesStdDev200Response) GetValuesOk() (*[]GetTimeSeriesStdDev200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesStdDev200Response) SetValues(v []GetTimeSeriesStdDev200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesStdDev200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesStdDev200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesStdDev200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


