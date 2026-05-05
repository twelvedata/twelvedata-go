// Twelve Data API client for Go
//
// NOTE: This code is auto generated, please do not edit it manually.

package twelvedata

import (
	"encoding/json"
)

// checks if the CashFlowStructInvestingActivities type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CashFlowStructInvestingActivities{}

// CashFlowStructInvestingActivities Investing activities section
type CashFlowStructInvestingActivities struct {
	// Capital expenditures (CapEx) are funds used by a company to acquire, upgrade, and maintain physical assets (PPE)
	CapitalExpenditures *float64 `json:"capital_expenditures,omitempty"`
	// Represents purchase of a not physical asset
	NetIntangibles *float64 `json:"net_intangibles,omitempty"`
	// Refers to net amount of business purchase and sale
	NetAcquisitions *float64 `json:"net_acquisitions,omitempty"`
	// Represents how much money has been used in making investments, including purchases of physical assets, investments in securities
	PurchaseOfInvestments *float64 `json:"purchase_of_investments,omitempty"`
	// Represents how much money has been generated from the sale of securities or assets
	SaleOfInvestments *float64 `json:"sale_of_investments,omitempty"`
	// Represents other investing activity
	OtherInvestingActivity *float64 `json:"other_investing_activity,omitempty"`
	// Returns total amount of cash flow used in investments
	InvestingCashFlow *float64 `json:"investing_cash_flow,omitempty"`
}

// NewCashFlowStructInvestingActivities instantiates a new CashFlowStructInvestingActivities object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCashFlowStructInvestingActivities() *CashFlowStructInvestingActivities {
	this := CashFlowStructInvestingActivities{}
	return &this
}

// NewCashFlowStructInvestingActivitiesWithDefaults instantiates a new CashFlowStructInvestingActivities object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCashFlowStructInvestingActivitiesWithDefaults() *CashFlowStructInvestingActivities {
	this := CashFlowStructInvestingActivities{}
	return &this
}

// GetCapitalExpenditures returns the CapitalExpenditures field value if set, zero value otherwise.
func (o *CashFlowStructInvestingActivities) GetCapitalExpenditures() float64 {
	if o == nil || IsNil(o.CapitalExpenditures) {
		var ret float64
		return ret
	}
	return *o.CapitalExpenditures
}

// GetCapitalExpendituresOk returns a tuple with the CapitalExpenditures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowStructInvestingActivities) GetCapitalExpendituresOk() (*float64, bool) {
	if o == nil || IsNil(o.CapitalExpenditures) {
		return nil, false
	}
	return o.CapitalExpenditures, true
}

// HasCapitalExpenditures returns a boolean if a field has been set.
func (o *CashFlowStructInvestingActivities) HasCapitalExpenditures() bool {
	if o != nil && !IsNil(o.CapitalExpenditures) {
		return true
	}

	return false
}

// SetCapitalExpenditures gets a reference to the given float64 and assigns it to the CapitalExpenditures field.
func (o *CashFlowStructInvestingActivities) SetCapitalExpenditures(v float64) {
	o.CapitalExpenditures = &v
}

// GetNetIntangibles returns the NetIntangibles field value if set, zero value otherwise.
func (o *CashFlowStructInvestingActivities) GetNetIntangibles() float64 {
	if o == nil || IsNil(o.NetIntangibles) {
		var ret float64
		return ret
	}
	return *o.NetIntangibles
}

// GetNetIntangiblesOk returns a tuple with the NetIntangibles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowStructInvestingActivities) GetNetIntangiblesOk() (*float64, bool) {
	if o == nil || IsNil(o.NetIntangibles) {
		return nil, false
	}
	return o.NetIntangibles, true
}

// HasNetIntangibles returns a boolean if a field has been set.
func (o *CashFlowStructInvestingActivities) HasNetIntangibles() bool {
	if o != nil && !IsNil(o.NetIntangibles) {
		return true
	}

	return false
}

