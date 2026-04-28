/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the EquityInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EquityInfo{}

// EquityInfo EquityInfo represents equity information
type EquityInfo struct {
	// Total equity gross minority interest
	TotalEquityGrossMinorityInterest *float64 `json:"total_equity_gross_minority_interest,omitempty"`
	// Stockholders equity
	StockholdersEquity *float64 `json:"stockholders_equity,omitempty"`
	// Common stock equity
	CommonStockEquity *float64 `json:"common_stock_equity,omitempty"`
	// Preferred stock equity
	PreferredStockEquity *float64 `json:"preferred_stock_equity,omitempty"`
	// Other equity interest
	OtherEquityInterest *float64 `json:"other_equity_interest,omitempty"`
	// Minority interest
	MinorityInterest *float64 `json:"minority_interest,omitempty"`
	// Total capitalization
	TotalCapitalization *float64 `json:"total_capitalization,omitempty"`
	// Net tangible assets
	NetTangibleAssets *float64 `json:"net_tangible_assets,omitempty"`
	// Tangible book value
	TangibleBookValue *float64 `json:"tangible_book_value,omitempty"`
	// Invested capital
	InvestedCapital *float64 `json:"invested_capital,omitempty"`
	// Working capital
	WorkingCapital    *float64                     `json:"working_capital,omitempty"`
	CapitalStock      *EquityInfoCapitalStock      `json:"capital_stock,omitempty"`
	EquityAdjustments *EquityInfoEquityAdjustments `json:"equity_adjustments,omitempty"`
	// Net debt
	NetDebt *float64 `json:"net_debt,omitempty"`
	// Total debt
	TotalDebt *float64 `json:"total_debt,omitempty"`
}

// NewEquityInfo instantiates a new EquityInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquityInfo() *EquityInfo {
	this := EquityInfo{}
	return &this
}

// NewEquityInfoWithDefaults instantiates a new EquityInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquityInfoWithDefaults() *EquityInfo {
	this := EquityInfo{}
	return &this
}

// GetTotalEquityGrossMinorityInterest returns the TotalEquityGrossMinorityInterest field value if set, zero value otherwise.
func (o *EquityInfo) GetTotalEquityGrossMinorityInterest() float64 {
	if o == nil || IsNil(o.TotalEquityGrossMinorityInterest) {
		var ret float64
		return ret
	}
	return *o.TotalEquityGrossMinorityInterest
}

// GetTotalEquityGrossMinorityInterestOk returns a tuple with the TotalEquityGrossMinorityInterest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfo) GetTotalEquityGrossMinorityInterestOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalEquityGrossMinorityInterest) {
		return nil, false
	}
	return o.TotalEquityGrossMinorityInterest, true
}

// HasTotalEquityGrossMinorityInterest returns a boolean if a field has been set.
func (o *EquityInfo) HasTotalEquityGrossMinorityInterest() bool {
	if o != nil && !IsNil(o.TotalEquityGrossMinorityInterest) {
		return true
	}

	return false
}

// SetTotalEquityGrossMinorityInterest gets a reference to the given float64 and assigns it to the TotalEquityGrossMinorityInterest field.
func (o *EquityInfo) SetTotalEquityGrossMinorityInterest(v float64) {
	o.TotalEquityGrossMinorityInterest = &v
}

// GetStockholdersEquity returns the StockholdersEquity field value if set, zero value otherwise.
func (o *EquityInfo) GetStockholdersEquity() float64 {
	if o == nil || IsNil(o.StockholdersEquity) {
		var ret float64
		return ret
	}
	return *o.StockholdersEquity
}

// GetStockholdersEquityOk returns a tuple with the StockholdersEquity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfo) GetStockholdersEquityOk() (*float64, bool) {
	if o == nil || IsNil(o.StockholdersEquity) {
		return nil, false
	}
	return o.StockholdersEquity, true
}

