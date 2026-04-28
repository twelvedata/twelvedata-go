/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the AssetsInfoCurrentAssets type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetsInfoCurrentAssets{}

// AssetsInfoCurrentAssets Current assets information
type AssetsInfoCurrentAssets struct {
	// Total current assets
	TotalCurrentAssets *float64 `json:"total_current_assets,omitempty"`
	// Cash cash equivalents and short term investments
	CashCashEquivalentsAndShortTermInvestments *float64 `json:"cash_cash_equivalents_and_short_term_investments,omitempty"`
	// Cash and cash equivalents
	CashAndCashEquivalents *float64 `json:"cash_and_cash_equivalents,omitempty"`
	// Cash equivalents
	CashEquivalents *float64 `json:"cash_equivalents,omitempty"`
	// Cash financial
	CashFinancial *float64 `json:"cash_financial,omitempty"`
	// Other short term investments
	OtherShortTermInvestments *float64 `json:"other_short_term_investments,omitempty"`
	// Restricted cash
	RestrictedCash *float64                            `json:"restricted_cash,omitempty"`
	Receivables    *AssetsInfoCurrentAssetsReceivables `json:"receivables,omitempty"`
	Inventory      *AssetsInfoCurrentAssetsInventory   `json:"inventory,omitempty"`
	// Prepaid assets
	PrepaidAssets *float64 `json:"prepaid_assets,omitempty"`
	// Current deferred assets
	CurrentDeferredAssets *float64 `json:"current_deferred_assets,omitempty"`
	// Current deferred taxes assets
	CurrentDeferredTaxesAssets *float64 `json:"current_deferred_taxes_assets,omitempty"`
	// Assets held for sale current
	AssetsHeldForSaleCurrent *float64 `json:"assets_held_for_sale_current,omitempty"`
	// Hedging assets current
	HedgingAssetsCurrent *float64 `json:"hedging_assets_current,omitempty"`
	// Other current assets
	OtherCurrentAssets *float64 `json:"other_current_assets,omitempty"`
}

// NewAssetsInfoCurrentAssets instantiates a new AssetsInfoCurrentAssets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetsInfoCurrentAssets() *AssetsInfoCurrentAssets {
	this := AssetsInfoCurrentAssets{}
	return &this
}

// NewAssetsInfoCurrentAssetsWithDefaults instantiates a new AssetsInfoCurrentAssets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetsInfoCurrentAssetsWithDefaults() *AssetsInfoCurrentAssets {
	this := AssetsInfoCurrentAssets{}
	return &this
}

// GetTotalCurrentAssets returns the TotalCurrentAssets field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssets) GetTotalCurrentAssets() float64 {
	if o == nil || IsNil(o.TotalCurrentAssets) {
		var ret float64
		return ret
	}
	return *o.TotalCurrentAssets
}

// GetTotalCurrentAssetsOk returns a tuple with the TotalCurrentAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssets) GetTotalCurrentAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalCurrentAssets) {
		return nil, false
	}
	return o.TotalCurrentAssets, true
}

// HasTotalCurrentAssets returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssets) HasTotalCurrentAssets() bool {
	if o != nil && !IsNil(o.TotalCurrentAssets) {
		return true
	}

	return false
}

// SetTotalCurrentAssets gets a reference to the given float64 and assigns it to the TotalCurrentAssets field.
func (o *AssetsInfoCurrentAssets) SetTotalCurrentAssets(v float64) {
	o.TotalCurrentAssets = &v
}

// GetCashCashEquivalentsAndShortTermInvestments returns the CashCashEquivalentsAndShortTermInvestments field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssets) GetCashCashEquivalentsAndShortTermInvestments() float64 {
	if o == nil || IsNil(o.CashCashEquivalentsAndShortTermInvestments) {
		var ret float64
		return ret
	}
	return *o.CashCashEquivalentsAndShortTermInvestments
}

// GetCashCashEquivalentsAndShortTermInvestmentsOk returns a tuple with the CashCashEquivalentsAndShortTermInvestments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssets) GetCashCashEquivalentsAndShortTermInvestmentsOk() (*float64, bool) {
	if o == nil || IsNil(o.CashCashEquivalentsAndShortTermInvestments) {
		return nil, false
	}
	return o.CashCashEquivalentsAndShortTermInvestments, true
}

