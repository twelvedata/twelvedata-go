# GetTimeSeriesBop200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesBop200ResponseMeta**](GetTimeSeriesBop200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesBop200ResponseValuesInner**](GetTimeSeriesBop200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesBop200Response

`func NewGetTimeSeriesBop200Response(meta GetTimeSeriesBop200ResponseMeta, values []GetTimeSeriesBop200ResponseValuesInner, status string, ) *GetTimeSeriesBop200Response`

NewGetTimeSeriesBop200Response instantiates a new GetTimeSeriesBop200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesBop200ResponseWithDefaults

`func NewGetTimeSeriesBop200ResponseWithDefaults() *GetTimeSeriesBop200Response`

NewGetTimeSeriesBop200ResponseWithDefaults instantiates a new GetTimeSeriesBop200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesBop200Response) GetMeta() GetTimeSeriesBop200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesBop200Response) GetMetaOk() (*GetTimeSeriesBop200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesBop200Response) SetMeta(v GetTimeSeriesBop200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesBop200Response) GetValues() []GetTimeSeriesBop200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesBop200Response) GetValuesOk() (*[]GetTimeSeriesBop200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesBop200Response) SetValues(v []GetTimeSeriesBop200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesBop200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesBop200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesBop200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


