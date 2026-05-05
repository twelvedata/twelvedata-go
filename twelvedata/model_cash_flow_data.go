// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CashFlowData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CashFlowData{}

// CashFlowData CashFlowData represents cash flow data for a specific fiscal date
type CashFlowData struct {
	// Date of fiscal period ending
	FiscalDate string `json:"fiscal_date"`
	// Year of the cash flow statement
	Year                            *int64                                       `json:"year,omitempty"`
	CashFlowFromOperatingActivities *CashFlowDataCashFlowFromOperatingActivities `json:"cash_flow_from_operating_activities,omitempty"`
	CashFlowFromInvestingActivities *CashFlowDataCashFlowFromInvestingActivities `json:"cash_flow_from_investing_activities,omitempty"`
	CashFlowFromFinancingActivities *CashFlowDataCashFlowFromFinancingActivities `json:"cash_flow_from_financing_activities,omitempty"`
	SupplementalData                *CashFlowDataSupplementalData                `json:"supplemental_data,omitempty"`
	ForeignAndDomesticSales         *CashFlowDataForeignAndDomesticSales         `json:"foreign_and_domestic_sales,omitempty"`
	CashPosition                    *CashFlowDataCashPosition                    `json:"cash_position,omitempty"`
	DirectMethodCashFlow            *CashFlowDataDirectMethodCashFlow            `json:"direct_method_cash_flow,omitempty"`
}

type _CashFlowData CashFlowData

// NewCashFlowData instantiates a new CashFlowData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCashFlowData(fiscalDate string) *CashFlowData {
	this := CashFlowData{}
	this.FiscalDate = fiscalDate
	return &this
}

// NewCashFlowDataWithDefaults instantiates a new CashFlowData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCashFlowDataWithDefaults() *CashFlowData {
	this := CashFlowData{}
	return &this
}

// GetFiscalDate returns the FiscalDate field value
func (o *CashFlowData) GetFiscalDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FiscalDate
}

// GetFiscalDateOk returns a tuple with the FiscalDate field value
// and a boolean to check if the value has been set.
func (o *CashFlowData) GetFiscalDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FiscalDate, true
}

// SetFiscalDate sets field value
func (o *CashFlowData) SetFiscalDate(v string) {
	o.FiscalDate = v
}

// GetYear returns the Year field value if set, zero value otherwise.
func (o *CashFlowData) GetYear() int64 {
	if o == nil || IsNil(o.Year) {
		var ret int64
		return ret
	}
	return *o.Year
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowData) GetYearOk() (*int64, bool) {
	if o == nil || IsNil(o.Year) {
		return nil, false
	}
	return o.Year, true
}

// HasYear returns a boolean if a field has been set.
func (o *CashFlowData) HasYear() bool {
	if o != nil && !IsNil(o.Year) {
		return true
	}

	return false
}

// SetYear gets a reference to the given int64 and assigns it to the Year field.
func (o *CashFlowData) SetYear(v int64) {
	o.Year = &v
}

// GetCashFlowFromOperatingActivities returns the CashFlowFromOperatingActivities field value if set, zero value otherwise.
func (o *CashFlowData) GetCashFlowFromOperatingActivities() CashFlowDataCashFlowFromOperatingActivities {
	if o == nil || IsNil(o.CashFlowFromOperatingActivities) {
		var ret CashFlowDataCashFlowFromOperatingActivities
		return ret
	}
	return *o.CashFlowFromOperatingActivities
}

// GetCashFlowFromOperatingActivitiesOk returns a tuple with the CashFlowFromOperatingActivities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowData) GetCashFlowFromOperatingActivitiesOk() (*CashFlowDataCashFlowFromOperatingActivities, bool) {
	if o == nil || IsNil(o.CashFlowFromOperatingActivities) {
		return nil, false
	}
	return o.CashFlowFromOperatingActivities, true
}

// HasCashFlowFromOperatingActivities returns a boolean if a field has been set.
func (o *CashFlowData) HasCashFlowFromOperatingActivities() bool {
	if o != nil && !IsNil(o.CashFlowFromOperatingActivities) {
		return true
	}

	return false
}

// SetCashFlowFromOperatingActivities gets a reference to the given CashFlowDataCashFlowFromOperatingActivities and assigns it to the CashFlowFromOperatingActivities field.
func (o *CashFlowData) SetCashFlowFromOperatingActivities(v CashFlowDataCashFlowFromOperatingActivities) {
	o.CashFlowFromOperatingActivities = &v
}

// GetCashFlowFromInvestingActivities returns the CashFlowFromInvestingActivities field value if set, zero value otherwise.
func (o *CashFlowData) GetCashFlowFromInvestingActivities() CashFlowDataCashFlowFromInvestingActivities {
	if o == nil || IsNil(o.CashFlowFromInvestingActivities) {
		var ret CashFlowDataCashFlowFromInvestingActivities
		return ret
	}
	return *o.CashFlowFromInvestingActivities
}

