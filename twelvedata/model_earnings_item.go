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

// checks if the EarningsItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EarningsItem{}

// EarningsItem struct for EarningsItem
type EarningsItem struct {
	// Date of earning release
	Date string `json:"date"`
	// Time of earning release, can be either of the following values: `Pre Market`, `After Hours`, `Time Not Supplied`
	Time *string `json:"time,omitempty"`
	// Analyst estimate of the future company earning
	EpsEstimate *float64 `json:"eps_estimate,omitempty"`
	// Actual value of reported earning
	EpsActual *float64 `json:"eps_actual,omitempty"`
	// Delta between `eps_actual` and `eps_estimate`
	Difference *float64 `json:"difference,omitempty"`
	// Surprise in the percentage of the `eps_actual` related to `eps_estimate`
	SurprisePrc *float64 `json:"surprise_prc,omitempty"`
}

type _EarningsItem EarningsItem

// NewEarningsItem instantiates a new EarningsItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEarningsItem(date string) *EarningsItem {
	this := EarningsItem{}
	this.Date = date
	return &this
}

// NewEarningsItemWithDefaults instantiates a new EarningsItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEarningsItemWithDefaults() *EarningsItem {
	this := EarningsItem{}
	return &this
}

// GetDate returns the Date field value
func (o *EarningsItem) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *EarningsItem) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *EarningsItem) SetDate(v string) {
	o.Date = v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *EarningsItem) GetTime() string {
	if o == nil || IsNil(o.Time) {
		var ret string
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EarningsItem) GetTimeOk() (*string, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *EarningsItem) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given string and assigns it to the Time field.
func (o *EarningsItem) SetTime(v string) {
	o.Time = &v
}

// GetEpsEstimate returns the EpsEstimate field value if set, zero value otherwise.
func (o *EarningsItem) GetEpsEstimate() float64 {
	if o == nil || IsNil(o.EpsEstimate) {
		var ret float64
		return ret
	}
	return *o.EpsEstimate
}

// GetEpsEstimateOk returns a tuple with the EpsEstimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EarningsItem) GetEpsEstimateOk() (*float64, bool) {
	if o == nil || IsNil(o.EpsEstimate) {
		return nil, false
	}
	return o.EpsEstimate, true
}

// HasEpsEstimate returns a boolean if a field has been set.
func (o *EarningsItem) HasEpsEstimate() bool {
	if o != nil && !IsNil(o.EpsEstimate) {
		return true
	}

	return false
}

// SetEpsEstimate gets a reference to the given float64 and assigns it to the EpsEstimate field.
func (o *EarningsItem) SetEpsEstimate(v float64) {
	o.EpsEstimate = &v
}

// GetEpsActual returns the EpsActual field value if set, zero value otherwise.
func (o *EarningsItem) GetEpsActual() float64 {
	if o == nil || IsNil(o.EpsActual) {
		var ret float64
		return ret
	}
	return *o.EpsActual
}

// GetEpsActualOk returns a tuple with the EpsActual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EarningsItem) GetEpsActualOk() (*float64, bool) {
	if o == nil || IsNil(o.EpsActual) {
		return nil, false
	}
	return o.EpsActual, true
}

// HasEpsActual returns a boolean if a field has been set.
func (o *EarningsItem) HasEpsActual() bool {
	if o != nil && !IsNil(o.EpsActual) {
		return true
	}

	return false
}

// SetEpsActual gets a reference to the given float64 and assigns it to the EpsActual field.
func (o *EarningsItem) SetEpsActual(v float64) {
	o.EpsActual = &v
}

// GetDifference returns the Difference field value if set, zero value otherwise.
func (o *EarningsItem) GetDifference() float64 {
	if o == nil || IsNil(o.Difference) {
		var ret float64
		return ret
	}
	return *o.Difference
}

// GetDifferenceOk returns a tuple with the Difference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EarningsItem) GetDifferenceOk() (*float64, bool) {
	if o == nil || IsNil(o.Difference) {
		return nil, false
	}
	return o.Difference, true
}

// HasDifference returns a boolean if a field has been set.
func (o *EarningsItem) HasDifference() bool {
	if o != nil && !IsNil(o.Difference) {
		return true
	}

	return false
}

// SetDifference gets a reference to the given float64 and assigns it to the Difference field.
func (o *EarningsItem) SetDifference(v float64) {
	o.Difference = &v
}

// GetSurprisePrc returns the SurprisePrc field value if set, zero value otherwise.
func (o *EarningsItem) GetSurprisePrc() float64 {
	if o == nil || IsNil(o.SurprisePrc) {
		var ret float64
		return ret
	}
	return *o.SurprisePrc
}

// GetSurprisePrcOk returns a tuple with the SurprisePrc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EarningsItem) GetSurprisePrcOk() (*float64, bool) {
	if o == nil || IsNil(o.SurprisePrc) {
		return nil, false
	}
	return o.SurprisePrc, true
}

// HasSurprisePrc returns a boolean if a field has been set.
func (o *EarningsItem) HasSurprisePrc() bool {
	if o != nil && !IsNil(o.SurprisePrc) {
		return true
	}

	return false
}

// SetSurprisePrc gets a reference to the given float64 and assigns it to the SurprisePrc field.
func (o *EarningsItem) SetSurprisePrc(v float64) {
	o.SurprisePrc = &v
}

func (o EarningsItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EarningsItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["date"] = o.Date
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.EpsEstimate) {
		toSerialize["eps_estimate"] = o.EpsEstimate
	}
	if !IsNil(o.EpsActual) {
		toSerialize["eps_actual"] = o.EpsActual
	}
	if !IsNil(o.Difference) {
		toSerialize["difference"] = o.Difference
	}
	if !IsNil(o.SurprisePrc) {
		toSerialize["surprise_prc"] = o.SurprisePrc
	}
	return toSerialize, nil
}

func (o *EarningsItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"date",
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

	varEarningsItem := _EarningsItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEarningsItem)

	if err != nil {
		return err
	}

	*o = EarningsItem(varEarningsItem)

	return err
}

type NullableEarningsItem struct {
	value *EarningsItem
	isSet bool
}

func (v NullableEarningsItem) Get() *EarningsItem {
	return v.value
}

func (v *NullableEarningsItem) Set(val *EarningsItem) {
	v.value = val
	v.isSet = true
}

func (v NullableEarningsItem) IsSet() bool {
	return v.isSet
}

func (v *NullableEarningsItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEarningsItem(val *EarningsItem) *NullableEarningsItem {
	return &NullableEarningsItem{value: val, isSet: true}
}

func (v NullableEarningsItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEarningsItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
