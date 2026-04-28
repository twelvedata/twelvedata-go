# GetTimeSeriesMcGinleyDynamic200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesMcGinleyDynamic200ResponseMeta**](GetTimeSeriesMcGinleyDynamic200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesMcGinleyDynamic200ResponseValuesInner**](GetTimeSeriesMcGinleyDynamic200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesMcGinleyDynamic200Response

`func NewGetTimeSeriesMcGinleyDynamic200Response(meta GetTimeSeriesMcGinleyDynamic200ResponseMeta, values []GetTimeSeriesMcGinleyDynamic200ResponseValuesInner, status string, ) *GetTimeSeriesMcGinleyDynamic200Response`

NewGetTimeSeriesMcGinleyDynamic200Response instantiates a new GetTimeSeriesMcGinleyDynamic200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesMcGinleyDynamic200ResponseWithDefaults

`func NewGetTimeSeriesMcGinleyDynamic200ResponseWithDefaults() *GetTimeSeriesMcGinleyDynamic200Response`

NewGetTimeSeriesMcGinleyDynamic200ResponseWithDefaults instantiates a new GetTimeSeriesMcGinleyDynamic200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesMcGinleyDynamic200Response) GetMeta() GetTimeSeriesMcGinleyDynamic200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesMcGinleyDynamic200Response) GetMetaOk() (*GetTimeSeriesMcGinleyDynamic200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesMcGinleyDynamic200Response) SetMeta(v GetTimeSeriesMcGinleyDynamic200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesMcGinleyDynamic200Response) GetValues() []GetTimeSeriesMcGinleyDynamic200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesMcGinleyDynamic200Response) GetValuesOk() (*[]GetTimeSeriesMcGinleyDynamic200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesMcGinleyDynamic200Response) SetValues(v []GetTimeSeriesMcGinleyDynamic200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesMcGinleyDynamic200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesMcGinleyDynamic200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesMcGinleyDynamic200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


