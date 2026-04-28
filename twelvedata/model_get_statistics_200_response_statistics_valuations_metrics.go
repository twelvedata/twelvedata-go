/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the GetStatistics200ResponseStatisticsValuationsMetrics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetStatistics200ResponseStatisticsValuationsMetrics{}

// GetStatistics200ResponseStatisticsValuationsMetrics Valuation metrics of the company
type GetStatistics200ResponseStatisticsValuationsMetrics struct {
	// Refers to the market value of the company's outstanding shares
	MarketCapitalization *float64 `json:"market_capitalization,omitempty"`
	// Refers to enterprise value (EV) of the company, often used as a more comprehensive alternative to market capitalization
	EnterpriseValue *float64 `json:"enterprise_value,omitempty"`
	// Refers to the trailing price-to-earnings (P/E). It is calculated by taking the current stock price and dividing it by the trailing earnings per share (EPS) for the past 12 months
	TrailingPe *float64 `json:"trailing_pe,omitempty"`
	// Refers to the forward price-to-earnings ratio. It is calculated by dividing the current share price by the estimated future earnings per share
	ForwardPe *float64 `json:"forward_pe,omitempty"`
	// The price/earnings to growth (PEG) ratio is a price-to-earnings ratio divided by the growth rate of the earnings
	PegRatio *float64 `json:"peg_ratio,omitempty"`
	// The price-to-sales (P/S) ratio is a valuation ratio that compares the market capitalization to its revenues over the last 12 months
	PriceToSalesTtm *float64 `json:"price_to_sales_ttm,omitempty"`
	// The price-to-book (P/B) ratio is equal to the current share price divided by the book value of all shares (BVPS) over the last quarter
	PriceToBookMrq *float64 `json:"price_to_book_mrq,omitempty"`
	// The enterprise value-to-revenue multiple (EV/R) is a measure that compares enterprise value to revenue
	EnterpriseToRevenue *float64 `json:"enterprise_to_revenue,omitempty"`
	// The enterprise value-to-ebitda multiple (EV/EBITDA) is a measure that compares enterprise value to EBITDA
	EnterpriseToEbitda *float64 `json:"enterprise_to_ebitda,omitempty"`
}

// NewGetStatistics200ResponseStatisticsValuationsMetrics instantiates a new GetStatistics200ResponseStatisticsValuationsMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStatistics200ResponseStatisticsValuationsMetrics() *GetStatistics200ResponseStatisticsValuationsMetrics {
	this := GetStatistics200ResponseStatisticsValuationsMetrics{}
	return &this
}

// NewGetStatistics200ResponseStatisticsValuationsMetricsWithDefaults instantiates a new GetStatistics200ResponseStatisticsValuationsMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStatistics200ResponseStatisticsValuationsMetricsWithDefaults() *GetStatistics200ResponseStatisticsValuationsMetrics {
	this := GetStatistics200ResponseStatisticsValuationsMetrics{}
	return &this
}

// GetMarketCapitalization returns the MarketCapitalization field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetMarketCapitalization() float64 {
	if o == nil || IsNil(o.MarketCapitalization) {
		var ret float64
		return ret
	}
	return *o.MarketCapitalization
}

// GetMarketCapitalizationOk returns a tuple with the MarketCapitalization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetMarketCapitalizationOk() (*float64, bool) {
	if o == nil || IsNil(o.MarketCapitalization) {
		return nil, false
	}
	return o.MarketCapitalization, true
}

// HasMarketCapitalization returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) HasMarketCapitalization() bool {
	if o != nil && !IsNil(o.MarketCapitalization) {
		return true
	}

	return false
}

// SetMarketCapitalization gets a reference to the given float64 and assigns it to the MarketCapitalization field.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) SetMarketCapitalization(v float64) {
	o.MarketCapitalization = &v
}

// GetEnterpriseValue returns the EnterpriseValue field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetEnterpriseValue() float64 {
	if o == nil || IsNil(o.EnterpriseValue) {
		var ret float64
		return ret
	}
	return *o.EnterpriseValue
}

// GetEnterpriseValueOk returns a tuple with the EnterpriseValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetEnterpriseValueOk() (*float64, bool) {
	if o == nil || IsNil(o.EnterpriseValue) {
		return nil, false
	}
	return o.EnterpriseValue, true
}

// HasEnterpriseValue returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) HasEnterpriseValue() bool {
	if o != nil && !IsNil(o.EnterpriseValue) {
		return true
	}

	return false
}

