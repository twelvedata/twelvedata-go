// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the GetStatistics200ResponseStatisticsStockPriceSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetStatistics200ResponseStatisticsStockPriceSummary{}

// GetStatistics200ResponseStatisticsStockPriceSummary Stock price summary of the company
type GetStatistics200ResponseStatisticsStockPriceSummary struct {
	// Refers to the lowest price at which stock traded during a year
	FiftyTwoWeekLow *float64 `json:"fifty_two_week_low,omitempty"`
	// Refers to the highest price at which stock traded during a year
	FiftyTwoWeekHigh *float64 `json:"fifty_two_week_high,omitempty"`
	// Refers to the change between lowest and highest prices during a year
	FiftyTwoWeekChange *float64 `json:"fifty_two_week_change,omitempty"`
	// Refers to beta measure relative to the primary benchmark (index) of the country
	Beta *float64 `json:"beta,omitempty"`
	// Refers to the 50-day simple moving average
	Day50Ma *float64 `json:"day_50_ma,omitempty"`
	// Refers to the 200-day simple moving average
	Day200Ma *float64 `json:"day_200_ma,omitempty"`
}

// NewGetStatistics200ResponseStatisticsStockPriceSummary instantiates a new GetStatistics200ResponseStatisticsStockPriceSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStatistics200ResponseStatisticsStockPriceSummary() *GetStatistics200ResponseStatisticsStockPriceSummary {
	this := GetStatistics200ResponseStatisticsStockPriceSummary{}
	return &this
}

// NewGetStatistics200ResponseStatisticsStockPriceSummaryWithDefaults instantiates a new GetStatistics200ResponseStatisticsStockPriceSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStatistics200ResponseStatisticsStockPriceSummaryWithDefaults() *GetStatistics200ResponseStatisticsStockPriceSummary {
	this := GetStatistics200ResponseStatisticsStockPriceSummary{}
	return &this
}

// GetFiftyTwoWeekLow returns the FiftyTwoWeekLow field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsStockPriceSummary) GetFiftyTwoWeekLow() float64 {
	if o == nil || IsNil(o.FiftyTwoWeekLow) {
		var ret float64
		return ret
	}
	return *o.FiftyTwoWeekLow
}

// GetFiftyTwoWeekLowOk returns a tuple with the FiftyTwoWeekLow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsStockPriceSummary) GetFiftyTwoWeekLowOk() (*float64, bool) {
	if o == nil || IsNil(o.FiftyTwoWeekLow) {
		return nil, false
	}
	return o.FiftyTwoWeekLow, true
}

// HasFiftyTwoWeekLow returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsStockPriceSummary) HasFiftyTwoWeekLow() bool {
	if o != nil && !IsNil(o.FiftyTwoWeekLow) {
		return true
	}

	return false
}

// SetFiftyTwoWeekLow gets a reference to the given float64 and assigns it to the FiftyTwoWeekLow field.
func (o *GetStatistics200ResponseStatisticsStockPriceSummary) SetFiftyTwoWeekLow(v float64) {
	o.FiftyTwoWeekLow = &v
}

// GetFiftyTwoWeekHigh returns the FiftyTwoWeekHigh field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsStockPriceSummary) GetFiftyTwoWeekHigh() float64 {
	if o == nil || IsNil(o.FiftyTwoWeekHigh) {
		var ret float64
		return ret
	}
	return *o.FiftyTwoWeekHigh
}

// GetFiftyTwoWeekHighOk returns a tuple with the FiftyTwoWeekHigh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsStockPriceSummary) GetFiftyTwoWeekHighOk() (*float64, bool) {
	if o == nil || IsNil(o.FiftyTwoWeekHigh) {
		return nil, false
	}
	return o.FiftyTwoWeekHigh, true
}

// HasFiftyTwoWeekHigh returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsStockPriceSummary) HasFiftyTwoWeekHigh() bool {
	if o != nil && !IsNil(o.FiftyTwoWeekHigh) {
		return true
	}

	return false
}

// SetFiftyTwoWeekHigh gets a reference to the given float64 and assigns it to the FiftyTwoWeekHigh field.
func (o *GetStatistics200ResponseStatisticsStockPriceSummary) SetFiftyTwoWeekHigh(v float64) {
	o.FiftyTwoWeekHigh = &v
}

// GetFiftyTwoWeekChange returns the FiftyTwoWeekChange field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsStockPriceSummary) GetFiftyTwoWeekChange() float64 {
	if o == nil || IsNil(o.FiftyTwoWeekChange) {
		var ret float64
		return ret
	}
	return *o.FiftyTwoWeekChange
}

// GetFiftyTwoWeekChangeOk returns a tuple with the FiftyTwoWeekChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsStockPriceSummary) GetFiftyTwoWeekChangeOk() (*float64, bool) {
	if o == nil || IsNil(o.FiftyTwoWeekChange) {
		return nil, false
	}
	return o.FiftyTwoWeekChange, true
}

