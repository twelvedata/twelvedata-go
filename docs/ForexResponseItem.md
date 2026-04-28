# ForexResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | Currency pair according to ISO 4217 standard codes with slash(/) delimiter | 
**CurrencyGroup** | **string** | Group to which currency pair belongs to, could be: Major, Minor, Exotic and Exotic-Cross | 
**CurrencyBase** | **string** | Base currency name according to ISO 4217 standard | 
**CurrencyQuote** | **string** | Quote currency name according to ISO 4217 standard | 

## Methods

### NewForexResponseItem

`func NewForexResponseItem(symbol string, currencyGroup string, currencyBase string, currencyQuote string, ) *ForexResponseItem`

NewForexResponseItem instantiates a new ForexResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForexResponseItemWithDefaults

`func NewForexResponseItemWithDefaults() *ForexResponseItem`

NewForexResponseItemWithDefaults instantiates a new ForexResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *ForexResponseItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ForexResponseItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ForexResponseItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetCurrencyGroup

`func (o *ForexResponseItem) GetCurrencyGroup() string`

GetCurrencyGroup returns the CurrencyGroup field if non-nil, zero value otherwise.

### GetCurrencyGroupOk

`func (o *ForexResponseItem) GetCurrencyGroupOk() (*string, bool)`

GetCurrencyGroupOk returns a tuple with the CurrencyGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyGroup

`func (o *ForexResponseItem) SetCurrencyGroup(v string)`

SetCurrencyGroup sets CurrencyGroup field to given value.


### GetCurrencyBase

`func (o *ForexResponseItem) GetCurrencyBase() string`

GetCurrencyBase returns the CurrencyBase field if non-nil, zero value otherwise.

### GetCurrencyBaseOk

`func (o *ForexResponseItem) GetCurrencyBaseOk() (*string, bool)`

GetCurrencyBaseOk returns a tuple with the CurrencyBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyBase

`func (o *ForexResponseItem) SetCurrencyBase(v string)`

SetCurrencyBase sets CurrencyBase field to given value.


### GetCurrencyQuote

`func (o *ForexResponseItem) GetCurrencyQuote() string`

GetCurrencyQuote returns the CurrencyQuote field if non-nil, zero value otherwise.

### GetCurrencyQuoteOk

`func (o *ForexResponseItem) GetCurrencyQuoteOk() (*string, bool)`

GetCurrencyQuoteOk returns a tuple with the CurrencyQuote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyQuote

`func (o *ForexResponseItem) SetCurrencyQuote(v string)`

SetCurrencyQuote sets CurrencyQuote field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


