/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the CashFlowDataDirectMethodCashFlow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CashFlowDataDirectMethodCashFlow{}

// CashFlowDataDirectMethodCashFlow Direct method cash flow
type CashFlowDataDirectMethodCashFlow struct {
	// Classes of cash receipts from operating activities
	ClassesOfCashReceiptsFromOperatingActivities *float64 `json:"classes_of_cash_receipts_from_operating_activities,omitempty"`
	// Other cash receipts from operating activities
	OtherCashReceiptsFromOperatingActivities *float64 `json:"other_cash_receipts_from_operating_activities,omitempty"`
	// Receipts from government grants
	ReceiptsFromGovernmentGrants *float64 `json:"receipts_from_government_grants,omitempty"`
	// Receipts from customers
	ReceiptsFromCustomers *float64 `json:"receipts_from_customers,omitempty"`
	// Classes of cash payments
	ClassesOfCashPayments *float64 `json:"classes_of_cash_payments,omitempty"`
	// Other cash payments from operating activities
	OtherCashPaymentsFromOperatingActivities *float64 `json:"other_cash_payments_from_operating_activities,omitempty"`
	// Payments on behalf of employees
	PaymentsOnBehalfOfEmployees *float64 `json:"payments_on_behalf_of_employees,omitempty"`
	// Payments to suppliers for goods and services
	PaymentsToSuppliersForGoodsAndServices *float64 `json:"payments_to_suppliers_for_goods_and_services,omitempty"`
}

// NewCashFlowDataDirectMethodCashFlow instantiates a new CashFlowDataDirectMethodCashFlow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCashFlowDataDirectMethodCashFlow() *CashFlowDataDirectMethodCashFlow {
	this := CashFlowDataDirectMethodCashFlow{}
	return &this
}

// NewCashFlowDataDirectMethodCashFlowWithDefaults instantiates a new CashFlowDataDirectMethodCashFlow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCashFlowDataDirectMethodCashFlowWithDefaults() *CashFlowDataDirectMethodCashFlow {
	this := CashFlowDataDirectMethodCashFlow{}
	return &this
}

// GetClassesOfCashReceiptsFromOperatingActivities returns the ClassesOfCashReceiptsFromOperatingActivities field value if set, zero value otherwise.
func (o *CashFlowDataDirectMethodCashFlow) GetClassesOfCashReceiptsFromOperatingActivities() float64 {
	if o == nil || IsNil(o.ClassesOfCashReceiptsFromOperatingActivities) {
		var ret float64
		return ret
	}
	return *o.ClassesOfCashReceiptsFromOperatingActivities
}

// GetClassesOfCashReceiptsFromOperatingActivitiesOk returns a tuple with the ClassesOfCashReceiptsFromOperatingActivities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataDirectMethodCashFlow) GetClassesOfCashReceiptsFromOperatingActivitiesOk() (*float64, bool) {
	if o == nil || IsNil(o.ClassesOfCashReceiptsFromOperatingActivities) {
		return nil, false
	}
	return o.ClassesOfCashReceiptsFromOperatingActivities, true
}

// HasClassesOfCashReceiptsFromOperatingActivities returns a boolean if a field has been set.
func (o *CashFlowDataDirectMethodCashFlow) HasClassesOfCashReceiptsFromOperatingActivities() bool {
	if o != nil && !IsNil(o.ClassesOfCashReceiptsFromOperatingActivities) {
		return true
	}

	return false
}

// SetClassesOfCashReceiptsFromOperatingActivities gets a reference to the given float64 and assigns it to the ClassesOfCashReceiptsFromOperatingActivities field.
func (o *CashFlowDataDirectMethodCashFlow) SetClassesOfCashReceiptsFromOperatingActivities(v float64) {
	o.ClassesOfCashReceiptsFromOperatingActivities = &v
}

// GetOtherCashReceiptsFromOperatingActivities returns the OtherCashReceiptsFromOperatingActivities field value if set, zero value otherwise.
func (o *CashFlowDataDirectMethodCashFlow) GetOtherCashReceiptsFromOperatingActivities() float64 {
	if o == nil || IsNil(o.OtherCashReceiptsFromOperatingActivities) {
		var ret float64
		return ret
	}
	return *o.OtherCashReceiptsFromOperatingActivities
}

