// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the GetStatistics200ResponseStatisticsFinancialsBalanceSheet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetStatistics200ResponseStatisticsFinancialsBalanceSheet{}

// GetStatistics200ResponseStatisticsFinancialsBalanceSheet Balance sheet information
type GetStatistics200ResponseStatisticsFinancialsBalanceSheet struct {
	// Refers to total cash measure for the most recent quarter
	TotalCashMrq *float64 `json:"total_cash_mrq,omitempty"`
	// Refers to total cash per share measure for the most recent quarter
	TotalCashPerShareMrq *float64 `json:"total_cash_per_share_mrq,omitempty"`
	// Refers to total debt measure for the most recent quarter
	TotalDebtMrq *float64 `json:"total_debt_mrq,omitempty"`
	// Refers to total debt to equity measure for the most recent quarter
	TotalDebtToEquityMrq *float64 `json:"total_debt_to_equity_mrq,omitempty"`
	// Refers to current ratio (total assets / total liabilities) ratio for the most recent quarter
	CurrentRatioMrq *float64 `json:"current_ratio_mrq,omitempty"`
	// Refers to book value per share (BVPS) ratio for the most recent quarter
	BookValuePerShareMrq *float64 `json:"book_value_per_share_mrq,omitempty"`
}

// NewGetStatistics200ResponseStatisticsFinancialsBalanceSheet instantiates a new GetStatistics200ResponseStatisticsFinancialsBalanceSheet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStatistics200ResponseStatisticsFinancialsBalanceSheet() *GetStatistics200ResponseStatisticsFinancialsBalanceSheet {
	this := GetStatistics200ResponseStatisticsFinancialsBalanceSheet{}
	return &this
}

// NewGetStatistics200ResponseStatisticsFinancialsBalanceSheetWithDefaults instantiates a new GetStatistics200ResponseStatisticsFinancialsBalanceSheet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStatistics200ResponseStatisticsFinancialsBalanceSheetWithDefaults() *GetStatistics200ResponseStatisticsFinancialsBalanceSheet {
	this := GetStatistics200ResponseStatisticsFinancialsBalanceSheet{}
	return &this
}

// GetTotalCashMrq returns the TotalCashMrq field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsFinancialsBalanceSheet) GetTotalCashMrq() float64 {
	if o == nil || IsNil(o.TotalCashMrq) {
		var ret float64
		return ret
	}
	return *o.TotalCashMrq
}

// GetTotalCashMrqOk returns a tuple with the TotalCashMrq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsFinancialsBalanceSheet) GetTotalCashMrqOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalCashMrq) {
		return nil, false
	}
	return o.TotalCashMrq, true
}

// HasTotalCashMrq returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsFinancialsBalanceSheet) HasTotalCashMrq() bool {
	if o != nil && !IsNil(o.TotalCashMrq) {
		return true
	}

	return false
}

// SetTotalCashMrq gets a reference to the given float64 and assigns it to the TotalCashMrq field.
func (o *GetStatistics200ResponseStatisticsFinancialsBalanceSheet) SetTotalCashMrq(v float64) {
	o.TotalCashMrq = &v
}

// GetTotalCashPerShareMrq returns the TotalCashPerShareMrq field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsFinancialsBalanceSheet) GetTotalCashPerShareMrq() float64 {
	if o == nil || IsNil(o.TotalCashPerShareMrq) {
		var ret float64
		return ret
	}
	return *o.TotalCashPerShareMrq
}

// GetTotalCashPerShareMrqOk returns a tuple with the TotalCashPerShareMrq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsFinancialsBalanceSheet) GetTotalCashPerShareMrqOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalCashPerShareMrq) {
		return nil, false
	}
	return o.TotalCashPerShareMrq, true
}

// HasTotalCashPerShareMrq returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsFinancialsBalanceSheet) HasTotalCashPerShareMrq() bool {
	if o != nil && !IsNil(o.TotalCashPerShareMrq) {
		return true
	}

	return false
}

// SetTotalCashPerShareMrq gets a reference to the given float64 and assigns it to the TotalCashPerShareMrq field.
func (o *GetStatistics200ResponseStatisticsFinancialsBalanceSheet) SetTotalCashPerShareMrq(v float64) {
	o.TotalCashPerShareMrq = &v
}

// GetTotalDebtMrq returns the TotalDebtMrq field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsFinancialsBalanceSheet) GetTotalDebtMrq() float64 {
	if o == nil || IsNil(o.TotalDebtMrq) {
		var ret float64
		return ret
	}
	return *o.TotalDebtMrq
}

// GetTotalDebtMrqOk returns a tuple with the TotalDebtMrq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsFinancialsBalanceSheet) GetTotalDebtMrqOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalDebtMrq) {
		return nil, false
	}
	return o.TotalDebtMrq, true
}

// HasTotalDebtMrq returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsFinancialsBalanceSheet) HasTotalDebtMrq() bool {
	if o != nil && !IsNil(o.TotalDebtMrq) {
		return true
	}

	return false
}

// SetTotalDebtMrq gets a reference to the given float64 and assigns it to the TotalDebtMrq field.
func (o *GetStatistics200ResponseStatisticsFinancialsBalanceSheet) SetTotalDebtMrq(v float64) {
	o.TotalDebtMrq = &v
}

// GetTotalDebtToEquityMrq returns the TotalDebtToEquityMrq field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsFinancialsBalanceSheet) GetTotalDebtToEquityMrq() float64 {
	if o == nil || IsNil(o.TotalDebtToEquityMrq) {
		var ret float64
		return ret
	}
	return *o.TotalDebtToEquityMrq
}

