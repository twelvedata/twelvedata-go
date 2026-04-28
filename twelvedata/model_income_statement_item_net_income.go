/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the IncomeStatementItemNetIncome type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IncomeStatementItemNetIncome{}

// IncomeStatementItemNetIncome Net income information
type IncomeStatementItemNetIncome struct {
	// Net income value
	NetIncomeValue *float64 `json:"net_income_value,omitempty"`
	// Net income common stockholders
	NetIncomeCommonStockholders *float64 `json:"net_income_common_stockholders,omitempty"`
	// Net income including noncontrolling interests
	NetIncomeIncludingNoncontrollingInterests *float64 `json:"net_income_including_noncontrolling_interests,omitempty"`
	// Net income from tax loss carryforward
	NetIncomeFromTaxLossCarryforward *float64 `json:"net_income_from_tax_loss_carryforward,omitempty"`
	// Net income extraordinary
	NetIncomeExtraordinary *float64 `json:"net_income_extraordinary,omitempty"`
	// Net income discontinuous operations
	NetIncomeDiscontinuousOperations *float64 `json:"net_income_discontinuous_operations,omitempty"`
	// Net income continuous operations
	NetIncomeContinuousOperations *float64 `json:"net_income_continuous_operations,omitempty"`
	// Net income from continuing operation net minority interest
	NetIncomeFromContinuingOperationNetMinorityInterest *float64 `json:"net_income_from_continuing_operation_net_minority_interest,omitempty"`
	// Net income from continuing and discontinued operation
	NetIncomeFromContinuingAndDiscontinuedOperation *float64 `json:"net_income_from_continuing_and_discontinued_operation,omitempty"`
	// Normalized income
	NormalizedIncome *float64 `json:"normalized_income,omitempty"`
	// Minority interests
	MinorityInterests *float64 `json:"minority_interests,omitempty"`
}

// NewIncomeStatementItemNetIncome instantiates a new IncomeStatementItemNetIncome object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncomeStatementItemNetIncome() *IncomeStatementItemNetIncome {
	this := IncomeStatementItemNetIncome{}
	return &this
}

// NewIncomeStatementItemNetIncomeWithDefaults instantiates a new IncomeStatementItemNetIncome object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncomeStatementItemNetIncomeWithDefaults() *IncomeStatementItemNetIncome {
	this := IncomeStatementItemNetIncome{}
	return &this
}

// GetNetIncomeValue returns the NetIncomeValue field value if set, zero value otherwise.
func (o *IncomeStatementItemNetIncome) GetNetIncomeValue() float64 {
	if o == nil || IsNil(o.NetIncomeValue) {
		var ret float64
		return ret
	}
	return *o.NetIncomeValue
}

// GetNetIncomeValueOk returns a tuple with the NetIncomeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemNetIncome) GetNetIncomeValueOk() (*float64, bool) {
	if o == nil || IsNil(o.NetIncomeValue) {
		return nil, false
	}
	return o.NetIncomeValue, true
}

// HasNetIncomeValue returns a boolean if a field has been set.
func (o *IncomeStatementItemNetIncome) HasNetIncomeValue() bool {
	if o != nil && !IsNil(o.NetIncomeValue) {
		return true
	}

	return false
}

// SetNetIncomeValue gets a reference to the given float64 and assigns it to the NetIncomeValue field.
func (o *IncomeStatementItemNetIncome) SetNetIncomeValue(v float64) {
	o.NetIncomeValue = &v
}

// GetNetIncomeCommonStockholders returns the NetIncomeCommonStockholders field value if set, zero value otherwise.
func (o *IncomeStatementItemNetIncome) GetNetIncomeCommonStockholders() float64 {
	if o == nil || IsNil(o.NetIncomeCommonStockholders) {
		var ret float64
		return ret
	}
	return *o.NetIncomeCommonStockholders
}

// GetNetIncomeCommonStockholdersOk returns a tuple with the NetIncomeCommonStockholders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemNetIncome) GetNetIncomeCommonStockholdersOk() (*float64, bool) {
	if o == nil || IsNil(o.NetIncomeCommonStockholders) {
		return nil, false
	}
	return o.NetIncomeCommonStockholders, true
}

