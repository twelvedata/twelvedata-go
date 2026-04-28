# GetCashFlow200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetCashFlow200ResponseMeta**](GetCashFlow200ResponseMeta.md) |  | 
**CashFlow** | [**[]CashFlowStruct**](CashFlowStruct.md) | Cash flow data | 

## Methods

### NewGetCashFlow200Response

`func NewGetCashFlow200Response(meta GetCashFlow200ResponseMeta, cashFlow []CashFlowStruct, ) *GetCashFlow200Response`

NewGetCashFlow200Response instantiates a new GetCashFlow200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCashFlow200ResponseWithDefaults

`func NewGetCashFlow200ResponseWithDefaults() *GetCashFlow200Response`

NewGetCashFlow200ResponseWithDefaults instantiates a new GetCashFlow200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetCashFlow200Response) GetMeta() GetCashFlow200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetCashFlow200Response) GetMetaOk() (*GetCashFlow200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetCashFlow200Response) SetMeta(v GetCashFlow200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetCashFlow

`func (o *GetCashFlow200Response) GetCashFlow() []CashFlowStruct`

GetCashFlow returns the CashFlow field if non-nil, zero value otherwise.

### GetCashFlowOk

`func (o *GetCashFlow200Response) GetCashFlowOk() (*[]CashFlowStruct, bool)`

GetCashFlowOk returns a tuple with the CashFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashFlow

`func (o *GetCashFlow200Response) SetCashFlow(v []CashFlowStruct)`

SetCashFlow sets CashFlow field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


