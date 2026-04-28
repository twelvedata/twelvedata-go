# GetEpsTrend200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetEarningsEstimate200ResponseMeta**](GetEarningsEstimate200ResponseMeta.md) |  | 
**EpsTrend** | [**[]GetEpsTrend200ResponseEpsTrendInner**](GetEpsTrend200ResponseEpsTrendInner.md) | EPS trend data | 
**Status** | **string** | Status of the response | 

## Methods

### NewGetEpsTrend200Response

`func NewGetEpsTrend200Response(meta GetEarningsEstimate200ResponseMeta, epsTrend []GetEpsTrend200ResponseEpsTrendInner, status string, ) *GetEpsTrend200Response`

NewGetEpsTrend200Response instantiates a new GetEpsTrend200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEpsTrend200ResponseWithDefaults

`func NewGetEpsTrend200ResponseWithDefaults() *GetEpsTrend200Response`

NewGetEpsTrend200ResponseWithDefaults instantiates a new GetEpsTrend200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetEpsTrend200Response) GetMeta() GetEarningsEstimate200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetEpsTrend200Response) GetMetaOk() (*GetEarningsEstimate200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetEpsTrend200Response) SetMeta(v GetEarningsEstimate200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetEpsTrend

`func (o *GetEpsTrend200Response) GetEpsTrend() []GetEpsTrend200ResponseEpsTrendInner`

GetEpsTrend returns the EpsTrend field if non-nil, zero value otherwise.

### GetEpsTrendOk

`func (o *GetEpsTrend200Response) GetEpsTrendOk() (*[]GetEpsTrend200ResponseEpsTrendInner, bool)`

GetEpsTrendOk returns a tuple with the EpsTrend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpsTrend

`func (o *GetEpsTrend200Response) SetEpsTrend(v []GetEpsTrend200ResponseEpsTrendInner)`

SetEpsTrend sets EpsTrend field to given value.


### GetStatus

`func (o *GetEpsTrend200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetEpsTrend200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetEpsTrend200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


