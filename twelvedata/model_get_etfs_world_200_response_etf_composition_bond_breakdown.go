/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the GetETFsWorld200ResponseEtfCompositionBondBreakdown type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetETFsWorld200ResponseEtfCompositionBondBreakdown{}

// GetETFsWorld200ResponseEtfCompositionBondBreakdown Breakdown of the fund’s portfolio by bond holding characteristics
type GetETFsWorld200ResponseEtfCompositionBondBreakdown struct {
	AverageMaturity *GetETFsWorld200ResponseEtfCompositionBondBreakdownAverageMaturity `json:"average_maturity,omitempty"`
	AverageDuration *GetETFsWorld200ResponseEtfCompositionBondBreakdownAverageDuration `json:"average_duration,omitempty"`
	// Breakdown of the fund’s bond holdings by credit rating and their respective portfolio weights
	CreditQuality []GetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInner `json:"credit_quality,omitempty"`
}

// NewGetETFsWorld200ResponseEtfCompositionBondBreakdown instantiates a new GetETFsWorld200ResponseEtfCompositionBondBreakdown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetETFsWorld200ResponseEtfCompositionBondBreakdown() *GetETFsWorld200ResponseEtfCompositionBondBreakdown {
	this := GetETFsWorld200ResponseEtfCompositionBondBreakdown{}
	return &this
}

// NewGetETFsWorld200ResponseEtfCompositionBondBreakdownWithDefaults instantiates a new GetETFsWorld200ResponseEtfCompositionBondBreakdown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetETFsWorld200ResponseEtfCompositionBondBreakdownWithDefaults() *GetETFsWorld200ResponseEtfCompositionBondBreakdown {
	this := GetETFsWorld200ResponseEtfCompositionBondBreakdown{}
	return &this
}

// GetAverageMaturity returns the AverageMaturity field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfCompositionBondBreakdown) GetAverageMaturity() GetETFsWorld200ResponseEtfCompositionBondBreakdownAverageMaturity {
	if o == nil || IsNil(o.AverageMaturity) {
		var ret GetETFsWorld200ResponseEtfCompositionBondBreakdownAverageMaturity
		return ret
	}
	return *o.AverageMaturity
}

// GetAverageMaturityOk returns a tuple with the AverageMaturity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfCompositionBondBreakdown) GetAverageMaturityOk() (*GetETFsWorld200ResponseEtfCompositionBondBreakdownAverageMaturity, bool) {
	if o == nil || IsNil(o.AverageMaturity) {
		return nil, false
	}
	return o.AverageMaturity, true
}

// HasAverageMaturity returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfCompositionBondBreakdown) HasAverageMaturity() bool {
	if o != nil && !IsNil(o.AverageMaturity) {
		return true
	}

	return false
}

// SetAverageMaturity gets a reference to the given GetETFsWorld200ResponseEtfCompositionBondBreakdownAverageMaturity and assigns it to the AverageMaturity field.
func (o *GetETFsWorld200ResponseEtfCompositionBondBreakdown) SetAverageMaturity(v GetETFsWorld200ResponseEtfCompositionBondBreakdownAverageMaturity) {
	o.AverageMaturity = &v
}

// GetAverageDuration returns the AverageDuration field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfCompositionBondBreakdown) GetAverageDuration() GetETFsWorld200ResponseEtfCompositionBondBreakdownAverageDuration {
	if o == nil || IsNil(o.AverageDuration) {
		var ret GetETFsWorld200ResponseEtfCompositionBondBreakdownAverageDuration
		return ret
	}
	return *o.AverageDuration
}

// GetAverageDurationOk returns a tuple with the AverageDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfCompositionBondBreakdown) GetAverageDurationOk() (*GetETFsWorld200ResponseEtfCompositionBondBreakdownAverageDuration, bool) {
	if o == nil || IsNil(o.AverageDuration) {
		return nil, false
	}
	return o.AverageDuration, true
}

// HasAverageDuration returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfCompositionBondBreakdown) HasAverageDuration() bool {
	if o != nil && !IsNil(o.AverageDuration) {
		return true
	}

	return false
}

// SetAverageDuration gets a reference to the given GetETFsWorld200ResponseEtfCompositionBondBreakdownAverageDuration and assigns it to the AverageDuration field.
func (o *GetETFsWorld200ResponseEtfCompositionBondBreakdown) SetAverageDuration(v GetETFsWorld200ResponseEtfCompositionBondBreakdownAverageDuration) {
	o.AverageDuration = &v
}

// GetCreditQuality returns the CreditQuality field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfCompositionBondBreakdown) GetCreditQuality() []GetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInner {
	if o == nil || IsNil(o.CreditQuality) {
		var ret []GetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInner
		return ret
	}
	return o.CreditQuality
}

// GetCreditQualityOk returns a tuple with the CreditQuality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfCompositionBondBreakdown) GetCreditQualityOk() ([]GetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInner, bool) {
	if o == nil || IsNil(o.CreditQuality) {
		return nil, false
	}
	return o.CreditQuality, true
}

// HasCreditQuality returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfCompositionBondBreakdown) HasCreditQuality() bool {
	if o != nil && !IsNil(o.CreditQuality) {
		return true
	}

	return false
}

// SetCreditQuality gets a reference to the given []GetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInner and assigns it to the CreditQuality field.
func (o *GetETFsWorld200ResponseEtfCompositionBondBreakdown) SetCreditQuality(v []GetETFsWorld200ResponseEtfCompositionBondBreakdownCreditQualityInner) {
	o.CreditQuality = v
}

func (o GetETFsWorld200ResponseEtfCompositionBondBreakdown) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetETFsWorld200ResponseEtfCompositionBondBreakdown) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AverageMaturity) {
		toSerialize["average_maturity"] = o.AverageMaturity
	}
	if !IsNil(o.AverageDuration) {
		toSerialize["average_duration"] = o.AverageDuration
	}
	if !IsNil(o.CreditQuality) {
		toSerialize["credit_quality"] = o.CreditQuality
	}
	return toSerialize, nil
}

type NullableGetETFsWorld200ResponseEtfCompositionBondBreakdown struct {
	value *GetETFsWorld200ResponseEtfCompositionBondBreakdown
	isSet bool
}

func (v NullableGetETFsWorld200ResponseEtfCompositionBondBreakdown) Get() *GetETFsWorld200ResponseEtfCompositionBondBreakdown {
	return v.value
}

func (v *NullableGetETFsWorld200ResponseEtfCompositionBondBreakdown) Set(val *GetETFsWorld200ResponseEtfCompositionBondBreakdown) {
	v.value = val
	v.isSet = true
}

func (v NullableGetETFsWorld200ResponseEtfCompositionBondBreakdown) IsSet() bool {
	return v.isSet
}

func (v *NullableGetETFsWorld200ResponseEtfCompositionBondBreakdown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetETFsWorld200ResponseEtfCompositionBondBreakdown(val *GetETFsWorld200ResponseEtfCompositionBondBreakdown) *NullableGetETFsWorld200ResponseEtfCompositionBondBreakdown {
	return &NullableGetETFsWorld200ResponseEtfCompositionBondBreakdown{value: val, isSet: true}
}

func (v NullableGetETFsWorld200ResponseEtfCompositionBondBreakdown) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetETFsWorld200ResponseEtfCompositionBondBreakdown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
