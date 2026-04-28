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

// checks if the CashFlowStruct type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CashFlowStruct{}

// CashFlowStruct struct for CashFlowStruct
type CashFlowStruct struct {
	// Date of fiscal period ending
	FiscalDate string `json:"fiscal_date"`
	// Fiscal quarter. Visible when `&period=quarterly`
	Quarter *string `json:"quarter,omitempty"`
	// Fiscal year
	Year                *int64                             `json:"year,omitempty"`
	OperatingActivities *CashFlowStructOperatingActivities `json:"operating_activities,omitempty"`
	InvestingActivities *CashFlowStructInvestingActivities `json:"investing_activities,omitempty"`
	FinancingActivities *CashFlowStructFinancingActivities `json:"financing_activities,omitempty"`
	// Returns the amount of cash a company has when adding the change in cash and beginning cash balance for the current fiscal period
	EndCashPosition *float64 `json:"end_cash_position,omitempty"`
	// Refers to supplemental data about income tax paid
	IncomeTaxPaid *float64 `json:"income_tax_paid,omitempty"`
	// Refers to supplemental data about interest paid
	InterestPaid *float64 `json:"interest_paid,omitempty"`
	// Represents the cash a company generates after accounting for cash outflows to support operations and maintain its capital assets
	FreeCashFlow *float64 `json:"free_cash_flow,omitempty"`
}

type _CashFlowStruct CashFlowStruct

// NewCashFlowStruct instantiates a new CashFlowStruct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCashFlowStruct(fiscalDate string) *CashFlowStruct {
	this := CashFlowStruct{}
	this.FiscalDate = fiscalDate
	return &this
}

// NewCashFlowStructWithDefaults instantiates a new CashFlowStruct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCashFlowStructWithDefaults() *CashFlowStruct {
	this := CashFlowStruct{}
	return &this
}

// GetFiscalDate returns the FiscalDate field value
func (o *CashFlowStruct) GetFiscalDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FiscalDate
}

// GetFiscalDateOk returns a tuple with the FiscalDate field value
// and a boolean to check if the value has been set.
func (o *CashFlowStruct) GetFiscalDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FiscalDate, true
}

// SetFiscalDate sets field value
func (o *CashFlowStruct) SetFiscalDate(v string) {
	o.FiscalDate = v
}

// GetQuarter returns the Quarter field value if set, zero value otherwise.
func (o *CashFlowStruct) GetQuarter() string {
	if o == nil || IsNil(o.Quarter) {
		var ret string
		return ret
	}
	return *o.Quarter
}

// GetQuarterOk returns a tuple with the Quarter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowStruct) GetQuarterOk() (*string, bool) {
	if o == nil || IsNil(o.Quarter) {
		return nil, false
	}
	return o.Quarter, true
}

// HasQuarter returns a boolean if a field has been set.
func (o *CashFlowStruct) HasQuarter() bool {
	if o != nil && !IsNil(o.Quarter) {
		return true
	}

	return false
}

// SetQuarter gets a reference to the given string and assigns it to the Quarter field.
func (o *CashFlowStruct) SetQuarter(v string) {
	o.Quarter = &v
}

// GetYear returns the Year field value if set, zero value otherwise.
func (o *CashFlowStruct) GetYear() int64 {
	if o == nil || IsNil(o.Year) {
		var ret int64
		return ret
	}
	return *o.Year
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowStruct) GetYearOk() (*int64, bool) {
	if o == nil || IsNil(o.Year) {
		return nil, false
	}
	return o.Year, true
}

// HasYear returns a boolean if a field has been set.
func (o *CashFlowStruct) HasYear() bool {
	if o != nil && !IsNil(o.Year) {
		return true
	}

	return false
}

// SetYear gets a reference to the given int64 and assigns it to the Year field.
func (o *CashFlowStruct) SetYear(v int64) {
	o.Year = &v
}

// GetOperatingActivities returns the OperatingActivities field value if set, zero value otherwise.
func (o *CashFlowStruct) GetOperatingActivities() CashFlowStructOperatingActivities {
	if o == nil || IsNil(o.OperatingActivities) {
		var ret CashFlowStructOperatingActivities
		return ret
	}
	return *o.OperatingActivities
}

// GetOperatingActivitiesOk returns a tuple with the OperatingActivities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowStruct) GetOperatingActivitiesOk() (*CashFlowStructOperatingActivities, bool) {
	if o == nil || IsNil(o.OperatingActivities) {
		return nil, false
	}
	return o.OperatingActivities, true
}

