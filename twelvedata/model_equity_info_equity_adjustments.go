// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the EquityInfoEquityAdjustments type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EquityInfoEquityAdjustments{}

// EquityInfoEquityAdjustments Equity adjustments information
type EquityInfoEquityAdjustments struct {
	// Gains losses not affecting retained earnings
	GainsLossesNotAffectingRetainedEarnings *float64 `json:"gains_losses_not_affecting_retained_earnings,omitempty"`
	// Other equity adjustments
	OtherEquityAdjustments *float64 `json:"other_equity_adjustments,omitempty"`
	// Fixed assets revaluation reserve
	FixedAssetsRevaluationReserve *float64 `json:"fixed_assets_revaluation_reserve,omitempty"`
	// Foreign currency translation adjustments
	ForeignCurrencyTranslationAdjustments *float64 `json:"foreign_currency_translation_adjustments,omitempty"`
	// Minimum pension liabilities
	MinimumPensionLiabilities *float64 `json:"minimum_pension_liabilities,omitempty"`
	// Unrealized gain loss
	UnrealizedGainLoss *float64 `json:"unrealized_gain_loss,omitempty"`
}

// NewEquityInfoEquityAdjustments instantiates a new EquityInfoEquityAdjustments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquityInfoEquityAdjustments() *EquityInfoEquityAdjustments {
	this := EquityInfoEquityAdjustments{}
	return &this
}

// NewEquityInfoEquityAdjustmentsWithDefaults instantiates a new EquityInfoEquityAdjustments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquityInfoEquityAdjustmentsWithDefaults() *EquityInfoEquityAdjustments {
	this := EquityInfoEquityAdjustments{}
	return &this
}

// GetGainsLossesNotAffectingRetainedEarnings returns the GainsLossesNotAffectingRetainedEarnings field value if set, zero value otherwise.
func (o *EquityInfoEquityAdjustments) GetGainsLossesNotAffectingRetainedEarnings() float64 {
	if o == nil || IsNil(o.GainsLossesNotAffectingRetainedEarnings) {
		var ret float64
		return ret
	}
	return *o.GainsLossesNotAffectingRetainedEarnings
}

// GetGainsLossesNotAffectingRetainedEarningsOk returns a tuple with the GainsLossesNotAffectingRetainedEarnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfoEquityAdjustments) GetGainsLossesNotAffectingRetainedEarningsOk() (*float64, bool) {
	if o == nil || IsNil(o.GainsLossesNotAffectingRetainedEarnings) {
		return nil, false
	}
	return o.GainsLossesNotAffectingRetainedEarnings, true
}

// HasGainsLossesNotAffectingRetainedEarnings returns a boolean if a field has been set.
func (o *EquityInfoEquityAdjustments) HasGainsLossesNotAffectingRetainedEarnings() bool {
	if o != nil && !IsNil(o.GainsLossesNotAffectingRetainedEarnings) {
		return true
	}

	return false
}

// SetGainsLossesNotAffectingRetainedEarnings gets a reference to the given float64 and assigns it to the GainsLossesNotAffectingRetainedEarnings field.
func (o *EquityInfoEquityAdjustments) SetGainsLossesNotAffectingRetainedEarnings(v float64) {
	o.GainsLossesNotAffectingRetainedEarnings = &v
}

// GetOtherEquityAdjustments returns the OtherEquityAdjustments field value if set, zero value otherwise.
func (o *EquityInfoEquityAdjustments) GetOtherEquityAdjustments() float64 {
	if o == nil || IsNil(o.OtherEquityAdjustments) {
		var ret float64
		return ret
	}
	return *o.OtherEquityAdjustments
}

// GetOtherEquityAdjustmentsOk returns a tuple with the OtherEquityAdjustments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfoEquityAdjustments) GetOtherEquityAdjustmentsOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherEquityAdjustments) {
		return nil, false
	}
	return o.OtherEquityAdjustments, true
}

