// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the GetStatistics200ResponseStatistics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetStatistics200ResponseStatistics{}

// GetStatistics200ResponseStatistics Statistics of the company
type GetStatistics200ResponseStatistics struct {
	ValuationsMetrics  *GetStatistics200ResponseStatisticsValuationsMetrics  `json:"valuations_metrics,omitempty"`
	Financials         *GetStatistics200ResponseStatisticsFinancials         `json:"financials,omitempty"`
	StockStatistics    *GetStatistics200ResponseStatisticsStockStatistics    `json:"stock_statistics,omitempty"`
	StockPriceSummary  *GetStatistics200ResponseStatisticsStockPriceSummary  `json:"stock_price_summary,omitempty"`
	DividendsAndSplits *GetStatistics200ResponseStatisticsDividendsAndSplits `json:"dividends_and_splits,omitempty"`
}

// NewGetStatistics200ResponseStatistics instantiates a new GetStatistics200ResponseStatistics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStatistics200ResponseStatistics() *GetStatistics200ResponseStatistics {
	this := GetStatistics200ResponseStatistics{}
	return &this
}

// NewGetStatistics200ResponseStatisticsWithDefaults instantiates a new GetStatistics200ResponseStatistics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStatistics200ResponseStatisticsWithDefaults() *GetStatistics200ResponseStatistics {
	this := GetStatistics200ResponseStatistics{}
	return &this
}

// GetValuationsMetrics returns the ValuationsMetrics field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatistics) GetValuationsMetrics() GetStatistics200ResponseStatisticsValuationsMetrics {
	if o == nil || IsNil(o.ValuationsMetrics) {
		var ret GetStatistics200ResponseStatisticsValuationsMetrics
		return ret
	}
	return *o.ValuationsMetrics
}

// GetValuationsMetricsOk returns a tuple with the ValuationsMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatistics) GetValuationsMetricsOk() (*GetStatistics200ResponseStatisticsValuationsMetrics, bool) {
	if o == nil || IsNil(o.ValuationsMetrics) {
		return nil, false
	}
	return o.ValuationsMetrics, true
}

// HasValuationsMetrics returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatistics) HasValuationsMetrics() bool {
	if o != nil && !IsNil(o.ValuationsMetrics) {
		return true
	}

	return false
}

// SetValuationsMetrics gets a reference to the given GetStatistics200ResponseStatisticsValuationsMetrics and assigns it to the ValuationsMetrics field.
func (o *GetStatistics200ResponseStatistics) SetValuationsMetrics(v GetStatistics200ResponseStatisticsValuationsMetrics) {
	o.ValuationsMetrics = &v
}

// GetFinancials returns the Financials field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatistics) GetFinancials() GetStatistics200ResponseStatisticsFinancials {
	if o == nil || IsNil(o.Financials) {
		var ret GetStatistics200ResponseStatisticsFinancials
		return ret
	}
	return *o.Financials
}

// GetFinancialsOk returns a tuple with the Financials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatistics) GetFinancialsOk() (*GetStatistics200ResponseStatisticsFinancials, bool) {
	if o == nil || IsNil(o.Financials) {
		return nil, false
	}
	return o.Financials, true
}

// HasFinancials returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatistics) HasFinancials() bool {
	if o != nil && !IsNil(o.Financials) {
		return true
	}

	return false
}

// SetFinancials gets a reference to the given GetStatistics200ResponseStatisticsFinancials and assigns it to the Financials field.
func (o *GetStatistics200ResponseStatistics) SetFinancials(v GetStatistics200ResponseStatisticsFinancials) {
	o.Financials = &v
}

// GetStockStatistics returns the StockStatistics field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatistics) GetStockStatistics() GetStatistics200ResponseStatisticsStockStatistics {
	if o == nil || IsNil(o.StockStatistics) {
		var ret GetStatistics200ResponseStatisticsStockStatistics
		return ret
	}
	return *o.StockStatistics
}

// GetStockStatisticsOk returns a tuple with the StockStatistics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatistics) GetStockStatisticsOk() (*GetStatistics200ResponseStatisticsStockStatistics, bool) {
	if o == nil || IsNil(o.StockStatistics) {
		return nil, false
	}
	return o.StockStatistics, true
}

// HasStockStatistics returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatistics) HasStockStatistics() bool {
	if o != nil && !IsNil(o.StockStatistics) {
		return true
	}

	return false
}

