# GetTimeSeriesBBands200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesBBands200ResponseMeta**](GetTimeSeriesBBands200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesBBands200ResponseValuesInner**](GetTimeSeriesBBands200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesBBands200Response

`func NewGetTimeSeriesBBands200Response(meta GetTimeSeriesBBands200ResponseMeta, values []GetTimeSeriesBBands200ResponseValuesInner, status string, ) *GetTimeSeriesBBands200Response`

NewGetTimeSeriesBBands200Response instantiates a new GetTimeSeriesBBands200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesBBands200ResponseWithDefaults

`func NewGetTimeSeriesBBands200ResponseWithDefaults() *GetTimeSeriesBBands200Response`

NewGetTimeSeriesBBands200ResponseWithDefaults instantiates a new GetTimeSeriesBBands200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesBBands200Response) GetMeta() GetTimeSeriesBBands200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesBBands200Response) GetMetaOk() (*GetTimeSeriesBBands200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesBBands200Response) SetMeta(v GetTimeSeriesBBands200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesBBands200Response) GetValues() []GetTimeSeriesBBands200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesBBands200Response) GetValuesOk() (*[]GetTimeSeriesBBands200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesBBands200Response) SetValues(v []GetTimeSeriesBBands200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesBBands200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesBBands200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesBBands200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


