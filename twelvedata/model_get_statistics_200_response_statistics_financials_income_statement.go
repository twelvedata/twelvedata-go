/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the GetStatistics200ResponseStatisticsFinancialsIncomeStatement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetStatistics200ResponseStatisticsFinancialsIncomeStatement{}

// GetStatistics200ResponseStatisticsFinancialsIncomeStatement Income statement information
type GetStatistics200ResponseStatisticsFinancialsIncomeStatement struct {
	// Refers to total revenue over the last 12 months
	RevenueTtm *float64 `json:"revenue_ttm,omitempty"`
	// Refers to revenue per share over the last 12 months
	RevenuePerShareTtm *float64 `json:"revenue_per_share_ttm,omitempty"`
	// Refers to quarterly revenue growth year over year
	QuarterlyRevenueGrowth *float64 `json:"quarterly_revenue_growth,omitempty"`
	// Refers to gross profit over the last 12 months
	GrossProfitTtm *float64 `json:"gross_profit_ttm,omitempty"`
	// Refers to EBITDA (earnings before interest, taxes, depreciation, and amortization) measure; EBITDA is not calculated for banks
	Ebitda *float64 `json:"ebitda,omitempty"`
	// Refers to net income applicable to common shares over the last 12 months
	NetIncomeToCommonTtm *float64 `json:"net_income_to_common_ttm,omitempty"`
	// Refers to diluted EPS measure over the last 12 months
	DilutedEpsTtm *float64 `json:"diluted_eps_ttm,omitempty"`
	// Refers to quarterly earnings growth year over year
	QuarterlyEarningsGrowthYoy *float64 `json:"quarterly_earnings_growth_yoy,omitempty"`
}

// NewGetStatistics200ResponseStatisticsFinancialsIncomeStatement instantiates a new GetStatistics200ResponseStatisticsFinancialsIncomeStatement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStatistics200ResponseStatisticsFinancialsIncomeStatement() *GetStatistics200ResponseStatisticsFinancialsIncomeStatement {
	this := GetStatistics200ResponseStatisticsFinancialsIncomeStatement{}
	return &this
}

// NewGetStatistics200ResponseStatisticsFinancialsIncomeStatementWithDefaults instantiates a new GetStatistics200ResponseStatisticsFinancialsIncomeStatement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStatistics200ResponseStatisticsFinancialsIncomeStatementWithDefaults() *GetStatistics200ResponseStatisticsFinancialsIncomeStatement {
	this := GetStatistics200ResponseStatisticsFinancialsIncomeStatement{}
	return &this
}

// GetRevenueTtm returns the RevenueTtm field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsFinancialsIncomeStatement) GetRevenueTtm() float64 {
	if o == nil || IsNil(o.RevenueTtm) {
		var ret float64
		return ret
	}
	return *o.RevenueTtm
}

// GetRevenueTtmOk returns a tuple with the RevenueTtm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsFinancialsIncomeStatement) GetRevenueTtmOk() (*float64, bool) {
	if o == nil || IsNil(o.RevenueTtm) {
		return nil, false
	}
	return o.RevenueTtm, true
}

// HasRevenueTtm returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsFinancialsIncomeStatement) HasRevenueTtm() bool {
	if o != nil && !IsNil(o.RevenueTtm) {
		return true
	}

	return false
}

// SetRevenueTtm gets a reference to the given float64 and assigns it to the RevenueTtm field.
func (o *GetStatistics200ResponseStatisticsFinancialsIncomeStatement) SetRevenueTtm(v float64) {
	o.RevenueTtm = &v
}

// GetRevenuePerShareTtm returns the RevenuePerShareTtm field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsFinancialsIncomeStatement) GetRevenuePerShareTtm() float64 {
	if o == nil || IsNil(o.RevenuePerShareTtm) {
		var ret float64
		return ret
	}
	return *o.RevenuePerShareTtm
}

// GetRevenuePerShareTtmOk returns a tuple with the RevenuePerShareTtm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsFinancialsIncomeStatement) GetRevenuePerShareTtmOk() (*float64, bool) {
	if o == nil || IsNil(o.RevenuePerShareTtm) {
		return nil, false
	}
	return o.RevenuePerShareTtm, true
}

// HasRevenuePerShareTtm returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsFinancialsIncomeStatement) HasRevenuePerShareTtm() bool {
	if o != nil && !IsNil(o.RevenuePerShareTtm) {
		return true
	}

	return false
}

// SetRevenuePerShareTtm gets a reference to the given float64 and assigns it to the RevenuePerShareTtm field.
func (o *GetStatistics200ResponseStatisticsFinancialsIncomeStatement) SetRevenuePerShareTtm(v float64) {
	o.RevenuePerShareTtm = &v
}

