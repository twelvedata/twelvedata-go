/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the CashFlowDataSupplementalData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CashFlowDataSupplementalData{}

// CashFlowDataSupplementalData Supplemental data
type CashFlowDataSupplementalData struct {
	// Interest paid supplemental data
	InterestPaidSupplementalData *float64 `json:"interest_paid_supplemental_data,omitempty"`
	// Income tax paid supplemental data
	IncomeTaxPaidSupplementalData *float64 `json:"income_tax_paid_supplemental_data,omitempty"`
}

// NewCashFlowDataSupplementalData instantiates a new CashFlowDataSupplementalData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCashFlowDataSupplementalData() *CashFlowDataSupplementalData {
	this := CashFlowDataSupplementalData{}
	return &this
}

// NewCashFlowDataSupplementalDataWithDefaults instantiates a new CashFlowDataSupplementalData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCashFlowDataSupplementalDataWithDefaults() *CashFlowDataSupplementalData {
	this := CashFlowDataSupplementalData{}
	return &this
}

// GetInterestPaidSupplementalData returns the InterestPaidSupplementalData field value if set, zero value otherwise.
func (o *CashFlowDataSupplementalData) GetInterestPaidSupplementalData() float64 {
	if o == nil || IsNil(o.InterestPaidSupplementalData) {
		var ret float64
		return ret
	}
	return *o.InterestPaidSupplementalData
}

// GetInterestPaidSupplementalDataOk returns a tuple with the InterestPaidSupplementalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataSupplementalData) GetInterestPaidSupplementalDataOk() (*float64, bool) {
	if o == nil || IsNil(o.InterestPaidSupplementalData) {
		return nil, false
	}
	return o.InterestPaidSupplementalData, true
}

// HasInterestPaidSupplementalData returns a boolean if a field has been set.
func (o *CashFlowDataSupplementalData) HasInterestPaidSupplementalData() bool {
	if o != nil && !IsNil(o.InterestPaidSupplementalData) {
		return true
	}

	return false
}

// SetInterestPaidSupplementalData gets a reference to the given float64 and assigns it to the InterestPaidSupplementalData field.
func (o *CashFlowDataSupplementalData) SetInterestPaidSupplementalData(v float64) {
	o.InterestPaidSupplementalData = &v
}

// GetIncomeTaxPaidSupplementalData returns the IncomeTaxPaidSupplementalData field value if set, zero value otherwise.
func (o *CashFlowDataSupplementalData) GetIncomeTaxPaidSupplementalData() float64 {
	if o == nil || IsNil(o.IncomeTaxPaidSupplementalData) {
		var ret float64
		return ret
	}
	return *o.IncomeTaxPaidSupplementalData
}

// GetIncomeTaxPaidSupplementalDataOk returns a tuple with the IncomeTaxPaidSupplementalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataSupplementalData) GetIncomeTaxPaidSupplementalDataOk() (*float64, bool) {
	if o == nil || IsNil(o.IncomeTaxPaidSupplementalData) {
		return nil, false
	}
	return o.IncomeTaxPaidSupplementalData, true
}

// HasIncomeTaxPaidSupplementalData returns a boolean if a field has been set.
func (o *CashFlowDataSupplementalData) HasIncomeTaxPaidSupplementalData() bool {
	if o != nil && !IsNil(o.IncomeTaxPaidSupplementalData) {
		return true
	}

	return false
}

// SetIncomeTaxPaidSupplementalData gets a reference to the given float64 and assigns it to the IncomeTaxPaidSupplementalData field.
func (o *CashFlowDataSupplementalData) SetIncomeTaxPaidSupplementalData(v float64) {
	o.IncomeTaxPaidSupplementalData = &v
}

func (o CashFlowDataSupplementalData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CashFlowDataSupplementalData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InterestPaidSupplementalData) {
		toSerialize["interest_paid_supplemental_data"] = o.InterestPaidSupplementalData
	}
	if !IsNil(o.IncomeTaxPaidSupplementalData) {
		toSerialize["income_tax_paid_supplemental_data"] = o.IncomeTaxPaidSupplementalData
	}
	return toSerialize, nil
}

type NullableCashFlowDataSupplementalData struct {
	value *CashFlowDataSupplementalData
	isSet bool
}

func (v NullableCashFlowDataSupplementalData) Get() *CashFlowDataSupplementalData {
	return v.value
}

func (v *NullableCashFlowDataSupplementalData) Set(val *CashFlowDataSupplementalData) {
	v.value = val
	v.isSet = true
}

func (v NullableCashFlowDataSupplementalData) IsSet() bool {
	return v.isSet
}

func (v *NullableCashFlowDataSupplementalData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCashFlowDataSupplementalData(val *CashFlowDataSupplementalData) *NullableCashFlowDataSupplementalData {
	return &NullableCashFlowDataSupplementalData{value: val, isSet: true}
}

func (v NullableCashFlowDataSupplementalData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCashFlowDataSupplementalData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
