# GetTimeSeriesKeltner200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesKeltner200ResponseMeta**](GetTimeSeriesKeltner200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesKeltner200ResponseValuesInner**](GetTimeSeriesKeltner200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesKeltner200Response

`func NewGetTimeSeriesKeltner200Response(meta GetTimeSeriesKeltner200ResponseMeta, values []GetTimeSeriesKeltner200ResponseValuesInner, status string, ) *GetTimeSeriesKeltner200Response`

NewGetTimeSeriesKeltner200Response instantiates a new GetTimeSeriesKeltner200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesKeltner200ResponseWithDefaults

`func NewGetTimeSeriesKeltner200ResponseWithDefaults() *GetTimeSeriesKeltner200Response`

NewGetTimeSeriesKeltner200ResponseWithDefaults instantiates a new GetTimeSeriesKeltner200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesKeltner200Response) GetMeta() GetTimeSeriesKeltner200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesKeltner200Response) GetMetaOk() (*GetTimeSeriesKeltner200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesKeltner200Response) SetMeta(v GetTimeSeriesKeltner200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesKeltner200Response) GetValues() []GetTimeSeriesKeltner200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesKeltner200Response) GetValuesOk() (*[]GetTimeSeriesKeltner200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesKeltner200Response) SetValues(v []GetTimeSeriesKeltner200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesKeltner200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesKeltner200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesKeltner200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


