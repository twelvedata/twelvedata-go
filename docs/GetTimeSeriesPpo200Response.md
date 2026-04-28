# GetTimeSeriesPpo200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesPpo200ResponseMeta**](GetTimeSeriesPpo200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesPpo200ResponseValuesInner**](GetTimeSeriesPpo200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesPpo200Response

`func NewGetTimeSeriesPpo200Response(meta GetTimeSeriesPpo200ResponseMeta, values []GetTimeSeriesPpo200ResponseValuesInner, status string, ) *GetTimeSeriesPpo200Response`

NewGetTimeSeriesPpo200Response instantiates a new GetTimeSeriesPpo200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesPpo200ResponseWithDefaults

`func NewGetTimeSeriesPpo200ResponseWithDefaults() *GetTimeSeriesPpo200Response`

NewGetTimeSeriesPpo200ResponseWithDefaults instantiates a new GetTimeSeriesPpo200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesPpo200Response) GetMeta() GetTimeSeriesPpo200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesPpo200Response) GetMetaOk() (*GetTimeSeriesPpo200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesPpo200Response) SetMeta(v GetTimeSeriesPpo200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesPpo200Response) GetValues() []GetTimeSeriesPpo200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesPpo200Response) GetValuesOk() (*[]GetTimeSeriesPpo200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesPpo200Response) SetValues(v []GetTimeSeriesPpo200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesPpo200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesPpo200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesPpo200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


