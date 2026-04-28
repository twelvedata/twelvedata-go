/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetEpsTrend200ResponseEpsTrendInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEpsTrend200ResponseEpsTrendInner{}

// GetEpsTrend200ResponseEpsTrendInner struct for GetEpsTrend200ResponseEpsTrendInner
type GetEpsTrend200ResponseEpsTrendInner struct {
	// Date of the estimation
	Date string `json:"date"`
	// Period of estimation, can be `current_quarter`, `next_quarter`, `current_year`, or `next_year`
	Period string `json:"period"`
	// Actual EPS estimation for the period
	CurrentEstimate *float64 `json:"current_estimate,omitempty"`
	// EPS estimation value 7 days ago
	Var7DaysAgo *float64 `json:"7_days_ago,omitempty"`
	// EPS estimation value 30 days ago
	Var30DaysAgo *float64 `json:"30_days_ago,omitempty"`
	// EPS estimation value 60 days ago
	Var60DaysAgo *float64 `json:"60_days_ago,omitempty"`
	// EPS estimation value 90 days ago
	Var90DaysAgo *float64 `json:"90_days_ago,omitempty"`
}

type _GetEpsTrend200ResponseEpsTrendInner GetEpsTrend200ResponseEpsTrendInner

// NewGetEpsTrend200ResponseEpsTrendInner instantiates a new GetEpsTrend200ResponseEpsTrendInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEpsTrend200ResponseEpsTrendInner(date string, period string) *GetEpsTrend200ResponseEpsTrendInner {
	this := GetEpsTrend200ResponseEpsTrendInner{}
	this.Date = date
	this.Period = period
	return &this
}

// NewGetEpsTrend200ResponseEpsTrendInnerWithDefaults instantiates a new GetEpsTrend200ResponseEpsTrendInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEpsTrend200ResponseEpsTrendInnerWithDefaults() *GetEpsTrend200ResponseEpsTrendInner {
	this := GetEpsTrend200ResponseEpsTrendInner{}
	return &this
}

// GetDate returns the Date field value
func (o *GetEpsTrend200ResponseEpsTrendInner) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *GetEpsTrend200ResponseEpsTrendInner) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *GetEpsTrend200ResponseEpsTrendInner) SetDate(v string) {
	o.Date = v
}

// GetPeriod returns the Period field value
func (o *GetEpsTrend200ResponseEpsTrendInner) GetPeriod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Period
}

// GetPeriodOk returns a tuple with the Period field value
// and a boolean to check if the value has been set.
func (o *GetEpsTrend200ResponseEpsTrendInner) GetPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Period, true
}

// SetPeriod sets field value
func (o *GetEpsTrend200ResponseEpsTrendInner) SetPeriod(v string) {
	o.Period = v
}

// GetCurrentEstimate returns the CurrentEstimate field value if set, zero value otherwise.
func (o *GetEpsTrend200ResponseEpsTrendInner) GetCurrentEstimate() float64 {
	if o == nil || IsNil(o.CurrentEstimate) {
		var ret float64
		return ret
	}
	return *o.CurrentEstimate
}

// GetCurrentEstimateOk returns a tuple with the CurrentEstimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEpsTrend200ResponseEpsTrendInner) GetCurrentEstimateOk() (*float64, bool) {
	if o == nil || IsNil(o.CurrentEstimate) {
		return nil, false
	}
	return o.CurrentEstimate, true
}

// HasCurrentEstimate returns a boolean if a field has been set.
func (o *GetEpsTrend200ResponseEpsTrendInner) HasCurrentEstimate() bool {
	if o != nil && !IsNil(o.CurrentEstimate) {
		return true
	}

	return false
}

// SetCurrentEstimate gets a reference to the given float64 and assigns it to the CurrentEstimate field.
func (o *GetEpsTrend200ResponseEpsTrendInner) SetCurrentEstimate(v float64) {
	o.CurrentEstimate = &v
}

// GetVar7DaysAgo returns the Var7DaysAgo field value if set, zero value otherwise.
func (o *GetEpsTrend200ResponseEpsTrendInner) GetVar7DaysAgo() float64 {
	if o == nil || IsNil(o.Var7DaysAgo) {
		var ret float64
		return ret
	}
	return *o.Var7DaysAgo
}

// GetVar7DaysAgoOk returns a tuple with the Var7DaysAgo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEpsTrend200ResponseEpsTrendInner) GetVar7DaysAgoOk() (*float64, bool) {
	if o == nil || IsNil(o.Var7DaysAgo) {
		return nil, false
	}
	return o.Var7DaysAgo, true
}

// HasVar7DaysAgo returns a boolean if a field has been set.
func (o *GetEpsTrend200ResponseEpsTrendInner) HasVar7DaysAgo() bool {
	if o != nil && !IsNil(o.Var7DaysAgo) {
		return true
	}

	return false
}

// SetVar7DaysAgo gets a reference to the given float64 and assigns it to the Var7DaysAgo field.
func (o *GetEpsTrend200ResponseEpsTrendInner) SetVar7DaysAgo(v float64) {
	o.Var7DaysAgo = &v
}

// GetVar30DaysAgo returns the Var30DaysAgo field value if set, zero value otherwise.
func (o *GetEpsTrend200ResponseEpsTrendInner) GetVar30DaysAgo() float64 {
	if o == nil || IsNil(o.Var30DaysAgo) {
		var ret float64
		return ret
	}
	return *o.Var30DaysAgo
}

