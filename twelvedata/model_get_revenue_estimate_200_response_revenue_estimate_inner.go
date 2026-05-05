// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the GetRevenueEstimate200ResponseRevenueEstimateInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRevenueEstimate200ResponseRevenueEstimateInner{}

// GetRevenueEstimate200ResponseRevenueEstimateInner struct for GetRevenueEstimate200ResponseRevenueEstimateInner
type GetRevenueEstimate200ResponseRevenueEstimateInner struct {
	// Date of the estimate
	Date *string `json:"date,omitempty"`
	// Period of estimation, can be `current_quarter`, `next_quarter`, `current_year`, or `next_year`
	Period *string `json:"period,omitempty"`
	// Number of analysts that made the estimation
	NumberOfAnalysts *int64 `json:"number_of_analysts,omitempty"`
	// Average estimation across analysts
	AvgEstimate *float64 `json:"avg_estimate,omitempty"`
	// Lowest estimation given by an analyst
	LowEstimate *float64 `json:"low_estimate,omitempty"`
	// Highest estimation given by an analyst
	HighEstimate *float64 `json:"high_estimate,omitempty"`
	// Total revenue received a year ago relative to period
	YearAgoSales *float64 `json:"year_ago_sales,omitempty"`
	// Estimated sales growth of the period in relation to year-ago sales in prc (%)
	SalesGrowth *float64 `json:"sales_growth,omitempty"`
}

// NewGetRevenueEstimate200ResponseRevenueEstimateInner instantiates a new GetRevenueEstimate200ResponseRevenueEstimateInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRevenueEstimate200ResponseRevenueEstimateInner() *GetRevenueEstimate200ResponseRevenueEstimateInner {
	this := GetRevenueEstimate200ResponseRevenueEstimateInner{}
	return &this
}

// NewGetRevenueEstimate200ResponseRevenueEstimateInnerWithDefaults instantiates a new GetRevenueEstimate200ResponseRevenueEstimateInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRevenueEstimate200ResponseRevenueEstimateInnerWithDefaults() *GetRevenueEstimate200ResponseRevenueEstimateInner {
	this := GetRevenueEstimate200ResponseRevenueEstimateInner{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) SetDate(v string) {
	o.Date = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) GetPeriod() string {
	if o == nil || IsNil(o.Period) {
		var ret string
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) GetPeriodOk() (*string, bool) {
	if o == nil || IsNil(o.Period) {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) HasPeriod() bool {
	if o != nil && !IsNil(o.Period) {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given string and assigns it to the Period field.
func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) SetPeriod(v string) {
	o.Period = &v
}

// GetNumberOfAnalysts returns the NumberOfAnalysts field value if set, zero value otherwise.
func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) GetNumberOfAnalysts() int64 {
	if o == nil || IsNil(o.NumberOfAnalysts) {
		var ret int64
		return ret
	}
	return *o.NumberOfAnalysts
}

// GetNumberOfAnalystsOk returns a tuple with the NumberOfAnalysts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) GetNumberOfAnalystsOk() (*int64, bool) {
	if o == nil || IsNil(o.NumberOfAnalysts) {
		return nil, false
	}
	return o.NumberOfAnalysts, true
}

// HasNumberOfAnalysts returns a boolean if a field has been set.
func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) HasNumberOfAnalysts() bool {
	if o != nil && !IsNil(o.NumberOfAnalysts) {
		return true
	}

	return false
}

// SetNumberOfAnalysts gets a reference to the given int64 and assigns it to the NumberOfAnalysts field.
func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) SetNumberOfAnalysts(v int64) {
	o.NumberOfAnalysts = &v
}

// GetAvgEstimate returns the AvgEstimate field value if set, zero value otherwise.
func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) GetAvgEstimate() float64 {
	if o == nil || IsNil(o.AvgEstimate) {
		var ret float64
		return ret
	}
	return *o.AvgEstimate
}

// GetAvgEstimateOk returns a tuple with the AvgEstimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) GetAvgEstimateOk() (*float64, bool) {
	if o == nil || IsNil(o.AvgEstimate) {
		return nil, false
	}
	return o.AvgEstimate, true
}

// HasAvgEstimate returns a boolean if a field has been set.
func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) HasAvgEstimate() bool {
	if o != nil && !IsNil(o.AvgEstimate) {
		return true
	}

	return false
}

// SetAvgEstimate gets a reference to the given float64 and assigns it to the AvgEstimate field.
func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) SetAvgEstimate(v float64) {
	o.AvgEstimate = &v
}

// GetLowEstimate returns the LowEstimate field value if set, zero value otherwise.
func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) GetLowEstimate() float64 {
	if o == nil || IsNil(o.LowEstimate) {
		var ret float64
		return ret
	}
	return *o.LowEstimate
}

