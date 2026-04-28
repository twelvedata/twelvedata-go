# GetTimeSeriesMedPrice200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesMedPrice200ResponseMeta**](GetTimeSeriesMedPrice200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesMedPrice200ResponseValuesInner**](GetTimeSeriesMedPrice200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesMedPrice200Response

`func NewGetTimeSeriesMedPrice200Response(meta GetTimeSeriesMedPrice200ResponseMeta, values []GetTimeSeriesMedPrice200ResponseValuesInner, status string, ) *GetTimeSeriesMedPrice200Response`

NewGetTimeSeriesMedPrice200Response instantiates a new GetTimeSeriesMedPrice200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesMedPrice200ResponseWithDefaults

`func NewGetTimeSeriesMedPrice200ResponseWithDefaults() *GetTimeSeriesMedPrice200Response`

NewGetTimeSeriesMedPrice200ResponseWithDefaults instantiates a new GetTimeSeriesMedPrice200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesMedPrice200Response) GetMeta() GetTimeSeriesMedPrice200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesMedPrice200Response) GetMetaOk() (*GetTimeSeriesMedPrice200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesMedPrice200Response) SetMeta(v GetTimeSeriesMedPrice200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesMedPrice200Response) GetValues() []GetTimeSeriesMedPrice200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesMedPrice200Response) GetValuesOk() (*[]GetTimeSeriesMedPrice200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesMedPrice200Response) SetValues(v []GetTimeSeriesMedPrice200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesMedPrice200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesMedPrice200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesMedPrice200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


