# GetTimeSeriesIchimoku200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesIchimoku200ResponseMeta**](GetTimeSeriesIchimoku200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesIchimoku200ResponseValuesInner**](GetTimeSeriesIchimoku200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesIchimoku200Response

`func NewGetTimeSeriesIchimoku200Response(meta GetTimeSeriesIchimoku200ResponseMeta, values []GetTimeSeriesIchimoku200ResponseValuesInner, status string, ) *GetTimeSeriesIchimoku200Response`

NewGetTimeSeriesIchimoku200Response instantiates a new GetTimeSeriesIchimoku200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesIchimoku200ResponseWithDefaults

`func NewGetTimeSeriesIchimoku200ResponseWithDefaults() *GetTimeSeriesIchimoku200Response`

NewGetTimeSeriesIchimoku200ResponseWithDefaults instantiates a new GetTimeSeriesIchimoku200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesIchimoku200Response) GetMeta() GetTimeSeriesIchimoku200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesIchimoku200Response) GetMetaOk() (*GetTimeSeriesIchimoku200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesIchimoku200Response) SetMeta(v GetTimeSeriesIchimoku200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesIchimoku200Response) GetValues() []GetTimeSeriesIchimoku200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesIchimoku200Response) GetValuesOk() (*[]GetTimeSeriesIchimoku200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesIchimoku200Response) SetValues(v []GetTimeSeriesIchimoku200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesIchimoku200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesIchimoku200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesIchimoku200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