// HasCashCashEquivalentsAndShortTermInvestments returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssets) HasCashCashEquivalentsAndShortTermInvestments() bool {
	if o != nil && !IsNil(o.CashCashEquivalentsAndShortTermInvestments) {
		return true
	}

	return false
}

// SetCashCashEquivalentsAndShortTermInvestments gets a reference to the given float64 and assigns it to the CashCashEquivalentsAndShortTermInvestments field.
func (o *AssetsInfoCurrentAssets) SetCashCashEquivalentsAndShortTermInvestments(v float64) {
	o.CashCashEquivalentsAndShortTermInvestments = &v
}

// GetCashAndCashEquivalents returns the CashAndCashEquivalents field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssets) GetCashAndCashEquivalents() float64 {
	if o == nil || IsNil(o.CashAndCashEquivalents) {
		var ret float64
		return ret
	}
	return *o.CashAndCashEquivalents
}

// GetCashAndCashEquivalentsOk returns a tuple with the CashAndCashEquivalents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssets) GetCashAndCashEquivalentsOk() (*float64, bool) {
	if o == nil || IsNil(o.CashAndCashEquivalents) {
		return nil, false
	}
	return o.CashAndCashEquivalents, true
}

// HasCashAndCashEquivalents returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssets) HasCashAndCashEquivalents() bool {
	if o != nil && !IsNil(o.CashAndCashEquivalents) {
		return true
	}

	return false
}

// SetCashAndCashEquivalents gets a reference to the given float64 and assigns it to the CashAndCashEquivalents field.
func (o *AssetsInfoCurrentAssets) SetCashAndCashEquivalents(v float64) {
	o.CashAndCashEquivalents = &v
}

// GetCashEquivalents returns the CashEquivalents field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssets) GetCashEquivalents() float64 {
	if o == nil || IsNil(o.CashEquivalents) {
		var ret float64
		return ret
	}
	return *o.CashEquivalents
}

// GetCashEquivalentsOk returns a tuple with the CashEquivalents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssets) GetCashEquivalentsOk() (*float64, bool) {
	if o == nil || IsNil(o.CashEquivalents) {
		return nil, false
	}
	return o.CashEquivalents, true
}

// HasCashEquivalents returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssets) HasCashEquivalents() bool {
	if o != nil && !IsNil(o.CashEquivalents) {
		return true
	}

	return false
}

// SetCashEquivalents gets a reference to the given float64 and assigns it to the CashEquivalents field.
func (o *AssetsInfoCurrentAssets) SetCashEquivalents(v float64) {
	o.CashEquivalents = &v
}

// GetCashFinancial returns the CashFinancial field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssets) GetCashFinancial() float64 {
	if o == nil || IsNil(o.CashFinancial) {
		var ret float64
		return ret
	}
	return *o.CashFinancial
}

// GetCashFinancialOk returns a tuple with the CashFinancial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssets) GetCashFinancialOk() (*float64, bool) {
	if o == nil || IsNil(o.CashFinancial) {
		return nil, false
	}
	return o.CashFinancial, true
}

// HasCashFinancial returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssets) HasCashFinancial() bool {
	if o != nil && !IsNil(o.CashFinancial) {
		return true
	}

	return false
}

// SetCashFinancial gets a reference to the given float64 and assigns it to the CashFinancial field.
func (o *AssetsInfoCurrentAssets) SetCashFinancial(v float64) {
	o.CashFinancial = &v
}

// GetOtherShortTermInvestments returns the OtherShortTermInvestments field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssets) GetOtherShortTermInvestments() float64 {
	if o == nil || IsNil(o.OtherShortTermInvestments) {
		var ret float64
		return ret
	}
	return *o.OtherShortTermInvestments
}

// GetOtherShortTermInvestmentsOk returns a tuple with the OtherShortTermInvestments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssets) GetOtherShortTermInvestmentsOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherShortTermInvestments) {
		return nil, false
	}
	return o.OtherShortTermInvestments, true
}

// HasOtherShortTermInvestments returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssets) HasOtherShortTermInvestments() bool {
	if o != nil && !IsNil(o.OtherShortTermInvestments) {
		return true
	}

	return false
}

