# GetIpoCalendar200ResponseValueInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | Ticker of the company | 
**Name** | **string** | Name of the company | 
**Exchange** | **string** | Exchange name where the company is listed | 
**MicCode** | **string** | Market Identifier Code (MIC) under ISO 10383 standard | 
**PriceRangeLow** | Pointer to **float64** | The lower bound of stock price range if available | [optional] 
**PriceRangeHigh** | Pointer to **float64** | The upper bound of stock price range if available | [optional] 
**OfferPrice** | Pointer to **float64** | Initial offer price if available | [optional] 
**Currency** | **string** | Currency of the stock | 
**Shares** | Pointer to **int64** | Number of shares, if available | [optional] 

## Methods

### NewGetIpoCalendar200ResponseValueInner

`func NewGetIpoCalendar200ResponseValueInner(symbol string, name string, exchange string, micCode string, currency string, ) *GetIpoCalendar200ResponseValueInner`

NewGetIpoCalendar200ResponseValueInner instantiates a new GetIpoCalendar200ResponseValueInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIpoCalendar200ResponseValueInnerWithDefaults

`func NewGetIpoCalendar200ResponseValueInnerWithDefaults() *GetIpoCalendar200ResponseValueInner`

NewGetIpoCalendar200ResponseValueInnerWithDefaults instantiates a new GetIpoCalendar200ResponseValueInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *GetIpoCalendar200ResponseValueInner) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *GetIpoCalendar200ResponseValueInner) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *GetIpoCalendar200ResponseValueInner) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetName

`func (o *GetIpoCalendar200ResponseValueInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetIpoCalendar200ResponseValueInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetIpoCalendar200ResponseValueInner) SetName(v string)`

SetName sets Name field to given value.


### GetExchange

`func (o *GetIpoCalendar200ResponseValueInner) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *GetIpoCalendar200ResponseValueInner) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *GetIpoCalendar200ResponseValueInner) SetExchange(v string)`

SetExchange sets Exchange field to given value.


### GetMicCode

`func (o *GetIpoCalendar200ResponseValueInner) GetMicCode() string`

GetMicCode returns the MicCode field if non-nil, zero value otherwise.

### GetMicCodeOk

`func (o *GetIpoCalendar200ResponseValueInner) GetMicCodeOk() (*string, bool)`

GetMicCodeOk returns a tuple with the MicCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicCode

`func (o *GetIpoCalendar200ResponseValueInner) SetMicCode(v string)`

SetMicCode sets MicCode field to given value.


### GetPriceRangeLow

`func (o *GetIpoCalendar200ResponseValueInner) GetPriceRangeLow() float64`

GetPriceRangeLow returns the PriceRangeLow field if non-nil, zero value otherwise.

### GetPriceRangeLowOk

`func (o *GetIpoCalendar200ResponseValueInner) GetPriceRangeLowOk() (*float64, bool)`

GetPriceRangeLowOk returns a tuple with the PriceRangeLow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceRangeLow

`func (o *GetIpoCalendar200ResponseValueInner) SetPriceRangeLow(v float64)`

SetPriceRangeLow sets PriceRangeLow field to given value.

### HasPriceRangeLow

`func (o *GetIpoCalendar200ResponseValueInner) HasPriceRangeLow() bool`

HasPriceRangeLow returns a boolean if a field has been set.

### GetPriceRangeHigh

`func (o *GetIpoCalendar200ResponseValueInner) GetPriceRangeHigh() float64`

GetPriceRangeHigh returns the PriceRangeHigh field if non-nil, zero value otherwise.

### GetPriceRangeHighOk

`func (o *GetIpoCalendar200ResponseValueInner) GetPriceRangeHighOk() (*float64, bool)`

GetPriceRangeHighOk returns a tuple with the PriceRangeHigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceRangeHigh

`func (o *GetIpoCalendar200ResponseValueInner) SetPriceRangeHigh(v float64)`

SetPriceRangeHigh sets PriceRangeHigh field to given value.

### HasPriceRangeHigh

`func (o *GetIpoCalendar200ResponseValueInner) HasPriceRangeHigh() bool`

HasPriceRangeHigh returns a boolean if a field has been set.

### GetOfferPrice

`func (o *GetIpoCalendar200ResponseValueInner) GetOfferPrice() float64`

GetOfferPrice returns the OfferPrice field if non-nil, zero value otherwise.

### GetOfferPriceOk

`func (o *GetIpoCalendar200ResponseValueInner) GetOfferPriceOk() (*float64, bool)`

GetOfferPriceOk returns a tuple with the OfferPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferPrice

`func (o *GetIpoCalendar200ResponseValueInner) SetOfferPrice(v float64)`

SetOfferPrice sets OfferPrice field to given value.

### HasOfferPrice

`func (o *GetIpoCalendar200ResponseValueInner) HasOfferPrice() bool`

HasOfferPrice returns a boolean if a field has been set.

### GetCurrency

`func (o *GetIpoCalendar200ResponseValueInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetIpoCalendar200ResponseValueInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetIpoCalendar200ResponseValueInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetShares

`func (o *GetIpoCalendar200ResponseValueInner) GetShares() int64`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *GetIpoCalendar200ResponseValueInner) GetSharesOk() (*int64, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *GetIpoCalendar200ResponseValueInner) SetShares(v int64)`

SetShares sets Shares field to given value.

### HasShares

`func (o *GetIpoCalendar200ResponseValueInner) HasShares() bool`

HasShares returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


