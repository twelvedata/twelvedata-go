# GetTimeSeriesRvol200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesRvol200ResponseMeta**](GetTimeSeriesRvol200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesRvol200ResponseValuesInner**](GetTimeSeriesRvol200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesRvol200Response

`func NewGetTimeSeriesRvol200Response(meta GetTimeSeriesRvol200ResponseMeta, values []GetTimeSeriesRvol200ResponseValuesInner, status string, ) *GetTimeSeriesRvol200Response`

NewGetTimeSeriesRvol200Response instantiates a new GetTimeSeriesRvol200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesRvol200ResponseWithDefaults

`func NewGetTimeSeriesRvol200ResponseWithDefaults() *GetTimeSeriesRvol200Response`

NewGetTimeSeriesRvol200ResponseWithDefaults instantiates a new GetTimeSeriesRvol200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesRvol200Response) GetMeta() GetTimeSeriesRvol200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesRvol200Response) GetMetaOk() (*GetTimeSeriesRvol200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesRvol200Response) SetMeta(v GetTimeSeriesRvol200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesRvol200Response) GetValues() []GetTimeSeriesRvol200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesRvol200Response) GetValuesOk() (*[]GetTimeSeriesRvol200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesRvol200Response) SetValues(v []GetTimeSeriesRvol200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesRvol200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesRvol200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesRvol200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