// HasNetIncomeCommonStockholders returns a boolean if a field has been set.
func (o *IncomeStatementItemNetIncome) HasNetIncomeCommonStockholders() bool {
	if o != nil && !IsNil(o.NetIncomeCommonStockholders) {
		return true
	}

	return false
}

// SetNetIncomeCommonStockholders gets a reference to the given float64 and assigns it to the NetIncomeCommonStockholders field.
func (o *IncomeStatementItemNetIncome) SetNetIncomeCommonStockholders(v float64) {
	o.NetIncomeCommonStockholders = &v
}

// GetNetIncomeIncludingNoncontrollingInterests returns the NetIncomeIncludingNoncontrollingInterests field value if set, zero value otherwise.
func (o *IncomeStatementItemNetIncome) GetNetIncomeIncludingNoncontrollingInterests() float64 {
	if o == nil || IsNil(o.NetIncomeIncludingNoncontrollingInterests) {
		var ret float64
		return ret
	}
	return *o.NetIncomeIncludingNoncontrollingInterests
}

// GetNetIncomeIncludingNoncontrollingInterestsOk returns a tuple with the NetIncomeIncludingNoncontrollingInterests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemNetIncome) GetNetIncomeIncludingNoncontrollingInterestsOk() (*float64, bool) {
	if o == nil || IsNil(o.NetIncomeIncludingNoncontrollingInterests) {
		return nil, false
	}
	return o.NetIncomeIncludingNoncontrollingInterests, true
}

// HasNetIncomeIncludingNoncontrollingInterests returns a boolean if a field has been set.
func (o *IncomeStatementItemNetIncome) HasNetIncomeIncludingNoncontrollingInterests() bool {
	if o != nil && !IsNil(o.NetIncomeIncludingNoncontrollingInterests) {
		return true
	}

	return false
}

// SetNetIncomeIncludingNoncontrollingInterests gets a reference to the given float64 and assigns it to the NetIncomeIncludingNoncontrollingInterests field.
func (o *IncomeStatementItemNetIncome) SetNetIncomeIncludingNoncontrollingInterests(v float64) {
	o.NetIncomeIncludingNoncontrollingInterests = &v
}

// GetNetIncomeFromTaxLossCarryforward returns the NetIncomeFromTaxLossCarryforward field value if set, zero value otherwise.
func (o *IncomeStatementItemNetIncome) GetNetIncomeFromTaxLossCarryforward() float64 {
	if o == nil || IsNil(o.NetIncomeFromTaxLossCarryforward) {
		var ret float64
		return ret
	}
	return *o.NetIncomeFromTaxLossCarryforward
}

// GetNetIncomeFromTaxLossCarryforwardOk returns a tuple with the NetIncomeFromTaxLossCarryforward field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemNetIncome) GetNetIncomeFromTaxLossCarryforwardOk() (*float64, bool) {
	if o == nil || IsNil(o.NetIncomeFromTaxLossCarryforward) {
		return nil, false
	}
	return o.NetIncomeFromTaxLossCarryforward, true
}

// HasNetIncomeFromTaxLossCarryforward returns a boolean if a field has been set.
func (o *IncomeStatementItemNetIncome) HasNetIncomeFromTaxLossCarryforward() bool {
	if o != nil && !IsNil(o.NetIncomeFromTaxLossCarryforward) {
		return true
	}

	return false
}

// SetNetIncomeFromTaxLossCarryforward gets a reference to the given float64 and assigns it to the NetIncomeFromTaxLossCarryforward field.
func (o *IncomeStatementItemNetIncome) SetNetIncomeFromTaxLossCarryforward(v float64) {
	o.NetIncomeFromTaxLossCarryforward = &v
}

// GetNetIncomeExtraordinary returns the NetIncomeExtraordinary field value if set, zero value otherwise.
func (o *IncomeStatementItemNetIncome) GetNetIncomeExtraordinary() float64 {
	if o == nil || IsNil(o.NetIncomeExtraordinary) {
		var ret float64
		return ret
	}
	return *o.NetIncomeExtraordinary
}

