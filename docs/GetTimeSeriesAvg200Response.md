# GetTimeSeriesAvg200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesAvg200ResponseMeta**](GetTimeSeriesAvg200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesAvg200ResponseValuesInner**](GetTimeSeriesAvg200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesAvg200Response

`func NewGetTimeSeriesAvg200Response(meta GetTimeSeriesAvg200ResponseMeta, values []GetTimeSeriesAvg200ResponseValuesInner, status string, ) *GetTimeSeriesAvg200Response`

NewGetTimeSeriesAvg200Response instantiates a new GetTimeSeriesAvg200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesAvg200ResponseWithDefaults

`func NewGetTimeSeriesAvg200ResponseWithDefaults() *GetTimeSeriesAvg200Response`

NewGetTimeSeriesAvg200ResponseWithDefaults instantiates a new GetTimeSeriesAvg200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesAvg200Response) GetMeta() GetTimeSeriesAvg200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesAvg200Response) GetMetaOk() (*GetTimeSeriesAvg200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesAvg200Response) SetMeta(v GetTimeSeriesAvg200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesAvg200Response) GetValues() []GetTimeSeriesAvg200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesAvg200Response) GetValuesOk() (*[]GetTimeSeriesAvg200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesAvg200Response) SetValues(v []GetTimeSeriesAvg200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesAvg200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesAvg200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesAvg200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


