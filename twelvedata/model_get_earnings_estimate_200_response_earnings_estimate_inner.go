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

// checks if the GetEarningsEstimate200ResponseEarningsEstimateInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEarningsEstimate200ResponseEarningsEstimateInner{}

// GetEarningsEstimate200ResponseEarningsEstimateInner struct for GetEarningsEstimate200ResponseEarningsEstimateInner
type GetEarningsEstimate200ResponseEarningsEstimateInner struct {
	// Date of the earnings estimate
	Date string `json:"date"`
	// Period of estimation, can be `current_quarter`, `next_quarter`, `current_year`, or `next_year`
	Period string `json:"period"`
	// Number of analysts that made the estimation
	NumberOfAnalysts *int64 `json:"number_of_analysts,omitempty"`
	// Average estimation across analysts
	AvgEstimate *float64 `json:"avg_estimate,omitempty"`
	// Lowest estimation given by an analyst
	LowEstimate *float64 `json:"low_estimate,omitempty"`
	// Highest estimation given by an analyst
	HighEstimate *float64 `json:"high_estimate,omitempty"`
	// Average estimation of this period's earnings given a year ago
	YearAgoEps *float64 `json:"year_ago_eps,omitempty"`
}

type _GetEarningsEstimate200ResponseEarningsEstimateInner GetEarningsEstimate200ResponseEarningsEstimateInner

// NewGetEarningsEstimate200ResponseEarningsEstimateInner instantiates a new GetEarningsEstimate200ResponseEarningsEstimateInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEarningsEstimate200ResponseEarningsEstimateInner(date string, period string) *GetEarningsEstimate200ResponseEarningsEstimateInner {
	this := GetEarningsEstimate200ResponseEarningsEstimateInner{}
	this.Date = date
	this.Period = period
	return &this
}

// NewGetEarningsEstimate200ResponseEarningsEstimateInnerWithDefaults instantiates a new GetEarningsEstimate200ResponseEarningsEstimateInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEarningsEstimate200ResponseEarningsEstimateInnerWithDefaults() *GetEarningsEstimate200ResponseEarningsEstimateInner {
	this := GetEarningsEstimate200ResponseEarningsEstimateInner{}
	return &this
}

// GetDate returns the Date field value
func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) SetDate(v string) {
	o.Date = v
}

// GetPeriod returns the Period field value
func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) GetPeriod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Period
}

// GetPeriodOk returns a tuple with the Period field value
// and a boolean to check if the value has been set.
func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) GetPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Period, true
}

// SetPeriod sets field value
func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) SetPeriod(v string) {
	o.Period = v
}

// GetNumberOfAnalysts returns the NumberOfAnalysts field value if set, zero value otherwise.
func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) GetNumberOfAnalysts() int64 {
	if o == nil || IsNil(o.NumberOfAnalysts) {
		var ret int64
		return ret
	}
	return *o.NumberOfAnalysts
}

// GetNumberOfAnalystsOk returns a tuple with the NumberOfAnalysts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) GetNumberOfAnalystsOk() (*int64, bool) {
	if o == nil || IsNil(o.NumberOfAnalysts) {
		return nil, false
	}
	return o.NumberOfAnalysts, true
}

// HasNumberOfAnalysts returns a boolean if a field has been set.
func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) HasNumberOfAnalysts() bool {
	if o != nil && !IsNil(o.NumberOfAnalysts) {
		return true
	}

	return false
}

// SetNumberOfAnalysts gets a reference to the given int64 and assigns it to the NumberOfAnalysts field.
func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) SetNumberOfAnalysts(v int64) {
	o.NumberOfAnalysts = &v
}

// GetAvgEstimate returns the AvgEstimate field value if set, zero value otherwise.
func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) GetAvgEstimate() float64 {
	if o == nil || IsNil(o.AvgEstimate) {
		var ret float64
		return ret
	}
	return *o.AvgEstimate
}

// GetAvgEstimateOk returns a tuple with the AvgEstimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) GetAvgEstimateOk() (*float64, bool) {
	if o == nil || IsNil(o.AvgEstimate) {
		return nil, false
	}
	return o.AvgEstimate, true
}

// HasAvgEstimate returns a boolean if a field has been set.
func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) HasAvgEstimate() bool {
	if o != nil && !IsNil(o.AvgEstimate) {
		return true
	}

	return false
}

// SetAvgEstimate gets a reference to the given float64 and assigns it to the AvgEstimate field.
func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) SetAvgEstimate(v float64) {
	o.AvgEstimate = &v
}

// GetLowEstimate returns the LowEstimate field value if set, zero value otherwise.
func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) GetLowEstimate() float64 {
	if o == nil || IsNil(o.LowEstimate) {
		var ret float64
		return ret
	}
	return *o.LowEstimate
}