// HasOperatingActivities returns a boolean if a field has been set.
func (o *CashFlowStruct) HasOperatingActivities() bool {
	if o != nil && !IsNil(o.OperatingActivities) {
		return true
	}

	return false
}

// SetOperatingActivities gets a reference to the given CashFlowStructOperatingActivities and assigns it to the OperatingActivities field.
func (o *CashFlowStruct) SetOperatingActivities(v CashFlowStructOperatingActivities) {
	o.OperatingActivities = &v
}

// GetInvestingActivities returns the InvestingActivities field value if set, zero value otherwise.
func (o *CashFlowStruct) GetInvestingActivities() CashFlowStructInvestingActivities {
	if o == nil || IsNil(o.InvestingActivities) {
		var ret CashFlowStructInvestingActivities
		return ret
	}
	return *o.InvestingActivities
}

// GetInvestingActivitiesOk returns a tuple with the InvestingActivities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowStruct) GetInvestingActivitiesOk() (*CashFlowStructInvestingActivities, bool) {
	if o == nil || IsNil(o.InvestingActivities) {
		return nil, false
	}
	return o.InvestingActivities, true
}

// HasInvestingActivities returns a boolean if a field has been set.
func (o *CashFlowStruct) HasInvestingActivities() bool {
	if o != nil && !IsNil(o.InvestingActivities) {
		return true
	}

	return false
}

// SetInvestingActivities gets a reference to the given CashFlowStructInvestingActivities and assigns it to the InvestingActivities field.
func (o *CashFlowStruct) SetInvestingActivities(v CashFlowStructInvestingActivities) {
	o.InvestingActivities = &v
}

// GetFinancingActivities returns the FinancingActivities field value if set, zero value otherwise.
func (o *CashFlowStruct) GetFinancingActivities() CashFlowStructFinancingActivities {
	if o == nil || IsNil(o.FinancingActivities) {
		var ret CashFlowStructFinancingActivities
		return ret
	}
	return *o.FinancingActivities
}

// GetFinancingActivitiesOk returns a tuple with the FinancingActivities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowStruct) GetFinancingActivitiesOk() (*CashFlowStructFinancingActivities, bool) {
	if o == nil || IsNil(o.FinancingActivities) {
		return nil, false
	}
	return o.FinancingActivities, true
}

// HasFinancingActivities returns a boolean if a field has been set.
func (o *CashFlowStruct) HasFinancingActivities() bool {
	if o != nil && !IsNil(o.FinancingActivities) {
		return true
	}

	return false
}

// SetFinancingActivities gets a reference to the given CashFlowStructFinancingActivities and assigns it to the FinancingActivities field.
func (o *CashFlowStruct) SetFinancingActivities(v CashFlowStructFinancingActivities) {
	o.FinancingActivities = &v
}

// GetEndCashPosition returns the EndCashPosition field value if set, zero value otherwise.
func (o *CashFlowStruct) GetEndCashPosition() float64 {
	if o == nil || IsNil(o.EndCashPosition) {
		var ret float64
		return ret
	}
	return *o.EndCashPosition
}

// GetEndCashPositionOk returns a tuple with the EndCashPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowStruct) GetEndCashPositionOk() (*float64, bool) {
	if o == nil || IsNil(o.EndCashPosition) {
		return nil, false
	}
	return o.EndCashPosition, true
}

// HasEndCashPosition returns a boolean if a field has been set.
func (o *CashFlowStruct) HasEndCashPosition() bool {
	if o != nil && !IsNil(o.EndCashPosition) {
		return true
	}

	return false
}

// SetEndCashPosition gets a reference to the given float64 and assigns it to the EndCashPosition field.
func (o *CashFlowStruct) SetEndCashPosition(v float64) {
	o.EndCashPosition = &v
}

// GetIncomeTaxPaid returns the IncomeTaxPaid field value if set, zero value otherwise.
func (o *CashFlowStruct) GetIncomeTaxPaid() float64 {
	if o == nil || IsNil(o.IncomeTaxPaid) {
		var ret float64
		return ret
	}
	return *o.IncomeTaxPaid
}

// GetIncomeTaxPaidOk returns a tuple with the IncomeTaxPaid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowStruct) GetIncomeTaxPaidOk() (*float64, bool) {
	if o == nil || IsNil(o.IncomeTaxPaid) {
		return nil, false
	}
	return o.IncomeTaxPaid, true
}

