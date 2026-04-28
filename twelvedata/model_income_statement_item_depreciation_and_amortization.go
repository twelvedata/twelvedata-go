/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the IncomeStatementItemDepreciationAndAmortization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IncomeStatementItemDepreciationAndAmortization{}

// IncomeStatementItemDepreciationAndAmortization Depreciation and amortization information
type IncomeStatementItemDepreciationAndAmortization struct {
	// Depreciation amortization depletion
	DepreciationAmortizationDepletion *float64 `json:"depreciation_amortization_depletion,omitempty"`
	// Amortization of intangibles
	AmortizationOfIntangibles *float64 `json:"amortization_of_intangibles,omitempty"`
	// Depreciation
	Depreciation *float64 `json:"depreciation,omitempty"`
	// Amortization
	Amortization *float64 `json:"amortization,omitempty"`
	// Depletion
	Depletion *float64 `json:"depletion,omitempty"`
	// Depreciation and amortization in income statement
	DepreciationAndAmortizationInIncomeStatement *float64 `json:"depreciation_and_amortization_in_income_statement,omitempty"`
}

// NewIncomeStatementItemDepreciationAndAmortization instantiates a new IncomeStatementItemDepreciationAndAmortization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeStatementItemDepreciationAndAmortization() *IncomeStatementItemDepreciationAndAmortization {
	this := IncomeStatementItemDepreciationAndAmortization{}
	return &this
}

// NewIncomeStatementItemDepreciationAndAmortizationWithDefaults instantiates a new IncomeStatementItemDepreciationAndAmortization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeStatementItemDepreciationAndAmortizationWithDefaults() *IncomeStatementItemDepreciationAndAmortization {
	this := IncomeStatementItemDepreciationAndAmortization{}
	return &this
}

// GetDepreciationAmortizationDepletion returns the DepreciationAmortizationDepletion field value if set, zero value otherwise.
func (o *IncomeStatementItemDepreciationAndAmortization) GetDepreciationAmortizationDepletion() float64 {
	if o == nil || IsNil(o.DepreciationAmortizationDepletion) {
		var ret float64
		return ret
	}
	return *o.DepreciationAmortizationDepletion
}

// GetDepreciationAmortizationDepletionOk returns a tuple with the DepreciationAmortizationDepletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemDepreciationAndAmortization) GetDepreciationAmortizationDepletionOk() (*float64, bool) {
	if o == nil || IsNil(o.DepreciationAmortizationDepletion) {
		return nil, false
	}
	return o.DepreciationAmortizationDepletion, true
}

// HasDepreciationAmortizationDepletion returns a boolean if a field has been set.
func (o *IncomeStatementItemDepreciationAndAmortization) HasDepreciationAmortizationDepletion() bool {
	if o != nil && !IsNil(o.DepreciationAmortizationDepletion) {
		return true
	}

	return false
}

// SetDepreciationAmortizationDepletion gets a reference to the given float64 and assigns it to the DepreciationAmortizationDepletion field.
func (o *IncomeStatementItemDepreciationAndAmortization) SetDepreciationAmortizationDepletion(v float64) {
	o.DepreciationAmortizationDepletion = &v
}

// GetAmortizationOfIntangibles returns the AmortizationOfIntangibles field value if set, zero value otherwise.
func (o *IncomeStatementItemDepreciationAndAmortization) GetAmortizationOfIntangibles() float64 {
	if o == nil || IsNil(o.AmortizationOfIntangibles) {
		var ret float64
		return ret
	}
	return *o.AmortizationOfIntangibles
}

// GetAmortizationOfIntangiblesOk returns a tuple with the AmortizationOfIntangibles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemDepreciationAndAmortization) GetAmortizationOfIntangiblesOk() (*float64, bool) {
	if o == nil || IsNil(o.AmortizationOfIntangibles) {
		return nil, false
	}
	return o.AmortizationOfIntangibles, true
}

// HasAmortizationOfIntangibles returns a boolean if a field has been set.
func (o *IncomeStatementItemDepreciationAndAmortization) HasAmortizationOfIntangibles() bool {
	if o != nil && !IsNil(o.AmortizationOfIntangibles) {
		return true
	}

	return false
}

// SetAmortizationOfIntangibles gets a reference to the given float64 and assigns it to the AmortizationOfIntangibles field.
func (o *IncomeStatementItemDepreciationAndAmortization) SetAmortizationOfIntangibles(v float64) {
	o.AmortizationOfIntangibles = &v
}

// GetDepreciation returns the Depreciation field value if set, zero value otherwise.
func (o *IncomeStatementItemDepreciationAndAmortization) GetDepreciation() float64 {
	if o == nil || IsNil(o.Depreciation) {
		var ret float64
		return ret
	}
	return *o.Depreciation
}

// GetDepreciationOk returns a tuple with the Depreciation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemDepreciationAndAmortization) GetDepreciationOk() (*float64, bool) {
	if o == nil || IsNil(o.Depreciation) {
		return nil, false
	}
	return o.Depreciation, true
}

// HasDepreciation returns a boolean if a field has been set.
func (o *IncomeStatementItemDepreciationAndAmortization) HasDepreciation() bool {
	if o != nil && !IsNil(o.Depreciation) {
		return true
	}

	return false
}

// SetDepreciation gets a reference to the given float64 and assigns it to the Depreciation field.
func (o *IncomeStatementItemDepreciationAndAmortization) SetDepreciation(v float64) {
	o.Depreciation = &v
}