// SetEnterpriseValue gets a reference to the given float64 and assigns it to the EnterpriseValue field.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) SetEnterpriseValue(v float64) {
	o.EnterpriseValue = &v
}

// GetTrailingPe returns the TrailingPe field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetTrailingPe() float64 {
	if o == nil || IsNil(o.TrailingPe) {
		var ret float64
		return ret
	}
	return *o.TrailingPe
}

// GetTrailingPeOk returns a tuple with the TrailingPe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetTrailingPeOk() (*float64, bool) {
	if o == nil || IsNil(o.TrailingPe) {
		return nil, false
	}
	return o.TrailingPe, true
}

// HasTrailingPe returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) HasTrailingPe() bool {
	if o != nil && !IsNil(o.TrailingPe) {
		return true
	}

	return false
}

// SetTrailingPe gets a reference to the given float64 and assigns it to the TrailingPe field.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) SetTrailingPe(v float64) {
	o.TrailingPe = &v
}

// GetForwardPe returns the ForwardPe field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetForwardPe() float64 {
	if o == nil || IsNil(o.ForwardPe) {
		var ret float64
		return ret
	}
	return *o.ForwardPe
}

// GetForwardPeOk returns a tuple with the ForwardPe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetForwardPeOk() (*float64, bool) {
	if o == nil || IsNil(o.ForwardPe) {
		return nil, false
	}
	return o.ForwardPe, true
}

// HasForwardPe returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) HasForwardPe() bool {
	if o != nil && !IsNil(o.ForwardPe) {
		return true
	}

	return false
}

// SetForwardPe gets a reference to the given float64 and assigns it to the ForwardPe field.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) SetForwardPe(v float64) {
	o.ForwardPe = &v
}

// GetPegRatio returns the PegRatio field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetPegRatio() float64 {
	if o == nil || IsNil(o.PegRatio) {
		var ret float64
		return ret
	}
	return *o.PegRatio
}

// GetPegRatioOk returns a tuple with the PegRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetPegRatioOk() (*float64, bool) {
	if o == nil || IsNil(o.PegRatio) {
		return nil, false
	}
	return o.PegRatio, true
}

// HasPegRatio returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) HasPegRatio() bool {
	if o != nil && !IsNil(o.PegRatio) {
		return true
	}

	return false
}

// SetPegRatio gets a reference to the given float64 and assigns it to the PegRatio field.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) SetPegRatio(v float64) {
	o.PegRatio = &v
}

// GetPriceToSalesTtm returns the PriceToSalesTtm field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetPriceToSalesTtm() float64 {
	if o == nil || IsNil(o.PriceToSalesTtm) {
		var ret float64
		return ret
	}
	return *o.PriceToSalesTtm
}

// GetPriceToSalesTtmOk returns a tuple with the PriceToSalesTtm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetPriceToSalesTtmOk() (*float64, bool) {
	if o == nil || IsNil(o.PriceToSalesTtm) {
		return nil, false
	}
	return o.PriceToSalesTtm, true
}

// HasPriceToSalesTtm returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) HasPriceToSalesTtm() bool {
	if o != nil && !IsNil(o.PriceToSalesTtm) {
		return true
	}

	return false
}

// SetPriceToSalesTtm gets a reference to the given float64 and assigns it to the PriceToSalesTtm field.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) SetPriceToSalesTtm(v float64) {
	o.PriceToSalesTtm = &v
}

// GetPriceToBookMrq returns the PriceToBookMrq field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetPriceToBookMrq() float64 {
	if o == nil || IsNil(o.PriceToBookMrq) {
		var ret float64
		return ret
	}
	return *o.PriceToBookMrq
}

// GetPriceToBookMrqOk returns a tuple with the PriceToBookMrq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetPriceToBookMrqOk() (*float64, bool) {
	if o == nil || IsNil(o.PriceToBookMrq) {
		return nil, false
	}
	return o.PriceToBookMrq, true
}

// HasPriceToBookMrq returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) HasPriceToBookMrq() bool {
	if o != nil && !IsNil(o.PriceToBookMrq) {
		return true
	}

	return false
}

// SetPriceToBookMrq gets a reference to the given float64 and assigns it to the PriceToBookMrq field.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) SetPriceToBookMrq(v float64) {
	o.PriceToBookMrq = &v
}