// HasStockholdersEquity returns a boolean if a field has been set.
func (o *EquityInfo) HasStockholdersEquity() bool {
	if o != nil && !IsNil(o.StockholdersEquity) {
		return true
	}

	return false
}

// SetStockholdersEquity gets a reference to the given float64 and assigns it to the StockholdersEquity field.
func (o *EquityInfo) SetStockholdersEquity(v float64) {
	o.StockholdersEquity = &v
}

// GetCommonStockEquity returns the CommonStockEquity field value if set, zero value otherwise.
func (o *EquityInfo) GetCommonStockEquity() float64 {
	if o == nil || IsNil(o.CommonStockEquity) {
		var ret float64
		return ret
	}
	return *o.CommonStockEquity
}

// GetCommonStockEquityOk returns a tuple with the CommonStockEquity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfo) GetCommonStockEquityOk() (*float64, bool) {
	if o == nil || IsNil(o.CommonStockEquity) {
		return nil, false
	}
	return o.CommonStockEquity, true
}

// HasCommonStockEquity returns a boolean if a field has been set.
func (o *EquityInfo) HasCommonStockEquity() bool {
	if o != nil && !IsNil(o.CommonStockEquity) {
		return true
	}

	return false
}

// SetCommonStockEquity gets a reference to the given float64 and assigns it to the CommonStockEquity field.
func (o *EquityInfo) SetCommonStockEquity(v float64) {
	o.CommonStockEquity = &v
}

// GetPreferredStockEquity returns the PreferredStockEquity field value if set, zero value otherwise.
func (o *EquityInfo) GetPreferredStockEquity() float64 {
	if o == nil || IsNil(o.PreferredStockEquity) {
		var ret float64
		return ret
	}
	return *o.PreferredStockEquity
}

// GetPreferredStockEquityOk returns a tuple with the PreferredStockEquity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfo) GetPreferredStockEquityOk() (*float64, bool) {
	if o == nil || IsNil(o.PreferredStockEquity) {
		return nil, false
	}
	return o.PreferredStockEquity, true
}

// HasPreferredStockEquity returns a boolean if a field has been set.
func (o *EquityInfo) HasPreferredStockEquity() bool {
	if o != nil && !IsNil(o.PreferredStockEquity) {
		return true
	}

	return false
}

// SetPreferredStockEquity gets a reference to the given float64 and assigns it to the PreferredStockEquity field.
func (o *EquityInfo) SetPreferredStockEquity(v float64) {
	o.PreferredStockEquity = &v
}

// GetOtherEquityInterest returns the OtherEquityInterest field value if set, zero value otherwise.
func (o *EquityInfo) GetOtherEquityInterest() float64 {
	if o == nil || IsNil(o.OtherEquityInterest) {
		var ret float64
		return ret
	}
	return *o.OtherEquityInterest
}

// GetOtherEquityInterestOk returns a tuple with the OtherEquityInterest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfo) GetOtherEquityInterestOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherEquityInterest) {
		return nil, false
	}
	return o.OtherEquityInterest, true
}

// HasOtherEquityInterest returns a boolean if a field has been set.
func (o *EquityInfo) HasOtherEquityInterest() bool {
	if o != nil && !IsNil(o.OtherEquityInterest) {
		return true
	}

	return false
}

// SetOtherEquityInterest gets a reference to the given float64 and assigns it to the OtherEquityInterest field.
func (o *EquityInfo) SetOtherEquityInterest(v float64) {
	o.OtherEquityInterest = &v
}

// GetMinorityInterest returns the MinorityInterest field value if set, zero value otherwise.
func (o *EquityInfo) GetMinorityInterest() float64 {
	if o == nil || IsNil(o.MinorityInterest) {
		var ret float64
		return ret
	}
	return *o.MinorityInterest
}

// GetMinorityInterestOk returns a tuple with the MinorityInterest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfo) GetMinorityInterestOk() (*float64, bool) {
	if o == nil || IsNil(o.MinorityInterest) {
		return nil, false
	}
	return o.MinorityInterest, true
}

