// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the IncomeStatementItemPretaxIncome type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IncomeStatementItemPretaxIncome{}

// IncomeStatementItemPretaxIncome Pretax income information
type IncomeStatementItemPretaxIncome struct {
	// Pretax income value
	PretaxIncomeValue *float64 `json:"pretax_income_value,omitempty"`
}

// NewIncomeStatementItemPretaxIncome instantiates a new IncomeStatementItemPretaxIncome object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeStatementItemPretaxIncome() *IncomeStatementItemPretaxIncome {
	this := IncomeStatementItemPretaxIncome{}
	return &this
}

// NewIncomeStatementItemPretaxIncomeWithDefaults instantiates a new IncomeStatementItemPretaxIncome object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeStatementItemPretaxIncomeWithDefaults() *IncomeStatementItemPretaxIncome {
	this := IncomeStatementItemPretaxIncome{}
	return &this
}

// GetPretaxIncomeValue returns the PretaxIncomeValue field value if set, zero value otherwise.
func (o *IncomeStatementItemPretaxIncome) GetPretaxIncomeValue() float64 {
	if o == nil || IsNil(o.PretaxIncomeValue) {
		var ret float64
		return ret
	}
	return *o.PretaxIncomeValue
}

// GetPretaxIncomeValueOk returns a tuple with the PretaxIncomeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemPretaxIncome) GetPretaxIncomeValueOk() (*float64, bool) {
	if o == nil || IsNil(o.PretaxIncomeValue) {
		return nil, false
	}
	return o.PretaxIncomeValue, true
}

// HasPretaxIncomeValue returns a boolean if a field has been set.
func (o *IncomeStatementItemPretaxIncome) HasPretaxIncomeValue() bool {
	if o != nil && !IsNil(o.PretaxIncomeValue) {
		return true
	}

	return false
}

// SetPretaxIncomeValue gets a reference to the given float64 and assigns it to the PretaxIncomeValue field.
func (o *IncomeStatementItemPretaxIncome) SetPretaxIncomeValue(v float64) {
	o.PretaxIncomeValue = &v
}

func (o IncomeStatementItemPretaxIncome) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IncomeStatementItemPretaxIncome) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PretaxIncomeValue) {
		toSerialize["pretax_income_value"] = o.PretaxIncomeValue
	}
	return toSerialize, nil
}

type NullableIncomeStatementItemPretaxIncome struct {
	value *IncomeStatementItemPretaxIncome
	isSet bool
}

func (v NullableIncomeStatementItemPretaxIncome) Get() *IncomeStatementItemPretaxIncome {
	return v.value
}

func (v *NullableIncomeStatementItemPretaxIncome) Set(val *IncomeStatementItemPretaxIncome) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeStatementItemPretaxIncome) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeStatementItemPretaxIncome) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeStatementItemPretaxIncome(val *IncomeStatementItemPretaxIncome) *NullableIncomeStatementItemPretaxIncome {
	return &NullableIncomeStatementItemPretaxIncome{value: val, isSet: true}
}

func (v NullableIncomeStatementItemPretaxIncome) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeStatementItemPretaxIncome) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
