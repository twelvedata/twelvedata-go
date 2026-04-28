/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the InlineObject2Expenses type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject2Expenses{}

// InlineObject2Expenses struct for InlineObject2Expenses
type InlineObject2Expenses struct {
	ExpenseRatioGross *float64 `json:"expense_ratio_gross,omitempty"`
	ExpenseRatioNet   *float64 `json:"expense_ratio_net,omitempty"`
}

// NewInlineObject2Expenses instantiates a new InlineObject2Expenses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject2Expenses() *InlineObject2Expenses {
	this := InlineObject2Expenses{}
	return &this
}

// NewInlineObject2ExpensesWithDefaults instantiates a new InlineObject2Expenses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject2ExpensesWithDefaults() *InlineObject2Expenses {
	this := InlineObject2Expenses{}
	return &this
}

// GetExpenseRatioGross returns the ExpenseRatioGross field value if set, zero value otherwise.
func (o *InlineObject2Expenses) GetExpenseRatioGross() float64 {
	if o == nil || IsNil(o.ExpenseRatioGross) {
		var ret float64
		return ret
	}
	return *o.ExpenseRatioGross
}

// GetExpenseRatioGrossOk returns a tuple with the ExpenseRatioGross field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2Expenses) GetExpenseRatioGrossOk() (*float64, bool) {
	if o == nil || IsNil(o.ExpenseRatioGross) {
		return nil, false
	}
	return o.ExpenseRatioGross, true
}

// HasExpenseRatioGross returns a boolean if a field has been set.
func (o *InlineObject2Expenses) HasExpenseRatioGross() bool {
	if o != nil && !IsNil(o.ExpenseRatioGross) {
		return true
	}

	return false
}

// SetExpenseRatioGross gets a reference to the given float64 and assigns it to the ExpenseRatioGross field.
func (o *InlineObject2Expenses) SetExpenseRatioGross(v float64) {
	o.ExpenseRatioGross = &v
}

// GetExpenseRatioNet returns the ExpenseRatioNet field value if set, zero value otherwise.
func (o *InlineObject2Expenses) GetExpenseRatioNet() float64 {
	if o == nil || IsNil(o.ExpenseRatioNet) {
		var ret float64
		return ret
	}
	return *o.ExpenseRatioNet
}

// GetExpenseRatioNetOk returns a tuple with the ExpenseRatioNet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2Expenses) GetExpenseRatioNetOk() (*float64, bool) {
	if o == nil || IsNil(o.ExpenseRatioNet) {
		return nil, false
	}
	return o.ExpenseRatioNet, true
}

// HasExpenseRatioNet returns a boolean if a field has been set.
func (o *InlineObject2Expenses) HasExpenseRatioNet() bool {
	if o != nil && !IsNil(o.ExpenseRatioNet) {
		return true
	}

	return false
}

// SetExpenseRatioNet gets a reference to the given float64 and assigns it to the ExpenseRatioNet field.
func (o *InlineObject2Expenses) SetExpenseRatioNet(v float64) {
	o.ExpenseRatioNet = &v
}

func (o InlineObject2Expenses) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject2Expenses) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExpenseRatioGross) {
		toSerialize["expense_ratio_gross"] = o.ExpenseRatioGross
	}
	if !IsNil(o.ExpenseRatioNet) {
		toSerialize["expense_ratio_net"] = o.ExpenseRatioNet
	}
	return toSerialize, nil
}

type NullableInlineObject2Expenses struct {
	value *InlineObject2Expenses
	isSet bool
}

func (v NullableInlineObject2Expenses) Get() *InlineObject2Expenses {
	return v.value
}

func (v *NullableInlineObject2Expenses) Set(val *InlineObject2Expenses) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject2Expenses) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject2Expenses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject2Expenses(val *InlineObject2Expenses) *NullableInlineObject2Expenses {
	return &NullableInlineObject2Expenses{value: val, isSet: true}
}

func (v NullableInlineObject2Expenses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject2Expenses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
