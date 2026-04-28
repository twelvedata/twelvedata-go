# StocksResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | Instrument symbol (ticker) | 
**Name** | **string** | Full name of instrument | 
**Currency** | **string** | Currency of the instrument according to the ISO 4217 standard | 
**Exchange** | **string** | Exchange where instrument is traded | 
**MicCode** | **string** | Market identifier code (MIC) under ISO 10383 standard | 
**Country** | **string** | Country where exchange is located | 
**Type** | **string** | Common issue type | 
**FigiCode** | **string** | Financial instrument global identifier (FIGI) | 
**CfiCode** | **string** | Classification of Financial Instruments (CFI) | 
**Isin** | **string** | International securities identification number (ISIN), available by individual request to support | 
**Cusip** | **string** | A unique nine-character alphanumeric code used to identify financial securities, ensuring accurate data retrieval for the specified asset | 
**Access** | Pointer to [**EtfResponseItemAccess**](EtfResponseItemAccess.md) |  | [optional] 

## Methods

### NewStocksResponseItem

`func NewStocksResponseItem(symbol string, name string, currency string, exchange string, micCode string, country string, type_ string, figiCode string, cfiCode string, isin string, cusip string, ) *StocksResponseItem`

NewStocksResponseItem instantiates a new StocksResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStocksResponseItemWithDefaults

`func NewStocksResponseItemWithDefaults() *StocksResponseItem`

NewStocksResponseItemWithDefaults instantiates a new StocksResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *StocksResponseItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *StocksResponseItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *StocksResponseItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetName

`func (o *StocksResponseItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StocksResponseItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StocksResponseItem) SetName(v string)`

SetName sets Name field to given value.


### GetCurrency

`func (o *StocksResponseItem) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *StocksResponseItem) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *StocksResponseItem) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetExchange

`func (o *StocksResponseItem) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *StocksResponseItem) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *StocksResponseItem) SetExchange(v string)`

SetExchange sets Exchange field to given value.


### GetMicCode

`func (o *StocksResponseItem) GetMicCode() string`

GetMicCode returns the MicCode field if non-nil, zero value otherwise.

### GetMicCodeOk

`func (o *StocksResponseItem) GetMicCodeOk() (*string, bool)`

GetMicCodeOk returns a tuple with the MicCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicCode

`func (o *StocksResponseItem) SetMicCode(v string)`

SetMicCode sets MicCode field to given value.


### GetCountry

`func (o *StocksResponseItem) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *StocksResponseItem) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *StocksResponseItem) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetType

`func (o *StocksResponseItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StocksResponseItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StocksResponseItem) SetType(v string)`

SetType sets Type field to given value.


### GetFigiCode

`func (o *StocksResponseItem) GetFigiCode() string`

GetFigiCode returns the FigiCode field if non-nil, zero value otherwise.

### GetFigiCodeOk

`func (o *StocksResponseItem) GetFigiCodeOk() (*string, bool)`

GetFigiCodeOk returns a tuple with the FigiCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFigiCode

`func (o *StocksResponseItem) SetFigiCode(v string)`

SetFigiCode sets FigiCode field to given value.


### GetCfiCode

`func (o *StocksResponseItem) GetCfiCode() string`

GetCfiCode returns the CfiCode field if non-nil, zero value otherwise.

### GetCfiCodeOk

`func (o *StocksResponseItem) GetCfiCodeOk() (*string, bool)`

GetCfiCodeOk returns a tuple with the CfiCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfiCode

`func (o *StocksResponseItem) SetCfiCode(v string)`

SetCfiCode sets CfiCode field to given value.


### GetIsin

`func (o *StocksResponseItem) GetIsin() string`

GetIsin returns the Isin field if non-nil, zero value otherwise.

### GetIsinOk

`func (o *StocksResponseItem) GetIsinOk() (*string, bool)`

GetIsinOk returns a tuple with the Isin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsin

`func (o *StocksResponseItem) SetIsin(v string)`

SetIsin sets Isin field to given value.


### GetCusip

`func (o *StocksResponseItem) GetCusip() string`

GetCusip returns the Cusip field if non-nil, zero value otherwise.

### GetCusipOk

`func (o *StocksResponseItem) GetCusipOk() (*string, bool)`

GetCusipOk returns a tuple with the Cusip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCusip

`func (o *StocksResponseItem) SetCusip(v string)`

SetCusip sets Cusip field to given value.


### GetAccess

`func (o *StocksResponseItem) GetAccess() EtfResponseItemAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *StocksResponseItem) GetAccessOk() (*EtfResponseItemAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *StocksResponseItem) SetAccess(v EtfResponseItemAccess)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *StocksResponseItem) HasAccess() bool`

HasAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