// HasOtherEquityAdjustments returns a boolean if a field has been set.
func (o *EquityInfoEquityAdjustments) HasOtherEquityAdjustments() bool {
	if o != nil && !IsNil(o.OtherEquityAdjustments) {
		return true
	}

	return false
}

// SetOtherEquityAdjustments gets a reference to the given float64 and assigns it to the OtherEquityAdjustments field.
func (o *EquityInfoEquityAdjustments) SetOtherEquityAdjustments(v float64) {
	o.OtherEquityAdjustments = &v
}

// GetFixedAssetsRevaluationReserve returns the FixedAssetsRevaluationReserve field value if set, zero value otherwise.
func (o *EquityInfoEquityAdjustments) GetFixedAssetsRevaluationReserve() float64 {
	if o == nil || IsNil(o.FixedAssetsRevaluationReserve) {
		var ret float64
		return ret
	}
	return *o.FixedAssetsRevaluationReserve
}

// GetFixedAssetsRevaluationReserveOk returns a tuple with the FixedAssetsRevaluationReserve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfoEquityAdjustments) GetFixedAssetsRevaluationReserveOk() (*float64, bool) {
	if o == nil || IsNil(o.FixedAssetsRevaluationReserve) {
		return nil, false
	}
	return o.FixedAssetsRevaluationReserve, true
}

// HasFixedAssetsRevaluationReserve returns a boolean if a field has been set.
func (o *EquityInfoEquityAdjustments) HasFixedAssetsRevaluationReserve() bool {
	if o != nil && !IsNil(o.FixedAssetsRevaluationReserve) {
		return true
	}

	return false
}

// SetFixedAssetsRevaluationReserve gets a reference to the given float64 and assigns it to the FixedAssetsRevaluationReserve field.
func (o *EquityInfoEquityAdjustments) SetFixedAssetsRevaluationReserve(v float64) {
	o.FixedAssetsRevaluationReserve = &v
}

// GetForeignCurrencyTranslationAdjustments returns the ForeignCurrencyTranslationAdjustments field value if set, zero value otherwise.
func (o *EquityInfoEquityAdjustments) GetForeignCurrencyTranslationAdjustments() float64 {
	if o == nil || IsNil(o.ForeignCurrencyTranslationAdjustments) {
		var ret float64
		return ret
	}
	return *o.ForeignCurrencyTranslationAdjustments
}

// GetForeignCurrencyTranslationAdjustmentsOk returns a tuple with the ForeignCurrencyTranslationAdjustments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfoEquityAdjustments) GetForeignCurrencyTranslationAdjustmentsOk() (*float64, bool) {
	if o == nil || IsNil(o.ForeignCurrencyTranslationAdjustments) {
		return nil, false
	}
	return o.ForeignCurrencyTranslationAdjustments, true
}

// HasForeignCurrencyTranslationAdjustments returns a boolean if a field has been set.
func (o *EquityInfoEquityAdjustments) HasForeignCurrencyTranslationAdjustments() bool {
	if o != nil && !IsNil(o.ForeignCurrencyTranslationAdjustments) {
		return true
	}

	return false
}

// SetForeignCurrencyTranslationAdjustments gets a reference to the given float64 and assigns it to the ForeignCurrencyTranslationAdjustments field.
func (o *EquityInfoEquityAdjustments) SetForeignCurrencyTranslationAdjustments(v float64) {
	o.ForeignCurrencyTranslationAdjustments = &v
}

// GetMinimumPensionLiabilities returns the MinimumPensionLiabilities field value if set, zero value otherwise.
func (o *EquityInfoEquityAdjustments) GetMinimumPensionLiabilities() float64 {
	if o == nil || IsNil(o.MinimumPensionLiabilities) {
		var ret float64
		return ret
	}
	return *o.MinimumPensionLiabilities
}

// GetMinimumPensionLiabilitiesOk returns a tuple with the MinimumPensionLiabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfoEquityAdjustments) GetMinimumPensionLiabilitiesOk() (*float64, bool) {
	if o == nil || IsNil(o.MinimumPensionLiabilities) {
		return nil, false
	}
	return o.MinimumPensionLiabilities, true
}

