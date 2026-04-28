/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the GetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner{}

// GetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner struct for GetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner
type GetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner struct {
	// Year of a fund quarter return
	Year *int64 `json:"year,omitempty"`
	// Total return (%) of a fund in the first quarter
	Q1 *float64 `json:"q1,omitempty"`
	// Total return (%) of a fund in the second quarter
	Q2 *float64 `json:"q2,omitempty"`
	// Total return (%) of a fund in the third quarter
	Q3 *float64 `json:"q3,omitempty"`
	// Total return (%) of a fund in the fourth quarter
	Q4 *float64 `json:"q4,omitempty"`
}

// NewGetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner instantiates a new GetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner() *GetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner {
	this := GetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner{}
	return &this
}

// NewGetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInnerWithDefaults instantiates a new GetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInnerWithDefaults() *GetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner {
	this := GetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner{}
	return &this
}

// GetYear returns the Year field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner) GetYear() int64 {
	if o == nil || IsNil(o.Year) {
		var ret int64
		return ret
	}
	return *o.Year
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner) GetYearOk() (*int64, bool) {
	if o == nil || IsNil(o.Year) {
		return nil, false
	}
	return o.Year, true
}

// HasYear returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner) HasYear() bool {
	if o != nil && !IsNil(o.Year) {
		return true
	}

	return false
}

// SetYear gets a reference to the given int64 and assigns it to the Year field.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner) SetYear(v int64) {
	o.Year = &v
}

// GetQ1 returns the Q1 field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner) GetQ1() float64 {
	if o == nil || IsNil(o.Q1) {
		var ret float64
		return ret
	}
	return *o.Q1
}

// GetQ1Ok returns a tuple with the Q1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner) GetQ1Ok() (*float64, bool) {
	if o == nil || IsNil(o.Q1) {
		return nil, false
	}
	return o.Q1, true
}

// HasQ1 returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner) HasQ1() bool {
	if o != nil && !IsNil(o.Q1) {
		return true
	}

	return false
}

// SetQ1 gets a reference to the given float64 and assigns it to the Q1 field.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner) SetQ1(v float64) {
	o.Q1 = &v
}

// GetQ2 returns the Q2 field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner) GetQ2() float64 {
	if o == nil || IsNil(o.Q2) {
		var ret float64
		return ret
	}
	return *o.Q2
}

// GetQ2Ok returns a tuple with the Q2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner) GetQ2Ok() (*float64, bool) {
	if o == nil || IsNil(o.Q2) {
		return nil, false
	}
	return o.Q2, true
}

// HasQ2 returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner) HasQ2() bool {
	if o != nil && !IsNil(o.Q2) {
		return true
	}

	return false
}

// SetQ2 gets a reference to the given float64 and assigns it to the Q2 field.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner) SetQ2(v float64) {
	o.Q2 = &v
}

// GetQ3 returns the Q3 field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner) GetQ3() float64 {
	if o == nil || IsNil(o.Q3) {
		var ret float64
		return ret
	}
	return *o.Q3
}

// GetQ3Ok returns a tuple with the Q3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner) GetQ3Ok() (*float64, bool) {
	if o == nil || IsNil(o.Q3) {
		return nil, false
	}
	return o.Q3, true
}

// HasQ3 returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner) HasQ3() bool {
	if o != nil && !IsNil(o.Q3) {
		return true
	}

	return false
}

// SetQ3 gets a reference to the given float64 and assigns it to the Q3 field.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner) SetQ3(v float64) {
	o.Q3 = &v
}

// GetQ4 returns the Q4 field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner) GetQ4() float64 {
	if o == nil || IsNil(o.Q4) {
		var ret float64
		return ret
	}
	return *o.Q4
}

// GetQ4Ok returns a tuple with the Q4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner) GetQ4Ok() (*float64, bool) {
	if o == nil || IsNil(o.Q4) {
		return nil, false
	}
	return o.Q4, true
}

// HasQ4 returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner) HasQ4() bool {
	if o != nil && !IsNil(o.Q4) {
		return true
	}

	return false
}

// SetQ4 gets a reference to the given float64 and assigns it to the Q4 field.
func (o *GetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner) SetQ4(v float64) {
	o.Q4 = &v
}

func (o GetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Year) {
		toSerialize["year"] = o.Year
	}
	if !IsNil(o.Q1) {
		toSerialize["q1"] = o.Q1
	}
	if !IsNil(o.Q2) {
		toSerialize["q2"] = o.Q2
	}
	if !IsNil(o.Q3) {
		toSerialize["q3"] = o.Q3
	}
	if !IsNil(o.Q4) {
		toSerialize["q4"] = o.Q4
	}
	return toSerialize, nil
}

type NullableGetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner struct {
	value *GetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner
	isSet bool
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner) Get() *GetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner {
	return v.value
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner) Set(val *GetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner(val *GetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner) *NullableGetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner {
	return &NullableGetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner{value: val, isSet: true}
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundPerformanceQuarterlyTotalReturnsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