// SetNetIntangibles gets a reference to the given float64 and assigns it to the NetIntangibles field.
func (o *CashFlowStructInvestingActivities) SetNetIntangibles(v float64) {
	o.NetIntangibles = &v
}

// GetNetAcquisitions returns the NetAcquisitions field value if set, zero value otherwise.
func (o *CashFlowStructInvestingActivities) GetNetAcquisitions() float64 {
	if o == nil || IsNil(o.NetAcquisitions) {
		var ret float64
		return ret
	}
	return *o.NetAcquisitions
}

// GetNetAcquisitionsOk returns a tuple with the NetAcquisitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowStructInvestingActivities) GetNetAcquisitionsOk() (*float64, bool) {
	if o == nil || IsNil(o.NetAcquisitions) {
		return nil, false
	}
	return o.NetAcquisitions, true
}

// HasNetAcquisitions returns a boolean if a field has been set.
func (o *CashFlowStructInvestingActivities) HasNetAcquisitions() bool {
	if o != nil && !IsNil(o.NetAcquisitions) {
		return true
	}

	return false
}

// SetNetAcquisitions gets a reference to the given float64 and assigns it to the NetAcquisitions field.
func (o *CashFlowStructInvestingActivities) SetNetAcquisitions(v float64) {
	o.NetAcquisitions = &v
}

// GetPurchaseOfInvestments returns the PurchaseOfInvestments field value if set, zero value otherwise.
func (o *CashFlowStructInvestingActivities) GetPurchaseOfInvestments() float64 {
	if o == nil || IsNil(o.PurchaseOfInvestments) {
		var ret float64
		return ret
	}
	return *o.PurchaseOfInvestments
}

// GetPurchaseOfInvestmentsOk returns a tuple with the PurchaseOfInvestments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowStructInvestingActivities) GetPurchaseOfInvestmentsOk() (*float64, bool) {
	if o == nil || IsNil(o.PurchaseOfInvestments) {
		return nil, false
	}
	return o.PurchaseOfInvestments, true
}

// HasPurchaseOfInvestments returns a boolean if a field has been set.
func (o *CashFlowStructInvestingActivities) HasPurchaseOfInvestments() bool {
	if o != nil && !IsNil(o.PurchaseOfInvestments) {
		return true
	}

	return false
}

// SetPurchaseOfInvestments gets a reference to the given float64 and assigns it to the PurchaseOfInvestments field.
func (o *CashFlowStructInvestingActivities) SetPurchaseOfInvestments(v float64) {
	o.PurchaseOfInvestments = &v
}

// GetSaleOfInvestments returns the SaleOfInvestments field value if set, zero value otherwise.
func (o *CashFlowStructInvestingActivities) GetSaleOfInvestments() float64 {
	if o == nil || IsNil(o.SaleOfInvestments) {
		var ret float64
		return ret
	}
	return *o.SaleOfInvestments
}

// GetSaleOfInvestmentsOk returns a tuple with the SaleOfInvestments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowStructInvestingActivities) GetSaleOfInvestmentsOk() (*float64, bool) {
	if o == nil || IsNil(o.SaleOfInvestments) {
		return nil, false
	}
	return o.SaleOfInvestments, true
}

// HasSaleOfInvestments returns a boolean if a field has been set.
func (o *CashFlowStructInvestingActivities) HasSaleOfInvestments() bool {
	if o != nil && !IsNil(o.SaleOfInvestments) {
		return true
	}

	return false
}

// SetSaleOfInvestments gets a reference to the given float64 and assigns it to the SaleOfInvestments field.
func (o *CashFlowStructInvestingActivities) SetSaleOfInvestments(v float64) {
	o.SaleOfInvestments = &v
}

// GetOtherInvestingActivity returns the OtherInvestingActivity field value if set, zero value otherwise.
func (o *CashFlowStructInvestingActivities) GetOtherInvestingActivity() float64 {
	if o == nil || IsNil(o.OtherInvestingActivity) {
		var ret float64
		return ret
	}
	return *o.OtherInvestingActivity
}

