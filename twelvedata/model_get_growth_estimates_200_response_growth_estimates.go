// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the GetGrowthEstimates200ResponseGrowthEstimates type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetGrowthEstimates200ResponseGrowthEstimates{}

// GetGrowthEstimates200ResponseGrowthEstimates Growth estimates data
type GetGrowthEstimates200ResponseGrowthEstimates struct {
	// Projected growth of the current quarter in percentage (%)
	CurrentQuarter *float64 `json:"current_quarter,omitempty"`
	// Projected growth of the next quarter in percentage (%)
	NextQuarter *float64 `json:"next_quarter,omitempty"`
	// Projected growth of the current year in percentage (%)
	CurrentYear *float64 `json:"current_year,omitempty"`
	// Projected growth of the next year in percentage (%)
	NextYear *float64 `json:"next_year,omitempty"`
	// Projected growth during the next 5 years in percentage (%) per annum
	Next5YearsPa *float64 `json:"next_5_years_pa,omitempty"`
	// Actual growth over the last 5 years in percentage (%) per annum
	Past5YearsPa *float64 `json:"past_5_years_pa,omitempty"`
}

// NewGetGrowthEstimates200ResponseGrowthEstimates instantiates a new GetGrowthEstimates200ResponseGrowthEstimates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetGrowthEstimates200ResponseGrowthEstimates() *GetGrowthEstimates200ResponseGrowthEstimates {
	this := GetGrowthEstimates200ResponseGrowthEstimates{}
	return &this
}

// NewGetGrowthEstimates200ResponseGrowthEstimatesWithDefaults instantiates a new GetGrowthEstimates200ResponseGrowthEstimates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetGrowthEstimates200ResponseGrowthEstimatesWithDefaults() *GetGrowthEstimates200ResponseGrowthEstimates {
	this := GetGrowthEstimates200ResponseGrowthEstimates{}
	return &this
}

// GetCurrentQuarter returns the CurrentQuarter field value if set, zero value otherwise.
func (o *GetGrowthEstimates200ResponseGrowthEstimates) GetCurrentQuarter() float64 {
	if o == nil || IsNil(o.CurrentQuarter) {
		var ret float64
		return ret
	}
	return *o.CurrentQuarter
}

// GetCurrentQuarterOk returns a tuple with the CurrentQuarter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetGrowthEstimates200ResponseGrowthEstimates) GetCurrentQuarterOk() (*float64, bool) {
	if o == nil || IsNil(o.CurrentQuarter) {
		return nil, false
	}
	return o.CurrentQuarter, true
}

// HasCurrentQuarter returns a boolean if a field has been set.
func (o *GetGrowthEstimates200ResponseGrowthEstimates) HasCurrentQuarter() bool {
	if o != nil && !IsNil(o.CurrentQuarter) {
		return true
	}

	return false
}

// SetCurrentQuarter gets a reference to the given float64 and assigns it to the CurrentQuarter field.
func (o *GetGrowthEstimates200ResponseGrowthEstimates) SetCurrentQuarter(v float64) {
	o.CurrentQuarter = &v
}

// GetNextQuarter returns the NextQuarter field value if set, zero value otherwise.
func (o *GetGrowthEstimates200ResponseGrowthEstimates) GetNextQuarter() float64 {
	if o == nil || IsNil(o.NextQuarter) {
		var ret float64
		return ret
	}
	return *o.NextQuarter
}

// GetNextQuarterOk returns a tuple with the NextQuarter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetGrowthEstimates200ResponseGrowthEstimates) GetNextQuarterOk() (*float64, bool) {
	if o == nil || IsNil(o.NextQuarter) {
		return nil, false
	}
	return o.NextQuarter, true
}

// HasNextQuarter returns a boolean if a field has been set.
func (o *GetGrowthEstimates200ResponseGrowthEstimates) HasNextQuarter() bool {
	if o != nil && !IsNil(o.NextQuarter) {
		return true
	}

	return false
}

// SetNextQuarter gets a reference to the given float64 and assigns it to the NextQuarter field.
func (o *GetGrowthEstimates200ResponseGrowthEstimates) SetNextQuarter(v float64) {
	o.NextQuarter = &v
}

// GetCurrentYear returns the CurrentYear field value if set, zero value otherwise.
func (o *GetGrowthEstimates200ResponseGrowthEstimates) GetCurrentYear() float64 {
	if o == nil || IsNil(o.CurrentYear) {
		var ret float64
		return ret
	}
	return *o.CurrentYear
}

// GetCurrentYearOk returns a tuple with the CurrentYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetGrowthEstimates200ResponseGrowthEstimates) GetCurrentYearOk() (*float64, bool) {
	if o == nil || IsNil(o.CurrentYear) {
		return nil, false
	}
	return o.CurrentYear, true
}

// HasCurrentYear returns a boolean if a field has been set.
func (o *GetGrowthEstimates200ResponseGrowthEstimates) HasCurrentYear() bool {
	if o != nil && !IsNil(o.CurrentYear) {
		return true
	}

	return false
}

// SetCurrentYear gets a reference to the given float64 and assigns it to the CurrentYear field.
func (o *GetGrowthEstimates200ResponseGrowthEstimates) SetCurrentYear(v float64) {
	o.CurrentYear = &v
}