// GetLowEstimateOk returns a tuple with the LowEstimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) GetLowEstimateOk() (*float64, bool) {
	if o == nil || IsNil(o.LowEstimate) {
		return nil, false
	}
	return o.LowEstimate, true
}

// HasLowEstimate returns a boolean if a field has been set.
func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) HasLowEstimate() bool {
	if o != nil && !IsNil(o.LowEstimate) {
		return true
	}

	return false
}

// SetLowEstimate gets a reference to the given float64 and assigns it to the LowEstimate field.
func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) SetLowEstimate(v float64) {
	o.LowEstimate = &v
}

// GetHighEstimate returns the HighEstimate field value if set, zero value otherwise.
func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) GetHighEstimate() float64 {
	if o == nil || IsNil(o.HighEstimate) {
		var ret float64
		return ret
	}
	return *o.HighEstimate
}

// GetHighEstimateOk returns a tuple with the HighEstimate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) GetHighEstimateOk() (*float64, bool) {
	if o == nil || IsNil(o.HighEstimate) {
		return nil, false
	}
	return o.HighEstimate, true
}

// HasHighEstimate returns a boolean if a field has been set.
func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) HasHighEstimate() bool {
	if o != nil && !IsNil(o.HighEstimate) {
		return true
	}

	return false
}

// SetHighEstimate gets a reference to the given float64 and assigns it to the HighEstimate field.
func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) SetHighEstimate(v float64) {
	o.HighEstimate = &v
}

// GetYearAgoSales returns the YearAgoSales field value if set, zero value otherwise.
func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) GetYearAgoSales() float64 {
	if o == nil || IsNil(o.YearAgoSales) {
		var ret float64
		return ret
	}
	return *o.YearAgoSales
}

// GetYearAgoSalesOk returns a tuple with the YearAgoSales field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) GetYearAgoSalesOk() (*float64, bool) {
	if o == nil || IsNil(o.YearAgoSales) {
		return nil, false
	}
	return o.YearAgoSales, true
}

// HasYearAgoSales returns a boolean if a field has been set.
func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) HasYearAgoSales() bool {
	if o != nil && !IsNil(o.YearAgoSales) {
		return true
	}

	return false
}

// SetYearAgoSales gets a reference to the given float64 and assigns it to the YearAgoSales field.
func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) SetYearAgoSales(v float64) {
	o.YearAgoSales = &v
}

// GetSalesGrowth returns the SalesGrowth field value if set, zero value otherwise.
func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) GetSalesGrowth() float64 {
	if o == nil || IsNil(o.SalesGrowth) {
		var ret float64
		return ret
	}
	return *o.SalesGrowth
}

// GetSalesGrowthOk returns a tuple with the SalesGrowth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) GetSalesGrowthOk() (*float64, bool) {
	if o == nil || IsNil(o.SalesGrowth) {
		return nil, false
	}
	return o.SalesGrowth, true
}

// HasSalesGrowth returns a boolean if a field has been set.
func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) HasSalesGrowth() bool {
	if o != nil && !IsNil(o.SalesGrowth) {
		return true
	}

	return false
}

// SetSalesGrowth gets a reference to the given float64 and assigns it to the SalesGrowth field.
func (o *GetRevenueEstimate200ResponseRevenueEstimateInner) SetSalesGrowth(v float64) {
	o.SalesGrowth = &v
}

func (o GetRevenueEstimate200ResponseRevenueEstimateInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRevenueEstimate200ResponseRevenueEstimateInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.Period) {
		toSerialize["period"] = o.Period
	}
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
	if !IsNil(o.YearAgoSales) {
		toSerialize["year_ago_sales"] = o.YearAgoSales
	}
	if !IsNil(o.SalesGrowth) {
		toSerialize["sales_growth"] = o.SalesGrowth
	}
	return toSerialize, nil
}

type NullableGetRevenueEstimate200ResponseRevenueEstimateInner struct {
	value *GetRevenueEstimate200ResponseRevenueEstimateInner
	isSet bool
}

func (v NullableGetRevenueEstimate200ResponseRevenueEstimateInner) Get() *GetRevenueEstimate200ResponseRevenueEstimateInner {
	return v.value
}

func (v *NullableGetRevenueEstimate200ResponseRevenueEstimateInner) Set(val *GetRevenueEstimate200ResponseRevenueEstimateInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRevenueEstimate200ResponseRevenueEstimateInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRevenueEstimate200ResponseRevenueEstimateInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRevenueEstimate200ResponseRevenueEstimateInner(val *GetRevenueEstimate200ResponseRevenueEstimateInner) *NullableGetRevenueEstimate200ResponseRevenueEstimateInner {
	return &NullableGetRevenueEstimate200ResponseRevenueEstimateInner{value: val, isSet: true}
}

func (v NullableGetRevenueEstimate200ResponseRevenueEstimateInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRevenueEstimate200ResponseRevenueEstimateInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
