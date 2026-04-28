# GetTimeSeriesLn200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesLn200ResponseMeta**](GetTimeSeriesLn200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesLn200ResponseValuesInner**](GetTimeSeriesLn200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesLn200Response

`func NewGetTimeSeriesLn200Response(meta GetTimeSeriesLn200ResponseMeta, values []GetTimeSeriesLn200ResponseValuesInner, status string, ) *GetTimeSeriesLn200Response`

NewGetTimeSeriesLn200Response instantiates a new GetTimeSeriesLn200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesLn200ResponseWithDefaults

`func NewGetTimeSeriesLn200ResponseWithDefaults() *GetTimeSeriesLn200Response`

NewGetTimeSeriesLn200ResponseWithDefaults instantiates a new GetTimeSeriesLn200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesLn200Response) GetMeta() GetTimeSeriesLn200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesLn200Response) GetMetaOk() (*GetTimeSeriesLn200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesLn200Response) SetMeta(v GetTimeSeriesLn200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesLn200Response) GetValues() []GetTimeSeriesLn200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesLn200Response) GetValuesOk() (*[]GetTimeSeriesLn200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesLn200Response) SetValues(v []GetTimeSeriesLn200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesLn200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesLn200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesLn200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


