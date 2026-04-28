# GetEarningsEstimate200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetEarningsEstimate200ResponseMeta**](GetEarningsEstimate200ResponseMeta.md) |  | 
**EarningsEstimate** | [**[]GetEarningsEstimate200ResponseEarningsEstimateInner**](GetEarningsEstimate200ResponseEarningsEstimateInner.md) | List of earnings estimates | 
**Status** | **string** | Response status | 

## Methods

### NewGetEarningsEstimate200Response

`func NewGetEarningsEstimate200Response(meta GetEarningsEstimate200ResponseMeta, earningsEstimate []GetEarningsEstimate200ResponseEarningsEstimateInner, status string, ) *GetEarningsEstimate200Response`

NewGetEarningsEstimate200Response instantiates a new GetEarningsEstimate200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEarningsEstimate200ResponseWithDefaults

`func NewGetEarningsEstimate200ResponseWithDefaults() *GetEarningsEstimate200Response`

NewGetEarningsEstimate200ResponseWithDefaults instantiates a new GetEarningsEstimate200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetEarningsEstimate200Response) GetMeta() GetEarningsEstimate200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetEarningsEstimate200Response) GetMetaOk() (*GetEarningsEstimate200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetEarningsEstimate200Response) SetMeta(v GetEarningsEstimate200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetEarningsEstimate

`func (o *GetEarningsEstimate200Response) GetEarningsEstimate() []GetEarningsEstimate200ResponseEarningsEstimateInner`

GetEarningsEstimate returns the EarningsEstimate field if non-nil, zero value otherwise.

### GetEarningsEstimateOk

`func (o *GetEarningsEstimate200Response) GetEarningsEstimateOk() (*[]GetEarningsEstimate200ResponseEarningsEstimateInner, bool)`

GetEarningsEstimateOk returns a tuple with the EarningsEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarningsEstimate

`func (o *GetEarningsEstimate200Response) SetEarningsEstimate(v []GetEarningsEstimate200ResponseEarningsEstimateInner)`

SetEarningsEstimate sets EarningsEstimate field to given value.


### GetStatus

`func (o *GetEarningsEstimate200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetEarningsEstimate200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetEarningsEstimate200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


