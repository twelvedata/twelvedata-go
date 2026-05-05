// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities{}

// GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities Non-current liabilities section
type GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities struct {
	// Represents money set aside for employee benefits such as gratuity
	LongTermProvisions *float64 `json:"long_term_provisions,omitempty"`
	// Represents amount of outstanding debt that has a maturity of 12 months or longer
	LongTermDebt *float64 `json:"long_term_debt,omitempty"`
	// Represents funds set aside as assets to pay for anticipated future losses
	ProvisionForRisksAndCharges *float64 `json:"provision_for_risks_and_charges,omitempty"`
	// Represents revenue producing activity for which revenue has not yet been recognized, and is not expected to be recognized in the next 12 months
	DeferredLiabilities *float64 `json:"deferred_liabilities,omitempty"`
	// Represents the value of derivative financial instruments that a company has issued
	DerivativeProductLiabilities *float64 `json:"derivative_product_liabilities,omitempty"`
	// Represents other non-current liabilities
	OtherNonCurrentLiabilities *float64 `json:"other_non_current_liabilities,omitempty"`
	// Represents total non-current liabilities
	TotalNonCurrentLiabilities *float64 `json:"total_non_current_liabilities,omitempty"`
}

// NewGetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities instantiates a new GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities() *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities {
	this := GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities{}
	return &this
}

// NewGetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilitiesWithDefaults instantiates a new GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilitiesWithDefaults() *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities {
	this := GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities{}
	return &this
}

// GetLongTermProvisions returns the LongTermProvisions field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) GetLongTermProvisions() float64 {
	if o == nil || IsNil(o.LongTermProvisions) {
		var ret float64
		return ret
	}
	return *o.LongTermProvisions
}

// GetLongTermProvisionsOk returns a tuple with the LongTermProvisions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) GetLongTermProvisionsOk() (*float64, bool) {
	if o == nil || IsNil(o.LongTermProvisions) {
		return nil, false
	}
	return o.LongTermProvisions, true
}

// HasLongTermProvisions returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) HasLongTermProvisions() bool {
	if o != nil && !IsNil(o.LongTermProvisions) {
		return true
	}

	return false
}

// SetLongTermProvisions gets a reference to the given float64 and assigns it to the LongTermProvisions field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) SetLongTermProvisions(v float64) {
	o.LongTermProvisions = &v
}

// GetLongTermDebt returns the LongTermDebt field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) GetLongTermDebt() float64 {
	if o == nil || IsNil(o.LongTermDebt) {
		var ret float64
		return ret
	}
	return *o.LongTermDebt
}

// GetLongTermDebtOk returns a tuple with the LongTermDebt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) GetLongTermDebtOk() (*float64, bool) {
	if o == nil || IsNil(o.LongTermDebt) {
		return nil, false
	}
	return o.LongTermDebt, true
}

// HasLongTermDebt returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) HasLongTermDebt() bool {
	if o != nil && !IsNil(o.LongTermDebt) {
		return true
	}

	return false
}

// SetLongTermDebt gets a reference to the given float64 and assigns it to the LongTermDebt field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) SetLongTermDebt(v float64) {
	o.LongTermDebt = &v
}

// GetProvisionForRisksAndCharges returns the ProvisionForRisksAndCharges field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) GetProvisionForRisksAndCharges() float64 {
	if o == nil || IsNil(o.ProvisionForRisksAndCharges) {
		var ret float64
		return ret
	}
	return *o.ProvisionForRisksAndCharges
}

// GetProvisionForRisksAndChargesOk returns a tuple with the ProvisionForRisksAndCharges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) GetProvisionForRisksAndChargesOk() (*float64, bool) {
	if o == nil || IsNil(o.ProvisionForRisksAndCharges) {
		return nil, false
	}
	return o.ProvisionForRisksAndCharges, true
}

// HasProvisionForRisksAndCharges returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) HasProvisionForRisksAndCharges() bool {
	if o != nil && !IsNil(o.ProvisionForRisksAndCharges) {
		return true
	}

	return false
}

// SetProvisionForRisksAndCharges gets a reference to the given float64 and assigns it to the ProvisionForRisksAndCharges field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) SetProvisionForRisksAndCharges(v float64) {
	o.ProvisionForRisksAndCharges = &v
}

// GetDeferredLiabilities returns the DeferredLiabilities field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) GetDeferredLiabilities() float64 {
	if o == nil || IsNil(o.DeferredLiabilities) {
		var ret float64
		return ret
	}
	return *o.DeferredLiabilities
}

// GetDeferredLiabilitiesOk returns a tuple with the DeferredLiabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) GetDeferredLiabilitiesOk() (*float64, bool) {
	if o == nil || IsNil(o.DeferredLiabilities) {
		return nil, false
	}
	return o.DeferredLiabilities, true
}

// HasDeferredLiabilities returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) HasDeferredLiabilities() bool {
	if o != nil && !IsNil(o.DeferredLiabilities) {
		return true
	}

	return false
}

