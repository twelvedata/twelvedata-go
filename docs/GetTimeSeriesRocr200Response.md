# GetTimeSeriesRocr200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesRocr200ResponseMeta**](GetTimeSeriesRocr200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesRocr200ResponseValuesInner**](GetTimeSeriesRocr200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesRocr200Response

`func NewGetTimeSeriesRocr200Response(meta GetTimeSeriesRocr200ResponseMeta, values []GetTimeSeriesRocr200ResponseValuesInner, status string, ) *GetTimeSeriesRocr200Response`

NewGetTimeSeriesRocr200Response instantiates a new GetTimeSeriesRocr200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesRocr200ResponseWithDefaults

`func NewGetTimeSeriesRocr200ResponseWithDefaults() *GetTimeSeriesRocr200Response`

NewGetTimeSeriesRocr200ResponseWithDefaults instantiates a new GetTimeSeriesRocr200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesRocr200Response) GetMeta() GetTimeSeriesRocr200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesRocr200Response) GetMetaOk() (*GetTimeSeriesRocr200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesRocr200Response) SetMeta(v GetTimeSeriesRocr200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesRocr200Response) GetValues() []GetTimeSeriesRocr200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesRocr200Response) GetValuesOk() (*[]GetTimeSeriesRocr200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesRocr200Response) SetValues(v []GetTimeSeriesRocr200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesRocr200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesRocr200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesRocr200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