// GetNetIncomeExtraordinaryOk returns a tuple with the NetIncomeExtraordinary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemNetIncome) GetNetIncomeExtraordinaryOk() (*float64, bool) {
	if o == nil || IsNil(o.NetIncomeExtraordinary) {
		return nil, false
	}
	return o.NetIncomeExtraordinary, true
}

// HasNetIncomeExtraordinary returns a boolean if a field has been set.
func (o *IncomeStatementItemNetIncome) HasNetIncomeExtraordinary() bool {
	if o != nil && !IsNil(o.NetIncomeExtraordinary) {
		return true
	}

	return false
}

// SetNetIncomeExtraordinary gets a reference to the given float64 and assigns it to the NetIncomeExtraordinary field.
func (o *IncomeStatementItemNetIncome) SetNetIncomeExtraordinary(v float64) {
	o.NetIncomeExtraordinary = &v
}

// GetNetIncomeDiscontinuousOperations returns the NetIncomeDiscontinuousOperations field value if set, zero value otherwise.
func (o *IncomeStatementItemNetIncome) GetNetIncomeDiscontinuousOperations() float64 {
	if o == nil || IsNil(o.NetIncomeDiscontinuousOperations) {
		var ret float64
		return ret
	}
	return *o.NetIncomeDiscontinuousOperations
}

// GetNetIncomeDiscontinuousOperationsOk returns a tuple with the NetIncomeDiscontinuousOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemNetIncome) GetNetIncomeDiscontinuousOperationsOk() (*float64, bool) {
	if o == nil || IsNil(o.NetIncomeDiscontinuousOperations) {
		return nil, false
	}
	return o.NetIncomeDiscontinuousOperations, true
}

// HasNetIncomeDiscontinuousOperations returns a boolean if a field has been set.
func (o *IncomeStatementItemNetIncome) HasNetIncomeDiscontinuousOperations() bool {
	if o != nil && !IsNil(o.NetIncomeDiscontinuousOperations) {
		return true
	}

	return false
}

// SetNetIncomeDiscontinuousOperations gets a reference to the given float64 and assigns it to the NetIncomeDiscontinuousOperations field.
func (o *IncomeStatementItemNetIncome) SetNetIncomeDiscontinuousOperations(v float64) {
	o.NetIncomeDiscontinuousOperations = &v
}

// GetNetIncomeContinuousOperations returns the NetIncomeContinuousOperations field value if set, zero value otherwise.
func (o *IncomeStatementItemNetIncome) GetNetIncomeContinuousOperations() float64 {
	if o == nil || IsNil(o.NetIncomeContinuousOperations) {
		var ret float64
		return ret
	}
	return *o.NetIncomeContinuousOperations
}

// GetNetIncomeContinuousOperationsOk returns a tuple with the NetIncomeContinuousOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemNetIncome) GetNetIncomeContinuousOperationsOk() (*float64, bool) {
	if o == nil || IsNil(o.NetIncomeContinuousOperations) {
		return nil, false
	}
	return o.NetIncomeContinuousOperations, true
}

// HasNetIncomeContinuousOperations returns a boolean if a field has been set.
func (o *IncomeStatementItemNetIncome) HasNetIncomeContinuousOperations() bool {
	if o != nil && !IsNil(o.NetIncomeContinuousOperations) {
		return true
	}

	return false
}

// SetNetIncomeContinuousOperations gets a reference to the given float64 and assigns it to the NetIncomeContinuousOperations field.
func (o *IncomeStatementItemNetIncome) SetNetIncomeContinuousOperations(v float64) {
	o.NetIncomeContinuousOperations = &v
}

// GetNetIncomeFromContinuingOperationNetMinorityInterest returns the NetIncomeFromContinuingOperationNetMinorityInterest field value if set, zero value otherwise.
func (o *IncomeStatementItemNetIncome) GetNetIncomeFromContinuingOperationNetMinorityInterest() float64 {
	if o == nil || IsNil(o.NetIncomeFromContinuingOperationNetMinorityInterest) {
		var ret float64
		return ret
	}
	return *o.NetIncomeFromContinuingOperationNetMinorityInterest
}

