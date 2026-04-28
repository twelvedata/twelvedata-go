# GetTimeSeriesMinMaxIndex200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesMinMaxIndex200ResponseMeta**](GetTimeSeriesMinMaxIndex200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesMinMaxIndex200ResponseValuesInner**](GetTimeSeriesMinMaxIndex200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesMinMaxIndex200Response

`func NewGetTimeSeriesMinMaxIndex200Response(meta GetTimeSeriesMinMaxIndex200ResponseMeta, values []GetTimeSeriesMinMaxIndex200ResponseValuesInner, status string, ) *GetTimeSeriesMinMaxIndex200Response`

NewGetTimeSeriesMinMaxIndex200Response instantiates a new GetTimeSeriesMinMaxIndex200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesMinMaxIndex200ResponseWithDefaults

`func NewGetTimeSeriesMinMaxIndex200ResponseWithDefaults() *GetTimeSeriesMinMaxIndex200Response`

NewGetTimeSeriesMinMaxIndex200ResponseWithDefaults instantiates a new GetTimeSeriesMinMaxIndex200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesMinMaxIndex200Response) GetMeta() GetTimeSeriesMinMaxIndex200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesMinMaxIndex200Response) GetMetaOk() (*GetTimeSeriesMinMaxIndex200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesMinMaxIndex200Response) SetMeta(v GetTimeSeriesMinMaxIndex200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesMinMaxIndex200Response) GetValues() []GetTimeSeriesMinMaxIndex200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesMinMaxIndex200Response) GetValuesOk() (*[]GetTimeSeriesMinMaxIndex200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesMinMaxIndex200Response) SetValues(v []GetTimeSeriesMinMaxIndex200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesMinMaxIndex200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesMinMaxIndex200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesMinMaxIndex200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