// HasFiftyTwoWeekChange returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsStockPriceSummary) HasFiftyTwoWeekChange() bool {
	if o != nil && !IsNil(o.FiftyTwoWeekChange) {
		return true
	}

	return false
}

// SetFiftyTwoWeekChange gets a reference to the given float64 and assigns it to the FiftyTwoWeekChange field.
func (o *GetStatistics200ResponseStatisticsStockPriceSummary) SetFiftyTwoWeekChange(v float64) {
	o.FiftyTwoWeekChange = &v
}

// GetBeta returns the Beta field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsStockPriceSummary) GetBeta() float64 {
	if o == nil || IsNil(o.Beta) {
		var ret float64
		return ret
	}
	return *o.Beta
}

// GetBetaOk returns a tuple with the Beta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsStockPriceSummary) GetBetaOk() (*float64, bool) {
	if o == nil || IsNil(o.Beta) {
		return nil, false
	}
	return o.Beta, true
}

// HasBeta returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsStockPriceSummary) HasBeta() bool {
	if o != nil && !IsNil(o.Beta) {
		return true
	}

	return false
}

// SetBeta gets a reference to the given float64 and assigns it to the Beta field.
func (o *GetStatistics200ResponseStatisticsStockPriceSummary) SetBeta(v float64) {
	o.Beta = &v
}

// GetDay50Ma returns the Day50Ma field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsStockPriceSummary) GetDay50Ma() float64 {
	if o == nil || IsNil(o.Day50Ma) {
		var ret float64
		return ret
	}
	return *o.Day50Ma
}

// GetDay50MaOk returns a tuple with the Day50Ma field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsStockPriceSummary) GetDay50MaOk() (*float64, bool) {
	if o == nil || IsNil(o.Day50Ma) {
		return nil, false
	}
	return o.Day50Ma, true
}

// HasDay50Ma returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsStockPriceSummary) HasDay50Ma() bool {
	if o != nil && !IsNil(o.Day50Ma) {
		return true
	}

	return false
}

// SetDay50Ma gets a reference to the given float64 and assigns it to the Day50Ma field.
func (o *GetStatistics200ResponseStatisticsStockPriceSummary) SetDay50Ma(v float64) {
	o.Day50Ma = &v
}

// GetDay200Ma returns the Day200Ma field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsStockPriceSummary) GetDay200Ma() float64 {
	if o == nil || IsNil(o.Day200Ma) {
		var ret float64
		return ret
	}
	return *o.Day200Ma
}

// GetDay200MaOk returns a tuple with the Day200Ma field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsStockPriceSummary) GetDay200MaOk() (*float64, bool) {
	if o == nil || IsNil(o.Day200Ma) {
		return nil, false
	}
	return o.Day200Ma, true
}

// HasDay200Ma returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsStockPriceSummary) HasDay200Ma() bool {
	if o != nil && !IsNil(o.Day200Ma) {
		return true
	}

	return false
}

// SetDay200Ma gets a reference to the given float64 and assigns it to the Day200Ma field.
func (o *GetStatistics200ResponseStatisticsStockPriceSummary) SetDay200Ma(v float64) {
	o.Day200Ma = &v
}

func (o GetStatistics200ResponseStatisticsStockPriceSummary) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetStatistics200ResponseStatisticsStockPriceSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FiftyTwoWeekLow) {
		toSerialize["fifty_two_week_low"] = o.FiftyTwoWeekLow
	}
	if !IsNil(o.FiftyTwoWeekHigh) {
		toSerialize["fifty_two_week_high"] = o.FiftyTwoWeekHigh
	}
	if !IsNil(o.FiftyTwoWeekChange) {
		toSerialize["fifty_two_week_change"] = o.FiftyTwoWeekChange
	}
	if !IsNil(o.Beta) {
		toSerialize["beta"] = o.Beta
	}
	if !IsNil(o.Day50Ma) {
		toSerialize["day_50_ma"] = o.Day50Ma
	}
	if !IsNil(o.Day200Ma) {
		toSerialize["day_200_ma"] = o.Day200Ma
	}
	return toSerialize, nil
}

type NullableGetStatistics200ResponseStatisticsStockPriceSummary struct {
	value *GetStatistics200ResponseStatisticsStockPriceSummary
	isSet bool
}

func (v NullableGetStatistics200ResponseStatisticsStockPriceSummary) Get() *GetStatistics200ResponseStatisticsStockPriceSummary {
	return v.value
}

func (v *NullableGetStatistics200ResponseStatisticsStockPriceSummary) Set(val *GetStatistics200ResponseStatisticsStockPriceSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStatistics200ResponseStatisticsStockPriceSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStatistics200ResponseStatisticsStockPriceSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStatistics200ResponseStatisticsStockPriceSummary(val *GetStatistics200ResponseStatisticsStockPriceSummary) *NullableGetStatistics200ResponseStatisticsStockPriceSummary {
	return &NullableGetStatistics200ResponseStatisticsStockPriceSummary{value: val, isSet: true}
}

func (v NullableGetStatistics200ResponseStatisticsStockPriceSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStatistics200ResponseStatisticsStockPriceSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
