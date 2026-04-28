/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the IncomeStatementItemOperatingIncome type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IncomeStatementItemOperatingIncome{}

// IncomeStatementItemOperatingIncome Operating income information
type IncomeStatementItemOperatingIncome struct {
	// Operating income value
	OperatingIncomeValue *float64 `json:"operating_income_value,omitempty"`
	// Total operating income as reported
	TotalOperatingIncomeAsReported *float64 `json:"total_operating_income_as_reported,omitempty"`
	// Operating expense
	OperatingExpense *float64 `json:"operating_expense,omitempty"`
	// Other operating expenses
	OtherOperatingExpenses *float64 `json:"other_operating_expenses,omitempty"`
	// Total expenses
	TotalExpenses *float64 `json:"total_expenses,omitempty"`
}

// NewIncomeStatementItemOperatingIncome instantiates a new IncomeStatementItemOperatingIncome object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeStatementItemOperatingIncome() *IncomeStatementItemOperatingIncome {
	this := IncomeStatementItemOperatingIncome{}
	return &this
}

// NewIncomeStatementItemOperatingIncomeWithDefaults instantiates a new IncomeStatementItemOperatingIncome object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeStatementItemOperatingIncomeWithDefaults() *IncomeStatementItemOperatingIncome {
	this := IncomeStatementItemOperatingIncome{}
	return &this
}

// GetOperatingIncomeValue returns the OperatingIncomeValue field value if set, zero value otherwise.
func (o *IncomeStatementItemOperatingIncome) GetOperatingIncomeValue() float64 {
	if o == nil || IsNil(o.OperatingIncomeValue) {
		var ret float64
		return ret
	}
	return *o.OperatingIncomeValue
}

// GetOperatingIncomeValueOk returns a tuple with the OperatingIncomeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemOperatingIncome) GetOperatingIncomeValueOk() (*float64, bool) {
	if o == nil || IsNil(o.OperatingIncomeValue) {
		return nil, false
	}
	return o.OperatingIncomeValue, true
}

// HasOperatingIncomeValue returns a boolean if a field has been set.
func (o *IncomeStatementItemOperatingIncome) HasOperatingIncomeValue() bool {
	if o != nil && !IsNil(o.OperatingIncomeValue) {
		return true
	}

	return false
}

// SetOperatingIncomeValue gets a reference to the given float64 and assigns it to the OperatingIncomeValue field.
func (o *IncomeStatementItemOperatingIncome) SetOperatingIncomeValue(v float64) {
	o.OperatingIncomeValue = &v
}

// GetTotalOperatingIncomeAsReported returns the TotalOperatingIncomeAsReported field value if set, zero value otherwise.
func (o *IncomeStatementItemOperatingIncome) GetTotalOperatingIncomeAsReported() float64 {
	if o == nil || IsNil(o.TotalOperatingIncomeAsReported) {
		var ret float64
		return ret
	}
	return *o.TotalOperatingIncomeAsReported
}

// GetTotalOperatingIncomeAsReportedOk returns a tuple with the TotalOperatingIncomeAsReported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemOperatingIncome) GetTotalOperatingIncomeAsReportedOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalOperatingIncomeAsReported) {
		return nil, false
	}
	return o.TotalOperatingIncomeAsReported, true
}

// HasTotalOperatingIncomeAsReported returns a boolean if a field has been set.
func (o *IncomeStatementItemOperatingIncome) HasTotalOperatingIncomeAsReported() bool {
	if o != nil && !IsNil(o.TotalOperatingIncomeAsReported) {
		return true
	}

	return false
}

// SetTotalOperatingIncomeAsReported gets a reference to the given float64 and assigns it to the TotalOperatingIncomeAsReported field.
func (o *IncomeStatementItemOperatingIncome) SetTotalOperatingIncomeAsReported(v float64) {
	o.TotalOperatingIncomeAsReported = &v
}

// GetOperatingExpense returns the OperatingExpense field value if set, zero value otherwise.
func (o *IncomeStatementItemOperatingIncome) GetOperatingExpense() float64 {
	if o == nil || IsNil(o.OperatingExpense) {
		var ret float64
		return ret
	}
	return *o.OperatingExpense
}

// GetOperatingExpenseOk returns a tuple with the OperatingExpense field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemOperatingIncome) GetOperatingExpenseOk() (*float64, bool) {
	if o == nil || IsNil(o.OperatingExpense) {
		return nil, false
	}
	return o.OperatingExpense, true
}