// HasMinorityInterest returns a boolean if a field has been set.
func (o *EquityInfo) HasMinorityInterest() bool {
	if o != nil && !IsNil(o.MinorityInterest) {
		return true
	}

	return false
}

// SetMinorityInterest gets a reference to the given float64 and assigns it to the MinorityInterest field.
func (o *EquityInfo) SetMinorityInterest(v float64) {
	o.MinorityInterest = &v
}

// GetTotalCapitalization returns the TotalCapitalization field value if set, zero value otherwise.
func (o *EquityInfo) GetTotalCapitalization() float64 {
	if o == nil || IsNil(o.TotalCapitalization) {
		var ret float64
		return ret
	}
	return *o.TotalCapitalization
}

// GetTotalCapitalizationOk returns a tuple with the TotalCapitalization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfo) GetTotalCapitalizationOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalCapitalization) {
		return nil, false
	}
	return o.TotalCapitalization, true
}

// HasTotalCapitalization returns a boolean if a field has been set.
func (o *EquityInfo) HasTotalCapitalization() bool {
	if o != nil && !IsNil(o.TotalCapitalization) {
		return true
	}

	return false
}

// SetTotalCapitalization gets a reference to the given float64 and assigns it to the TotalCapitalization field.
func (o *EquityInfo) SetTotalCapitalization(v float64) {
	o.TotalCapitalization = &v
}

// GetNetTangibleAssets returns the NetTangibleAssets field value if set, zero value otherwise.
func (o *EquityInfo) GetNetTangibleAssets() float64 {
	if o == nil || IsNil(o.NetTangibleAssets) {
		var ret float64
		return ret
	}
	return *o.NetTangibleAssets
}

// GetNetTangibleAssetsOk returns a tuple with the NetTangibleAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfo) GetNetTangibleAssetsOk() (*float64, bool) {
	if o == nil || IsNil(o.NetTangibleAssets) {
		return nil, false
	}
	return o.NetTangibleAssets, true
}

// HasNetTangibleAssets returns a boolean if a field has been set.
func (o *EquityInfo) HasNetTangibleAssets() bool {
	if o != nil && !IsNil(o.NetTangibleAssets) {
		return true
	}

	return false
}

// SetNetTangibleAssets gets a reference to the given float64 and assigns it to the NetTangibleAssets field.
func (o *EquityInfo) SetNetTangibleAssets(v float64) {
	o.NetTangibleAssets = &v
}

// GetTangibleBookValue returns the TangibleBookValue field value if set, zero value otherwise.
func (o *EquityInfo) GetTangibleBookValue() float64 {
	if o == nil || IsNil(o.TangibleBookValue) {
		var ret float64
		return ret
	}
	return *o.TangibleBookValue
}

// GetTangibleBookValueOk returns a tuple with the TangibleBookValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfo) GetTangibleBookValueOk() (*float64, bool) {
	if o == nil || IsNil(o.TangibleBookValue) {
		return nil, false
	}
	return o.TangibleBookValue, true
}

// HasTangibleBookValue returns a boolean if a field has been set.
func (o *EquityInfo) HasTangibleBookValue() bool {
	if o != nil && !IsNil(o.TangibleBookValue) {
		return true
	}

	return false
}

// SetTangibleBookValue gets a reference to the given float64 and assigns it to the TangibleBookValue field.
func (o *EquityInfo) SetTangibleBookValue(v float64) {
	o.TangibleBookValue = &v
}

// GetInvestedCapital returns the InvestedCapital field value if set, zero value otherwise.
func (o *EquityInfo) GetInvestedCapital() float64 {
	if o == nil || IsNil(o.InvestedCapital) {
		var ret float64
		return ret
	}
	return *o.InvestedCapital
}

// GetInvestedCapitalOk returns a tuple with the InvestedCapital field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfo) GetInvestedCapitalOk() (*float64, bool) {
	if o == nil || IsNil(o.InvestedCapital) {
		return nil, false
	}
	return o.InvestedCapital, true
}

