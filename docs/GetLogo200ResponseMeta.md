# GetLogo200ResponseMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | The ticker symbol of an instrument | 
**Exchange** | Pointer to **string** | The exchange where the instrument is traded (for &#x60;crypto&#x60; only) | [optional] 

## Methods

### NewGetLogo200ResponseMeta

`func NewGetLogo200ResponseMeta(symbol string, ) *GetLogo200ResponseMeta`

NewGetLogo200ResponseMeta instantiates a new GetLogo200ResponseMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLogo200ResponseMetaWithDefaults

`func NewGetLogo200ResponseMetaWithDefaults() *GetLogo200ResponseMeta`

NewGetLogo200ResponseMetaWithDefaults instantiates a new GetLogo200ResponseMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *GetLogo200ResponseMeta) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetLogo200ResponseMeta) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetLogo200ResponseMeta) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetExchange

`func (o *GetLogo200ResponseMeta) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *GetLogo200ResponseMeta) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *GetLogo200ResponseMeta) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *GetLogo200ResponseMeta) HasExchange() bool`

HasExchange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


