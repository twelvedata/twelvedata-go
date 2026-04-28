# GetTimeSeriesCmo200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesCmo200ResponseMeta**](GetTimeSeriesCmo200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesCmo200ResponseValuesInner**](GetTimeSeriesCmo200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesCmo200Response

`func NewGetTimeSeriesCmo200Response(meta GetTimeSeriesCmo200ResponseMeta, values []GetTimeSeriesCmo200ResponseValuesInner, status string, ) *GetTimeSeriesCmo200Response`

NewGetTimeSeriesCmo200Response instantiates a new GetTimeSeriesCmo200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesCmo200ResponseWithDefaults

`func NewGetTimeSeriesCmo200ResponseWithDefaults() *GetTimeSeriesCmo200Response`

NewGetTimeSeriesCmo200ResponseWithDefaults instantiates a new GetTimeSeriesCmo200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesCmo200Response) GetMeta() GetTimeSeriesCmo200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesCmo200Response) GetMetaOk() (*GetTimeSeriesCmo200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesCmo200Response) SetMeta(v GetTimeSeriesCmo200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesCmo200Response) GetValues() []GetTimeSeriesCmo200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesCmo200Response) GetValuesOk() (*[]GetTimeSeriesCmo200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesCmo200Response) SetValues(v []GetTimeSeriesCmo200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesCmo200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesCmo200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesCmo200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


