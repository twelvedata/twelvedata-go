/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the EquityInfoCapitalStock type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EquityInfoCapitalStock{}

// EquityInfoCapitalStock Capital stock information
type EquityInfoCapitalStock struct {
	// Common stock
	CommonStock *float64 `json:"common_stock,omitempty"`
	// Preferred stock
	PreferredStock *float64 `json:"preferred_stock,omitempty"`
	// Total partnership capital
	TotalPartnershipCapital *float64 `json:"total_partnership_capital,omitempty"`
	// General partnership capital
	GeneralPartnershipCapital *float64 `json:"general_partnership_capital,omitempty"`
	// Limited partnership capital
	LimitedPartnershipCapital *float64 `json:"limited_partnership_capital,omitempty"`
	// Capital stock
	CapitalStock *float64 `json:"capital_stock,omitempty"`
	// Other capital stock
	OtherCapitalStock *float64 `json:"other_capital_stock,omitempty"`
	// Additional paid in capital
	AdditionalPaidInCapital *float64 `json:"additional_paid_in_capital,omitempty"`
	// Retained earnings
	RetainedEarnings *float64 `json:"retained_earnings,omitempty"`
	// Treasury stock
	TreasuryStock *float64 `json:"treasury_stock,omitempty"`
	// Treasury shares number
	TreasurySharesNumber *float64 `json:"treasury_shares_number,omitempty"`
	// Ordinary shares number
	OrdinarySharesNumber *float64 `json:"ordinary_shares_number,omitempty"`
	// Preferred shares number
	PreferredSharesNumber *float64 `json:"preferred_shares_number,omitempty"`
	// Share issued
	ShareIssued *float64 `json:"share_issued,omitempty"`
}

// NewEquityInfoCapitalStock instantiates a new EquityInfoCapitalStock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquityInfoCapitalStock() *EquityInfoCapitalStock {
	this := EquityInfoCapitalStock{}
	return &this
}

// NewEquityInfoCapitalStockWithDefaults instantiates a new EquityInfoCapitalStock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquityInfoCapitalStockWithDefaults() *EquityInfoCapitalStock {
	this := EquityInfoCapitalStock{}
	return &this
}

// GetCommonStock returns the CommonStock field value if set, zero value otherwise.
func (o *EquityInfoCapitalStock) GetCommonStock() float64 {
	if o == nil || IsNil(o.CommonStock) {
		var ret float64
		return ret
	}
	return *o.CommonStock
}

// GetCommonStockOk returns a tuple with the CommonStock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfoCapitalStock) GetCommonStockOk() (*float64, bool) {
	if o == nil || IsNil(o.CommonStock) {
		return nil, false
	}
	return o.CommonStock, true
}

// HasCommonStock returns a boolean if a field has been set.
func (o *EquityInfoCapitalStock) HasCommonStock() bool {
	if o != nil && !IsNil(o.CommonStock) {
		return true
	}

	return false
}

// SetCommonStock gets a reference to the given float64 and assigns it to the CommonStock field.
func (o *EquityInfoCapitalStock) SetCommonStock(v float64) {
	o.CommonStock = &v
}

// GetPreferredStock returns the PreferredStock field value if set, zero value otherwise.
func (o *EquityInfoCapitalStock) GetPreferredStock() float64 {
	if o == nil || IsNil(o.PreferredStock) {
		var ret float64
		return ret
	}
	return *o.PreferredStock
}

// GetPreferredStockOk returns a tuple with the PreferredStock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfoCapitalStock) GetPreferredStockOk() (*float64, bool) {
	if o == nil || IsNil(o.PreferredStock) {
		return nil, false
	}
	return o.PreferredStock, true
}

// HasPreferredStock returns a boolean if a field has been set.
func (o *EquityInfoCapitalStock) HasPreferredStock() bool {
	if o != nil && !IsNil(o.PreferredStock) {
		return true
	}

	return false
}

// SetPreferredStock gets a reference to the given float64 and assigns it to the PreferredStock field.
func (o *EquityInfoCapitalStock) SetPreferredStock(v float64) {
	o.PreferredStock = &v
}

