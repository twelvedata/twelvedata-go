// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the IncomeStatementItemGrossProfit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IncomeStatementItemGrossProfit{}

// IncomeStatementItemGrossProfit Gross profit information
type IncomeStatementItemGrossProfit struct {
	// Gross profit value
	GrossProfitValue *float64                                     `json:"gross_profit_value,omitempty"`
	CostOfRevenue    *IncomeStatementItemGrossProfitCostOfRevenue `json:"cost_of_revenue,omitempty"`
}

// NewIncomeStatementItemGrossProfit instantiates a new IncomeStatementItemGrossProfit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeStatementItemGrossProfit() *IncomeStatementItemGrossProfit {
	this := IncomeStatementItemGrossProfit{}
	return &this
}

// NewIncomeStatementItemGrossProfitWithDefaults instantiates a new IncomeStatementItemGrossProfit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeStatementItemGrossProfitWithDefaults() *IncomeStatementItemGrossProfit {
	this := IncomeStatementItemGrossProfit{}
	return &this
}

// GetGrossProfitValue returns the GrossProfitValue field value if set, zero value otherwise.
func (o *IncomeStatementItemGrossProfit) GetGrossProfitValue() float64 {
	if o == nil || IsNil(o.GrossProfitValue) {
		var ret float64
		return ret
	}
	return *o.GrossProfitValue
}

// GetGrossProfitValueOk returns a tuple with the GrossProfitValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemGrossProfit) GetGrossProfitValueOk() (*float64, bool) {
	if o == nil || IsNil(o.GrossProfitValue) {
		return nil, false
	}
	return o.GrossProfitValue, true
}

// HasGrossProfitValue returns a boolean if a field has been set.
func (o *IncomeStatementItemGrossProfit) HasGrossProfitValue() bool {
	if o != nil && !IsNil(o.GrossProfitValue) {
		return true
	}

	return false
}

// SetGrossProfitValue gets a reference to the given float64 and assigns it to the GrossProfitValue field.
func (o *IncomeStatementItemGrossProfit) SetGrossProfitValue(v float64) {
	o.GrossProfitValue = &v
}

// GetCostOfRevenue returns the CostOfRevenue field value if set, zero value otherwise.
func (o *IncomeStatementItemGrossProfit) GetCostOfRevenue() IncomeStatementItemGrossProfitCostOfRevenue {
	if o == nil || IsNil(o.CostOfRevenue) {
		var ret IncomeStatementItemGrossProfitCostOfRevenue
		return ret
	}
	return *o.CostOfRevenue
}

// GetCostOfRevenueOk returns a tuple with the CostOfRevenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemGrossProfit) GetCostOfRevenueOk() (*IncomeStatementItemGrossProfitCostOfRevenue, bool) {
	if o == nil || IsNil(o.CostOfRevenue) {
		return nil, false
	}
	return o.CostOfRevenue, true
}

// HasCostOfRevenue returns a boolean if a field has been set.
func (o *IncomeStatementItemGrossProfit) HasCostOfRevenue() bool {
	if o != nil && !IsNil(o.CostOfRevenue) {
		return true
	}

	return false
}

// SetCostOfRevenue gets a reference to the given IncomeStatementItemGrossProfitCostOfRevenue and assigns it to the CostOfRevenue field.
func (o *IncomeStatementItemGrossProfit) SetCostOfRevenue(v IncomeStatementItemGrossProfitCostOfRevenue) {
	o.CostOfRevenue = &v
}

func (o IncomeStatementItemGrossProfit) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IncomeStatementItemGrossProfit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GrossProfitValue) {
		toSerialize["gross_profit_value"] = o.GrossProfitValue
	}
	if !IsNil(o.CostOfRevenue) {
		toSerialize["cost_of_revenue"] = o.CostOfRevenue
	}
	return toSerialize, nil
}

type NullableIncomeStatementItemGrossProfit struct {
	value *IncomeStatementItemGrossProfit
	isSet bool
}

func (v NullableIncomeStatementItemGrossProfit) Get() *IncomeStatementItemGrossProfit {
	return v.value
}

func (v *NullableIncomeStatementItemGrossProfit) Set(val *IncomeStatementItemGrossProfit) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeStatementItemGrossProfit) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeStatementItemGrossProfit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeStatementItemGrossProfit(val *IncomeStatementItemGrossProfit) *NullableIncomeStatementItemGrossProfit {
	return &NullableIncomeStatementItemGrossProfit{value: val, isSet: true}
}

func (v NullableIncomeStatementItemGrossProfit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeStatementItemGrossProfit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
