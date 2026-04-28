/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the ResponseMutualFundWorldRisk type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseMutualFundWorldRisk{}

// ResponseMutualFundWorldRisk Risk and volatility statistics of the fund and its category over different periods
type ResponseMutualFundWorldRisk struct {
	// Volatility statistics of the fund
	VolatilityMeasures []GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner `json:"volatility_measures,omitempty"`
	ValuationMetrics   *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics         `json:"valuation_metrics,omitempty"`
}

// NewResponseMutualFundWorldRisk instantiates a new ResponseMutualFundWorldRisk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseMutualFundWorldRisk() *ResponseMutualFundWorldRisk {
	this := ResponseMutualFundWorldRisk{}
	return &this
}

// NewResponseMutualFundWorldRiskWithDefaults instantiates a new ResponseMutualFundWorldRisk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseMutualFundWorldRiskWithDefaults() *ResponseMutualFundWorldRisk {
	this := ResponseMutualFundWorldRisk{}
	return &this
}

// GetVolatilityMeasures returns the VolatilityMeasures field value if set, zero value otherwise.
func (o *ResponseMutualFundWorldRisk) GetVolatilityMeasures() []GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner {
	if o == nil || IsNil(o.VolatilityMeasures) {
		var ret []GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner
		return ret
	}
	return o.VolatilityMeasures
}

// GetVolatilityMeasuresOk returns a tuple with the VolatilityMeasures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMutualFundWorldRisk) GetVolatilityMeasuresOk() ([]GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner, bool) {
	if o == nil || IsNil(o.VolatilityMeasures) {
		return nil, false
	}
	return o.VolatilityMeasures, true
}

// HasVolatilityMeasures returns a boolean if a field has been set.
func (o *ResponseMutualFundWorldRisk) HasVolatilityMeasures() bool {
	if o != nil && !IsNil(o.VolatilityMeasures) {
		return true
	}

	return false
}

// SetVolatilityMeasures gets a reference to the given []GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner and assigns it to the VolatilityMeasures field.
func (o *ResponseMutualFundWorldRisk) SetVolatilityMeasures(v []GetMutualFundsWorld200ResponseMutualFundRiskVolatilityMeasuresInner) {
	o.VolatilityMeasures = v
}

// GetValuationMetrics returns the ValuationMetrics field value if set, zero value otherwise.
func (o *ResponseMutualFundWorldRisk) GetValuationMetrics() GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics {
	if o == nil || IsNil(o.ValuationMetrics) {
		var ret GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics
		return ret
	}
	return *o.ValuationMetrics
}

// GetValuationMetricsOk returns a tuple with the ValuationMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMutualFundWorldRisk) GetValuationMetricsOk() (*GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics, bool) {
	if o == nil || IsNil(o.ValuationMetrics) {
		return nil, false
	}
	return o.ValuationMetrics, true
}

// HasValuationMetrics returns a boolean if a field has been set.
func (o *ResponseMutualFundWorldRisk) HasValuationMetrics() bool {
	if o != nil && !IsNil(o.ValuationMetrics) {
		return true
	}

	return false
}

// SetValuationMetrics gets a reference to the given GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics and assigns it to the ValuationMetrics field.
func (o *ResponseMutualFundWorldRisk) SetValuationMetrics(v GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) {
	o.ValuationMetrics = &v
}

func (o ResponseMutualFundWorldRisk) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseMutualFundWorldRisk) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VolatilityMeasures) {
		toSerialize["volatility_measures"] = o.VolatilityMeasures
	}
	if !IsNil(o.ValuationMetrics) {
		toSerialize["valuation_metrics"] = o.ValuationMetrics
	}
	return toSerialize, nil
}

type NullableResponseMutualFundWorldRisk struct {
	value *ResponseMutualFundWorldRisk
	isSet bool
}

func (v NullableResponseMutualFundWorldRisk) Get() *ResponseMutualFundWorldRisk {
	return v.value
}

func (v *NullableResponseMutualFundWorldRisk) Set(val *ResponseMutualFundWorldRisk) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseMutualFundWorldRisk) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseMutualFundWorldRisk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseMutualFundWorldRisk(val *ResponseMutualFundWorldRisk) *NullableResponseMutualFundWorldRisk {
	return &NullableResponseMutualFundWorldRisk{value: val, isSet: true}
}

func (v NullableResponseMutualFundWorldRisk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseMutualFundWorldRisk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
