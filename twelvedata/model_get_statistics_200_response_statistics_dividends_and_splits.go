// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the GetStatistics200ResponseStatisticsDividendsAndSplits type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetStatistics200ResponseStatisticsDividendsAndSplits{}

// GetStatistics200ResponseStatisticsDividendsAndSplits Dividends and splits information of the company
type GetStatistics200ResponseStatisticsDividendsAndSplits struct {
	// Refers to the forward dividend yield estimation in the currency of instrument
	ForwardAnnualDividendRate *float64 `json:"forward_annual_dividend_rate,omitempty"`
	// Refers to the forward dividend yield percentage relative to stock price
	ForwardAnnualDividendYield *float64 `json:"forward_annual_dividend_yield,omitempty"`
	// Refers to the trailing dividend yield rate in the currency of instrument over the last 12 months
	TrailingAnnualDividendRate *float64 `json:"trailing_annual_dividend_rate,omitempty"`
	// Refers to the trailing dividend yield percentage relative to stock price
	TrailingAnnualDividendYield *float64 `json:"trailing_annual_dividend_yield,omitempty"`
	// Refers to the average 5 years dividend yield
	Var5YearAverageDividendYield *float64 `json:"5_year_average_dividend_yield,omitempty"`
	// Refers to payout ratio, showing the proportion of earnings a company pays its shareholders in the form of dividends
	PayoutRatio *float64 `json:"payout_ratio,omitempty"`
	// Refers to how often a stock or fund pays dividends
	DividendFrequency *string `json:"dividend_frequency,omitempty"`
	// Refers to the last dividend payout date
	DividendDate *string `json:"dividend_date,omitempty"`
	// Refers to the last ex-dividend payout date
	ExDividendDate *string `json:"ex_dividend_date,omitempty"`
	// Specification of the last split event
	LastSplitFactor *string `json:"last_split_factor,omitempty"`
	// Refers for the last split date
	LastSplitDate *string `json:"last_split_date,omitempty"`
}

// NewGetStatistics200ResponseStatisticsDividendsAndSplits instantiates a new GetStatistics200ResponseStatisticsDividendsAndSplits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStatistics200ResponseStatisticsDividendsAndSplits() *GetStatistics200ResponseStatisticsDividendsAndSplits {
	this := GetStatistics200ResponseStatisticsDividendsAndSplits{}
	return &this
}

// NewGetStatistics200ResponseStatisticsDividendsAndSplitsWithDefaults instantiates a new GetStatistics200ResponseStatisticsDividendsAndSplits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStatistics200ResponseStatisticsDividendsAndSplitsWithDefaults() *GetStatistics200ResponseStatisticsDividendsAndSplits {
	this := GetStatistics200ResponseStatisticsDividendsAndSplits{}
	return &this
}

// GetForwardAnnualDividendRate returns the ForwardAnnualDividendRate field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetForwardAnnualDividendRate() float64 {
	if o == nil || IsNil(o.ForwardAnnualDividendRate) {
		var ret float64
		return ret
	}
	return *o.ForwardAnnualDividendRate
}

// GetForwardAnnualDividendRateOk returns a tuple with the ForwardAnnualDividendRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetForwardAnnualDividendRateOk() (*float64, bool) {
	if o == nil || IsNil(o.ForwardAnnualDividendRate) {
		return nil, false
	}
	return o.ForwardAnnualDividendRate, true
}

// HasForwardAnnualDividendRate returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) HasForwardAnnualDividendRate() bool {
	if o != nil && !IsNil(o.ForwardAnnualDividendRate) {
		return true
	}

	return false
}

// SetForwardAnnualDividendRate gets a reference to the given float64 and assigns it to the ForwardAnnualDividendRate field.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) SetForwardAnnualDividendRate(v float64) {
	o.ForwardAnnualDividendRate = &v
}

