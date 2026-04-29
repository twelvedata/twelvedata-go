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

// checks if the GetEpsRevisions200ResponseEpsRevisionInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEpsRevisions200ResponseEpsRevisionInner{}

// GetEpsRevisions200ResponseEpsRevisionInner struct for GetEpsRevisions200ResponseEpsRevisionInner
type GetEpsRevisions200ResponseEpsRevisionInner struct {
	// Date of the EPS estimate
	Date string `json:"date"`
	// Period of estimation, can be `current_quarter`, `next_quarter`, `current_year`, or `next_year`
	Period string `json:"period"`
	// Number of up revisions over the last 7 days
	UpLastWeek *int64 `json:"up_last_week,omitempty"`
	// Number of up revisions over the last 30 days
	UpLastMonth *int64 `json:"up_last_month,omitempty"`
	// Number of down revisions over the last 7 days
	DownLastWeek *int64 `json:"down_last_week,omitempty"`
	// Number of down revisions over the last 30 days
	DownLastMonth *int64 `json:"down_last_month,omitempty"`
}

type _GetEpsRevisions200ResponseEpsRevisionInner GetEpsRevisions200ResponseEpsRevisionInner

// NewGetEpsRevisions200ResponseEpsRevisionInner instantiates a new GetEpsRevisions200ResponseEpsRevisionInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEpsRevisions200ResponseEpsRevisionInner(date string, period string) *GetEpsRevisions200ResponseEpsRevisionInner {
	this := GetEpsRevisions200ResponseEpsRevisionInner{}
	this.Date = date
	this.Period = period
	return &this
}

// NewGetEpsRevisions200ResponseEpsRevisionInnerWithDefaults instantiates a new GetEpsRevisions200ResponseEpsRevisionInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEpsRevisions200ResponseEpsRevisionInnerWithDefaults() *GetEpsRevisions200ResponseEpsRevisionInner {
	this := GetEpsRevisions200ResponseEpsRevisionInner{}
	return &this
}

// GetDate returns the Date field value
func (o *GetEpsRevisions200ResponseEpsRevisionInner) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *GetEpsRevisions200ResponseEpsRevisionInner) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *GetEpsRevisions200ResponseEpsRevisionInner) SetDate(v string) {
	o.Date = v
}

// GetPeriod returns the Period field value
func (o *GetEpsRevisions200ResponseEpsRevisionInner) GetPeriod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Period
}

// GetPeriodOk returns a tuple with the Period field value
// and a boolean to check if the value has been set.
func (o *GetEpsRevisions200ResponseEpsRevisionInner) GetPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Period, true
}

// SetPeriod sets field value
func (o *GetEpsRevisions200ResponseEpsRevisionInner) SetPeriod(v string) {
	o.Period = v
}

// GetUpLastWeek returns the UpLastWeek field value if set, zero value otherwise.
func (o *GetEpsRevisions200ResponseEpsRevisionInner) GetUpLastWeek() int64 {
	if o == nil || IsNil(o.UpLastWeek) {
		var ret int64
		return ret
	}
	return *o.UpLastWeek
}

// GetUpLastWeekOk returns a tuple with the UpLastWeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEpsRevisions200ResponseEpsRevisionInner) GetUpLastWeekOk() (*int64, bool) {
	if o == nil || IsNil(o.UpLastWeek) {
		return nil, false
	}
	return o.UpLastWeek, true
}

// HasUpLastWeek returns a boolean if a field has been set.
func (o *GetEpsRevisions200ResponseEpsRevisionInner) HasUpLastWeek() bool {
	if o != nil && !IsNil(o.UpLastWeek) {
		return true
	}

	return false
}

// SetUpLastWeek gets a reference to the given int64 and assigns it to the UpLastWeek field.
func (o *GetEpsRevisions200ResponseEpsRevisionInner) SetUpLastWeek(v int64) {
	o.UpLastWeek = &v
}

// GetUpLastMonth returns the UpLastMonth field value if set, zero value otherwise.
func (o *GetEpsRevisions200ResponseEpsRevisionInner) GetUpLastMonth() int64 {
	if o == nil || IsNil(o.UpLastMonth) {
		var ret int64
		return ret
	}
	return *o.UpLastMonth
}

// GetUpLastMonthOk returns a tuple with the UpLastMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEpsRevisions200ResponseEpsRevisionInner) GetUpLastMonthOk() (*int64, bool) {
	if o == nil || IsNil(o.UpLastMonth) {
		return nil, false
	}
	return o.UpLastMonth, true
}

// HasUpLastMonth returns a boolean if a field has been set.
func (o *GetEpsRevisions200ResponseEpsRevisionInner) HasUpLastMonth() bool {
	if o != nil && !IsNil(o.UpLastMonth) {
		return true
	}

	return false
}