// GetOtherCashReceiptsFromOperatingActivitiesOk returns a tuple with the OtherCashReceiptsFromOperatingActivities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataDirectMethodCashFlow) GetOtherCashReceiptsFromOperatingActivitiesOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherCashReceiptsFromOperatingActivities) {
		return nil, false
	}
	return o.OtherCashReceiptsFromOperatingActivities, true
}

// HasOtherCashReceiptsFromOperatingActivities returns a boolean if a field has been set.
func (o *CashFlowDataDirectMethodCashFlow) HasOtherCashReceiptsFromOperatingActivities() bool {
	if o != nil && !IsNil(o.OtherCashReceiptsFromOperatingActivities) {
		return true
	}

	return false
}

// SetOtherCashReceiptsFromOperatingActivities gets a reference to the given float64 and assigns it to the OtherCashReceiptsFromOperatingActivities field.
func (o *CashFlowDataDirectMethodCashFlow) SetOtherCashReceiptsFromOperatingActivities(v float64) {
	o.OtherCashReceiptsFromOperatingActivities = &v
}

// GetReceiptsFromGovernmentGrants returns the ReceiptsFromGovernmentGrants field value if set, zero value otherwise.
func (o *CashFlowDataDirectMethodCashFlow) GetReceiptsFromGovernmentGrants() float64 {
	if o == nil || IsNil(o.ReceiptsFromGovernmentGrants) {
		var ret float64
		return ret
	}
	return *o.ReceiptsFromGovernmentGrants
}

// GetReceiptsFromGovernmentGrantsOk returns a tuple with the ReceiptsFromGovernmentGrants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataDirectMethodCashFlow) GetReceiptsFromGovernmentGrantsOk() (*float64, bool) {
	if o == nil || IsNil(o.ReceiptsFromGovernmentGrants) {
		return nil, false
	}
	return o.ReceiptsFromGovernmentGrants, true
}

// HasReceiptsFromGovernmentGrants returns a boolean if a field has been set.
func (o *CashFlowDataDirectMethodCashFlow) HasReceiptsFromGovernmentGrants() bool {
	if o != nil && !IsNil(o.ReceiptsFromGovernmentGrants) {
		return true
	}

	return false
}

// SetReceiptsFromGovernmentGrants gets a reference to the given float64 and assigns it to the ReceiptsFromGovernmentGrants field.
func (o *CashFlowDataDirectMethodCashFlow) SetReceiptsFromGovernmentGrants(v float64) {
	o.ReceiptsFromGovernmentGrants = &v
}

// GetReceiptsFromCustomers returns the ReceiptsFromCustomers field value if set, zero value otherwise.
func (o *CashFlowDataDirectMethodCashFlow) GetReceiptsFromCustomers() float64 {
	if o == nil || IsNil(o.ReceiptsFromCustomers) {
		var ret float64
		return ret
	}
	return *o.ReceiptsFromCustomers
}

// GetReceiptsFromCustomersOk returns a tuple with the ReceiptsFromCustomers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataDirectMethodCashFlow) GetReceiptsFromCustomersOk() (*float64, bool) {
	if o == nil || IsNil(o.ReceiptsFromCustomers) {
		return nil, false
	}
	return o.ReceiptsFromCustomers, true
}

// HasReceiptsFromCustomers returns a boolean if a field has been set.
func (o *CashFlowDataDirectMethodCashFlow) HasReceiptsFromCustomers() bool {
	if o != nil && !IsNil(o.ReceiptsFromCustomers) {
		return true
	}

	return false
}

// SetReceiptsFromCustomers gets a reference to the given float64 and assigns it to the ReceiptsFromCustomers field.
func (o *CashFlowDataDirectMethodCashFlow) SetReceiptsFromCustomers(v float64) {
	o.ReceiptsFromCustomers = &v
}

// GetClassesOfCashPayments returns the ClassesOfCashPayments field value if set, zero value otherwise.
func (o *CashFlowDataDirectMethodCashFlow) GetClassesOfCashPayments() float64 {
	if o == nil || IsNil(o.ClassesOfCashPayments) {
		var ret float64
		return ret
	}
	return *o.ClassesOfCashPayments
}

