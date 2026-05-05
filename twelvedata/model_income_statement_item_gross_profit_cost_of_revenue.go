// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the IncomeStatementItemGrossProfitCostOfRevenue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IncomeStatementItemGrossProfitCostOfRevenue{}

// IncomeStatementItemGrossProfitCostOfRevenue Cost of revenue information
type IncomeStatementItemGrossProfitCostOfRevenue struct {
	// Cost of revenue value
	CostOfRevenueValue *float64 `json:"cost_of_revenue_value,omitempty"`
	// Excise taxes
	ExciseTaxes *float64 `json:"excise_taxes,omitempty"`
	// Reconciled cost of revenue
	ReconciledCostOfRevenue *float64 `json:"reconciled_cost_of_revenue,omitempty"`
}

// NewIncomeStatementItemGrossProfitCostOfRevenue instantiates a new IncomeStatementItemGrossProfitCostOfRevenue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeStatementItemGrossProfitCostOfRevenue() *IncomeStatementItemGrossProfitCostOfRevenue {
	this := IncomeStatementItemGrossProfitCostOfRevenue{}
	return &this
}

// NewIncomeStatementItemGrossProfitCostOfRevenueWithDefaults instantiates a new IncomeStatementItemGrossProfitCostOfRevenue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeStatementItemGrossProfitCostOfRevenueWithDefaults() *IncomeStatementItemGrossProfitCostOfRevenue {
	this := IncomeStatementItemGrossProfitCostOfRevenue{}
	return &this
}

// GetCostOfRevenueValue returns the CostOfRevenueValue field value if set, zero value otherwise.
func (o *IncomeStatementItemGrossProfitCostOfRevenue) GetCostOfRevenueValue() float64 {
	if o == nil || IsNil(o.CostOfRevenueValue) {
		var ret float64
		return ret
	}
	return *o.CostOfRevenueValue
}

// GetCostOfRevenueValueOk returns a tuple with the CostOfRevenueValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemGrossProfitCostOfRevenue) GetCostOfRevenueValueOk() (*float64, bool) {
	if o == nil || IsNil(o.CostOfRevenueValue) {
		return nil, false
	}
	return o.CostOfRevenueValue, true
}

// HasCostOfRevenueValue returns a boolean if a field has been set.
func (o *IncomeStatementItemGrossProfitCostOfRevenue) HasCostOfRevenueValue() bool {
	if o != nil && !IsNil(o.CostOfRevenueValue) {
		return true
	}

	return false
}

// SetCostOfRevenueValue gets a reference to the given float64 and assigns it to the CostOfRevenueValue field.
func (o *IncomeStatementItemGrossProfitCostOfRevenue) SetCostOfRevenueValue(v float64) {
	o.CostOfRevenueValue = &v
}

// GetExciseTaxes returns the ExciseTaxes field value if set, zero value otherwise.
func (o *IncomeStatementItemGrossProfitCostOfRevenue) GetExciseTaxes() float64 {
	if o == nil || IsNil(o.ExciseTaxes) {
		var ret float64
		return ret
	}
	return *o.ExciseTaxes
}

// GetExciseTaxesOk returns a tuple with the ExciseTaxes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemGrossProfitCostOfRevenue) GetExciseTaxesOk() (*float64, bool) {
	if o == nil || IsNil(o.ExciseTaxes) {
		return nil, false
	}
	return o.ExciseTaxes, true
}

// HasExciseTaxes returns a boolean if a field has been set.
func (o *IncomeStatementItemGrossProfitCostOfRevenue) HasExciseTaxes() bool {
	if o != nil && !IsNil(o.ExciseTaxes) {
		return true
	}

	return false
}

// SetExciseTaxes gets a reference to the given float64 and assigns it to the ExciseTaxes field.
func (o *IncomeStatementItemGrossProfitCostOfRevenue) SetExciseTaxes(v float64) {
	o.ExciseTaxes = &v
}

// GetReconciledCostOfRevenue returns the ReconciledCostOfRevenue field value if set, zero value otherwise.
func (o *IncomeStatementItemGrossProfitCostOfRevenue) GetReconciledCostOfRevenue() float64 {
	if o == nil || IsNil(o.ReconciledCostOfRevenue) {
		var ret float64
		return ret
	}
	return *o.ReconciledCostOfRevenue
}

// GetReconciledCostOfRevenueOk returns a tuple with the ReconciledCostOfRevenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemGrossProfitCostOfRevenue) GetReconciledCostOfRevenueOk() (*float64, bool) {
	if o == nil || IsNil(o.ReconciledCostOfRevenue) {
		return nil, false
	}
	return o.ReconciledCostOfRevenue, true
}

// HasReconciledCostOfRevenue returns a boolean if a field has been set.
func (o *IncomeStatementItemGrossProfitCostOfRevenue) HasReconciledCostOfRevenue() bool {
	if o != nil && !IsNil(o.ReconciledCostOfRevenue) {
		return true
	}

	return false
}

// SetReconciledCostOfRevenue gets a reference to the given float64 and assigns it to the ReconciledCostOfRevenue field.
func (o *IncomeStatementItemGrossProfitCostOfRevenue) SetReconciledCostOfRevenue(v float64) {
	o.ReconciledCostOfRevenue = &v
}

func (o IncomeStatementItemGrossProfitCostOfRevenue) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IncomeStatementItemGrossProfitCostOfRevenue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CostOfRevenueValue) {
		toSerialize["cost_of_revenue_value"] = o.CostOfRevenueValue
	}
	if !IsNil(o.ExciseTaxes) {
		toSerialize["excise_taxes"] = o.ExciseTaxes
	}
	if !IsNil(o.ReconciledCostOfRevenue) {
		toSerialize["reconciled_cost_of_revenue"] = o.ReconciledCostOfRevenue
	}
	return toSerialize, nil
}

type NullableIncomeStatementItemGrossProfitCostOfRevenue struct {
	value *IncomeStatementItemGrossProfitCostOfRevenue
	isSet bool
}

func (v NullableIncomeStatementItemGrossProfitCostOfRevenue) Get() *IncomeStatementItemGrossProfitCostOfRevenue {
	return v.value
}

func (v *NullableIncomeStatementItemGrossProfitCostOfRevenue) Set(val *IncomeStatementItemGrossProfitCostOfRevenue) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeStatementItemGrossProfitCostOfRevenue) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeStatementItemGrossProfitCostOfRevenue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeStatementItemGrossProfitCostOfRevenue(val *IncomeStatementItemGrossProfitCostOfRevenue) *NullableIncomeStatementItemGrossProfitCostOfRevenue {
	return &NullableIncomeStatementItemGrossProfitCostOfRevenue{value: val, isSet: true}
}

func (v NullableIncomeStatementItemGrossProfitCostOfRevenue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeStatementItemGrossProfitCostOfRevenue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
