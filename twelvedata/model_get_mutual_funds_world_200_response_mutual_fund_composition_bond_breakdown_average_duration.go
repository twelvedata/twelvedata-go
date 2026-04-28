/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageDuration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageDuration{}

// GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageDuration Average duration of bond holdings for the fund and its category
type GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageDuration struct {
	// Average duration of bond holding of a fund
	Fund *float64 `json:"fund,omitempty"`
	// Average duration of bond holding of funds in the same category
	Category *float64 `json:"category,omitempty"`
}

// NewGetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageDuration instantiates a new GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageDuration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageDuration() *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageDuration {
	this := GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageDuration{}
	return &this
}

// NewGetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageDurationWithDefaults instantiates a new GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageDuration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageDurationWithDefaults() *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageDuration {
	this := GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageDuration{}
	return &this
}

// GetFund returns the Fund field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageDuration) GetFund() float64 {
	if o == nil || IsNil(o.Fund) {
		var ret float64
		return ret
	}
	return *o.Fund
}

// GetFundOk returns a tuple with the Fund field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageDuration) GetFundOk() (*float64, bool) {
	if o == nil || IsNil(o.Fund) {
		return nil, false
	}
	return o.Fund, true
}

// HasFund returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageDuration) HasFund() bool {
	if o != nil && !IsNil(o.Fund) {
		return true
	}

	return false
}

// SetFund gets a reference to the given float64 and assigns it to the Fund field.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageDuration) SetFund(v float64) {
	o.Fund = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageDuration) GetCategory() float64 {
	if o == nil || IsNil(o.Category) {
		var ret float64
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageDuration) GetCategoryOk() (*float64, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageDuration) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given float64 and assigns it to the Category field.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageDuration) SetCategory(v float64) {
	o.Category = &v
}

func (o GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageDuration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageDuration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Fund) {
		toSerialize["fund"] = o.Fund
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	return toSerialize, nil
}

type NullableGetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageDuration struct {
	value *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageDuration
	isSet bool
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageDuration) Get() *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageDuration {
	return v.value
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageDuration) Set(val *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageDuration) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageDuration) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageDuration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageDuration(val *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageDuration) *NullableGetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageDuration {
	return &NullableGetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageDuration{value: val, isSet: true}
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageDuration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageDuration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