// GetClassesOfCashPaymentsOk returns a tuple with the ClassesOfCashPayments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataDirectMethodCashFlow) GetClassesOfCashPaymentsOk() (*float64, bool) {
	if o == nil || IsNil(o.ClassesOfCashPayments) {
		return nil, false
	}
	return o.ClassesOfCashPayments, true
}

// HasClassesOfCashPayments returns a boolean if a field has been set.
func (o *CashFlowDataDirectMethodCashFlow) HasClassesOfCashPayments() bool {
	if o != nil && !IsNil(o.ClassesOfCashPayments) {
		return true
	}

	return false
}

// SetClassesOfCashPayments gets a reference to the given float64 and assigns it to the ClassesOfCashPayments field.
func (o *CashFlowDataDirectMethodCashFlow) SetClassesOfCashPayments(v float64) {
	o.ClassesOfCashPayments = &v
}

// GetOtherCashPaymentsFromOperatingActivities returns the OtherCashPaymentsFromOperatingActivities field value if set, zero value otherwise.
func (o *CashFlowDataDirectMethodCashFlow) GetOtherCashPaymentsFromOperatingActivities() float64 {
	if o == nil || IsNil(o.OtherCashPaymentsFromOperatingActivities) {
		var ret float64
		return ret
	}
	return *o.OtherCashPaymentsFromOperatingActivities
}

// GetOtherCashPaymentsFromOperatingActivitiesOk returns a tuple with the OtherCashPaymentsFromOperatingActivities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataDirectMethodCashFlow) GetOtherCashPaymentsFromOperatingActivitiesOk() (*float64, bool) {
	if o == nil || IsNil(o.OtherCashPaymentsFromOperatingActivities) {
		return nil, false
	}
	return o.OtherCashPaymentsFromOperatingActivities, true
}

// HasOtherCashPaymentsFromOperatingActivities returns a boolean if a field has been set.
func (o *CashFlowDataDirectMethodCashFlow) HasOtherCashPaymentsFromOperatingActivities() bool {
	if o != nil && !IsNil(o.OtherCashPaymentsFromOperatingActivities) {
		return true
	}

	return false
}

// SetOtherCashPaymentsFromOperatingActivities gets a reference to the given float64 and assigns it to the OtherCashPaymentsFromOperatingActivities field.
func (o *CashFlowDataDirectMethodCashFlow) SetOtherCashPaymentsFromOperatingActivities(v float64) {
	o.OtherCashPaymentsFromOperatingActivities = &v
}

// GetPaymentsOnBehalfOfEmployees returns the PaymentsOnBehalfOfEmployees field value if set, zero value otherwise.
func (o *CashFlowDataDirectMethodCashFlow) GetPaymentsOnBehalfOfEmployees() float64 {
	if o == nil || IsNil(o.PaymentsOnBehalfOfEmployees) {
		var ret float64
		return ret
	}
	return *o.PaymentsOnBehalfOfEmployees
}

// GetPaymentsOnBehalfOfEmployeesOk returns a tuple with the PaymentsOnBehalfOfEmployees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataDirectMethodCashFlow) GetPaymentsOnBehalfOfEmployeesOk() (*float64, bool) {
	if o == nil || IsNil(o.PaymentsOnBehalfOfEmployees) {
		return nil, false
	}
	return o.PaymentsOnBehalfOfEmployees, true
}

// HasPaymentsOnBehalfOfEmployees returns a boolean if a field has been set.
func (o *CashFlowDataDirectMethodCashFlow) HasPaymentsOnBehalfOfEmployees() bool {
	if o != nil && !IsNil(o.PaymentsOnBehalfOfEmployees) {
		return true
	}

	return false
}

// SetPaymentsOnBehalfOfEmployees gets a reference to the given float64 and assigns it to the PaymentsOnBehalfOfEmployees field.
func (o *CashFlowDataDirectMethodCashFlow) SetPaymentsOnBehalfOfEmployees(v float64) {
	o.PaymentsOnBehalfOfEmployees = &v
}