// SetOtherShortTermInvestments gets a reference to the given float64 and assigns it to the OtherShortTermInvestments field.
func (o *AssetsInfoCurrentAssets) SetOtherShortTermInvestments(v float64) {
	o.OtherShortTermInvestments = &v
}

// GetRestrictedCash returns the RestrictedCash field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssets) GetRestrictedCash() float64 {
	if o == nil || IsNil(o.RestrictedCash) {
		var ret float64
		return ret
	}
	return *o.RestrictedCash
}

// GetRestrictedCashOk returns a tuple with the RestrictedCash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssets) GetRestrictedCashOk() (*float64, bool) {
	if o == nil || IsNil(o.RestrictedCash) {
		return nil, false
	}
	return o.RestrictedCash, true
}

// HasRestrictedCash returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssets) HasRestrictedCash() bool {
	if o != nil && !IsNil(o.RestrictedCash) {
		return true
	}

	return false
}

// SetRestrictedCash gets a reference to the given float64 and assigns it to the RestrictedCash field.
func (o *AssetsInfoCurrentAssets) SetRestrictedCash(v float64) {
	o.RestrictedCash = &v
}

// GetReceivables returns the Receivables field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssets) GetReceivables() AssetsInfoCurrentAssetsReceivables {
	if o == nil || IsNil(o.Receivables) {
		var ret AssetsInfoCurrentAssetsReceivables
		return ret
	}
	return *o.Receivables
}

// GetReceivablesOk returns a tuple with the Receivables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssets) GetReceivablesOk() (*AssetsInfoCurrentAssetsReceivables, bool) {
	if o == nil || IsNil(o.Receivables) {
		return nil, false
	}
	return o.Receivables, true
}

// HasReceivables returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssets) HasReceivables() bool {
	if o != nil && !IsNil(o.Receivables) {
		return true
	}

	return false
}

// SetReceivables gets a reference to the given AssetsInfoCurrentAssetsReceivables and assigns it to the Receivables field.
func (o *AssetsInfoCurrentAssets) SetReceivables(v AssetsInfoCurrentAssetsReceivables) {
	o.Receivables = &v
}

// GetInventory returns the Inventory field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssets) GetInventory() AssetsInfoCurrentAssetsInventory {
	if o == nil || IsNil(o.Inventory) {
		var ret AssetsInfoCurrentAssetsInventory
		return ret
	}
	return *o.Inventory
}

// GetInventoryOk returns a tuple with the Inventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssets) GetInventoryOk() (*AssetsInfoCurrentAssetsInventory, bool) {
	if o == nil || IsNil(o.Inventory) {
		return nil, false
	}
	return o.Inventory, true
}

// HasInventory returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssets) HasInventory() bool {
	if o != nil && !IsNil(o.Inventory) {
		return true
	}

	return false
}

// SetInventory gets a reference to the given AssetsInfoCurrentAssetsInventory and assigns it to the Inventory field.
func (o *AssetsInfoCurrentAssets) SetInventory(v AssetsInfoCurrentAssetsInventory) {
	o.Inventory = &v
}

// GetPrepaidAssets returns the PrepaidAssets field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssets) GetPrepaidAssets() float64 {
	if o == nil || IsNil(o.PrepaidAssets) {
		var ret float64
		return ret
	}
	return *o.PrepaidAssets
}

// GetPrepaidAssetsOk returns a tuple with the PrepaidAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssets) GetPrepaidAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.PrepaidAssets) {
		return nil, false
	}
	return o.PrepaidAssets, true
}

// HasPrepaidAssets returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssets) HasPrepaidAssets() bool {
	if o != nil && !IsNil(o.PrepaidAssets) {
		return true
	}

	return false
}

// SetPrepaidAssets gets a reference to the given float64 and assigns it to the PrepaidAssets field.
func (o *AssetsInfoCurrentAssets) SetPrepaidAssets(v float64) {
	o.PrepaidAssets = &v
}

// GetCurrentDeferredAssets returns the CurrentDeferredAssets field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssets) GetCurrentDeferredAssets() float64 {
	if o == nil || IsNil(o.CurrentDeferredAssets) {
		var ret float64
		return ret
	}
	return *o.CurrentDeferredAssets
}

