// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the CashFlowStructFinancingActivities type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CashFlowStructFinancingActivities{}

// CashFlowStructFinancingActivities Financing activities section
type CashFlowStructFinancingActivities struct {
	// Refers to the issuance of any financial obligations that extend beyond a 12 months period
	LongTermDebtIssuance *float64 `json:"long_term_debt_issuance,omitempty"`
	// Refers to the payments of any financial obligations that extend beyond a 12 months period
	LongTermDebtPayments *float64 `json:"long_term_debt_payments,omitempty"`
	// Refers to the issuance of any financial obligations that are expected to be paid off within a year
	ShortTermDebtIssuance *float64 `json:"short_term_debt_issuance,omitempty"`
	// Represents a transaction whereby a company issues its own shares to the marketplace
	CommonStockIssuance *float64 `json:"common_stock_issuance,omitempty"`
	// Represents a transaction whereby a company buys back its own shares from the marketplace
	CommonStockRepurchase *float64 `json:"common_stock_repurchase,omitempty"`
	// Returns value of payment doled out by a company to its stockholders in the form of periodic distributions of cash
	CommonDividends *float64 `json:"common_dividends,omitempty"`
	// Represents other financing charges
	OtherFinancingCharges *float64 `json:"other_financing_charges,omitempty"`
	// Returns cash flow from financing activities (CFF), which shows the net flows of cash that are used to fund the company
	FinancingCashFlow *float64 `json:"financing_cash_flow,omitempty"`
}

// NewCashFlowStructFinancingActivities instantiates a new CashFlowStructFinancingActivities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCashFlowStructFinancingActivities() *CashFlowStructFinancingActivities {
	this := CashFlowStructFinancingActivities{}
	return &this
}

// NewCashFlowStructFinancingActivitiesWithDefaults instantiates a new CashFlowStructFinancingActivities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCashFlowStructFinancingActivitiesWithDefaults() *CashFlowStructFinancingActivities {
	this := CashFlowStructFinancingActivities{}
	return &this
}

// GetLongTermDebtIssuance returns the LongTermDebtIssuance field value if set, zero value otherwise.
func (o *CashFlowStructFinancingActivities) GetLongTermDebtIssuance() float64 {
	if o == nil || IsNil(o.LongTermDebtIssuance) {
		var ret float64
		return ret
	}
	return *o.LongTermDebtIssuance
}

// GetLongTermDebtIssuanceOk returns a tuple with the LongTermDebtIssuance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowStructFinancingActivities) GetLongTermDebtIssuanceOk() (*float64, bool) {
	if o == nil || IsNil(o.LongTermDebtIssuance) {
		return nil, false
	}
	return o.LongTermDebtIssuance, true
}

// HasLongTermDebtIssuance returns a boolean if a field has been set.
func (o *CashFlowStructFinancingActivities) HasLongTermDebtIssuance() bool {
	if o != nil && !IsNil(o.LongTermDebtIssuance) {
		return true
	}

	return false
}

// SetLongTermDebtIssuance gets a reference to the given float64 and assigns it to the LongTermDebtIssuance field.
func (o *CashFlowStructFinancingActivities) SetLongTermDebtIssuance(v float64) {
	o.LongTermDebtIssuance = &v
}

// GetLongTermDebtPayments returns the LongTermDebtPayments field value if set, zero value otherwise.
func (o *CashFlowStructFinancingActivities) GetLongTermDebtPayments() float64 {
	if o == nil || IsNil(o.LongTermDebtPayments) {
		var ret float64
		return ret
	}
	return *o.LongTermDebtPayments
}

// GetLongTermDebtPaymentsOk returns a tuple with the LongTermDebtPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowStructFinancingActivities) GetLongTermDebtPaymentsOk() (*float64, bool) {
	if o == nil || IsNil(o.LongTermDebtPayments) {
		return nil, false
	}
	return o.LongTermDebtPayments, true
}

// HasLongTermDebtPayments returns a boolean if a field has been set.
func (o *CashFlowStructFinancingActivities) HasLongTermDebtPayments() bool {
	if o != nil && !IsNil(o.LongTermDebtPayments) {
		return true
	}

	return false
}

// SetLongTermDebtPayments gets a reference to the given float64 and assigns it to the LongTermDebtPayments field.
func (o *CashFlowStructFinancingActivities) SetLongTermDebtPayments(v float64) {
	o.LongTermDebtPayments = &v
}

