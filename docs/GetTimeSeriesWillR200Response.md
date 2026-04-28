# GetTimeSeriesWillR200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesWillR200ResponseMeta**](GetTimeSeriesWillR200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesWillR200ResponseValuesInner**](GetTimeSeriesWillR200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesWillR200Response

`func NewGetTimeSeriesWillR200Response(meta GetTimeSeriesWillR200ResponseMeta, values []GetTimeSeriesWillR200ResponseValuesInner, status string, ) *GetTimeSeriesWillR200Response`

NewGetTimeSeriesWillR200Response instantiates a new GetTimeSeriesWillR200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesWillR200ResponseWithDefaults

`func NewGetTimeSeriesWillR200ResponseWithDefaults() *GetTimeSeriesWillR200Response`

NewGetTimeSeriesWillR200ResponseWithDefaults instantiates a new GetTimeSeriesWillR200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesWillR200Response) GetMeta() GetTimeSeriesWillR200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesWillR200Response) GetMetaOk() (*GetTimeSeriesWillR200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesWillR200Response) SetMeta(v GetTimeSeriesWillR200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesWillR200Response) GetValues() []GetTimeSeriesWillR200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesWillR200Response) GetValuesOk() (*[]GetTimeSeriesWillR200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesWillR200Response) SetValues(v []GetTimeSeriesWillR200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesWillR200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesWillR200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesWillR200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