// GetForwardAnnualDividendYield returns the ForwardAnnualDividendYield field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetForwardAnnualDividendYield() float64 {
	if o == nil || IsNil(o.ForwardAnnualDividendYield) {
		var ret float64
		return ret
	}
	return *o.ForwardAnnualDividendYield
}

// GetForwardAnnualDividendYieldOk returns a tuple with the ForwardAnnualDividendYield field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetForwardAnnualDividendYieldOk() (*float64, bool) {
	if o == nil || IsNil(o.ForwardAnnualDividendYield) {
		return nil, false
	}
	return o.ForwardAnnualDividendYield, true
}

// HasForwardAnnualDividendYield returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) HasForwardAnnualDividendYield() bool {
	if o != nil && !IsNil(o.ForwardAnnualDividendYield) {
		return true
	}

	return false
}

// SetForwardAnnualDividendYield gets a reference to the given float64 and assigns it to the ForwardAnnualDividendYield field.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) SetForwardAnnualDividendYield(v float64) {
	o.ForwardAnnualDividendYield = &v
}

// GetTrailingAnnualDividendRate returns the TrailingAnnualDividendRate field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetTrailingAnnualDividendRate() float64 {
	if o == nil || IsNil(o.TrailingAnnualDividendRate) {
		var ret float64
		return ret
	}
	return *o.TrailingAnnualDividendRate
}

// GetTrailingAnnualDividendRateOk returns a tuple with the TrailingAnnualDividendRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetTrailingAnnualDividendRateOk() (*float64, bool) {
	if o == nil || IsNil(o.TrailingAnnualDividendRate) {
		return nil, false
	}
	return o.TrailingAnnualDividendRate, true
}

// HasTrailingAnnualDividendRate returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) HasTrailingAnnualDividendRate() bool {
	if o != nil && !IsNil(o.TrailingAnnualDividendRate) {
		return true
	}

	return false
}

// SetTrailingAnnualDividendRate gets a reference to the given float64 and assigns it to the TrailingAnnualDividendRate field.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) SetTrailingAnnualDividendRate(v float64) {
	o.TrailingAnnualDividendRate = &v
}

// GetTrailingAnnualDividendYield returns the TrailingAnnualDividendYield field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetTrailingAnnualDividendYield() float64 {
	if o == nil || IsNil(o.TrailingAnnualDividendYield) {
		var ret float64
		return ret
	}
	return *o.TrailingAnnualDividendYield
}

// GetTrailingAnnualDividendYieldOk returns a tuple with the TrailingAnnualDividendYield field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetTrailingAnnualDividendYieldOk() (*float64, bool) {
	if o == nil || IsNil(o.TrailingAnnualDividendYield) {
		return nil, false
	}
	return o.TrailingAnnualDividendYield, true
}

// HasTrailingAnnualDividendYield returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) HasTrailingAnnualDividendYield() bool {
	if o != nil && !IsNil(o.TrailingAnnualDividendYield) {
		return true
	}

	return false
}

// SetTrailingAnnualDividendYield gets a reference to the given float64 and assigns it to the TrailingAnnualDividendYield field.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) SetTrailingAnnualDividendYield(v float64) {
	o.TrailingAnnualDividendYield = &v
}

// GetVar5YearAverageDividendYield returns the Var5YearAverageDividendYield field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetVar5YearAverageDividendYield() float64 {
	if o == nil || IsNil(o.Var5YearAverageDividendYield) {
		var ret float64
		return ret
	}
	return *o.Var5YearAverageDividendYield
}

// GetVar5YearAverageDividendYieldOk returns a tuple with the Var5YearAverageDividendYield field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetVar5YearAverageDividendYieldOk() (*float64, bool) {
	if o == nil || IsNil(o.Var5YearAverageDividendYield) {
		return nil, false
	}
	return o.Var5YearAverageDividendYield, true
}

// HasVar5YearAverageDividendYield returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) HasVar5YearAverageDividendYield() bool {
	if o != nil && !IsNil(o.Var5YearAverageDividendYield) {
		return true
	}

	return false
}

