/**
 * Twelve Data API client for Go
 *
 * NOTE: This code is auto generated, please do not edit it manually.
 */
package twelvedata

import (
	"encoding/json"
)

// checks if the GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageMaturity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageMaturity{}

// GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageMaturity Average maturity of bond holdings for the fund and its category
type GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageMaturity struct {
	// Average maturity of bond holding of a fund
	Fund *float64 `json:"fund,omitempty"`
	// Average maturity of bond holding of funds in the same category
	Category *float64 `json:"category,omitempty"`
}

// NewGetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageMaturity instantiates a new GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageMaturity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageMaturity() *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageMaturity {
	this := GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageMaturity{}
	return &this
}

// NewGetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageMaturityWithDefaults instantiates a new GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageMaturity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageMaturityWithDefaults() *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageMaturity {
	this := GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageMaturity{}
	return &this
}

// GetFund returns the Fund field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageMaturity) GetFund() float64 {
	if o == nil || IsNil(o.Fund) {
		var ret float64
		return ret
	}
	return *o.Fund
}

// GetFundOk returns a tuple with the Fund field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageMaturity) GetFundOk() (*float64, bool) {
	if o == nil || IsNil(o.Fund) {
		return nil, false
	}
	return o.Fund, true
}

// HasFund returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageMaturity) HasFund() bool {
	if o != nil && !IsNil(o.Fund) {
		return true
	}

	return false
}

// SetFund gets a reference to the given float64 and assigns it to the Fund field.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageMaturity) SetFund(v float64) {
	o.Fund = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageMaturity) GetCategory() float64 {
	if o == nil || IsNil(o.Category) {
		var ret float64
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageMaturity) GetCategoryOk() (*float64, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageMaturity) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given float64 and assigns it to the Category field.
func (o *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageMaturity) SetCategory(v float64) {
	o.Category = &v
}

func (o GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageMaturity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageMaturity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Fund) {
		toSerialize["fund"] = o.Fund
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	return toSerialize, nil
}

type NullableGetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageMaturity struct {
	value *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageMaturity
	isSet bool
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageMaturity) Get() *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageMaturity {
	return v.value
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageMaturity) Set(val *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageMaturity) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageMaturity) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageMaturity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageMaturity(val *GetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageMaturity) *NullableGetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageMaturity {
	return &NullableGetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageMaturity{value: val, isSet: true}
}

func (v NullableGetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageMaturity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMutualFundsWorld200ResponseMutualFundCompositionBondBreakdownAverageMaturity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
