# GetFundHolders200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetFundHolders200ResponseMeta**](GetFundHolders200ResponseMeta.md) |  | 
**FundHolders** | [**[]FundHolderItem**](FundHolderItem.md) | List of fund holders for the financial instrument | 

## Methods

### NewGetFundHolders200Response

`func NewGetFundHolders200Response(meta GetFundHolders200ResponseMeta, fundHolders []FundHolderItem, ) *GetFundHolders200Response`

NewGetFundHolders200Response instantiates a new GetFundHolders200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFundHolders200ResponseWithDefaults

`func NewGetFundHolders200ResponseWithDefaults() *GetFundHolders200Response`

NewGetFundHolders200ResponseWithDefaults instantiates a new GetFundHolders200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetFundHolders200Response) GetMeta() GetFundHolders200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetFundHolders200Response) GetMetaOk() (*GetFundHolders200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetFundHolders200Response) SetMeta(v GetFundHolders200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetFundHolders

`func (o *GetFundHolders200Response) GetFundHolders() []FundHolderItem`

GetFundHolders returns the FundHolders field if non-nil, zero value otherwise.

### GetFundHoldersOk

`func (o *GetFundHolders200Response) GetFundHoldersOk() (*[]FundHolderItem, bool)`

GetFundHoldersOk returns a tuple with the FundHolders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundHolders

`func (o *GetFundHolders200Response) SetFundHolders(v []FundHolderItem)`

SetFundHolders sets FundHolders field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