// GetNetIncomeFromContinuingOperationNetMinorityInterestOk returns a tuple with the NetIncomeFromContinuingOperationNetMinorityInterest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemNetIncome) GetNetIncomeFromContinuingOperationNetMinorityInterestOk() (*float64, bool) {
	if o == nil || IsNil(o.NetIncomeFromContinuingOperationNetMinorityInterest) {
		return nil, false
	}
	return o.NetIncomeFromContinuingOperationNetMinorityInterest, true
}

// HasNetIncomeFromContinuingOperationNetMinorityInterest returns a boolean if a field has been set.
func (o *IncomeStatementItemNetIncome) HasNetIncomeFromContinuingOperationNetMinorityInterest() bool {
	if o != nil && !IsNil(o.NetIncomeFromContinuingOperationNetMinorityInterest) {
		return true
	}

	return false
}

// SetNetIncomeFromContinuingOperationNetMinorityInterest gets a reference to the given float64 and assigns it to the NetIncomeFromContinuingOperationNetMinorityInterest field.
func (o *IncomeStatementItemNetIncome) SetNetIncomeFromContinuingOperationNetMinorityInterest(v float64) {
	o.NetIncomeFromContinuingOperationNetMinorityInterest = &v
}

// GetNetIncomeFromContinuingAndDiscontinuedOperation returns the NetIncomeFromContinuingAndDiscontinuedOperation field value if set, zero value otherwise.
func (o *IncomeStatementItemNetIncome) GetNetIncomeFromContinuingAndDiscontinuedOperation() float64 {
	if o == nil || IsNil(o.NetIncomeFromContinuingAndDiscontinuedOperation) {
		var ret float64
		return ret
	}
	return *o.NetIncomeFromContinuingAndDiscontinuedOperation
}

// GetNetIncomeFromContinuingAndDiscontinuedOperationOk returns a tuple with the NetIncomeFromContinuingAndDiscontinuedOperation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemNetIncome) GetNetIncomeFromContinuingAndDiscontinuedOperationOk() (*float64, bool) {
	if o == nil || IsNil(o.NetIncomeFromContinuingAndDiscontinuedOperation) {
		return nil, false
	}
	return o.NetIncomeFromContinuingAndDiscontinuedOperation, true
}

// HasNetIncomeFromContinuingAndDiscontinuedOperation returns a boolean if a field has been set.
func (o *IncomeStatementItemNetIncome) HasNetIncomeFromContinuingAndDiscontinuedOperation() bool {
	if o != nil && !IsNil(o.NetIncomeFromContinuingAndDiscontinuedOperation) {
		return true
	}

	return false
}

// SetNetIncomeFromContinuingAndDiscontinuedOperation gets a reference to the given float64 and assigns it to the NetIncomeFromContinuingAndDiscontinuedOperation field.
func (o *IncomeStatementItemNetIncome) SetNetIncomeFromContinuingAndDiscontinuedOperation(v float64) {
	o.NetIncomeFromContinuingAndDiscontinuedOperation = &v
}

// GetNormalizedIncome returns the NormalizedIncome field value if set, zero value otherwise.
func (o *IncomeStatementItemNetIncome) GetNormalizedIncome() float64 {
	if o == nil || IsNil(o.NormalizedIncome) {
		var ret float64
		return ret
	}
	return *o.NormalizedIncome
}

// GetNormalizedIncomeOk returns a tuple with the NormalizedIncome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemNetIncome) GetNormalizedIncomeOk() (*float64, bool) {
	if o == nil || IsNil(o.NormalizedIncome) {
		return nil, false
	}
	return o.NormalizedIncome, true
}

// HasNormalizedIncome returns a boolean if a field has been set.
func (o *IncomeStatementItemNetIncome) HasNormalizedIncome() bool {
	if o != nil && !IsNil(o.NormalizedIncome) {
		return true
	}

	return false
}

// SetNormalizedIncome gets a reference to the given float64 and assigns it to the NormalizedIncome field.
func (o *IncomeStatementItemNetIncome) SetNormalizedIncome(v float64) {
	o.NormalizedIncome = &v
}