// GetCurrentDeferredAssetsOk returns a tuple with the CurrentDeferredAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssets) GetCurrentDeferredAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.CurrentDeferredAssets) {
		return nil, false
	}
	return o.CurrentDeferredAssets, true
}

// HasCurrentDeferredAssets returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssets) HasCurrentDeferredAssets() bool {
	if o != nil && !IsNil(o.CurrentDeferredAssets) {
		return true
	}

	return false
}

// SetCurrentDeferredAssets gets a reference to the given float64 and assigns it to the CurrentDeferredAssets field.
func (o *AssetsInfoCurrentAssets) SetCurrentDeferredAssets(v float64) {
	o.CurrentDeferredAssets = &v
}

// GetCurrentDeferredTaxesAssets returns the CurrentDeferredTaxesAssets field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssets) GetCurrentDeferredTaxesAssets() float64 {
	if o == nil || IsNil(o.CurrentDeferredTaxesAssets) {
		var ret float64
		return ret
	}
	return *o.CurrentDeferredTaxesAssets
}

// GetCurrentDeferredTaxesAssetsOk returns a tuple with the CurrentDeferredTaxesAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssets) GetCurrentDeferredTaxesAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.CurrentDeferredTaxesAssets) {
		return nil, false
	}
	return o.CurrentDeferredTaxesAssets, true
}

// HasCurrentDeferredTaxesAssets returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssets) HasCurrentDeferredTaxesAssets() bool {
	if o != nil && !IsNil(o.CurrentDeferredTaxesAssets) {
		return true
	}

	return false
}

// SetCurrentDeferredTaxesAssets gets a reference to the given float64 and assigns it to the CurrentDeferredTaxesAssets field.
func (o *AssetsInfoCurrentAssets) SetCurrentDeferredTaxesAssets(v float64) {
	o.CurrentDeferredTaxesAssets = &v
}

// GetAssetsHeldForSaleCurrent returns the AssetsHeldForSaleCurrent field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssets) GetAssetsHeldForSaleCurrent() float64 {
	if o == nil || IsNil(o.AssetsHeldForSaleCurrent) {
		var ret float64
		return ret
	}
	return *o.AssetsHeldForSaleCurrent
}

// GetAssetsHeldForSaleCurrentOk returns a tuple with the AssetsHeldForSaleCurrent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssets) GetAssetsHeldForSaleCurrentOk() (*float64, bool) {
	if o == nil || IsNil(o.AssetsHeldForSaleCurrent) {
		return nil, false
	}
	return o.AssetsHeldForSaleCurrent, true
}

// HasAssetsHeldForSaleCurrent returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssets) HasAssetsHeldForSaleCurrent() bool {
	if o != nil && !IsNil(o.AssetsHeldForSaleCurrent) {
		return true
	}

	return false
}

// SetAssetsHeldForSaleCurrent gets a reference to the given float64 and assigns it to the AssetsHeldForSaleCurrent field.
func (o *AssetsInfoCurrentAssets) SetAssetsHeldForSaleCurrent(v float64) {
	o.AssetsHeldForSaleCurrent = &v
}

// GetHedgingAssetsCurrent returns the HedgingAssetsCurrent field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssets) GetHedgingAssetsCurrent() float64 {
	if o == nil || IsNil(o.HedgingAssetsCurrent) {
		var ret float64
		return ret
	}
	return *o.HedgingAssetsCurrent
}

// GetHedgingAssetsCurrentOk returns a tuple with the HedgingAssetsCurrent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssets) GetHedgingAssetsCurrentOk() (*float64, bool) {
	if o == nil || IsNil(o.HedgingAssetsCurrent) {
		return nil, false
	}
	return o.HedgingAssetsCurrent, true
}

// HasHedgingAssetsCurrent returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssets) HasHedgingAssetsCurrent() bool {
	if o != nil && !IsNil(o.HedgingAssetsCurrent) {
		return true
	}

	return false
}