// GetShortTermDebtIssuance returns the ShortTermDebtIssuance field value if set, zero value otherwise.
func (o *CashFlowStructFinancingActivities) GetShortTermDebtIssuance() float64 {
	if o == nil || IsNil(o.ShortTermDebtIssuance) {
		var ret float64
		return ret
	}
	return *o.ShortTermDebtIssuance
}

// GetShortTermDebtIssuanceOk returns a tuple with the ShortTermDebtIssuance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowStructFinancingActivities) GetShortTermDebtIssuanceOk() (*float64, bool) {
	if o == nil || IsNil(o.ShortTermDebtIssuance) {
		return nil, false
	}
	return o.ShortTermDebtIssuance, true
}

// HasShortTermDebtIssuance returns a boolean if a field has been set.
func (o *CashFlowStructFinancingActivities) HasShortTermDebtIssuance() bool {
	if o != nil && !IsNil(o.ShortTermDebtIssuance) {
		return true
	}

	return false
}

// SetShortTermDebtIssuance gets a reference to the given float64 and assigns it to the ShortTermDebtIssuance field.
func (o *CashFlowStructFinancingActivities) SetShortTermDebtIssuance(v float64) {
	o.ShortTermDebtIssuance = &v
}

// GetCommonStockIssuance returns the CommonStockIssuance field value if set, zero value otherwise.
func (o *CashFlowStructFinancingActivities) GetCommonStockIssuance() float64 {
	if o == nil || IsNil(o.CommonStockIssuance) {
		var ret float64
		return ret
	}
	return *o.CommonStockIssuance
}

// GetCommonStockIssuanceOk returns a tuple with the CommonStockIssuance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowStructFinancingActivities) GetCommonStockIssuanceOk() (*float64, bool) {
	if o == nil || IsNil(o.CommonStockIssuance) {
		return nil, false
	}
	return o.CommonStockIssuance, true
}

// HasCommonStockIssuance returns a boolean if a field has been set.
func (o *CashFlowStructFinancingActivities) HasCommonStockIssuance() bool {
	if o != nil && !IsNil(o.CommonStockIssuance) {
		return true
	}

	return false
}

// SetCommonStockIssuance gets a reference to the given float64 and assigns it to the CommonStockIssuance field.
func (o *CashFlowStructFinancingActivities) SetCommonStockIssuance(v float64) {
	o.CommonStockIssuance = &v
}

// GetCommonStockRepurchase returns the CommonStockRepurchase field value if set, zero value otherwise.
func (o *CashFlowStructFinancingActivities) GetCommonStockRepurchase() float64 {
	if o == nil || IsNil(o.CommonStockRepurchase) {
		var ret float64
		return ret
	}
	return *o.CommonStockRepurchase
}

// GetCommonStockRepurchaseOk returns a tuple with the CommonStockRepurchase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowStructFinancingActivities) GetCommonStockRepurchaseOk() (*float64, bool) {
	if o == nil || IsNil(o.CommonStockRepurchase) {
		return nil, false
	}
	return o.CommonStockRepurchase, true
}

// HasCommonStockRepurchase returns a boolean if a field has been set.
func (o *CashFlowStructFinancingActivities) HasCommonStockRepurchase() bool {
	if o != nil && !IsNil(o.CommonStockRepurchase) {
		return true
	}

	return false
}

// SetCommonStockRepurchase gets a reference to the given float64 and assigns it to the CommonStockRepurchase field.
func (o *CashFlowStructFinancingActivities) SetCommonStockRepurchase(v float64) {
	o.CommonStockRepurchase = &v
}

// GetCommonDividends returns the CommonDividends field value if set, zero value otherwise.
func (o *CashFlowStructFinancingActivities) GetCommonDividends() float64 {
	if o == nil || IsNil(o.CommonDividends) {
		var ret float64
		return ret
	}
	return *o.CommonDividends
}

// GetCommonDividendsOk returns a tuple with the CommonDividends field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowStructFinancingActivities) GetCommonDividendsOk() (*float64, bool) {
	if o == nil || IsNil(o.CommonDividends) {
		return nil, false
	}
	return o.CommonDividends, true
}

// HasCommonDividends returns a boolean if a field has been set.
func (o *CashFlowStructFinancingActivities) HasCommonDividends() bool {
	if o != nil && !IsNil(o.CommonDividends) {
		return true
	}

	return false
}

// SetCommonDividends gets a reference to the given float64 and assigns it to the CommonDividends field.
func (o *CashFlowStructFinancingActivities) SetCommonDividends(v float64) {
	o.CommonDividends = &v
}

