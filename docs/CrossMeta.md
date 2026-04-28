# CrossMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseInstrument** | **string** | Base instrument symbol | 
**BaseCurrency** | **string** | Base currency | 
**BaseExchange** | **string** | Base exchange | 
**Interval** | **string** | Interval between two consecutive points in time series | 
**QuoteInstrument** | **string** | Quote instrument symbol | 
**QuoteCurrency** | **string** | Quote currency | 
**QuoteExchange** | **string** | Quote exchange | 

## Methods

### NewCrossMeta

`func NewCrossMeta(baseInstrument string, baseCurrency string, baseExchange string, interval string, quoteInstrument string, quoteCurrency string, quoteExchange string, ) *CrossMeta`

NewCrossMeta instantiates a new CrossMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossMetaWithDefaults

`func NewCrossMetaWithDefaults() *CrossMeta`

NewCrossMetaWithDefaults instantiates a new CrossMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseInstrument

`func (o *CrossMeta) GetBaseInstrument() string`

GetBaseInstrument returns the BaseInstrument field if non-nil, zero value otherwise.

### GetBaseInstrumentOk

`func (o *CrossMeta) GetBaseInstrumentOk() (*string, bool)`

GetBaseInstrumentOk returns a tuple with the BaseInstrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseInstrument

`func (o *CrossMeta) SetBaseInstrument(v string)`

SetBaseInstrument sets BaseInstrument field to given value.


### GetBaseCurrency

`func (o *CrossMeta) GetBaseCurrency() string`

GetBaseCurrency returns the BaseCurrency field if non-nil, zero value otherwise.

### GetBaseCurrencyOk

`func (o *CrossMeta) GetBaseCurrencyOk() (*string, bool)`

GetBaseCurrencyOk returns a tuple with the BaseCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCurrency

`func (o *CrossMeta) SetBaseCurrency(v string)`

SetBaseCurrency sets BaseCurrency field to given value.


### GetBaseExchange

`func (o *CrossMeta) GetBaseExchange() string`

GetBaseExchange returns the BaseExchange field if non-nil, zero value otherwise.

### GetBaseExchangeOk

`func (o *CrossMeta) GetBaseExchangeOk() (*string, bool)`

GetBaseExchangeOk returns a tuple with the BaseExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseExchange

`func (o *CrossMeta) SetBaseExchange(v string)`

SetBaseExchange sets BaseExchange field to given value.


### GetInterval

`func (o *CrossMeta) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *CrossMeta) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *CrossMeta) SetInterval(v string)`

SetInterval sets Interval field to given value.


### GetQuoteInstrument

`func (o *CrossMeta) GetQuoteInstrument() string`

GetQuoteInstrument returns the QuoteInstrument field if non-nil, zero value otherwise.

### GetQuoteInstrumentOk

`func (o *CrossMeta) GetQuoteInstrumentOk() (*string, bool)`

GetQuoteInstrumentOk returns a tuple with the QuoteInstrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteInstrument

`func (o *CrossMeta) SetQuoteInstrument(v string)`

SetQuoteInstrument sets QuoteInstrument field to given value.


### GetQuoteCurrency

`func (o *CrossMeta) GetQuoteCurrency() string`

GetQuoteCurrency returns the QuoteCurrency field if non-nil, zero value otherwise.

### GetQuoteCurrencyOk

`func (o *CrossMeta) GetQuoteCurrencyOk() (*string, bool)`

GetQuoteCurrencyOk returns a tuple with the QuoteCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteCurrency

`func (o *CrossMeta) SetQuoteCurrency(v string)`

SetQuoteCurrency sets QuoteCurrency field to given value.


### GetQuoteExchange

`func (o *CrossMeta) GetQuoteExchange() string`

GetQuoteExchange returns the QuoteExchange field if non-nil, zero value otherwise.

### GetQuoteExchangeOk

`func (o *CrossMeta) GetQuoteExchangeOk() (*string, bool)`

GetQuoteExchangeOk returns a tuple with the QuoteExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteExchange

`func (o *CrossMeta) SetQuoteExchange(v string)`

SetQuoteExchange sets QuoteExchange field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


