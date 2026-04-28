/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the CashFlowDataForeignAndDomesticSales type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CashFlowDataForeignAndDomesticSales{}

// CashFlowDataForeignAndDomesticSales Foreign and domestic sales
type CashFlowDataForeignAndDomesticSales struct {
	// Foreign sales
	ForeignSales *float64 `json:"foreign_sales,omitempty"`
	// Domestic sales
	DomesticSales *float64 `json:"domestic_sales,omitempty"`
	// Adjusted geography segment data
	AdjustedGeographySegmentData *float64 `json:"adjusted_geography_segment_data,omitempty"`
}

// NewCashFlowDataForeignAndDomesticSales instantiates a new CashFlowDataForeignAndDomesticSales object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCashFlowDataForeignAndDomesticSales() *CashFlowDataForeignAndDomesticSales {
	this := CashFlowDataForeignAndDomesticSales{}
	return &this
}

// NewCashFlowDataForeignAndDomesticSalesWithDefaults instantiates a new CashFlowDataForeignAndDomesticSales object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCashFlowDataForeignAndDomesticSalesWithDefaults() *CashFlowDataForeignAndDomesticSales {
	this := CashFlowDataForeignAndDomesticSales{}
	return &this
}

// GetForeignSales returns the ForeignSales field value if set, zero value otherwise.
func (o *CashFlowDataForeignAndDomesticSales) GetForeignSales() float64 {
	if o == nil || IsNil(o.ForeignSales) {
		var ret float64
		return ret
	}
	return *o.ForeignSales
}

// GetForeignSalesOk returns a tuple with the ForeignSales field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataForeignAndDomesticSales) GetForeignSalesOk() (*float64, bool) {
	if o == nil || IsNil(o.ForeignSales) {
		return nil, false
	}
	return o.ForeignSales, true
}

// HasForeignSales returns a boolean if a field has been set.
func (o *CashFlowDataForeignAndDomesticSales) HasForeignSales() bool {
	if o != nil && !IsNil(o.ForeignSales) {
		return true
	}

	return false
}

// SetForeignSales gets a reference to the given float64 and assigns it to the ForeignSales field.
func (o *CashFlowDataForeignAndDomesticSales) SetForeignSales(v float64) {
	o.ForeignSales = &v
}

// GetDomesticSales returns the DomesticSales field value if set, zero value otherwise.
func (o *CashFlowDataForeignAndDomesticSales) GetDomesticSales() float64 {
	if o == nil || IsNil(o.DomesticSales) {
		var ret float64
		return ret
	}
	return *o.DomesticSales
}

// GetDomesticSalesOk returns a tuple with the DomesticSales field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataForeignAndDomesticSales) GetDomesticSalesOk() (*float64, bool) {
	if o == nil || IsNil(o.DomesticSales) {
		return nil, false
	}
	return o.DomesticSales, true
}

// HasDomesticSales returns a boolean if a field has been set.
func (o *CashFlowDataForeignAndDomesticSales) HasDomesticSales() bool {
	if o != nil && !IsNil(o.DomesticSales) {
		return true
	}

	return false
}

// SetDomesticSales gets a reference to the given float64 and assigns it to the DomesticSales field.
func (o *CashFlowDataForeignAndDomesticSales) SetDomesticSales(v float64) {
	o.DomesticSales = &v
}

// GetAdjustedGeographySegmentData returns the AdjustedGeographySegmentData field value if set, zero value otherwise.
func (o *CashFlowDataForeignAndDomesticSales) GetAdjustedGeographySegmentData() float64 {
	if o == nil || IsNil(o.AdjustedGeographySegmentData) {
		var ret float64
		return ret
	}
	return *o.AdjustedGeographySegmentData
}

// GetAdjustedGeographySegmentDataOk returns a tuple with the AdjustedGeographySegmentData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataForeignAndDomesticSales) GetAdjustedGeographySegmentDataOk() (*float64, bool) {
	if o == nil || IsNil(o.AdjustedGeographySegmentData) {
		return nil, false
	}
	return o.AdjustedGeographySegmentData, true
}

// HasAdjustedGeographySegmentData returns a boolean if a field has been set.
func (o *CashFlowDataForeignAndDomesticSales) HasAdjustedGeographySegmentData() bool {
	if o != nil && !IsNil(o.AdjustedGeographySegmentData) {
		return true
	}

	return false
}

// SetAdjustedGeographySegmentData gets a reference to the given float64 and assigns it to the AdjustedGeographySegmentData field.
func (o *CashFlowDataForeignAndDomesticSales) SetAdjustedGeographySegmentData(v float64) {
	o.AdjustedGeographySegmentData = &v
}

func (o CashFlowDataForeignAndDomesticSales) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CashFlowDataForeignAndDomesticSales) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ForeignSales) {
		toSerialize["foreign_sales"] = o.ForeignSales
	}
	if !IsNil(o.DomesticSales) {
		toSerialize["domestic_sales"] = o.DomesticSales
	}
	if !IsNil(o.AdjustedGeographySegmentData) {
		toSerialize["adjusted_geography_segment_data"] = o.AdjustedGeographySegmentData
	}
	return toSerialize, nil
}

type NullableCashFlowDataForeignAndDomesticSales struct {
	value *CashFlowDataForeignAndDomesticSales
	isSet bool
}

func (v NullableCashFlowDataForeignAndDomesticSales) Get() *CashFlowDataForeignAndDomesticSales {
	return v.value
}

func (v *NullableCashFlowDataForeignAndDomesticSales) Set(val *CashFlowDataForeignAndDomesticSales) {
	v.value = val
	v.isSet = true
}

func (v NullableCashFlowDataForeignAndDomesticSales) IsSet() bool {
	return v.isSet
}

func (v *NullableCashFlowDataForeignAndDomesticSales) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCashFlowDataForeignAndDomesticSales(val *CashFlowDataForeignAndDomesticSales) *NullableCashFlowDataForeignAndDomesticSales {
	return &NullableCashFlowDataForeignAndDomesticSales{value: val, isSet: true}
}

func (v NullableCashFlowDataForeignAndDomesticSales) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCashFlowDataForeignAndDomesticSales) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
