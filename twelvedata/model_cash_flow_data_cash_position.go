/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the CashFlowDataCashPosition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CashFlowDataCashPosition{}

// CashFlowDataCashPosition Cash position
type CashFlowDataCashPosition struct {
	// Beginning cash position
	BeginningCashPosition *float64 `json:"beginning_cash_position,omitempty"`
	// End cash position
	EndCashPosition *float64 `json:"end_cash_position,omitempty"`
	// Changes in cash
	ChangesInCash *float64 `json:"changes_in_cash,omitempty"`
	// Other cash adjustment outside change in cash
	OtherCashAdjustmentOutsideChangeInCash *float64 `json:"other_cash_adjustment_outside_change_in_cash,omitempty"`
	// Other cash adjustment inside change in cash
	OtherCashAdjustmentInsideChangeInCash *float64 `json:"other_cash_adjustment_inside_change_in_cash,omitempty"`
	// Effect of exchange rate changes
	EffectOfExchangeRateChanges *float64 `json:"effect_of_exchange_rate_changes,omitempty"`
}

// NewCashFlowDataCashPosition instantiates a new CashFlowDataCashPosition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCashFlowDataCashPosition() *CashFlowDataCashPosition {
	this := CashFlowDataCashPosition{}
	return &this
}

// NewCashFlowDataCashPositionWithDefaults instantiates a new CashFlowDataCashPosition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCashFlowDataCashPositionWithDefaults() *CashFlowDataCashPosition {
	this := CashFlowDataCashPosition{}
	return &this
}

// GetBeginningCashPosition returns the BeginningCashPosition field value if set, zero value otherwise.
func (o *CashFlowDataCashPosition) GetBeginningCashPosition() float64 {
	if o == nil || IsNil(o.BeginningCashPosition) {
		var ret float64
		return ret
	}
	return *o.BeginningCashPosition
}

// GetBeginningCashPositionOk returns a tuple with the BeginningCashPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashPosition) GetBeginningCashPositionOk() (*float64, bool) {
	if o == nil || IsNil(o.BeginningCashPosition) {
		return nil, false
	}
	return o.BeginningCashPosition, true
}

// HasBeginningCashPosition returns a boolean if a field has been set.
func (o *CashFlowDataCashPosition) HasBeginningCashPosition() bool {
	if o != nil && !IsNil(o.BeginningCashPosition) {
		return true
	}

	return false
}

// SetBeginningCashPosition gets a reference to the given float64 and assigns it to the BeginningCashPosition field.
func (o *CashFlowDataCashPosition) SetBeginningCashPosition(v float64) {
	o.BeginningCashPosition = &v
}

// GetEndCashPosition returns the EndCashPosition field value if set, zero value otherwise.
func (o *CashFlowDataCashPosition) GetEndCashPosition() float64 {
	if o == nil || IsNil(o.EndCashPosition) {
		var ret float64
		return ret
	}
	return *o.EndCashPosition
}

// GetEndCashPositionOk returns a tuple with the EndCashPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashPosition) GetEndCashPositionOk() (*float64, bool) {
	if o == nil || IsNil(o.EndCashPosition) {
		return nil, false
	}
	return o.EndCashPosition, true
}

// HasEndCashPosition returns a boolean if a field has been set.
func (o *CashFlowDataCashPosition) HasEndCashPosition() bool {
	if o != nil && !IsNil(o.EndCashPosition) {
		return true
	}

	return false
}

// SetEndCashPosition gets a reference to the given float64 and assigns it to the EndCashPosition field.
func (o *CashFlowDataCashPosition) SetEndCashPosition(v float64) {
	o.EndCashPosition = &v
}

// GetChangesInCash returns the ChangesInCash field value if set, zero value otherwise.
func (o *CashFlowDataCashPosition) GetChangesInCash() float64 {
	if o == nil || IsNil(o.ChangesInCash) {
		var ret float64
		return ret
	}
	return *o.ChangesInCash
}

// GetChangesInCashOk returns a tuple with the ChangesInCash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashPosition) GetChangesInCashOk() (*float64, bool) {
	if o == nil || IsNil(o.ChangesInCash) {
		return nil, false
	}
	return o.ChangesInCash, true
}

// HasChangesInCash returns a boolean if a field has been set.
func (o *CashFlowDataCashPosition) HasChangesInCash() bool {
	if o != nil && !IsNil(o.ChangesInCash) {
		return true
	}

	return false
}

// SetChangesInCash gets a reference to the given float64 and assigns it to the ChangesInCash field.
func (o *CashFlowDataCashPosition) SetChangesInCash(v float64) {
	o.ChangesInCash = &v
}

// GetOtherCashAdjustmentOutsideChangeInCash returns the OtherCashAdjustmentOutsideChangeInCash field value if set, zero value otherwise.
func (o *CashFlowDataCashPosition) GetOtherCashAdjustmentOutsideChangeInCash() float64 {
	if o == nil || IsNil(o.OtherCashAdjustmentOutsideChangeInCash) {
		var ret float64
		return ret
	}
	return *o.OtherCashAdjustmentOutsideChangeInCash
}

// GetOtherCashAdjustmentOutsideChangeInCashOk returns a tuple with the OtherCashAdjustmentOutsideChangeInCash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashPosition) GetOtherCashAdjustmentOutsideChangeInCashOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherCashAdjustmentOutsideChangeInCash) {
		return nil, false
	}
	return o.OtherCashAdjustmentOutsideChangeInCash, true
}