// SetVar5YearAverageDividendYield gets a reference to the given float64 and assigns it to the Var5YearAverageDividendYield field.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) SetVar5YearAverageDividendYield(v float64) {
	o.Var5YearAverageDividendYield = &v
}

// GetPayoutRatio returns the PayoutRatio field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetPayoutRatio() float64 {
	if o == nil || IsNil(o.PayoutRatio) {
		var ret float64
		return ret
	}
	return *o.PayoutRatio
}

// GetPayoutRatioOk returns a tuple with the PayoutRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetPayoutRatioOk() (*float64, bool) {
	if o == nil || IsNil(o.PayoutRatio) {
		return nil, false
	}
	return o.PayoutRatio, true
}

// HasPayoutRatio returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) HasPayoutRatio() bool {
	if o != nil && !IsNil(o.PayoutRatio) {
		return true
	}

	return false
}

// SetPayoutRatio gets a reference to the given float64 and assigns it to the PayoutRatio field.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) SetPayoutRatio(v float64) {
	o.PayoutRatio = &v
}

// GetDividendFrequency returns the DividendFrequency field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetDividendFrequency() string {
	if o == nil || IsNil(o.DividendFrequency) {
		var ret string
		return ret
	}
	return *o.DividendFrequency
}

// GetDividendFrequencyOk returns a tuple with the DividendFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetDividendFrequencyOk() (*string, bool) {
	if o == nil || IsNil(o.DividendFrequency) {
		return nil, false
	}
	return o.DividendFrequency, true
}

// HasDividendFrequency returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) HasDividendFrequency() bool {
	if o != nil && !IsNil(o.DividendFrequency) {
		return true
	}

	return false
}

// SetDividendFrequency gets a reference to the given string and assigns it to the DividendFrequency field.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) SetDividendFrequency(v string) {
	o.DividendFrequency = &v
}

// GetDividendDate returns the DividendDate field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetDividendDate() string {
	if o == nil || IsNil(o.DividendDate) {
		var ret string
		return ret
	}
	return *o.DividendDate
}

// GetDividendDateOk returns a tuple with the DividendDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetDividendDateOk() (*string, bool) {
	if o == nil || IsNil(o.DividendDate) {
		return nil, false
	}
	return o.DividendDate, true
}

// HasDividendDate returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) HasDividendDate() bool {
	if o != nil && !IsNil(o.DividendDate) {
		return true
	}

	return false
}

// SetDividendDate gets a reference to the given string and assigns it to the DividendDate field.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) SetDividendDate(v string) {
	o.DividendDate = &v
}

// GetExDividendDate returns the ExDividendDate field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetExDividendDate() string {
	if o == nil || IsNil(o.ExDividendDate) {
		var ret string
		return ret
	}
	return *o.ExDividendDate
}

// GetExDividendDateOk returns a tuple with the ExDividendDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetExDividendDateOk() (*string, bool) {
	if o == nil || IsNil(o.ExDividendDate) {
		return nil, false
	}
	return o.ExDividendDate, true
}

// HasExDividendDate returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) HasExDividendDate() bool {
	if o != nil && !IsNil(o.ExDividendDate) {
		return true
	}

	return false
}

// SetExDividendDate gets a reference to the given string and assigns it to the ExDividendDate field.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) SetExDividendDate(v string) {
	o.ExDividendDate = &v
}

// GetLastSplitFactor returns the LastSplitFactor field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetLastSplitFactor() string {
	if o == nil || IsNil(o.LastSplitFactor) {
		var ret string
		return ret
	}
	return *o.LastSplitFactor
}

// GetLastSplitFactorOk returns a tuple with the LastSplitFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetLastSplitFactorOk() (*string, bool) {
	if o == nil || IsNil(o.LastSplitFactor) {
		return nil, false
	}
	return o.LastSplitFactor, true
}

// HasLastSplitFactor returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) HasLastSplitFactor() bool {
	if o != nil && !IsNil(o.LastSplitFactor) {
		return true
	}

	return false
}