// GetNextYear returns the NextYear field value if set, zero value otherwise.
func (o *GetGrowthEstimates200ResponseGrowthEstimates) GetNextYear() float64 {
	if o == nil || IsNil(o.NextYear) {
		var ret float64
		return ret
	}
	return *o.NextYear
}

// GetNextYearOk returns a tuple with the NextYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetGrowthEstimates200ResponseGrowthEstimates) GetNextYearOk() (*float64, bool) {
	if o == nil || IsNil(o.NextYear) {
		return nil, false
	}
	return o.NextYear, true
}

// HasNextYear returns a boolean if a field has been set.
func (o *GetGrowthEstimates200ResponseGrowthEstimates) HasNextYear() bool {
	if o != nil && !IsNil(o.NextYear) {
		return true
	}

	return false
}

// SetNextYear gets a reference to the given float64 and assigns it to the NextYear field.
func (o *GetGrowthEstimates200ResponseGrowthEstimates) SetNextYear(v float64) {
	o.NextYear = &v
}

// GetNext5YearsPa returns the Next5YearsPa field value if set, zero value otherwise.
func (o *GetGrowthEstimates200ResponseGrowthEstimates) GetNext5YearsPa() float64 {
	if o == nil || IsNil(o.Next5YearsPa) {
		var ret float64
		return ret
	}
	return *o.Next5YearsPa
}

// GetNext5YearsPaOk returns a tuple with the Next5YearsPa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetGrowthEstimates200ResponseGrowthEstimates) GetNext5YearsPaOk() (*float64, bool) {
	if o == nil || IsNil(o.Next5YearsPa) {
		return nil, false
	}
	return o.Next5YearsPa, true
}

// HasNext5YearsPa returns a boolean if a field has been set.
func (o *GetGrowthEstimates200ResponseGrowthEstimates) HasNext5YearsPa() bool {
	if o != nil && !IsNil(o.Next5YearsPa) {
		return true
	}

	return false
}

// SetNext5YearsPa gets a reference to the given float64 and assigns it to the Next5YearsPa field.
func (o *GetGrowthEstimates200ResponseGrowthEstimates) SetNext5YearsPa(v float64) {
	o.Next5YearsPa = &v
}

// GetPast5YearsPa returns the Past5YearsPa field value if set, zero value otherwise.
func (o *GetGrowthEstimates200ResponseGrowthEstimates) GetPast5YearsPa() float64 {
	if o == nil || IsNil(o.Past5YearsPa) {
		var ret float64
		return ret
	}
	return *o.Past5YearsPa
}

// GetPast5YearsPaOk returns a tuple with the Past5YearsPa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetGrowthEstimates200ResponseGrowthEstimates) GetPast5YearsPaOk() (*float64, bool) {
	if o == nil || IsNil(o.Past5YearsPa) {
		return nil, false
	}
	return o.Past5YearsPa, true
}

// HasPast5YearsPa returns a boolean if a field has been set.
func (o *GetGrowthEstimates200ResponseGrowthEstimates) HasPast5YearsPa() bool {
	if o != nil && !IsNil(o.Past5YearsPa) {
		return true
	}

	return false
}

// SetPast5YearsPa gets a reference to the given float64 and assigns it to the Past5YearsPa field.
func (o *GetGrowthEstimates200ResponseGrowthEstimates) SetPast5YearsPa(v float64) {
	o.Past5YearsPa = &v
}

func (o GetGrowthEstimates200ResponseGrowthEstimates) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetGrowthEstimates200ResponseGrowthEstimates) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CurrentQuarter) {
		toSerialize["current_quarter"] = o.CurrentQuarter
	}
	if !IsNil(o.NextQuarter) {
		toSerialize["next_quarter"] = o.NextQuarter
	}
	if !IsNil(o.CurrentYear) {
		toSerialize["current_year"] = o.CurrentYear
	}
	if !IsNil(o.NextYear) {
		toSerialize["next_year"] = o.NextYear
	}
	if !IsNil(o.Next5YearsPa) {
		toSerialize["next_5_years_pa"] = o.Next5YearsPa
	}
	if !IsNil(o.Past5YearsPa) {
		toSerialize["past_5_years_pa"] = o.Past5YearsPa
	}
	return toSerialize, nil
}

type NullableGetGrowthEstimates200ResponseGrowthEstimates struct {
	value *GetGrowthEstimates200ResponseGrowthEstimates
	isSet bool
}

func (v NullableGetGrowthEstimates200ResponseGrowthEstimates) Get() *GetGrowthEstimates200ResponseGrowthEstimates {
	return v.value
}

func (v *NullableGetGrowthEstimates200ResponseGrowthEstimates) Set(val *GetGrowthEstimates200ResponseGrowthEstimates) {
	v.value = val
	v.isSet = true
}

func (v NullableGetGrowthEstimates200ResponseGrowthEstimates) IsSet() bool {
	return v.isSet
}

func (v *NullableGetGrowthEstimates200ResponseGrowthEstimates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetGrowthEstimates200ResponseGrowthEstimates(val *GetGrowthEstimates200ResponseGrowthEstimates) *NullableGetGrowthEstimates200ResponseGrowthEstimates {
	return &NullableGetGrowthEstimates200ResponseGrowthEstimates{value: val, isSet: true}
}

func (v NullableGetGrowthEstimates200ResponseGrowthEstimates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetGrowthEstimates200ResponseGrowthEstimates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
