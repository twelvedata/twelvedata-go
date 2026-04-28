/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the ResponseMutualFundWorldPerformance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseMutualFundWorldPerformance{}

// ResponseMutualFundWorldPerformance Detailed performance of a mutual fund
type ResponseMutualFundWorldPerformance struct {
	// Trailing returns of the fund
	TrailingReturns []GetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner `json:"trailing_returns,omitempty"`
	// Annual total returns of the fund
	AnnualTotalReturns []GetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner `json:"annual_total_returns,omitempty"`
	// Quarterly total returns of the fund
	QuarterlyTotalReturns []GetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner `json:"quarterly_total_returns,omitempty"`
	// Load adjusted return of the fund
	LoadAdjustedReturn []GetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInner `json:"load_adjusted_return,omitempty"`
}

// NewResponseMutualFundWorldPerformance instantiates a new ResponseMutualFundWorldPerformance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseMutualFundWorldPerformance() *ResponseMutualFundWorldPerformance {
	this := ResponseMutualFundWorldPerformance{}
	return &this
}

// NewResponseMutualFundWorldPerformanceWithDefaults instantiates a new ResponseMutualFundWorldPerformance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseMutualFundWorldPerformanceWithDefaults() *ResponseMutualFundWorldPerformance {
	this := ResponseMutualFundWorldPerformance{}
	return &this
}

// GetTrailingReturns returns the TrailingReturns field value if set, zero value otherwise.
func (o *ResponseMutualFundWorldPerformance) GetTrailingReturns() []GetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner {
	if o == nil || IsNil(o.TrailingReturns) {
		var ret []GetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner
		return ret
	}
	return o.TrailingReturns
}

// GetTrailingReturnsOk returns a tuple with the TrailingReturns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMutualFundWorldPerformance) GetTrailingReturnsOk() ([]GetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner, bool) {
	if o == nil || IsNil(o.TrailingReturns) {
		return nil, false
	}
	return o.TrailingReturns, true
}

// HasTrailingReturns returns a boolean if a field has been set.
func (o *ResponseMutualFundWorldPerformance) HasTrailingReturns() bool {
	if o != nil && !IsNil(o.TrailingReturns) {
		return true
	}

	return false
}

// SetTrailingReturns gets a reference to the given []GetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner and assigns it to the TrailingReturns field.
func (o *ResponseMutualFundWorldPerformance) SetTrailingReturns(v []GetMutualFundsWorld200ResponseMutualFundPerformanceTrailingReturnsInner) {
	o.TrailingReturns = v
}

// GetAnnualTotalReturns returns the AnnualTotalReturns field value if set, zero value otherwise.
func (o *ResponseMutualFundWorldPerformance) GetAnnualTotalReturns() []GetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner {
	if o == nil || IsNil(o.AnnualTotalReturns) {
		var ret []GetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner
		return ret
	}
	return o.AnnualTotalReturns
}

// GetAnnualTotalReturnsOk returns a tuple with the AnnualTotalReturns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMutualFundWorldPerformance) GetAnnualTotalReturnsOk() ([]GetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner, bool) {
	if o == nil || IsNil(o.AnnualTotalReturns) {
		return nil, false
	}
	return o.AnnualTotalReturns, true
}

// HasAnnualTotalReturns returns a boolean if a field has been set.
func (o *ResponseMutualFundWorldPerformance) HasAnnualTotalReturns() bool {
	if o != nil && !IsNil(o.AnnualTotalReturns) {
		return true
	}

	return false
}

// SetAnnualTotalReturns gets a reference to the given []GetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner and assigns it to the AnnualTotalReturns field.
func (o *ResponseMutualFundWorldPerformance) SetAnnualTotalReturns(v []GetMutualFundsWorld200ResponseMutualFundPerformanceAnnualTotalReturnsInner) {
	o.AnnualTotalReturns = v
}

