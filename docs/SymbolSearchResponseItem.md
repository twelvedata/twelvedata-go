# SymbolSearchResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | Ticker symbol of instrument | 
**InstrumentName** | **string** | Name of exchange | 
**Exchange** | **string** | Exchange where instrument is traded | 
**MicCode** | **string** | Market identifier code (MIC) under ISO 10383 standard | 
**ExchangeTimezone** | **string** | Time zone where exchange is located | 
**InstrumentType** | **string** | Type of instrument | 
**Country** | **string** | Country to which stock exchange belongs to | 
**Currency** | **string** | Currency in which the instrument is traded | 
**Access** | Pointer to [**SymbolSearchResponseItemAccess**](SymbolSearchResponseItemAccess.md) |  | [optional] 

## Methods

### NewSymbolSearchResponseItem

`func NewSymbolSearchResponseItem(symbol string, instrumentName string, exchange string, micCode string, exchangeTimezone string, instrumentType string, country string, currency string, ) *SymbolSearchResponseItem`

NewSymbolSearchResponseItem instantiates a new SymbolSearchResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSymbolSearchResponseItemWithDefaults

`func NewSymbolSearchResponseItemWithDefaults() *SymbolSearchResponseItem`

NewSymbolSearchResponseItemWithDefaults instantiates a new SymbolSearchResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *SymbolSearchResponseItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *SymbolSearchResponseItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *SymbolSearchResponseItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetInstrumentName

`func (o *SymbolSearchResponseItem) GetInstrumentName() string`

GetInstrumentName returns the InstrumentName field if non-nil, zero value otherwise.

### GetInstrumentNameOk

`func (o *SymbolSearchResponseItem) GetInstrumentNameOk() (*string, bool)`

GetInstrumentNameOk returns a tuple with the InstrumentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentName

`func (o *SymbolSearchResponseItem) SetInstrumentName(v string)`

SetInstrumentName sets InstrumentName field to given value.


### GetExchange

`func (o *SymbolSearchResponseItem) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *SymbolSearchResponseItem) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *SymbolSearchResponseItem) SetExchange(v string)`

SetExchange sets Exchange field to given value.


### GetMicCode

`func (o *SymbolSearchResponseItem) GetMicCode() string`

GetMicCode returns the MicCode field if non-nil, zero value otherwise.

### GetMicCodeOk

`func (o *SymbolSearchResponseItem) GetMicCodeOk() (*string, bool)`

GetMicCodeOk returns a tuple with the MicCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicCode

`func (o *SymbolSearchResponseItem) SetMicCode(v string)`

SetMicCode sets MicCode field to given value.


### GetExchangeTimezone

`func (o *SymbolSearchResponseItem) GetExchangeTimezone() string`

GetExchangeTimezone returns the ExchangeTimezone field if non-nil, zero value otherwise.

### GetExchangeTimezoneOk

`func (o *SymbolSearchResponseItem) GetExchangeTimezoneOk() (*string, bool)`

GetExchangeTimezoneOk returns a tuple with the ExchangeTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeTimezone

`func (o *SymbolSearchResponseItem) SetExchangeTimezone(v string)`

SetExchangeTimezone sets ExchangeTimezone field to given value.


### GetInstrumentType

`func (o *SymbolSearchResponseItem) GetInstrumentType() string`

GetInstrumentType returns the InstrumentType field if non-nil, zero value otherwise.

### GetInstrumentTypeOk

`func (o *SymbolSearchResponseItem) GetInstrumentTypeOk() (*string, bool)`

GetInstrumentTypeOk returns a tuple with the InstrumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentType

`func (o *SymbolSearchResponseItem) SetInstrumentType(v string)`

SetInstrumentType sets InstrumentType field to given value.


### GetCountry

`func (o *SymbolSearchResponseItem) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *SymbolSearchResponseItem) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *SymbolSearchResponseItem) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetCurrency

`func (o *SymbolSearchResponseItem) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SymbolSearchResponseItem) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SymbolSearchResponseItem) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetAccess

`func (o *SymbolSearchResponseItem) GetAccess() SymbolSearchResponseItemAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *SymbolSearchResponseItem) GetAccessOk() (*SymbolSearchResponseItemAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *SymbolSearchResponseItem) SetAccess(v SymbolSearchResponseItemAccess)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *SymbolSearchResponseItem) HasAccess() bool`

HasAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


