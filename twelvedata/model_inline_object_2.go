/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the InlineObject2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject2{}

// InlineObject2 struct for InlineObject2
type InlineObject2 struct {
	Brokerages []string               `json:"brokerages,omitempty"`
	Expenses   *InlineObject2Expenses `json:"expenses,omitempty"`
	Minimums   *InlineObject2Minimums `json:"minimums,omitempty"`
	Pricing    *InlineObject2Pricing  `json:"pricing,omitempty"`
}

// NewInlineObject2 instantiates a new InlineObject2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject2() *InlineObject2 {
	this := InlineObject2{}
	return &this
}

// NewInlineObject2WithDefaults instantiates a new InlineObject2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject2WithDefaults() *InlineObject2 {
	this := InlineObject2{}
	return &this
}

// GetBrokerages returns the Brokerages field value if set, zero value otherwise.
func (o *InlineObject2) GetBrokerages() []string {
	if o == nil || IsNil(o.Brokerages) {
		var ret []string
		return ret
	}
	return o.Brokerages
}

// GetBrokeragesOk returns a tuple with the Brokerages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetBrokeragesOk() ([]string, bool) {
	if o == nil || IsNil(o.Brokerages) {
		return nil, false
	}
	return o.Brokerages, true
}

// HasBrokerages returns a boolean if a field has been set.
func (o *InlineObject2) HasBrokerages() bool {
	if o != nil && !IsNil(o.Brokerages) {
		return true
	}

	return false
}

// SetBrokerages gets a reference to the given []string and assigns it to the Brokerages field.
func (o *InlineObject2) SetBrokerages(v []string) {
	o.Brokerages = v
}

// GetExpenses returns the Expenses field value if set, zero value otherwise.
func (o *InlineObject2) GetExpenses() InlineObject2Expenses {
	if o == nil || IsNil(o.Expenses) {
		var ret InlineObject2Expenses
		return ret
	}
	return *o.Expenses
}

// GetExpensesOk returns a tuple with the Expenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetExpensesOk() (*InlineObject2Expenses, bool) {
	if o == nil || IsNil(o.Expenses) {
		return nil, false
	}
	return o.Expenses, true
}

// HasExpenses returns a boolean if a field has been set.
func (o *InlineObject2) HasExpenses() bool {
	if o != nil && !IsNil(o.Expenses) {
		return true
	}

	return false
}

// SetExpenses gets a reference to the given InlineObject2Expenses and assigns it to the Expenses field.
func (o *InlineObject2) SetExpenses(v InlineObject2Expenses) {
	o.Expenses = &v
}

// GetMinimums returns the Minimums field value if set, zero value otherwise.
func (o *InlineObject2) GetMinimums() InlineObject2Minimums {
	if o == nil || IsNil(o.Minimums) {
		var ret InlineObject2Minimums
		return ret
	}
	return *o.Minimums
}

// GetMinimumsOk returns a tuple with the Minimums field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetMinimumsOk() (*InlineObject2Minimums, bool) {
	if o == nil || IsNil(o.Minimums) {
		return nil, false
	}
	return o.Minimums, true
}

// HasMinimums returns a boolean if a field has been set.
func (o *InlineObject2) HasMinimums() bool {
	if o != nil && !IsNil(o.Minimums) {
		return true
	}

	return false
}

// SetMinimums gets a reference to the given InlineObject2Minimums and assigns it to the Minimums field.
func (o *InlineObject2) SetMinimums(v InlineObject2Minimums) {
	o.Minimums = &v
}

// GetPricing returns the Pricing field value if set, zero value otherwise.
func (o *InlineObject2) GetPricing() InlineObject2Pricing {
	if o == nil || IsNil(o.Pricing) {
		var ret InlineObject2Pricing
		return ret
	}
	return *o.Pricing
}

// GetPricingOk returns a tuple with the Pricing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetPricingOk() (*InlineObject2Pricing, bool) {
	if o == nil || IsNil(o.Pricing) {
		return nil, false
	}
	return o.Pricing, true
}

// HasPricing returns a boolean if a field has been set.
func (o *InlineObject2) HasPricing() bool {
	if o != nil && !IsNil(o.Pricing) {
		return true
	}

	return false
}

// SetPricing gets a reference to the given InlineObject2Pricing and assigns it to the Pricing field.
func (o *InlineObject2) SetPricing(v InlineObject2Pricing) {
	o.Pricing = &v
}

func (o InlineObject2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Brokerages) {
		toSerialize["brokerages"] = o.Brokerages
	}
	if !IsNil(o.Expenses) {
		toSerialize["expenses"] = o.Expenses
	}
	if !IsNil(o.Minimums) {
		toSerialize["minimums"] = o.Minimums
	}
	if !IsNil(o.Pricing) {
		toSerialize["pricing"] = o.Pricing
	}
	return toSerialize, nil
}

type NullableInlineObject2 struct {
	value *InlineObject2
	isSet bool
}

func (v NullableInlineObject2) Get() *InlineObject2 {
	return v.value
}

func (v *NullableInlineObject2) Set(val *InlineObject2) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject2) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject2(val *InlineObject2) *NullableInlineObject2 {
	return &NullableInlineObject2{value: val, isSet: true}
}

func (v NullableInlineObject2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
