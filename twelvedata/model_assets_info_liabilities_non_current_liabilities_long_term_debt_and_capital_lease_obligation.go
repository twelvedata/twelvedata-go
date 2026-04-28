/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the AssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation{}

// AssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation Long term debt and capital lease obligation information
type AssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation struct {
	// Total long term debt and capital lease obligation
	TotalLongTermDebtAndCapitalLeaseObligation *float64 `json:"total_long_term_debt_and_capital_lease_obligation,omitempty"`
	// Long term debt
	LongTermDebt *float64 `json:"long_term_debt,omitempty"`
	// Long term capital lease obligation
	LongTermCapitalLeaseObligation *float64 `json:"long_term_capital_lease_obligation,omitempty"`
}

// NewAssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation instantiates a new AssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation() *AssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation {
	this := AssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation{}
	return &this
}

// NewAssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligationWithDefaults instantiates a new AssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligationWithDefaults() *AssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation {
	this := AssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation{}
	return &this
}

// GetTotalLongTermDebtAndCapitalLeaseObligation returns the TotalLongTermDebtAndCapitalLeaseObligation field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation) GetTotalLongTermDebtAndCapitalLeaseObligation() float64 {
	if o == nil || IsNil(o.TotalLongTermDebtAndCapitalLeaseObligation) {
		var ret float64
		return ret
	}
	return *o.TotalLongTermDebtAndCapitalLeaseObligation
}

// GetTotalLongTermDebtAndCapitalLeaseObligationOk returns a tuple with the TotalLongTermDebtAndCapitalLeaseObligation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation) GetTotalLongTermDebtAndCapitalLeaseObligationOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalLongTermDebtAndCapitalLeaseObligation) {
		return nil, false
	}
	return o.TotalLongTermDebtAndCapitalLeaseObligation, true
}

// HasTotalLongTermDebtAndCapitalLeaseObligation returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation) HasTotalLongTermDebtAndCapitalLeaseObligation() bool {
	if o != nil && !IsNil(o.TotalLongTermDebtAndCapitalLeaseObligation) {
		return true
	}

	return false
}

// SetTotalLongTermDebtAndCapitalLeaseObligation gets a reference to the given float64 and assigns it to the TotalLongTermDebtAndCapitalLeaseObligation field.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation) SetTotalLongTermDebtAndCapitalLeaseObligation(v float64) {
	o.TotalLongTermDebtAndCapitalLeaseObligation = &v
}

// GetLongTermDebt returns the LongTermDebt field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation) GetLongTermDebt() float64 {
	if o == nil || IsNil(o.LongTermDebt) {
		var ret float64
		return ret
	}
	return *o.LongTermDebt
}

// GetLongTermDebtOk returns a tuple with the LongTermDebt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation) GetLongTermDebtOk() (*float64, bool) {
	if o == nil || IsNil(o.LongTermDebt) {
		return nil, false
	}
	return o.LongTermDebt, true
}

// HasLongTermDebt returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation) HasLongTermDebt() bool {
	if o != nil && !IsNil(o.LongTermDebt) {
		return true
	}

	return false
}

// SetLongTermDebt gets a reference to the given float64 and assigns it to the LongTermDebt field.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation) SetLongTermDebt(v float64) {
	o.LongTermDebt = &v
}

// GetLongTermCapitalLeaseObligation returns the LongTermCapitalLeaseObligation field value if set, zero value otherwise.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation) GetLongTermCapitalLeaseObligation() float64 {
	if o == nil || IsNil(o.LongTermCapitalLeaseObligation) {
		var ret float64
		return ret
	}
	return *o.LongTermCapitalLeaseObligation
}

// GetLongTermCapitalLeaseObligationOk returns a tuple with the LongTermCapitalLeaseObligation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation) GetLongTermCapitalLeaseObligationOk() (*float64, bool) {
	if o == nil || IsNil(o.LongTermCapitalLeaseObligation) {
		return nil, false
	}
	return o.LongTermCapitalLeaseObligation, true
}

// HasLongTermCapitalLeaseObligation returns a boolean if a field has been set.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation) HasLongTermCapitalLeaseObligation() bool {
	if o != nil && !IsNil(o.LongTermCapitalLeaseObligation) {
		return true
	}

	return false
}

// SetLongTermCapitalLeaseObligation gets a reference to the given float64 and assigns it to the LongTermCapitalLeaseObligation field.
func (o *AssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation) SetLongTermCapitalLeaseObligation(v float64) {
	o.LongTermCapitalLeaseObligation = &v
}

func (o AssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalLongTermDebtAndCapitalLeaseObligation) {
		toSerialize["total_long_term_debt_and_capital_lease_obligation"] = o.TotalLongTermDebtAndCapitalLeaseObligation
	}
	if !IsNil(o.LongTermDebt) {
		toSerialize["long_term_debt"] = o.LongTermDebt
	}
	if !IsNil(o.LongTermCapitalLeaseObligation) {
		toSerialize["long_term_capital_lease_obligation"] = o.LongTermCapitalLeaseObligation
	}
	return toSerialize, nil
}

type NullableAssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation struct {
	value *AssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation
	isSet bool
}

func (v NullableAssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation) Get() *AssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation {
	return v.value
}

func (v *NullableAssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation) Set(val *AssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation(val *AssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation) *NullableAssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation {
	return &NullableAssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation{value: val, isSet: true}
}

func (v NullableAssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetsInfoLiabilitiesNonCurrentLiabilitiesLongTermDebtAndCapitalLeaseObligation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
