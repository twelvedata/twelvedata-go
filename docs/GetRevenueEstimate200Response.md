# GetRevenueEstimate200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetEarningsEstimate200ResponseMeta**](GetEarningsEstimate200ResponseMeta.md) |  | 
**RevenueEstimate** | [**[]GetRevenueEstimate200ResponseRevenueEstimateInner**](GetRevenueEstimate200ResponseRevenueEstimateInner.md) | Revenue estimate data | 
**Status** | **string** | Status of the response | 

## Methods

### NewGetRevenueEstimate200Response

`func NewGetRevenueEstimate200Response(meta GetEarningsEstimate200ResponseMeta, revenueEstimate []GetRevenueEstimate200ResponseRevenueEstimateInner, status string, ) *GetRevenueEstimate200Response`

NewGetRevenueEstimate200Response instantiates a new GetRevenueEstimate200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRevenueEstimate200ResponseWithDefaults

`func NewGetRevenueEstimate200ResponseWithDefaults() *GetRevenueEstimate200Response`

NewGetRevenueEstimate200ResponseWithDefaults instantiates a new GetRevenueEstimate200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetRevenueEstimate200Response) GetMeta() GetEarningsEstimate200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetRevenueEstimate200Response) GetMetaOk() (*GetEarningsEstimate200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetRevenueEstimate200Response) SetMeta(v GetEarningsEstimate200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetRevenueEstimate

`func (o *GetRevenueEstimate200Response) GetRevenueEstimate() []GetRevenueEstimate200ResponseRevenueEstimateInner`

GetRevenueEstimate returns the RevenueEstimate field if non-nil, zero value otherwise.

### GetRevenueEstimateOk

`func (o *GetRevenueEstimate200Response) GetRevenueEstimateOk() (*[]GetRevenueEstimate200ResponseRevenueEstimateInner, bool)`

GetRevenueEstimateOk returns a tuple with the RevenueEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueEstimate

`func (o *GetRevenueEstimate200Response) SetRevenueEstimate(v []GetRevenueEstimate200ResponseRevenueEstimateInner)`

SetRevenueEstimate sets RevenueEstimate field to given value.


### GetStatus

`func (o *GetRevenueEstimate200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetRevenueEstimate200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetRevenueEstimate200Response) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


