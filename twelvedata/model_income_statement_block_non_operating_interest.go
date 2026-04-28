/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the IncomeStatementBlockNonOperatingInterest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IncomeStatementBlockNonOperatingInterest{}

// IncomeStatementBlockNonOperatingInterest Non-operating interest details
type IncomeStatementBlockNonOperatingInterest struct {
	// Refers to non-operating interest income
	Income *int64 `json:"income,omitempty"`
	// Refers to non-operating interest expense
	Expense *int64 `json:"expense,omitempty"`
}

// NewIncomeStatementBlockNonOperatingInterest instantiates a new IncomeStatementBlockNonOperatingInterest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeStatementBlockNonOperatingInterest() *IncomeStatementBlockNonOperatingInterest {
	this := IncomeStatementBlockNonOperatingInterest{}
	return &this
}

// NewIncomeStatementBlockNonOperatingInterestWithDefaults instantiates a new IncomeStatementBlockNonOperatingInterest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeStatementBlockNonOperatingInterestWithDefaults() *IncomeStatementBlockNonOperatingInterest {
	this := IncomeStatementBlockNonOperatingInterest{}
	return &this
}

// GetIncome returns the Income field value if set, zero value otherwise.
func (o *IncomeStatementBlockNonOperatingInterest) GetIncome() int64 {
	if o == nil || IsNil(o.Income) {
		var ret int64
		return ret
	}
	return *o.Income
}

// GetIncomeOk returns a tuple with the Income field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementBlockNonOperatingInterest) GetIncomeOk() (*int64, bool) {
	if o == nil || IsNil(o.Income) {
		return nil, false
	}
	return o.Income, true
}

// HasIncome returns a boolean if a field has been set.
func (o *IncomeStatementBlockNonOperatingInterest) HasIncome() bool {
	if o != nil && !IsNil(o.Income) {
		return true
	}

	return false
}

// SetIncome gets a reference to the given int64 and assigns it to the Income field.
func (o *IncomeStatementBlockNonOperatingInterest) SetIncome(v int64) {
	o.Income = &v
}

// GetExpense returns the Expense field value if set, zero value otherwise.
func (o *IncomeStatementBlockNonOperatingInterest) GetExpense() int64 {
	if o == nil || IsNil(o.Expense) {
		var ret int64
		return ret
	}
	return *o.Expense
}

// GetExpenseOk returns a tuple with the Expense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementBlockNonOperatingInterest) GetExpenseOk() (*int64, bool) {
	if o == nil || IsNil(o.Expense) {
		return nil, false
	}
	return o.Expense, true
}

// HasExpense returns a boolean if a field has been set.
func (o *IncomeStatementBlockNonOperatingInterest) HasExpense() bool {
	if o != nil && !IsNil(o.Expense) {
		return true
	}

	return false
}

// SetExpense gets a reference to the given int64 and assigns it to the Expense field.
func (o *IncomeStatementBlockNonOperatingInterest) SetExpense(v int64) {
	o.Expense = &v
}

func (o IncomeStatementBlockNonOperatingInterest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IncomeStatementBlockNonOperatingInterest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Income) {
		toSerialize["income"] = o.Income
	}
	if !IsNil(o.Expense) {
		toSerialize["expense"] = o.Expense
	}
	return toSerialize, nil
}

type NullableIncomeStatementBlockNonOperatingInterest struct {
	value *IncomeStatementBlockNonOperatingInterest
	isSet bool
}

func (v NullableIncomeStatementBlockNonOperatingInterest) Get() *IncomeStatementBlockNonOperatingInterest {
	return v.value
}

func (v *NullableIncomeStatementBlockNonOperatingInterest) Set(val *IncomeStatementBlockNonOperatingInterest) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeStatementBlockNonOperatingInterest) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeStatementBlockNonOperatingInterest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeStatementBlockNonOperatingInterest(val *IncomeStatementBlockNonOperatingInterest) *NullableIncomeStatementBlockNonOperatingInterest {
	return &NullableIncomeStatementBlockNonOperatingInterest{value: val, isSet: true}
}

func (v NullableIncomeStatementBlockNonOperatingInterest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeStatementBlockNonOperatingInterest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
