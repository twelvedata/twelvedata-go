/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the IncomeStatementBlockOperatingExpense type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IncomeStatementBlockOperatingExpense{}

// IncomeStatementBlockOperatingExpense Operating expense details
type IncomeStatementBlockOperatingExpense struct {
	// Refers to research & development (R&D) expenses
	ResearchAndDevelopment *int64 `json:"research_and_development,omitempty"`
	// Refers to selling, general and administrative (SG&A) expenses
	SellingGeneralAndAdministrative *int64 `json:"selling_general_and_administrative,omitempty"`
	// Refers to other operating expenses
	OtherOperatingExpenses *int64 `json:"other_operating_expenses,omitempty"`
}

// NewIncomeStatementBlockOperatingExpense instantiates a new IncomeStatementBlockOperatingExpense object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeStatementBlockOperatingExpense() *IncomeStatementBlockOperatingExpense {
	this := IncomeStatementBlockOperatingExpense{}
	return &this
}

// NewIncomeStatementBlockOperatingExpenseWithDefaults instantiates a new IncomeStatementBlockOperatingExpense object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeStatementBlockOperatingExpenseWithDefaults() *IncomeStatementBlockOperatingExpense {
	this := IncomeStatementBlockOperatingExpense{}
	return &this
}

// GetResearchAndDevelopment returns the ResearchAndDevelopment field value if set, zero value otherwise.
func (o *IncomeStatementBlockOperatingExpense) GetResearchAndDevelopment() int64 {
	if o == nil || IsNil(o.ResearchAndDevelopment) {
		var ret int64
		return ret
	}
	return *o.ResearchAndDevelopment
}

// GetResearchAndDevelopmentOk returns a tuple with the ResearchAndDevelopment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementBlockOperatingExpense) GetResearchAndDevelopmentOk() (*int64, bool) {
	if o == nil || IsNil(o.ResearchAndDevelopment) {
		return nil, false
	}
	return o.ResearchAndDevelopment, true
}

// HasResearchAndDevelopment returns a boolean if a field has been set.
func (o *IncomeStatementBlockOperatingExpense) HasResearchAndDevelopment() bool {
	if o != nil && !IsNil(o.ResearchAndDevelopment) {
		return true
	}

	return false
}

// SetResearchAndDevelopment gets a reference to the given int64 and assigns it to the ResearchAndDevelopment field.
func (o *IncomeStatementBlockOperatingExpense) SetResearchAndDevelopment(v int64) {
	o.ResearchAndDevelopment = &v
}

// GetSellingGeneralAndAdministrative returns the SellingGeneralAndAdministrative field value if set, zero value otherwise.
func (o *IncomeStatementBlockOperatingExpense) GetSellingGeneralAndAdministrative() int64 {
	if o == nil || IsNil(o.SellingGeneralAndAdministrative) {
		var ret int64
		return ret
	}
	return *o.SellingGeneralAndAdministrative
}

// GetSellingGeneralAndAdministrativeOk returns a tuple with the SellingGeneralAndAdministrative field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementBlockOperatingExpense) GetSellingGeneralAndAdministrativeOk() (*int64, bool) {
	if o == nil || IsNil(o.SellingGeneralAndAdministrative) {
		return nil, false
	}
	return o.SellingGeneralAndAdministrative, true
}

// HasSellingGeneralAndAdministrative returns a boolean if a field has been set.
func (o *IncomeStatementBlockOperatingExpense) HasSellingGeneralAndAdministrative() bool {
	if o != nil && !IsNil(o.SellingGeneralAndAdministrative) {
		return true
	}

	return false
}

// SetSellingGeneralAndAdministrative gets a reference to the given int64 and assigns it to the SellingGeneralAndAdministrative field.
func (o *IncomeStatementBlockOperatingExpense) SetSellingGeneralAndAdministrative(v int64) {
	o.SellingGeneralAndAdministrative = &v
}

// GetOtherOperatingExpenses returns the OtherOperatingExpenses field value if set, zero value otherwise.
func (o *IncomeStatementBlockOperatingExpense) GetOtherOperatingExpenses() int64 {
	if o == nil || IsNil(o.OtherOperatingExpenses) {
		var ret int64
		return ret
	}
	return *o.OtherOperatingExpenses
}

// GetOtherOperatingExpensesOk returns a tuple with the OtherOperatingExpenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementBlockOperatingExpense) GetOtherOperatingExpensesOk() (*int64, bool) {
	if o == nil || IsNil(o.OtherOperatingExpenses) {
		return nil, false
	}
	return o.OtherOperatingExpenses, true
}

// HasOtherOperatingExpenses returns a boolean if a field has been set.
func (o *IncomeStatementBlockOperatingExpense) HasOtherOperatingExpenses() bool {
	if o != nil && !IsNil(o.OtherOperatingExpenses) {
		return true
	}

	return false
}

// SetOtherOperatingExpenses gets a reference to the given int64 and assigns it to the OtherOperatingExpenses field.
func (o *IncomeStatementBlockOperatingExpense) SetOtherOperatingExpenses(v int64) {
	o.OtherOperatingExpenses = &v
}

func (o IncomeStatementBlockOperatingExpense) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IncomeStatementBlockOperatingExpense) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResearchAndDevelopment) {
		toSerialize["research_and_development"] = o.ResearchAndDevelopment
	}
	if !IsNil(o.SellingGeneralAndAdministrative) {
		toSerialize["selling_general_and_administrative"] = o.SellingGeneralAndAdministrative
	}
	if !IsNil(o.OtherOperatingExpenses) {
		toSerialize["other_operating_expenses"] = o.OtherOperatingExpenses
	}
	return toSerialize, nil
}

type NullableIncomeStatementBlockOperatingExpense struct {
	value *IncomeStatementBlockOperatingExpense
	isSet bool
}

func (v NullableIncomeStatementBlockOperatingExpense) Get() *IncomeStatementBlockOperatingExpense {
	return v.value
}

func (v *NullableIncomeStatementBlockOperatingExpense) Set(val *IncomeStatementBlockOperatingExpense) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeStatementBlockOperatingExpense) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeStatementBlockOperatingExpense) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeStatementBlockOperatingExpense(val *IncomeStatementBlockOperatingExpense) *NullableIncomeStatementBlockOperatingExpense {
	return &NullableIncomeStatementBlockOperatingExpense{value: val, isSet: true}
}

func (v NullableIncomeStatementBlockOperatingExpense) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeStatementBlockOperatingExpense) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
