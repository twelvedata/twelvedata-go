// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the GetStatistics200ResponseStatisticsFinancialsCashFlow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetStatistics200ResponseStatisticsFinancialsCashFlow{}

// GetStatistics200ResponseStatisticsFinancialsCashFlow Cash flow information
type GetStatistics200ResponseStatisticsFinancialsCashFlow struct {
	// Refers to operating cash flow measure over the last 12 months
	OperatingCashFlowTtm *float64 `json:"operating_cash_flow_ttm,omitempty"`
	// Refers to levered free cash flow measure over the last 12 months
	LeveredFreeCashFlowTtm *float64 `json:"levered_free_cash_flow_ttm,omitempty"`
}

// NewGetStatistics200ResponseStatisticsFinancialsCashFlow instantiates a new GetStatistics200ResponseStatisticsFinancialsCashFlow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStatistics200ResponseStatisticsFinancialsCashFlow() *GetStatistics200ResponseStatisticsFinancialsCashFlow {
	this := GetStatistics200ResponseStatisticsFinancialsCashFlow{}
	return &this
}

// NewGetStatistics200ResponseStatisticsFinancialsCashFlowWithDefaults instantiates a new GetStatistics200ResponseStatisticsFinancialsCashFlow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStatistics200ResponseStatisticsFinancialsCashFlowWithDefaults() *GetStatistics200ResponseStatisticsFinancialsCashFlow {
	this := GetStatistics200ResponseStatisticsFinancialsCashFlow{}
	return &this
}

// GetOperatingCashFlowTtm returns the OperatingCashFlowTtm field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsFinancialsCashFlow) GetOperatingCashFlowTtm() float64 {
	if o == nil || IsNil(o.OperatingCashFlowTtm) {
		var ret float64
		return ret
	}
	return *o.OperatingCashFlowTtm
}

// GetOperatingCashFlowTtmOk returns a tuple with the OperatingCashFlowTtm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsFinancialsCashFlow) GetOperatingCashFlowTtmOk() (*float64, bool) {
	if o == nil || IsNil(o.OperatingCashFlowTtm) {
		return nil, false
	}
	return o.OperatingCashFlowTtm, true
}

// HasOperatingCashFlowTtm returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsFinancialsCashFlow) HasOperatingCashFlowTtm() bool {
	if o != nil && !IsNil(o.OperatingCashFlowTtm) {
		return true
	}

	return false
}

// SetOperatingCashFlowTtm gets a reference to the given float64 and assigns it to the OperatingCashFlowTtm field.
func (o *GetStatistics200ResponseStatisticsFinancialsCashFlow) SetOperatingCashFlowTtm(v float64) {
	o.OperatingCashFlowTtm = &v
}

// GetLeveredFreeCashFlowTtm returns the LeveredFreeCashFlowTtm field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsFinancialsCashFlow) GetLeveredFreeCashFlowTtm() float64 {
	if o == nil || IsNil(o.LeveredFreeCashFlowTtm) {
		var ret float64
		return ret
	}
	return *o.LeveredFreeCashFlowTtm
}

// GetLeveredFreeCashFlowTtmOk returns a tuple with the LeveredFreeCashFlowTtm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsFinancialsCashFlow) GetLeveredFreeCashFlowTtmOk() (*float64, bool) {
	if o == nil || IsNil(o.LeveredFreeCashFlowTtm) {
		return nil, false
	}
	return o.LeveredFreeCashFlowTtm, true
}

// HasLeveredFreeCashFlowTtm returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsFinancialsCashFlow) HasLeveredFreeCashFlowTtm() bool {
	if o != nil && !IsNil(o.LeveredFreeCashFlowTtm) {
		return true
	}

	return false
}

// SetLeveredFreeCashFlowTtm gets a reference to the given float64 and assigns it to the LeveredFreeCashFlowTtm field.
func (o *GetStatistics200ResponseStatisticsFinancialsCashFlow) SetLeveredFreeCashFlowTtm(v float64) {
	o.LeveredFreeCashFlowTtm = &v
}

func (o GetStatistics200ResponseStatisticsFinancialsCashFlow) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetStatistics200ResponseStatisticsFinancialsCashFlow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OperatingCashFlowTtm) {
		toSerialize["operating_cash_flow_ttm"] = o.OperatingCashFlowTtm
	}
	if !IsNil(o.LeveredFreeCashFlowTtm) {
		toSerialize["levered_free_cash_flow_ttm"] = o.LeveredFreeCashFlowTtm
	}
	return toSerialize, nil
}

type NullableGetStatistics200ResponseStatisticsFinancialsCashFlow struct {
	value *GetStatistics200ResponseStatisticsFinancialsCashFlow
	isSet bool
}

func (v NullableGetStatistics200ResponseStatisticsFinancialsCashFlow) Get() *GetStatistics200ResponseStatisticsFinancialsCashFlow {
	return v.value
}

func (v *NullableGetStatistics200ResponseStatisticsFinancialsCashFlow) Set(val *GetStatistics200ResponseStatisticsFinancialsCashFlow) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStatistics200ResponseStatisticsFinancialsCashFlow) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStatistics200ResponseStatisticsFinancialsCashFlow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStatistics200ResponseStatisticsFinancialsCashFlow(val *GetStatistics200ResponseStatisticsFinancialsCashFlow) *NullableGetStatistics200ResponseStatisticsFinancialsCashFlow {
	return &NullableGetStatistics200ResponseStatisticsFinancialsCashFlow{value: val, isSet: true}
}

func (v NullableGetStatistics200ResponseStatisticsFinancialsCashFlow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStatistics200ResponseStatisticsFinancialsCashFlow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
