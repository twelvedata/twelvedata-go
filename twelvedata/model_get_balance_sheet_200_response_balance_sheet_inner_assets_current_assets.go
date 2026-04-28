/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets{}

// GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets Current assets section
type GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets struct {
	// Cash includes currency, bank accounts, and undeposited checks
	Cash *float64 `json:"cash,omitempty"`
	// Represents cash equivalents that have high credit quality and are highly liquid
	CashEquivalents *float64 `json:"cash_equivalents,omitempty"`
	// Combined value of cash and cash equivalents when company doesn't report a breakdown
	CashAndCashEquivalents *float64 `json:"cash_and_cash_equivalents,omitempty"`
	// Represents other short term investments
	OtherShortTermInvestments *float64 `json:"other_short_term_investments,omitempty"`
	// Represents the balance of money due to a firm for goods or services delivered or used but not yet paid for by customers
	AccountsReceivable *float64 `json:"accounts_receivable,omitempty"`
	// Represents other receivables
	OtherReceivables *float64 `json:"other_receivables,omitempty"`
	// Represents the goods available for sale and raw materials used to produce goods available for sale
	Inventory *float64 `json:"inventory,omitempty"`
	// Represents expense that has already been paid for, but which has not yet been consumed
	PrepaidAssets *float64 `json:"prepaid_assets,omitempty"`
	// Represents money that is held for a specific purpose and thus not available to the company for immediate or general business use
	RestrictedCash *float64 `json:"restricted_cash,omitempty"`
	// Represents assets which company plans to sell
	AssetsHeldForSale *float64 `json:"assets_held_for_sale,omitempty"`
	// Represents money that is spent on hedging assets
	HedgingAssets *float64 `json:"hedging_assets,omitempty"`
	// Represents other current assets
	OtherCurrentAssets *float64 `json:"other_current_assets,omitempty"`
	// All current assets values in a total
	TotalCurrentAssets *float64 `json:"total_current_assets,omitempty"`
}

// NewGetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets instantiates a new GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets() *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets {
	this := GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets{}
	return &this
}

// NewGetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssetsWithDefaults instantiates a new GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssetsWithDefaults() *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets {
	this := GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets{}
	return &this
}

// GetCash returns the Cash field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetCash() float64 {
	if o == nil || IsNil(o.Cash) {
		var ret float64
		return ret
	}
	return *o.Cash
}

// GetCashOk returns a tuple with the Cash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetCashOk() (*float64, bool) {
	if o == nil || IsNil(o.Cash) {
		return nil, false
	}
	return o.Cash, true
}

// HasCash returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) HasCash() bool {
	if o != nil && !IsNil(o.Cash) {
		return true
	}

	return false
}

// SetCash gets a reference to the given float64 and assigns it to the Cash field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) SetCash(v float64) {
	o.Cash = &v
}

// GetCashEquivalents returns the CashEquivalents field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetCashEquivalents() float64 {
	if o == nil || IsNil(o.CashEquivalents) {
		var ret float64
		return ret
	}
	return *o.CashEquivalents
}

// GetCashEquivalentsOk returns a tuple with the CashEquivalents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetCashEquivalentsOk() (*float64, bool) {
	if o == nil || IsNil(o.CashEquivalents) {
		return nil, false
	}
	return o.CashEquivalents, true
}

// HasCashEquivalents returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) HasCashEquivalents() bool {
	if o != nil && !IsNil(o.CashEquivalents) {
		return true
	}

	return false
}

// SetCashEquivalents gets a reference to the given float64 and assigns it to the CashEquivalents field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) SetCashEquivalents(v float64) {
	o.CashEquivalents = &v
}

// GetCashAndCashEquivalents returns the CashAndCashEquivalents field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetCashAndCashEquivalents() float64 {
	if o == nil || IsNil(o.CashAndCashEquivalents) {
		var ret float64
		return ret
	}
	return *o.CashAndCashEquivalents
}

// GetCashAndCashEquivalentsOk returns a tuple with the CashAndCashEquivalents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetCashAndCashEquivalentsOk() (*float64, bool) {
	if o == nil || IsNil(o.CashAndCashEquivalents) {
		return nil, false
	}
	return o.CashAndCashEquivalents, true
}

// HasCashAndCashEquivalents returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) HasCashAndCashEquivalents() bool {
	if o != nil && !IsNil(o.CashAndCashEquivalents) {
		return true
	}

	return false
}

// SetCashAndCashEquivalents gets a reference to the given float64 and assigns it to the CashAndCashEquivalents field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) SetCashAndCashEquivalents(v float64) {
	o.CashAndCashEquivalents = &v
}

// GetOtherShortTermInvestments returns the OtherShortTermInvestments field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetOtherShortTermInvestments() float64 {
	if o == nil || IsNil(o.OtherShortTermInvestments) {
		var ret float64
		return ret
	}
	return *o.OtherShortTermInvestments
}

