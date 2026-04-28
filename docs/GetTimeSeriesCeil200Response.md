# GetTimeSeriesCeil200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesCeil200ResponseMeta**](GetTimeSeriesCeil200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesCeil200ResponseValuesInner**](GetTimeSeriesCeil200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesCeil200Response

`func NewGetTimeSeriesCeil200Response(meta GetTimeSeriesCeil200ResponseMeta, values []GetTimeSeriesCeil200ResponseValuesInner, status string, ) *GetTimeSeriesCeil200Response`

NewGetTimeSeriesCeil200Response instantiates a new GetTimeSeriesCeil200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesCeil200ResponseWithDefaults

`func NewGetTimeSeriesCeil200ResponseWithDefaults() *GetTimeSeriesCeil200Response`

NewGetTimeSeriesCeil200ResponseWithDefaults instantiates a new GetTimeSeriesCeil200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesCeil200Response) GetMeta() GetTimeSeriesCeil200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesCeil200Response) GetMetaOk() (*GetTimeSeriesCeil200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesCeil200Response) SetMeta(v GetTimeSeriesCeil200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesCeil200Response) GetValues() []GetTimeSeriesCeil200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesCeil200Response) GetValuesOk() (*[]GetTimeSeriesCeil200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesCeil200Response) SetValues(v []GetTimeSeriesCeil200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesCeil200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesCeil200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesCeil200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


