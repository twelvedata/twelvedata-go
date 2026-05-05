// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity{}

// GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity Shareholders' equity section of the balance sheet
type GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity struct {
	// Represents net worth of investors shares, which is equal to the total_assets - total_liabilities
	CommonStock *float64 `json:"common_stock,omitempty"`
	// Refers to the profits earned minus dividends paid
	RetainedEarnings *float64 `json:"retained_earnings,omitempty"`
	// Represents other not affecting retained earnings gains and looses
	OtherShareholdersEquity *float64 `json:"other_shareholders_equity,omitempty"`
	// Represents the net worth of a company, which is the amount that would be returned to shareholders if a company's total assets were liquidated, and all of its debts were repaid
	TotalShareholdersEquity *float64 `json:"total_shareholders_equity,omitempty"`
	// Represents the additional paid-in capital, which is the amount shareholders have invested in a company above the par value of its stock
	AdditionalPaidInCapital *float64 `json:"additional_paid_in_capital,omitempty"`
	// Represents the value of shares that have been repurchased by the company and are held in its treasury
	TreasuryStock *float64 `json:"treasury_stock,omitempty"`
	// Represents the portion of shareholders' equity that is attributable to minority shareholders in a subsidiary company
	MinorityInterest *float64 `json:"minority_interest,omitempty"`
}

// NewGetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity instantiates a new GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity() *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity {
	this := GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity{}
	return &this
}

// NewGetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquityWithDefaults instantiates a new GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquityWithDefaults() *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity {
	this := GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity{}
	return &this
}

// GetCommonStock returns the CommonStock field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) GetCommonStock() float64 {
	if o == nil || IsNil(o.CommonStock) {
		var ret float64
		return ret
	}
	return *o.CommonStock
}

// GetCommonStockOk returns a tuple with the CommonStock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) GetCommonStockOk() (*float64, bool) {
	if o == nil || IsNil(o.CommonStock) {
		return nil, false
	}
	return o.CommonStock, true
}

// HasCommonStock returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) HasCommonStock() bool {
	if o != nil && !IsNil(o.CommonStock) {
		return true
	}

	return false
}

// SetCommonStock gets a reference to the given float64 and assigns it to the CommonStock field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) SetCommonStock(v float64) {
	o.CommonStock = &v
}

// GetRetainedEarnings returns the RetainedEarnings field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) GetRetainedEarnings() float64 {
	if o == nil || IsNil(o.RetainedEarnings) {
		var ret float64
		return ret
	}
	return *o.RetainedEarnings
}

// GetRetainedEarningsOk returns a tuple with the RetainedEarnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) GetRetainedEarningsOk() (*float64, bool) {
	if o == nil || IsNil(o.RetainedEarnings) {
		return nil, false
	}
	return o.RetainedEarnings, true
}

// HasRetainedEarnings returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) HasRetainedEarnings() bool {
	if o != nil && !IsNil(o.RetainedEarnings) {
		return true
	}

	return false
}

// SetRetainedEarnings gets a reference to the given float64 and assigns it to the RetainedEarnings field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) SetRetainedEarnings(v float64) {
	o.RetainedEarnings = &v
}

// GetOtherShareholdersEquity returns the OtherShareholdersEquity field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) GetOtherShareholdersEquity() float64 {
	if o == nil || IsNil(o.OtherShareholdersEquity) {
		var ret float64
		return ret
	}
	return *o.OtherShareholdersEquity
}

// GetOtherShareholdersEquityOk returns a tuple with the OtherShareholdersEquity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) GetOtherShareholdersEquityOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherShareholdersEquity) {
		return nil, false
	}
	return o.OtherShareholdersEquity, true
}

// HasOtherShareholdersEquity returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) HasOtherShareholdersEquity() bool {
	if o != nil && !IsNil(o.OtherShareholdersEquity) {
		return true
	}

	return false
}

// SetOtherShareholdersEquity gets a reference to the given float64 and assigns it to the OtherShareholdersEquity field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) SetOtherShareholdersEquity(v float64) {
	o.OtherShareholdersEquity = &v
}

// GetTotalShareholdersEquity returns the TotalShareholdersEquity field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) GetTotalShareholdersEquity() float64 {
	if o == nil || IsNil(o.TotalShareholdersEquity) {
		var ret float64
		return ret
	}
	return *o.TotalShareholdersEquity
}

// GetTotalShareholdersEquityOk returns a tuple with the TotalShareholdersEquity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) GetTotalShareholdersEquityOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalShareholdersEquity) {
		return nil, false
	}
	return o.TotalShareholdersEquity, true
}