// GetQuarterlyTotalReturns returns the QuarterlyTotalReturns field value if set, zero value otherwise.
func (o *ResponseMutualFundWorldPerformance) GetQuarterlyTotalReturns() []GetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner {
	if o == nil || IsNil(o.QuarterlyTotalReturns) {
		var ret []GetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner
		return ret
	}
	return o.QuarterlyTotalReturns
}

// GetQuarterlyTotalReturnsOk returns a tuple with the QuarterlyTotalReturns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMutualFundWorldPerformance) GetQuarterlyTotalReturnsOk() ([]GetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner, bool) {
	if o == nil || IsNil(o.QuarterlyTotalReturns) {
		return nil, false
	}
	return o.QuarterlyTotalReturns, true
}

// HasQuarterlyTotalReturns returns a boolean if a field has been set.
func (o *ResponseMutualFundWorldPerformance) HasQuarterlyTotalReturns() bool {
	if o != nil && !IsNil(o.QuarterlyTotalReturns) {
		return true
	}

	return false
}

// SetQuarterlyTotalReturns gets a reference to the given []GetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner and assigns it to the QuarterlyTotalReturns field.
func (o *ResponseMutualFundWorldPerformance) SetQuarterlyTotalReturns(v []GetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner) {
	o.QuarterlyTotalReturns = v
}

// GetLoadAdjustedReturn returns the LoadAdjustedReturn field value if set, zero value otherwise.
func (o *ResponseMutualFundWorldPerformance) GetLoadAdjustedReturn() []GetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInner {
	if o == nil || IsNil(o.LoadAdjustedReturn) {
		var ret []GetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInner
		return ret
	}
	return o.LoadAdjustedReturn
}

// GetLoadAdjustedReturnOk returns a tuple with the LoadAdjustedReturn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMutualFundWorldPerformance) GetLoadAdjustedReturnOk() ([]GetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInner, bool) {
	if o == nil || IsNil(o.LoadAdjustedReturn) {
		return nil, false
	}
	return o.LoadAdjustedReturn, true
}

// HasLoadAdjustedReturn returns a boolean if a field has been set.
func (o *ResponseMutualFundWorldPerformance) HasLoadAdjustedReturn() bool {
	if o != nil && !IsNil(o.LoadAdjustedReturn) {
		return true
	}

	return false
}

// SetLoadAdjustedReturn gets a reference to the given []GetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInner and assigns it to the LoadAdjustedReturn field.
func (o *ResponseMutualFundWorldPerformance) SetLoadAdjustedReturn(v []GetMutualFundsWorld200ResponseMutualFundPerformanceLoadAdjustedReturnInner) {
	o.LoadAdjustedReturn = v
}

func (o ResponseMutualFundWorldPerformance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseMutualFundWorldPerformance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TrailingReturns) {
		toSerialize["trailing_returns"] = o.TrailingReturns
	}
	if !IsNil(o.AnnualTotalReturns) {
		toSerialize["annual_total_returns"] = o.AnnualTotalReturns
	}
	if !IsNil(o.QuarterlyTotalReturns) {
		toSerialize["quarterly_total_returns"] = o.QuarterlyTotalReturns
	}
	if !IsNil(o.LoadAdjustedReturn) {
		toSerialize["load_adjusted_return"] = o.LoadAdjustedReturn
	}
	return toSerialize, nil
}

type NullableResponseMutualFundWorldPerformance struct {
	value *ResponseMutualFundWorldPerformance
	isSet bool
}

func (v NullableResponseMutualFundWorldPerformance) Get() *ResponseMutualFundWorldPerformance {
	return v.value
}

func (v *NullableResponseMutualFundWorldPerformance) Set(val *ResponseMutualFundWorldPerformance) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseMutualFundWorldPerformance) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseMutualFundWorldPerformance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseMutualFundWorldPerformance(val *ResponseMutualFundWorldPerformance) *NullableResponseMutualFundWorldPerformance {
	return &NullableResponseMutualFundWorldPerformance{value: val, isSet: true}
}

func (v NullableResponseMutualFundWorldPerformance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseMutualFundWorldPerformance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
