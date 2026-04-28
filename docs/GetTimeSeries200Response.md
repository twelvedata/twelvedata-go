# GetTimeSeries200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeries200ResponseMeta**](GetTimeSeries200ResponseMeta.md) |  | 
**Values** | [**[]TimeSeriesItem**](TimeSeriesItem.md) | List of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeries200Response

`func NewGetTimeSeries200Response(meta GetTimeSeries200ResponseMeta, values []TimeSeriesItem, status string, ) *GetTimeSeries200Response`

NewGetTimeSeries200Response instantiates a new GetTimeSeries200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeries200ResponseWithDefaults

`func NewGetTimeSeries200ResponseWithDefaults() *GetTimeSeries200Response`

NewGetTimeSeries200ResponseWithDefaults instantiates a new GetTimeSeries200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeries200Response) GetMeta() GetTimeSeries200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeries200Response) GetMetaOk() (*GetTimeSeries200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeries200Response) SetMeta(v GetTimeSeries200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeries200Response) GetValues() []TimeSeriesItem`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeries200Response) GetValuesOk() (*[]TimeSeriesItem, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeries200Response) SetValues(v []TimeSeriesItem)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeries200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeries200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeries200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


