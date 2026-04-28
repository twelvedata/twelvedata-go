/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the IncomeStatementItemEbitda type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IncomeStatementItemEbitda{}

// IncomeStatementItemEbitda EBITDA information
type IncomeStatementItemEbitda struct {
	// EBITDA value
	EbitdaValue *float64 `json:"ebitda_value,omitempty"`
	// Normalized EBITDA value
	NormalizedEbitdaValue *float64 `json:"normalized_ebitda_value,omitempty"`
	// EBIT value
	EbitValue *float64 `json:"ebit_value,omitempty"`
}

// NewIncomeStatementItemEbitda instantiates a new IncomeStatementItemEbitda object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeStatementItemEbitda() *IncomeStatementItemEbitda {
	this := IncomeStatementItemEbitda{}
	return &this
}

// NewIncomeStatementItemEbitdaWithDefaults instantiates a new IncomeStatementItemEbitda object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeStatementItemEbitdaWithDefaults() *IncomeStatementItemEbitda {
	this := IncomeStatementItemEbitda{}
	return &this
}

// GetEbitdaValue returns the EbitdaValue field value if set, zero value otherwise.
func (o *IncomeStatementItemEbitda) GetEbitdaValue() float64 {
	if o == nil || IsNil(o.EbitdaValue) {
		var ret float64
		return ret
	}
	return *o.EbitdaValue
}

// GetEbitdaValueOk returns a tuple with the EbitdaValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEbitda) GetEbitdaValueOk() (*float64, bool) {
	if o == nil || IsNil(o.EbitdaValue) {
		return nil, false
	}
	return o.EbitdaValue, true
}

// HasEbitdaValue returns a boolean if a field has been set.
func (o *IncomeStatementItemEbitda) HasEbitdaValue() bool {
	if o != nil && !IsNil(o.EbitdaValue) {
		return true
	}

	return false
}

// SetEbitdaValue gets a reference to the given float64 and assigns it to the EbitdaValue field.
func (o *IncomeStatementItemEbitda) SetEbitdaValue(v float64) {
	o.EbitdaValue = &v
}

// GetNormalizedEbitdaValue returns the NormalizedEbitdaValue field value if set, zero value otherwise.
func (o *IncomeStatementItemEbitda) GetNormalizedEbitdaValue() float64 {
	if o == nil || IsNil(o.NormalizedEbitdaValue) {
		var ret float64
		return ret
	}
	return *o.NormalizedEbitdaValue
}

// GetNormalizedEbitdaValueOk returns a tuple with the NormalizedEbitdaValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEbitda) GetNormalizedEbitdaValueOk() (*float64, bool) {
	if o == nil || IsNil(o.NormalizedEbitdaValue) {
		return nil, false
	}
	return o.NormalizedEbitdaValue, true
}

// HasNormalizedEbitdaValue returns a boolean if a field has been set.
func (o *IncomeStatementItemEbitda) HasNormalizedEbitdaValue() bool {
	if o != nil && !IsNil(o.NormalizedEbitdaValue) {
		return true
	}

	return false
}

// SetNormalizedEbitdaValue gets a reference to the given float64 and assigns it to the NormalizedEbitdaValue field.
func (o *IncomeStatementItemEbitda) SetNormalizedEbitdaValue(v float64) {
	o.NormalizedEbitdaValue = &v
}

// GetEbitValue returns the EbitValue field value if set, zero value otherwise.
func (o *IncomeStatementItemEbitda) GetEbitValue() float64 {
	if o == nil || IsNil(o.EbitValue) {
		var ret float64
		return ret
	}
	return *o.EbitValue
}

// GetEbitValueOk returns a tuple with the EbitValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemEbitda) GetEbitValueOk() (*float64, bool) {
	if o == nil || IsNil(o.EbitValue) {
		return nil, false
	}
	return o.EbitValue, true
}

// HasEbitValue returns a boolean if a field has been set.
func (o *IncomeStatementItemEbitda) HasEbitValue() bool {
	if o != nil && !IsNil(o.EbitValue) {
		return true
	}

	return false
}

// SetEbitValue gets a reference to the given float64 and assigns it to the EbitValue field.
func (o *IncomeStatementItemEbitda) SetEbitValue(v float64) {
	o.EbitValue = &v
}

func (o IncomeStatementItemEbitda) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IncomeStatementItemEbitda) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EbitdaValue) {
		toSerialize["ebitda_value"] = o.EbitdaValue
	}
	if !IsNil(o.NormalizedEbitdaValue) {
		toSerialize["normalized_ebitda_value"] = o.NormalizedEbitdaValue
	}
	if !IsNil(o.EbitValue) {
		toSerialize["ebit_value"] = o.EbitValue
	}
	return toSerialize, nil
}

type NullableIncomeStatementItemEbitda struct {
	value *IncomeStatementItemEbitda
	isSet bool
}

func (v NullableIncomeStatementItemEbitda) Get() *IncomeStatementItemEbitda {
	return v.value
}

func (v *NullableIncomeStatementItemEbitda) Set(val *IncomeStatementItemEbitda) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeStatementItemEbitda) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeStatementItemEbitda) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeStatementItemEbitda(val *IncomeStatementItemEbitda) *NullableIncomeStatementItemEbitda {
	return &NullableIncomeStatementItemEbitda{value: val, isSet: true}
}

func (v NullableIncomeStatementItemEbitda) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeStatementItemEbitda) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
