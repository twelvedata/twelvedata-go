# GetTimeSeriesTsf200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesTsf200ResponseMeta**](GetTimeSeriesTsf200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesTsf200ResponseValuesInner**](GetTimeSeriesTsf200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesTsf200Response

`func NewGetTimeSeriesTsf200Response(meta GetTimeSeriesTsf200ResponseMeta, values []GetTimeSeriesTsf200ResponseValuesInner, status string, ) *GetTimeSeriesTsf200Response`

NewGetTimeSeriesTsf200Response instantiates a new GetTimeSeriesTsf200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesTsf200ResponseWithDefaults

`func NewGetTimeSeriesTsf200ResponseWithDefaults() *GetTimeSeriesTsf200Response`

NewGetTimeSeriesTsf200ResponseWithDefaults instantiates a new GetTimeSeriesTsf200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesTsf200Response) GetMeta() GetTimeSeriesTsf200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesTsf200Response) GetMetaOk() (*GetTimeSeriesTsf200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesTsf200Response) SetMeta(v GetTimeSeriesTsf200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesTsf200Response) GetValues() []GetTimeSeriesTsf200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesTsf200Response) GetValuesOk() (*[]GetTimeSeriesTsf200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesTsf200Response) SetValues(v []GetTimeSeriesTsf200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesTsf200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesTsf200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesTsf200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


