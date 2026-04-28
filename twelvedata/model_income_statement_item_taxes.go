/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the IncomeStatementItemTaxes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IncomeStatementItemTaxes{}

// IncomeStatementItemTaxes Taxes information
type IncomeStatementItemTaxes struct {
	// Tax provision
	TaxProvision *float64 `json:"tax_provision,omitempty"`
	// Tax effect of unusual items
	TaxEffectOfUnusualItems *float64 `json:"tax_effect_of_unusual_items,omitempty"`
	// Tax rate for calculations
	TaxRateForCalculations *float64 `json:"tax_rate_for_calculations,omitempty"`
	// Other taxes
	OtherTaxes *float64 `json:"other_taxes,omitempty"`
}

// NewIncomeStatementItemTaxes instantiates a new IncomeStatementItemTaxes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeStatementItemTaxes() *IncomeStatementItemTaxes {
	this := IncomeStatementItemTaxes{}
	return &this
}

// NewIncomeStatementItemTaxesWithDefaults instantiates a new IncomeStatementItemTaxes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeStatementItemTaxesWithDefaults() *IncomeStatementItemTaxes {
	this := IncomeStatementItemTaxes{}
	return &this
}

// GetTaxProvision returns the TaxProvision field value if set, zero value otherwise.
func (o *IncomeStatementItemTaxes) GetTaxProvision() float64 {
	if o == nil || IsNil(o.TaxProvision) {
		var ret float64
		return ret
	}
	return *o.TaxProvision
}

// GetTaxProvisionOk returns a tuple with the TaxProvision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemTaxes) GetTaxProvisionOk() (*float64, bool) {
	if o == nil || IsNil(o.TaxProvision) {
		return nil, false
	}
	return o.TaxProvision, true
}

// HasTaxProvision returns a boolean if a field has been set.
func (o *IncomeStatementItemTaxes) HasTaxProvision() bool {
	if o != nil && !IsNil(o.TaxProvision) {
		return true
	}

	return false
}

// SetTaxProvision gets a reference to the given float64 and assigns it to the TaxProvision field.
func (o *IncomeStatementItemTaxes) SetTaxProvision(v float64) {
	o.TaxProvision = &v
}

// GetTaxEffectOfUnusualItems returns the TaxEffectOfUnusualItems field value if set, zero value otherwise.
func (o *IncomeStatementItemTaxes) GetTaxEffectOfUnusualItems() float64 {
	if o == nil || IsNil(o.TaxEffectOfUnusualItems) {
		var ret float64
		return ret
	}
	return *o.TaxEffectOfUnusualItems
}

// GetTaxEffectOfUnusualItemsOk returns a tuple with the TaxEffectOfUnusualItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemTaxes) GetTaxEffectOfUnusualItemsOk() (*float64, bool) {
	if o == nil || IsNil(o.TaxEffectOfUnusualItems) {
		return nil, false
	}
	return o.TaxEffectOfUnusualItems, true
}

// HasTaxEffectOfUnusualItems returns a boolean if a field has been set.
func (o *IncomeStatementItemTaxes) HasTaxEffectOfUnusualItems() bool {
	if o != nil && !IsNil(o.TaxEffectOfUnusualItems) {
		return true
	}

	return false
}

// SetTaxEffectOfUnusualItems gets a reference to the given float64 and assigns it to the TaxEffectOfUnusualItems field.
func (o *IncomeStatementItemTaxes) SetTaxEffectOfUnusualItems(v float64) {
	o.TaxEffectOfUnusualItems = &v
}

// GetTaxRateForCalculations returns the TaxRateForCalculations field value if set, zero value otherwise.
func (o *IncomeStatementItemTaxes) GetTaxRateForCalculations() float64 {
	if o == nil || IsNil(o.TaxRateForCalculations) {
		var ret float64
		return ret
	}
	return *o.TaxRateForCalculations
}

// GetTaxRateForCalculationsOk returns a tuple with the TaxRateForCalculations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemTaxes) GetTaxRateForCalculationsOk() (*float64, bool) {
	if o == nil || IsNil(o.TaxRateForCalculations) {
		return nil, false
	}
	return o.TaxRateForCalculations, true
}

// HasTaxRateForCalculations returns a boolean if a field has been set.
func (o *IncomeStatementItemTaxes) HasTaxRateForCalculations() bool {
	if o != nil && !IsNil(o.TaxRateForCalculations) {
		return true
	}

	return false
}

// SetTaxRateForCalculations gets a reference to the given float64 and assigns it to the TaxRateForCalculations field.
func (o *IncomeStatementItemTaxes) SetTaxRateForCalculations(v float64) {
	o.TaxRateForCalculations = &v
}

// GetOtherTaxes returns the OtherTaxes field value if set, zero value otherwise.
func (o *IncomeStatementItemTaxes) GetOtherTaxes() float64 {
	if o == nil || IsNil(o.OtherTaxes) {
		var ret float64
		return ret
	}
	return *o.OtherTaxes
}

// GetOtherTaxesOk returns a tuple with the OtherTaxes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemTaxes) GetOtherTaxesOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherTaxes) {
		return nil, false
	}
	return o.OtherTaxes, true
}

// HasOtherTaxes returns a boolean if a field has been set.
func (o *IncomeStatementItemTaxes) HasOtherTaxes() bool {
	if o != nil && !IsNil(o.OtherTaxes) {
		return true
	}

	return false
}

// SetOtherTaxes gets a reference to the given float64 and assigns it to the OtherTaxes field.
func (o *IncomeStatementItemTaxes) SetOtherTaxes(v float64) {
	o.OtherTaxes = &v
}

func (o IncomeStatementItemTaxes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IncomeStatementItemTaxes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TaxProvision) {
		toSerialize["tax_provision"] = o.TaxProvision
	}
	if !IsNil(o.TaxEffectOfUnusualItems) {
		toSerialize["tax_effect_of_unusual_items"] = o.TaxEffectOfUnusualItems
	}
	if !IsNil(o.TaxRateForCalculations) {
		toSerialize["tax_rate_for_calculations"] = o.TaxRateForCalculations
	}
	if !IsNil(o.OtherTaxes) {
		toSerialize["other_taxes"] = o.OtherTaxes
	}
	return toSerialize, nil
}

type NullableIncomeStatementItemTaxes struct {
	value *IncomeStatementItemTaxes
	isSet bool
}

func (v NullableIncomeStatementItemTaxes) Get() *IncomeStatementItemTaxes {
	return v.value
}

func (v *NullableIncomeStatementItemTaxes) Set(val *IncomeStatementItemTaxes) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeStatementItemTaxes) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeStatementItemTaxes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeStatementItemTaxes(val *IncomeStatementItemTaxes) *NullableIncomeStatementItemTaxes {
	return &NullableIncomeStatementItemTaxes{value: val, isSet: true}
}

func (v NullableIncomeStatementItemTaxes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeStatementItemTaxes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
