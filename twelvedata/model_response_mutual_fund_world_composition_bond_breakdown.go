/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the ResponseMutualFundWorldCompositionBondBreakdown type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseMutualFundWorldCompositionBondBreakdown{}

// ResponseMutualFundWorldCompositionBondBreakdown Breakdown of the fund’s bond holdings by maturity, duration, and credit quality
type ResponseMutualFundWorldCompositionBondBreakdown struct {
	AverageMaturity *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageMaturity `json:"average_maturity,omitempty"`
	// Breakdown of the fund’s bond holdings by credit rating and their respective portfolio weights
	CreditQuality   []GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInner `json:"credit_quality,omitempty"`
	AverageDuration *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageDuration     `json:"average_duration,omitempty"`
}

// NewResponseMutualFundWorldCompositionBondBreakdown instantiates a new ResponseMutualFundWorldCompositionBondBreakdown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseMutualFundWorldCompositionBondBreakdown() *ResponseMutualFundWorldCompositionBondBreakdown {
	this := ResponseMutualFundWorldCompositionBondBreakdown{}
	return &this
}

// NewResponseMutualFundWorldCompositionBondBreakdownWithDefaults instantiates a new ResponseMutualFundWorldCompositionBondBreakdown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseMutualFundWorldCompositionBondBreakdownWithDefaults() *ResponseMutualFundWorldCompositionBondBreakdown {
	this := ResponseMutualFundWorldCompositionBondBreakdown{}
	return &this
}

// GetAverageMaturity returns the AverageMaturity field value if set, zero value otherwise.
func (o *ResponseMutualFundWorldCompositionBondBreakdown) GetAverageMaturity() GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageMaturity {
	if o == nil || IsNil(o.AverageMaturity) {
		var ret GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageMaturity
		return ret
	}
	return *o.AverageMaturity
}

// GetAverageMaturityOk returns a tuple with the AverageMaturity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMutualFundWorldCompositionBondBreakdown) GetAverageMaturityOk() (*GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageMaturity, bool) {
	if o == nil || IsNil(o.AverageMaturity) {
		return nil, false
	}
	return o.AverageMaturity, true
}

// HasAverageMaturity returns a boolean if a field has been set.
func (o *ResponseMutualFundWorldCompositionBondBreakdown) HasAverageMaturity() bool {
	if o != nil && !IsNil(o.AverageMaturity) {
		return true
	}

	return false
}

// SetAverageMaturity gets a reference to the given GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageMaturity and assigns it to the AverageMaturity field.
func (o *ResponseMutualFundWorldCompositionBondBreakdown) SetAverageMaturity(v GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageMaturity) {
	o.AverageMaturity = &v
}

// GetCreditQuality returns the CreditQuality field value if set, zero value otherwise.
func (o *ResponseMutualFundWorldCompositionBondBreakdown) GetCreditQuality() []GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInner {
	if o == nil || IsNil(o.CreditQuality) {
		var ret []GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInner
		return ret
	}
	return o.CreditQuality
}

// GetCreditQualityOk returns a tuple with the CreditQuality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMutualFundWorldCompositionBondBreakdown) GetCreditQualityOk() ([]GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInner, bool) {
	if o == nil || IsNil(o.CreditQuality) {
		return nil, false
	}
	return o.CreditQuality, true
}

// HasCreditQuality returns a boolean if a field has been set.
func (o *ResponseMutualFundWorldCompositionBondBreakdown) HasCreditQuality() bool {
	if o != nil && !IsNil(o.CreditQuality) {
		return true
	}

	return false
}

// SetCreditQuality gets a reference to the given []GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInner and assigns it to the CreditQuality field.
func (o *ResponseMutualFundWorldCompositionBondBreakdown) SetCreditQuality(v []GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownCreditQualityInner) {
	o.CreditQuality = v
}

// GetAverageDuration returns the AverageDuration field value if set, zero value otherwise.
func (o *ResponseMutualFundWorldCompositionBondBreakdown) GetAverageDuration() GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageDuration {
	if o == nil || IsNil(o.AverageDuration) {
		var ret GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageDuration
		return ret
	}
	return *o.AverageDuration
}

// GetAverageDurationOk returns a tuple with the AverageDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMutualFundWorldCompositionBondBreakdown) GetAverageDurationOk() (*GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageDuration, bool) {
	if o == nil || IsNil(o.AverageDuration) {
		return nil, false
	}
	return o.AverageDuration, true
}

// HasAverageDuration returns a boolean if a field has been set.
func (o *ResponseMutualFundWorldCompositionBondBreakdown) HasAverageDuration() bool {
	if o != nil && !IsNil(o.AverageDuration) {
		return true
	}

	return false
}

// SetAverageDuration gets a reference to the given GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageDuration and assigns it to the AverageDuration field.
func (o *ResponseMutualFundWorldCompositionBondBreakdown) SetAverageDuration(v GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageDuration) {
	o.AverageDuration = &v
}

func (o ResponseMutualFundWorldCompositionBondBreakdown) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseMutualFundWorldCompositionBondBreakdown) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AverageMaturity) {
		toSerialize["average_maturity"] = o.AverageMaturity
	}
	if !IsNil(o.CreditQuality) {
		toSerialize["credit_quality"] = o.CreditQuality
	}
	if !IsNil(o.AverageDuration) {
		toSerialize["average_duration"] = o.AverageDuration
	}
	return toSerialize, nil
}

type NullableResponseMutualFundWorldCompositionBondBreakdown struct {
	value *ResponseMutualFundWorldCompositionBondBreakdown
	isSet bool
}

func (v NullableResponseMutualFundWorldCompositionBondBreakdown) Get() *ResponseMutualFundWorldCompositionBondBreakdown {
	return v.value
}

func (v *NullableResponseMutualFundWorldCompositionBondBreakdown) Set(val *ResponseMutualFundWorldCompositionBondBreakdown) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseMutualFundWorldCompositionBondBreakdown) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseMutualFundWorldCompositionBondBreakdown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseMutualFundWorldCompositionBondBreakdown(val *ResponseMutualFundWorldCompositionBondBreakdown) *NullableResponseMutualFundWorldCompositionBondBreakdown {
	return &NullableResponseMutualFundWorldCompositionBondBreakdown{value: val, isSet: true}
}

func (v NullableResponseMutualFundWorldCompositionBondBreakdown) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseMutualFundWorldCompositionBondBreakdown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
