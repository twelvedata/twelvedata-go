# GetBalanceSheet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetBalanceSheet200ResponseMeta**](GetBalanceSheet200ResponseMeta.md) |  | 
**BalanceSheet** | [**[]GetBalanceSheet200ResponseBalanceSheetInner**](GetBalanceSheet200ResponseBalanceSheetInner.md) | Array of balance sheet records | 

## Methods

### NewGetBalanceSheet200Response

`func NewGetBalanceSheet200Response(meta GetBalanceSheet200ResponseMeta, balanceSheet []GetBalanceSheet200ResponseBalanceSheetInner, ) *GetBalanceSheet200Response`

NewGetBalanceSheet200Response instantiates a new GetBalanceSheet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBalanceSheet200ResponseWithDefaults

`func NewGetBalanceSheet200ResponseWithDefaults() *GetBalanceSheet200Response`

NewGetBalanceSheet200ResponseWithDefaults instantiates a new GetBalanceSheet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetBalanceSheet200Response) GetMeta() GetBalanceSheet200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetBalanceSheet200Response) GetMetaOk() (*GetBalanceSheet200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetBalanceSheet200Response) SetMeta(v GetBalanceSheet200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetBalanceSheet

`func (o *GetBalanceSheet200Response) GetBalanceSheet() []GetBalanceSheet200ResponseBalanceSheetInner`

GetBalanceSheet returns the BalanceSheet field if non-nil, zero value otherwise.

### GetBalanceSheetOk

`func (o *GetBalanceSheet200Response) GetBalanceSheetOk() (*[]GetBalanceSheet200ResponseBalanceSheetInner, bool)`

GetBalanceSheetOk returns a tuple with the BalanceSheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceSheet

`func (o *GetBalanceSheet200Response) SetBalanceSheet(v []GetBalanceSheet200ResponseBalanceSheetInner)`

SetBalanceSheet sets BalanceSheet field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


