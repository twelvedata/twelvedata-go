# GetStatistics200ResponseStatisticsDividendsAndSplits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForwardAnnualDividendRate** | Pointer to **float64** | Refers to the forward dividend yield estimation in the currency of instrument | [optional] 
**ForwardAnnualDividendYield** | Pointer to **float64** | Refers to the forward dividend yield percentage relative to stock price | [optional] 
**TrailingAnnualDividendRate** | Pointer to **float64** | Refers to the trailing dividend yield rate in the currency of instrument over the last 12 months | [optional] 
**TrailingAnnualDividendYield** | Pointer to **float64** | Refers to the trailing dividend yield percentage relative to stock price | [optional] 
**Var5YearAverageDividendYield** | Pointer to **float64** | Refers to the average 5 years dividend yield | [optional] 
**PayoutRatio** | Pointer to **float64** | Refers to payout ratio, showing the proportion of earnings a company pays its shareholders in the form of dividends | [optional] 
**DividendFrequency** | Pointer to **string** | Refers to how often a stock or fund pays dividends | [optional] 
**DividendDate** | Pointer to **string** | Refers to the last dividend payout date | [optional] 
**ExDividendDate** | Pointer to **string** | Refers to the last ex-dividend payout date | [optional] 
**LastSplitFactor** | Pointer to **string** | Specification of the last split event | [optional] 
**LastSplitDate** | Pointer to **string** | Refers for the last split date | [optional] 

## Methods

### NewGetStatistics200ResponseStatisticsDividendsAndSplits

`func NewGetStatistics200ResponseStatisticsDividendsAndSplits() *GetStatistics200ResponseStatisticsDividendsAndSplits`

NewGetStatistics200ResponseStatisticsDividendsAndSplits instantiates a new GetStatistics200ResponseStatisticsDividendsAndSplits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStatistics200ResponseStatisticsDividendsAndSplitsWithDefaults

`func NewGetStatistics200ResponseStatisticsDividendsAndSplitsWithDefaults() *GetStatistics200ResponseStatisticsDividendsAndSplits`

NewGetStatistics200ResponseStatisticsDividendsAndSplitsWithDefaults instantiates a new GetStatistics200ResponseStatisticsDividendsAndSplits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForwardAnnualDividendRate

`func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetForwardAnnualDividendRate() float64`

GetForwardAnnualDividendRate returns the ForwardAnnualDividendRate field if non-nil, zero value otherwise.

### GetForwardAnnualDividendRateOk

`func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetForwardAnnualDividendRateOk() (*float64, bool)`

GetForwardAnnualDividendRateOk returns a tuple with the ForwardAnnualDividendRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardAnnualDividendRate

`func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) SetForwardAnnualDividendRate(v float64)`

SetForwardAnnualDividendRate sets ForwardAnnualDividendRate field to given value.

### HasForwardAnnualDividendRate

`func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) HasForwardAnnualDividendRate() bool`

HasForwardAnnualDividendRate returns a boolean if a field has been set.

### GetForwardAnnualDividendYield

`func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetForwardAnnualDividendYield() float64`

GetForwardAnnualDividendYield returns the ForwardAnnualDividendYield field if non-nil, zero value otherwise.

### GetForwardAnnualDividendYieldOk

`func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetForwardAnnualDividendYieldOk() (*float64, bool)`

GetForwardAnnualDividendYieldOk returns a tuple with the ForwardAnnualDividendYield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardAnnualDividendYield

`func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) SetForwardAnnualDividendYield(v float64)`

SetForwardAnnualDividendYield sets ForwardAnnualDividendYield field to given value.

### HasForwardAnnualDividendYield

`func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) HasForwardAnnualDividendYield() bool`

HasForwardAnnualDividendYield returns a boolean if a field has been set.

### GetTrailingAnnualDividendRate

`func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetTrailingAnnualDividendRate() float64`

GetTrailingAnnualDividendRate returns the TrailingAnnualDividendRate field if non-nil, zero value otherwise.

### GetTrailingAnnualDividendRateOk

`func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetTrailingAnnualDividendRateOk() (*float64, bool)`

GetTrailingAnnualDividendRateOk returns a tuple with the TrailingAnnualDividendRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailingAnnualDividendRate

`func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) SetTrailingAnnualDividendRate(v float64)`

SetTrailingAnnualDividendRate sets TrailingAnnualDividendRate field to given value.

### HasTrailingAnnualDividendRate

`func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) HasTrailingAnnualDividendRate() bool`

HasTrailingAnnualDividendRate returns a boolean if a field has been set.

### GetTrailingAnnualDividendYield

`func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetTrailingAnnualDividendYield() float64`

GetTrailingAnnualDividendYield returns the TrailingAnnualDividendYield field if non-nil, zero value otherwise.