// GetTotalDebtToEquityMrqOk returns a tuple with the TotalDebtToEquityMrq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsFinancialsBalanceSheet) GetTotalDebtToEquityMrqOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalDebtToEquityMrq) {
		return nil, false
	}
	return o.TotalDebtToEquityMrq, true
}

// HasTotalDebtToEquityMrq returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsFinancialsBalanceSheet) HasTotalDebtToEquityMrq() bool {
	if o != nil && !IsNil(o.TotalDebtToEquityMrq) {
		return true
	}

	return false
}

// SetTotalDebtToEquityMrq gets a reference to the given float64 and assigns it to the TotalDebtToEquityMrq field.
func (o *GetStatistics200ResponseStatisticsFinancialsBalanceSheet) SetTotalDebtToEquityMrq(v float64) {
	o.TotalDebtToEquityMrq = &v
}

// GetCurrentRatioMrq returns the CurrentRatioMrq field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsFinancialsBalanceSheet) GetCurrentRatioMrq() float64 {
	if o == nil || IsNil(o.CurrentRatioMrq) {
		var ret float64
		return ret
	}
	return *o.CurrentRatioMrq
}

// GetCurrentRatioMrqOk returns a tuple with the CurrentRatioMrq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsFinancialsBalanceSheet) GetCurrentRatioMrqOk() (*float64, bool) {
	if o == nil || IsNil(o.CurrentRatioMrq) {
		return nil, false
	}
	return o.CurrentRatioMrq, true
}

// HasCurrentRatioMrq returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsFinancialsBalanceSheet) HasCurrentRatioMrq() bool {
	if o != nil && !IsNil(o.CurrentRatioMrq) {
		return true
	}

	return false
}

// SetCurrentRatioMrq gets a reference to the given float64 and assigns it to the CurrentRatioMrq field.
func (o *GetStatistics200ResponseStatisticsFinancialsBalanceSheet) SetCurrentRatioMrq(v float64) {
	o.CurrentRatioMrq = &v
}

// GetBookValuePerShareMrq returns the BookValuePerShareMrq field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsFinancialsBalanceSheet) GetBookValuePerShareMrq() float64 {
	if o == nil || IsNil(o.BookValuePerShareMrq) {
		var ret float64
		return ret
	}
	return *o.BookValuePerShareMrq
}

// GetBookValuePerShareMrqOk returns a tuple with the BookValuePerShareMrq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsFinancialsBalanceSheet) GetBookValuePerShareMrqOk() (*float64, bool) {
	if o == nil || IsNil(o.BookValuePerShareMrq) {
		return nil, false
	}
	return o.BookValuePerShareMrq, true
}

// HasBookValuePerShareMrq returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsFinancialsBalanceSheet) HasBookValuePerShareMrq() bool {
	if o != nil && !IsNil(o.BookValuePerShareMrq) {
		return true
	}

	return false
}

// SetBookValuePerShareMrq gets a reference to the given float64 and assigns it to the BookValuePerShareMrq field.
func (o *GetStatistics200ResponseStatisticsFinancialsBalanceSheet) SetBookValuePerShareMrq(v float64) {
	o.BookValuePerShareMrq = &v
}

func (o GetStatistics200ResponseStatisticsFinancialsBalanceSheet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetStatistics200ResponseStatisticsFinancialsBalanceSheet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalCashMrq) {
		toSerialize["total_cash_mrq"] = o.TotalCashMrq
	}
	if !IsNil(o.TotalCashPerShareMrq) {
		toSerialize["total_cash_per_share_mrq"] = o.TotalCashPerShareMrq
	}
	if !IsNil(o.TotalDebtMrq) {
		toSerialize["total_debt_mrq"] = o.TotalDebtMrq
	}
	if !IsNil(o.TotalDebtToEquityMrq) {
		toSerialize["total_debt_to_equity_mrq"] = o.TotalDebtToEquityMrq
	}
	if !IsNil(o.CurrentRatioMrq) {
		toSerialize["current_ratio_mrq"] = o.CurrentRatioMrq
	}
	if !IsNil(o.BookValuePerShareMrq) {
		toSerialize["book_value_per_share_mrq"] = o.BookValuePerShareMrq
	}
	return toSerialize, nil
}

type NullableGetStatistics200ResponseStatisticsFinancialsBalanceSheet struct {
	value *GetStatistics200ResponseStatisticsFinancialsBalanceSheet
	isSet bool
}

func (v NullableGetStatistics200ResponseStatisticsFinancialsBalanceSheet) Get() *GetStatistics200ResponseStatisticsFinancialsBalanceSheet {
	return v.value
}

func (v *NullableGetStatistics200ResponseStatisticsFinancialsBalanceSheet) Set(val *GetStatistics200ResponseStatisticsFinancialsBalanceSheet) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStatistics200ResponseStatisticsFinancialsBalanceSheet) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStatistics200ResponseStatisticsFinancialsBalanceSheet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStatistics200ResponseStatisticsFinancialsBalanceSheet(val *GetStatistics200ResponseStatisticsFinancialsBalanceSheet) *NullableGetStatistics200ResponseStatisticsFinancialsBalanceSheet {
	return &NullableGetStatistics200ResponseStatisticsFinancialsBalanceSheet{value: val, isSet: true}
}

func (v NullableGetStatistics200ResponseStatisticsFinancialsBalanceSheet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStatistics200ResponseStatisticsFinancialsBalanceSheet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