// GetTotalPartnershipCapital returns the TotalPartnershipCapital field value if set, zero value otherwise.
func (o *EquityInfoCapitalStock) GetTotalPartnershipCapital() float64 {
	if o == nil || IsNil(o.TotalPartnershipCapital) {
		var ret float64
		return ret
	}
	return *o.TotalPartnershipCapital
}

// GetTotalPartnershipCapitalOk returns a tuple with the TotalPartnershipCapital field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfoCapitalStock) GetTotalPartnershipCapitalOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalPartnershipCapital) {
		return nil, false
	}
	return o.TotalPartnershipCapital, true
}

// HasTotalPartnershipCapital returns a boolean if a field has been set.
func (o *EquityInfoCapitalStock) HasTotalPartnershipCapital() bool {
	if o != nil && !IsNil(o.TotalPartnershipCapital) {
		return true
	}

	return false
}

// SetTotalPartnershipCapital gets a reference to the given float64 and assigns it to the TotalPartnershipCapital field.
func (o *EquityInfoCapitalStock) SetTotalPartnershipCapital(v float64) {
	o.TotalPartnershipCapital = &v
}

// GetGeneralPartnershipCapital returns the GeneralPartnershipCapital field value if set, zero value otherwise.
func (o *EquityInfoCapitalStock) GetGeneralPartnershipCapital() float64 {
	if o == nil || IsNil(o.GeneralPartnershipCapital) {
		var ret float64
		return ret
	}
	return *o.GeneralPartnershipCapital
}

// GetGeneralPartnershipCapitalOk returns a tuple with the GeneralPartnershipCapital field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfoCapitalStock) GetGeneralPartnershipCapitalOk() (*float64, bool) {
	if o == nil || IsNil(o.GeneralPartnershipCapital) {
		return nil, false
	}
	return o.GeneralPartnershipCapital, true
}

// HasGeneralPartnershipCapital returns a boolean if a field has been set.
func (o *EquityInfoCapitalStock) HasGeneralPartnershipCapital() bool {
	if o != nil && !IsNil(o.GeneralPartnershipCapital) {
		return true
	}

	return false
}

// SetGeneralPartnershipCapital gets a reference to the given float64 and assigns it to the GeneralPartnershipCapital field.
func (o *EquityInfoCapitalStock) SetGeneralPartnershipCapital(v float64) {
	o.GeneralPartnershipCapital = &v
}

// GetLimitedPartnershipCapital returns the LimitedPartnershipCapital field value if set, zero value otherwise.
func (o *EquityInfoCapitalStock) GetLimitedPartnershipCapital() float64 {
	if o == nil || IsNil(o.LimitedPartnershipCapital) {
		var ret float64
		return ret
	}
	return *o.LimitedPartnershipCapital
}

// GetLimitedPartnershipCapitalOk returns a tuple with the LimitedPartnershipCapital field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfoCapitalStock) GetLimitedPartnershipCapitalOk() (*float64, bool) {
	if o == nil || IsNil(o.LimitedPartnershipCapital) {
		return nil, false
	}
	return o.LimitedPartnershipCapital, true
}

// HasLimitedPartnershipCapital returns a boolean if a field has been set.
func (o *EquityInfoCapitalStock) HasLimitedPartnershipCapital() bool {
	if o != nil && !IsNil(o.LimitedPartnershipCapital) {
		return true
	}

	return false
}

// SetLimitedPartnershipCapital gets a reference to the given float64 and assigns it to the LimitedPartnershipCapital field.
func (o *EquityInfoCapitalStock) SetLimitedPartnershipCapital(v float64) {
	o.LimitedPartnershipCapital = &v
}

// GetCapitalStock returns the CapitalStock field value if set, zero value otherwise.
func (o *EquityInfoCapitalStock) GetCapitalStock() float64 {
	if o == nil || IsNil(o.CapitalStock) {
		var ret float64
		return ret
	}
	return *o.CapitalStock
}

