# GetTimeSeriesSma200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesSma200ResponseMeta**](GetTimeSeriesSma200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesSma200ResponseValuesInner**](GetTimeSeriesSma200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesSma200Response

`func NewGetTimeSeriesSma200Response(meta GetTimeSeriesSma200ResponseMeta, values []GetTimeSeriesSma200ResponseValuesInner, status string, ) *GetTimeSeriesSma200Response`

NewGetTimeSeriesSma200Response instantiates a new GetTimeSeriesSma200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesSma200ResponseWithDefaults

`func NewGetTimeSeriesSma200ResponseWithDefaults() *GetTimeSeriesSma200Response`

NewGetTimeSeriesSma200ResponseWithDefaults instantiates a new GetTimeSeriesSma200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesSma200Response) GetMeta() GetTimeSeriesSma200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesSma200Response) GetMetaOk() (*GetTimeSeriesSma200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesSma200Response) SetMeta(v GetTimeSeriesSma200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesSma200Response) GetValues() []GetTimeSeriesSma200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesSma200Response) GetValuesOk() (*[]GetTimeSeriesSma200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesSma200Response) SetValues(v []GetTimeSeriesSma200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesSma200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesSma200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesSma200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