// SetDeferredLiabilities gets a reference to the given float64 and assigns it to the DeferredLiabilities field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) SetDeferredLiabilities(v float64) {
	o.DeferredLiabilities = &v
}

// GetDerivativeProductLiabilities returns the DerivativeProductLiabilities field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) GetDerivativeProductLiabilities() float64 {
	if o == nil || IsNil(o.DerivativeProductLiabilities) {
		var ret float64
		return ret
	}
	return *o.DerivativeProductLiabilities
}

// GetDerivativeProductLiabilitiesOk returns a tuple with the DerivativeProductLiabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) GetDerivativeProductLiabilitiesOk() (*float64, bool) {
	if o == nil || IsNil(o.DerivativeProductLiabilities) {
		return nil, false
	}
	return o.DerivativeProductLiabilities, true
}

// HasDerivativeProductLiabilities returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) HasDerivativeProductLiabilities() bool {
	if o != nil && !IsNil(o.DerivativeProductLiabilities) {
		return true
	}

	return false
}

// SetDerivativeProductLiabilities gets a reference to the given float64 and assigns it to the DerivativeProductLiabilities field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) SetDerivativeProductLiabilities(v float64) {
	o.DerivativeProductLiabilities = &v
}

// GetOtherNonCurrentLiabilities returns the OtherNonCurrentLiabilities field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) GetOtherNonCurrentLiabilities() float64 {
	if o == nil || IsNil(o.OtherNonCurrentLiabilities) {
		var ret float64
		return ret
	}
	return *o.OtherNonCurrentLiabilities
}

// GetOtherNonCurrentLiabilitiesOk returns a tuple with the OtherNonCurrentLiabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) GetOtherNonCurrentLiabilitiesOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherNonCurrentLiabilities) {
		return nil, false
	}
	return o.OtherNonCurrentLiabilities, true
}

// HasOtherNonCurrentLiabilities returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) HasOtherNonCurrentLiabilities() bool {
	if o != nil && !IsNil(o.OtherNonCurrentLiabilities) {
		return true
	}

	return false
}

// SetOtherNonCurrentLiabilities gets a reference to the given float64 and assigns it to the OtherNonCurrentLiabilities field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) SetOtherNonCurrentLiabilities(v float64) {
	o.OtherNonCurrentLiabilities = &v
}

// GetTotalNonCurrentLiabilities returns the TotalNonCurrentLiabilities field value if set, zero value otherwise.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) GetTotalNonCurrentLiabilities() float64 {
	if o == nil || IsNil(o.TotalNonCurrentLiabilities) {
		var ret float64
		return ret
	}
	return *o.TotalNonCurrentLiabilities
}

// GetTotalNonCurrentLiabilitiesOk returns a tuple with the TotalNonCurrentLiabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) GetTotalNonCurrentLiabilitiesOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalNonCurrentLiabilities) {
		return nil, false
	}
	return o.TotalNonCurrentLiabilities, true
}

// HasTotalNonCurrentLiabilities returns a boolean if a field has been set.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) HasTotalNonCurrentLiabilities() bool {
	if o != nil && !IsNil(o.TotalNonCurrentLiabilities) {
		return true
	}

	return false
}

// SetTotalNonCurrentLiabilities gets a reference to the given float64 and assigns it to the TotalNonCurrentLiabilities field.
func (o *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) SetTotalNonCurrentLiabilities(v float64) {
	o.TotalNonCurrentLiabilities = &v
}

func (o GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LongTermProvisions) {
		toSerialize["long_term_provisions"] = o.LongTermProvisions
	}
	if !IsNil(o.LongTermDebt) {
		toSerialize["long_term_debt"] = o.LongTermDebt
	}
	if !IsNil(o.ProvisionForRisksAndCharges) {
		toSerialize["provision_for_risks_and_charges"] = o.ProvisionForRisksAndCharges
	}
	if !IsNil(o.DeferredLiabilities) {
		toSerialize["deferred_liabilities"] = o.DeferredLiabilities
	}
	if !IsNil(o.DerivativeProductLiabilities) {
		toSerialize["derivative_product_liabilities"] = o.DerivativeProductLiabilities
	}
	if !IsNil(o.OtherNonCurrentLiabilities) {
		toSerialize["other_non_current_liabilities"] = o.OtherNonCurrentLiabilities
	}
	if !IsNil(o.TotalNonCurrentLiabilities) {
		toSerialize["total_non_current_liabilities"] = o.TotalNonCurrentLiabilities
	}
	return toSerialize, nil
}

type NullableGetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities struct {
	value *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities
	isSet bool
}

func (v NullableGetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) Get() *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities {
	return v.value
}

func (v *NullableGetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) Set(val *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities(val *GetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) *NullableGetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities {
	return &NullableGetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities{value: val, isSet: true}
}

func (v NullableGetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBalanceSheet200ResponseBalanceSheetInnerLiabilitiesNonCurrentLiabilities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