// SetUpLastMonth gets a reference to the given int64 and assigns it to the UpLastMonth field.
func (o *GetEpsRevisions200ResponseEpsRevisionInner) SetUpLastMonth(v int64) {
	o.UpLastMonth = &v
}

// GetDownLastWeek returns the DownLastWeek field value if set, zero value otherwise.
func (o *GetEpsRevisions200ResponseEpsRevisionInner) GetDownLastWeek() int64 {
	if o == nil || IsNil(o.DownLastWeek) {
		var ret int64
		return ret
	}
	return *o.DownLastWeek
}

// GetDownLastWeekOk returns a tuple with the DownLastWeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEpsRevisions200ResponseEpsRevisionInner) GetDownLastWeekOk() (*int64, bool) {
	if o == nil || IsNil(o.DownLastWeek) {
		return nil, false
	}
	return o.DownLastWeek, true
}

// HasDownLastWeek returns a boolean if a field has been set.
func (o *GetEpsRevisions200ResponseEpsRevisionInner) HasDownLastWeek() bool {
	if o != nil && !IsNil(o.DownLastWeek) {
		return true
	}

	return false
}

// SetDownLastWeek gets a reference to the given int64 and assigns it to the DownLastWeek field.
func (o *GetEpsRevisions200ResponseEpsRevisionInner) SetDownLastWeek(v int64) {
	o.DownLastWeek = &v
}

// GetDownLastMonth returns the DownLastMonth field value if set, zero value otherwise.
func (o *GetEpsRevisions200ResponseEpsRevisionInner) GetDownLastMonth() int64 {
	if o == nil || IsNil(o.DownLastMonth) {
		var ret int64
		return ret
	}
	return *o.DownLastMonth
}

// GetDownLastMonthOk returns a tuple with the DownLastMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEpsRevisions200ResponseEpsRevisionInner) GetDownLastMonthOk() (*int64, bool) {
	if o == nil || IsNil(o.DownLastMonth) {
		return nil, false
	}
	return o.DownLastMonth, true
}

// HasDownLastMonth returns a boolean if a field has been set.
func (o *GetEpsRevisions200ResponseEpsRevisionInner) HasDownLastMonth() bool {
	if o != nil && !IsNil(o.DownLastMonth) {
		return true
	}

	return false
}

// SetDownLastMonth gets a reference to the given int64 and assigns it to the DownLastMonth field.
func (o *GetEpsRevisions200ResponseEpsRevisionInner) SetDownLastMonth(v int64) {
	o.DownLastMonth = &v
}

func (o GetEpsRevisions200ResponseEpsRevisionInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEpsRevisions200ResponseEpsRevisionInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["date"] = o.Date
	toSerialize["period"] = o.Period
	if !IsNil(o.UpLastWeek) {
		toSerialize["up_last_week"] = o.UpLastWeek
	}
	if !IsNil(o.UpLastMonth) {
		toSerialize["up_last_month"] = o.UpLastMonth
	}
	if !IsNil(o.DownLastWeek) {
		toSerialize["down_last_week"] = o.DownLastWeek
	}
	if !IsNil(o.DownLastMonth) {
		toSerialize["down_last_month"] = o.DownLastMonth
	}
	return toSerialize, nil
}

func (o *GetEpsRevisions200ResponseEpsRevisionInner) UnmarshalJSON(data []byte) (err error) {
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

	varGetEpsRevisions200ResponseEpsRevisionInner := _GetEpsRevisions200ResponseEpsRevisionInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetEpsRevisions200ResponseEpsRevisionInner)

	if err != nil {
		return err
	}

	*o = GetEpsRevisions200ResponseEpsRevisionInner(varGetEpsRevisions200ResponseEpsRevisionInner)

	return err
}

type NullableGetEpsRevisions200ResponseEpsRevisionInner struct {
	value *GetEpsRevisions200ResponseEpsRevisionInner
	isSet bool
}

func (v NullableGetEpsRevisions200ResponseEpsRevisionInner) Get() *GetEpsRevisions200ResponseEpsRevisionInner {
	return v.value
}

func (v *NullableGetEpsRevisions200ResponseEpsRevisionInner) Set(val *GetEpsRevisions200ResponseEpsRevisionInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEpsRevisions200ResponseEpsRevisionInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEpsRevisions200ResponseEpsRevisionInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEpsRevisions200ResponseEpsRevisionInner(val *GetEpsRevisions200ResponseEpsRevisionInner) *NullableGetEpsRevisions200ResponseEpsRevisionInner {
	return &NullableGetEpsRevisions200ResponseEpsRevisionInner{value: val, isSet: true}
}

func (v NullableGetEpsRevisions200ResponseEpsRevisionInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEpsRevisions200ResponseEpsRevisionInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