// GetAmortization returns the Amortization field value if set, zero value otherwise.
func (o *IncomeStatementItemDepreciationAndAmortization) GetAmortization() float64 {
	if o == nil || IsNil(o.Amortization) {
		var ret float64
		return ret
	}
	return *o.Amortization
}

// GetAmortizationOk returns a tuple with the Amortization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemDepreciationAndAmortization) GetAmortizationOk() (*float64, bool) {
	if o == nil || IsNil(o.Amortization) {
		return nil, false
	}
	return o.Amortization, true
}

// HasAmortization returns a boolean if a field has been set.
func (o *IncomeStatementItemDepreciationAndAmortization) HasAmortization() bool {
	if o != nil && !IsNil(o.Amortization) {
		return true
	}

	return false
}

// SetAmortization gets a reference to the given float64 and assigns it to the Amortization field.
func (o *IncomeStatementItemDepreciationAndAmortization) SetAmortization(v float64) {
	o.Amortization = &v
}

// GetDepletion returns the Depletion field value if set, zero value otherwise.
func (o *IncomeStatementItemDepreciationAndAmortization) GetDepletion() float64 {
	if o == nil || IsNil(o.Depletion) {
		var ret float64
		return ret
	}
	return *o.Depletion
}

// GetDepletionOk returns a tuple with the Depletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemDepreciationAndAmortization) GetDepletionOk() (*float64, bool) {
	if o == nil || IsNil(o.Depletion) {
		return nil, false
	}
	return o.Depletion, true
}

// HasDepletion returns a boolean if a field has been set.
func (o *IncomeStatementItemDepreciationAndAmortization) HasDepletion() bool {
	if o != nil && !IsNil(o.Depletion) {
		return true
	}

	return false
}

// SetDepletion gets a reference to the given float64 and assigns it to the Depletion field.
func (o *IncomeStatementItemDepreciationAndAmortization) SetDepletion(v float64) {
	o.Depletion = &v
}

// GetDepreciationAndAmortizationInIncomeStatement returns the DepreciationAndAmortizationInIncomeStatement field value if set, zero value otherwise.
func (o *IncomeStatementItemDepreciationAndAmortization) GetDepreciationAndAmortizationInIncomeStatement() float64 {
	if o == nil || IsNil(o.DepreciationAndAmortizationInIncomeStatement) {
		var ret float64
		return ret
	}
	return *o.DepreciationAndAmortizationInIncomeStatement
}

// GetDepreciationAndAmortizationInIncomeStatementOk returns a tuple with the DepreciationAndAmortizationInIncomeStatement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemDepreciationAndAmortization) GetDepreciationAndAmortizationInIncomeStatementOk() (*float64, bool) {
	if o == nil || IsNil(o.DepreciationAndAmortizationInIncomeStatement) {
		return nil, false
	}
	return o.DepreciationAndAmortizationInIncomeStatement, true
}

// HasDepreciationAndAmortizationInIncomeStatement returns a boolean if a field has been set.
func (o *IncomeStatementItemDepreciationAndAmortization) HasDepreciationAndAmortizationInIncomeStatement() bool {
	if o != nil && !IsNil(o.DepreciationAndAmortizationInIncomeStatement) {
		return true
	}

	return false
}

// SetDepreciationAndAmortizationInIncomeStatement gets a reference to the given float64 and assigns it to the DepreciationAndAmortizationInIncomeStatement field.
func (o *IncomeStatementItemDepreciationAndAmortization) SetDepreciationAndAmortizationInIncomeStatement(v float64) {
	o.DepreciationAndAmortizationInIncomeStatement = &v
}

func (o IncomeStatementItemDepreciationAndAmortization) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IncomeStatementItemDepreciationAndAmortization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DepreciationAmortizationDepletion) {
		toSerialize["depreciation_amortization_depletion"] = o.DepreciationAmortizationDepletion
	}
	if !IsNil(o.AmortizationOfIntangibles) {
		toSerialize["amortization_of_intangibles"] = o.AmortizationOfIntangibles
	}
	if !IsNil(o.Depreciation) {
		toSerialize["depreciation"] = o.Depreciation
	}
	if !IsNil(o.Amortization) {
		toSerialize["amortization"] = o.Amortization
	}
	if !IsNil(o.Depletion) {
		toSerialize["depletion"] = o.Depletion
	}
	if !IsNil(o.DepreciationAndAmortizationInIncomeStatement) {
		toSerialize["depreciation_and_amortization_in_income_statement"] = o.DepreciationAndAmortizationInIncomeStatement
	}
	return toSerialize, nil
}

type NullableIncomeStatementItemDepreciationAndAmortization struct {
	value *IncomeStatementItemDepreciationAndAmortization
	isSet bool
}

func (v NullableIncomeStatementItemDepreciationAndAmortization) Get() *IncomeStatementItemDepreciationAndAmortization {
	return v.value
}

func (v *NullableIncomeStatementItemDepreciationAndAmortization) Set(val *IncomeStatementItemDepreciationAndAmortization) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeStatementItemDepreciationAndAmortization) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeStatementItemDepreciationAndAmortization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeStatementItemDepreciationAndAmortization(val *IncomeStatementItemDepreciationAndAmortization) *NullableIncomeStatementItemDepreciationAndAmortization {
	return &NullableIncomeStatementItemDepreciationAndAmortization{value: val, isSet: true}
}

func (v NullableIncomeStatementItemDepreciationAndAmortization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeStatementItemDepreciationAndAmortization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
