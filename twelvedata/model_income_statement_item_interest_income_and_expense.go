/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the IncomeStatementItemInterestIncomeAndExpense type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IncomeStatementItemInterestIncomeAndExpense{}

// IncomeStatementItemInterestIncomeAndExpense Interest income and expense information
type IncomeStatementItemInterestIncomeAndExpense struct {
	// Interest income
	InterestIncome *float64 `json:"interest_income,omitempty"`
	// Interest expense
	InterestExpense *float64 `json:"interest_expense,omitempty"`
	// Net interest income
	NetInterestIncome *float64 `json:"net_interest_income,omitempty"`
	// Net non operating interest income expense
	NetNonOperatingInterestIncomeExpense *float64 `json:"net_non_operating_interest_income_expense,omitempty"`
	// Interest expense non operating
	InterestExpenseNonOperating *float64 `json:"interest_expense_non_operating,omitempty"`
	// Interest income non operating
	InterestIncomeNonOperating *float64 `json:"interest_income_non_operating,omitempty"`
}

// NewIncomeStatementItemInterestIncomeAndExpense instantiates a new IncomeStatementItemInterestIncomeAndExpense object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeStatementItemInterestIncomeAndExpense() *IncomeStatementItemInterestIncomeAndExpense {
	this := IncomeStatementItemInterestIncomeAndExpense{}
	return &this
}

// NewIncomeStatementItemInterestIncomeAndExpenseWithDefaults instantiates a new IncomeStatementItemInterestIncomeAndExpense object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeStatementItemInterestIncomeAndExpenseWithDefaults() *IncomeStatementItemInterestIncomeAndExpense {
	this := IncomeStatementItemInterestIncomeAndExpense{}
	return &this
}

// GetInterestIncome returns the InterestIncome field value if set, zero value otherwise.
func (o *IncomeStatementItemInterestIncomeAndExpense) GetInterestIncome() float64 {
	if o == nil || IsNil(o.InterestIncome) {
		var ret float64
		return ret
	}
	return *o.InterestIncome
}

// GetInterestIncomeOk returns a tuple with the InterestIncome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemInterestIncomeAndExpense) GetInterestIncomeOk() (*float64, bool) {
	if o == nil || IsNil(o.InterestIncome) {
		return nil, false
	}
	return o.InterestIncome, true
}

// HasInterestIncome returns a boolean if a field has been set.
func (o *IncomeStatementItemInterestIncomeAndExpense) HasInterestIncome() bool {
	if o != nil && !IsNil(o.InterestIncome) {
		return true
	}

	return false
}

// SetInterestIncome gets a reference to the given float64 and assigns it to the InterestIncome field.
func (o *IncomeStatementItemInterestIncomeAndExpense) SetInterestIncome(v float64) {
	o.InterestIncome = &v
}

// GetInterestExpense returns the InterestExpense field value if set, zero value otherwise.
func (o *IncomeStatementItemInterestIncomeAndExpense) GetInterestExpense() float64 {
	if o == nil || IsNil(o.InterestExpense) {
		var ret float64
		return ret
	}
	return *o.InterestExpense
}

// GetInterestExpenseOk returns a tuple with the InterestExpense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemInterestIncomeAndExpense) GetInterestExpenseOk() (*float64, bool) {
	if o == nil || IsNil(o.InterestExpense) {
		return nil, false
	}
	return o.InterestExpense, true
}

// HasInterestExpense returns a boolean if a field has been set.
func (o *IncomeStatementItemInterestIncomeAndExpense) HasInterestExpense() bool {
	if o != nil && !IsNil(o.InterestExpense) {
		return true
	}

	return false
}

// SetInterestExpense gets a reference to the given float64 and assigns it to the InterestExpense field.
func (o *IncomeStatementItemInterestIncomeAndExpense) SetInterestExpense(v float64) {
	o.InterestExpense = &v
}

// GetNetInterestIncome returns the NetInterestIncome field value if set, zero value otherwise.
func (o *IncomeStatementItemInterestIncomeAndExpense) GetNetInterestIncome() float64 {
	if o == nil || IsNil(o.NetInterestIncome) {
		var ret float64
		return ret
	}
	return *o.NetInterestIncome
}

// GetNetInterestIncomeOk returns a tuple with the NetInterestIncome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemInterestIncomeAndExpense) GetNetInterestIncomeOk() (*float64, bool) {
	if o == nil || IsNil(o.NetInterestIncome) {
		return nil, false
	}
	return o.NetInterestIncome, true
}

// HasNetInterestIncome returns a boolean if a field has been set.
func (o *IncomeStatementItemInterestIncomeAndExpense) HasNetInterestIncome() bool {
	if o != nil && !IsNil(o.NetInterestIncome) {
		return true
	}

	return false
}

// SetNetInterestIncome gets a reference to the given float64 and assigns it to the NetInterestIncome field.
func (o *IncomeStatementItemInterestIncomeAndExpense) SetNetInterestIncome(v float64) {
	o.NetInterestIncome = &v
}

// GetNetNonOperatingInterestIncomeExpense returns the NetNonOperatingInterestIncomeExpense field value if set, zero value otherwise.
func (o *IncomeStatementItemInterestIncomeAndExpense) GetNetNonOperatingInterestIncomeExpense() float64 {
	if o == nil || IsNil(o.NetNonOperatingInterestIncomeExpense) {
		var ret float64
		return ret
	}
	return *o.NetNonOperatingInterestIncomeExpense
}

