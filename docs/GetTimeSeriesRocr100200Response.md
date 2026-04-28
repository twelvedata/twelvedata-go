# GetTimeSeriesRocr100200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetTimeSeriesRocr100200ResponseMeta**](GetTimeSeriesRocr100200ResponseMeta.md) |  | 
**Values** | [**[]GetTimeSeriesRocr100200ResponseValuesInner**](GetTimeSeriesRocr100200ResponseValuesInner.md) | Array of time series data points | 
**Status** | **string** | Response status | 

## Methods

### NewGetTimeSeriesRocr100200Response

`func NewGetTimeSeriesRocr100200Response(meta GetTimeSeriesRocr100200ResponseMeta, values []GetTimeSeriesRocr100200ResponseValuesInner, status string, ) *GetTimeSeriesRocr100200Response`

NewGetTimeSeriesRocr100200Response instantiates a new GetTimeSeriesRocr100200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTimeSeriesRocr100200ResponseWithDefaults

`func NewGetTimeSeriesRocr100200ResponseWithDefaults() *GetTimeSeriesRocr100200Response`

NewGetTimeSeriesRocr100200ResponseWithDefaults instantiates a new GetTimeSeriesRocr100200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetTimeSeriesRocr100200Response) GetMeta() GetTimeSeriesRocr100200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetTimeSeriesRocr100200Response) GetMetaOk() (*GetTimeSeriesRocr100200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetTimeSeriesRocr100200Response) SetMeta(v GetTimeSeriesRocr100200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetValues

`func (o *GetTimeSeriesRocr100200Response) GetValues() []GetTimeSeriesRocr100200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetTimeSeriesRocr100200Response) GetValuesOk() (*[]GetTimeSeriesRocr100200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetTimeSeriesRocr100200Response) SetValues(v []GetTimeSeriesRocr100200ResponseValuesInner)`

SetValues sets Values field to given value.


### GetStatus

`func (o *GetTimeSeriesRocr100200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTimeSeriesRocr100200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTimeSeriesRocr100200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