// HasOperatingExpense returns a boolean if a field has been set.
func (o *IncomeStatementItemOperatingIncome) HasOperatingExpense() bool {
	if o != nil && !IsNil(o.OperatingExpense) {
		return true
	}

	return false
}

// SetOperatingExpense gets a reference to the given float64 and assigns it to the OperatingExpense field.
func (o *IncomeStatementItemOperatingIncome) SetOperatingExpense(v float64) {
	o.OperatingExpense = &v
}

// GetOtherOperatingExpenses returns the OtherOperatingExpenses field value if set, zero value otherwise.
func (o *IncomeStatementItemOperatingIncome) GetOtherOperatingExpenses() float64 {
	if o == nil || IsNil(o.OtherOperatingExpenses) {
		var ret float64
		return ret
	}
	return *o.OtherOperatingExpenses
}

// GetOtherOperatingExpensesOk returns a tuple with the OtherOperatingExpenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemOperatingIncome) GetOtherOperatingExpensesOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherOperatingExpenses) {
		return nil, false
	}
	return o.OtherOperatingExpenses, true
}

// HasOtherOperatingExpenses returns a boolean if a field has been set.
func (o *IncomeStatementItemOperatingIncome) HasOtherOperatingExpenses() bool {
	if o != nil && !IsNil(o.OtherOperatingExpenses) {
		return true
	}

	return false
}

// SetOtherOperatingExpenses gets a reference to the given float64 and assigns it to the OtherOperatingExpenses field.
func (o *IncomeStatementItemOperatingIncome) SetOtherOperatingExpenses(v float64) {
	o.OtherOperatingExpenses = &v
}

// GetTotalExpenses returns the TotalExpenses field value if set, zero value otherwise.
func (o *IncomeStatementItemOperatingIncome) GetTotalExpenses() float64 {
	if o == nil || IsNil(o.TotalExpenses) {
		var ret float64
		return ret
	}
	return *o.TotalExpenses
}

// GetTotalExpensesOk returns a tuple with the TotalExpenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemOperatingIncome) GetTotalExpensesOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalExpenses) {
		return nil, false
	}
	return o.TotalExpenses, true
}

// HasTotalExpenses returns a boolean if a field has been set.
func (o *IncomeStatementItemOperatingIncome) HasTotalExpenses() bool {
	if o != nil && !IsNil(o.TotalExpenses) {
		return true
	}

	return false
}

// SetTotalExpenses gets a reference to the given float64 and assigns it to the TotalExpenses field.
func (o *IncomeStatementItemOperatingIncome) SetTotalExpenses(v float64) {
	o.TotalExpenses = &v
}

func (o IncomeStatementItemOperatingIncome) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IncomeStatementItemOperatingIncome) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OperatingIncomeValue) {
		toSerialize["operating_income_value"] = o.OperatingIncomeValue
	}
	if !IsNil(o.TotalOperatingIncomeAsReported) {
		toSerialize["total_operating_income_as_reported"] = o.TotalOperatingIncomeAsReported
	}
	if !IsNil(o.OperatingExpense) {
		toSerialize["operating_expense"] = o.OperatingExpense
	}
	if !IsNil(o.OtherOperatingExpenses) {
		toSerialize["other_operating_expenses"] = o.OtherOperatingExpenses
	}
	if !IsNil(o.TotalExpenses) {
		toSerialize["total_expenses"] = o.TotalExpenses
	}
	return toSerialize, nil
}

type NullableIncomeStatementItemOperatingIncome struct {
	value *IncomeStatementItemOperatingIncome
	isSet bool
}

func (v NullableIncomeStatementItemOperatingIncome) Get() *IncomeStatementItemOperatingIncome {
	return v.value
}

func (v *NullableIncomeStatementItemOperatingIncome) Set(val *IncomeStatementItemOperatingIncome) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeStatementItemOperatingIncome) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeStatementItemOperatingIncome) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeStatementItemOperatingIncome(val *IncomeStatementItemOperatingIncome) *NullableIncomeStatementItemOperatingIncome {
	return &NullableIncomeStatementItemOperatingIncome{value: val, isSet: true}
}

func (v NullableIncomeStatementItemOperatingIncome) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeStatementItemOperatingIncome) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