// GetNetNonOperatingInterestIncomeExpenseOk returns a tuple with the NetNonOperatingInterestIncomeExpense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemInterestIncomeAndExpense) GetNetNonOperatingInterestIncomeExpenseOk() (*float64, bool) {
	if o == nil || IsNil(o.NetNonOperatingInterestIncomeExpense) {
		return nil, false
	}
	return o.NetNonOperatingInterestIncomeExpense, true
}

// HasNetNonOperatingInterestIncomeExpense returns a boolean if a field has been set.
func (o *IncomeStatementItemInterestIncomeAndExpense) HasNetNonOperatingInterestIncomeExpense() bool {
	if o != nil && !IsNil(o.NetNonOperatingInterestIncomeExpense) {
		return true
	}

	return false
}

// SetNetNonOperatingInterestIncomeExpense gets a reference to the given float64 and assigns it to the NetNonOperatingInterestIncomeExpense field.
func (o *IncomeStatementItemInterestIncomeAndExpense) SetNetNonOperatingInterestIncomeExpense(v float64) {
	o.NetNonOperatingInterestIncomeExpense = &v
}

// GetInterestExpenseNonOperating returns the InterestExpenseNonOperating field value if set, zero value otherwise.
func (o *IncomeStatementItemInterestIncomeAndExpense) GetInterestExpenseNonOperating() float64 {
	if o == nil || IsNil(o.InterestExpenseNonOperating) {
		var ret float64
		return ret
	}
	return *o.InterestExpenseNonOperating
}

// GetInterestExpenseNonOperatingOk returns a tuple with the InterestExpenseNonOperating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemInterestIncomeAndExpense) GetInterestExpenseNonOperatingOk() (*float64, bool) {
	if o == nil || IsNil(o.InterestExpenseNonOperating) {
		return nil, false
	}
	return o.InterestExpenseNonOperating, true
}

// HasInterestExpenseNonOperating returns a boolean if a field has been set.
func (o *IncomeStatementItemInterestIncomeAndExpense) HasInterestExpenseNonOperating() bool {
	if o != nil && !IsNil(o.InterestExpenseNonOperating) {
		return true
	}

	return false
}

// SetInterestExpenseNonOperating gets a reference to the given float64 and assigns it to the InterestExpenseNonOperating field.
func (o *IncomeStatementItemInterestIncomeAndExpense) SetInterestExpenseNonOperating(v float64) {
	o.InterestExpenseNonOperating = &v
}

// GetInterestIncomeNonOperating returns the InterestIncomeNonOperating field value if set, zero value otherwise.
func (o *IncomeStatementItemInterestIncomeAndExpense) GetInterestIncomeNonOperating() float64 {
	if o == nil || IsNil(o.InterestIncomeNonOperating) {
		var ret float64
		return ret
	}
	return *o.InterestIncomeNonOperating
}

// GetInterestIncomeNonOperatingOk returns a tuple with the InterestIncomeNonOperating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemInterestIncomeAndExpense) GetInterestIncomeNonOperatingOk() (*float64, bool) {
	if o == nil || IsNil(o.InterestIncomeNonOperating) {
		return nil, false
	}
	return o.InterestIncomeNonOperating, true
}

// HasInterestIncomeNonOperating returns a boolean if a field has been set.
func (o *IncomeStatementItemInterestIncomeAndExpense) HasInterestIncomeNonOperating() bool {
	if o != nil && !IsNil(o.InterestIncomeNonOperating) {
		return true
	}

	return false
}

// SetInterestIncomeNonOperating gets a reference to the given float64 and assigns it to the InterestIncomeNonOperating field.
func (o *IncomeStatementItemInterestIncomeAndExpense) SetInterestIncomeNonOperating(v float64) {
	o.InterestIncomeNonOperating = &v
}

func (o IncomeStatementItemInterestIncomeAndExpense) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IncomeStatementItemInterestIncomeAndExpense) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InterestIncome) {
		toSerialize["interest_income"] = o.InterestIncome
	}
	if !IsNil(o.InterestExpense) {
		toSerialize["interest_expense"] = o.InterestExpense
	}
	if !IsNil(o.NetInterestIncome) {
		toSerialize["net_interest_income"] = o.NetInterestIncome
	}
	if !IsNil(o.NetNonOperatingInterestIncomeExpense) {
		toSerialize["net_non_operating_interest_income_expense"] = o.NetNonOperatingInterestIncomeExpense
	}
	if !IsNil(o.InterestExpenseNonOperating) {
		toSerialize["interest_expense_non_operating"] = o.InterestExpenseNonOperating
	}
	if !IsNil(o.InterestIncomeNonOperating) {
		toSerialize["interest_income_non_operating"] = o.InterestIncomeNonOperating
	}
	return toSerialize, nil
}

type NullableIncomeStatementItemInterestIncomeAndExpense struct {
	value *IncomeStatementItemInterestIncomeAndExpense
	isSet bool
}

func (v NullableIncomeStatementItemInterestIncomeAndExpense) Get() *IncomeStatementItemInterestIncomeAndExpense {
	return v.value
}

func (v *NullableIncomeStatementItemInterestIncomeAndExpense) Set(val *IncomeStatementItemInterestIncomeAndExpense) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeStatementItemInterestIncomeAndExpense) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeStatementItemInterestIncomeAndExpense) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeStatementItemInterestIncomeAndExpense(val *IncomeStatementItemInterestIncomeAndExpense) *NullableIncomeStatementItemInterestIncomeAndExpense {
	return &NullableIncomeStatementItemInterestIncomeAndExpense{value: val, isSet: true}
}

func (v NullableIncomeStatementItemInterestIncomeAndExpense) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeStatementItemInterestIncomeAndExpense) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