// GetQuarterlyRevenueGrowth returns the QuarterlyRevenueGrowth field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsFinancialsIncomeStatement) GetQuarterlyRevenueGrowth() float64 {
	if o == nil || IsNil(o.QuarterlyRevenueGrowth) {
		var ret float64
		return ret
	}
	return *o.QuarterlyRevenueGrowth
}

// GetQuarterlyRevenueGrowthOk returns a tuple with the QuarterlyRevenueGrowth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsFinancialsIncomeStatement) GetQuarterlyRevenueGrowthOk() (*float64, bool) {
	if o == nil || IsNil(o.QuarterlyRevenueGrowth) {
		return nil, false
	}
	return o.QuarterlyRevenueGrowth, true
}

// HasQuarterlyRevenueGrowth returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsFinancialsIncomeStatement) HasQuarterlyRevenueGrowth() bool {
	if o != nil && !IsNil(o.QuarterlyRevenueGrowth) {
		return true
	}

	return false
}

// SetQuarterlyRevenueGrowth gets a reference to the given float64 and assigns it to the QuarterlyRevenueGrowth field.
func (o *GetStatistics200ResponseStatisticsFinancialsIncomeStatement) SetQuarterlyRevenueGrowth(v float64) {
	o.QuarterlyRevenueGrowth = &v
}

// GetGrossProfitTtm returns the GrossProfitTtm field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsFinancialsIncomeStatement) GetGrossProfitTtm() float64 {
	if o == nil || IsNil(o.GrossProfitTtm) {
		var ret float64
		return ret
	}
	return *o.GrossProfitTtm
}

// GetGrossProfitTtmOk returns a tuple with the GrossProfitTtm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsFinancialsIncomeStatement) GetGrossProfitTtmOk() (*float64, bool) {
	if o == nil || IsNil(o.GrossProfitTtm) {
		return nil, false
	}
	return o.GrossProfitTtm, true
}

// HasGrossProfitTtm returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsFinancialsIncomeStatement) HasGrossProfitTtm() bool {
	if o != nil && !IsNil(o.GrossProfitTtm) {
		return true
	}

	return false
}

// SetGrossProfitTtm gets a reference to the given float64 and assigns it to the GrossProfitTtm field.
func (o *GetStatistics200ResponseStatisticsFinancialsIncomeStatement) SetGrossProfitTtm(v float64) {
	o.GrossProfitTtm = &v
}

// GetEbitda returns the Ebitda field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsFinancialsIncomeStatement) GetEbitda() float64 {
	if o == nil || IsNil(o.Ebitda) {
		var ret float64
		return ret
	}
	return *o.Ebitda
}

// GetEbitdaOk returns a tuple with the Ebitda field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsFinancialsIncomeStatement) GetEbitdaOk() (*float64, bool) {
	if o == nil || IsNil(o.Ebitda) {
		return nil, false
	}
	return o.Ebitda, true
}

// HasEbitda returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsFinancialsIncomeStatement) HasEbitda() bool {
	if o != nil && !IsNil(o.Ebitda) {
		return true
	}

	return false
}

// SetEbitda gets a reference to the given float64 and assigns it to the Ebitda field.
func (o *GetStatistics200ResponseStatisticsFinancialsIncomeStatement) SetEbitda(v float64) {
	o.Ebitda = &v
}

// GetNetIncomeToCommonTtm returns the NetIncomeToCommonTtm field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsFinancialsIncomeStatement) GetNetIncomeToCommonTtm() float64 {
	if o == nil || IsNil(o.NetIncomeToCommonTtm) {
		var ret float64
		return ret
	}
	return *o.NetIncomeToCommonTtm
}

// GetNetIncomeToCommonTtmOk returns a tuple with the NetIncomeToCommonTtm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsFinancialsIncomeStatement) GetNetIncomeToCommonTtmOk() (*float64, bool) {
	if o == nil || IsNil(o.NetIncomeToCommonTtm) {
		return nil, false
	}
	return o.NetIncomeToCommonTtm, true
}

// HasNetIncomeToCommonTtm returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsFinancialsIncomeStatement) HasNetIncomeToCommonTtm() bool {
	if o != nil && !IsNil(o.NetIncomeToCommonTtm) {
		return true
	}

	return false
}

// SetNetIncomeToCommonTtm gets a reference to the given float64 and assigns it to the NetIncomeToCommonTtm field.
func (o *GetStatistics200ResponseStatisticsFinancialsIncomeStatement) SetNetIncomeToCommonTtm(v float64) {
	o.NetIncomeToCommonTtm = &v
}

// GetDilutedEpsTtm returns the DilutedEpsTtm field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsFinancialsIncomeStatement) GetDilutedEpsTtm() float64 {
	if o == nil || IsNil(o.DilutedEpsTtm) {
		var ret float64
		return ret
	}
	return *o.DilutedEpsTtm
}