// HasInvestedCapital returns a boolean if a field has been set.
func (o *EquityInfo) HasInvestedCapital() bool {
	if o != nil && !IsNil(o.InvestedCapital) {
		return true
	}

	return false
}

// SetInvestedCapital gets a reference to the given float64 and assigns it to the InvestedCapital field.
func (o *EquityInfo) SetInvestedCapital(v float64) {
	o.InvestedCapital = &v
}

// GetWorkingCapital returns the WorkingCapital field value if set, zero value otherwise.
func (o *EquityInfo) GetWorkingCapital() float64 {
	if o == nil || IsNil(o.WorkingCapital) {
		var ret float64
		return ret
	}
	return *o.WorkingCapital
}

// GetWorkingCapitalOk returns a tuple with the WorkingCapital field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfo) GetWorkingCapitalOk() (*float64, bool) {
	if o == nil || IsNil(o.WorkingCapital) {
		return nil, false
	}
	return o.WorkingCapital, true
}

// HasWorkingCapital returns a boolean if a field has been set.
func (o *EquityInfo) HasWorkingCapital() bool {
	if o != nil && !IsNil(o.WorkingCapital) {
		return true
	}

	return false
}

// SetWorkingCapital gets a reference to the given float64 and assigns it to the WorkingCapital field.
func (o *EquityInfo) SetWorkingCapital(v float64) {
	o.WorkingCapital = &v
}

// GetCapitalStock returns the CapitalStock field value if set, zero value otherwise.
func (o *EquityInfo) GetCapitalStock() EquityInfoCapitalStock {
	if o == nil || IsNil(o.CapitalStock) {
		var ret EquityInfoCapitalStock
		return ret
	}
	return *o.CapitalStock
}

// GetCapitalStockOk returns a tuple with the CapitalStock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfo) GetCapitalStockOk() (*EquityInfoCapitalStock, bool) {
	if o == nil || IsNil(o.CapitalStock) {
		return nil, false
	}
	return o.CapitalStock, true
}

// HasCapitalStock returns a boolean if a field has been set.
func (o *EquityInfo) HasCapitalStock() bool {
	if o != nil && !IsNil(o.CapitalStock) {
		return true
	}

	return false
}

// SetCapitalStock gets a reference to the given EquityInfoCapitalStock and assigns it to the CapitalStock field.
func (o *EquityInfo) SetCapitalStock(v EquityInfoCapitalStock) {
	o.CapitalStock = &v
}

// GetEquityAdjustments returns the EquityAdjustments field value if set, zero value otherwise.
func (o *EquityInfo) GetEquityAdjustments() EquityInfoEquityAdjustments {
	if o == nil || IsNil(o.EquityAdjustments) {
		var ret EquityInfoEquityAdjustments
		return ret
	}
	return *o.EquityAdjustments
}

// GetEquityAdjustmentsOk returns a tuple with the EquityAdjustments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfo) GetEquityAdjustmentsOk() (*EquityInfoEquityAdjustments, bool) {
	if o == nil || IsNil(o.EquityAdjustments) {
		return nil, false
	}
	return o.EquityAdjustments, true
}

// HasEquityAdjustments returns a boolean if a field has been set.
func (o *EquityInfo) HasEquityAdjustments() bool {
	if o != nil && !IsNil(o.EquityAdjustments) {
		return true
	}

	return false
}

// SetEquityAdjustments gets a reference to the given EquityInfoEquityAdjustments and assigns it to the EquityAdjustments field.
func (o *EquityInfo) SetEquityAdjustments(v EquityInfoEquityAdjustments) {
	o.EquityAdjustments = &v
}

// GetNetDebt returns the NetDebt field value if set, zero value otherwise.
func (o *EquityInfo) GetNetDebt() float64 {
	if o == nil || IsNil(o.NetDebt) {
		var ret float64
		return ret
	}
	return *o.NetDebt
}