### GetTrailingAnnualDividendYieldOk

`func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetTrailingAnnualDividendYieldOk() (*float64, bool)`

GetTrailingAnnualDividendYieldOk returns a tuple with the TrailingAnnualDividendYield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrailingAnnualDividendYield

`func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) SetTrailingAnnualDividendYield(v float64)`

SetTrailingAnnualDividendYield sets TrailingAnnualDividendYield field to given value.

### HasTrailingAnnualDividendYield

`func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) HasTrailingAnnualDividendYield() bool`

HasTrailingAnnualDividendYield returns a boolean if a field has been set.

### GetVar5YearAverageDividendYield

`func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetVar5YearAverageDividendYield() float64`

GetVar5YearAverageDividendYield returns the Var5YearAverageDividendYield field if non-nil, zero value otherwise.

### GetVar5YearAverageDividendYieldOk

`func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetVar5YearAverageDividendYieldOk() (*float64, bool)`

GetVar5YearAverageDividendYieldOk returns a tuple with the Var5YearAverageDividendYield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar5YearAverageDividendYield

`func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) SetVar5YearAverageDividendYield(v float64)`

SetVar5YearAverageDividendYield sets Var5YearAverageDividendYield field to given value.

### HasVar5YearAverageDividendYield

`func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) HasVar5YearAverageDividendYield() bool`

HasVar5YearAverageDividendYield returns a boolean if a field has been set.

### GetPayoutRatio

`func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetPayoutRatio() float64`

GetPayoutRatio returns the PayoutRatio field if non-nil, zero value otherwise.

### GetPayoutRatioOk

`func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetPayoutRatioOk() (*float64, bool)`

GetPayoutRatioOk returns a tuple with the PayoutRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutRatio

`func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) SetPayoutRatio(v float64)`

SetPayoutRatio sets PayoutRatio field to given value.

### HasPayoutRatio

`func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) HasPayoutRatio() bool`

HasPayoutRatio returns a boolean if a field has been set.

### GetDividendFrequency

`func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetDividendFrequency() string`

GetDividendFrequency returns the DividendFrequency field if non-nil, zero value otherwise.

### GetDividendFrequencyOk

`func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetDividendFrequencyOk() (*string, bool)`

GetDividendFrequencyOk returns a tuple with the DividendFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDividendFrequency

`func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) SetDividendFrequency(v string)`

SetDividendFrequency sets DividendFrequency field to given value.

### HasDividendFrequency

`func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) HasDividendFrequency() bool`

HasDividendFrequency returns a boolean if a field has been set.

### GetDividendDate

`func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetDividendDate() string`

GetDividendDate returns the DividendDate field if non-nil, zero value otherwise.

### GetDividendDateOk

`func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetDividendDateOk() (*string, bool)`

GetDividendDateOk returns a tuple with the DividendDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDividendDate

`func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) SetDividendDate(v string)`

SetDividendDate sets DividendDate field to given value.

### HasDividendDate

`func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) HasDividendDate() bool`

HasDividendDate returns a boolean if a field has been set.

### GetExDividendDate

`func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetExDividendDate() string`

GetExDividendDate returns the ExDividendDate field if non-nil, zero value otherwise.

### GetExDividendDateOk

`func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetExDividendDateOk() (*string, bool)`

GetExDividendDateOk returns a tuple with the ExDividendDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExDividendDate

`func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) SetExDividendDate(v string)`

SetExDividendDate sets ExDividendDate field to given value.

### HasExDividendDate

`func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) HasExDividendDate() bool`

HasExDividendDate returns a boolean if a field has been set.

### GetLastSplitFactor

`func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetLastSplitFactor() string`

GetLastSplitFactor returns the LastSplitFactor field if non-nil, zero value otherwise.

### GetLastSplitFactorOk

`func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetLastSplitFactorOk() (*string, bool)`

GetLastSplitFactorOk returns a tuple with the LastSplitFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSplitFactor

`func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) SetLastSplitFactor(v string)`

SetLastSplitFactor sets LastSplitFactor field to given value.

### HasLastSplitFactor

`func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) HasLastSplitFactor() bool`

HasLastSplitFactor returns a boolean if a field has been set.

### GetLastSplitDate

`func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetLastSplitDate() string`

GetLastSplitDate returns the LastSplitDate field if non-nil, zero value otherwise.

### GetLastSplitDateOk

`func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetLastSplitDateOk() (*string, bool)`

GetLastSplitDateOk returns a tuple with the LastSplitDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSplitDate

`func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) SetLastSplitDate(v string)`

SetLastSplitDate sets LastSplitDate field to given value.

### HasLastSplitDate

`func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) HasLastSplitDate() bool`

HasLastSplitDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


