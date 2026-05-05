// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the GetETFsWorld200ResponseEtfRiskValuationMetrics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetETFsWorld200ResponseEtfRiskValuationMetrics{}

// GetETFsWorld200ResponseEtfRiskValuationMetrics Valuation ratios and metrics of the fund and its category
type GetETFsWorld200ResponseEtfRiskValuationMetrics struct {
	// Fund price to earnings metric
	PriceToEarnings *float64 `json:"price_to_earnings,omitempty"`
	// Fund price to book metric
	PriceToBook *float64 `json:"price_to_book,omitempty"`
	// Fund price to sales metric
	PriceToSales *float64 `json:"price_to_sales,omitempty"`
	// Fund price to cashflow metric
	PriceToCashflow *float64 `json:"price_to_cashflow,omitempty"`
}

// NewGetETFsWorld200ResponseEtfRiskValuationMetrics instantiates a new GetETFsWorld200ResponseEtfRiskValuationMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetETFsWorld200ResponseEtfRiskValuationMetrics() *GetETFsWorld200ResponseEtfRiskValuationMetrics {
	this := GetETFsWorld200ResponseEtfRiskValuationMetrics{}
	return &this
}

// NewGetETFsWorld200ResponseEtfRiskValuationMetricsWithDefaults instantiates a new GetETFsWorld200ResponseEtfRiskValuationMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetETFsWorld200ResponseEtfRiskValuationMetricsWithDefaults() *GetETFsWorld200ResponseEtfRiskValuationMetrics {
	this := GetETFsWorld200ResponseEtfRiskValuationMetrics{}
	return &this
}

// GetPriceToEarnings returns the PriceToEarnings field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfRiskValuationMetrics) GetPriceToEarnings() float64 {
	if o == nil || IsNil(o.PriceToEarnings) {
		var ret float64
		return ret
	}
	return *o.PriceToEarnings
}

// GetPriceToEarningsOk returns a tuple with the PriceToEarnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfRiskValuationMetrics) GetPriceToEarningsOk() (*float64, bool) {
	if o == nil || IsNil(o.PriceToEarnings) {
		return nil, false
	}
	return o.PriceToEarnings, true
}

// HasPriceToEarnings returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfRiskValuationMetrics) HasPriceToEarnings() bool {
	if o != nil && !IsNil(o.PriceToEarnings) {
		return true
	}

	return false
}

// SetPriceToEarnings gets a reference to the given float64 and assigns it to the PriceToEarnings field.
func (o *GetETFsWorld200ResponseEtfRiskValuationMetrics) SetPriceToEarnings(v float64) {
	o.PriceToEarnings = &v
}

// GetPriceToBook returns the PriceToBook field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfRiskValuationMetrics) GetPriceToBook() float64 {
	if o == nil || IsNil(o.PriceToBook) {
		var ret float64
		return ret
	}
	return *o.PriceToBook
}

// GetPriceToBookOk returns a tuple with the PriceToBook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfRiskValuationMetrics) GetPriceToBookOk() (*float64, bool) {
	if o == nil || IsNil(o.PriceToBook) {
		return nil, false
	}
	return o.PriceToBook, true
}

// HasPriceToBook returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfRiskValuationMetrics) HasPriceToBook() bool {
	if o != nil && !IsNil(o.PriceToBook) {
		return true
	}

	return false
}

// SetPriceToBook gets a reference to the given float64 and assigns it to the PriceToBook field.
func (o *GetETFsWorld200ResponseEtfRiskValuationMetrics) SetPriceToBook(v float64) {
	o.PriceToBook = &v
}

// GetPriceToSales returns the PriceToSales field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfRiskValuationMetrics) GetPriceToSales() float64 {
	if o == nil || IsNil(o.PriceToSales) {
		var ret float64
		return ret
	}
	return *o.PriceToSales
}

// GetPriceToSalesOk returns a tuple with the PriceToSales field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfRiskValuationMetrics) GetPriceToSalesOk() (*float64, bool) {
	if o == nil || IsNil(o.PriceToSales) {
		return nil, false
	}
	return o.PriceToSales, true
}

// HasPriceToSales returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfRiskValuationMetrics) HasPriceToSales() bool {
	if o != nil && !IsNil(o.PriceToSales) {
		return true
	}

	return false
}

// SetPriceToSales gets a reference to the given float64 and assigns it to the PriceToSales field.
func (o *GetETFsWorld200ResponseEtfRiskValuationMetrics) SetPriceToSales(v float64) {
	o.PriceToSales = &v
}

// GetPriceToCashflow returns the PriceToCashflow field value if set, zero value otherwise.
func (o *GetETFsWorld200ResponseEtfRiskValuationMetrics) GetPriceToCashflow() float64 {
	if o == nil || IsNil(o.PriceToCashflow) {
		var ret float64
		return ret
	}
	return *o.PriceToCashflow
}

// GetPriceToCashflowOk returns a tuple with the PriceToCashflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetETFsWorld200ResponseEtfRiskValuationMetrics) GetPriceToCashflowOk() (*float64, bool) {
	if o == nil || IsNil(o.PriceToCashflow) {
		return nil, false
	}
	return o.PriceToCashflow, true
}

// HasPriceToCashflow returns a boolean if a field has been set.
func (o *GetETFsWorld200ResponseEtfRiskValuationMetrics) HasPriceToCashflow() bool {
	if o != nil && !IsNil(o.PriceToCashflow) {
		return true
	}

	return false
}

// SetPriceToCashflow gets a reference to the given float64 and assigns it to the PriceToCashflow field.
func (o *GetETFsWorld200ResponseEtfRiskValuationMetrics) SetPriceToCashflow(v float64) {
	o.PriceToCashflow = &v
}

func (o GetETFsWorld200ResponseEtfRiskValuationMetrics) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetETFsWorld200ResponseEtfRiskValuationMetrics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PriceToEarnings) {
		toSerialize["price_to_earnings"] = o.PriceToEarnings
	}
	if !IsNil(o.PriceToBook) {
		toSerialize["price_to_book"] = o.PriceToBook
	}
	if !IsNil(o.PriceToSales) {
		toSerialize["price_to_sales"] = o.PriceToSales
	}
	if !IsNil(o.PriceToCashflow) {
		toSerialize["price_to_cashflow"] = o.PriceToCashflow
	}
	return toSerialize, nil
}

type NullableGetETFsWorld200ResponseEtfRiskValuationMetrics struct {
	value *GetETFsWorld200ResponseEtfRiskValuationMetrics
	isSet bool
}

func (v NullableGetETFsWorld200ResponseEtfRiskValuationMetrics) Get() *GetETFsWorld200ResponseEtfRiskValuationMetrics {
	return v.value
}

func (v *NullableGetETFsWorld200ResponseEtfRiskValuationMetrics) Set(val *GetETFsWorld200ResponseEtfRiskValuationMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableGetETFsWorld200ResponseEtfRiskValuationMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableGetETFsWorld200ResponseEtfRiskValuationMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetETFsWorld200ResponseEtfRiskValuationMetrics(val *GetETFsWorld200ResponseEtfRiskValuationMetrics) *NullableGetETFsWorld200ResponseEtfRiskValuationMetrics {
	return &NullableGetETFsWorld200ResponseEtfRiskValuationMetrics{value: val, isSet: true}
}

func (v NullableGetETFsWorld200ResponseEtfRiskValuationMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetETFsWorld200ResponseEtfRiskValuationMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
