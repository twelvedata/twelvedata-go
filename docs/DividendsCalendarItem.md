# DividendsCalendarItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | Pointer to **string** | Ticker symbol of instrument | [optional] 
**MicCode** | Pointer to **string** | Market identifier code (MIC) under ISO 10383 standard | [optional] 
**Exchange** | Pointer to **string** | Exchange where instrument is traded | [optional] 
**ExDate** | **string** | Stands for the ex date | 
**Amount** | **float64** | Dividend payment amount | 

## Methods

### NewDividendsCalendarItem

`func NewDividendsCalendarItem(exDate string, amount float64, ) *DividendsCalendarItem`

NewDividendsCalendarItem instantiates a new DividendsCalendarItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDividendsCalendarItemWithDefaults

`func NewDividendsCalendarItemWithDefaults() *DividendsCalendarItem`

NewDividendsCalendarItemWithDefaults instantiates a new DividendsCalendarItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *DividendsCalendarItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *DividendsCalendarItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *DividendsCalendarItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *DividendsCalendarItem) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetMicCode

`func (o *DividendsCalendarItem) GetMicCode() string`

GetMicCode returns the MicCode field if non-nil, zero value otherwise.

### GetMicCodeOk

`func (o *DividendsCalendarItem) GetMicCodeOk() (*string, bool)`

GetMicCodeOk returns a tuple with the MicCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicCode

`func (o *DividendsCalendarItem) SetMicCode(v string)`

SetMicCode sets MicCode field to given value.

### HasMicCode

`func (o *DividendsCalendarItem) HasMicCode() bool`

HasMicCode returns a boolean if a field has been set.

### GetExchange

`func (o *DividendsCalendarItem) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *DividendsCalendarItem) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *DividendsCalendarItem) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *DividendsCalendarItem) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetExDate

`func (o *DividendsCalendarItem) GetExDate() string`

GetExDate returns the ExDate field if non-nil, zero value otherwise.

### GetExDateOk

`func (o *DividendsCalendarItem) GetExDateOk() (*string, bool)`

GetExDateOk returns a tuple with the ExDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExDate

`func (o *DividendsCalendarItem) SetExDate(v string)`

SetExDate sets ExDate field to given value.


### GetAmount

`func (o *DividendsCalendarItem) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DividendsCalendarItem) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DividendsCalendarItem) SetAmount(v float64)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