// SetLastSplitFactor gets a reference to the given string and assigns it to the LastSplitFactor field.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) SetLastSplitFactor(v string) {
	o.LastSplitFactor = &v
}

// GetLastSplitDate returns the LastSplitDate field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetLastSplitDate() string {
	if o == nil || IsNil(o.LastSplitDate) {
		var ret string
		return ret
	}
	return *o.LastSplitDate
}

// GetLastSplitDateOk returns a tuple with the LastSplitDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) GetLastSplitDateOk() (*string, bool) {
	if o == nil || IsNil(o.LastSplitDate) {
		return nil, false
	}
	return o.LastSplitDate, true
}

// HasLastSplitDate returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) HasLastSplitDate() bool {
	if o != nil && !IsNil(o.LastSplitDate) {
		return true
	}

	return false
}

// SetLastSplitDate gets a reference to the given string and assigns it to the LastSplitDate field.
func (o *GetStatistics200ResponseStatisticsDividendsAndSplits) SetLastSplitDate(v string) {
	o.LastSplitDate = &v
}

func (o GetStatistics200ResponseStatisticsDividendsAndSplits) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetStatistics200ResponseStatisticsDividendsAndSplits) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ForwardAnnualDividendRate) {
		toSerialize["forward_annual_dividend_rate"] = o.ForwardAnnualDividendRate
	}
	if !IsNil(o.ForwardAnnualDividendYield) {
		toSerialize["forward_annual_dividend_yield"] = o.ForwardAnnualDividendYield
	}
	if !IsNil(o.TrailingAnnualDividendRate) {
		toSerialize["trailing_annual_dividend_rate"] = o.TrailingAnnualDividendRate
	}
	if !IsNil(o.TrailingAnnualDividendYield) {
		toSerialize["trailing_annual_dividend_yield"] = o.TrailingAnnualDividendYield
	}
	if !IsNil(o.Var5YearAverageDividendYield) {
		toSerialize["5_year_average_dividend_yield"] = o.Var5YearAverageDividendYield
	}
	if !IsNil(o.PayoutRatio) {
		toSerialize["payout_ratio"] = o.PayoutRatio
	}
	if !IsNil(o.DividendFrequency) {
		toSerialize["dividend_frequency"] = o.DividendFrequency
	}
	if !IsNil(o.DividendDate) {
		toSerialize["dividend_date"] = o.DividendDate
	}
	if !IsNil(o.ExDividendDate) {
		toSerialize["ex_dividend_date"] = o.ExDividendDate
	}
	if !IsNil(o.LastSplitFactor) {
		toSerialize["last_split_factor"] = o.LastSplitFactor
	}
	if !IsNil(o.LastSplitDate) {
		toSerialize["last_split_date"] = o.LastSplitDate
	}
	return toSerialize, nil
}

type NullableGetStatistics200ResponseStatisticsDividendsAndSplits struct {
	value *GetStatistics200ResponseStatisticsDividendsAndSplits
	isSet bool
}

func (v NullableGetStatistics200ResponseStatisticsDividendsAndSplits) Get() *GetStatistics200ResponseStatisticsDividendsAndSplits {
	return v.value
}

func (v *NullableGetStatistics200ResponseStatisticsDividendsAndSplits) Set(val *GetStatistics200ResponseStatisticsDividendsAndSplits) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStatistics200ResponseStatisticsDividendsAndSplits) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStatistics200ResponseStatisticsDividendsAndSplits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStatistics200ResponseStatisticsDividendsAndSplits(val *GetStatistics200ResponseStatisticsDividendsAndSplits) *NullableGetStatistics200ResponseStatisticsDividendsAndSplits {
	return &NullableGetStatistics200ResponseStatisticsDividendsAndSplits{value: val, isSet: true}
}

func (v NullableGetStatistics200ResponseStatisticsDividendsAndSplits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStatistics200ResponseStatisticsDividendsAndSplits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