// GetEnterpriseToRevenue returns the EnterpriseToRevenue field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetEnterpriseToRevenue() float64 {
	if o == nil || IsNil(o.EnterpriseToRevenue) {
		var ret float64
		return ret
	}
	return *o.EnterpriseToRevenue
}

// GetEnterpriseToRevenueOk returns a tuple with the EnterpriseToRevenue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetEnterpriseToRevenueOk() (*float64, bool) {
	if o == nil || IsNil(o.EnterpriseToRevenue) {
		return nil, false
	}
	return o.EnterpriseToRevenue, true
}

// HasEnterpriseToRevenue returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) HasEnterpriseToRevenue() bool {
	if o != nil && !IsNil(o.EnterpriseToRevenue) {
		return true
	}

	return false
}

// SetEnterpriseToRevenue gets a reference to the given float64 and assigns it to the EnterpriseToRevenue field.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) SetEnterpriseToRevenue(v float64) {
	o.EnterpriseToRevenue = &v
}

// GetEnterpriseToEbitda returns the EnterpriseToEbitda field value if set, zero value otherwise.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetEnterpriseToEbitda() float64 {
	if o == nil || IsNil(o.EnterpriseToEbitda) {
		var ret float64
		return ret
	}
	return *o.EnterpriseToEbitda
}

// GetEnterpriseToEbitdaOk returns a tuple with the EnterpriseToEbitda field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) GetEnterpriseToEbitdaOk() (*float64, bool) {
	if o == nil || IsNil(o.EnterpriseToEbitda) {
		return nil, false
	}
	return o.EnterpriseToEbitda, true
}

// HasEnterpriseToEbitda returns a boolean if a field has been set.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) HasEnterpriseToEbitda() bool {
	if o != nil && !IsNil(o.EnterpriseToEbitda) {
		return true
	}

	return false
}

// SetEnterpriseToEbitda gets a reference to the given float64 and assigns it to the EnterpriseToEbitda field.
func (o *GetStatistics200ResponseStatisticsValuationsMetrics) SetEnterpriseToEbitda(v float64) {
	o.EnterpriseToEbitda = &v
}

func (o GetStatistics200ResponseStatisticsValuationsMetrics) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetStatistics200ResponseStatisticsValuationsMetrics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MarketCapitalization) {
		toSerialize["market_capitalization"] = o.MarketCapitalization
	}
	if !IsNil(o.EnterpriseValue) {
		toSerialize["enterprise_value"] = o.EnterpriseValue
	}
	if !IsNil(o.TrailingPe) {
		toSerialize["trailing_pe"] = o.TrailingPe
	}
	if !IsNil(o.ForwardPe) {
		toSerialize["forward_pe"] = o.ForwardPe
	}
	if !IsNil(o.PegRatio) {
		toSerialize["peg_ratio"] = o.PegRatio
	}
	if !IsNil(o.PriceToSalesTtm) {
		toSerialize["price_to_sales_ttm"] = o.PriceToSalesTtm
	}
	if !IsNil(o.PriceToBookMrq) {
		toSerialize["price_to_book_mrq"] = o.PriceToBookMrq
	}
	if !IsNil(o.EnterpriseToRevenue) {
		toSerialize["enterprise_to_revenue"] = o.EnterpriseToRevenue
	}
	if !IsNil(o.EnterpriseToEbitda) {
		toSerialize["enterprise_to_ebitda"] = o.EnterpriseToEbitda
	}
	return toSerialize, nil
}

type NullableGetStatistics200ResponseStatisticsValuationsMetrics struct {
	value *GetStatistics200ResponseStatisticsValuationsMetrics
	isSet bool
}

func (v NullableGetStatistics200ResponseStatisticsValuationsMetrics) Get() *GetStatistics200ResponseStatisticsValuationsMetrics {
	return v.value
}

func (v *NullableGetStatistics200ResponseStatisticsValuationsMetrics) Set(val *GetStatistics200ResponseStatisticsValuationsMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStatistics200ResponseStatisticsValuationsMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStatistics200ResponseStatisticsValuationsMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStatistics200ResponseStatisticsValuationsMetrics(val *GetStatistics200ResponseStatisticsValuationsMetrics) *NullableGetStatistics200ResponseStatisticsValuationsMetrics {
	return &NullableGetStatistics200ResponseStatisticsValuationsMetrics{value: val, isSet: true}
}

func (v NullableGetStatistics200ResponseStatisticsValuationsMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStatistics200ResponseStatisticsValuationsMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
