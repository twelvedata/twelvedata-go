# GetTimeSeriesWma200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesWma200ResponseMeta**](GetTimeSeriesWma200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesWma200ResponseValuesInner**](GetTimeSeriesWma200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesWma200Response

`func NewGetTimeSeriesWma200Response(meta GetTimeSeriesWma200ResponseMeta, values []GetTimeSeriesWma200ResponseValuesInner, status string, ) *GetTimeSeriesWma200Response`

NewGetTimeSeriesWma200Response instantiates a new GetTimeSeriesWma200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesWma200ResponseWithDefaults

`func NewGetTimeSeriesWma200ResponseWithDefaults() *GetTimeSeriesWma200Response`

NewGetTimeSeriesWma200ResponseWithDefaults instantiates a new GetTimeSeriesWma200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesWma200Response) GetMeta() GetTimeSeriesWma200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesWma200Response) GetMetaOk() (*GetTimeSeriesWma200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesWma200Response) SetMeta(v GetTimeSeriesWma200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesWma200Response) GetValues() []GetTimeSeriesWma200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesWma200Response) GetValuesOk() (*[]GetTimeSeriesWma200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesWma200Response) SetValues(v []GetTimeSeriesWma200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesWma200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesWma200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesWma200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


