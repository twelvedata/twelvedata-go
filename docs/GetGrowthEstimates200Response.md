# GetGrowthEstimates200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetEarningsEstimate200ResponseMeta**](GetEarningsEstimate200ResponseMeta.md) |  | 
**GrowthEstimates** | Pointer to [**GetGrowthEstimates200ResponseGrowthEstimates**](GetGrowthEstimates200ResponseGrowthEstimates.md) |  | [optional] 
**Status** | **string** | Status of the request | 

## Methods

### NewGetGrowthEstimates200Response

`func NewGetGrowthEstimates200Response(meta GetEarningsEstimate200ResponseMeta, status string, ) *GetGrowthEstimates200Response`

NewGetGrowthEstimates200Response instantiates a new GetGrowthEstimates200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGrowthEstimates200ResponseWithDefaults

`func NewGetGrowthEstimates200ResponseWithDefaults() *GetGrowthEstimates200Response`

NewGetGrowthEstimates200ResponseWithDefaults instantiates a new GetGrowthEstimates200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetGrowthEstimates200Response) GetMeta() GetEarningsEstimate200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetGrowthEstimates200Response) GetMetaOk() (*GetEarningsEstimate200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetGrowthEstimates200Response) SetMeta(v GetEarningsEstimate200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetGrowthEstimates

`func (o *GetGrowthEstimates200Response) GetGrowthEstimates() GetGrowthEstimates200ResponseGrowthEstimates`

GetGrowthEstimates returns the GrowthEstimates field if non-nil, zero value otherwise.

### GetGrowthEstimatesOk

`func (o *GetGrowthEstimates200Response) GetGrowthEstimatesOk() (*GetGrowthEstimates200ResponseGrowthEstimates, bool)`

GetGrowthEstimatesOk returns a tuple with the GrowthEstimates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrowthEstimates

`func (o *GetGrowthEstimates200Response) SetGrowthEstimates(v GetGrowthEstimates200ResponseGrowthEstimates)`

SetGrowthEstimates sets GrowthEstimates field to given value.

### HasGrowthEstimates

`func (o *GetGrowthEstimates200Response) HasGrowthEstimates() bool`

HasGrowthEstimates returns a boolean if a field has been set.

### GetStatus

`func (o *GetGrowthEstimates200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetGrowthEstimates200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetGrowthEstimates200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


