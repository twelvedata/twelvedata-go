# GetTimeSeriesAdOsc200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesAdOsc200ResponseMeta**](GetTimeSeriesAdOsc200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesAdOsc200ResponseValuesInner**](GetTimeSeriesAdOsc200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesAdOsc200Response

`func NewGetTimeSeriesAdOsc200Response(meta GetTimeSeriesAdOsc200ResponseMeta, values []GetTimeSeriesAdOsc200ResponseValuesInner, status string, ) *GetTimeSeriesAdOsc200Response`

NewGetTimeSeriesAdOsc200Response instantiates a new GetTimeSeriesAdOsc200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesAdOsc200ResponseWithDefaults

`func NewGetTimeSeriesAdOsc200ResponseWithDefaults() *GetTimeSeriesAdOsc200Response`

NewGetTimeSeriesAdOsc200ResponseWithDefaults instantiates a new GetTimeSeriesAdOsc200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesAdOsc200Response) GetMeta() GetTimeSeriesAdOsc200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesAdOsc200Response) GetMetaOk() (*GetTimeSeriesAdOsc200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesAdOsc200Response) SetMeta(v GetTimeSeriesAdOsc200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesAdOsc200Response) GetValues() []GetTimeSeriesAdOsc200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesAdOsc200Response) GetValuesOk() (*[]GetTimeSeriesAdOsc200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesAdOsc200Response) SetValues(v []GetTimeSeriesAdOsc200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesAdOsc200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesAdOsc200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesAdOsc200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