// GetMinorityInterests returns the MinorityInterests field value if set, zero value otherwise.
func (o *IncomeStatementItemNetIncome) GetMinorityInterests() float64 {
	if o == nil || IsNil(o.MinorityInterests) {
		var ret float64
		return ret
	}
	return *o.MinorityInterests
}

// GetMinorityInterestsOk returns a tuple with the MinorityInterests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncomeStatementItemNetIncome) GetMinorityInterestsOk() (*float64, bool) {
	if o == nil || IsNil(o.MinorityInterests) {
		return nil, false
	}
	return o.MinorityInterests, true
}

// HasMinorityInterests returns a boolean if a field has been set.
func (o *IncomeStatementItemNetIncome) HasMinorityInterests() bool {
	if o != nil && !IsNil(o.MinorityInterests) {
		return true
	}

	return false
}

// SetMinorityInterests gets a reference to the given float64 and assigns it to the MinorityInterests field.
func (o *IncomeStatementItemNetIncome) SetMinorityInterests(v float64) {
	o.MinorityInterests = &v
}

func (o IncomeStatementItemNetIncome) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IncomeStatementItemNetIncome) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NetIncomeValue) {
		toSerialize["net_income_value"] = o.NetIncomeValue
	}
	if !IsNil(o.NetIncomeCommonStockholders) {
		toSerialize["net_income_common_stockholders"] = o.NetIncomeCommonStockholders
	}
	if !IsNil(o.NetIncomeIncludingNoncontrollingInterests) {
		toSerialize["net_income_including_noncontrolling_interests"] = o.NetIncomeIncludingNoncontrollingInterests
	}
	if !IsNil(o.NetIncomeFromTaxLossCarryforward) {
		toSerialize["net_income_from_tax_loss_carryforward"] = o.NetIncomeFromTaxLossCarryforward
	}
	if !IsNil(o.NetIncomeExtraordinary) {
		toSerialize["net_income_extraordinary"] = o.NetIncomeExtraordinary
	}
	if !IsNil(o.NetIncomeDiscontinuousOperations) {
		toSerialize["net_income_discontinuous_operations"] = o.NetIncomeDiscontinuousOperations
	}
	if !IsNil(o.NetIncomeContinuousOperations) {
		toSerialize["net_income_continuous_operations"] = o.NetIncomeContinuousOperations
	}
	if !IsNil(o.NetIncomeFromContinuingOperationNetMinorityInterest) {
		toSerialize["net_income_from_continuing_operation_net_minority_interest"] = o.NetIncomeFromContinuingOperationNetMinorityInterest
	}
	if !IsNil(o.NetIncomeFromContinuingAndDiscontinuedOperation) {
		toSerialize["net_income_from_continuing_and_discontinued_operation"] = o.NetIncomeFromContinuingAndDiscontinuedOperation
	}
	if !IsNil(o.NormalizedIncome) {
		toSerialize["normalized_income"] = o.NormalizedIncome
	}
	if !IsNil(o.MinorityInterests) {
		toSerialize["minority_interests"] = o.MinorityInterests
	}
	return toSerialize, nil
}

type NullableIncomeStatementItemNetIncome struct {
	value *IncomeStatementItemNetIncome
	isSet bool
}

func (v NullableIncomeStatementItemNetIncome) Get() *IncomeStatementItemNetIncome {
	return v.value
}

func (v *NullableIncomeStatementItemNetIncome) Set(val *IncomeStatementItemNetIncome) {
	v.value = val
	v.isSet = true
}

func (v NullableIncomeStatementItemNetIncome) IsSet() bool {
	return v.isSet
}

func (v *NullableIncomeStatementItemNetIncome) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncomeStatementItemNetIncome(val *IncomeStatementItemNetIncome) *NullableIncomeStatementItemNetIncome {
	return &NullableIncomeStatementItemNetIncome{value: val, isSet: true}
}

func (v NullableIncomeStatementItemNetIncome) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncomeStatementItemNetIncome) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