// GetCapitalStockOk returns a tuple with the CapitalStock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfoCapitalStock) GetCapitalStockOk() (*float64, bool) {
	if o == nil || IsNil(o.CapitalStock) {
		return nil, false
	}
	return o.CapitalStock, true
}

// HasCapitalStock returns a boolean if a field has been set.
func (o *EquityInfoCapitalStock) HasCapitalStock() bool {
	if o != nil && !IsNil(o.CapitalStock) {
		return true
	}

	return false
}

// SetCapitalStock gets a reference to the given float64 and assigns it to the CapitalStock field.
func (o *EquityInfoCapitalStock) SetCapitalStock(v float64) {
	o.CapitalStock = &v
}

// GetOtherCapitalStock returns the OtherCapitalStock field value if set, zero value otherwise.
func (o *EquityInfoCapitalStock) GetOtherCapitalStock() float64 {
	if o == nil || IsNil(o.OtherCapitalStock) {
		var ret float64
		return ret
	}
	return *o.OtherCapitalStock
}

// GetOtherCapitalStockOk returns a tuple with the OtherCapitalStock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfoCapitalStock) GetOtherCapitalStockOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherCapitalStock) {
		return nil, false
	}
	return o.OtherCapitalStock, true
}

// HasOtherCapitalStock returns a boolean if a field has been set.
func (o *EquityInfoCapitalStock) HasOtherCapitalStock() bool {
	if o != nil && !IsNil(o.OtherCapitalStock) {
		return true
	}

	return false
}

// SetOtherCapitalStock gets a reference to the given float64 and assigns it to the OtherCapitalStock field.
func (o *EquityInfoCapitalStock) SetOtherCapitalStock(v float64) {
	o.OtherCapitalStock = &v
}

// GetAdditionalPaidInCapital returns the AdditionalPaidInCapital field value if set, zero value otherwise.
func (o *EquityInfoCapitalStock) GetAdditionalPaidInCapital() float64 {
	if o == nil || IsNil(o.AdditionalPaidInCapital) {
		var ret float64
		return ret
	}
	return *o.AdditionalPaidInCapital
}

// GetAdditionalPaidInCapitalOk returns a tuple with the AdditionalPaidInCapital field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfoCapitalStock) GetAdditionalPaidInCapitalOk() (*float64, bool) {
	if o == nil || IsNil(o.AdditionalPaidInCapital) {
		return nil, false
	}
	return o.AdditionalPaidInCapital, true
}

// HasAdditionalPaidInCapital returns a boolean if a field has been set.
func (o *EquityInfoCapitalStock) HasAdditionalPaidInCapital() bool {
	if o != nil && !IsNil(o.AdditionalPaidInCapital) {
		return true
	}

	return false
}

// SetAdditionalPaidInCapital gets a reference to the given float64 and assigns it to the AdditionalPaidInCapital field.
func (o *EquityInfoCapitalStock) SetAdditionalPaidInCapital(v float64) {
	o.AdditionalPaidInCapital = &v
}

// GetRetainedEarnings returns the RetainedEarnings field value if set, zero value otherwise.
func (o *EquityInfoCapitalStock) GetRetainedEarnings() float64 {
	if o == nil || IsNil(o.RetainedEarnings) {
		var ret float64
		return ret
	}
	return *o.RetainedEarnings
}

// GetRetainedEarningsOk returns a tuple with the RetainedEarnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfoCapitalStock) GetRetainedEarningsOk() (*float64, bool) {
	if o == nil || IsNil(o.RetainedEarnings) {
		return nil, false
	}
	return o.RetainedEarnings, true
}

// HasRetainedEarnings returns a boolean if a field has been set.
func (o *EquityInfoCapitalStock) HasRetainedEarnings() bool {
	if o != nil && !IsNil(o.RetainedEarnings) {
		return true
	}

	return false
}

// SetRetainedEarnings gets a reference to the given float64 and assigns it to the RetainedEarnings field.
func (o *EquityInfoCapitalStock) SetRetainedEarnings(v float64) {
	o.RetainedEarnings = &v
}