// GetCashFlowFromInvestingActivitiesOk returns a tuple with the CashFlowFromInvestingActivities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowData) GetCashFlowFromInvestingActivitiesOk() (*CashFlowDataCashFlowFromInvestingActivities, bool) {
	if o == nil || IsNil(o.CashFlowFromInvestingActivities) {
		return nil, false
	}
	return o.CashFlowFromInvestingActivities, true
}

// HasCashFlowFromInvestingActivities returns a boolean if a field has been set.
func (o *CashFlowData) HasCashFlowFromInvestingActivities() bool {
	if o != nil && !IsNil(o.CashFlowFromInvestingActivities) {
		return true
	}

	return false
}

// SetCashFlowFromInvestingActivities gets a reference to the given CashFlowDataCashFlowFromInvestingActivities and assigns it to the CashFlowFromInvestingActivities field.
func (o *CashFlowData) SetCashFlowFromInvestingActivities(v CashFlowDataCashFlowFromInvestingActivities) {
	o.CashFlowFromInvestingActivities = &v
}

// GetCashFlowFromFinancingActivities returns the CashFlowFromFinancingActivities field value if set, zero value otherwise.
func (o *CashFlowData) GetCashFlowFromFinancingActivities() CashFlowDataCashFlowFromFinancingActivities {
	if o == nil || IsNil(o.CashFlowFromFinancingActivities) {
		var ret CashFlowDataCashFlowFromFinancingActivities
		return ret
	}
	return *o.CashFlowFromFinancingActivities
}

// GetCashFlowFromFinancingActivitiesOk returns a tuple with the CashFlowFromFinancingActivities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowData) GetCashFlowFromFinancingActivitiesOk() (*CashFlowDataCashFlowFromFinancingActivities, bool) {
	if o == nil || IsNil(o.CashFlowFromFinancingActivities) {
		return nil, false
	}
	return o.CashFlowFromFinancingActivities, true
}

// HasCashFlowFromFinancingActivities returns a boolean if a field has been set.
func (o *CashFlowData) HasCashFlowFromFinancingActivities() bool {
	if o != nil && !IsNil(o.CashFlowFromFinancingActivities) {
		return true
	}

	return false
}

// SetCashFlowFromFinancingActivities gets a reference to the given CashFlowDataCashFlowFromFinancingActivities and assigns it to the CashFlowFromFinancingActivities field.
func (o *CashFlowData) SetCashFlowFromFinancingActivities(v CashFlowDataCashFlowFromFinancingActivities) {
	o.CashFlowFromFinancingActivities = &v
}

// GetSupplementalData returns the SupplementalData field value if set, zero value otherwise.
func (o *CashFlowData) GetSupplementalData() CashFlowDataSupplementalData {
	if o == nil || IsNil(o.SupplementalData) {
		var ret CashFlowDataSupplementalData
		return ret
	}
	return *o.SupplementalData
}

// GetSupplementalDataOk returns a tuple with the SupplementalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowData) GetSupplementalDataOk() (*CashFlowDataSupplementalData, bool) {
	if o == nil || IsNil(o.SupplementalData) {
		return nil, false
	}
	return o.SupplementalData, true
}

// HasSupplementalData returns a boolean if a field has been set.
func (o *CashFlowData) HasSupplementalData() bool {
	if o != nil && !IsNil(o.SupplementalData) {
		return true
	}

	return false
}

// SetSupplementalData gets a reference to the given CashFlowDataSupplementalData and assigns it to the SupplementalData field.
func (o *CashFlowData) SetSupplementalData(v CashFlowDataSupplementalData) {
	o.SupplementalData = &v
}

// GetForeignAndDomesticSales returns the ForeignAndDomesticSales field value if set, zero value otherwise.
func (o *CashFlowData) GetForeignAndDomesticSales() CashFlowDataForeignAndDomesticSales {
	if o == nil || IsNil(o.ForeignAndDomesticSales) {
		var ret CashFlowDataForeignAndDomesticSales
		return ret
	}
	return *o.ForeignAndDomesticSales
}

// GetForeignAndDomesticSalesOk returns a tuple with the ForeignAndDomesticSales field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowData) GetForeignAndDomesticSalesOk() (*CashFlowDataForeignAndDomesticSales, bool) {
	if o == nil || IsNil(o.ForeignAndDomesticSales) {
		return nil, false
	}
	return o.ForeignAndDomesticSales, true
}

// HasForeignAndDomesticSales returns a boolean if a field has been set.
func (o *CashFlowData) HasForeignAndDomesticSales() bool {
	if o != nil && !IsNil(o.ForeignAndDomesticSales) {
		return true
	}

	return false
}