// GetOtherShortTermInvestmentsOk returns a tuple with the OtherShortTermInvestments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetOtherShortTermInvestmentsOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherShortTermInvestments) {
		return nil, false
	}
	return o.OtherShortTermInvestments, true
}

// HasOtherShortTermInvestments returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) HasOtherShortTermInvestments() bool {
	if o != nil && !IsNil(o.OtherShortTermInvestments) {
		return true
	}

	return false
}

// SetOtherShortTermInvestments gets a reference to the given float64 and assigns it to the OtherShortTermInvestments field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) SetOtherShortTermInvestments(v float64) {
	o.OtherShortTermInvestments = &v
}

// GetAccountsReceivable returns the AccountsReceivable field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetAccountsReceivable() float64 {
	if o == nil || IsNil(o.AccountsReceivable) {
		var ret float64
		return ret
	}
	return *o.AccountsReceivable
}

// GetAccountsReceivableOk returns a tuple with the AccountsReceivable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetAccountsReceivableOk() (*float64, bool) {
	if o == nil || IsNil(o.AccountsReceivable) {
		return nil, false
	}
	return o.AccountsReceivable, true
}

// HasAccountsReceivable returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) HasAccountsReceivable() bool {
	if o != nil && !IsNil(o.AccountsReceivable) {
		return true
	}

	return false
}

// SetAccountsReceivable gets a reference to the given float64 and assigns it to the AccountsReceivable field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) SetAccountsReceivable(v float64) {
	o.AccountsReceivable = &v
}

// GetOtherReceivables returns the OtherReceivables field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetOtherReceivables() float64 {
	if o == nil || IsNil(o.OtherReceivables) {
		var ret float64
		return ret
	}
	return *o.OtherReceivables
}

// GetOtherReceivablesOk returns a tuple with the OtherReceivables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetOtherReceivablesOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherReceivables) {
		return nil, false
	}
	return o.OtherReceivables, true
}

// HasOtherReceivables returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) HasOtherReceivables() bool {
	if o != nil && !IsNil(o.OtherReceivables) {
		return true
	}

	return false
}

// SetOtherReceivables gets a reference to the given float64 and assigns it to the OtherReceivables field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) SetOtherReceivables(v float64) {
	o.OtherReceivables = &v
}

// GetInventory returns the Inventory field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetInventory() float64 {
	if o == nil || IsNil(o.Inventory) {
		var ret float64
		return ret
	}
	return *o.Inventory
}

// GetInventoryOk returns a tuple with the Inventory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetInventoryOk() (*float64, bool) {
	if o == nil || IsNil(o.Inventory) {
		return nil, false
	}
	return o.Inventory, true
}

// HasInventory returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) HasInventory() bool {
	if o != nil && !IsNil(o.Inventory) {
		return true
	}

	return false
}

// SetInventory gets a reference to the given float64 and assigns it to the Inventory field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) SetInventory(v float64) {
	o.Inventory = &v
}

// GetPrepaidAssets returns the PrepaidAssets field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetPrepaidAssets() float64 {
	if o == nil || IsNil(o.PrepaidAssets) {
		var ret float64
		return ret
	}
	return *o.PrepaidAssets
}

// GetPrepaidAssetsOk returns a tuple with the PrepaidAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetPrepaidAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.PrepaidAssets) {
		return nil, false
	}
	return o.PrepaidAssets, true
}

// HasPrepaidAssets returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) HasPrepaidAssets() bool {
	if o != nil && !IsNil(o.PrepaidAssets) {
		return true
	}

	return false
}

// SetPrepaidAssets gets a reference to the given float64 and assigns it to the PrepaidAssets field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) SetPrepaidAssets(v float64) {
	o.PrepaidAssets = &v
}

// GetRestrictedCash returns the RestrictedCash field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetRestrictedCash() float64 {
	if o == nil || IsNil(o.RestrictedCash) {
		var ret float64
		return ret
	}
	return *o.RestrictedCash
}

// GetRestrictedCashOk returns a tuple with the RestrictedCash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetRestrictedCashOk() (*float64, bool) {
	if o == nil || IsNil(o.RestrictedCash) {
		return nil, false
	}
	return o.RestrictedCash, true
}

// HasRestrictedCash returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) HasRestrictedCash() bool {
	if o != nil && !IsNil(o.RestrictedCash) {
		return true
	}

	return false
}

// SetRestrictedCash gets a reference to the given float64 and assigns it to the RestrictedCash field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) SetRestrictedCash(v float64) {
	o.RestrictedCash = &v
}

// GetAssetsHeldForSale returns the AssetsHeldForSale field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetAssetsHeldForSale() float64 {
	if o == nil || IsNil(o.AssetsHeldForSale) {
		var ret float64
		return ret
	}
	return *o.AssetsHeldForSale
}

// GetAssetsHeldForSaleOk returns a tuple with the AssetsHeldForSale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetAssetsHeldForSaleOk() (*float64, bool) {
	if o == nil || IsNil(o.AssetsHeldForSale) {
		return nil, false
	}
	return o.AssetsHeldForSale, true
}