// GetTreasuryStock returns the TreasuryStock field value if set, zero value otherwise.
func (o *EquityInfoCapitalStock) GetTreasuryStock() float64 {
	if o == nil || IsNil(o.TreasuryStock) {
		var ret float64
		return ret
	}
	return *o.TreasuryStock
}

// GetTreasuryStockOk returns a tuple with the TreasuryStock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfoCapitalStock) GetTreasuryStockOk() (*float64, bool) {
	if o == nil || IsNil(o.TreasuryStock) {
		return nil, false
	}
	return o.TreasuryStock, true
}

// HasTreasuryStock returns a boolean if a field has been set.
func (o *EquityInfoCapitalStock) HasTreasuryStock() bool {
	if o != nil && !IsNil(o.TreasuryStock) {
		return true
	}

	return false
}

// SetTreasuryStock gets a reference to the given float64 and assigns it to the TreasuryStock field.
func (o *EquityInfoCapitalStock) SetTreasuryStock(v float64) {
	o.TreasuryStock = &v
}

// GetTreasurySharesNumber returns the TreasurySharesNumber field value if set, zero value otherwise.
func (o *EquityInfoCapitalStock) GetTreasurySharesNumber() float64 {
	if o == nil || IsNil(o.TreasurySharesNumber) {
		var ret float64
		return ret
	}
	return *o.TreasurySharesNumber
}

// GetTreasurySharesNumberOk returns a tuple with the TreasurySharesNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfoCapitalStock) GetTreasurySharesNumberOk() (*float64, bool) {
	if o == nil || IsNil(o.TreasurySharesNumber) {
		return nil, false
	}
	return o.TreasurySharesNumber, true
}

// HasTreasurySharesNumber returns a boolean if a field has been set.
func (o *EquityInfoCapitalStock) HasTreasurySharesNumber() bool {
	if o != nil && !IsNil(o.TreasurySharesNumber) {
		return true
	}

	return false
}

// SetTreasurySharesNumber gets a reference to the given float64 and assigns it to the TreasurySharesNumber field.
func (o *EquityInfoCapitalStock) SetTreasurySharesNumber(v float64) {
	o.TreasurySharesNumber = &v
}

// GetOrdinarySharesNumber returns the OrdinarySharesNumber field value if set, zero value otherwise.
func (o *EquityInfoCapitalStock) GetOrdinarySharesNumber() float64 {
	if o == nil || IsNil(o.OrdinarySharesNumber) {
		var ret float64
		return ret
	}
	return *o.OrdinarySharesNumber
}

// GetOrdinarySharesNumberOk returns a tuple with the OrdinarySharesNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfoCapitalStock) GetOrdinarySharesNumberOk() (*float64, bool) {
	if o == nil || IsNil(o.OrdinarySharesNumber) {
		return nil, false
	}
	return o.OrdinarySharesNumber, true
}

// HasOrdinarySharesNumber returns a boolean if a field has been set.
func (o *EquityInfoCapitalStock) HasOrdinarySharesNumber() bool {
	if o != nil && !IsNil(o.OrdinarySharesNumber) {
		return true
	}

	return false
}

// SetOrdinarySharesNumber gets a reference to the given float64 and assigns it to the OrdinarySharesNumber field.
func (o *EquityInfoCapitalStock) SetOrdinarySharesNumber(v float64) {
	o.OrdinarySharesNumber = &v
}

// GetPreferredSharesNumber returns the PreferredSharesNumber field value if set, zero value otherwise.
func (o *EquityInfoCapitalStock) GetPreferredSharesNumber() float64 {
	if o == nil || IsNil(o.PreferredSharesNumber) {
		var ret float64
		return ret
	}
	return *o.PreferredSharesNumber
}

// GetPreferredSharesNumberOk returns a tuple with the PreferredSharesNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfoCapitalStock) GetPreferredSharesNumberOk() (*float64, bool) {
	if o == nil || IsNil(o.PreferredSharesNumber) {
		return nil, false
	}
	return o.PreferredSharesNumber, true
}

// HasPreferredSharesNumber returns a boolean if a field has been set.
func (o *EquityInfoCapitalStock) HasPreferredSharesNumber() bool {
	if o != nil && !IsNil(o.PreferredSharesNumber) {
		return true
	}

	return false
}

