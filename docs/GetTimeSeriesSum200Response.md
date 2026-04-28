# GetTimeSeriesSum200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesSum200ResponseMeta**](GetTimeSeriesSum200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesSum200ResponseValuesInner**](GetTimeSeriesSum200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesSum200Response

`func NewGetTimeSeriesSum200Response(meta GetTimeSeriesSum200ResponseMeta, values []GetTimeSeriesSum200ResponseValuesInner, status string, ) *GetTimeSeriesSum200Response`

NewGetTimeSeriesSum200Response instantiates a new GetTimeSeriesSum200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesSum200ResponseWithDefaults

`func NewGetTimeSeriesSum200ResponseWithDefaults() *GetTimeSeriesSum200Response`

NewGetTimeSeriesSum200ResponseWithDefaults instantiates a new GetTimeSeriesSum200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesSum200Response) GetMeta() GetTimeSeriesSum200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesSum200Response) GetMetaOk() (*GetTimeSeriesSum200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesSum200Response) SetMeta(v GetTimeSeriesSum200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesSum200Response) GetValues() []GetTimeSeriesSum200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesSum200Response) GetValuesOk() (*[]GetTimeSeriesSum200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesSum200Response) SetValues(v []GetTimeSeriesSum200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesSum200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesSum200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesSum200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


