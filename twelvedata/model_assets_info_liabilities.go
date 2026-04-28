/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the AssetsInfoLiabilities type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetsInfoLiabilities{}

// AssetsInfoLiabilities Liabilities information
type AssetsInfoLiabilities struct {
	// Total liabilities net minority interest
	TotalLiabilitiesNetMinorityInterest *float64                                    `json:"total_liabilities_net_minority_interest,omitempty"`
	CurrentLiabilities                  *AssetsInfoLiabilitiesCurrentLiabilities    `json:"current_liabilities,omitempty"`
	NonCurrentLiabilities               *AssetsInfoLiabilitiesNonCurrentLiabilities `json:"non_current_liabilities,omitempty"`
	Equity                              *EquityInfo                                 `json:"equity,omitempty"`
}

// NewAssetsInfoLiabilities instantiates a new AssetsInfoLiabilities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetsInfoLiabilities() *AssetsInfoLiabilities {
	this := AssetsInfoLiabilities{}
	return &this
}

// NewAssetsInfoLiabilitiesWithDefaults instantiates a new AssetsInfoLiabilities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetsInfoLiabilitiesWithDefaults() *AssetsInfoLiabilities {
	this := AssetsInfoLiabilities{}
	return &this
}

// GetTotalLiabilitiesNetMinorityInterest returns the TotalLiabilitiesNetMinorityInterest field value if set, zero value otherwise.
func (o *AssetsInfoLiabilities) GetTotalLiabilitiesNetMinorityInterest() float64 {
	if o == nil || IsNil(o.TotalLiabilitiesNetMinorityInterest) {
		var ret float64
		return ret
	}
	return *o.TotalLiabilitiesNetMinorityInterest
}

// GetTotalLiabilitiesNetMinorityInterestOk returns a tuple with the TotalLiabilitiesNetMinorityInterest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilities) GetTotalLiabilitiesNetMinorityInterestOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalLiabilitiesNetMinorityInterest) {
		return nil, false
	}
	return o.TotalLiabilitiesNetMinorityInterest, true
}

// HasTotalLiabilitiesNetMinorityInterest returns a boolean if a field has been set.
func (o *AssetsInfoLiabilities) HasTotalLiabilitiesNetMinorityInterest() bool {
	if o != nil && !IsNil(o.TotalLiabilitiesNetMinorityInterest) {
		return true
	}

	return false
}

// SetTotalLiabilitiesNetMinorityInterest gets a reference to the given float64 and assigns it to the TotalLiabilitiesNetMinorityInterest field.
func (o *AssetsInfoLiabilities) SetTotalLiabilitiesNetMinorityInterest(v float64) {
	o.TotalLiabilitiesNetMinorityInterest = &v
}

// GetCurrentLiabilities returns the CurrentLiabilities field value if set, zero value otherwise.
func (o *AssetsInfoLiabilities) GetCurrentLiabilities() AssetsInfoLiabilitiesCurrentLiabilities {
	if o == nil || IsNil(o.CurrentLiabilities) {
		var ret AssetsInfoLiabilitiesCurrentLiabilities
		return ret
	}
	return *o.CurrentLiabilities
}

// GetCurrentLiabilitiesOk returns a tuple with the CurrentLiabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilities) GetCurrentLiabilitiesOk() (*AssetsInfoLiabilitiesCurrentLiabilities, bool) {
	if o == nil || IsNil(o.CurrentLiabilities) {
		return nil, false
	}
	return o.CurrentLiabilities, true
}

// HasCurrentLiabilities returns a boolean if a field has been set.
func (o *AssetsInfoLiabilities) HasCurrentLiabilities() bool {
	if o != nil && !IsNil(o.CurrentLiabilities) {
		return true
	}

	return false
}

