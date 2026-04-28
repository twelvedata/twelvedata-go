# GetTimeSeriesCci200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesCci200ResponseMeta**](GetTimeSeriesCci200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesCci200ResponseValuesInner**](GetTimeSeriesCci200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesCci200Response

`func NewGetTimeSeriesCci200Response(meta GetTimeSeriesCci200ResponseMeta, values []GetTimeSeriesCci200ResponseValuesInner, status string, ) *GetTimeSeriesCci200Response`

NewGetTimeSeriesCci200Response instantiates a new GetTimeSeriesCci200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesCci200ResponseWithDefaults

`func NewGetTimeSeriesCci200ResponseWithDefaults() *GetTimeSeriesCci200Response`

NewGetTimeSeriesCci200ResponseWithDefaults instantiates a new GetTimeSeriesCci200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesCci200Response) GetMeta() GetTimeSeriesCci200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesCci200Response) GetMetaOk() (*GetTimeSeriesCci200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesCci200Response) SetMeta(v GetTimeSeriesCci200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesCci200Response) GetValues() []GetTimeSeriesCci200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesCci200Response) GetValuesOk() (*[]GetTimeSeriesCci200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesCci200Response) SetValues(v []GetTimeSeriesCci200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesCci200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesCci200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesCci200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


