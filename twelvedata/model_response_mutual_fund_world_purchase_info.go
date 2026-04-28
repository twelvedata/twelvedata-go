/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the ResponseMutualFundWorldPurchaseInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseMutualFundWorldPurchaseInfo{}

// ResponseMutualFundWorldPurchaseInfo Purchase information for the mutual fund
type ResponseMutualFundWorldPurchaseInfo struct {
	Expenses *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoExpenses `json:"expenses,omitempty"`
	Minimums *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums `json:"minimums,omitempty"`
	Pricing  *GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing  `json:"pricing,omitempty"`
	// List of brokerages where mutual fund can be purchased
	Brokerages []string `json:"brokerages,omitempty"`
}

// NewResponseMutualFundWorldPurchaseInfo instantiates a new ResponseMutualFundWorldPurchaseInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseMutualFundWorldPurchaseInfo() *ResponseMutualFundWorldPurchaseInfo {
	this := ResponseMutualFundWorldPurchaseInfo{}
	return &this
}

// NewResponseMutualFundWorldPurchaseInfoWithDefaults instantiates a new ResponseMutualFundWorldPurchaseInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseMutualFundWorldPurchaseInfoWithDefaults() *ResponseMutualFundWorldPurchaseInfo {
	this := ResponseMutualFundWorldPurchaseInfo{}
	return &this
}

// GetExpenses returns the Expenses field value if set, zero value otherwise.
func (o *ResponseMutualFundWorldPurchaseInfo) GetExpenses() GetMutualFundsWorld200ResponseMutualFundPurchaseInfoExpenses {
	if o == nil || IsNil(o.Expenses) {
		var ret GetMutualFundsWorld200ResponseMutualFundPurchaseInfoExpenses
		return ret
	}
	return *o.Expenses
}

// GetExpensesOk returns a tuple with the Expenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMutualFundWorldPurchaseInfo) GetExpensesOk() (*GetMutualFundsWorld200ResponseMutualFundPurchaseInfoExpenses, bool) {
	if o == nil || IsNil(o.Expenses) {
		return nil, false
	}
	return o.Expenses, true
}

// HasExpenses returns a boolean if a field has been set.
func (o *ResponseMutualFundWorldPurchaseInfo) HasExpenses() bool {
	if o != nil && !IsNil(o.Expenses) {
		return true
	}

	return false
}

// SetExpenses gets a reference to the given GetMutualFundsWorld200ResponseMutualFundPurchaseInfoExpenses and assigns it to the Expenses field.
func (o *ResponseMutualFundWorldPurchaseInfo) SetExpenses(v GetMutualFundsWorld200ResponseMutualFundPurchaseInfoExpenses) {
	o.Expenses = &v
}

// GetMinimums returns the Minimums field value if set, zero value otherwise.
func (o *ResponseMutualFundWorldPurchaseInfo) GetMinimums() GetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums {
	if o == nil || IsNil(o.Minimums) {
		var ret GetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums
		return ret
	}
	return *o.Minimums
}

// GetMinimumsOk returns a tuple with the Minimums field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMutualFundWorldPurchaseInfo) GetMinimumsOk() (*GetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums, bool) {
	if o == nil || IsNil(o.Minimums) {
		return nil, false
	}
	return o.Minimums, true
}

// HasMinimums returns a boolean if a field has been set.
func (o *ResponseMutualFundWorldPurchaseInfo) HasMinimums() bool {
	if o != nil && !IsNil(o.Minimums) {
		return true
	}

	return false
}

// SetMinimums gets a reference to the given GetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums and assigns it to the Minimums field.
func (o *ResponseMutualFundWorldPurchaseInfo) SetMinimums(v GetMutualFundsWorld200ResponseMutualFundPurchaseInfoMinimums) {
	o.Minimums = &v
}

// GetPricing returns the Pricing field value if set, zero value otherwise.
func (o *ResponseMutualFundWorldPurchaseInfo) GetPricing() GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing {
	if o == nil || IsNil(o.Pricing) {
		var ret GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing
		return ret
	}
	return *o.Pricing
}

// GetPricingOk returns a tuple with the Pricing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMutualFundWorldPurchaseInfo) GetPricingOk() (*GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing, bool) {
	if o == nil || IsNil(o.Pricing) {
		return nil, false
	}
	return o.Pricing, true
}

// HasPricing returns a boolean if a field has been set.
func (o *ResponseMutualFundWorldPurchaseInfo) HasPricing() bool {
	if o != nil && !IsNil(o.Pricing) {
		return true
	}

	return false
}

// SetPricing gets a reference to the given GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing and assigns it to the Pricing field.
func (o *ResponseMutualFundWorldPurchaseInfo) SetPricing(v GetMutualFundsWorld200ResponseMutualFundPurchaseInfoPricing) {
	o.Pricing = &v
}

// GetBrokerages returns the Brokerages field value if set, zero value otherwise.
func (o *ResponseMutualFundWorldPurchaseInfo) GetBrokerages() []string {
	if o == nil || IsNil(o.Brokerages) {
		var ret []string
		return ret
	}
	return o.Brokerages
}

// GetBrokeragesOk returns a tuple with the Brokerages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseMutualFundWorldPurchaseInfo) GetBrokeragesOk() ([]string, bool) {
	if o == nil || IsNil(o.Brokerages) {
		return nil, false
	}
	return o.Brokerages, true
}

// HasBrokerages returns a boolean if a field has been set.
func (o *ResponseMutualFundWorldPurchaseInfo) HasBrokerages() bool {
	if o != nil && !IsNil(o.Brokerages) {
		return true
	}

	return false
}

// SetBrokerages gets a reference to the given []string and assigns it to the Brokerages field.
func (o *ResponseMutualFundWorldPurchaseInfo) SetBrokerages(v []string) {
	o.Brokerages = v
}

func (o ResponseMutualFundWorldPurchaseInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseMutualFundWorldPurchaseInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Expenses) {
		toSerialize["expenses"] = o.Expenses
	}
	if !IsNil(o.Minimums) {
		toSerialize["minimums"] = o.Minimums
	}
	if !IsNil(o.Pricing) {
		toSerialize["pricing"] = o.Pricing
	}
	if !IsNil(o.Brokerages) {
		toSerialize["brokerages"] = o.Brokerages
	}
	return toSerialize, nil
}

type NullableResponseMutualFundWorldPurchaseInfo struct {
	value *ResponseMutualFundWorldPurchaseInfo
	isSet bool
}

func (v NullableResponseMutualFundWorldPurchaseInfo) Get() *ResponseMutualFundWorldPurchaseInfo {
	return v.value
}

func (v *NullableResponseMutualFundWorldPurchaseInfo) Set(val *ResponseMutualFundWorldPurchaseInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseMutualFundWorldPurchaseInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseMutualFundWorldPurchaseInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseMutualFundWorldPurchaseInfo(val *ResponseMutualFundWorldPurchaseInfo) *NullableResponseMutualFundWorldPurchaseInfo {
	return &NullableResponseMutualFundWorldPurchaseInfo{value: val, isSet: true}
}

func (v NullableResponseMutualFundWorldPurchaseInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseMutualFundWorldPurchaseInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