// SetStockStatistics gets a reference to the given GetStatistics200ResponseStatisticsStockStatistics and assigns it to the StockStatistics field.
func (o *GetStatistics200ResponseStatistics) SetStockStatistics(v GetStatistics200ResponseStatisticsStockStatistics) {
	o.StockStatistics = &v
}

// GetStockPriceSummary returns the StockPriceSummary field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatistics) GetStockPriceSummary() GetStatistics200ResponseStatisticsStockPriceSummary {
	if o == nil || IsNil(o.StockPriceSummary) {
		var ret GetStatistics200ResponseStatisticsStockPriceSummary
		return ret
	}
	return *o.StockPriceSummary
}

// GetStockPriceSummaryOk returns a tuple with the StockPriceSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatistics) GetStockPriceSummaryOk() (*GetStatistics200ResponseStatisticsStockPriceSummary, bool) {
	if o == nil || IsNil(o.StockPriceSummary) {
		return nil, false
	}
	return o.StockPriceSummary, true
}

// HasStockPriceSummary returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatistics) HasStockPriceSummary() bool {
	if o != nil && !IsNil(o.StockPriceSummary) {
		return true
	}

	return false
}

// SetStockPriceSummary gets a reference to the given GetStatistics200ResponseStatisticsStockPriceSummary and assigns it to the StockPriceSummary field.
func (o *GetStatistics200ResponseStatistics) SetStockPriceSummary(v GetStatistics200ResponseStatisticsStockPriceSummary) {
	o.StockPriceSummary = &v
}

// GetDividendsAndSplits returns the DividendsAndSplits field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatistics) GetDividendsAndSplits() GetStatistics200ResponseStatisticsDividendsAndSplits {
	if o == nil || IsNil(o.DividendsAndSplits) {
		var ret GetStatistics200ResponseStatisticsDividendsAndSplits
		return ret
	}
	return *o.DividendsAndSplits
}

// GetDividendsAndSplitsOk returns a tuple with the DividendsAndSplits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatistics) GetDividendsAndSplitsOk() (*GetStatistics200ResponseStatisticsDividendsAndSplits, bool) {
	if o == nil || IsNil(o.DividendsAndSplits) {
		return nil, false
	}
	return o.DividendsAndSplits, true
}

// HasDividendsAndSplits returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatistics) HasDividendsAndSplits() bool {
	if o != nil && !IsNil(o.DividendsAndSplits) {
		return true
	}

	return false
}

// SetDividendsAndSplits gets a reference to the given GetStatistics200ResponseStatisticsDividendsAndSplits and assigns it to the DividendsAndSplits field.
func (o *GetStatistics200ResponseStatistics) SetDividendsAndSplits(v GetStatistics200ResponseStatisticsDividendsAndSplits) {
	o.DividendsAndSplits = &v
}

func (o GetStatistics200ResponseStatistics) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetStatistics200ResponseStatistics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ValuationsMetrics) {
		toSerialize["valuations_metrics"] = o.ValuationsMetrics
	}
	if !IsNil(o.Financials) {
		toSerialize["financials"] = o.Financials
	}
	if !IsNil(o.StockStatistics) {
		toSerialize["stock_statistics"] = o.StockStatistics
	}
	if !IsNil(o.StockPriceSummary) {
		toSerialize["stock_price_summary"] = o.StockPriceSummary
	}
	if !IsNil(o.DividendsAndSplits) {
		toSerialize["dividends_and_splits"] = o.DividendsAndSplits
	}
	return toSerialize, nil
}

type NullableGetStatistics200ResponseStatistics struct {
	value *GetStatistics200ResponseStatistics
	isSet bool
}

func (v NullableGetStatistics200ResponseStatistics) Get() *GetStatistics200ResponseStatistics {
	return v.value
}

func (v *NullableGetStatistics200ResponseStatistics) Set(val *GetStatistics200ResponseStatistics) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStatistics200ResponseStatistics) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStatistics200ResponseStatistics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStatistics200ResponseStatistics(val *GetStatistics200ResponseStatistics) *NullableGetStatistics200ResponseStatistics {
	return &NullableGetStatistics200ResponseStatistics{value: val, isSet: true}
}

func (v NullableGetStatistics200ResponseStatistics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStatistics200ResponseStatistics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