// SetHedgingAssetsCurrent gets a reference to the given float64 and assigns it to the HedgingAssetsCurrent field.
func (o *AssetsInfoCurrentAssets) SetHedgingAssetsCurrent(v float64) {
	o.HedgingAssetsCurrent = &v
}

// GetOtherCurrentAssets returns the OtherCurrentAssets field value if set, zero value otherwise.
func (o *AssetsInfoCurrentAssets) GetOtherCurrentAssets() float64 {
	if o == nil || IsNil(o.OtherCurrentAssets) {
		var ret float64
		return ret
	}
	return *o.OtherCurrentAssets
}

// GetOtherCurrentAssetsOk returns a tuple with the OtherCurrentAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoCurrentAssets) GetOtherCurrentAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherCurrentAssets) {
		return nil, false
	}
	return o.OtherCurrentAssets, true
}

// HasOtherCurrentAssets returns a boolean if a field has been set.
func (o *AssetsInfoCurrentAssets) HasOtherCurrentAssets() bool {
	if o != nil && !IsNil(o.OtherCurrentAssets) {
		return true
	}

	return false
}

// SetOtherCurrentAssets gets a reference to the given float64 and assigns it to the OtherCurrentAssets field.
func (o *AssetsInfoCurrentAssets) SetOtherCurrentAssets(v float64) {
	o.OtherCurrentAssets = &v
}

func (o AssetsInfoCurrentAssets) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetsInfoCurrentAssets) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalCurrentAssets) {
		toSerialize["total_current_assets"] = o.TotalCurrentAssets
	}
	if !IsNil(o.CashCashEquivalentsAndShortTermInvestments) {
		toSerialize["cash_cash_equivalents_and_short_term_investments"] = o.CashCashEquivalentsAndShortTermInvestments
	}
	if !IsNil(o.CashAndCashEquivalents) {
		toSerialize["cash_and_cash_equivalents"] = o.CashAndCashEquivalents
	}
	if !IsNil(o.CashEquivalents) {
		toSerialize["cash_equivalents"] = o.CashEquivalents
	}
	if !IsNil(o.CashFinancial) {
		toSerialize["cash_financial"] = o.CashFinancial
	}
	if !IsNil(o.OtherShortTermInvestments) {
		toSerialize["other_short_term_investments"] = o.OtherShortTermInvestments
	}
	if !IsNil(o.RestrictedCash) {
		toSerialize["restricted_cash"] = o.RestrictedCash
	}
	if !IsNil(o.Receivables) {
		toSerialize["receivables"] = o.Receivables
	}
	if !IsNil(o.Inventory) {
		toSerialize["inventory"] = o.Inventory
	}
	if !IsNil(o.PrepaidAssets) {
		toSerialize["prepaid_assets"] = o.PrepaidAssets
	}
	if !IsNil(o.CurrentDeferredAssets) {
		toSerialize["current_deferred_assets"] = o.CurrentDeferredAssets
	}
	if !IsNil(o.CurrentDeferredTaxesAssets) {
		toSerialize["current_deferred_taxes_assets"] = o.CurrentDeferredTaxesAssets
	}
	if !IsNil(o.AssetsHeldForSaleCurrent) {
		toSerialize["assets_held_for_sale_current"] = o.AssetsHeldForSaleCurrent
	}
	if !IsNil(o.HedgingAssetsCurrent) {
		toSerialize["hedging_assets_current"] = o.HedgingAssetsCurrent
	}
	if !IsNil(o.OtherCurrentAssets) {
		toSerialize["other_current_assets"] = o.OtherCurrentAssets
	}
	return toSerialize, nil
}

type NullableAssetsInfoCurrentAssets struct {
	value *AssetsInfoCurrentAssets
	isSet bool
}

func (v NullableAssetsInfoCurrentAssets) Get() *AssetsInfoCurrentAssets {
	return v.value
}

func (v *NullableAssetsInfoCurrentAssets) Set(val *AssetsInfoCurrentAssets) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetsInfoCurrentAssets) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetsInfoCurrentAssets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetsInfoCurrentAssets(val *AssetsInfoCurrentAssets) *NullableAssetsInfoCurrentAssets {
	return &NullableAssetsInfoCurrentAssets{value: val, isSet: true}
}

func (v NullableAssetsInfoCurrentAssets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetsInfoCurrentAssets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
