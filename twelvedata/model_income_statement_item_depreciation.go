// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the IncomeStatementItemDepreciation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IncomeStatementItemDepreciation{}

// IncomeStatementItemDepreciation Depreciation information
type IncomeStatementItemDepreciation struct {
	// Reconciled depreciation
	ReconciledDepreciation *float64 `json:"reconciled_depreciation,omitempty"`
}

// NewIncomeStatementItemDepreciation instantiates a new IncomeStatementItemDepreciation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeStatementItemDepreciation() *IncomeStatementItemDepreciation {
	this := IncomeStatementItemDepreciation{}
	return &this
}

// NewIncomeStatementItemDepreciationWithDefaults instantiates a new IncomeStatementItemDepreciation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeStatementItemDepreciationWithDefaults() *IncomeStatementItemDepreciation {
	this := IncomeStatementItemDepreciation{}
	return &this
}

// GetReconciledDepreciation returns the ReconciledDepreciation field value if set, zero value otherwise.
func (o *IncomeStatementItemDepreciation) GetReconciledDepreciation() float64 {
	if o == nil || IsNil(o.ReconciledDepreciation) {
		var ret float64
		return ret
	}
	return *o.ReconciledDepreciation
}

// GetReconciledDepreciationOk returns a tuple with the ReconciledDepreciation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemDepreciation) GetReconciledDepreciationOk() (*float64, bool) {
	if o == nil || IsNil(o.ReconciledDepreciation) {
		return nil, false
	}
	return o.ReconciledDepreciation, true
}

// HasReconciledDepreciation returns a boolean if a field has been set.
func (o *IncomeStatementItemDepreciation) HasReconciledDepreciation() bool {
	if o != nil && !IsNil(o.ReconciledDepreciation) {
		return true
	}

	return false
}

// SetReconciledDepreciation gets a reference to the given float64 and assigns it to the ReconciledDepreciation field.
func (o *IncomeStatementItemDepreciation) SetReconciledDepreciation(v float64) {
	o.ReconciledDepreciation = &v
}

func (o IncomeStatementItemDepreciation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IncomeStatementItemDepreciation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReconciledDepreciation) {
		toSerialize["reconciled_depreciation"] = o.ReconciledDepreciation
	}
	return toSerialize, nil
}

type NullableIncomeStatementItemDepreciation struct {
	value *IncomeStatementItemDepreciation
	isSet bool
}

func (v NullableIncomeStatementItemDepreciation) Get() *IncomeStatementItemDepreciation {
	return v.value
}

func (v *NullableIncomeStatementItemDepreciation) Set(val *IncomeStatementItemDepreciation) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeStatementItemDepreciation) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeStatementItemDepreciation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeStatementItemDepreciation(val *IncomeStatementItemDepreciation) *NullableIncomeStatementItemDepreciation {
	return &NullableIncomeStatementItemDepreciation{value: val, isSet: true}
}

func (v NullableIncomeStatementItemDepreciation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeStatementItemDepreciation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