// GetDilutedEpsTtmOk returns a tuple with the DilutedEpsTtm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsFinancialsIncomeStatement) GetDilutedEpsTtmOk() (*float64, bool) {
	if o == nil || IsNil(o.DilutedEpsTtm) {
		return nil, false
	}
	return o.DilutedEpsTtm, true
}

// HasDilutedEpsTtm returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsFinancialsIncomeStatement) HasDilutedEpsTtm() bool {
	if o != nil && !IsNil(o.DilutedEpsTtm) {
		return true
	}

	return false
}

// SetDilutedEpsTtm gets a reference to the given float64 and assigns it to the DilutedEpsTtm field.
func (o *GetStatistics200ResponseStatisticsFinancialsIncomeStatement) SetDilutedEpsTtm(v float64) {
	o.DilutedEpsTtm = &v
}

// GetQuarterlyEarningsGrowthYoy returns the QuarterlyEarningsGrowthYoy field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsFinancialsIncomeStatement) GetQuarterlyEarningsGrowthYoy() float64 {
	if o == nil || IsNil(o.QuarterlyEarningsGrowthYoy) {
		var ret float64
		return ret
	}
	return *o.QuarterlyEarningsGrowthYoy
}

// GetQuarterlyEarningsGrowthYoyOk returns a tuple with the QuarterlyEarningsGrowthYoy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsFinancialsIncomeStatement) GetQuarterlyEarningsGrowthYoyOk() (*float64, bool) {
	if o == nil || IsNil(o.QuarterlyEarningsGrowthYoy) {
		return nil, false
	}
	return o.QuarterlyEarningsGrowthYoy, true
}

// HasQuarterlyEarningsGrowthYoy returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsFinancialsIncomeStatement) HasQuarterlyEarningsGrowthYoy() bool {
	if o != nil && !IsNil(o.QuarterlyEarningsGrowthYoy) {
		return true
	}

	return false
}

// SetQuarterlyEarningsGrowthYoy gets a reference to the given float64 and assigns it to the QuarterlyEarningsGrowthYoy field.
func (o *GetStatistics200ResponseStatisticsFinancialsIncomeStatement) SetQuarterlyEarningsGrowthYoy(v float64) {
	o.QuarterlyEarningsGrowthYoy = &v
}

func (o GetStatistics200ResponseStatisticsFinancialsIncomeStatement) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetStatistics200ResponseStatisticsFinancialsIncomeStatement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RevenueTtm) {
		toSerialize["revenue_ttm"] = o.RevenueTtm
	}
	if !IsNil(o.RevenuePerShareTtm) {
		toSerialize["revenue_per_share_ttm"] = o.RevenuePerShareTtm
	}
	if !IsNil(o.QuarterlyRevenueGrowth) {
		toSerialize["quarterly_revenue_growth"] = o.QuarterlyRevenueGrowth
	}
	if !IsNil(o.GrossProfitTtm) {
		toSerialize["gross_profit_ttm"] = o.GrossProfitTtm
	}
	if !IsNil(o.Ebitda) {
		toSerialize["ebitda"] = o.Ebitda
	}
	if !IsNil(o.NetIncomeToCommonTtm) {
		toSerialize["net_income_to_common_ttm"] = o.NetIncomeToCommonTtm
	}
	if !IsNil(o.DilutedEpsTtm) {
		toSerialize["diluted_eps_ttm"] = o.DilutedEpsTtm
	}
	if !IsNil(o.QuarterlyEarningsGrowthYoy) {
		toSerialize["quarterly_earnings_growth_yoy"] = o.QuarterlyEarningsGrowthYoy
	}
	return toSerialize, nil
}

type NullableGetStatistics200ResponseStatisticsFinancialsIncomeStatement struct {
	value *GetStatistics200ResponseStatisticsFinancialsIncomeStatement
	isSet bool
}

func (v NullableGetStatistics200ResponseStatisticsFinancialsIncomeStatement) Get() *GetStatistics200ResponseStatisticsFinancialsIncomeStatement {
	return v.value
}

func (v *NullableGetStatistics200ResponseStatisticsFinancialsIncomeStatement) Set(val *GetStatistics200ResponseStatisticsFinancialsIncomeStatement) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStatistics200ResponseStatisticsFinancialsIncomeStatement) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStatistics200ResponseStatisticsFinancialsIncomeStatement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStatistics200ResponseStatisticsFinancialsIncomeStatement(val *GetStatistics200ResponseStatisticsFinancialsIncomeStatement) *NullableGetStatistics200ResponseStatisticsFinancialsIncomeStatement {
	return &NullableGetStatistics200ResponseStatisticsFinancialsIncomeStatement{value: val, isSet: true}
}

func (v NullableGetStatistics200ResponseStatisticsFinancialsIncomeStatement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStatistics200ResponseStatisticsFinancialsIncomeStatement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
