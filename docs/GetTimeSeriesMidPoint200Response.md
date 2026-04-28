# GetTimeSeriesMidPoint200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesMidPoint200ResponseMeta**](GetTimeSeriesMidPoint200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesMidPoint200ResponseValuesInner**](GetTimeSeriesMidPoint200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesMidPoint200Response

`func NewGetTimeSeriesMidPoint200Response(meta GetTimeSeriesMidPoint200ResponseMeta, values []GetTimeSeriesMidPoint200ResponseValuesInner, status string, ) *GetTimeSeriesMidPoint200Response`

NewGetTimeSeriesMidPoint200Response instantiates a new GetTimeSeriesMidPoint200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesMidPoint200ResponseWithDefaults

`func NewGetTimeSeriesMidPoint200ResponseWithDefaults() *GetTimeSeriesMidPoint200Response`

NewGetTimeSeriesMidPoint200ResponseWithDefaults instantiates a new GetTimeSeriesMidPoint200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesMidPoint200Response) GetMeta() GetTimeSeriesMidPoint200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesMidPoint200Response) GetMetaOk() (*GetTimeSeriesMidPoint200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesMidPoint200Response) SetMeta(v GetTimeSeriesMidPoint200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesMidPoint200Response) GetValues() []GetTimeSeriesMidPoint200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesMidPoint200Response) GetValuesOk() (*[]GetTimeSeriesMidPoint200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesMidPoint200Response) SetValues(v []GetTimeSeriesMidPoint200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesMidPoint200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesMidPoint200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesMidPoint200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


