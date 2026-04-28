# InlineObject11Meta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | The ticker symbol of an instrument for which data was requested. | 
**Interval** | **string** | The time gap between consecutive data points. | 
**Currency** | **string** | The currency of a traded instrument. | 
**ExchangeTimezone** | **string** | The timezone of the exchange where the instrument is traded. | 
**Exchange** | **string** | The exchange name where the instrument is traded. | 
**MicCode** | **string** | The Market Identifier Code (MIC) of the exchange where the instrument is traded. | 
**Type** | **string** | The asset class to which the instrument belongs. | 
**Indicator** | [**InlineObject11MetaIndicator**](InlineObject11MetaIndicator.md) |  | 

## Methods

### NewInlineObject11Meta

`func NewInlineObject11Meta(symbol string, interval string, currency string, exchangeTimezone string, exchange string, micCode string, type_ string, indicator InlineObject11MetaIndicator, ) *InlineObject11Meta`

NewInlineObject11Meta instantiates a new InlineObject11Meta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject11MetaWithDefaults

`func NewInlineObject11MetaWithDefaults() *InlineObject11Meta`

NewInlineObject11MetaWithDefaults instantiates a new InlineObject11Meta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *InlineObject11Meta) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *InlineObject11Meta) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *InlineObject11Meta) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetInterval

`func (o *InlineObject11Meta) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *InlineObject11Meta) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *InlineObject11Meta) SetInterval(v string)`

SetInterval sets Interval field to given value.


### GetCurrency

`func (o *InlineObject11Meta) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InlineObject11Meta) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InlineObject11Meta) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetExchangeTimezone

`func (o *InlineObject11Meta) GetExchangeTimezone() string`

GetExchangeTimezone returns the ExchangeTimezone field if non-nil, zero value otherwise.

### GetExchangeTimezoneOk

`func (o *InlineObject11Meta) GetExchangeTimezoneOk() (*string, bool)`

GetExchangeTimezoneOk returns a tuple with the ExchangeTimezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeTimezone

`func (o *InlineObject11Meta) SetExchangeTimezone(v string)`

SetExchangeTimezone sets ExchangeTimezone field to given value.


### GetExchange

`func (o *InlineObject11Meta) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *InlineObject11Meta) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *InlineObject11Meta) SetExchange(v string)`

SetExchange sets Exchange field to given value.


### GetMicCode

`func (o *InlineObject11Meta) GetMicCode() string`

GetMicCode returns the MicCode field if non-nil, zero value otherwise.

### GetMicCodeOk

`func (o *InlineObject11Meta) GetMicCodeOk() (*string, bool)`

GetMicCodeOk returns a tuple with the MicCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicCode

`func (o *InlineObject11Meta) SetMicCode(v string)`

SetMicCode sets MicCode field to given value.


### GetType

`func (o *InlineObject11Meta) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InlineObject11Meta) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InlineObject11Meta) SetType(v string)`

SetType sets Type field to given value.


### GetIndicator

`func (o *InlineObject11Meta) GetIndicator() InlineObject11MetaIndicator`

GetIndicator returns the Indicator field if non-nil, zero value otherwise.

### GetIndicatorOk

`func (o *InlineObject11Meta) GetIndicatorOk() (*InlineObject11MetaIndicator, bool)`

GetIndicatorOk returns a tuple with the Indicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndicator

`func (o *InlineObject11Meta) SetIndicator(v InlineObject11MetaIndicator)`

SetIndicator sets Indicator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


