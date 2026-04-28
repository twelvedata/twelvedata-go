# GetTimeSeriesPivotPointsHL200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesPivotPointsHL200ResponseMeta**](GetTimeSeriesPivotPointsHL200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesPivotPointsHL200ResponseValuesInner**](GetTimeSeriesPivotPointsHL200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesPivotPointsHL200Response

`func NewGetTimeSeriesPivotPointsHL200Response(meta GetTimeSeriesPivotPointsHL200ResponseMeta, values []GetTimeSeriesPivotPointsHL200ResponseValuesInner, status string, ) *GetTimeSeriesPivotPointsHL200Response`

NewGetTimeSeriesPivotPointsHL200Response instantiates a new GetTimeSeriesPivotPointsHL200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesPivotPointsHL200ResponseWithDefaults

`func NewGetTimeSeriesPivotPointsHL200ResponseWithDefaults() *GetTimeSeriesPivotPointsHL200Response`

NewGetTimeSeriesPivotPointsHL200ResponseWithDefaults instantiates a new GetTimeSeriesPivotPointsHL200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesPivotPointsHL200Response) GetMeta() GetTimeSeriesPivotPointsHL200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesPivotPointsHL200Response) GetMetaOk() (*GetTimeSeriesPivotPointsHL200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesPivotPointsHL200Response) SetMeta(v GetTimeSeriesPivotPointsHL200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesPivotPointsHL200Response) GetValues() []GetTimeSeriesPivotPointsHL200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesPivotPointsHL200Response) GetValuesOk() (*[]GetTimeSeriesPivotPointsHL200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesPivotPointsHL200Response) SetValues(v []GetTimeSeriesPivotPointsHL200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesPivotPointsHL200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesPivotPointsHL200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesPivotPointsHL200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


