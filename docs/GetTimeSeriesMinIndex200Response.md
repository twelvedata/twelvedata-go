# GetTimeSeriesMinIndex200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesMinIndex200ResponseMeta**](GetTimeSeriesMinIndex200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesMinIndex200ResponseValuesInner**](GetTimeSeriesMinIndex200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesMinIndex200Response

`func NewGetTimeSeriesMinIndex200Response(meta GetTimeSeriesMinIndex200ResponseMeta, values []GetTimeSeriesMinIndex200ResponseValuesInner, status string, ) *GetTimeSeriesMinIndex200Response`

NewGetTimeSeriesMinIndex200Response instantiates a new GetTimeSeriesMinIndex200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesMinIndex200ResponseWithDefaults

`func NewGetTimeSeriesMinIndex200ResponseWithDefaults() *GetTimeSeriesMinIndex200Response`

NewGetTimeSeriesMinIndex200ResponseWithDefaults instantiates a new GetTimeSeriesMinIndex200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesMinIndex200Response) GetMeta() GetTimeSeriesMinIndex200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesMinIndex200Response) GetMetaOk() (*GetTimeSeriesMinIndex200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesMinIndex200Response) SetMeta(v GetTimeSeriesMinIndex200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesMinIndex200Response) GetValues() []GetTimeSeriesMinIndex200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesMinIndex200Response) GetValuesOk() (*[]GetTimeSeriesMinIndex200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesMinIndex200Response) SetValues(v []GetTimeSeriesMinIndex200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesMinIndex200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesMinIndex200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesMinIndex200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


