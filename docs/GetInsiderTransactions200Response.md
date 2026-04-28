# GetInsiderTransactions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetInsiderTransactions200ResponseMeta**](GetInsiderTransactions200ResponseMeta.md) |  | 
**InsiderTransactions** | [**[]GetInsiderTransactions200ResponseInsiderTransactionsInner**](GetInsiderTransactions200ResponseInsiderTransactionsInner.md) | List of insider transactions | 

## Methods

### NewGetInsiderTransactions200Response

`func NewGetInsiderTransactions200Response(meta GetInsiderTransactions200ResponseMeta, insiderTransactions []GetInsiderTransactions200ResponseInsiderTransactionsInner, ) *GetInsiderTransactions200Response`

NewGetInsiderTransactions200Response instantiates a new GetInsiderTransactions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInsiderTransactions200ResponseWithDefaults

`func NewGetInsiderTransactions200ResponseWithDefaults() *GetInsiderTransactions200Response`

NewGetInsiderTransactions200ResponseWithDefaults instantiates a new GetInsiderTransactions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetInsiderTransactions200Response) GetMeta() GetInsiderTransactions200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetInsiderTransactions200Response) GetMetaOk() (*GetInsiderTransactions200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetInsiderTransactions200Response) SetMeta(v GetInsiderTransactions200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetInsiderTransactions

`func (o *GetInsiderTransactions200Response) GetInsiderTransactions() []GetInsiderTransactions200ResponseInsiderTransactionsInner`

GetInsiderTransactions returns the InsiderTransactions field if non-nil, zero value otherwise.

### GetInsiderTransactionsOk

`func (o *GetInsiderTransactions200Response) GetInsiderTransactionsOk() (*[]GetInsiderTransactions200ResponseInsiderTransactionsInner, bool)`

GetInsiderTransactionsOk returns a tuple with the InsiderTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsiderTransactions

`func (o *GetInsiderTransactions200Response) SetInsiderTransactions(v []GetInsiderTransactions200ResponseInsiderTransactionsInner)`

SetInsiderTransactions sets InsiderTransactions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


