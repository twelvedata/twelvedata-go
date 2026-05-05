// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the GetETFsWorld200ResponseEtfRisk type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetETFsWorld200ResponseEtfRisk{}

// GetETFsWorld200ResponseEtfRisk Risk metrics of a etf
type GetETFsWorld200ResponseEtfRisk struct {
	// Risk and volatility statistics of the fund and its category over different periods
	VolatilityMeasures []GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner `json:"volatility_measures,omitempty"`
	ValuationMetrics   *GetETFsWorld200ResponseEtfRiskValuationMetrics         `json:"valuation_metrics,omitempty"`
}

// NewGetETFsWorld200ResponseEtfRisk instantiates a new GetETFsWorld200ResponseEtfRisk object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetETFsWorld200ResponseEtfRisk() *GetETFsWorld200ResponseEtfRisk {
	this := GetETFsWorld200ResponseEtfRisk{}
	return &this
}

// NewGetETFsWorld200ResponseEtfRiskWithDefaults instantiates a new GetETFsWorld200ResponseEtfRisk object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetETFsWorld200ResponseEtfRiskWithDefaults() *GetETFsWorld200ResponseEtfRisk {
	this := GetETFsWorld200ResponseEtfRisk{}
	return &this
}

// GetVolatilityMeasures returns the VolatilityMeasures field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfRisk) GetVolatilityMeasures() []GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner {
	if o == nil || IsNil(o.VolatilityMeasures) {
		var ret []GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner
		return ret
	}
	return o.VolatilityMeasures
}

// GetVolatilityMeasuresOk returns a tuple with the VolatilityMeasures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfRisk) GetVolatilityMeasuresOk() ([]GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner, bool) {
	if o == nil || IsNil(o.VolatilityMeasures) {
		return nil, false
	}
	return o.VolatilityMeasures, true
}

// HasVolatilityMeasures returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfRisk) HasVolatilityMeasures() bool {
	if o != nil && !IsNil(o.VolatilityMeasures) {
		return true
	}

	return false
}

// SetVolatilityMeasures gets a reference to the given []GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner and assigns it to the VolatilityMeasures field.
func (o *GetETFsWorld200ResponseEtfRisk) SetVolatilityMeasures(v []GetETFsWorld200ResponseEtfRiskVolatilityMeasuresInner) {
	o.VolatilityMeasures = v
}

// GetValuationMetrics returns the ValuationMetrics field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfRisk) GetValuationMetrics() GetETFsWorld200ResponseEtfRiskValuationMetrics {
	if o == nil || IsNil(o.ValuationMetrics) {
		var ret GetETFsWorld200ResponseEtfRiskValuationMetrics
		return ret
	}
	return *o.ValuationMetrics
}

// GetValuationMetricsOk returns a tuple with the ValuationMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfRisk) GetValuationMetricsOk() (*GetETFsWorld200ResponseEtfRiskValuationMetrics, bool) {
	if o == nil || IsNil(o.ValuationMetrics) {
		return nil, false
	}
	return o.ValuationMetrics, true
}

// HasValuationMetrics returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfRisk) HasValuationMetrics() bool {
	if o != nil && !IsNil(o.ValuationMetrics) {
		return true
	}

	return false
}

// SetValuationMetrics gets a reference to the given GetETFsWorld200ResponseEtfRiskValuationMetrics and assigns it to the ValuationMetrics field.
func (o *GetETFsWorld200ResponseEtfRisk) SetValuationMetrics(v GetETFsWorld200ResponseEtfRiskValuationMetrics) {
	o.ValuationMetrics = &v
}

func (o GetETFsWorld200ResponseEtfRisk) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetETFsWorld200ResponseEtfRisk) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VolatilityMeasures) {
		toSerialize["volatility_measures"] = o.VolatilityMeasures
	}
	if !IsNil(o.ValuationMetrics) {
		toSerialize["valuation_metrics"] = o.ValuationMetrics
	}
	return toSerialize, nil
}

type NullableGetETFsWorld200ResponseEtfRisk struct {
	value *GetETFsWorld200ResponseEtfRisk
	isSet bool
}

func (v NullableGetETFsWorld200ResponseEtfRisk) Get() *GetETFsWorld200ResponseEtfRisk {
	return v.value
}

func (v *NullableGetETFsWorld200ResponseEtfRisk) Set(val *GetETFsWorld200ResponseEtfRisk) {
	v.value = val
	v.isSet = true
}

func (v NullableGetETFsWorld200ResponseEtfRisk) IsSet() bool {
	return v.isSet
}

func (v *NullableGetETFsWorld200ResponseEtfRisk) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetETFsWorld200ResponseEtfRisk(val *GetETFsWorld200ResponseEtfRisk) *NullableGetETFsWorld200ResponseEtfRisk {
	return &NullableGetETFsWorld200ResponseEtfRisk{value: val, isSet: true}
}

func (v NullableGetETFsWorld200ResponseEtfRisk) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetETFsWorld200ResponseEtfRisk) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