// GetLowEstimateOk returns a tuple with the LowEstimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) GetLowEstimateOk() (*float64, bool) {
	if o == nil || IsNil(o.LowEstimate) {
		return nil, false
	}
	return o.LowEstimate, true
}

// HasLowEstimate returns a boolean if a field has been set.
func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) HasLowEstimate() bool {
	if o != nil && !IsNil(o.LowEstimate) {
		return true
	}

	return false
}

// SetLowEstimate gets a reference to the given float64 and assigns it to the LowEstimate field.
func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) SetLowEstimate(v float64) {
	o.LowEstimate = &v
}

// GetHighEstimate returns the HighEstimate field value if set, zero value otherwise.
func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) GetHighEstimate() float64 {
	if o == nil || IsNil(o.HighEstimate) {
		var ret float64
		return ret
	}
	return *o.HighEstimate
}

// GetHighEstimateOk returns a tuple with the HighEstimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) GetHighEstimateOk() (*float64, bool) {
	if o == nil || IsNil(o.HighEstimate) {
		return nil, false
	}
	return o.HighEstimate, true
}

// HasHighEstimate returns a boolean if a field has been set.
func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) HasHighEstimate() bool {
	if o != nil && !IsNil(o.HighEstimate) {
		return true
	}

	return false
}

// SetHighEstimate gets a reference to the given float64 and assigns it to the HighEstimate field.
func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) SetHighEstimate(v float64) {
	o.HighEstimate = &v
}

// GetYearAgoEps returns the YearAgoEps field value if set, zero value otherwise.
func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) GetYearAgoEps() float64 {
	if o == nil || IsNil(o.YearAgoEps) {
		var ret float64
		return ret
	}
	return *o.YearAgoEps
}

// GetYearAgoEpsOk returns a tuple with the YearAgoEps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) GetYearAgoEpsOk() (*float64, bool) {
	if o == nil || IsNil(o.YearAgoEps) {
		return nil, false
	}
	return o.YearAgoEps, true
}

// HasYearAgoEps returns a boolean if a field has been set.
func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) HasYearAgoEps() bool {
	if o != nil && !IsNil(o.YearAgoEps) {
		return true
	}

	return false
}

// SetYearAgoEps gets a reference to the given float64 and assigns it to the YearAgoEps field.
func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) SetYearAgoEps(v float64) {
	o.YearAgoEps = &v
}

func (o GetEarningsEstimate200ResponseEarningsEstimateInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEarningsEstimate200ResponseEarningsEstimateInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["date"] = o.Date
	toSerialize["period"] = o.Period
	if !IsNil(o.NumberOfAnalysts) {
		toSerialize["number_of_analysts"] = o.NumberOfAnalysts
	}
	if !IsNil(o.AvgEstimate) {
		toSerialize["avg_estimate"] = o.AvgEstimate
	}
	if !IsNil(o.LowEstimate) {
		toSerialize["low_estimate"] = o.LowEstimate
	}
	if !IsNil(o.HighEstimate) {
		toSerialize["high_estimate"] = o.HighEstimate
	}
	if !IsNil(o.YearAgoEps) {
		toSerialize["year_ago_eps"] = o.YearAgoEps
	}
	return toSerialize, nil
}

func (o *GetEarningsEstimate200ResponseEarningsEstimateInner) UnmarshalJSON(data []byte) (err error) {
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

	varGetEarningsEstimate200ResponseEarningsEstimateInner := _GetEarningsEstimate200ResponseEarningsEstimateInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varGetEarningsEstimate200ResponseEarningsEstimateInner)

	if err != nil {
		return err
	}

	*o = GetEarningsEstimate200ResponseEarningsEstimateInner(varGetEarningsEstimate200ResponseEarningsEstimateInner)

	return err
}

type NullableGetEarningsEstimate200ResponseEarningsEstimateInner struct {
	value *GetEarningsEstimate200ResponseEarningsEstimateInner
	isSet bool
}

func (v NullableGetEarningsEstimate200ResponseEarningsEstimateInner) Get() *GetEarningsEstimate200ResponseEarningsEstimateInner {
	return v.value
}

func (v *NullableGetEarningsEstimate200ResponseEarningsEstimateInner) Set(val *GetEarningsEstimate200ResponseEarningsEstimateInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEarningsEstimate200ResponseEarningsEstimateInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEarningsEstimate200ResponseEarningsEstimateInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEarningsEstimate200ResponseEarningsEstimateInner(val *GetEarningsEstimate200ResponseEarningsEstimateInner) *NullableGetEarningsEstimate200ResponseEarningsEstimateInner {
	return &NullableGetEarningsEstimate200ResponseEarningsEstimateInner{value: val, isSet: true}
}

func (v NullableGetEarningsEstimate200ResponseEarningsEstimateInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEarningsEstimate200ResponseEarningsEstimateInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