// SetForeignAndDomesticSales gets a reference to the given CashFlowDataForeignAndDomesticSales and assigns it to the ForeignAndDomesticSales field.
func (o *CashFlowData) SetForeignAndDomesticSales(v CashFlowDataForeignAndDomesticSales) {
	o.ForeignAndDomesticSales = &v
}

// GetCashPosition returns the CashPosition field value if set, zero value otherwise.
func (o *CashFlowData) GetCashPosition() CashFlowDataCashPosition {
	if o == nil || IsNil(o.CashPosition) {
		var ret CashFlowDataCashPosition
		return ret
	}
	return *o.CashPosition
}

// GetCashPositionOk returns a tuple with the CashPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowData) GetCashPositionOk() (*CashFlowDataCashPosition, bool) {
	if o == nil || IsNil(o.CashPosition) {
		return nil, false
	}
	return o.CashPosition, true
}

// HasCashPosition returns a boolean if a field has been set.
func (o *CashFlowData) HasCashPosition() bool {
	if o != nil && !IsNil(o.CashPosition) {
		return true
	}

	return false
}

// SetCashPosition gets a reference to the given CashFlowDataCashPosition and assigns it to the CashPosition field.
func (o *CashFlowData) SetCashPosition(v CashFlowDataCashPosition) {
	o.CashPosition = &v
}

// GetDirectMethodCashFlow returns the DirectMethodCashFlow field value if set, zero value otherwise.
func (o *CashFlowData) GetDirectMethodCashFlow() CashFlowDataDirectMethodCashFlow {
	if o == nil || IsNil(o.DirectMethodCashFlow) {
		var ret CashFlowDataDirectMethodCashFlow
		return ret
	}
	return *o.DirectMethodCashFlow
}

// GetDirectMethodCashFlowOk returns a tuple with the DirectMethodCashFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowData) GetDirectMethodCashFlowOk() (*CashFlowDataDirectMethodCashFlow, bool) {
	if o == nil || IsNil(o.DirectMethodCashFlow) {
		return nil, false
	}
	return o.DirectMethodCashFlow, true
}

// HasDirectMethodCashFlow returns a boolean if a field has been set.
func (o *CashFlowData) HasDirectMethodCashFlow() bool {
	if o != nil && !IsNil(o.DirectMethodCashFlow) {
		return true
	}

	return false
}

// SetDirectMethodCashFlow gets a reference to the given CashFlowDataDirectMethodCashFlow and assigns it to the DirectMethodCashFlow field.
func (o *CashFlowData) SetDirectMethodCashFlow(v CashFlowDataDirectMethodCashFlow) {
	o.DirectMethodCashFlow = &v
}

func (o CashFlowData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CashFlowData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fiscal_date"] = o.FiscalDate
	if !IsNil(o.Year) {
		toSerialize["year"] = o.Year
	}
	if !IsNil(o.CashFlowFromOperatingActivities) {
		toSerialize["cash_flow_from_operating_activities"] = o.CashFlowFromOperatingActivities
	}
	if !IsNil(o.CashFlowFromInvestingActivities) {
		toSerialize["cash_flow_from_investing_activities"] = o.CashFlowFromInvestingActivities
	}
	if !IsNil(o.CashFlowFromFinancingActivities) {
		toSerialize["cash_flow_from_financing_activities"] = o.CashFlowFromFinancingActivities
	}
	if !IsNil(o.SupplementalData) {
		toSerialize["supplemental_data"] = o.SupplementalData
	}
	if !IsNil(o.ForeignAndDomesticSales) {
		toSerialize["foreign_and_domestic_sales"] = o.ForeignAndDomesticSales
	}
	if !IsNil(o.CashPosition) {
		toSerialize["cash_position"] = o.CashPosition
	}
	if !IsNil(o.DirectMethodCashFlow) {
		toSerialize["direct_method_cash_flow"] = o.DirectMethodCashFlow
	}
	return toSerialize, nil
}

func (o *CashFlowData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fiscal_date",
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

	varCashFlowData := _CashFlowData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varCashFlowData)

	if err != nil {
		return err
	}

	*o = CashFlowData(varCashFlowData)

	return err
}

type NullableCashFlowData struct {
	value *CashFlowData
	isSet bool
}

func (v NullableCashFlowData) Get() *CashFlowData {
	return v.value
}

func (v *NullableCashFlowData) Set(val *CashFlowData) {
	v.value = val
	v.isSet = true
}

func (v NullableCashFlowData) IsSet() bool {
	return v.isSet
}

func (v *NullableCashFlowData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCashFlowData(val *CashFlowData) *NullableCashFlowData {
	return &NullableCashFlowData{value: val, isSet: true}
}

func (v NullableCashFlowData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCashFlowData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
