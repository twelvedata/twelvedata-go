// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics{}

// GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics Valuation ratios and metrics of the fund and its category
type GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics struct {
	// Fund price to earnings metric
	PriceToEarnings *float64 `json:"price_to_earnings,omitempty"`
	// Average price to earnings metric of funds in the same category
	PriceToEarningsCategory *float64 `json:"price_to_earnings_category,omitempty"`
	// Fund price to book metric
	PriceToBook *float64 `json:"price_to_book,omitempty"`
	// Average price to book metric of funds in the same category
	PriceToBookCategory *float64 `json:"price_to_book_category,omitempty"`
	// Fund price to sales metric
	PriceToSales *float64 `json:"price_to_sales,omitempty"`
	// Average price to sales metric of funds in the same category
	PriceToSalesCategory *float64 `json:"price_to_sales_category,omitempty"`
	// Fund price to cashflow metric
	PriceToCashflow *float64 `json:"price_to_cashflow,omitempty"`
	// Average price to cashflow metric of funds in the same category
	PriceToCashflowCategory *float64 `json:"price_to_cashflow_category,omitempty"`
	// Median market capitalization of a fund
	MedianMarketCapitalization *int64 `json:"median_market_capitalization,omitempty"`
	// Median market capitalization of funds in the same category
	MedianMarketCapitalizationCategory *int64 `json:"median_market_capitalization_category,omitempty"`
	// Earnings growth over the last three years
	Var3YearEarningsGrowth *float64 `json:"3_year_earnings_growth,omitempty"`
	// Earnings growth over the last three years of funds in the same category
	Var3YearEarningsGrowthsCategory *float64 `json:"3_year_earnings_growths_category,omitempty"`
}

// NewGetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics instantiates a new GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics() *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics {
	this := GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics{}
	return &this
}

// NewGetMutualFundsWorld200ResponseMutualFundRiskValuationMetricsWithDefaults instantiates a new GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMutualFundsWorld200ResponseMutualFundRiskValuationMetricsWithDefaults() *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics {
	this := GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics{}
	return &this
}

// GetPriceToEarnings returns the PriceToEarnings field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetPriceToEarnings() float64 {
	if o == nil || IsNil(o.PriceToEarnings) {
		var ret float64
		return ret
	}
	return *o.PriceToEarnings
}

// GetPriceToEarningsOk returns a tuple with the PriceToEarnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetPriceToEarningsOk() (*float64, bool) {
	if o == nil || IsNil(o.PriceToEarnings) {
		return nil, false
	}
	return o.PriceToEarnings, true
}

// HasPriceToEarnings returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) HasPriceToEarnings() bool {
	if o != nil && !IsNil(o.PriceToEarnings) {
		return true
	}

	return false
}

// SetPriceToEarnings gets a reference to the given float64 and assigns it to the PriceToEarnings field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) SetPriceToEarnings(v float64) {
	o.PriceToEarnings = &v
}

// GetPriceToEarningsCategory returns the PriceToEarningsCategory field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetPriceToEarningsCategory() float64 {
	if o == nil || IsNil(o.PriceToEarningsCategory) {
		var ret float64
		return ret
	}
	return *o.PriceToEarningsCategory
}

// GetPriceToEarningsCategoryOk returns a tuple with the PriceToEarningsCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetPriceToEarningsCategoryOk() (*float64, bool) {
	if o == nil || IsNil(o.PriceToEarningsCategory) {
		return nil, false
	}
	return o.PriceToEarningsCategory, true
}

// HasPriceToEarningsCategory returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) HasPriceToEarningsCategory() bool {
	if o != nil && !IsNil(o.PriceToEarningsCategory) {
		return true
	}

	return false
}

// SetPriceToEarningsCategory gets a reference to the given float64 and assigns it to the PriceToEarningsCategory field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) SetPriceToEarningsCategory(v float64) {
	o.PriceToEarningsCategory = &v
}

