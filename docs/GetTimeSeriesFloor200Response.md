# GetTimeSeriesFloor200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesFloor200ResponseMeta**](GetTimeSeriesFloor200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesFloor200ResponseValuesInner**](GetTimeSeriesFloor200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesFloor200Response

`func NewGetTimeSeriesFloor200Response(meta GetTimeSeriesFloor200ResponseMeta, values []GetTimeSeriesFloor200ResponseValuesInner, status string, ) *GetTimeSeriesFloor200Response`

NewGetTimeSeriesFloor200Response instantiates a new GetTimeSeriesFloor200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesFloor200ResponseWithDefaults

`func NewGetTimeSeriesFloor200ResponseWithDefaults() *GetTimeSeriesFloor200Response`

NewGetTimeSeriesFloor200ResponseWithDefaults instantiates a new GetTimeSeriesFloor200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesFloor200Response) GetMeta() GetTimeSeriesFloor200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesFloor200Response) GetMetaOk() (*GetTimeSeriesFloor200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesFloor200Response) SetMeta(v GetTimeSeriesFloor200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesFloor200Response) GetValues() []GetTimeSeriesFloor200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesFloor200Response) GetValuesOk() (*[]GetTimeSeriesFloor200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesFloor200Response) SetValues(v []GetTimeSeriesFloor200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesFloor200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesFloor200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesFloor200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