// HasIncomeTaxPaid returns a boolean if a field has been set.
func (o *CashFlowStruct) HasIncomeTaxPaid() bool {
	if o != nil && !IsNil(o.IncomeTaxPaid) {
		return true
	}

	return false
}

// SetIncomeTaxPaid gets a reference to the given float64 and assigns it to the IncomeTaxPaid field.
func (o *CashFlowStruct) SetIncomeTaxPaid(v float64) {
	o.IncomeTaxPaid = &v
}

// GetInterestPaid returns the InterestPaid field value if set, zero value otherwise.
func (o *CashFlowStruct) GetInterestPaid() float64 {
	if o == nil || IsNil(o.InterestPaid) {
		var ret float64
		return ret
	}
	return *o.InterestPaid
}

// GetInterestPaidOk returns a tuple with the InterestPaid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowStruct) GetInterestPaidOk() (*float64, bool) {
	if o == nil || IsNil(o.InterestPaid) {
		return nil, false
	}
	return o.InterestPaid, true
}

// HasInterestPaid returns a boolean if a field has been set.
func (o *CashFlowStruct) HasInterestPaid() bool {
	if o != nil && !IsNil(o.InterestPaid) {
		return true
	}

	return false
}

// SetInterestPaid gets a reference to the given float64 and assigns it to the InterestPaid field.
func (o *CashFlowStruct) SetInterestPaid(v float64) {
	o.InterestPaid = &v
}

// GetFreeCashFlow returns the FreeCashFlow field value if set, zero value otherwise.
func (o *CashFlowStruct) GetFreeCashFlow() float64 {
	if o == nil || IsNil(o.FreeCashFlow) {
		var ret float64
		return ret
	}
	return *o.FreeCashFlow
}

// GetFreeCashFlowOk returns a tuple with the FreeCashFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowStruct) GetFreeCashFlowOk() (*float64, bool) {
	if o == nil || IsNil(o.FreeCashFlow) {
		return nil, false
	}
	return o.FreeCashFlow, true
}

// HasFreeCashFlow returns a boolean if a field has been set.
func (o *CashFlowStruct) HasFreeCashFlow() bool {
	if o != nil && !IsNil(o.FreeCashFlow) {
		return true
	}

	return false
}

// SetFreeCashFlow gets a reference to the given float64 and assigns it to the FreeCashFlow field.
func (o *CashFlowStruct) SetFreeCashFlow(v float64) {
	o.FreeCashFlow = &v
}

func (o CashFlowStruct) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CashFlowStruct) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fiscal_date"] = o.FiscalDate
	if !IsNil(o.Quarter) {
		toSerialize["quarter"] = o.Quarter
	}
	if !IsNil(o.Year) {
		toSerialize["year"] = o.Year
	}
	if !IsNil(o.OperatingActivities) {
		toSerialize["operating_activities"] = o.OperatingActivities
	}
	if !IsNil(o.InvestingActivities) {
		toSerialize["investing_activities"] = o.InvestingActivities
	}
	if !IsNil(o.FinancingActivities) {
		toSerialize["financing_activities"] = o.FinancingActivities
	}
	if !IsNil(o.EndCashPosition) {
		toSerialize["end_cash_position"] = o.EndCashPosition
	}
	if !IsNil(o.IncomeTaxPaid) {
		toSerialize["income_tax_paid"] = o.IncomeTaxPaid
	}
	if !IsNil(o.InterestPaid) {
		toSerialize["interest_paid"] = o.InterestPaid
	}
	if !IsNil(o.FreeCashFlow) {
		toSerialize["free_cash_flow"] = o.FreeCashFlow
	}
	return toSerialize, nil
}

func (o *CashFlowStruct) UnmarshalJSON(data []byte) (err error) {
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

	varCashFlowStruct := _CashFlowStruct{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCashFlowStruct)

	if err != nil {
		return err
	}

	*o = CashFlowStruct(varCashFlowStruct)

	return err
}

type NullableCashFlowStruct struct {
	value *CashFlowStruct
	isSet bool
}

func (v NullableCashFlowStruct) Get() *CashFlowStruct {
	return v.value
}

func (v *NullableCashFlowStruct) Set(val *CashFlowStruct) {
	v.value = val
	v.isSet = true
}

func (v NullableCashFlowStruct) IsSet() bool {
	return v.isSet
}

func (v *NullableCashFlowStruct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCashFlowStruct(val *CashFlowStruct) *NullableCashFlowStruct {
	return &NullableCashFlowStruct{value: val, isSet: true}
}

func (v NullableCashFlowStruct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCashFlowStruct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