// GetVar30DaysAgoOk returns a tuple with the Var30DaysAgo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEpsTrend200ResponseEpsTrendInner) GetVar30DaysAgoOk() (*float64, bool) {
	if o == nil || IsNil(o.Var30DaysAgo) {
		return nil, false
	}
	return o.Var30DaysAgo, true
}

// HasVar30DaysAgo returns a boolean if a field has been set.
func (o *GetEpsTrend200ResponseEpsTrendInner) HasVar30DaysAgo() bool {
	if o != nil && !IsNil(o.Var30DaysAgo) {
		return true
	}

	return false
}

// SetVar30DaysAgo gets a reference to the given float64 and assigns it to the Var30DaysAgo field.
func (o *GetEpsTrend200ResponseEpsTrendInner) SetVar30DaysAgo(v float64) {
	o.Var30DaysAgo = &v
}

// GetVar60DaysAgo returns the Var60DaysAgo field value if set, zero value otherwise.
func (o *GetEpsTrend200ResponseEpsTrendInner) GetVar60DaysAgo() float64 {
	if o == nil || IsNil(o.Var60DaysAgo) {
		var ret float64
		return ret
	}
	return *o.Var60DaysAgo
}

// GetVar60DaysAgoOk returns a tuple with the Var60DaysAgo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEpsTrend200ResponseEpsTrendInner) GetVar60DaysAgoOk() (*float64, bool) {
	if o == nil || IsNil(o.Var60DaysAgo) {
		return nil, false
	}
	return o.Var60DaysAgo, true
}

// HasVar60DaysAgo returns a boolean if a field has been set.
func (o *GetEpsTrend200ResponseEpsTrendInner) HasVar60DaysAgo() bool {
	if o != nil && !IsNil(o.Var60DaysAgo) {
		return true
	}

	return false
}

// SetVar60DaysAgo gets a reference to the given float64 and assigns it to the Var60DaysAgo field.
func (o *GetEpsTrend200ResponseEpsTrendInner) SetVar60DaysAgo(v float64) {
	o.Var60DaysAgo = &v
}

// GetVar90DaysAgo returns the Var90DaysAgo field value if set, zero value otherwise.
func (o *GetEpsTrend200ResponseEpsTrendInner) GetVar90DaysAgo() float64 {
	if o == nil || IsNil(o.Var90DaysAgo) {
		var ret float64
		return ret
	}
	return *o.Var90DaysAgo
}

// GetVar90DaysAgoOk returns a tuple with the Var90DaysAgo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEpsTrend200ResponseEpsTrendInner) GetVar90DaysAgoOk() (*float64, bool) {
	if o == nil || IsNil(o.Var90DaysAgo) {
		return nil, false
	}
	return o.Var90DaysAgo, true
}

// HasVar90DaysAgo returns a boolean if a field has been set.
func (o *GetEpsTrend200ResponseEpsTrendInner) HasVar90DaysAgo() bool {
	if o != nil && !IsNil(o.Var90DaysAgo) {
		return true
	}

	return false
}

// SetVar90DaysAgo gets a reference to the given float64 and assigns it to the Var90DaysAgo field.
func (o *GetEpsTrend200ResponseEpsTrendInner) SetVar90DaysAgo(v float64) {
	o.Var90DaysAgo = &v
}

func (o GetEpsTrend200ResponseEpsTrendInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEpsTrend200ResponseEpsTrendInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["date"] = o.Date
	toSerialize["period"] = o.Period
	if !IsNil(o.CurrentEstimate) {
		toSerialize["current_estimate"] = o.CurrentEstimate
	}
	if !IsNil(o.Var7DaysAgo) {
		toSerialize["7_days_ago"] = o.Var7DaysAgo
	}
	if !IsNil(o.Var30DaysAgo) {
		toSerialize["30_days_ago"] = o.Var30DaysAgo
	}
	if !IsNil(o.Var60DaysAgo) {
		toSerialize["60_days_ago"] = o.Var60DaysAgo
	}
	if !IsNil(o.Var90DaysAgo) {
		toSerialize["90_days_ago"] = o.Var90DaysAgo
	}
	return toSerialize, nil
}

func (o *GetEpsTrend200ResponseEpsTrendInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"date",
		"period",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGetEpsTrend200ResponseEpsTrendInner := _GetEpsTrend200ResponseEpsTrendInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetEpsTrend200ResponseEpsTrendInner)

	if err != nil {
		return err
	}

	*o = GetEpsTrend200ResponseEpsTrendInner(varGetEpsTrend200ResponseEpsTrendInner)

	return err
}

type NullableGetEpsTrend200ResponseEpsTrendInner struct {
	value *GetEpsTrend200ResponseEpsTrendInner
	isSet bool
}

func (v NullableGetEpsTrend200ResponseEpsTrendInner) Get() *GetEpsTrend200ResponseEpsTrendInner {
	return v.value
}

func (v *NullableGetEpsTrend200ResponseEpsTrendInner) Set(val *GetEpsTrend200ResponseEpsTrendInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEpsTrend200ResponseEpsTrendInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEpsTrend200ResponseEpsTrendInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEpsTrend200ResponseEpsTrendInner(val *GetEpsTrend200ResponseEpsTrendInner) *NullableGetEpsTrend200ResponseEpsTrendInner {
	return &NullableGetEpsTrend200ResponseEpsTrendInner{value: val, isSet: true}
}

func (v NullableGetEpsTrend200ResponseEpsTrendInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEpsTrend200ResponseEpsTrendInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