// HasTotalShareholdersEquity returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) HasTotalShareholdersEquity() bool {
	if o != nil && !IsNil(o.TotalShareholdersEquity) {
		return true
	}

	return false
}

// SetTotalShareholdersEquity gets a reference to the given float64 and assigns it to the TotalShareholdersEquity field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) SetTotalShareholdersEquity(v float64) {
	o.TotalShareholdersEquity = &v
}

// GetAdditionalPaidInCapital returns the AdditionalPaidInCapital field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) GetAdditionalPaidInCapital() float64 {
	if o == nil || IsNil(o.AdditionalPaidInCapital) {
		var ret float64
		return ret
	}
	return *o.AdditionalPaidInCapital
}

// GetAdditionalPaidInCapitalOk returns a tuple with the AdditionalPaidInCapital field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) GetAdditionalPaidInCapitalOk() (*float64, bool) {
	if o == nil || IsNil(o.AdditionalPaidInCapital) {
		return nil, false
	}
	return o.AdditionalPaidInCapital, true
}

// HasAdditionalPaidInCapital returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) HasAdditionalPaidInCapital() bool {
	if o != nil && !IsNil(o.AdditionalPaidInCapital) {
		return true
	}

	return false
}

// SetAdditionalPaidInCapital gets a reference to the given float64 and assigns it to the AdditionalPaidInCapital field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) SetAdditionalPaidInCapital(v float64) {
	o.AdditionalPaidInCapital = &v
}

// GetTreasuryStock returns the TreasuryStock field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) GetTreasuryStock() float64 {
	if o == nil || IsNil(o.TreasuryStock) {
		var ret float64
		return ret
	}
	return *o.TreasuryStock
}

// GetTreasuryStockOk returns a tuple with the TreasuryStock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) GetTreasuryStockOk() (*float64, bool) {
	if o == nil || IsNil(o.TreasuryStock) {
		return nil, false
	}
	return o.TreasuryStock, true
}

// HasTreasuryStock returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) HasTreasuryStock() bool {
	if o != nil && !IsNil(o.TreasuryStock) {
		return true
	}

	return false
}

// SetTreasuryStock gets a reference to the given float64 and assigns it to the TreasuryStock field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) SetTreasuryStock(v float64) {
	o.TreasuryStock = &v
}

// GetMinorityInterest returns the MinorityInterest field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) GetMinorityInterest() float64 {
	if o == nil || IsNil(o.MinorityInterest) {
		var ret float64
		return ret
	}
	return *o.MinorityInterest
}

// GetMinorityInterestOk returns a tuple with the MinorityInterest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) GetMinorityInterestOk() (*float64, bool) {
	if o == nil || IsNil(o.MinorityInterest) {
		return nil, false
	}
	return o.MinorityInterest, true
}

// HasMinorityInterest returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) HasMinorityInterest() bool {
	if o != nil && !IsNil(o.MinorityInterest) {
		return true
	}

	return false
}

// SetMinorityInterest gets a reference to the given float64 and assigns it to the MinorityInterest field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) SetMinorityInterest(v float64) {
	o.MinorityInterest = &v
}

func (o GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CommonStock) {
		toSerialize["common_stock"] = o.CommonStock
	}
	if !IsNil(o.RetainedEarnings) {
		toSerialize["retained_earnings"] = o.RetainedEarnings
	}
	if !IsNil(o.OtherShareholdersEquity) {
		toSerialize["other_shareholders_equity"] = o.OtherShareholdersEquity
	}
	if !IsNil(o.TotalShareholdersEquity) {
		toSerialize["total_shareholders_equity"] = o.TotalShareholdersEquity
	}
	if !IsNil(o.AdditionalPaidInCapital) {
		toSerialize["additional_paid_in_capital"] = o.AdditionalPaidInCapital
	}
	if !IsNil(o.TreasuryStock) {
		toSerialize["treasury_stock"] = o.TreasuryStock
	}
	if !IsNil(o.MinorityInterest) {
		toSerialize["minority_interest"] = o.MinorityInterest
	}
	return toSerialize, nil
}

type NullableGetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity struct {
	value *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity
	isSet bool
}

func (v NullableGetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) Get() *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity {
	return v.value
}

func (v *NullableGetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) Set(val *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity(val *GetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) *NullableGetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity {
	return &NullableGetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity{value: val, isSet: true}
}

func (v NullableGetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBalanceSheet200ResponseBalanceSheetInnerShareholdersEquity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
