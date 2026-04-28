# BondResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | Bond symbol | 
**Name** | **string** | Full name of the bond | 
**Country** | **string** | Country where the bond is located | 
**Currency** | **string** | Currency of the bond according to the ISO 4217 standard | 
**Exchange** | **string** | Exchange where the bond is traded | 
**MicCode** | **string** | Market identifier code (MIC) under ISO 10383 standard | 
**Type** | **string** | Type of the bond | 
**Access** | Pointer to [**BondsResponseItemAccess**](BondsResponseItemAccess.md) |  | [optional] 

## Methods

### NewBondResponseItem

`func NewBondResponseItem(symbol string, name string, country string, currency string, exchange string, micCode string, type_ string, ) *BondResponseItem`

NewBondResponseItem instantiates a new BondResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBondResponseItemWithDefaults

`func NewBondResponseItemWithDefaults() *BondResponseItem`

NewBondResponseItemWithDefaults instantiates a new BondResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *BondResponseItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *BondResponseItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *BondResponseItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetName

`func (o *BondResponseItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BondResponseItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BondResponseItem) SetName(v string)`

SetName sets Name field to given value.


### GetCountry

`func (o *BondResponseItem) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *BondResponseItem) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *BondResponseItem) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetCurrency

`func (o *BondResponseItem) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BondResponseItem) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BondResponseItem) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetExchange

`func (o *BondResponseItem) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *BondResponseItem) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *BondResponseItem) SetExchange(v string)`

SetExchange sets Exchange field to given value.


### GetMicCode

`func (o *BondResponseItem) GetMicCode() string`

GetMicCode returns the MicCode field if non-nil, zero value otherwise.

### GetMicCodeOk

`func (o *BondResponseItem) GetMicCodeOk() (*string, bool)`

GetMicCodeOk returns a tuple with the MicCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicCode

`func (o *BondResponseItem) SetMicCode(v string)`

SetMicCode sets MicCode field to given value.


### GetType

`func (o *BondResponseItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BondResponseItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BondResponseItem) SetType(v string)`

SetType sets Type field to given value.


### GetAccess

`func (o *BondResponseItem) GetAccess() BondsResponseItemAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *BondResponseItem) GetAccessOk() (*BondsResponseItemAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *BondResponseItem) SetAccess(v BondsResponseItemAccess)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *BondResponseItem) HasAccess() bool`

HasAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