// GetPaymentsToSuppliersForGoodsAndServices returns the PaymentsToSuppliersForGoodsAndServices field value if set, zero value otherwise.
func (o *CashFlowDataDirectMethodCashFlow) GetPaymentsToSuppliersForGoodsAndServices() float64 {
	if o == nil || IsNil(o.PaymentsToSuppliersForGoodsAndServices) {
		var ret float64
		return ret
	}
	return *o.PaymentsToSuppliersForGoodsAndServices
}

// GetPaymentsToSuppliersForGoodsAndServicesOk returns a tuple with the PaymentsToSuppliersForGoodsAndServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CashFlowDataDirectMethodCashFlow) GetPaymentsToSuppliersForGoodsAndServicesOk() (*float64, bool) {
	if o == nil || IsNil(o.PaymentsToSuppliersForGoodsAndServices) {
		return nil, false
	}
	return o.PaymentsToSuppliersForGoodsAndServices, true
}

// HasPaymentsToSuppliersForGoodsAndServices returns a boolean if a field has been set.
func (o *CashFlowDataDirectMethodCashFlow) HasPaymentsToSuppliersForGoodsAndServices() bool {
	if o != nil && !IsNil(o.PaymentsToSuppliersForGoodsAndServices) {
		return true
	}

	return false
}

// SetPaymentsToSuppliersForGoodsAndServices gets a reference to the given float64 and assigns it to the PaymentsToSuppliersForGoodsAndServices field.
func (o *CashFlowDataDirectMethodCashFlow) SetPaymentsToSuppliersForGoodsAndServices(v float64) {
	o.PaymentsToSuppliersForGoodsAndServices = &v
}

func (o CashFlowDataDirectMethodCashFlow) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CashFlowDataDirectMethodCashFlow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClassesOfCashReceiptsFromOperatingActivities) {
		toSerialize["classes_of_cash_receipts_from_operating_activities"] = o.ClassesOfCashReceiptsFromOperatingActivities
	}
	if !IsNil(o.OtherCashReceiptsFromOperatingActivities) {
		toSerialize["other_cash_receipts_from_operating_activities"] = o.OtherCashReceiptsFromOperatingActivities
	}
	if !IsNil(o.ReceiptsFromGovernmentGrants) {
		toSerialize["receipts_from_government_grants"] = o.ReceiptsFromGovernmentGrants
	}
	if !IsNil(o.ReceiptsFromCustomers) {
		toSerialize["receipts_from_customers"] = o.ReceiptsFromCustomers
	}
	if !IsNil(o.ClassesOfCashPayments) {
		toSerialize["classes_of_cash_payments"] = o.ClassesOfCashPayments
	}
	if !IsNil(o.OtherCashPaymentsFromOperatingActivities) {
		toSerialize["other_cash_payments_from_operating_activities"] = o.OtherCashPaymentsFromOperatingActivities
	}
	if !IsNil(o.PaymentsOnBehalfOfEmployees) {
		toSerialize["payments_on_behalf_of_employees"] = o.PaymentsOnBehalfOfEmployees
	}
	if !IsNil(o.PaymentsToSuppliersForGoodsAndServices) {
		toSerialize["payments_to_suppliers_for_goods_and_services"] = o.PaymentsToSuppliersForGoodsAndServices
	}
	return toSerialize, nil
}

type NullableCashFlowDataDirectMethodCashFlow struct {
	value *CashFlowDataDirectMethodCashFlow
	isSet bool
}

func (v NullableCashFlowDataDirectMethodCashFlow) Get() *CashFlowDataDirectMethodCashFlow {
	return v.value
}

func (v *NullableCashFlowDataDirectMethodCashFlow) Set(val *CashFlowDataDirectMethodCashFlow) {
	v.value = val
	v.isSet = true
}

func (v NullableCashFlowDataDirectMethodCashFlow) IsSet() bool {
	return v.isSet
}

func (v *NullableCashFlowDataDirectMethodCashFlow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCashFlowDataDirectMethodCashFlow(val *CashFlowDataDirectMethodCashFlow) *NullableCashFlowDataDirectMethodCashFlow {
	return &NullableCashFlowDataDirectMethodCashFlow{value: val, isSet: true}
}

func (v NullableCashFlowDataDirectMethodCashFlow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCashFlowDataDirectMethodCashFlow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