// SetCurrentLiabilities gets a reference to the given AssetsInfoLiabilitiesCurrentLiabilities and assigns it to the CurrentLiabilities field.
func (o *AssetsInfoLiabilities) SetCurrentLiabilities(v AssetsInfoLiabilitiesCurrentLiabilities) {
	o.CurrentLiabilities = &v
}

// GetNonCurrentLiabilities returns the NonCurrentLiabilities field value if set, zero value otherwise.
func (o *AssetsInfoLiabilities) GetNonCurrentLiabilities() AssetsInfoLiabilitiesNonCurrentLiabilities {
	if o == nil || IsNil(o.NonCurrentLiabilities) {
		var ret AssetsInfoLiabilitiesNonCurrentLiabilities
		return ret
	}
	return *o.NonCurrentLiabilities
}

// GetNonCurrentLiabilitiesOk returns a tuple with the NonCurrentLiabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilities) GetNonCurrentLiabilitiesOk() (*AssetsInfoLiabilitiesNonCurrentLiabilities, bool) {
	if o == nil || IsNil(o.NonCurrentLiabilities) {
		return nil, false
	}
	return o.NonCurrentLiabilities, true
}

// HasNonCurrentLiabilities returns a boolean if a field has been set.
func (o *AssetsInfoLiabilities) HasNonCurrentLiabilities() bool {
	if o != nil && !IsNil(o.NonCurrentLiabilities) {
		return true
	}

	return false
}

// SetNonCurrentLiabilities gets a reference to the given AssetsInfoLiabilitiesNonCurrentLiabilities and assigns it to the NonCurrentLiabilities field.
func (o *AssetsInfoLiabilities) SetNonCurrentLiabilities(v AssetsInfoLiabilitiesNonCurrentLiabilities) {
	o.NonCurrentLiabilities = &v
}

// GetEquity returns the Equity field value if set, zero value otherwise.
func (o *AssetsInfoLiabilities) GetEquity() EquityInfo {
	if o == nil || IsNil(o.Equity) {
		var ret EquityInfo
		return ret
	}
	return *o.Equity
}

// GetEquityOk returns a tuple with the Equity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsInfoLiabilities) GetEquityOk() (*EquityInfo, bool) {
	if o == nil || IsNil(o.Equity) {
		return nil, false
	}
	return o.Equity, true
}

// HasEquity returns a boolean if a field has been set.
func (o *AssetsInfoLiabilities) HasEquity() bool {
	if o != nil && !IsNil(o.Equity) {
		return true
	}

	return false
}

// SetEquity gets a reference to the given EquityInfo and assigns it to the Equity field.
func (o *AssetsInfoLiabilities) SetEquity(v EquityInfo) {
	o.Equity = &v
}

func (o AssetsInfoLiabilities) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetsInfoLiabilities) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalLiabilitiesNetMinorityInterest) {
		toSerialize["total_liabilities_net_minority_interest"] = o.TotalLiabilitiesNetMinorityInterest
	}
	if !IsNil(o.CurrentLiabilities) {
		toSerialize["current_liabilities"] = o.CurrentLiabilities
	}
	if !IsNil(o.NonCurrentLiabilities) {
		toSerialize["non_current_liabilities"] = o.NonCurrentLiabilities
	}
	if !IsNil(o.Equity) {
		toSerialize["equity"] = o.Equity
	}
	return toSerialize, nil
}

type NullableAssetsInfoLiabilities struct {
	value *AssetsInfoLiabilities
	isSet bool
}

func (v NullableAssetsInfoLiabilities) Get() *AssetsInfoLiabilities {
	return v.value
}

func (v *NullableAssetsInfoLiabilities) Set(val *AssetsInfoLiabilities) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetsInfoLiabilities) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetsInfoLiabilities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetsInfoLiabilities(val *AssetsInfoLiabilities) *NullableAssetsInfoLiabilities {
	return &NullableAssetsInfoLiabilities{value: val, isSet: true}
}

func (v NullableAssetsInfoLiabilities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetsInfoLiabilities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