// HasAssetsHeldForSale returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) HasAssetsHeldForSale() bool {
	if o != nil && !IsNil(o.AssetsHeldForSale) {
		return true
	}

	return false
}

// SetAssetsHeldForSale gets a reference to the given float64 and assigns it to the AssetsHeldForSale field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) SetAssetsHeldForSale(v float64) {
	o.AssetsHeldForSale = &v
}

// GetHedgingAssets returns the HedgingAssets field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetHedgingAssets() float64 {
	if o == nil || IsNil(o.HedgingAssets) {
		var ret float64
		return ret
	}
	return *o.HedgingAssets
}

// GetHedgingAssetsOk returns a tuple with the HedgingAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetHedgingAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.HedgingAssets) {
		return nil, false
	}
	return o.HedgingAssets, true
}

// HasHedgingAssets returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) HasHedgingAssets() bool {
	if o != nil && !IsNil(o.HedgingAssets) {
		return true
	}

	return false
}

// SetHedgingAssets gets a reference to the given float64 and assigns it to the HedgingAssets field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) SetHedgingAssets(v float64) {
	o.HedgingAssets = &v
}

// GetOtherCurrentAssets returns the OtherCurrentAssets field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetOtherCurrentAssets() float64 {
	if o == nil || IsNil(o.OtherCurrentAssets) {
		var ret float64
		return ret
	}
	return *o.OtherCurrentAssets
}

// GetOtherCurrentAssetsOk returns a tuple with the OtherCurrentAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetOtherCurrentAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherCurrentAssets) {
		return nil, false
	}
	return o.OtherCurrentAssets, true
}

// HasOtherCurrentAssets returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) HasOtherCurrentAssets() bool {
	if o != nil && !IsNil(o.OtherCurrentAssets) {
		return true
	}

	return false
}

// SetOtherCurrentAssets gets a reference to the given float64 and assigns it to the OtherCurrentAssets field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) SetOtherCurrentAssets(v float64) {
	o.OtherCurrentAssets = &v
}

// GetTotalCurrentAssets returns the TotalCurrentAssets field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetTotalCurrentAssets() float64 {
	if o == nil || IsNil(o.TotalCurrentAssets) {
		var ret float64
		return ret
	}
	return *o.TotalCurrentAssets
}

// GetTotalCurrentAssetsOk returns a tuple with the TotalCurrentAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) GetTotalCurrentAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalCurrentAssets) {
		return nil, false
	}
	return o.TotalCurrentAssets, true
}

// HasTotalCurrentAssets returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) HasTotalCurrentAssets() bool {
	if o != nil && !IsNil(o.TotalCurrentAssets) {
		return true
	}

	return false
}

// SetTotalCurrentAssets gets a reference to the given float64 and assigns it to the TotalCurrentAssets field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) SetTotalCurrentAssets(v float64) {
	o.TotalCurrentAssets = &v
}

func (o GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cash) {
		toSerialize["cash"] = o.Cash
	}
	if !IsNil(o.CashEquivalents) {
		toSerialize["cash_equivalents"] = o.CashEquivalents
	}
	if !IsNil(o.CashAndCashEquivalents) {
		toSerialize["cash_and_cash_equivalents"] = o.CashAndCashEquivalents
	}
	if !IsNil(o.OtherShortTermInvestments) {
		toSerialize["other_short_term_investments"] = o.OtherShortTermInvestments
	}
	if !IsNil(o.AccountsReceivable) {
		toSerialize["accounts_receivable"] = o.AccountsReceivable
	}
	if !IsNil(o.OtherReceivables) {
		toSerialize["other_receivables"] = o.OtherReceivables
	}
	if !IsNil(o.Inventory) {
		toSerialize["inventory"] = o.Inventory
	}
	if !IsNil(o.PrepaidAssets) {
		toSerialize["prepaid_assets"] = o.PrepaidAssets
	}
	if !IsNil(o.RestrictedCash) {
		toSerialize["restricted_cash"] = o.RestrictedCash
	}
	if !IsNil(o.AssetsHeldForSale) {
		toSerialize["assets_held_for_sale"] = o.AssetsHeldForSale
	}
	if !IsNil(o.HedgingAssets) {
		toSerialize["hedging_assets"] = o.HedgingAssets
	}
	if !IsNil(o.OtherCurrentAssets) {
		toSerialize["other_current_assets"] = o.OtherCurrentAssets
	}
	if !IsNil(o.TotalCurrentAssets) {
		toSerialize["total_current_assets"] = o.TotalCurrentAssets
	}
	return toSerialize, nil
}

type NullableGetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets struct {
	value *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets
	isSet bool
}

func (v NullableGetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) Get() *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets {
	return v.value
}

func (v *NullableGetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) Set(val *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets(val *GetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) *NullableGetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets {
	return &NullableGetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets{value: val, isSet: true}
}

func (v NullableGetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBalanceSheet200ResponseBalanceSheetInnerAssetsCurrentAssets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
