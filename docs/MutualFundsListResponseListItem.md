# MutualFundsListResponseListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Symbol** | **string** | Fund symbol ticker | 
**Name** | **string** | Fund name | 
**Country** | **string** | Country of fund incorporation | 
**FundFamily** | **string** | Investment company that manages the fund | 
**FundType** | **string** | Type of fund | 
**PerformanceRating** | Pointer to **int64** | Performance rating from &#x60;0&#x60; to &#x60;5&#x60; | [optional] 
**RiskRating** | Pointer to **int64** | Risk rating from &#x60;0&#x60; to &#x60;5&#x60; | [optional] 
**Currency** | **string** | Currency code in which the fund is denominated | 
**Exchange** | **string** | Exchange name where the fund is listed | 
**MicCode** | **string** | Market identifier code (MIC) under ISO 10383 standard | 

## Methods

### NewMutualFundsListResponseListItem

`func NewMutualFundsListResponseListItem(symbol string, name string, country string, fundFamily string, fundType string, currency string, exchange string, micCode string, ) *MutualFundsListResponseListItem`

NewMutualFundsListResponseListItem instantiates a new MutualFundsListResponseListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMutualFundsListResponseListItemWithDefaults

`func NewMutualFundsListResponseListItemWithDefaults() *MutualFundsListResponseListItem`

NewMutualFundsListResponseListItemWithDefaults instantiates a new MutualFundsListResponseListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSymbol

`func (o *MutualFundsListResponseListItem) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *MutualFundsListResponseListItem) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *MutualFundsListResponseListItem) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetName

`func (o *MutualFundsListResponseListItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MutualFundsListResponseListItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MutualFundsListResponseListItem) SetName(v string)`

SetName sets Name field to given value.


### GetCountry

`func (o *MutualFundsListResponseListItem) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *MutualFundsListResponseListItem) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *MutualFundsListResponseListItem) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetFundFamily

`func (o *MutualFundsListResponseListItem) GetFundFamily() string`

GetFundFamily returns the FundFamily field if non-nil, zero value otherwise.

### GetFundFamilyOk

`func (o *MutualFundsListResponseListItem) GetFundFamilyOk() (*string, bool)`

GetFundFamilyOk returns a tuple with the FundFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundFamily

`func (o *MutualFundsListResponseListItem) SetFundFamily(v string)`

SetFundFamily sets FundFamily field to given value.


### GetFundType

`func (o *MutualFundsListResponseListItem) GetFundType() string`

GetFundType returns the FundType field if non-nil, zero value otherwise.

### GetFundTypeOk

`func (o *MutualFundsListResponseListItem) GetFundTypeOk() (*string, bool)`

GetFundTypeOk returns a tuple with the FundType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundType

`func (o *MutualFundsListResponseListItem) SetFundType(v string)`

SetFundType sets FundType field to given value.


### GetPerformanceRating

`func (o *MutualFundsListResponseListItem) GetPerformanceRating() int64`

GetPerformanceRating returns the PerformanceRating field if non-nil, zero value otherwise.

### GetPerformanceRatingOk

`func (o *MutualFundsListResponseListItem) GetPerformanceRatingOk() (*int64, bool)`

GetPerformanceRatingOk returns a tuple with the PerformanceRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformanceRating

`func (o *MutualFundsListResponseListItem) SetPerformanceRating(v int64)`

SetPerformanceRating sets PerformanceRating field to given value.

### HasPerformanceRating

`func (o *MutualFundsListResponseListItem) HasPerformanceRating() bool`

HasPerformanceRating returns a boolean if a field has been set.

### GetRiskRating

`func (o *MutualFundsListResponseListItem) GetRiskRating() int64`

GetRiskRating returns the RiskRating field if non-nil, zero value otherwise.

### GetRiskRatingOk

`func (o *MutualFundsListResponseListItem) GetRiskRatingOk() (*int64, bool)`

GetRiskRatingOk returns a tuple with the RiskRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskRating

`func (o *MutualFundsListResponseListItem) SetRiskRating(v int64)`

SetRiskRating sets RiskRating field to given value.

### HasRiskRating

`func (o *MutualFundsListResponseListItem) HasRiskRating() bool`

HasRiskRating returns a boolean if a field has been set.

### GetCurrency

`func (o *MutualFundsListResponseListItem) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *MutualFundsListResponseListItem) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *MutualFundsListResponseListItem) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetExchange

`func (o *MutualFundsListResponseListItem) GetExchange() string`

GetExchange returns the Exchange field if non-nil, zero value otherwise.

### GetExchangeOk

`func (o *MutualFundsListResponseListItem) GetExchangeOk() (*string, bool)`

GetExchangeOk returns a tuple with the Exchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchange

`func (o *MutualFundsListResponseListItem) SetExchange(v string)`

SetExchange sets Exchange field to given value.


### GetMicCode

`func (o *MutualFundsListResponseListItem) GetMicCode() string`

GetMicCode returns the MicCode field if non-nil, zero value otherwise.

### GetMicCodeOk

`func (o *MutualFundsListResponseListItem) GetMicCodeOk() (*string, bool)`

GetMicCodeOk returns a tuple with the MicCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicCode

`func (o *MutualFundsListResponseListItem) SetMicCode(v string)`

SetMicCode sets MicCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