// GetPriceToBook returns the PriceToBook field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetPriceToBook() float64 {
	if o == nil || IsNil(o.PriceToBook) {
		var ret float64
		return ret
	}
	return *o.PriceToBook
}

// GetPriceToBookOk returns a tuple with the PriceToBook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetPriceToBookOk() (*float64, bool) {
	if o == nil || IsNil(o.PriceToBook) {
		return nil, false
	}
	return o.PriceToBook, true
}

// HasPriceToBook returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) HasPriceToBook() bool {
	if o != nil && !IsNil(o.PriceToBook) {
		return true
	}

	return false
}

// SetPriceToBook gets a reference to the given float64 and assigns it to the PriceToBook field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) SetPriceToBook(v float64) {
	o.PriceToBook = &v
}

// GetPriceToBookCategory returns the PriceToBookCategory field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetPriceToBookCategory() float64 {
	if o == nil || IsNil(o.PriceToBookCategory) {
		var ret float64
		return ret
	}
	return *o.PriceToBookCategory
}

// GetPriceToBookCategoryOk returns a tuple with the PriceToBookCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetPriceToBookCategoryOk() (*float64, bool) {
	if o == nil || IsNil(o.PriceToBookCategory) {
		return nil, false
	}
	return o.PriceToBookCategory, true
}

// HasPriceToBookCategory returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) HasPriceToBookCategory() bool {
	if o != nil && !IsNil(o.PriceToBookCategory) {
		return true
	}

	return false
}

// SetPriceToBookCategory gets a reference to the given float64 and assigns it to the PriceToBookCategory field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) SetPriceToBookCategory(v float64) {
	o.PriceToBookCategory = &v
}

// GetPriceToSales returns the PriceToSales field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetPriceToSales() float64 {
	if o == nil || IsNil(o.PriceToSales) {
		var ret float64
		return ret
	}
	return *o.PriceToSales
}

// GetPriceToSalesOk returns a tuple with the PriceToSales field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetPriceToSalesOk() (*float64, bool) {
	if o == nil || IsNil(o.PriceToSales) {
		return nil, false
	}
	return o.PriceToSales, true
}

// HasPriceToSales returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) HasPriceToSales() bool {
	if o != nil && !IsNil(o.PriceToSales) {
		return true
	}

	return false
}

// SetPriceToSales gets a reference to the given float64 and assigns it to the PriceToSales field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) SetPriceToSales(v float64) {
	o.PriceToSales = &v
}

// GetPriceToSalesCategory returns the PriceToSalesCategory field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetPriceToSalesCategory() float64 {
	if o == nil || IsNil(o.PriceToSalesCategory) {
		var ret float64
		return ret
	}
	return *o.PriceToSalesCategory
}

// GetPriceToSalesCategoryOk returns a tuple with the PriceToSalesCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetPriceToSalesCategoryOk() (*float64, bool) {
	if o == nil || IsNil(o.PriceToSalesCategory) {
		return nil, false
	}
	return o.PriceToSalesCategory, true
}

// HasPriceToSalesCategory returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) HasPriceToSalesCategory() bool {
	if o != nil && !IsNil(o.PriceToSalesCategory) {
		return true
	}

	return false
}

// SetPriceToSalesCategory gets a reference to the given float64 and assigns it to the PriceToSalesCategory field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) SetPriceToSalesCategory(v float64) {
	o.PriceToSalesCategory = &v
}

// GetPriceToCashflow returns the PriceToCashflow field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetPriceToCashflow() float64 {
	if o == nil || IsNil(o.PriceToCashflow) {
		var ret float64
		return ret
	}
	return *o.PriceToCashflow
}

// GetPriceToCashflowOk returns a tuple with the PriceToCashflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetPriceToCashflowOk() (*float64, bool) {
	if o == nil || IsNil(o.PriceToCashflow) {
		return nil, false
	}
	return o.PriceToCashflow, true
}

