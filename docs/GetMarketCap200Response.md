# GetMarketCap200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | [**GetMarketCap200ResponseMeta**](GetMarketCap200ResponseMeta.md) |  | 
**MarketCap** | [**[]GetMarketCap200ResponseMarketCapInner**](GetMarketCap200ResponseMarketCapInner.md) | Market capitalization values | 

## Methods

### NewGetMarketCap200Response

`func NewGetMarketCap200Response(meta GetMarketCap200ResponseMeta, marketCap []GetMarketCap200ResponseMarketCapInner, ) *GetMarketCap200Response`

NewGetMarketCap200Response instantiates a new GetMarketCap200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMarketCap200ResponseWithDefaults

`func NewGetMarketCap200ResponseWithDefaults() *GetMarketCap200Response`

NewGetMarketCap200ResponseWithDefaults instantiates a new GetMarketCap200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *GetMarketCap200Response) GetMeta() GetMarketCap200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetMarketCap200Response) GetMetaOk() (*GetMarketCap200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetMarketCap200Response) SetMeta(v GetMarketCap200ResponseMeta)`

SetMeta sets Meta field to given value.


### GetMarketCap

`func (o *GetMarketCap200Response) GetMarketCap() []GetMarketCap200ResponseMarketCapInner`

GetMarketCap returns the MarketCap field if non-nil, zero value otherwise.

### GetMarketCapOk

`func (o *GetMarketCap200Response) GetMarketCapOk() (*[]GetMarketCap200ResponseMarketCapInner, bool)`

GetMarketCapOk returns a tuple with the MarketCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketCap

`func (o *GetMarketCap200Response) SetMarketCap(v []GetMarketCap200ResponseMarketCapInner)`

SetMarketCap sets MarketCap field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


