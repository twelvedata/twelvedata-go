# GetTimeSeriesAd200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesAd200ResponseMeta**](GetTimeSeriesAd200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesAd200ResponseValuesInner**](GetTimeSeriesAd200ResponseValuesInner.md) |  | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesAd200Response

`func NewGetTimeSeriesAd200Response(meta GetTimeSeriesAd200ResponseMeta, values []GetTimeSeriesAd200ResponseValuesInner, status string, ) *GetTimeSeriesAd200Response`

NewGetTimeSeriesAd200Response instantiates a new GetTimeSeriesAd200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesAd200ResponseWithDefaults

`func NewGetTimeSeriesAd200ResponseWithDefaults() *GetTimeSeriesAd200Response`

NewGetTimeSeriesAd200ResponseWithDefaults instantiates a new GetTimeSeriesAd200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesAd200Response) GetMeta() GetTimeSeriesAd200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesAd200Response) GetMetaOk() (*GetTimeSeriesAd200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesAd200Response) SetMeta(v GetTimeSeriesAd200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesAd200Response) GetValues() []GetTimeSeriesAd200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesAd200Response) GetValuesOk() (*[]GetTimeSeriesAd200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesAd200Response) SetValues(v []GetTimeSeriesAd200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesAd200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesAd200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesAd200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