// GetOtherInvestingActivityOk returns a tuple with the OtherInvestingActivity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowStructInvestingActivities) GetOtherInvestingActivityOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherInvestingActivity) {
		return nil, false
	}
	return o.OtherInvestingActivity, true
}

// HasOtherInvestingActivity returns a boolean if a field has been set.
func (o *CashFlowStructInvestingActivities) HasOtherInvestingActivity() bool {
	if o != nil && !IsNil(o.OtherInvestingActivity) {
		return true
	}

	return false
}

// SetOtherInvestingActivity gets a reference to the given float64 and assigns it to the OtherInvestingActivity field.
func (o *CashFlowStructInvestingActivities) SetOtherInvestingActivity(v float64) {
	o.OtherInvestingActivity = &v
}

// GetInvestingCashFlow returns the InvestingCashFlow field value if set, zero value otherwise.
func (o *CashFlowStructInvestingActivities) GetInvestingCashFlow() float64 {
	if o == nil || IsNil(o.InvestingCashFlow) {
		var ret float64
		return ret
	}
	return *o.InvestingCashFlow
}

// GetInvestingCashFlowOk returns a tuple with the InvestingCashFlow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowStructInvestingActivities) GetInvestingCashFlowOk() (*float64, bool) {
	if o == nil || IsNil(o.InvestingCashFlow) {
		return nil, false
	}
	return o.InvestingCashFlow, true
}

// HasInvestingCashFlow returns a boolean if a field has been set.
func (o *CashFlowStructInvestingActivities) HasInvestingCashFlow() bool {
	if o != nil && !IsNil(o.InvestingCashFlow) {
		return true
	}

	return false
}

// SetInvestingCashFlow gets a reference to the given float64 and assigns it to the InvestingCashFlow field.
func (o *CashFlowStructInvestingActivities) SetInvestingCashFlow(v float64) {
	o.InvestingCashFlow = &v
}

func (o CashFlowStructInvestingActivities) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CashFlowStructInvestingActivities) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CapitalExpenditures) {
		toSerialize["capital_expenditures"] = o.CapitalExpenditures
	}
	if !IsNil(o.NetIntangibles) {
		toSerialize["net_intangibles"] = o.NetIntangibles
	}
	if !IsNil(o.NetAcquisitions) {
		toSerialize["net_acquisitions"] = o.NetAcquisitions
	}
	if !IsNil(o.PurchaseOfInvestments) {
		toSerialize["purchase_of_investments"] = o.PurchaseOfInvestments
	}
	if !IsNil(o.SaleOfInvestments) {
		toSerialize["sale_of_investments"] = o.SaleOfInvestments
	}
	if !IsNil(o.OtherInvestingActivity) {
		toSerialize["other_investing_activity"] = o.OtherInvestingActivity
	}
	if !IsNil(o.InvestingCashFlow) {
		toSerialize["investing_cash_flow"] = o.InvestingCashFlow
	}
	return toSerialize, nil
}

type NullableCashFlowStructInvestingActivities struct {
	value *CashFlowStructInvestingActivities
	isSet bool
}

func (v NullableCashFlowStructInvestingActivities) Get() *CashFlowStructInvestingActivities {
	return v.value
}

func (v *NullableCashFlowStructInvestingActivities) Set(val *CashFlowStructInvestingActivities) {
	v.value = val
	v.isSet = true
}

func (v NullableCashFlowStructInvestingActivities) IsSet() bool {
	return v.isSet
}

func (v *NullableCashFlowStructInvestingActivities) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCashFlowStructInvestingActivities(val *CashFlowStructInvestingActivities) *NullableCashFlowStructInvestingActivities {
	return &NullableCashFlowStructInvestingActivities{value: val, isSet: true}
}

func (v NullableCashFlowStructInvestingActivities) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCashFlowStructInvestingActivities) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
