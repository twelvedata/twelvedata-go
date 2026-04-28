# CryptocurrencyResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | Cryptocurrency pair codes with slash(/) delimiter | 
**AvailableExchanges** | **[]string** | List of exchanges where the cryptocurrency is available | 
**CurrencyBase** | **string** | Base currency of the cryptocurrency pair | 
**CurrencyQuote** | **string** | Quote currency of the cryptocurrency pair | 

## Methods

### NewCryptocurrencyResponseItem

`func NewCryptocurrencyResponseItem(symbol string, availableExchanges []string, currencyBase string, currencyQuote string, ) *CryptocurrencyResponseItem`

NewCryptocurrencyResponseItem instantiates a new CryptocurrencyResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCryptocurrencyResponseItemWithDefaults

`func NewCryptocurrencyResponseItemWithDefaults() *CryptocurrencyResponseItem`

NewCryptocurrencyResponseItemWithDefaults instantiates a new CryptocurrencyResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *CryptocurrencyResponseItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *CryptocurrencyResponseItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *CryptocurrencyResponseItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetAvailableExchanges

`func (o *CryptocurrencyResponseItem) GetAvailableExchanges() []string`

GetAvailableExchanges returns the AvailableExchanges field if non-nil, zero value otherwise.

### GetAvailableExchangesOk

`func (o *CryptocurrencyResponseItem) GetAvailableExchangesOk() (*[]string, bool)`

GetAvailableExchangesOk returns a tuple with the AvailableExchanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableExchanges

`func (o *CryptocurrencyResponseItem) SetAvailableExchanges(v []string)`

SetAvailableExchanges sets AvailableExchanges field to given value.


### GetCurrencyBase

`func (o *CryptocurrencyResponseItem) GetCurrencyBase() string`

GetCurrencyBase returns the CurrencyBase field if non-nil, zero value otherwise.

### GetCurrencyBaseOk

`func (o *CryptocurrencyResponseItem) GetCurrencyBaseOk() (*string, bool)`

GetCurrencyBaseOk returns a tuple with the CurrencyBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyBase

`func (o *CryptocurrencyResponseItem) SetCurrencyBase(v string)`

SetCurrencyBase sets CurrencyBase field to given value.


### GetCurrencyQuote

`func (o *CryptocurrencyResponseItem) GetCurrencyQuote() string`

GetCurrencyQuote returns the CurrencyQuote field if non-nil, zero value otherwise.

### GetCurrencyQuoteOk

`func (o *CryptocurrencyResponseItem) GetCurrencyQuoteOk() (*string, bool)`

GetCurrencyQuoteOk returns a tuple with the CurrencyQuote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyQuote

`func (o *CryptocurrencyResponseItem) SetCurrencyQuote(v string)`

SetCurrencyQuote sets CurrencyQuote field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