// HasPriceToCashflow returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) HasPriceToCashflow() bool {
	if o != nil && !IsNil(o.PriceToCashflow) {
		return true
	}

	return false
}

// SetPriceToCashflow gets a reference to the given float64 and assigns it to the PriceToCashflow field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) SetPriceToCashflow(v float64) {
	o.PriceToCashflow = &v
}

// GetPriceToCashflowCategory returns the PriceToCashflowCategory field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetPriceToCashflowCategory() float64 {
	if o == nil || IsNil(o.PriceToCashflowCategory) {
		var ret float64
		return ret
	}
	return *o.PriceToCashflowCategory
}

// GetPriceToCashflowCategoryOk returns a tuple with the PriceToCashflowCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetPriceToCashflowCategoryOk() (*float64, bool) {
	if o == nil || IsNil(o.PriceToCashflowCategory) {
		return nil, false
	}
	return o.PriceToCashflowCategory, true
}

// HasPriceToCashflowCategory returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) HasPriceToCashflowCategory() bool {
	if o != nil && !IsNil(o.PriceToCashflowCategory) {
		return true
	}

	return false
}

// SetPriceToCashflowCategory gets a reference to the given float64 and assigns it to the PriceToCashflowCategory field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) SetPriceToCashflowCategory(v float64) {
	o.PriceToCashflowCategory = &v
}

// GetMedianMarketCapitalization returns the MedianMarketCapitalization field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetMedianMarketCapitalization() int64 {
	if o == nil || IsNil(o.MedianMarketCapitalization) {
		var ret int64
		return ret
	}
	return *o.MedianMarketCapitalization
}

// GetMedianMarketCapitalizationOk returns a tuple with the MedianMarketCapitalization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetMedianMarketCapitalizationOk() (*int64, bool) {
	if o == nil || IsNil(o.MedianMarketCapitalization) {
		return nil, false
	}
	return o.MedianMarketCapitalization, true
}

// HasMedianMarketCapitalization returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) HasMedianMarketCapitalization() bool {
	if o != nil && !IsNil(o.MedianMarketCapitalization) {
		return true
	}

	return false
}

// SetMedianMarketCapitalization gets a reference to the given int64 and assigns it to the MedianMarketCapitalization field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) SetMedianMarketCapitalization(v int64) {
	o.MedianMarketCapitalization = &v
}

// GetMedianMarketCapitalizationCategory returns the MedianMarketCapitalizationCategory field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetMedianMarketCapitalizationCategory() int64 {
	if o == nil || IsNil(o.MedianMarketCapitalizationCategory) {
		var ret int64
		return ret
	}
	return *o.MedianMarketCapitalizationCategory
}

// GetMedianMarketCapitalizationCategoryOk returns a tuple with the MedianMarketCapitalizationCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetMedianMarketCapitalizationCategoryOk() (*int64, bool) {
	if o == nil || IsNil(o.MedianMarketCapitalizationCategory) {
		return nil, false
	}
	return o.MedianMarketCapitalizationCategory, true
}

// HasMedianMarketCapitalizationCategory returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) HasMedianMarketCapitalizationCategory() bool {
	if o != nil && !IsNil(o.MedianMarketCapitalizationCategory) {
		return true
	}

	return false
}

// SetMedianMarketCapitalizationCategory gets a reference to the given int64 and assigns it to the MedianMarketCapitalizationCategory field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) SetMedianMarketCapitalizationCategory(v int64) {
	o.MedianMarketCapitalizationCategory = &v
}

// GetVar3YearEarningsGrowth returns the Var3YearEarningsGrowth field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetVar3YearEarningsGrowth() float64 {
	if o == nil || IsNil(o.Var3YearEarningsGrowth) {
		var ret float64
		return ret
	}
	return *o.Var3YearEarningsGrowth
}

