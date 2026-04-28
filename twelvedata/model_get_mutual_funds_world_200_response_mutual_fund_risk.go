/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the GetMutualFundsWorld200ResponseMutualFundRisk type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMutualFundsWorld200ResponseMutualFundRisk{}

// GetMutualFundsWorld200ResponseMutualFundRisk Risk metrics of a mutual fund
type GetMutualFundsWorld200ResponseMutualFundRisk struct {
	// Volatility statistics of the fund
	VolatilityMeasures []GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner `json:"volatility_measures,omitempty"`
	ValuationMetrics   *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics         `json:"valuation_metrics,omitempty"`
}

// NewGetMutualFundsWorld200ResponseMutualFundRisk instantiates a new GetMutualFundsWorld200ResponseMutualFundRisk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMutualFundsWorld200ResponseMutualFundRisk() *GetMutualFundsWorld200ResponseMutualFundRisk {
	this := GetMutualFundsWorld200ResponseMutualFundRisk{}
	return &this
}

// NewGetMutualFundsWorld200ResponseMutualFundRiskWithDefaults instantiates a new GetMutualFundsWorld200ResponseMutualFundRisk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMutualFundsWorld200ResponseMutualFundRiskWithDefaults() *GetMutualFundsWorld200ResponseMutualFundRisk {
	this := GetMutualFundsWorld200ResponseMutualFundRisk{}
	return &this
}

// GetVolatilityMeasures returns the VolatilityMeasures field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRisk) GetVolatilityMeasures() []GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner {
	if o == nil || IsNil(o.VolatilityMeasures) {
		var ret []GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner
		return ret
	}
	return o.VolatilityMeasures
}

// GetVolatilityMeasuresOk returns a tuple with the VolatilityMeasures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRisk) GetVolatilityMeasuresOk() ([]GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner, bool) {
	if o == nil || IsNil(o.VolatilityMeasures) {
		return nil, false
	}
	return o.VolatilityMeasures, true
}

// HasVolatilityMeasures returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRisk) HasVolatilityMeasures() bool {
	if o != nil && !IsNil(o.VolatilityMeasures) {
		return true
	}

	return false
}

// SetVolatilityMeasures gets a reference to the given []GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner and assigns it to the VolatilityMeasures field.
func (o *GetMutualFundsWorld200ResponseMutualFundRisk) SetVolatilityMeasures(v []GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) {
	o.VolatilityMeasures = v
}

// GetValuationMetrics returns the ValuationMetrics field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRisk) GetValuationMetrics() GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics {
	if o == nil || IsNil(o.ValuationMetrics) {
		var ret GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics
		return ret
	}
	return *o.ValuationMetrics
}

// GetValuationMetricsOk returns a tuple with the ValuationMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRisk) GetValuationMetricsOk() (*GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics, bool) {
	if o == nil || IsNil(o.ValuationMetrics) {
		return nil, false
	}
	return o.ValuationMetrics, true
}

// HasValuationMetrics returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRisk) HasValuationMetrics() bool {
	if o != nil && !IsNil(o.ValuationMetrics) {
		return true
	}

	return false
}

// SetValuationMetrics gets a reference to the given GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics and assigns it to the ValuationMetrics field.
func (o *GetMutualFundsWorld200ResponseMutualFundRisk) SetValuationMetrics(v GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) {
	o.ValuationMetrics = &v
}

func (o GetMutualFundsWorld200ResponseMutualFundRisk) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMutualFundsWorld200ResponseMutualFundRisk) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VolatilityMeasures) {
		toSerialize["volatility_measures"] = o.VolatilityMeasures
	}
	if !IsNil(o.ValuationMetrics) {
		toSerialize["valuation_metrics"] = o.ValuationMetrics
	}
	return toSerialize, nil
}

type NullableGetMutualFundsWorld200ResponseMutualFundRisk struct {
	value *GetMutualFundsWorld200ResponseMutualFundRisk
	isSet bool
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundRisk) Get() *GetMutualFundsWorld200ResponseMutualFundRisk {
	return v.value
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundRisk) Set(val *GetMutualFundsWorld200ResponseMutualFundRisk) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundRisk) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundRisk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMutualFundsWorld200ResponseMutualFundRisk(val *GetMutualFundsWorld200ResponseMutualFundRisk) *NullableGetMutualFundsWorld200ResponseMutualFundRisk {
	return &NullableGetMutualFundsWorld200ResponseMutualFundRisk{value: val, isSet: true}
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundRisk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundRisk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
