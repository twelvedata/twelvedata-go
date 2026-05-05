// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the GetETFsWorld200ResponseEtfPerformance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetETFsWorld200ResponseEtfPerformance{}

// GetETFsWorld200ResponseEtfPerformance Detailed performance of a etf
type GetETFsWorld200ResponseEtfPerformance struct {
	// Performance returns of the fund and its category over various trailing time periods
	TrailingReturns []GetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner `json:"trailing_returns,omitempty"`
	// Fund and category total returns (%) for each calendar year
	AnnualTotalReturns []GetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner `json:"annual_total_returns,omitempty"`
}

// NewGetETFsWorld200ResponseEtfPerformance instantiates a new GetETFsWorld200ResponseEtfPerformance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetETFsWorld200ResponseEtfPerformance() *GetETFsWorld200ResponseEtfPerformance {
	this := GetETFsWorld200ResponseEtfPerformance{}
	return &this
}

// NewGetETFsWorld200ResponseEtfPerformanceWithDefaults instantiates a new GetETFsWorld200ResponseEtfPerformance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetETFsWorld200ResponseEtfPerformanceWithDefaults() *GetETFsWorld200ResponseEtfPerformance {
	this := GetETFsWorld200ResponseEtfPerformance{}
	return &this
}

// GetTrailingReturns returns the TrailingReturns field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfPerformance) GetTrailingReturns() []GetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner {
	if o == nil || IsNil(o.TrailingReturns) {
		var ret []GetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner
		return ret
	}
	return o.TrailingReturns
}

// GetTrailingReturnsOk returns a tuple with the TrailingReturns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfPerformance) GetTrailingReturnsOk() ([]GetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner, bool) {
	if o == nil || IsNil(o.TrailingReturns) {
		return nil, false
	}
	return o.TrailingReturns, true
}

// HasTrailingReturns returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfPerformance) HasTrailingReturns() bool {
	if o != nil && !IsNil(o.TrailingReturns) {
		return true
	}

	return false
}

// SetTrailingReturns gets a reference to the given []GetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner and assigns it to the TrailingReturns field.
func (o *GetETFsWorld200ResponseEtfPerformance) SetTrailingReturns(v []GetETFsWorld200ResponseEtfPerformanceTrailingReturnsInner) {
	o.TrailingReturns = v
}

// GetAnnualTotalReturns returns the AnnualTotalReturns field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfPerformance) GetAnnualTotalReturns() []GetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner {
	if o == nil || IsNil(o.AnnualTotalReturns) {
		var ret []GetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner
		return ret
	}
	return o.AnnualTotalReturns
}

// GetAnnualTotalReturnsOk returns a tuple with the AnnualTotalReturns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfPerformance) GetAnnualTotalReturnsOk() ([]GetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner, bool) {
	if o == nil || IsNil(o.AnnualTotalReturns) {
		return nil, false
	}
	return o.AnnualTotalReturns, true
}

// HasAnnualTotalReturns returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfPerformance) HasAnnualTotalReturns() bool {
	if o != nil && !IsNil(o.AnnualTotalReturns) {
		return true
	}

	return false
}

// SetAnnualTotalReturns gets a reference to the given []GetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner and assigns it to the AnnualTotalReturns field.
func (o *GetETFsWorld200ResponseEtfPerformance) SetAnnualTotalReturns(v []GetETFsWorld200ResponseEtfPerformanceAnnualTotalReturnsInner) {
	o.AnnualTotalReturns = v
}

func (o GetETFsWorld200ResponseEtfPerformance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetETFsWorld200ResponseEtfPerformance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TrailingReturns) {
		toSerialize["trailing_returns"] = o.TrailingReturns
	}
	if !IsNil(o.AnnualTotalReturns) {
		toSerialize["annual_total_returns"] = o.AnnualTotalReturns
	}
	return toSerialize, nil
}

type NullableGetETFsWorld200ResponseEtfPerformance struct {
	value *GetETFsWorld200ResponseEtfPerformance
	isSet bool
}

func (v NullableGetETFsWorld200ResponseEtfPerformance) Get() *GetETFsWorld200ResponseEtfPerformance {
	return v.value
}

func (v *NullableGetETFsWorld200ResponseEtfPerformance) Set(val *GetETFsWorld200ResponseEtfPerformance) {
	v.value = val
	v.isSet = true
}

func (v NullableGetETFsWorld200ResponseEtfPerformance) IsSet() bool {
	return v.isSet
}

func (v *NullableGetETFsWorld200ResponseEtfPerformance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetETFsWorld200ResponseEtfPerformance(val *GetETFsWorld200ResponseEtfPerformance) *NullableGetETFsWorld200ResponseEtfPerformance {
	return &NullableGetETFsWorld200ResponseEtfPerformance{value: val, isSet: true}
}

func (v NullableGetETFsWorld200ResponseEtfPerformance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetETFsWorld200ResponseEtfPerformance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
