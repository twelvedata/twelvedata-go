# GetTimeSeriesAvgPrice200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesAvgPrice200ResponseMeta**](GetTimeSeriesAvgPrice200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesAvgPrice200ResponseValuesInner**](GetTimeSeriesAvgPrice200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesAvgPrice200Response

`func NewGetTimeSeriesAvgPrice200Response(meta GetTimeSeriesAvgPrice200ResponseMeta, values []GetTimeSeriesAvgPrice200ResponseValuesInner, status string, ) *GetTimeSeriesAvgPrice200Response`

NewGetTimeSeriesAvgPrice200Response instantiates a new GetTimeSeriesAvgPrice200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesAvgPrice200ResponseWithDefaults

`func NewGetTimeSeriesAvgPrice200ResponseWithDefaults() *GetTimeSeriesAvgPrice200Response`

NewGetTimeSeriesAvgPrice200ResponseWithDefaults instantiates a new GetTimeSeriesAvgPrice200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesAvgPrice200Response) GetMeta() GetTimeSeriesAvgPrice200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesAvgPrice200Response) GetMetaOk() (*GetTimeSeriesAvgPrice200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesAvgPrice200Response) SetMeta(v GetTimeSeriesAvgPrice200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesAvgPrice200Response) GetValues() []GetTimeSeriesAvgPrice200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesAvgPrice200Response) GetValuesOk() (*[]GetTimeSeriesAvgPrice200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesAvgPrice200Response) SetValues(v []GetTimeSeriesAvgPrice200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesAvgPrice200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesAvgPrice200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesAvgPrice200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