// HasOtherCashAdjustmentOutsideChangeInCash returns a boolean if a field has been set.
func (o *CashFlowDataCashPosition) HasOtherCashAdjustmentOutsideChangeInCash() bool {
	if o != nil && !IsNil(o.OtherCashAdjustmentOutsideChangeInCash) {
		return true
	}

	return false
}

// SetOtherCashAdjustmentOutsideChangeInCash gets a reference to the given float64 and assigns it to the OtherCashAdjustmentOutsideChangeInCash field.
func (o *CashFlowDataCashPosition) SetOtherCashAdjustmentOutsideChangeInCash(v float64) {
	o.OtherCashAdjustmentOutsideChangeInCash = &v
}

// GetOtherCashAdjustmentInsideChangeInCash returns the OtherCashAdjustmentInsideChangeInCash field value if set, zero value otherwise.
func (o *CashFlowDataCashPosition) GetOtherCashAdjustmentInsideChangeInCash() float64 {
	if o == nil || IsNil(o.OtherCashAdjustmentInsideChangeInCash) {
		var ret float64
		return ret
	}
	return *o.OtherCashAdjustmentInsideChangeInCash
}

// GetOtherCashAdjustmentInsideChangeInCashOk returns a tuple with the OtherCashAdjustmentInsideChangeInCash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashPosition) GetOtherCashAdjustmentInsideChangeInCashOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherCashAdjustmentInsideChangeInCash) {
		return nil, false
	}
	return o.OtherCashAdjustmentInsideChangeInCash, true
}

// HasOtherCashAdjustmentInsideChangeInCash returns a boolean if a field has been set.
func (o *CashFlowDataCashPosition) HasOtherCashAdjustmentInsideChangeInCash() bool {
	if o != nil && !IsNil(o.OtherCashAdjustmentInsideChangeInCash) {
		return true
	}

	return false
}

// SetOtherCashAdjustmentInsideChangeInCash gets a reference to the given float64 and assigns it to the OtherCashAdjustmentInsideChangeInCash field.
func (o *CashFlowDataCashPosition) SetOtherCashAdjustmentInsideChangeInCash(v float64) {
	o.OtherCashAdjustmentInsideChangeInCash = &v
}

// GetEffectOfExchangeRateChanges returns the EffectOfExchangeRateChanges field value if set, zero value otherwise.
func (o *CashFlowDataCashPosition) GetEffectOfExchangeRateChanges() float64 {
	if o == nil || IsNil(o.EffectOfExchangeRateChanges) {
		var ret float64
		return ret
	}
	return *o.EffectOfExchangeRateChanges
}

// GetEffectOfExchangeRateChangesOk returns a tuple with the EffectOfExchangeRateChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataCashPosition) GetEffectOfExchangeRateChangesOk() (*float64, bool) {
	if o == nil || IsNil(o.EffectOfExchangeRateChanges) {
		return nil, false
	}
	return o.EffectOfExchangeRateChanges, true
}

// HasEffectOfExchangeRateChanges returns a boolean if a field has been set.
func (o *CashFlowDataCashPosition) HasEffectOfExchangeRateChanges() bool {
	if o != nil && !IsNil(o.EffectOfExchangeRateChanges) {
		return true
	}

	return false
}

// SetEffectOfExchangeRateChanges gets a reference to the given float64 and assigns it to the EffectOfExchangeRateChanges field.
func (o *CashFlowDataCashPosition) SetEffectOfExchangeRateChanges(v float64) {
	o.EffectOfExchangeRateChanges = &v
}

func (o CashFlowDataCashPosition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CashFlowDataCashPosition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BeginningCashPosition) {
		toSerialize["beginning_cash_position"] = o.BeginningCashPosition
	}
	if !IsNil(o.EndCashPosition) {
		toSerialize["end_cash_position"] = o.EndCashPosition
	}
	if !IsNil(o.ChangesInCash) {
		toSerialize["changes_in_cash"] = o.ChangesInCash
	}
	if !IsNil(o.OtherCashAdjustmentOutsideChangeInCash) {
		toSerialize["other_cash_adjustment_outside_change_in_cash"] = o.OtherCashAdjustmentOutsideChangeInCash
	}
	if !IsNil(o.OtherCashAdjustmentInsideChangeInCash) {
		toSerialize["other_cash_adjustment_inside_change_in_cash"] = o.OtherCashAdjustmentInsideChangeInCash
	}
	if !IsNil(o.EffectOfExchangeRateChanges) {
		toSerialize["effect_of_exchange_rate_changes"] = o.EffectOfExchangeRateChanges
	}
	return toSerialize, nil
}

type NullableCashFlowDataCashPosition struct {
	value *CashFlowDataCashPosition
	isSet bool
}

func (v NullableCashFlowDataCashPosition) Get() *CashFlowDataCashPosition {
	return v.value
}

func (v *NullableCashFlowDataCashPosition) Set(val *CashFlowDataCashPosition) {
	v.value = val
	v.isSet = true
}

func (v NullableCashFlowDataCashPosition) IsSet() bool {
	return v.isSet
}

func (v *NullableCashFlowDataCashPosition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCashFlowDataCashPosition(val *CashFlowDataCashPosition) *NullableCashFlowDataCashPosition {
	return &NullableCashFlowDataCashPosition{value: val, isSet: true}
}

func (v NullableCashFlowDataCashPosition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCashFlowDataCashPosition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