// SetPreferredSharesNumber gets a reference to the given float64 and assigns it to the PreferredSharesNumber field.
func (o *EquityInfoCapitalStock) SetPreferredSharesNumber(v float64) {
	o.PreferredSharesNumber = &v
}

// GetShareIssued returns the ShareIssued field value if set, zero value otherwise.
func (o *EquityInfoCapitalStock) GetShareIssued() float64 {
	if o == nil || IsNil(o.ShareIssued) {
		var ret float64
		return ret
	}
	return *o.ShareIssued
}

// GetShareIssuedOk returns a tuple with the ShareIssued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfoCapitalStock) GetShareIssuedOk() (*float64, bool) {
	if o == nil || IsNil(o.ShareIssued) {
		return nil, false
	}
	return o.ShareIssued, true
}

// HasShareIssued returns a boolean if a field has been set.
func (o *EquityInfoCapitalStock) HasShareIssued() bool {
	if o != nil && !IsNil(o.ShareIssued) {
		return true
	}

	return false
}

// SetShareIssued gets a reference to the given float64 and assigns it to the ShareIssued field.
func (o *EquityInfoCapitalStock) SetShareIssued(v float64) {
	o.ShareIssued = &v
}

func (o EquityInfoCapitalStock) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EquityInfoCapitalStock) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CommonStock) {
		toSerialize["common_stock"] = o.CommonStock
	}
	if !IsNil(o.PreferredStock) {
		toSerialize["preferred_stock"] = o.PreferredStock
	}
	if !IsNil(o.TotalPartnershipCapital) {
		toSerialize["total_partnership_capital"] = o.TotalPartnershipCapital
	}
	if !IsNil(o.GeneralPartnershipCapital) {
		toSerialize["general_partnership_capital"] = o.GeneralPartnershipCapital
	}
	if !IsNil(o.LimitedPartnershipCapital) {
		toSerialize["limited_partnership_capital"] = o.LimitedPartnershipCapital
	}
	if !IsNil(o.CapitalStock) {
		toSerialize["capital_stock"] = o.CapitalStock
	}
	if !IsNil(o.OtherCapitalStock) {
		toSerialize["other_capital_stock"] = o.OtherCapitalStock
	}
	if !IsNil(o.AdditionalPaidInCapital) {
		toSerialize["additional_paid_in_capital"] = o.AdditionalPaidInCapital
	}
	if !IsNil(o.RetainedEarnings) {
		toSerialize["retained_earnings"] = o.RetainedEarnings
	}
	if !IsNil(o.TreasuryStock) {
		toSerialize["treasury_stock"] = o.TreasuryStock
	}
	if !IsNil(o.TreasurySharesNumber) {
		toSerialize["treasury_shares_number"] = o.TreasurySharesNumber
	}
	if !IsNil(o.OrdinarySharesNumber) {
		toSerialize["ordinary_shares_number"] = o.OrdinarySharesNumber
	}
	if !IsNil(o.PreferredSharesNumber) {
		toSerialize["preferred_shares_number"] = o.PreferredSharesNumber
	}
	if !IsNil(o.ShareIssued) {
		toSerialize["share_issued"] = o.ShareIssued
	}
	return toSerialize, nil
}

type NullableEquityInfoCapitalStock struct {
	value *EquityInfoCapitalStock
	isSet bool
}

func (v NullableEquityInfoCapitalStock) Get() *EquityInfoCapitalStock {
	return v.value
}

func (v *NullableEquityInfoCapitalStock) Set(val *EquityInfoCapitalStock) {
	v.value = val
	v.isSet = true
}

func (v NullableEquityInfoCapitalStock) IsSet() bool {
	return v.isSet
}

func (v *NullableEquityInfoCapitalStock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquityInfoCapitalStock(val *EquityInfoCapitalStock) *NullableEquityInfoCapitalStock {
	return &NullableEquityInfoCapitalStock{value: val, isSet: true}
}

func (v NullableEquityInfoCapitalStock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquityInfoCapitalStock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
