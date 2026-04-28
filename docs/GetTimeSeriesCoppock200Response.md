# GetTimeSeriesCoppock200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesCoppock200ResponseMeta**](GetTimeSeriesCoppock200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesCoppock200ResponseValuesInner**](GetTimeSeriesCoppock200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesCoppock200Response

`func NewGetTimeSeriesCoppock200Response(meta GetTimeSeriesCoppock200ResponseMeta, values []GetTimeSeriesCoppock200ResponseValuesInner, status string, ) *GetTimeSeriesCoppock200Response`

NewGetTimeSeriesCoppock200Response instantiates a new GetTimeSeriesCoppock200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesCoppock200ResponseWithDefaults

`func NewGetTimeSeriesCoppock200ResponseWithDefaults() *GetTimeSeriesCoppock200Response`

NewGetTimeSeriesCoppock200ResponseWithDefaults instantiates a new GetTimeSeriesCoppock200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesCoppock200Response) GetMeta() GetTimeSeriesCoppock200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesCoppock200Response) GetMetaOk() (*GetTimeSeriesCoppock200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesCoppock200Response) SetMeta(v GetTimeSeriesCoppock200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesCoppock200Response) GetValues() []GetTimeSeriesCoppock200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesCoppock200Response) GetValuesOk() (*[]GetTimeSeriesCoppock200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesCoppock200Response) SetValues(v []GetTimeSeriesCoppock200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesCoppock200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesCoppock200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesCoppock200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


