/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the IncomeStatementItemSpecialIncomeCharges type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IncomeStatementItemSpecialIncomeCharges{}

// IncomeStatementItemSpecialIncomeCharges Special income charges information
type IncomeStatementItemSpecialIncomeCharges struct {
	// Special income charges value
	SpecialIncomeChargesValue *float64 `json:"special_income_charges_value,omitempty"`
}

// NewIncomeStatementItemSpecialIncomeCharges instantiates a new IncomeStatementItemSpecialIncomeCharges object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeStatementItemSpecialIncomeCharges() *IncomeStatementItemSpecialIncomeCharges {
	this := IncomeStatementItemSpecialIncomeCharges{}
	return &this
}

// NewIncomeStatementItemSpecialIncomeChargesWithDefaults instantiates a new IncomeStatementItemSpecialIncomeCharges object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeStatementItemSpecialIncomeChargesWithDefaults() *IncomeStatementItemSpecialIncomeCharges {
	this := IncomeStatementItemSpecialIncomeCharges{}
	return &this
}

// GetSpecialIncomeChargesValue returns the SpecialIncomeChargesValue field value if set, zero value otherwise.
func (o *IncomeStatementItemSpecialIncomeCharges) GetSpecialIncomeChargesValue() float64 {
	if o == nil || IsNil(o.SpecialIncomeChargesValue) {
		var ret float64
		return ret
	}
	return *o.SpecialIncomeChargesValue
}

// GetSpecialIncomeChargesValueOk returns a tuple with the SpecialIncomeChargesValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemSpecialIncomeCharges) GetSpecialIncomeChargesValueOk() (*float64, bool) {
	if o == nil || IsNil(o.SpecialIncomeChargesValue) {
		return nil, false
	}
	return o.SpecialIncomeChargesValue, true
}

// HasSpecialIncomeChargesValue returns a boolean if a field has been set.
func (o *IncomeStatementItemSpecialIncomeCharges) HasSpecialIncomeChargesValue() bool {
	if o != nil && !IsNil(o.SpecialIncomeChargesValue) {
		return true
	}

	return false
}

// SetSpecialIncomeChargesValue gets a reference to the given float64 and assigns it to the SpecialIncomeChargesValue field.
func (o *IncomeStatementItemSpecialIncomeCharges) SetSpecialIncomeChargesValue(v float64) {
	o.SpecialIncomeChargesValue = &v
}

func (o IncomeStatementItemSpecialIncomeCharges) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IncomeStatementItemSpecialIncomeCharges) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SpecialIncomeChargesValue) {
		toSerialize["special_income_charges_value"] = o.SpecialIncomeChargesValue
	}
	return toSerialize, nil
}

type NullableIncomeStatementItemSpecialIncomeCharges struct {
	value *IncomeStatementItemSpecialIncomeCharges
	isSet bool
}

func (v NullableIncomeStatementItemSpecialIncomeCharges) Get() *IncomeStatementItemSpecialIncomeCharges {
	return v.value
}

func (v *NullableIncomeStatementItemSpecialIncomeCharges) Set(val *IncomeStatementItemSpecialIncomeCharges) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeStatementItemSpecialIncomeCharges) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeStatementItemSpecialIncomeCharges) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeStatementItemSpecialIncomeCharges(val *IncomeStatementItemSpecialIncomeCharges) *NullableIncomeStatementItemSpecialIncomeCharges {
	return &NullableIncomeStatementItemSpecialIncomeCharges{value: val, isSet: true}
}

func (v NullableIncomeStatementItemSpecialIncomeCharges) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeStatementItemSpecialIncomeCharges) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
