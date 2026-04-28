# SplitsCalendarResponseItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** | Stands for the split date | 
**Symbol** | Pointer to **string** | Ticker of the company | [optional] 
**MicCode** | Pointer to **string** | Market Identifier Code (MIC) under ISO 10383 standard | [optional] 
**Exchange** | Pointer to **string** | Exchange name where the company is listed | [optional] 
**Description** | **string** | Specification of the split event | 
**Ratio** | **float64** | The ratio by which the number of a company&#39;s outstanding shares of stock are increased following a stock split. For example, a &#x60;4-for-1 split&#x60; results in four times as many outstanding shares, with each share selling at one forth of its pre-split price | 
**FromFactor** | **float64** | From factor of the split | 
**ToFactor** | **float64** | To factor of the split | 

## Methods

### NewSplitsCalendarResponseItem

`func NewSplitsCalendarResponseItem(date string, description string, ratio float64, fromFactor float64, toFactor float64, ) *SplitsCalendarResponseItem`

NewSplitsCalendarResponseItem instantiates a new SplitsCalendarResponseItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSplitsCalendarResponseItemWithDefaults

`func NewSplitsCalendarResponseItemWithDefaults() *SplitsCalendarResponseItem`

NewSplitsCalendarResponseItemWithDefaults instantiates a new SplitsCalendarResponseItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *SplitsCalendarResponseItem) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *SplitsCalendarResponseItem) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *SplitsCalendarResponseItem) SetDate(v string)`

SetDate sets Date field to given value.


### GetSymbol

`func (o *SplitsCalendarResponseItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *SplitsCalendarResponseItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *SplitsCalendarResponseItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *SplitsCalendarResponseItem) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetMicCode

`func (o *SplitsCalendarResponseItem) GetMicCode() string`

GetMicCode returns the MicCode field if non-nil, zero value otherwise.

### GetMicCodeOk

`func (o *SplitsCalendarResponseItem) GetMicCodeOk() (*string, bool)`

GetMicCodeOk returns a tuple with the MicCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicCode

`func (o *SplitsCalendarResponseItem) SetMicCode(v string)`

SetMicCode sets MicCode field to given value.

### HasMicCode

`func (o *SplitsCalendarResponseItem) HasMicCode() bool`

HasMicCode returns a boolean if a field has been set.

### GetExchange

`func (o *SplitsCalendarResponseItem) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *SplitsCalendarResponseItem) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *SplitsCalendarResponseItem) SetExchange(v string)`

SetExchange sets Exchange field to given value.

### HasExchange

`func (o *SplitsCalendarResponseItem) HasExchange() bool`

HasExchange returns a boolean if a field has been set.

### GetDescription

`func (o *SplitsCalendarResponseItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SplitsCalendarResponseItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SplitsCalendarResponseItem) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetRatio

`func (o *SplitsCalendarResponseItem) GetRatio() float64`

GetRatio returns the Ratio field if non-nil, zero value otherwise.

### GetRatioOk

`func (o *SplitsCalendarResponseItem) GetRatioOk() (*float64, bool)`

GetRatioOk returns a tuple with the Ratio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatio

`func (o *SplitsCalendarResponseItem) SetRatio(v float64)`

SetRatio sets Ratio field to given value.


### GetFromFactor

`func (o *SplitsCalendarResponseItem) GetFromFactor() float64`

GetFromFactor returns the FromFactor field if non-nil, zero value otherwise.

### GetFromFactorOk

`func (o *SplitsCalendarResponseItem) GetFromFactorOk() (*float64, bool)`

GetFromFactorOk returns a tuple with the FromFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromFactor

`func (o *SplitsCalendarResponseItem) SetFromFactor(v float64)`

SetFromFactor sets FromFactor field to given value.


### GetToFactor

`func (o *SplitsCalendarResponseItem) GetToFactor() float64`

GetToFactor returns the ToFactor field if non-nil, zero value otherwise.

### GetToFactorOk

`func (o *SplitsCalendarResponseItem) GetToFactorOk() (*float64, bool)`

GetToFactorOk returns a tuple with the ToFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToFactor

`func (o *SplitsCalendarResponseItem) SetToFactor(v float64)`

SetToFactor sets ToFactor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


