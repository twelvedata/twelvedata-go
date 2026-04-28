# GetTimeSeriesWclPrice200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesWclPrice200ResponseMeta**](GetTimeSeriesWclPrice200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesWclPrice200ResponseValuesInner**](GetTimeSeriesWclPrice200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesWclPrice200Response

`func NewGetTimeSeriesWclPrice200Response(meta GetTimeSeriesWclPrice200ResponseMeta, values []GetTimeSeriesWclPrice200ResponseValuesInner, status string, ) *GetTimeSeriesWclPrice200Response`

NewGetTimeSeriesWclPrice200Response instantiates a new GetTimeSeriesWclPrice200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesWclPrice200ResponseWithDefaults

`func NewGetTimeSeriesWclPrice200ResponseWithDefaults() *GetTimeSeriesWclPrice200Response`

NewGetTimeSeriesWclPrice200ResponseWithDefaults instantiates a new GetTimeSeriesWclPrice200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesWclPrice200Response) GetMeta() GetTimeSeriesWclPrice200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesWclPrice200Response) GetMetaOk() (*GetTimeSeriesWclPrice200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesWclPrice200Response) SetMeta(v GetTimeSeriesWclPrice200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesWclPrice200Response) GetValues() []GetTimeSeriesWclPrice200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesWclPrice200Response) GetValuesOk() (*[]GetTimeSeriesWclPrice200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesWclPrice200Response) SetValues(v []GetTimeSeriesWclPrice200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesWclPrice200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesWclPrice200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesWclPrice200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