// GetVar3YearEarningsGrowthOk returns a tuple with the Var3YearEarningsGrowth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetVar3YearEarningsGrowthOk() (*float64, bool) {
	if o == nil || IsNil(o.Var3YearEarningsGrowth) {
		return nil, false
	}
	return o.Var3YearEarningsGrowth, true
}

// HasVar3YearEarningsGrowth returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) HasVar3YearEarningsGrowth() bool {
	if o != nil && !IsNil(o.Var3YearEarningsGrowth) {
		return true
	}

	return false
}

// SetVar3YearEarningsGrowth gets a reference to the given float64 and assigns it to the Var3YearEarningsGrowth field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) SetVar3YearEarningsGrowth(v float64) {
	o.Var3YearEarningsGrowth = &v
}

// GetVar3YearEarningsGrowthsCategory returns the Var3YearEarningsGrowthsCategory field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetVar3YearEarningsGrowthsCategory() float64 {
	if o == nil || IsNil(o.Var3YearEarningsGrowthsCategory) {
		var ret float64
		return ret
	}
	return *o.Var3YearEarningsGrowthsCategory
}

// GetVar3YearEarningsGrowthsCategoryOk returns a tuple with the Var3YearEarningsGrowthsCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) GetVar3YearEarningsGrowthsCategoryOk() (*float64, bool) {
	if o == nil || IsNil(o.Var3YearEarningsGrowthsCategory) {
		return nil, false
	}
	return o.Var3YearEarningsGrowthsCategory, true
}

// HasVar3YearEarningsGrowthsCategory returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) HasVar3YearEarningsGrowthsCategory() bool {
	if o != nil && !IsNil(o.Var3YearEarningsGrowthsCategory) {
		return true
	}

	return false
}

// SetVar3YearEarningsGrowthsCategory gets a reference to the given float64 and assigns it to the Var3YearEarningsGrowthsCategory field.
func (o *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) SetVar3YearEarningsGrowthsCategory(v float64) {
	o.Var3YearEarningsGrowthsCategory = &v
}

func (o GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PriceToEarnings) {
		toSerialize["price_to_earnings"] = o.PriceToEarnings
	}
	if !IsNil(o.PriceToEarningsCategory) {
		toSerialize["price_to_earnings_category"] = o.PriceToEarningsCategory
	}
	if !IsNil(o.PriceToBook) {
		toSerialize["price_to_book"] = o.PriceToBook
	}
	if !IsNil(o.PriceToBookCategory) {
		toSerialize["price_to_book_category"] = o.PriceToBookCategory
	}
	if !IsNil(o.PriceToSales) {
		toSerialize["price_to_sales"] = o.PriceToSales
	}
	if !IsNil(o.PriceToSalesCategory) {
		toSerialize["price_to_sales_category"] = o.PriceToSalesCategory
	}
	if !IsNil(o.PriceToCashflow) {
		toSerialize["price_to_cashflow"] = o.PriceToCashflow
	}
	if !IsNil(o.PriceToCashflowCategory) {
		toSerialize["price_to_cashflow_category"] = o.PriceToCashflowCategory
	}
	if !IsNil(o.MedianMarketCapitalization) {
		toSerialize["median_market_capitalization"] = o.MedianMarketCapitalization
	}
	if !IsNil(o.MedianMarketCapitalizationCategory) {
		toSerialize["median_market_capitalization_category"] = o.MedianMarketCapitalizationCategory
	}
	if !IsNil(o.Var3YearEarningsGrowth) {
		toSerialize["3_year_earnings_growth"] = o.Var3YearEarningsGrowth
	}
	if !IsNil(o.Var3YearEarningsGrowthsCategory) {
		toSerialize["3_year_earnings_growths_category"] = o.Var3YearEarningsGrowthsCategory
	}
	return toSerialize, nil
}

type NullableGetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics struct {
	value *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics
	isSet bool
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) Get() *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics {
	return v.value
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) Set(val *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics(val *GetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) *NullableGetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics {
	return &NullableGetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics{value: val, isSet: true}
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundRiskValuationMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