// GetOtherFinancingCharges returns the OtherFinancingCharges field value if set, zero value otherwise.
func (o *CashFlowStructFinancingActivities) GetOtherFinancingCharges() float64 {
	if o == nil || IsNil(o.OtherFinancingCharges) {
		var ret float64
		return ret
	}
	return *o.OtherFinancingCharges
}

// GetOtherFinancingChargesOk returns a tuple with the OtherFinancingCharges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowStructFinancingActivities) GetOtherFinancingChargesOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherFinancingCharges) {
		return nil, false
	}
	return o.OtherFinancingCharges, true
}

// HasOtherFinancingCharges returns a boolean if a field has been set.
func (o *CashFlowStructFinancingActivities) HasOtherFinancingCharges() bool {
	if o != nil && !IsNil(o.OtherFinancingCharges) {
		return true
	}

	return false
}

// SetOtherFinancingCharges gets a reference to the given float64 and assigns it to the OtherFinancingCharges field.
func (o *CashFlowStructFinancingActivities) SetOtherFinancingCharges(v float64) {
	o.OtherFinancingCharges = &v
}

// GetFinancingCashFlow returns the FinancingCashFlow field value if set, zero value otherwise.
func (o *CashFlowStructFinancingActivities) GetFinancingCashFlow() float64 {
	if o == nil || IsNil(o.FinancingCashFlow) {
		var ret float64
		return ret
	}
	return *o.FinancingCashFlow
}

// GetFinancingCashFlowOk returns a tuple with the FinancingCashFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowStructFinancingActivities) GetFinancingCashFlowOk() (*float64, bool) {
	if o == nil || IsNil(o.FinancingCashFlow) {
		return nil, false
	}
	return o.FinancingCashFlow, true
}

// HasFinancingCashFlow returns a boolean if a field has been set.
func (o *CashFlowStructFinancingActivities) HasFinancingCashFlow() bool {
	if o != nil && !IsNil(o.FinancingCashFlow) {
		return true
	}

	return false
}

// SetFinancingCashFlow gets a reference to the given float64 and assigns it to the FinancingCashFlow field.
func (o *CashFlowStructFinancingActivities) SetFinancingCashFlow(v float64) {
	o.FinancingCashFlow = &v
}

func (o CashFlowStructFinancingActivities) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CashFlowStructFinancingActivities) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LongTermDebtIssuance) {
		toSerialize["long_term_debt_issuance"] = o.LongTermDebtIssuance
	}
	if !IsNil(o.LongTermDebtPayments) {
		toSerialize["long_term_debt_payments"] = o.LongTermDebtPayments
	}
	if !IsNil(o.ShortTermDebtIssuance) {
		toSerialize["short_term_debt_issuance"] = o.ShortTermDebtIssuance
	}
	if !IsNil(o.CommonStockIssuance) {
		toSerialize["common_stock_issuance"] = o.CommonStockIssuance
	}
	if !IsNil(o.CommonStockRepurchase) {
		toSerialize["common_stock_repurchase"] = o.CommonStockRepurchase
	}
	if !IsNil(o.CommonDividends) {
		toSerialize["common_dividends"] = o.CommonDividends
	}
	if !IsNil(o.OtherFinancingCharges) {
		toSerialize["other_financing_charges"] = o.OtherFinancingCharges
	}
	if !IsNil(o.FinancingCashFlow) {
		toSerialize["financing_cash_flow"] = o.FinancingCashFlow
	}
	return toSerialize, nil
}

type NullableCashFlowStructFinancingActivities struct {
	value *CashFlowStructFinancingActivities
	isSet bool
}

func (v NullableCashFlowStructFinancingActivities) Get() *CashFlowStructFinancingActivities {
	return v.value
}

func (v *NullableCashFlowStructFinancingActivities) Set(val *CashFlowStructFinancingActivities) {
	v.value = val
	v.isSet = true
}

func (v NullableCashFlowStructFinancingActivities) IsSet() bool {
	return v.isSet
}

func (v *NullableCashFlowStructFinancingActivities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCashFlowStructFinancingActivities(val *CashFlowStructFinancingActivities) *NullableCashFlowStructFinancingActivities {
	return &NullableCashFlowStructFinancingActivities{value: val, isSet: true}
}

func (v NullableCashFlowStructFinancingActivities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCashFlowStructFinancingActivities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