// HasMinimumPensionLiabilities returns a boolean if a field has been set.
func (o *EquityInfoEquityAdjustments) HasMinimumPensionLiabilities() bool {
	if o != nil && !IsNil(o.MinimumPensionLiabilities) {
		return true
	}

	return false
}

// SetMinimumPensionLiabilities gets a reference to the given float64 and assigns it to the MinimumPensionLiabilities field.
func (o *EquityInfoEquityAdjustments) SetMinimumPensionLiabilities(v float64) {
	o.MinimumPensionLiabilities = &v
}

// GetUnrealizedGainLoss returns the UnrealizedGainLoss field value if set, zero value otherwise.
func (o *EquityInfoEquityAdjustments) GetUnrealizedGainLoss() float64 {
	if o == nil || IsNil(o.UnrealizedGainLoss) {
		var ret float64
		return ret
	}
	return *o.UnrealizedGainLoss
}

// GetUnrealizedGainLossOk returns a tuple with the UnrealizedGainLoss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfoEquityAdjustments) GetUnrealizedGainLossOk() (*float64, bool) {
	if o == nil || IsNil(o.UnrealizedGainLoss) {
		return nil, false
	}
	return o.UnrealizedGainLoss, true
}

// HasUnrealizedGainLoss returns a boolean if a field has been set.
func (o *EquityInfoEquityAdjustments) HasUnrealizedGainLoss() bool {
	if o != nil && !IsNil(o.UnrealizedGainLoss) {
		return true
	}

	return false
}

// SetUnrealizedGainLoss gets a reference to the given float64 and assigns it to the UnrealizedGainLoss field.
func (o *EquityInfoEquityAdjustments) SetUnrealizedGainLoss(v float64) {
	o.UnrealizedGainLoss = &v
}

func (o EquityInfoEquityAdjustments) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EquityInfoEquityAdjustments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GainsLossesNotAffectingRetainedEarnings) {
		toSerialize["gains_losses_not_affecting_retained_earnings"] = o.GainsLossesNotAffectingRetainedEarnings
	}
	if !IsNil(o.OtherEquityAdjustments) {
		toSerialize["other_equity_adjustments"] = o.OtherEquityAdjustments
	}
	if !IsNil(o.FixedAssetsRevaluationReserve) {
		toSerialize["fixed_assets_revaluation_reserve"] = o.FixedAssetsRevaluationReserve
	}
	if !IsNil(o.ForeignCurrencyTranslationAdjustments) {
		toSerialize["foreign_currency_translation_adjustments"] = o.ForeignCurrencyTranslationAdjustments
	}
	if !IsNil(o.MinimumPensionLiabilities) {
		toSerialize["minimum_pension_liabilities"] = o.MinimumPensionLiabilities
	}
	if !IsNil(o.UnrealizedGainLoss) {
		toSerialize["unrealized_gain_loss"] = o.UnrealizedGainLoss
	}
	return toSerialize, nil
}

type NullableEquityInfoEquityAdjustments struct {
	value *EquityInfoEquityAdjustments
	isSet bool
}

func (v NullableEquityInfoEquityAdjustments) Get() *EquityInfoEquityAdjustments {
	return v.value
}

func (v *NullableEquityInfoEquityAdjustments) Set(val *EquityInfoEquityAdjustments) {
	v.value = val
	v.isSet = true
}

func (v NullableEquityInfoEquityAdjustments) IsSet() bool {
	return v.isSet
}

func (v *NullableEquityInfoEquityAdjustments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquityInfoEquityAdjustments(val *EquityInfoEquityAdjustments) *NullableEquityInfoEquityAdjustments {
	return &NullableEquityInfoEquityAdjustments{value: val, isSet: true}
}

func (v NullableEquityInfoEquityAdjustments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquityInfoEquityAdjustments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