// GetNetDebtOk returns a tuple with the NetDebt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfo) GetNetDebtOk() (*float64, bool) {
	if o == nil || IsNil(o.NetDebt) {
		return nil, false
	}
	return o.NetDebt, true
}

// HasNetDebt returns a boolean if a field has been set.
func (o *EquityInfo) HasNetDebt() bool {
	if o != nil && !IsNil(o.NetDebt) {
		return true
	}

	return false
}

// SetNetDebt gets a reference to the given float64 and assigns it to the NetDebt field.
func (o *EquityInfo) SetNetDebt(v float64) {
	o.NetDebt = &v
}

// GetTotalDebt returns the TotalDebt field value if set, zero value otherwise.
func (o *EquityInfo) GetTotalDebt() float64 {
	if o == nil || IsNil(o.TotalDebt) {
		var ret float64
		return ret
	}
	return *o.TotalDebt
}

// GetTotalDebtOk returns a tuple with the TotalDebt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquityInfo) GetTotalDebtOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalDebt) {
		return nil, false
	}
	return o.TotalDebt, true
}

// HasTotalDebt returns a boolean if a field has been set.
func (o *EquityInfo) HasTotalDebt() bool {
	if o != nil && !IsNil(o.TotalDebt) {
		return true
	}

	return false
}

// SetTotalDebt gets a reference to the given float64 and assigns it to the TotalDebt field.
func (o *EquityInfo) SetTotalDebt(v float64) {
	o.TotalDebt = &v
}

func (o EquityInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EquityInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalEquityGrossMinorityInterest) {
		toSerialize["total_equity_gross_minority_interest"] = o.TotalEquityGrossMinorityInterest
	}
	if !IsNil(o.StockholdersEquity) {
		toSerialize["stockholders_equity"] = o.StockholdersEquity
	}
	if !IsNil(o.CommonStockEquity) {
		toSerialize["common_stock_equity"] = o.CommonStockEquity
	}
	if !IsNil(o.PreferredStockEquity) {
		toSerialize["preferred_stock_equity"] = o.PreferredStockEquity
	}
	if !IsNil(o.OtherEquityInterest) {
		toSerialize["other_equity_interest"] = o.OtherEquityInterest
	}
	if !IsNil(o.MinorityInterest) {
		toSerialize["minority_interest"] = o.MinorityInterest
	}
	if !IsNil(o.TotalCapitalization) {
		toSerialize["total_capitalization"] = o.TotalCapitalization
	}
	if !IsNil(o.NetTangibleAssets) {
		toSerialize["net_tangible_assets"] = o.NetTangibleAssets
	}
	if !IsNil(o.TangibleBookValue) {
		toSerialize["tangible_book_value"] = o.TangibleBookValue
	}
	if !IsNil(o.InvestedCapital) {
		toSerialize["invested_capital"] = o.InvestedCapital
	}
	if !IsNil(o.WorkingCapital) {
		toSerialize["working_capital"] = o.WorkingCapital
	}
	if !IsNil(o.CapitalStock) {
		toSerialize["capital_stock"] = o.CapitalStock
	}
	if !IsNil(o.EquityAdjustments) {
		toSerialize["equity_adjustments"] = o.EquityAdjustments
	}
	if !IsNil(o.NetDebt) {
		toSerialize["net_debt"] = o.NetDebt
	}
	if !IsNil(o.TotalDebt) {
		toSerialize["total_debt"] = o.TotalDebt
	}
	return toSerialize, nil
}

type NullableEquityInfo struct {
	value *EquityInfo
	isSet bool
}

func (v NullableEquityInfo) Get() *EquityInfo {
	return v.value
}

func (v *NullableEquityInfo) Set(val *EquityInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableEquityInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableEquityInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquityInfo(val *EquityInfo) *NullableEquityInfo {
	return &NullableEquityInfo{value: val, isSet: true}
}

func (v NullableEquityInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquityInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
